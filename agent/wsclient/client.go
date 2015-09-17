// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package wsclient wraps the generated aws-sdk-go client to provide marshalling
// and unmarshalling of data over a websocket connection in the format expected
// by backend. It allows for bidirectional communication and acts as both a
// client-and-server in terms of requests, but only as a client in terms of
// connecting.
package wsclient

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/aws/amazon-ecs-agent/agent/logger"
	"github.com/aws/amazon-ecs-agent/agent/utils"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/internal/protocol/json/jsonutil"
	"github.com/gorilla/websocket"
)

var log = logger.ForModule("ws client")

const (
	// ServiceName defines the service name for the agent. This is used to sign messages
	// that are sent to the backend.
	ServiceName = "ecs"

	// wsConnectTimeout specifies the default connection timeout to the backend.
	wsConnectTimeout = 3 * time.Second

	// readBufSize is the size of the read buffer for the ws connection.
	readBufSize = 4096

	// writeBufSize is the size of the write buffer for the ws connection.
	writeBufSize = 32768
)

// ReceivedMessage is the intermediate message used to unmarshal a
// message from backend
type ReceivedMessage struct {
	Type    string          `json:"type"`
	Message json.RawMessage `json:"message"`
}

// RequestMessage is the intermediate message marshalled to send to backend.
type RequestMessage struct {
	Type    string          `json:"type"`
	Message json.RawMessage `json:"message"`
}

// WebsocketConn specifies the subset of gorilla/websocket's
// connection's methods that this client uses.
type WebsocketConn interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, data []byte, err error)
	Close() error
}

// RequestHandler would be func(*ecsacs.T for T in ecsacs.*) to be more proper, but it needs
// to be interface{} to properly capture that
type RequestHandler interface{}

// ClientServer is a combined client and server for the backend websocket connection
type ClientServer interface {
	AddRequestHandler(RequestHandler)
	// SetAnyRequestHandler takes a function with the signature 'func(i
	// interface{})' and calls it with every message the server passes down.
	// Only a single 'AnyRequestHandler' will be active at a given time for a
	// ClientServer
	SetAnyRequestHandler(RequestHandler)
	MakeRequest(input interface{}) error
	Connect() error
	Serve() error
	io.Closer
}

//go:generate mockgen.sh github.com/aws/amazon-ecs-agent/agent/wsclient ClientServer mock/$GOFILE

// ClientServerImpl wraps commonly used methods defined in ClientServer interface.
type ClientServerImpl struct {
	AcceptInvalidCert  bool
	Conn               WebsocketConn
	CredentialProvider *credentials.Credentials
	Region             string
	// RequestHandlers is a map from message types to handler functions of the
	// form:
	//     "FooMessage": func(message *ecsacs.FooMessage)
	RequestHandlers map[string]RequestHandler
	// AnyRequestHandler is a request handler that, if set, is called on every
	// message with said message. It will be called before a RequestHandler is
	// called. It must take a single interface{} argument.
	AnyRequestHandler RequestHandler
	// URL is the full url to the backend, including path, querystring, and so on.
	URL string
	ClientServer
	ServiceError
	TypeDecoder
}

// Connect opens a connection to the backend and upgrades it to a websocket. Calls to
// 'MakeRequest' can be made after calling this, but responss will not be
// receivable until 'Serve' is also called.
func (cs *ClientServerImpl) Connect() error {
	parsedURL, err := url.Parse(cs.URL)
	if err != nil {
		return err
	}

	// NewRequest never returns an error if the url parses and we just verified
	// it did above
	request, _ := http.NewRequest("GET", cs.URL, nil)
	// Sign the request; we'll send its headers via the websocket client which includes the signature
	utils.SignHTTPRequest(request, cs.Region, ServiceName, cs.CredentialProvider, nil)

	// url.Host might not have the port, but tls.Dial needs it
	dialHost := parsedURL.Host
	if !strings.Contains(dialHost, ":") {
		dialHost += ":443"
	}

	wsConn, err := websocketConn(dialHost, wsConnectTimeout, cs.AcceptInvalidCert)
	if err != nil {
		return err
	}

	websocketConn, httpResponse, err := websocket.NewClient(wsConn, parsedURL, request.Header, readBufSize, writeBufSize)
	if httpResponse != nil {
		defer httpResponse.Body.Close()
	}
	if err != nil {
		defer wsConn.Close()
		var resp []byte
		if httpResponse != nil {
			var readErr error
			resp, readErr = ioutil.ReadAll(httpResponse.Body)
			if readErr != nil {
				return errors.New("Unable to read websocket connection: " + readErr.Error() + ", " + err.Error())
			}
			// If there's a response, we can try to unmarshal it into one of the
			// modeled error types
			possibleError, _, decodeErr := DecodeData(resp, cs.TypeDecoder)
			if decodeErr == nil {
				return cs.NewError(possibleError)
			}
		}
		log.Warn("Error creating a websocket client", "err", err)
		return errors.New(string(resp) + ", " + err.Error())
	}
	cs.Conn = websocketConn
	return nil
}

