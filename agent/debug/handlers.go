// +build debug

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

package debug

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/aws/amazon-ecs-agent/agent/sighandlers/exitcodes"
	"github.com/cihub/seelog"
)

func logDebugHandler(w http.ResponseWriter, req *http.Request) {
	for i := 0; i < 1000; i++ {
		seelog.Criticalf("I'm logging some junk: %v\n", i)
	}
}

func init() {
	seelog.Info("running pprof debug handler on :6060")
	go func() {
		http.HandleFunc("/log", logDebugHandler)
		seelog.Critical(http.ListenAndServe(":6060", nil))
		os.Exit(exitcodes.ExitTerminal)
	}()
}
