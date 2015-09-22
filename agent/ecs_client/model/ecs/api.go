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

package ecs

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opCreateCluster = "CreateCluster"

// CreateClusterRequest generates a request for the CreateCluster operation.
func (c *ECS) CreateClusterRequest(input *CreateClusterInput) (req *request.Request, output *CreateClusterOutput) {
	op := &request.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateClusterOutput{}
	req.Data = output
	return
}

// Creates a new Amazon ECS cluster. By default, your account will receive a
// default cluster when you launch your first container instance. However, you
// can create your own cluster with a unique name with the CreateCluster action.
func (c *ECS) CreateCluster(input *CreateClusterInput) (*CreateClusterOutput, error) {
	req, out := c.CreateClusterRequest(input)
	err := req.Send()
	return out, err
}

const opCreateService = "CreateService"

// CreateServiceRequest generates a request for the CreateService operation.
func (c *ECS) CreateServiceRequest(input *CreateServiceInput) (req *request.Request, output *CreateServiceOutput) {
	op := &request.Operation{
		Name:       opCreateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateServiceOutput{}
	req.Data = output
	return
}

// Runs and maintains a desired number of tasks from a specified task definition.
// If the number of tasks running in a service drops below desiredCount, Amazon
// ECS will spawn another instantiation of the task in the specified cluster.
func (c *ECS) CreateService(input *CreateServiceInput) (*CreateServiceOutput, error) {
	req, out := c.CreateServiceRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteCluster = "DeleteCluster"

// DeleteClusterRequest generates a request for the DeleteCluster operation.
func (c *ECS) DeleteClusterRequest(input *DeleteClusterInput) (req *request.Request, output *DeleteClusterOutput) {
	op := &request.Operation{
		Name:       opDeleteCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteClusterInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteClusterOutput{}
	req.Data = output
	return
}

// Deletes the specified cluster. You must deregister all container instances
// from this cluster before you may delete it. You can list the container instances
// in a cluster with ListContainerInstances and deregister them with DeregisterContainerInstance.
func (c *ECS) DeleteCluster(input *DeleteClusterInput) (*DeleteClusterOutput, error) {
	req, out := c.DeleteClusterRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteService = "DeleteService"

// DeleteServiceRequest generates a request for the DeleteService operation.
func (c *ECS) DeleteServiceRequest(input *DeleteServiceInput) (req *request.Request, output *DeleteServiceOutput) {
	op := &request.Operation{
		Name:       opDeleteService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteServiceOutput{}
	req.Data = output
	return
}

// Deletes a specified service within a cluster.
func (c *ECS) DeleteService(input *DeleteServiceInput) (*DeleteServiceOutput, error) {
	req, out := c.DeleteServiceRequest(input)
	err := req.Send()
	return out, err
}

const opDeregisterContainerInstance = "DeregisterContainerInstance"

// DeregisterContainerInstanceRequest generates a request for the DeregisterContainerInstance operation.
func (c *ECS) DeregisterContainerInstanceRequest(input *DeregisterContainerInstanceInput) (req *request.Request, output *DeregisterContainerInstanceOutput) {
	op := &request.Operation{
		Name:       opDeregisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterContainerInstanceInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeregisterContainerInstanceOutput{}
	req.Data = output
	return
}

// Deregisters an Amazon ECS container instance from the specified cluster.
// This instance will no longer be available to run tasks.
//
// If you intend to use the container instance for some other purpose after
// deregistration, you should stop all of the tasks running on the container
// instance before deregistration to avoid any orphaned tasks from consuming
// resources.
//
// Deregistering a container instance removes the instance from a cluster,
// but it does not terminate the EC2 instance; if you are finished using the
// instance, be sure to terminate it in the Amazon EC2 console to stop billing.
//
// When you terminate a container instance, it is automatically deregistered
// from your cluster.
func (c *ECS) DeregisterContainerInstance(input *DeregisterContainerInstanceInput) (*DeregisterContainerInstanceOutput, error) {
	req, out := c.DeregisterContainerInstanceRequest(input)
	err := req.Send()
	return out, err
}

const opDeregisterTaskDefinition = "DeregisterTaskDefinition"

// DeregisterTaskDefinitionRequest generates a request for the DeregisterTaskDefinition operation.
func (c *ECS) DeregisterTaskDefinitionRequest(input *DeregisterTaskDefinitionInput) (req *request.Request, output *DeregisterTaskDefinitionOutput) {
	op := &request.Operation{
		Name:       opDeregisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterTaskDefinitionInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeregisterTaskDefinitionOutput{}
	req.Data = output
	return
}

// Deregisters the specified task definition by family and revision. Upon deregistration,
// the task definition is marked as INACTIVE. Existing tasks and services that
// reference an INACTIVE task definition continue to run without disruption.
// Existing services that reference an INACTIVE task definition can still scale
// up or down by modifying the service's desired count.
//
// You cannot use an INACTIVE task definition to run new tasks or create new
// services, and you cannot update an existing service to reference an INACTIVE
// task definition (although there may be up to a 10 minute window following
// deregistration where these restrictions have not yet taken effect).
func (c *ECS) DeregisterTaskDefinition(input *DeregisterTaskDefinitionInput) (*DeregisterTaskDefinitionOutput, error) {
	req, out := c.DeregisterTaskDefinitionRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeClusters = "DescribeClusters"

// DescribeClustersRequest generates a request for the DescribeClusters operation.
func (c *ECS) DescribeClustersRequest(input *DescribeClustersInput) (req *request.Request, output *DescribeClustersOutput) {
	op := &request.Operation{
		Name:       opDescribeClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeClustersInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeClustersOutput{}
	req.Data = output
	return
}

// Describes one or more of your clusters.
func (c *ECS) DescribeClusters(input *DescribeClustersInput) (*DescribeClustersOutput, error) {
	req, out := c.DescribeClustersRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeContainerInstances = "DescribeContainerInstances"

// DescribeContainerInstancesRequest generates a request for the DescribeContainerInstances operation.
func (c *ECS) DescribeContainerInstancesRequest(input *DescribeContainerInstancesInput) (req *request.Request, output *DescribeContainerInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeContainerInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeContainerInstancesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeContainerInstancesOutput{}
	req.Data = output
	return
}

// Describes Amazon EC2 Container Service container instances. Returns metadata
// about registered and remaining resources on each container instance requested.
func (c *ECS) DescribeContainerInstances(input *DescribeContainerInstancesInput) (*DescribeContainerInstancesOutput, error) {
	req, out := c.DescribeContainerInstancesRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeServices = "DescribeServices"

// DescribeServicesRequest generates a request for the DescribeServices operation.
func (c *ECS) DescribeServicesRequest(input *DescribeServicesInput) (req *request.Request, output *DescribeServicesOutput) {
	op := &request.Operation{
		Name:       opDescribeServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServicesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeServicesOutput{}
	req.Data = output
	return
}

// Describes the specified services running in your cluster.
func (c *ECS) DescribeServices(input *DescribeServicesInput) (*DescribeServicesOutput, error) {
	req, out := c.DescribeServicesRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeTaskDefinition = "DescribeTaskDefinition"

// DescribeTaskDefinitionRequest generates a request for the DescribeTaskDefinition operation.
func (c *ECS) DescribeTaskDefinitionRequest(input *DescribeTaskDefinitionInput) (req *request.Request, output *DescribeTaskDefinitionOutput) {
	op := &request.Operation{
		Name:       opDescribeTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTaskDefinitionInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeTaskDefinitionOutput{}
	req.Data = output
	return
}

// Describes a task definition. You can specify a family and revision to find
// information on a specific task definition, or you can simply specify the
// family to find the latest ACTIVE revision in that family.
//
//  You can only describe INACTIVE task definitions while an active task or
// service references them.
func (c *ECS) DescribeTaskDefinition(input *DescribeTaskDefinitionInput) (*DescribeTaskDefinitionOutput, error) {
	req, out := c.DescribeTaskDefinitionRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeTasks = "DescribeTasks"

// DescribeTasksRequest generates a request for the DescribeTasks operation.
func (c *ECS) DescribeTasksRequest(input *DescribeTasksInput) (req *request.Request, output *DescribeTasksOutput) {
	op := &request.Operation{
		Name:       opDescribeTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTasksInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeTasksOutput{}
	req.Data = output
	return
}

// Describes a specified task or tasks.
func (c *ECS) DescribeTasks(input *DescribeTasksInput) (*DescribeTasksOutput, error) {
	req, out := c.DescribeTasksRequest(input)
	err := req.Send()
	return out, err
}

const opDiscoverPollEndpoint = "DiscoverPollEndpoint"

// DiscoverPollEndpointRequest generates a request for the DiscoverPollEndpoint operation.
func (c *ECS) DiscoverPollEndpointRequest(input *DiscoverPollEndpointInput) (req *request.Request, output *DiscoverPollEndpointOutput) {
	op := &request.Operation{
		Name:       opDiscoverPollEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DiscoverPollEndpointInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DiscoverPollEndpointOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Returns an endpoint for the Amazon EC2 Container Service agent to poll for
// updates.
func (c *ECS) DiscoverPollEndpoint(input *DiscoverPollEndpointInput) (*DiscoverPollEndpointOutput, error) {
	req, out := c.DiscoverPollEndpointRequest(input)
	err := req.Send()
	return out, err
}

const opListClusters = "ListClusters"

// ListClustersRequest generates a request for the ListClusters operation.
func (c *ECS) ListClustersRequest(input *ListClustersInput) (req *request.Request, output *ListClustersOutput) {
	op := &request.Operation{
		Name:       opListClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListClustersOutput{}
	req.Data = output
	return
}

// Returns a list of existing clusters.
func (c *ECS) ListClusters(input *ListClustersInput) (*ListClustersOutput, error) {
	req, out := c.ListClustersRequest(input)
	err := req.Send()
	return out, err
}

func (c *ECS) ListClustersPages(input *ListClustersInput, fn func(p *ListClustersOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListClustersRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListClustersOutput), lastPage)
	})
}

const opListContainerInstances = "ListContainerInstances"

// ListContainerInstancesRequest generates a request for the ListContainerInstances operation.
func (c *ECS) ListContainerInstancesRequest(input *ListContainerInstancesInput) (req *request.Request, output *ListContainerInstancesOutput) {
	op := &request.Operation{
		Name:       opListContainerInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListContainerInstancesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListContainerInstancesOutput{}
	req.Data = output
	return
}

// Returns a list of container instances in a specified cluster.
func (c *ECS) ListContainerInstances(input *ListContainerInstancesInput) (*ListContainerInstancesOutput, error) {
	req, out := c.ListContainerInstancesRequest(input)
	err := req.Send()
	return out, err
}

func (c *ECS) ListContainerInstancesPages(input *ListContainerInstancesInput, fn func(p *ListContainerInstancesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListContainerInstancesRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListContainerInstancesOutput), lastPage)
	})
}

const opListServices = "ListServices"

// ListServicesRequest generates a request for the ListServices operation.
func (c *ECS) ListServicesRequest(input *ListServicesInput) (req *request.Request, output *ListServicesOutput) {
	op := &request.Operation{
		Name:       opListServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListServicesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListServicesOutput{}
	req.Data = output
	return
}

// Lists the services that are running in a specified cluster.
func (c *ECS) ListServices(input *ListServicesInput) (*ListServicesOutput, error) {
	req, out := c.ListServicesRequest(input)
	err := req.Send()
	return out, err
}

func (c *ECS) ListServicesPages(input *ListServicesInput, fn func(p *ListServicesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListServicesRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListServicesOutput), lastPage)
	})
}

const opListTaskDefinitionFamilies = "ListTaskDefinitionFamilies"

// ListTaskDefinitionFamiliesRequest generates a request for the ListTaskDefinitionFamilies operation.
func (c *ECS) ListTaskDefinitionFamiliesRequest(input *ListTaskDefinitionFamiliesInput) (req *request.Request, output *ListTaskDefinitionFamiliesOutput) {
	op := &request.Operation{
		Name:       opListTaskDefinitionFamilies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTaskDefinitionFamiliesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListTaskDefinitionFamiliesOutput{}
	req.Data = output
	return
}

// Returns a list of task definition families that are registered to your account
// (which may include task definition families that no longer have any ACTIVE
// task definitions). You can filter the results with the familyPrefix parameter.
func (c *ECS) ListTaskDefinitionFamilies(input *ListTaskDefinitionFamiliesInput) (*ListTaskDefinitionFamiliesOutput, error) {
	req, out := c.ListTaskDefinitionFamiliesRequest(input)
	err := req.Send()
	return out, err
}

func (c *ECS) ListTaskDefinitionFamiliesPages(input *ListTaskDefinitionFamiliesInput, fn func(p *ListTaskDefinitionFamiliesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListTaskDefinitionFamiliesRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListTaskDefinitionFamiliesOutput), lastPage)
	})
}

const opListTaskDefinitions = "ListTaskDefinitions"

// ListTaskDefinitionsRequest generates a request for the ListTaskDefinitions operation.
func (c *ECS) ListTaskDefinitionsRequest(input *ListTaskDefinitionsInput) (req *request.Request, output *ListTaskDefinitionsOutput) {
	op := &request.Operation{
		Name:       opListTaskDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTaskDefinitionsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListTaskDefinitionsOutput{}
	req.Data = output
	return
}

// Returns a list of task definitions that are registered to your account. You
// can filter the results by family name with the familyPrefix parameter or
// by status with the status parameter.
func (c *ECS) ListTaskDefinitions(input *ListTaskDefinitionsInput) (*ListTaskDefinitionsOutput, error) {
	req, out := c.ListTaskDefinitionsRequest(input)
	err := req.Send()
	return out, err
}

func (c *ECS) ListTaskDefinitionsPages(input *ListTaskDefinitionsInput, fn func(p *ListTaskDefinitionsOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListTaskDefinitionsRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListTaskDefinitionsOutput), lastPage)
	})
}

const opListTasks = "ListTasks"

// ListTasksRequest generates a request for the ListTasks operation.
func (c *ECS) ListTasksRequest(input *ListTasksInput) (req *request.Request, output *ListTasksOutput) {
	op := &request.Operation{
		Name:       opListTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTasksInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListTasksOutput{}
	req.Data = output
	return
}

// Returns a list of tasks for a specified cluster. You can filter the results
// by family name, by a particular container instance, or by the desired status
// of the task with the family, containerInstance, and desiredStatus parameters.
func (c *ECS) ListTasks(input *ListTasksInput) (*ListTasksOutput, error) {
	req, out := c.ListTasksRequest(input)
	err := req.Send()
	return out, err
}

func (c *ECS) ListTasksPages(input *ListTasksInput, fn func(p *ListTasksOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListTasksRequest(input)
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListTasksOutput), lastPage)
	})
}

const opRegisterContainerInstance = "RegisterContainerInstance"

// RegisterContainerInstanceRequest generates a request for the RegisterContainerInstance operation.
func (c *ECS) RegisterContainerInstanceRequest(input *RegisterContainerInstanceInput) (req *request.Request, output *RegisterContainerInstanceOutput) {
	op := &request.Operation{
		Name:       opRegisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterContainerInstanceInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RegisterContainerInstanceOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Registers an Amazon EC2 instance into the specified cluster. This instance
// will become available to place containers on.
func (c *ECS) RegisterContainerInstance(input *RegisterContainerInstanceInput) (*RegisterContainerInstanceOutput, error) {
	req, out := c.RegisterContainerInstanceRequest(input)
	err := req.Send()
	return out, err
}

const opRegisterTaskDefinition = "RegisterTaskDefinition"

// RegisterTaskDefinitionRequest generates a request for the RegisterTaskDefinition operation.
func (c *ECS) RegisterTaskDefinitionRequest(input *RegisterTaskDefinitionInput) (req *request.Request, output *RegisterTaskDefinitionOutput) {
	op := &request.Operation{
		Name:       opRegisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterTaskDefinitionInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RegisterTaskDefinitionOutput{}
	req.Data = output
	return
}

// Registers a new task definition from the supplied family and containerDefinitions.
// Optionally, you can add data volumes to your containers with the volumes
// parameter. For more information on task definition parameters and defaults,
// see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
// in the Amazon EC2 Container Service Developer Guide.
func (c *ECS) RegisterTaskDefinition(input *RegisterTaskDefinitionInput) (*RegisterTaskDefinitionOutput, error) {
	req, out := c.RegisterTaskDefinitionRequest(input)
	err := req.Send()
	return out, err
}

const opRunTask = "RunTask"

// RunTaskRequest generates a request for the RunTask operation.
func (c *ECS) RunTaskRequest(input *RunTaskInput) (req *request.Request, output *RunTaskOutput) {
	op := &request.Operation{
		Name:       opRunTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RunTaskInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RunTaskOutput{}
	req.Data = output
	return
}

// Start a task using random placement and the default Amazon ECS scheduler.
// If you want to use your own scheduler or place a task on a specific container
// instance, use StartTask instead.
//
//  The count parameter is limited to 10 tasks per call.
func (c *ECS) RunTask(input *RunTaskInput) (*RunTaskOutput, error) {
	req, out := c.RunTaskRequest(input)
	err := req.Send()
	return out, err
}

const opStartTask = "StartTask"

// StartTaskRequest generates a request for the StartTask operation.
func (c *ECS) StartTaskRequest(input *StartTaskInput) (req *request.Request, output *StartTaskOutput) {
	op := &request.Operation{
		Name:       opStartTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTaskInput{}
	}

	req = c.newRequest(op, input, output)
	output = &StartTaskOutput{}
	req.Data = output
	return
}

// Starts a new task from the specified task definition on the specified container
// instance or instances. If you want to use the default Amazon ECS scheduler
// to place your task, use RunTask instead.
//
//  The list of container instances to start tasks on is limited to 10.
func (c *ECS) StartTask(input *StartTaskInput) (*StartTaskOutput, error) {
	req, out := c.StartTaskRequest(input)
	err := req.Send()
	return out, err
}

const opStopTask = "StopTask"

// StopTaskRequest generates a request for the StopTask operation.
func (c *ECS) StopTaskRequest(input *StopTaskInput) (req *request.Request, output *StopTaskOutput) {
	op := &request.Operation{
		Name:       opStopTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopTaskInput{}
	}

	req = c.newRequest(op, input, output)
	output = &StopTaskOutput{}
	req.Data = output
	return
}

// Stops a running task.
//
// When StopTask is called on a task, the equivalent of docker stop is issued
// to the containers running in the task. This results in a SIGTERM and a 30-second
// timeout, after which SIGKILL is sent and the containers are forcibly stopped.
// If the container handles the SIGTERM gracefully and exits within 30 seconds
// from receiving it, no SIGKILL is sent.
func (c *ECS) StopTask(input *StopTaskInput) (*StopTaskOutput, error) {
	req, out := c.StopTaskRequest(input)
	err := req.Send()
	return out, err
}

const opSubmitContainerStateChange = "SubmitContainerStateChange"

// SubmitContainerStateChangeRequest generates a request for the SubmitContainerStateChange operation.
func (c *ECS) SubmitContainerStateChangeRequest(input *SubmitContainerStateChangeInput) (req *request.Request, output *SubmitContainerStateChangeOutput) {
	op := &request.Operation{
		Name:       opSubmitContainerStateChange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitContainerStateChangeInput{}
	}

	req = c.newRequest(op, input, output)
	output = &SubmitContainerStateChangeOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Sent to acknowledge that a container changed states.
func (c *ECS) SubmitContainerStateChange(input *SubmitContainerStateChangeInput) (*SubmitContainerStateChangeOutput, error) {
	req, out := c.SubmitContainerStateChangeRequest(input)
	err := req.Send()
	return out, err
}

const opSubmitTaskStateChange = "SubmitTaskStateChange"

// SubmitTaskStateChangeRequest generates a request for the SubmitTaskStateChange operation.
func (c *ECS) SubmitTaskStateChangeRequest(input *SubmitTaskStateChangeInput) (req *request.Request, output *SubmitTaskStateChangeOutput) {
	op := &request.Operation{
		Name:       opSubmitTaskStateChange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitTaskStateChangeInput{}
	}

	req = c.newRequest(op, input, output)
	output = &SubmitTaskStateChangeOutput{}
	req.Data = output
	return
}

// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Sent to acknowledge that a task changed states.
func (c *ECS) SubmitTaskStateChange(input *SubmitTaskStateChangeInput) (*SubmitTaskStateChangeOutput, error) {
	req, out := c.SubmitTaskStateChangeRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateContainerAgent = "UpdateContainerAgent"

// UpdateContainerAgentRequest generates a request for the UpdateContainerAgent operation.
func (c *ECS) UpdateContainerAgentRequest(input *UpdateContainerAgentInput) (req *request.Request, output *UpdateContainerAgentOutput) {
	op := &request.Operation{
		Name:       opUpdateContainerAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateContainerAgentInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateContainerAgentOutput{}
	req.Data = output
	return
}

// Updates the Amazon ECS container agent on a specified container instance.
// Updating the Amazon ECS container agent does not interrupt running tasks
// or services on the container instance. The process for updating the agent
// differs depending on whether your container instance was launched with the
// Amazon ECS-optimized AMI or another operating system.
//
// UpdateContainerAgent requires the Amazon ECS-optimized AMI or Amazon Linux
// with the ecs-init service installed and running. For help updating the Amazon
// ECS container agent on other operating systems, see Manually Updating the
// Amazon ECS Container Agent (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html#manually_update_agent)
// in the Amazon EC2 Container Service Developer Guide.
func (c *ECS) UpdateContainerAgent(input *UpdateContainerAgentInput) (*UpdateContainerAgentOutput, error) {
	req, out := c.UpdateContainerAgentRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateService = "UpdateService"

// UpdateServiceRequest generates a request for the UpdateService operation.
func (c *ECS) UpdateServiceRequest(input *UpdateServiceInput) (req *request.Request, output *UpdateServiceOutput) {
	op := &request.Operation{
		Name:       opUpdateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateServiceOutput{}
	req.Data = output
	return
}

// Modify the desired count or task definition used in a service.
//
// You can add to or subtract from the number of instantiations of a task definition
// in a service by specifying the cluster that the service is running in and
// a new desiredCount parameter.
//
// You can use UpdateService to modify your task definition and deploy a new
// version of your service, one task at a time. If you modify the task definition
// with UpdateService, Amazon ECS spawns a task with the new version of the
// task definition and then stops an old task after the new version is running.
// Because UpdateService starts a new version of the task before stopping an
// old version, your cluster must have capacity to support one more instantiation
// of the task when UpdateService is run. If your cluster cannot support another
// instantiation of the task used in your service, you can reduce the desired
// count of your service by one before modifying the task definition.
//
// When UpdateService replaces a task during an update, the equivalent of docker
// stop is issued to the containers running in the task. This results in a SIGTERM
// and a 30-second timeout, after which SIGKILL is sent and the containers are
// forcibly stopped. If the container handles the SIGTERM gracefully and exits
// within 30 seconds from receiving it, no SIGKILL is sent.
func (c *ECS) UpdateService(input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	req, out := c.UpdateServiceRequest(input)
	err := req.Send()
	return out, err
}

type Attribute struct {
	Name *string `locationName:"name" type:"string" required:"true"`

	Value *string `locationName:"value" type:"string"`

	metadataAttribute `json:"-" xml:"-"`
}

type metadataAttribute struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Attribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Attribute) GoString() string {
	return s.String()
}

// A regional grouping of one or more container instances on which you can run
// task requests. Each account receives a default cluster the first time you
// use the Amazon ECS service, but you may also create other clusters. Clusters
// may contain more than one instance type simultaneously.
type Cluster struct {
	// The number of services that are running on the cluster in an ACTIVE state.
	// You can view these services with ListServices.
	ActiveServicesCount *int64 `locationName:"activeServicesCount" type:"integer"`

	// The Amazon Resource Name (ARN) that identifies the cluster. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the cluster, the AWS
	// account ID of the cluster owner, the cluster namespace, and then the cluster
	// name. For example, arn:aws:ecs:region:012345678910:cluster/test.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// A user-generated string that you can use to identify your cluster.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The number of tasks in the cluster that are in the PENDING state.
	PendingTasksCount *int64 `locationName:"pendingTasksCount" type:"integer"`

	// The number of container instances registered into the cluster.
	RegisteredContainerInstancesCount *int64 `locationName:"registeredContainerInstancesCount" type:"integer"`

	// The number of tasks in the cluster that are in the RUNNING state.
	RunningTasksCount *int64 `locationName:"runningTasksCount" type:"integer"`

	// The status of the cluster. The valid values are ACTIVE or INACTIVE. ACTIVE
	// indicates that you can register container instances with the cluster and
	// the associated instances can accept tasks.
	Status *string `locationName:"status" type:"string"`

	metadataCluster `json:"-" xml:"-"`
}

type metadataCluster struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Cluster) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Cluster) GoString() string {
	return s.String()
}

// A docker container that is part of a task.
type Container struct {
	// The Amazon Resource Name (ARN) of the container.
	ContainerArn *string `locationName:"containerArn" type:"string"`

	// The exit code returned from the container.
	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	// The last known status of the container.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// The name of the container.
	Name *string `locationName:"name" type:"string"`

	// The network bindings associated with the container.
	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	// A short (255 max characters) human-readable string to provide additional
	// detail about a running or stopped container.
	Reason *string `locationName:"reason" type:"string"`

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string `locationName:"taskArn" type:"string"`

	metadataContainer `json:"-" xml:"-"`
}

type metadataContainer struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Container) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Container) GoString() string {
	return s.String()
}

// Container definitions are used in task definitions to describe the different
// containers that are launched as part of a task.
type ContainerDefinition struct {
	// The CMD that is passed to the container. For more information on the Docker
	// CMD parameter, see https://docs.docker.com/reference/builder/#cmd (https://docs.docker.com/reference/builder/#cmd).
	Command []*string `locationName:"command" type:"list"`

	// The number of cpu units reserved for the container. A container instance
	// has 1,024 cpu units for every CPU core. This parameter specifies the minimum
	// amount of CPU to reserve for a container, and containers share unallocated
	// CPU units with other containers on the instance with the same ratio as their
	// allocated amount.
	//
	//  You can determine the number of CPU units that are available per Amazon
	// EC2 instance type by multiplying the vCPUs listed for that instance type
	// on the Amazon EC2 Instances (http://aws.amazon.com/ec2/instance-types/) detail
	// page by 1,024.
	//
	//  For example, if you run a single-container task on a single-core instance
	// type with 512 CPU units specified for that container, and that is the only
	// task running on the container instance, that container could use the full
	// 1,024 CPU unit share at any given time. However, if you launched another
	// copy of the same task on that container instance, each task would be guaranteed
	// a minimum of 512 CPU units when needed, and each container could float to
	// higher CPU usage if the other container was not using it, but if both tasks
	// were 100% active all of the time, they would be limited to 512 CPU units.
	//
	// The Docker daemon on the container instance uses the CPU value to calculate
	// the relative CPU share ratios for running containers. For more information,
	// see CPU share constraint (https://docs.docker.com/reference/run/#cpu-share-constraint)
	// in the Docker documentation. The minimum valid CPU share value that the Linux
	// kernel will allow is 2; however, the CPU parameter is not required, and you
	// can use CPU values below 2 in your container definitions. For CPU values
	// below 2 (including null), the behavior varies based on your Amazon ECS container
	// agent version:
	//
	//   Agent versions less than or equal to 1.1.0: Null and zero CPU values are
	// passed to Docker as 0, which Docker then converts to 1,024 CPU shares. CPU
	// values of 1 are passed to Docker as 1, which the Linux kernel converts to
	// 2 CPU shares.  Agent versions greater than or equal to 1.2.0: Null, zero,
	// and CPU values of 1 are passed to Docker as 2.
	Cpu *int64 `locationName:"cpu" type:"integer"`

	DisableNetworking *bool `locationName:"disableNetworking" type:"boolean"`

	DnsSearchDomains []*string `locationName:"dnsSearchDomains" type:"list"`

	DnsServers []*string `locationName:"dnsServers" type:"list"`

	DockerLabels map[string]*string `locationName:"dockerLabels" type:"map"`

	DockerSecurityOptions []*string `locationName:"dockerSecurityOptions" type:"list"`

	// Early versions of the Amazon ECS container agent do not properly handle entryPoint
	// parameters. If you have problems using entryPoint, update your container
	// agent or enter your commands and arguments as command array items instead.
	//
	//  The ENTRYPOINT that is passed to the container. For more information on
	// the Docker ENTRYPOINT parameter, see https://docs.docker.com/reference/builder/#entrypoint
	// (https://docs.docker.com/reference/builder/#entrypoint).
	EntryPoint []*string `locationName:"entryPoint" type:"list"`

	// The environment variables to pass to a container.
	Environment []*KeyValuePair `locationName:"environment" type:"list"`

	// If the essential parameter of a container is marked as true, the failure
	// of that container will stop the task. If the essential parameter of a container
	// is marked as false, then its failure will not affect the rest of the containers
	// in a task. If this parameter is omitted, a container is assumed to be essential.
	//
	//  All tasks must have at least one essential container.
	Essential *bool `locationName:"essential" type:"boolean"`

	ExtraHosts []*HostEntry `locationName:"extraHosts" type:"list"`

	Hostname *string `locationName:"hostname" type:"string"`

	// The image used to start a container. This string is passed directly to the
	// Docker daemon. Images in the Docker Hub registry are available by default.
	// Other repositories are specified with repository-url/image:tag.
	Image *string `locationName:"image" type:"string"`

	// The link parameter allows containers to communicate with each other without
	// the need for port mappings, using the name parameter. The name:internalName
	// construct is analogous to name:alias in Docker links. For more information
	// on linking Docker containers, see https://docs.docker.com/userguide/dockerlinks/
	// (https://docs.docker.com/userguide/dockerlinks/).
	//
	//  Containers that are collocated on a single container instance may be able
	// to communicate with each other without requiring links or host port mappings.
	// Network isolation is achieved on the container instance using security groups
	// and VPC settings.
	Links []*string `locationName:"links" type:"list"`

	LogConfiguration *LogConfiguration `locationName:"logConfiguration" type:"structure"`

	// The number of MiB of memory reserved for the container. If your container
	// attempts to exceed the memory allocated here, the container is killed.
	Memory *int64 `locationName:"memory" type:"integer"`

	// The mount points for data volumes in your container.
	MountPoints []*MountPoint `locationName:"mountPoints" type:"list"`

	// The name of a container. If you are linking multiple containers together
	// in a task definition, the name of one container can be entered in the links
	// of another container to connect the containers.
	Name *string `locationName:"name" type:"string"`

	// The list of port mappings for the container.
	PortMappings []*PortMapping `locationName:"portMappings" type:"list"`

	Privileged *bool `locationName:"privileged" type:"boolean"`

	ReadonlyRootFilesystem *bool `locationName:"readonlyRootFilesystem" type:"boolean"`

	Ulimits []*Ulimit `locationName:"ulimits" type:"list"`

	User *string `locationName:"user" type:"string"`

	// Data volumes to mount from another container.
	VolumesFrom []*VolumeFrom `locationName:"volumesFrom" type:"list"`

	WorkingDirectory *string `locationName:"workingDirectory" type:"string"`

	metadataContainerDefinition `json:"-" xml:"-"`
}

type metadataContainerDefinition struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ContainerDefinition) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerDefinition) GoString() string {
	return s.String()
}

// An Amazon EC2 instance that is running the Amazon ECS agent and has been
// registered with a cluster.
type ContainerInstance struct {
	// This parameter returns true if the agent is actually connected to Amazon
	// ECS. Registered instances with an agent that may be unhealthy or stopped
	// will return false, and instances without a connected agent cannot accept
	// placement request.
	AgentConnected *bool `locationName:"agentConnected" type:"boolean"`

	// The status of the most recent agent update. If an update has never been requested,
	// this value is NULL.
	AgentUpdateStatus *string `locationName:"agentUpdateStatus" type:"string" enum:"AgentUpdateStatus"`

	Attributes []*Attribute `locationName:"attributes" type:"list"`

	// The Amazon Resource Name (ARN) of the container instance. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the container instance,
	// the AWS account ID of the container instance owner, the container-instance
	// namespace, and then the container instance UUID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_UUID.
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The Amazon EC2 instance ID of the container instance.
	Ec2InstanceId *string `locationName:"ec2InstanceId" type:"string"`

	// The number of tasks on the container instance that are in the PENDING status.
	PendingTasksCount *int64 `locationName:"pendingTasksCount" type:"integer"`

	// The registered resources on the container instance that are in use by current
	// tasks.
	RegisteredResources []*Resource `locationName:"registeredResources" type:"list"`

	// The remaining resources of the container instance that are available for
	// new tasks.
	RemainingResources []*Resource `locationName:"remainingResources" type:"list"`

	// The number of tasks on the container instance that are in the RUNNING status.
	RunningTasksCount *int64 `locationName:"runningTasksCount" type:"integer"`

	// The status of the container instance. The valid values are ACTIVE or INACTIVE.
	// ACTIVE indicates that the container instance can accept tasks.
	Status *string `locationName:"status" type:"string"`

	// The version information for the Amazon ECS container agent and Docker daemon
	// running on the container instance.
	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`

	metadataContainerInstance `json:"-" xml:"-"`
}

type metadataContainerInstance struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ContainerInstance) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerInstance) GoString() string {
	return s.String()
}

// The overrides that should be sent to a container.
type ContainerOverride struct {
	// The command to send to the container that overrides the default command from
	// the Docker image or the task definition.
	Command []*string `locationName:"command" type:"list"`

	// The environment variables to send to the container. You can add new environment
	// variables, which are added to the container at launch, or you can override
	// the existing environment variables from the Docker image or the task definition.
	Environment []*KeyValuePair `locationName:"environment" type:"list"`

	// The name of the container that receives the override.
	Name *string `locationName:"name" type:"string"`

	metadataContainerOverride `json:"-" xml:"-"`
}

type metadataContainerOverride struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ContainerOverride) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerOverride) GoString() string {
	return s.String()
}

type CreateClusterInput struct {
	// The name of your cluster. If you do not specify a name for your cluster,
	// you will create a cluster named default. Up to 255 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	ClusterName *string `locationName:"clusterName" type:"string"`

	metadataCreateClusterInput `json:"-" xml:"-"`
}

type metadataCreateClusterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterInput) GoString() string {
	return s.String()
}

type CreateClusterOutput struct {
	// The full description of your new cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`

	metadataCreateClusterOutput `json:"-" xml:"-"`
}

type metadataCreateClusterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterOutput) GoString() string {
	return s.String()
}

type CreateServiceInput struct {
	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. Up to 32 ASCII characters are allowed.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to run your service on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the specified task definition that you would
	// like to place and keep running on your cluster.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer" required:"true"`

	// A list of load balancer objects, containing the load balancer name, the container
	// name (as it appears in a container definition), and the container port to
	// access from the load balancer.
	LoadBalancers []*LoadBalancer `locationName:"loadBalancers" type:"list"`

	// The name or full Amazon Resource Name (ARN) of the IAM role that allows your
	// Amazon ECS container agent to make calls to your load balancer on your behalf.
	// This parameter is only required if you are using a load balancer with your
	// service.
	Role *string `locationName:"role" type:"string"`

	// The name of your service. Up to 255 letters (uppercase and lowercase), numbers,
	// hyphens, and underscores are allowed. Service names must be unique within
	// a cluster, but you can have similarly named services in multiple clusters
	// within a region or across multiple regions.
	ServiceName *string `locationName:"serviceName" type:"string" required:"true"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run in your service. If a revision
	// is not specified, the latest ACTIVE revision is used.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataCreateServiceInput `json:"-" xml:"-"`
}

type metadataCreateServiceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServiceInput) GoString() string {
	return s.String()
}

type CreateServiceOutput struct {
	// The full description of your service following the create call.
	Service *Service `locationName:"service" type:"structure"`

	metadataCreateServiceOutput `json:"-" xml:"-"`
}

type metadataCreateServiceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s CreateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServiceOutput) GoString() string {
	return s.String()
}

type DeleteClusterInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to delete.
	Cluster *string `locationName:"cluster" type:"string" required:"true"`

	metadataDeleteClusterInput `json:"-" xml:"-"`
}

type metadataDeleteClusterInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteClusterInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteClusterInput) GoString() string {
	return s.String()
}

type DeleteClusterOutput struct {
	// The full description of the deleted cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`

	metadataDeleteClusterOutput `json:"-" xml:"-"`
}

type metadataDeleteClusterOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteClusterOutput) GoString() string {
	return s.String()
}

type DeleteServiceInput struct {
	// The name of the cluster that hosts the service you want to delete.
	Cluster *string `locationName:"cluster" type:"string"`

	// The name of the service you want to delete.
	Service *string `locationName:"service" type:"string" required:"true"`

	metadataDeleteServiceInput `json:"-" xml:"-"`
}

type metadataDeleteServiceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteServiceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteServiceInput) GoString() string {
	return s.String()
}

type DeleteServiceOutput struct {
	// Details on a service within a cluster
	Service *Service `locationName:"service" type:"structure"`

	metadataDeleteServiceOutput `json:"-" xml:"-"`
}

type metadataDeleteServiceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeleteServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteServiceOutput) GoString() string {
	return s.String()
}

// The details of an Amazon ECS service deployment.
type Deployment struct {
	// The Unix time in seconds and milliseconds when the service was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The most recent desired count of tasks that was specified for the service
	// to deploy and/or maintain.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The ID of the deployment.
	Id *string `locationName:"id" type:"string"`

	// The number of tasks in the deployment that are in the PENDING status.
	PendingCount *int64 `locationName:"pendingCount" type:"integer"`

	// The number of tasks in the deployment that are in the RUNNING status.
	RunningCount *int64 `locationName:"runningCount" type:"integer"`

	// The status of the deployment. Valid values are PRIMARY (for the most recent
	// deployment), ACTIVE (for previous deployments that still have tasks running,
	// but are being replaced with the PRIMARY deployment), and INACTIVE (for deployments
	// that have been completely replaced).
	Status *string `locationName:"status" type:"string"`

	// The most recent task definition that was specified for the service to use.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	// The Unix time in seconds and milliseconds when the service was last updated.
	UpdatedAt *time.Time `locationName:"updatedAt" type:"timestamp" timestampFormat:"unix"`

	metadataDeployment `json:"-" xml:"-"`
}

type metadataDeployment struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Deployment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Deployment) GoString() string {
	return s.String()
}

type DeregisterContainerInstanceInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instance you want to deregister. If you do not specify a cluster,
	// the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) of the container
	// instance you want to deregister. The ARN contains the arn:aws:ecs namespace,
	// followed by the region of the container instance, the AWS account ID of the
	// container instance owner, the container-instance namespace, and then the
	// container instance UUID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_UUID.
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`

	// Force the deregistration of the container instance. If you have tasks running
	// on the container instance when you deregister it with the force option, these
	// tasks remain running and they will continue to pass Elastic Load Balancing
	// load balancer health checks until you terminate the instance or the tasks
	// stop through some other means, but they are orphaned (no longer monitored
	// or accounted for by Amazon ECS). If an orphaned task on your container instance
	// is part of an Amazon ECS service, then the service scheduler will start another
	// copy of that task on a different container instance if possible.
	Force *bool `locationName:"force" type:"boolean"`

	metadataDeregisterContainerInstanceInput `json:"-" xml:"-"`
}

type metadataDeregisterContainerInstanceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeregisterContainerInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterContainerInstanceInput) GoString() string {
	return s.String()
}

type DeregisterContainerInstanceOutput struct {
	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`

	metadataDeregisterContainerInstanceOutput `json:"-" xml:"-"`
}

type metadataDeregisterContainerInstanceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeregisterContainerInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterContainerInstanceOutput) GoString() string {
	return s.String()
}

type DeregisterTaskDefinitionInput struct {
	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to deregister. You must specify a revision.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataDeregisterTaskDefinitionInput `json:"-" xml:"-"`
}

type metadataDeregisterTaskDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeregisterTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterTaskDefinitionInput) GoString() string {
	return s.String()
}

