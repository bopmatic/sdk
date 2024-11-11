/* Copyright © 2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package golang

import (
	"fmt"
	"time"

	"github.com/bopmatic/sdk/golang/goswag"
	"github.com/bopmatic/sdk/golang/goswag/service_runner"
	"github.com/bopmatic/sdk/golang/models"
	"github.com/bopmatic/sdk/golang/pb"
)

func ListServices(projId string, envId string,
	opts ...DeployOption) ([]string, error) {

	deployOpts := fillDeployOptions(opts...)

	listServicesReq := &models.ListServicesRequest{
		Header: &models.ProjEnvHeader{
			ProjID: projId,
			EnvID:  envId,
		},
	}
	listServicesParams := service_runner.NewListServicesParams().
		WithBody(listServicesReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.ListServicesOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.ListServices(listServicesParams,
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
			fmt.Errorf("ListServices failure(%v): %v",
				*listReply.Result.Status, listReply.Result.StatusDetail)
	}

	return listReply.ServiceNames, nil
}

func DescribeService(projId string, envId string, svcName string,
	opts ...DeployOption) (*pb.DescribeServiceReply, error) {

	deployOpts := fillDeployOptions(opts...)

	describeServiceReq := &models.DescribeServiceRequest{
		SvcHeader: &models.ServiceHeader{
			ProjEnvHeader: &models.ProjEnvHeader{
				ProjID: projId,
				EnvID:  envId,
			},
			ServiceName: svcName,
		},
	}

	describeServiceParams := service_runner.NewDescribeServiceParams().
		WithBody(describeServiceReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.DescribeServiceOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeService(describeServiceParams,
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
	describeServiceReply := resp.GetPayload()

	if describeServiceReply.Result.Status != nil &&
		*describeServiceReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribeService failure(%v): %v",
			*describeServiceReply.Result.Status, describeServiceReply.Result.StatusDetail)
	}
	ret := &pb.DescribeServiceReply{
		Result: &pb.ServiceRunnerResult{
			Status:       0,
			StatusDetail: "",
		},
		Desc: &pb.ServiceDescription{
			SvcHeader: &pb.ServiceHeader{
				ProjEnvHeader: &pb.ProjEnvHeader{
					ProjId: describeServiceReply.Desc.SvcHeader.ProjEnvHeader.ProjID,
					EnvId:  describeServiceReply.Desc.SvcHeader.ProjEnvHeader.EnvID,
				},
				ServiceName: describeServiceReply.Desc.SvcHeader.ServiceName,
			},
			RpcEndpoints:   describeServiceReply.Desc.RPCEndpoints,
			DatabaseNames:  describeServiceReply.Desc.DatabaseNames,
			DatastoreNames: describeServiceReply.Desc.DatastoreNames,
			ApiDef:         describeServiceReply.Desc.APIDef,
			Port:           convertRESTIntToInt(describeServiceReply.Desc.Port),
		},
	}

	return ret, nil
}

func ListDatabases(projId string, envId string,
	opts ...DeployOption) ([]string, error) {

	deployOpts := fillDeployOptions(opts...)

	listDatabasesReq := &models.ListDatabasesRequest{
		Header: &models.ProjEnvHeader{
			ProjID: projId,
			EnvID:  envId,
		},
	}
	listDatabasesParams := service_runner.NewListDatabasesParams().
		WithBody(listDatabasesReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.ListDatabasesOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.ListDatabases(listDatabasesParams,
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
			fmt.Errorf("ListDatabases failure(%v): %v",
				*listReply.Result.Status, listReply.Result.StatusDetail)
	}

	return listReply.DatabaseNames, nil
}

func DescribeDatabase(projId string, envId string, dbName string,
	opts ...DeployOption) (*pb.DescribeDatabaseReply, error) {

	deployOpts := fillDeployOptions(opts...)

	describeDatabaseReq := &models.DescribeDatabaseRequest{
		DatabaseHeader: &models.DatabaseHeader{
			ProjEnvHeader: &models.ProjEnvHeader{
				ProjID: projId,
				EnvID:  envId,
			},
			DatabaseName: dbName,
		},
	}

	describeDatabaseParams := service_runner.NewDescribeDatabaseParams().
		WithBody(describeDatabaseReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.DescribeDatabaseOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeDatabase(describeDatabaseParams,
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
	describeDatabaseReply := resp.GetPayload()

	if describeDatabaseReply.Result.Status != nil &&
		*describeDatabaseReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribeDatabase failure(%v): %v",
			*describeDatabaseReply.Result.Status, describeDatabaseReply.Result.StatusDetail)
	}
	ret := &pb.DescribeDatabaseReply{
		Result: &pb.ServiceRunnerResult{
			Status:       0,
			StatusDetail: "",
		},
		Desc: &pb.DatabaseDescription{
			DatabaseHeader: &pb.DatabaseHeader{
				ProjEnvHeader: &pb.ProjEnvHeader{
					ProjId: describeDatabaseReply.Desc.DatabaseHeader.ProjEnvHeader.ProjID,
					EnvId:  describeDatabaseReply.Desc.DatabaseHeader.ProjEnvHeader.EnvID,
				},
				DatabaseName: describeDatabaseReply.Desc.DatabaseHeader.DatabaseName,
			},
			Tables:       make([]*pb.DatabaseTableDescription, 0),
			ServiceNames: describeDatabaseReply.Desc.ServiceNames,
		},
	}

	for _, tbl := range describeDatabaseReply.Desc.Tables {
		tbl2Append := &pb.DatabaseTableDescription{
			Name:         tbl.Name,
			NumRows:      convertRESTIntToInt(tbl.NumRows),
			Size:         convertRESTIntToInt(tbl.Size),
			ServiceNames: tbl.ServiceNames,
		}
		ret.Desc.Tables = append(ret.Desc.Tables, tbl2Append)
	}

	return ret, nil
}

func ListDatastores(projId string, envId string,
	opts ...DeployOption) ([]string, error) {

	deployOpts := fillDeployOptions(opts...)

	listDatastoresReq := &models.ListDatastoresRequest{
		Header: &models.ProjEnvHeader{
			ProjID: projId,
			EnvID:  envId,
		},
	}
	listDatastoresParams := service_runner.NewListDatastoresParams().
		WithBody(listDatastoresReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var err error
	var resp *service_runner.ListDatastoresOK

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.ListDatastores(listDatastoresParams,
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
			fmt.Errorf("ListDatastores failure(%v): %v",
				*listReply.Result.Status, listReply.Result.StatusDetail)
	}

	return listReply.DatastoreNames, nil
}

func DescribeDatastore(projId string, envId string, dstoreName string,
	opts ...DeployOption) (*pb.DescribeDatastoreReply, error) {

	deployOpts := fillDeployOptions(opts...)

	describeDatastoreReq := &models.DescribeDatastoreRequest{
		DatastoreHeader: &models.DatastoreHeader{
			ProjEnvHeader: &models.ProjEnvHeader{
				ProjID: projId,
				EnvID:  envId,
			},
			DatastoreName: dstoreName,
		},
	}

	describeDatastoreParams := service_runner.NewDescribeDatastoreParams().
		WithBody(describeDatastoreReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.DescribeDatastoreOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeDatastore(describeDatastoreParams,
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
	describeDatastoreReply := resp.GetPayload()

	if describeDatastoreReply.Result.Status != nil &&
		*describeDatastoreReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribeDatastore failure(%v): %v",
			*describeDatastoreReply.Result.Status, describeDatastoreReply.Result.StatusDetail)
	}
	ret := &pb.DescribeDatastoreReply{
		Result: &pb.ServiceRunnerResult{
			Status:       0,
			StatusDetail: "",
		},
		Desc: &pb.DatastoreDescription{
			DatastoreHeader: &pb.DatastoreHeader{
				ProjEnvHeader: &pb.ProjEnvHeader{
					ProjId: describeDatastoreReply.Desc.DatastoreHeader.ProjEnvHeader.ProjID,
					EnvId:  describeDatastoreReply.Desc.DatastoreHeader.ProjEnvHeader.EnvID,
				},
				DatastoreName: describeDatastoreReply.Desc.DatastoreHeader.DatastoreName,
			},
			NumObjects:              convertRESTIntToInt(describeDatastoreReply.Desc.NumObjects),
			CapacityConsumedInBytes: convertRESTIntToInt(describeDatastoreReply.Desc.CapacityConsumedInBytes),
			ServiceNames:            describeDatastoreReply.Desc.ServiceNames,
		},
	}

	return ret, nil
}

func DescribeSite(projId string, envId string,
	opts ...DeployOption) (*pb.DescribeSiteReply, error) {

	deployOpts := fillDeployOptions(opts...)

	describeSiteReq := &models.DescribeSiteRequest{
		ProjEnvHeader: &models.ProjEnvHeader{
			ProjID: projId,
			EnvID:  envId,
		},
	}

	describeSiteParams := service_runner.NewDescribeSiteParams().
		WithBody(describeSiteReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	var resp *service_runner.DescribeSiteOK

	var err error
	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.DescribeSite(describeSiteParams,
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
	describeSiteReply := resp.GetPayload()

	if describeSiteReply.Result.Status != nil &&
		*describeSiteReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return nil, fmt.Errorf("DescribeSite failure(%v): %v",
			*describeSiteReply.Result.Status, describeSiteReply.Result.StatusDetail)
	}
	ret := &pb.DescribeSiteReply{
		Result: &pb.ServiceRunnerResult{
			Status:       0,
			StatusDetail: "",
		},
		ProjEnvHeader: &pb.ProjEnvHeader{
			ProjId: describeSiteReply.ProjEnvHeader.ProjID,
			EnvId:  describeSiteReply.ProjEnvHeader.EnvID,
		},
		SiteEndpoint: describeSiteReply.SiteEndpoint,
	}

	return ret, nil
}
