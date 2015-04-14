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

package engine_test

import (
	"testing"

	"code.google.com/p/gomock/gomock"
	"github.com/aws/amazon-ecs-agent/agent/config"
	"github.com/aws/amazon-ecs-agent/agent/engine"
	"github.com/aws/amazon-ecs-agent/agent/engine/mocks"
)

func mocks(t *testing.T, cfg *config.Config) (*gomock.Controller, *mock_engine.MockDockerClient, engine.TaskEngine) {
	ctrl := gomock.NewController(t)
	client := mock_engine.NewMockDockerClient(ctrl)
	taskEngine := engine.NewTaskEngine(cfg)
	taskEngine.(*engine.DockerTaskEngine).SetDockerClient(client)
	return ctrl, client, taskEngine
}

func TestBatchContainerHappyPath(t *testing.T) {
	ctrl, client, taskEngine := mocks(t, &config.Config{})
	defer ctrl.Finish()

	eventStream := make(chan engine.DockerContainerChangeEvent)
	client.EXPECT().ContainerEvents().Return(eventStream, nil)

	err := taskEngine.Init()
	if err != nil {
		t.Fatal(err)
	}
	// TODO WIP
}
