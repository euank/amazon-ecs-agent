// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package wsclient

import (
	"crypto/tls"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cihub/seelog"
	"golang.org/x/net/proxy"
)

// websocketConn creates a connection that can be used for websocket communication.
// This exists to abstract away choosing a proxying connection if desired.
// Note, the 'timeout' argument is ignored if a proxy is used.
func websocketConn(host string, timeout time.Duration, acceptInvalidCert bool) (net.Conn, error) {
	if os.Getenv("all_proxy") != "" {
		// Don't just always go down the proxy path because this is
		// yet-another-dialer that doesn't take a timeout; prefer the default
		// timeoutable path whenever possible
		proxyDialer := proxy.FromEnvironment()
		seelog.Infof("Creating proxy websocket connection for %v", host)
		proxyConn, err := proxyDialer.Dial("tcp", host)
		if err != nil {
			return nil, err
		}
		colonPos := strings.LastIndex(host, ":")
		if colonPos == -1 {
			colonPos = len(host)
		}
		tlsProxyConn := tls.Client(proxyConn, &tls.Config{InsecureSkipVerify: acceptInvalidCert, ServerName: host[:colonPos]})
		return tlsProxyConn, nil
	} else {
		timeoutDialer := &net.Dialer{Timeout: wsConnectTimeout}
		seelog.Infof("Creating poll dialer for %v", host)
		return tls.DialWithDialer(timeoutDialer, "tcp", host, &tls.Config{InsecureSkipVerify: acceptInvalidCert})
	}
}