type DeregisterTaskDefinitionOutput struct {
	// The full description of the deregistered task.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`

	metadataDeregisterTaskDefinitionOutput `json:"-" xml:"-"`
}

type metadataDeregisterTaskDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DeregisterTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterTaskDefinitionOutput) GoString() string {
	return s.String()
}

type DescribeClustersInput struct {
	// A space-separated list of cluster names or full cluster Amazon Resource Name
	// (ARN) entries. If you do not specify a cluster, the default cluster is assumed.
	Clusters []*string `locationName:"clusters" type:"list"`

	metadataDescribeClustersInput `json:"-" xml:"-"`
}

type metadataDescribeClustersInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeClustersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeClustersInput) GoString() string {
	return s.String()
}

type DescribeClustersOutput struct {
	// The list of clusters.
	Clusters []*Cluster `locationName:"clusters" type:"list"`

	Failures []*Failure `locationName:"failures" type:"list"`

	metadataDescribeClustersOutput `json:"-" xml:"-"`
}

type metadataDescribeClustersOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeClustersOutput) GoString() string {
	return s.String()
}

type DescribeContainerInstancesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances you want to describe. If you do not specify a cluster,
	// the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A space-separated list of container instance UUIDs or full Amazon Resource
	// Name (ARN) entries.
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`

	metadataDescribeContainerInstancesInput `json:"-" xml:"-"`
}

type metadataDescribeContainerInstancesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeContainerInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContainerInstancesInput) GoString() string {
	return s.String()
}

type DescribeContainerInstancesOutput struct {
	// The list of container instances.
	ContainerInstances []*ContainerInstance `locationName:"containerInstances" type:"list"`

	Failures []*Failure `locationName:"failures" type:"list"`

	metadataDescribeContainerInstancesOutput `json:"-" xml:"-"`
}

type metadataDescribeContainerInstancesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeContainerInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContainerInstancesOutput) GoString() string {
	return s.String()
}

type DescribeServicesInput struct {
	// The name of the cluster that hosts the service you want to describe.
	Cluster *string `locationName:"cluster" type:"string"`

	// A list of services you want to describe.
	Services []*string `locationName:"services" type:"list" required:"true"`

	metadataDescribeServicesInput `json:"-" xml:"-"`
}

type metadataDescribeServicesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesInput) GoString() string {
	return s.String()
}

type DescribeServicesOutput struct {
	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`

	// The list of services described.
	Services []*Service `locationName:"services" type:"list"`

	metadataDescribeServicesOutput `json:"-" xml:"-"`
}

type metadataDescribeServicesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesOutput) GoString() string {
	return s.String()
}

type DescribeTaskDefinitionInput struct {
	// The family for the latest ACTIVE revision, family and revision (family:revision)
	// for a specific revision in the family, or full Amazon Resource Name (ARN)
	// of the task definition that you want to describe.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataDescribeTaskDefinitionInput `json:"-" xml:"-"`
}

type metadataDescribeTaskDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTaskDefinitionInput) GoString() string {
	return s.String()
}

type DescribeTaskDefinitionOutput struct {
	// The full task definition description.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`

	metadataDescribeTaskDefinitionOutput `json:"-" xml:"-"`
}

type metadataDescribeTaskDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTaskDefinitionOutput) GoString() string {
	return s.String()
}

type DescribeTasksInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task you want to describe. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A space-separated list of task UUIDs or full Amazon Resource Name (ARN) entries.
	Tasks []*string `locationName:"tasks" type:"list" required:"true"`

	metadataDescribeTasksInput `json:"-" xml:"-"`
}

type metadataDescribeTasksInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeTasksInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTasksInput) GoString() string {
	return s.String()
}

type DescribeTasksOutput struct {
	Failures []*Failure `locationName:"failures" type:"list"`

	// The list of tasks.
	Tasks []*Task `locationName:"tasks" type:"list"`

	metadataDescribeTasksOutput `json:"-" xml:"-"`
}

type metadataDescribeTasksOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DescribeTasksOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTasksOutput) GoString() string {
	return s.String()
}

type DiscoverPollEndpointInput struct {
	// The cluster that the container instance belongs to.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) of the container
	// instance. The ARN contains the arn:aws:ecs namespace, followed by the region
	// of the container instance, the AWS account ID of the container instance owner,
	// the container-instance namespace, and then the container instance UUID. For
	// example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_UUID.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	metadataDiscoverPollEndpointInput `json:"-" xml:"-"`
}

type metadataDiscoverPollEndpointInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DiscoverPollEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DiscoverPollEndpointInput) GoString() string {
	return s.String()
}

type DiscoverPollEndpointOutput struct {
	// The endpoint for the Amazon ECS agent to poll.
	Endpoint *string `locationName:"endpoint" type:"string"`

	// The telemetry endpoint for the Amazon ECS agent.
	TelemetryEndpoint *string `locationName:"telemetryEndpoint" type:"string"`

	metadataDiscoverPollEndpointOutput `json:"-" xml:"-"`
}

type metadataDiscoverPollEndpointOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s DiscoverPollEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DiscoverPollEndpointOutput) GoString() string {
	return s.String()
}

// A failed resource.
type Failure struct {
	// The Amazon Resource Name (ARN) of the failed resource.
	Arn *string `locationName:"arn" type:"string"`

	// The reason for the failure.
	Reason *string `locationName:"reason" type:"string"`

	metadataFailure `json:"-" xml:"-"`
}

