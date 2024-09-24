/* Copyright © 2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package golang

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bopmatic/sdk/golang/goswag"
	"github.com/bopmatic/sdk/golang/goswag/service_runner"
	"github.com/bopmatic/sdk/golang/models"
	"github.com/bopmatic/sdk/golang/pb"
)

type Deployment struct {
	PkgId    string
	ProjId   string
	EnvId    string
	DeployId string
}

// NewDeployment instantiates a new Deployment instance
func NewDeployment(pkgId string, projId string, envId string) *Deployment {
	return &Deployment{
		PkgId:    pkgId,
		ProjId:   projId,
		EnvId:    envId,
		DeployId: "",
	}
}

func (deployment *Deployment) Deploy(opts ...DeployOption) error {
	if deployment.DeployId != "" {
		return nil // already deployed
	}

	deployOpts := fillDeployOptions(opts...)

	deployPackageReq := &models.CreateDeploymentRequest{
		Header: &models.DeploymentHeader{
			EnvID:     deployment.EnvId,
			Initiator: models.DeploymentInitiatorCUSTOMER.Pointer(),
			PkgID:     deployment.PkgId,
			ProjID:    deployment.ProjId,
			Type:      models.DeploymentTypeNEWPACKAGE.Pointer(),
		},
	}

	deployPackageParams := service_runner.NewCreateDeploymentParams().
		WithBody(deployPackageReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.CreateDeploymentOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.CreateDeployment(deployPackageParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}
	deployReply := resp.GetPayload()

	if *deployReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("Deploy failure(%v): %v",
			*deployReply.Result.Status, deployReply.Result.StatusDetail)
	}
	deployment.DeployId = deployReply.ID

	return nil
}

// Describe() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func (deployment *Deployment) Describe(opts ...DeployOption) (*pb.DeploymentDescription, error) {

	deployOpts := fillDeployOptions(opts...)

	describeDeploymentReq := &models.DescribeDeploymentRequest{
		ID: deployment.DeployId,
	}

	// default endpoint is inferred from host field in sr.bopmatic.json
	describeDeploymentParams := service_runner.NewDescribeDeploymentParams().
		WithBody(describeDeploymentReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.DescribeDeploymentOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeDeployment(describeDeploymentParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("Client/HTTP failure: %v", err)
	}
	describeReply := resp.GetPayload()
	if *describeReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribeDescribement failure(%v): %v",
			*describeReply.Result.Status, describeReply.Result.StatusDetail)
	}

	deployTypeInt, ok :=
		pb.DeploymentType_value[string(*describeReply.Desc.Header.Type)]
	deployType := pb.DeploymentType(deployTypeInt)
	if !ok {
		deployType = pb.DeploymentType_UNKNOWN_DEPLOY_TYPE
	}
	deployInitiatorInt, ok :=
		pb.DeploymentInitiator_value[string(*describeReply.Desc.Header.Initiator)]
	deployInitiator := pb.DeploymentInitiator(deployInitiatorInt)
	if !ok {
		deployInitiator = pb.DeploymentInitiator_UNKNOWN_DEPLOY_INIT
	}
	deployStateInt, ok :=
		pb.DeploymentState_value[string(*describeReply.Desc.State)]
	deployState := pb.DeploymentState(deployStateInt)
	if !ok {
		deployState = pb.DeploymentState_UNKNOWN_DEPLOY_STATE
	}
	deployStateDetailInt, ok :=
		pb.DeploymentStateDetail_value[string(*describeReply.Desc.StateDetail)]
	deployStateDetail := pb.DeploymentStateDetail(deployStateDetailInt)
	if !ok {
		deployStateDetail = pb.DeploymentStateDetail_UNKNOWN_DEPLOY_STATE_DET
	}
	ret := &pb.DeploymentDescription{
		Id: deployment.DeployId,
		Header: &pb.DeploymentHeader{
			PkgId:     deployment.PkgId,
			ProjId:    deployment.ProjId,
			EnvId:     deployment.EnvId,
			Type:      deployType,
			Initiator: deployInitiator,
		},
		State:               deployState,
		StateDetail:         deployStateDetail,
		CreateTime:          convertRESTTimeToInt(describeReply.Desc.CreateTime),
		ValidationStartTime: convertRESTTimeToInt(describeReply.Desc.ValidationStartTime),
		BuildStartTime:      convertRESTTimeToInt(describeReply.Desc.BuildStartTime),
		DeployStartTime:     convertRESTTimeToInt(describeReply.Desc.DeployStartTime),
		EndTime:             convertRESTTimeToInt(describeReply.Desc.EndTime),
	}

	return ret, nil
}

func convertRESTTimeToInt(timeStr string) uint64 {
	secs, err := strconv.ParseUint(timeStr, 10, 64)
	if err != nil {
		secs = 0
	}

	return secs
}

func convertRESTIntToInt(intStr string) uint64 {
	return convertRESTTimeToInt(intStr)
}
