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

func ListApiKeys(opts ...DeployOption) ([]string, error) {

	deployOpts := fillDeployOptions(opts...)

	listApiKeysReq := &models.ListDeploymentsRequest{}
	listApiKeysParams := service_runner.NewListAPIKeysParams().
		WithBody(listApiKeysReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.ListAPIKeysOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.ListAPIKeys(listApiKeysParams,
			deployOpts)
		if err == nil {
			break
		}

		client = nil
		client = goswag.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return make([]string, 0), fmt.Errorf("Client/HTTP failure: %v", err)
	}
	listReply := resp.GetPayload()
	if listReply.Result.Status != nil &&
		*listReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return make([]string, 0),
			fmt.Errorf("ListApiKeys failure(%v): %v",
				*listReply.Result.Status, listReply.Result.StatusDetail)
	}

	return listReply.KeyIds, nil
}

func CreateApiKey(keyName string, desc string, expireTime time.Time,
	opts ...DeployOption) (*pb.CreateApiKeyReply, error) {

	deployOpts := fillDeployOptions(opts...)

	createApiKeyReq := &models.CreateAPIKeyRequest{
		Name:        keyName,
		Description: desc,
		ExpireTime:  strconv.FormatInt(expireTime.Unix(), 10),
	}

	createApiKeyParams := service_runner.NewCreateAPIKeyParams().
		WithBody(createApiKeyReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.CreateAPIKeyOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.CreateAPIKey(createApiKeyParams,
			deployOpts)
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
	createApiKeyReply := resp.GetPayload()

	if createApiKeyReply.Result.Status != nil &&
		*createApiKeyReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("CreateApiKey failure(%v): %v",
			*createApiKeyReply.Result.Status, createApiKeyReply.Result.StatusDetail)
	}
	ret := &pb.CreateApiKeyReply{
		Result: &pb.ServiceRunnerResult{
			Status:       0,
			StatusDetail: "",
		},
		KeyId:   createApiKeyReply.KeyID,
		KeyData: createApiKeyReply.KeyData,
	}

	return ret, nil
}

func DeleteApiKey(keyId string, opts ...DeployOption) error {

	deployOpts := fillDeployOptions(opts...)

	deleteApiKeyReq := &models.DeleteAPIKeyRequest{
		KeyID: keyId,
	}

	deleteApiKeyParams := service_runner.NewDeleteAPIKeyParams().
		WithBody(deleteApiKeyReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.DeleteAPIKeyOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DeleteAPIKey(deleteApiKeyParams,
			deployOpts)
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
	deleteApiKeyReply := resp.GetPayload()

	if deleteApiKeyReply.Result.Status != nil &&
		*deleteApiKeyReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("DeleteApiKey failure(%v): %v",
			*deleteApiKeyReply.Result.Status, deleteApiKeyReply.Result.StatusDetail)
	}

	return nil
}

func DescribeApiKey(keyId string,
	opts ...DeployOption) (*pb.DescribeApiKeyReply, error) {

	deployOpts := fillDeployOptions(opts...)

	describeApiKeyReq := &models.DescribeAPIKeyRequest{
		KeyID: keyId,
	}

	describeApiKeyParams := service_runner.NewDescribeAPIKeyParams().
		WithBody(describeApiKeyReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.DescribeAPIKeyOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeAPIKey(describeApiKeyParams,
			deployOpts)
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
	describeApiKeyReply := resp.GetPayload()

	if describeApiKeyReply.Result.Status != nil &&
		*describeApiKeyReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribeApiKey failure(%v): %v",
			*describeApiKeyReply.Result.Status, describeApiKeyReply.Result.StatusDetail)
	}
	ret := &pb.DescribeApiKeyReply{
		Result: &pb.ServiceRunnerResult{
			Status:       0,
			StatusDetail: "",
		},
		Desc: &pb.ApiKeyDescription{
			KeyId:       describeApiKeyReply.Desc.KeyID,
			Name:        describeApiKeyReply.Desc.Name,
			Description: describeApiKeyReply.Desc.Description,
			CreateTime:  convertRESTTimeToInt(describeApiKeyReply.Desc.CreateTime),
			ExpireTime:  convertRESTTimeToInt(describeApiKeyReply.Desc.ExpireTime),
			LastUsed:    convertRESTTimeToInt(describeApiKeyReply.Desc.LastUsed),
		},
	}

	return ret, nil
}