type metadataFailure struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Failure) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Failure) GoString() string {
	return s.String()
}

type HostEntry struct {
	Hostname *string `type:"string" required:"true"`

	IpAddress *string `type:"string" required:"true"`

	metadataHostEntry `json:"-" xml:"-"`
}

type metadataHostEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s HostEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HostEntry) GoString() string {
	return s.String()
}

// Details on a container instance host volume.
type HostVolumeProperties struct {
	// The path on the host container instance that is presented to the container.
	// If this parameter is empty, then the Docker daemon has assigned a host path
	// for you.
	SourcePath *string `locationName:"sourcePath" type:"string"`

	metadataHostVolumeProperties `json:"-" xml:"-"`
}

type metadataHostVolumeProperties struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s HostVolumeProperties) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HostVolumeProperties) GoString() string {
	return s.String()
}

// A key and value pair object.
type KeyValuePair struct {
	// The name of the key value pair. For environment variables, this is the name
	// of the environment variable.
	Name *string `locationName:"name" type:"string"`

	// The value of the key value pair. For environment variables, this is the value
	// of the environment variable.
	Value *string `locationName:"value" type:"string"`

	metadataKeyValuePair `json:"-" xml:"-"`
}

type metadataKeyValuePair struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s KeyValuePair) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s KeyValuePair) GoString() string {
	return s.String()
}

type ListClustersInput struct {
	// The maximum number of cluster results returned by ListClusters in paginated
	// output. When this parameter is used, ListClusters only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListClusters
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListClusters returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListClusters request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListClustersInput `json:"-" xml:"-"`
}

type metadataListClustersInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListClustersInput) GoString() string {
	return s.String()
}

type ListClustersOutput struct {
	// The list of full Amazon Resource Name (ARN) entries for each cluster associated
	// with your account.
	ClusterArns []*string `locationName:"clusterArns" type:"list"`

	// The nextToken value to include in a future ListClusters request. When the
	// results of a ListClusters request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListClustersOutput `json:"-" xml:"-"`
}

type metadataListClustersOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListClustersOutput) GoString() string {
	return s.String()
}

type ListContainerInstancesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances you want to list. If you do not specify a cluster,
	// the default cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of container instance results returned by ListContainerInstances
	// in paginated output. When this parameter is used, ListContainerInstances
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListContainerInstances request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListContainerInstances returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListContainerInstances
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListContainerInstancesInput `json:"-" xml:"-"`
}

type metadataListContainerInstancesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListContainerInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContainerInstancesInput) GoString() string {
	return s.String()
}

type ListContainerInstancesOutput struct {
	// The list of container instance full Amazon Resource Name (ARN) entries for
	// each container instance associated with the specified cluster.
	ContainerInstanceArns []*string `locationName:"containerInstanceArns" type:"list"`

	// The nextToken value to include in a future ListContainerInstances request.
	// When the results of a ListContainerInstances request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListContainerInstancesOutput `json:"-" xml:"-"`
}

type metadataListContainerInstancesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListContainerInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContainerInstancesOutput) GoString() string {
	return s.String()
}

type ListServicesInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the services you want to list. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of container instance results returned by ListServices
	// in paginated output. When this parameter is used, ListServices only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// ListServices request with the returned nextToken value. This value can be
	// between 1 and 100. If this parameter is not used, then ListServices returns
	// up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListServices request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListServicesInput `json:"-" xml:"-"`
}

type metadataListServicesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListServicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListServicesInput) GoString() string {
	return s.String()
}

type ListServicesOutput struct {
	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of full Amazon Resource Name (ARN) entries for each service associated
	// with the specified cluster.
	ServiceArns []*string `locationName:"serviceArns" type:"list"`

	metadataListServicesOutput `json:"-" xml:"-"`
}

type metadataListServicesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListServicesOutput) GoString() string {
	return s.String()
}

type ListTaskDefinitionFamiliesInput struct {
	// The familyPrefix is a string that is used to filter the results of ListTaskDefinitionFamilies.
	// If you specify a familyPrefix, only task definition family names that begin
	// with the familyPrefix string are returned.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition family results returned by ListTaskDefinitionFamilies
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitionFamilies request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListTaskDefinitionFamilies returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitionFamilies
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListTaskDefinitionFamiliesInput `json:"-" xml:"-"`
}

type metadataListTaskDefinitionFamiliesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionFamiliesInput) GoString() string {
	return s.String()
}

type ListTaskDefinitionFamiliesOutput struct {
	// The list of task definition family names that match the ListTaskDefinitionFamilies
	// request.
	Families []*string `locationName:"families" type:"list"`

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults,
	// this value can be used to retrieve the next page of results. This value is
	// null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	metadataListTaskDefinitionFamiliesOutput `json:"-" xml:"-"`
}

type metadataListTaskDefinitionFamiliesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionFamiliesOutput) GoString() string {
	return s.String()
}

type ListTaskDefinitionsInput struct {
	// The full family name that you want to filter the ListTaskDefinitions results
	// with. Specifying a familyPrefix will limit the listed task definitions to
	// task definition revisions that belong to that family.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition results returned by ListTaskDefinitions
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitions request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter is not used, then ListTaskDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The order in which to sort the results. Valid values are ASC and DESC. By
	// default (ASC), task definitions are listed lexicographically by family name
	// and in ascending numerical order by revision so that the newest task definitions
	// in a family are listed last. Setting this parameter to DESC reverses the
	// sort order on family name and revision so that the newest task definitions
	// in a family are listed first.
	Sort *string `locationName:"sort" type:"string" enum:"SortOrder"`

	// The task definition status that you want to filter the ListTaskDefinitions
	// results with. By default, only ACTIVE task definitions are listed. By setting
	// this parameter to INACTIVE, you can view task definitions that are INACTIVE
	// as long as an active task or service still references them. If you paginate
	// the resulting output, be sure to keep the status value constant in each subsequent
	// request.
	Status *string `locationName:"status" type:"string" enum:"TaskDefinitionStatus"`

	metadataListTaskDefinitionsInput `json:"-" xml:"-"`
}

type metadataListTaskDefinitionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListTaskDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionsInput) GoString() string {
	return s.String()
}

type ListTaskDefinitionsOutput struct {
	// The nextToken value to include in a future ListTaskDefinitions request. When
	// the results of a ListTaskDefinitions request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task definition Amazon Resource Name (ARN) entries for the ListTaskDefinitions
	// request.
	TaskDefinitionArns []*string `locationName:"taskDefinitionArns" type:"list"`

	metadataListTaskDefinitionsOutput `json:"-" xml:"-"`
}

type metadataListTaskDefinitionsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListTaskDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionsOutput) GoString() string {
	return s.String()
}

type ListTasksInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the tasks you want to list. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) of the container
	// instance that you want to filter the ListTasks results with. Specifying a
	// containerInstance will limit the results to tasks that belong to that container
	// instance.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	// The task status that you want to filter the ListTasks results with. Specifying
	// a desiredStatus of STOPPED will limit the results to tasks that are in the
	// STOPPED status, which can be useful for debugging tasks that are not starting
	// properly or have died or finished. The default status filter is RUNNING.
	DesiredStatus *string `locationName:"desiredStatus" type:"string" enum:"DesiredStatus"`

	// The name of the family that you want to filter the ListTasks results with.
	// Specifying a family will limit the results to tasks that belong to that family.
	Family *string `locationName:"family" type:"string"`

	// The maximum number of task results returned by ListTasks in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in
	// a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListTasks request with
	// the returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, then ListTasks returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTasks request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the service that you want to filter the ListTasks results with.
	// Specifying a serviceName will limit the results to tasks that belong to that
	// service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The startedBy value that you want to filter the task results with. Specifying
	// a startedBy value will limit the results to tasks that were started with
	// that value.
	StartedBy *string `locationName:"startedBy" type:"string"`

	metadataListTasksInput `json:"-" xml:"-"`
}

type metadataListTasksInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListTasksInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTasksInput) GoString() string {
	return s.String()
}

type ListTasksOutput struct {
	// The nextToken value to include in a future ListTasks request. When the results
	// of a ListTasks request exceed maxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task Amazon Resource Name (ARN) entries for the ListTasks request.
	TaskArns []*string `locationName:"taskArns" type:"list"`

	metadataListTasksOutput `json:"-" xml:"-"`
}

type metadataListTasksOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ListTasksOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTasksOutput) GoString() string {
	return s.String()
}

// Details on a load balancer that is used with a service.
type LoadBalancer struct {
	// The name of the container to associate with the load balancer.
	ContainerName *string `locationName:"containerName" type:"string"`

	// The port on the container to associate with the load balancer. This port
	// must correspond to a containerPort in the service's task definition. Your
	// container instances must allow ingress traffic on the hostPort of the port
	// mapping.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The name of the load balancer.
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string"`

	metadataLoadBalancer `json:"-" xml:"-"`
}

type metadataLoadBalancer struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s LoadBalancer) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancer) GoString() string {
	return s.String()
}

type LogConfiguration struct {
	LogDriver *string `locationName:"logDriver" type:"string" required:"true"`

	Options map[string]*string `locationName:"options" type:"map"`

	metadataLogConfiguration `json:"-" xml:"-"`
}

type metadataLogConfiguration struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s LogConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LogConfiguration) GoString() string {
	return s.String()
}

// Details on a volume mount point that is used in a container definition.
type MountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `locationName:"containerPath" type:"string"`

	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume. The default
	// value is false.
	ReadOnly *bool `locationName:"readOnly" type:"boolean"`

	// The name of the volume to mount.
	SourceVolume *string `locationName:"sourceVolume" type:"string"`

	metadataMountPoint `json:"-" xml:"-"`
}

type metadataMountPoint struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s MountPoint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MountPoint) GoString() string {
	return s.String()
}

// Details on the network bindings between a container and its host container
// instance.
type NetworkBinding struct {
	// The IP address that the container is bound to on the container instance.
	BindIP *string `locationName:"bindIP" type:"string"`

	// The port number on the container that is be used with the network binding.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The port number on the host that is used with the network binding.
	HostPort *int64 `locationName:"hostPort" type:"integer"`

	// The protocol used for the network binding.
	Protocol *string `locationName:"protocol" type:"string" enum:"TransportProtocol"`

	metadataNetworkBinding `json:"-" xml:"-"`
}

type metadataNetworkBinding struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s NetworkBinding) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkBinding) GoString() string {
	return s.String()
}

// Port mappings allow containers to access ports on the host container instance
// to send or receive traffic. Port mappings are specified as part of the container
// definition.
type PortMapping struct {
	// The port number on the container that is bound to the user-specified or automatically
	// assigned host port. If you specify a container port and not a host port,
	// your container will automatically receive a host port in the ephemeral port
	// range (for more information, see hostPort).
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The port number on the container instance to reserve for your container.
	// You can specify a non-reserved host port for your container port mapping,
	// or you can omit the hostPort (or set it to 0) while specifying a containerPort
	// and your container will automatically receive a port in the ephemeral port
	// range for your container instance operating system and Docker version.
	//
	// The default ephemeral port range is 49153 to 65535, and this range is used
	// for Docker versions prior to 1.6.0. For Docker version 1.6.0 and later, the
	// Docker daemon tries to read the ephemeral port range from /proc/sys/net/ipv4/ip_local_port_range;
	// if this kernel parameter is unavailable, the default ephemeral port range
	// is used. You should not attempt to specify a host port in the ephemeral port
	// range, since these are reserved for automatic assignment. In general, ports
	// below 32768 are outside of the ephemeral port range.
	//
	// The default reserved ports are 22 for SSH, the Docker ports 2375 and 2376,
	// and the Amazon ECS Container Agent port 51678. Any host port that was previously
	// specified in a running task is also reserved while the task is running (once
	// a task stops, the host port is released).The current reserved ports are displayed
	// in the remainingResources of DescribeContainerInstances output, and a container
	// instance may have up to 50 reserved ports at a time, including the default
	// reserved ports (automatically assigned ports do not count toward this limit).
	HostPort *int64 `locationName:"hostPort" type:"integer"`

	// The protocol used for the port mapping. Valid values are tcp and udp. The
	// default is tcp.
	Protocol *string `locationName:"protocol" type:"string" enum:"TransportProtocol"`

	metadataPortMapping `json:"-" xml:"-"`
}