// AddRequestHandler adds a request handler to this client.
// A request handler *must* be a function taking a single argument, and that
// argument *must* be a pointer to a recognized 'ecsacs' struct.
// E.g. if you desired to handle messages from acs of type 'FooMessage', you
// would pass the following handler in:
//     func(message *ecsacs.FooMessage)
// This function will panic if the passed in function does not have one pointer
// argument or the argument is not a recognized type.
// Additionally, the request handler will block processing of further messages
// on this connection so it's important that it return quickly.
func (cs *ClientServerImpl) AddRequestHandler(f RequestHandler) {
	firstArg := reflect.TypeOf(f).In(0)
	firstArgTypeStr := firstArg.Elem().Name()
	recognizedTypes := cs.GetRecognizedTypes()
	_, ok := recognizedTypes[firstArgTypeStr]
	if !ok {
		panic("AddRequestHandler called with invalid function; argument type not recognized: " + firstArgTypeStr)
	}
	cs.RequestHandlers[firstArgTypeStr] = f
}

func (cs *ClientServerImpl) SetAnyRequestHandler(f RequestHandler) {
	cs.AnyRequestHandler = f
}

// MakeRequest makes a request using the given input. Note, the input *MUST* be
// a pointer to a valid backend type that this client recognises
func (cs *ClientServerImpl) MakeRequest(input interface{}) error {
	send, err := cs.CreateRequestMessage(input)
	if err != nil {
		return err
	}

	// Over the wire we send something like
	// {"type":"AckRequest","message":{"messageId":"xyz"}}
	return cs.Conn.WriteMessage(websocket.TextMessage, send)
}

// ConsumeMessages reads messages from the websocket connection and handles read
// messages from an active connection.
func (cs *ClientServerImpl) ConsumeMessages() error {
	var err error
	for {
		messageType, message, cerr := cs.Conn.ReadMessage()
		err = cerr
		if err != nil {
			if err != io.EOF {
				if message != nil {
					log.Error("Error getting message from ws backend", "err", err, "message", message)
				} else {
					log.Error("Error getting message from ws backend", "err", err)
				}
			}
			break
		}
		if messageType != websocket.TextMessage {
			log.Error("Unexpected messageType", "type", messageType)
			// maybe not fatal though, we'll try to process it anyways
		}
		log.Debug("Got a message from websocket", "message", string(message[:]))
		cs.handleMessage(message)
	}
	return err
}

// CreateRequestMessage creates the request json message using the given input.
// Note, the input *MUST* be a pointer to a valid backend type that this
// client recognises.
func (cs *ClientServerImpl) CreateRequestMessage(input interface{}) ([]byte, error) {
	msg := &RequestMessage{}

	recognizedTypes := cs.GetRecognizedTypes()
	for typeStr, typeVal := range recognizedTypes {
		if reflect.TypeOf(input) == reflect.PtrTo(typeVal) {
			msg.Type = typeStr
			break
		}
	}
	if msg.Type == "" {
		return nil, &UnrecognizedWSRequestType{reflect.TypeOf(input).String()}
	}
	messageData, err := jsonutil.BuildJSON(input)
	if err != nil {
		return nil, &NotMarshallableWSRequest{msg.Type, err}
	}
	msg.Message = json.RawMessage(messageData)

	send, err := json.Marshal(msg)
	if err != nil {
		return nil, &NotMarshallableWSRequest{msg.Type, err}
	}
	return send, nil
}

// handleMessage dispatches a message to the correct 'requestHandler' for its
// type. If no request handler is found, the message is discarded.
func (cs *ClientServerImpl) handleMessage(data []byte) {
	typedMessage, typeStr, err := DecodeData(data, cs.TypeDecoder)
	if err != nil {
		log.Warn("Unable to handle message from backend", "err", err)
		return
	}

	if cs.AnyRequestHandler != nil {
		reflect.ValueOf(cs.AnyRequestHandler).Call([]reflect.Value{reflect.ValueOf(typedMessage)})
	}

	if handler, ok := cs.RequestHandlers[typeStr]; ok {
		reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(typedMessage)})
	} else {
		log.Info("No handler for message type", "type", typeStr)
	}
}
