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

package api

import (
	"reflect"
	"testing"

	"github.com/fsouza/go-dockerclient"
)

func TestPortBindingFromDockerPortBinding(t *testing.T) {
	pairs := []struct {
		dockerPortBindings map[docker.Port][]docker.PortBinding
		ecsPortBindings    []PortBinding
	}{
		{
			map[docker.Port][]docker.PortBinding{"53/udp": []docker.PortBinding{docker.PortBinding{"1.2.3.4", "55"}}},
			[]PortBinding{PortBinding{BindIp: "1.2.3.4", HostPort: 55, ContainerPort: 53, Protocol: TransportProtocolUDP}},
		},
		{
			map[docker.Port][]docker.PortBinding{"80/tcp": []docker.PortBinding{docker.PortBinding{"2.3.4.5", "8080"}, docker.PortBinding{"5.6.7.8", "80"}}},
			[]PortBinding{PortBinding{BindIp: "2.3.4.5", HostPort: 8080, ContainerPort: 80, Protocol: TransportProtocolTCP}, PortBinding{BindIp: "5.6.7.8", HostPort: 80, ContainerPort: 80, Protocol: TransportProtocolTCP}},
		},
	}

	for ndx, pair := range pairs {
		converted, err := PortBindingFromDockerPortBinding(pair.dockerPortBindings)
		if err != nil {
			t.Errorf("Error converting port binding pair #%v: %v", ndx, err)
		}
		if !reflect.DeepEqual(pair.ecsPortBindings, converted) {
			t.Errorf("Converted bindings didn't match expected for #%v: expected %+v, actual %+v", ndx, pair.ecsPortBindings, converted)
		}
	}
}

func TestPortBindingErrors(t *testing.T) {
	badInputs := []struct {
		dockerPortBindings map[docker.Port][]docker.PortBinding
		errorName          string
	}{
		{
			map[docker.Port][]docker.PortBinding{"woof/tcp": []docker.PortBinding{docker.PortBinding{"2.3.4.5", "8080"}, docker.PortBinding{"5.6.7.8", "80"}}},
			UnparseablePortErrorName,
		},
		{
			map[docker.Port][]docker.PortBinding{"80/tcp": []docker.PortBinding{docker.PortBinding{"2.3.4.5", "8080"}, docker.PortBinding{"5.6.7.8", "bark"}}},
			UnparseablePortErrorName,
		},
		{
			map[docker.Port][]docker.PortBinding{"80/tcp": []docker.PortBinding{docker.PortBinding{"2.3.4.5", "8080"}, docker.PortBinding{"5.6.7.8", "bark"}}},
			UnrecognizedTransportProtocolErrorName,
		},
	}

	for ndx, pair := range badInputs {
		_, err := PortBindingFromDockerPortBinding(pair.dockerPortBindings)
		if err == nil {
			t.Errorf("Expected error converting port binding pair #%v", ndx)
		}
	}
}