type metadataPortMapping struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s PortMapping) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PortMapping) GoString() string {
	return s.String()
}

type RegisterContainerInstanceInput struct {
	Attributes []*Attribute `locationName:"attributes" type:"list"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to register your container instance with. If you do not specify a cluster,
	// the default cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The Amazon Resource Name (ARN) of the container instance (if it was previously
	// registered).
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The instance identity document for the Amazon EC2 instance to register. This
	// document can be found by running the following command from the instance:
	// curl http://169.254.169.254/latest/dynamic/instance-identity/document/
	InstanceIdentityDocument *string `locationName:"instanceIdentityDocument" type:"string"`

	// The instance identity document signature for the Amazon EC2 instance to register.
	// This signature can be found by running the following command from the instance:
	// curl http://169.254.169.254/latest/dynamic/instance-identity/signature/
	InstanceIdentityDocumentSignature *string `locationName:"instanceIdentityDocumentSignature" type:"string"`

	// The resources available on the instance.
	TotalResources []*Resource `locationName:"totalResources" type:"list"`

	// The version information for the Amazon ECS container agent and Docker daemon
	// running on the container instance.
	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`

	metadataRegisterContainerInstanceInput `json:"-" xml:"-"`
}

type metadataRegisterContainerInstanceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RegisterContainerInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterContainerInstanceInput) GoString() string {
	return s.String()
}

type RegisterContainerInstanceOutput struct {
	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`

	metadataRegisterContainerInstanceOutput `json:"-" xml:"-"`
}

type metadataRegisterContainerInstanceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RegisterContainerInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterContainerInstanceOutput) GoString() string {
	return s.String()
}

type RegisterTaskDefinitionInput struct {
	// A list of container definitions in JSON format that describe the different
	// containers that make up your task.
	ContainerDefinitions []*ContainerDefinition `locationName:"containerDefinitions" type:"list" required:"true"`

	// You must specify a family for a task definition, which allows you to track
	// multiple versions of the same task definition. You can think of the family
	// as a name for your task definition. Up to 255 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	Family *string `locationName:"family" type:"string" required:"true"`

	// A list of volume definitions in JSON format that containers in your task
	// may use.
	Volumes []*Volume `locationName:"volumes" type:"list"`

	metadataRegisterTaskDefinitionInput `json:"-" xml:"-"`
}

type metadataRegisterTaskDefinitionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RegisterTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterTaskDefinitionInput) GoString() string {
	return s.String()
}

type RegisterTaskDefinitionOutput struct {
	// Details of a task definition.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`

	metadataRegisterTaskDefinitionOutput `json:"-" xml:"-"`
}

type metadataRegisterTaskDefinitionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RegisterTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterTaskDefinitionOutput) GoString() string {
	return s.String()
}

// Describes the resources available for a container instance.
type Resource struct {
	// When the doubleValue type is set, the value of the resource must be a double
	// precision floating-point type.
	DoubleValue *float64 `locationName:"doubleValue" type:"double"`

	// When the integerValue type is set, the value of the resource must be an integer.
	IntegerValue *int64 `locationName:"integerValue" type:"integer"`

	// When the longValue type is set, the value of the resource must be an extended
	// precision floating-point type.
	LongValue *int64 `locationName:"longValue" type:"long"`

	// The name of the resource, such as CPU, MEMORY, PORTS, or a user-defined resource.
	Name *string `locationName:"name" type:"string"`

	// When the stringSetValue type is set, the value of the resource must be a
	// string type.
	StringSetValue []*string `locationName:"stringSetValue" type:"list"`

	// The type of the resource, such as INTEGER, DOUBLE, LONG, or STRINGSET.
	Type *string `locationName:"type" type:"string"`

	metadataResource `json:"-" xml:"-"`
}

type metadataResource struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Resource) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Resource) GoString() string {
	return s.String()
}

type RunTaskInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to run your task on. If you do not specify a cluster, the default cluster
	// is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the specified task that you would like to
	// place on your cluster.
	//
	//  The count parameter is limited to 10 tasks per call.
	Count *int64 `locationName:"count" type:"integer"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	//
	//  A total of 8192 characters are allowed for overrides. This limit includes
	// the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// An optional tag specified when a task is started. For example if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run. If a revision is not specified,
	// the latest ACTIVE revision is used.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataRunTaskInput `json:"-" xml:"-"`
}

type metadataRunTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RunTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RunTaskInput) GoString() string {
	return s.String()
}

type RunTaskOutput struct {
	// Any failed tasks from your RunTask action are listed here.
	Failures []*Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were run. Each task that was successfully
	// placed on your cluster will be described here.
	Tasks []*Task `locationName:"tasks" type:"list"`

	metadataRunTaskOutput `json:"-" xml:"-"`
}

type metadataRunTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s RunTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RunTaskOutput) GoString() string {
	return s.String()
}

// Details on a service within a cluster
type Service struct {
	// The Amazon Resource Name (ARN) of the of the cluster that hosts the service.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The current state of deployments for the service.
	Deployments []*Deployment `locationName:"deployments" type:"list"`

	// The desired number of instantiations of the task definition to keep running
	// on the service. This value is specified when the service is created with
	// CreateService, and it can be modified with UpdateService.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The event stream for your service. A maximum of 100 of the latest events
	// are displayed.
	Events []*ServiceEvent `locationName:"events" type:"list"`

	// A list of load balancer objects, containing the load balancer name, the container
	// name (as it appears in a container definition), and the container port to
	// access from the load balancer.
	LoadBalancers []*LoadBalancer `locationName:"loadBalancers" type:"list"`

	// The number of tasks in the cluster that are in the PENDING state.
	PendingCount *int64 `locationName:"pendingCount" type:"integer"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the service
	// that allows the Amazon ECS container agent to register container instances
	// with a load balancer.
	RoleArn *string `locationName:"roleArn" type:"string"`

	// The number of tasks in the cluster that are in the RUNNING state.
	RunningCount *int64 `locationName:"runningCount" type:"integer"`

	// The Amazon Resource Name (ARN) that identifies the service. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the service, the AWS
	// account ID of the service owner, the service namespace, and then the service
	// name. For example, arn:aws:ecs:region:012345678910:service/my-service.
	ServiceArn *string `locationName:"serviceArn" type:"string"`

	// A user-generated string that you can use to identify your service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The status of the service. The valid values are ACTIVE, DRAINING, or INACTIVE.
	Status *string `locationName:"status" type:"string"`

	// The task definition to use for tasks in the service. This value is specified
	// when the service is created with CreateService, and it can be modified with
	// UpdateService.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	metadataService `json:"-" xml:"-"`
}

type metadataService struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Service) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Service) GoString() string {
	return s.String()
}

// Details on an event associated with a service.
type ServiceEvent struct {
	// The Unix time in seconds and milliseconds when the event was triggered.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The ID string of the event.
	Id *string `locationName:"id" type:"string"`

	// The event message.
	Message *string `locationName:"message" type:"string"`

	metadataServiceEvent `json:"-" xml:"-"`
}

type metadataServiceEvent struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s ServiceEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ServiceEvent) GoString() string {
	return s.String()
}

type StartTaskInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to start your task on. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUIDs or full Amazon Resource Name (ARN) entries for
	// the container instances on which you would like to place your task.
	//
	//  The list of container instances to start tasks on is limited to 10.
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	//
	//  A total of 8192 characters are allowed for overrides. This limit includes
	// the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// An optional tag specified when a task is started. For example if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to start. If a revision is not specified,
	// the latest ACTIVE revision is used.
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`

	metadataStartTaskInput `json:"-" xml:"-"`
}

type metadataStartTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s StartTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTaskInput) GoString() string {
	return s.String()
}

type StartTaskOutput struct {
	// Any failed tasks from your StartTask action are listed here.
	Failures []*Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were started. Each task that was successfully
	// placed on your container instances will be described here.
	Tasks []*Task `locationName:"tasks" type:"list"`

	metadataStartTaskOutput `json:"-" xml:"-"`
}

type metadataStartTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s StartTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTaskOutput) GoString() string {
	return s.String()
}

type StopTaskInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task you want to stop. If you do not specify a cluster, the default cluster
	// is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The task UUIDs or full Amazon Resource Name (ARN) entry of the task you would
	// like to stop.
	Task *string `locationName:"task" type:"string" required:"true"`

	metadataStopTaskInput `json:"-" xml:"-"`
}

type metadataStopTaskInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s StopTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopTaskInput) GoString() string {
	return s.String()
}

type StopTaskOutput struct {
	// Details on a task in a cluster.
	Task *Task `locationName:"task" type:"structure"`

	metadataStopTaskOutput `json:"-" xml:"-"`
}

type metadataStopTaskOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s StopTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopTaskOutput) GoString() string {
	return s.String()
}

type SubmitContainerStateChangeInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container.
	Cluster *string `locationName:"cluster" type:"string"`

	// The name of the container.
	ContainerName *string `locationName:"containerName" type:"string"`

	// The exit code returned for the state change request.
	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	// The network bindings of the container.
	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task UUID or full Amazon Resource Name (ARN) of the task that hosts the
	// container.
	Task *string `locationName:"task" type:"string"`

	metadataSubmitContainerStateChangeInput `json:"-" xml:"-"`
}

type metadataSubmitContainerStateChangeInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s SubmitContainerStateChangeInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitContainerStateChangeInput) GoString() string {
	return s.String()
}

type SubmitContainerStateChangeOutput struct {
	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`

	metadataSubmitContainerStateChangeOutput `json:"-" xml:"-"`
}

type metadataSubmitContainerStateChangeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s SubmitContainerStateChangeOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitContainerStateChangeOutput) GoString() string {
	return s.String()
}

type SubmitTaskStateChangeInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task.
	Cluster *string `locationName:"cluster" type:"string"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task UUID or full Amazon Resource Name (ARN) of the task in the state
	// change request.
	Task *string `locationName:"task" type:"string"`

	metadataSubmitTaskStateChangeInput `json:"-" xml:"-"`
}

type metadataSubmitTaskStateChangeInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s SubmitTaskStateChangeInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitTaskStateChangeInput) GoString() string {
	return s.String()
}

type SubmitTaskStateChangeOutput struct {
	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`

	metadataSubmitTaskStateChangeOutput `json:"-" xml:"-"`
}

type metadataSubmitTaskStateChangeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s SubmitTaskStateChangeOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitTaskStateChangeOutput) GoString() string {
	return s.String()
}

// Details on a task in a cluster.
type Task struct {
	// The Amazon Resource Name (ARN) of the of the cluster that hosts the task.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the container instances that host the task.
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The containers associated with the task.
	Containers []*Container `locationName:"containers" type:"list"`

	// The desired status of the task.
	DesiredStatus *string `locationName:"desiredStatus" type:"string"`

	// The last known status of the task.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// One or more container overrides.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// The tag specified when a task is started. If the task is started by an Amazon
	// ECS service, then the startedBy parameter contains the deployment ID of the
	// service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string `locationName:"taskArn" type:"string"`

	// The Amazon Resource Name (ARN) of the of the task definition that creates
	// the task.
	TaskDefinitionArn *string `locationName:"taskDefinitionArn" type:"string"`

	metadataTask `json:"-" xml:"-"`
}

type metadataTask struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Task) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Task) GoString() string {
	return s.String()
}

// Details of a task definition.
type TaskDefinition struct {
	// A list of container definitions in JSON format that describe the different
	// containers that make up your task. For more information on container definition
	// parameters and defaults, see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	ContainerDefinitions []*ContainerDefinition `locationName:"containerDefinitions" type:"list"`

	// The family of your task definition. You can think of the family as the name
	// of your task definition.
	Family *string `locationName:"family" type:"string"`

	RequiresAttributes []*Attribute `locationName:"requiresAttributes" type:"list"`

	// The revision of the task in a particular family. You can think of the revision
	// as a version number of a task definition in a family. When you register a
	// task definition for the first time, the revision is 1, and each time you
	// register a new revision of a task definition in the same family, the revision
	// value always increases by one (even if you have deregistered previous revisions
	// in this family).
	Revision *int64 `locationName:"revision" type:"integer"`

	// The status of the task definition.
	Status *string `locationName:"status" type:"string" enum:"TaskDefinitionStatus"`

	// The full Amazon Resource Name (ARN) of the of the task definition.
	TaskDefinitionArn *string `locationName:"taskDefinitionArn" type:"string"`

	// The list of volumes in a task. For more information on volume definition
	// parameters and defaults, see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	Volumes []*Volume `locationName:"volumes" type:"list"`

	metadataTaskDefinition `json:"-" xml:"-"`
}

type metadataTaskDefinition struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s TaskDefinition) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskDefinition) GoString() string {
	return s.String()
}

// The overrides associated with a task.
type TaskOverride struct {
	// One or more container overrides sent to a task.
	ContainerOverrides []*ContainerOverride `locationName:"containerOverrides" type:"list"`

	metadataTaskOverride `json:"-" xml:"-"`
}

type metadataTaskOverride struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s TaskOverride) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskOverride) GoString() string {
	return s.String()
}

type Ulimit struct {
	HardLimit *int64 `locationName:"hardLimit" type:"integer" required:"true"`

	Name *string `locationName:"name" type:"string" required:"true"`

	SoftLimit *int64 `locationName:"softLimit" type:"integer" required:"true"`

	metadataUlimit `json:"-" xml:"-"`
}

type metadataUlimit struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Ulimit) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Ulimit) GoString() string {
	return s.String()
}

type UpdateContainerAgentInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// container instance is running on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance UUID or full Amazon Resource Name (ARN) entries for
	// the container instance on which you would like to update the Amazon ECS container
	// agent.
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`

	metadataUpdateContainerAgentInput `json:"-" xml:"-"`
}

type metadataUpdateContainerAgentInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateContainerAgentInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContainerAgentInput) GoString() string {
	return s.String()
}

type UpdateContainerAgentOutput struct {
	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`

	metadataUpdateContainerAgentOutput `json:"-" xml:"-"`
}

type metadataUpdateContainerAgentOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateContainerAgentOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContainerAgentOutput) GoString() string {
	return s.String()
}

type UpdateServiceInput struct {
	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// service is running on. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the task that you would like to place and
	// keep running in your service.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The name of the service that you want to update.
	Service *string `locationName:"service" type:"string" required:"true"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run in your service. If a revision
	// is not specified, the latest ACTIVE revision is used. If you modify the task
	// definition with UpdateService, Amazon ECS spawns a task with the new version
	// of the task definition and then stops an old task after the new version is
	// running.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	metadataUpdateServiceInput `json:"-" xml:"-"`
}

type metadataUpdateServiceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateServiceInput) GoString() string {
	return s.String()
}

type UpdateServiceOutput struct {
	// The full description of your service following the update call.
	Service *Service `locationName:"service" type:"structure"`

	metadataUpdateServiceOutput `json:"-" xml:"-"`
}

type metadataUpdateServiceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s UpdateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateServiceOutput) GoString() string {
	return s.String()
}

// The Docker and Amazon ECS container agent version information on a container
// instance.
type VersionInfo struct {
	// The Git commit hash for the Amazon ECS container agent build on the amazon-ecs-agent
	//  (https://github.com/aws/amazon-ecs-agent/commits/master) GitHub repository.
	AgentHash *string `locationName:"agentHash" type:"string"`

	// The version number of the Amazon ECS container agent.
	AgentVersion *string `locationName:"agentVersion" type:"string"`

	// The Docker version running on the container instance.
	DockerVersion *string `locationName:"dockerVersion" type:"string"`

	metadataVersionInfo `json:"-" xml:"-"`
}

type metadataVersionInfo struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s VersionInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s VersionInfo) GoString() string {
	return s.String()
}

// A data volume used in a task definition.
type Volume struct {
	// The path on the host container instance that is presented to the containers
	// which access the volume. If this parameter is empty, then the Docker daemon
	// assigns a host path for you.
	Host *HostVolumeProperties `locationName:"host" type:"structure"`

	// The name of the volume. This name is referenced in the sourceVolume parameter
	// of container definition mountPoints.
	Name *string `locationName:"name" type:"string"`

	metadataVolume `json:"-" xml:"-"`
}

type metadataVolume struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s Volume) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Volume) GoString() string {
	return s.String()
}

// Details on a data volume from another container.
type VolumeFrom struct {
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume. The default
	// value is false.
	ReadOnly *bool `locationName:"readOnly" type:"boolean"`

	// The name of the container to mount volumes from.
	SourceContainer *string `locationName:"sourceContainer" type:"string"`

	metadataVolumeFrom `json:"-" xml:"-"`
}

type metadataVolumeFrom struct {
	SDKShapeTraits bool `type:"structure"`
}

// String returns the string representation
func (s VolumeFrom) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeFrom) GoString() string {
	return s.String()
}

const (
	// @enum AgentUpdateStatus
	AgentUpdateStatusPending = "PENDING"
	// @enum AgentUpdateStatus
	AgentUpdateStatusStaging = "STAGING"
	// @enum AgentUpdateStatus
	AgentUpdateStatusStaged = "STAGED"
	// @enum AgentUpdateStatus
	AgentUpdateStatusUpdating = "UPDATING"
	// @enum AgentUpdateStatus
	AgentUpdateStatusUpdated = "UPDATED"
	// @enum AgentUpdateStatus
	AgentUpdateStatusFailed = "FAILED"
)

const (
	// @enum DesiredStatus
	DesiredStatusRunning = "RUNNING"
	// @enum DesiredStatus
	DesiredStatusPending = "PENDING"
	// @enum DesiredStatus
	DesiredStatusStopped = "STOPPED"
)

const (
	// @enum SortOrder
	SortOrderAsc = "ASC"
	// @enum SortOrder
	SortOrderDesc = "DESC"
)

const (
	// @enum TaskDefinitionStatus
	TaskDefinitionStatusActive = "ACTIVE"
	// @enum TaskDefinitionStatus
	TaskDefinitionStatusInactive = "INACTIVE"
)

const (
	// @enum TransportProtocol
	TransportProtocolTcp = "tcp"
	// @enum TransportProtocol
	TransportProtocolUdp = "udp"
)
