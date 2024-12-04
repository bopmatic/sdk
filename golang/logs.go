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
)

// GetLogs() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
//
// svcName can be left blank to retrieve logs for all services
func GetLogs(projId string, envId string, svcName string, startTime time.Time,
	endTime time.Time, opts ...DeployOption) error {

	deployOpts := fillDeployOptions(opts...)

	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	getLogsReq := &models.GetLogsRequest{
		ProjID:      projId,
		EnvID:       envId,
		ServiceName: svcName,
		StartTime:   strconv.FormatInt(startTime.UnixMilli(), 10),
		EndTime:     strconv.FormatInt(endTime.UnixMilli(), 10),
	}

	getLogsParams := service_runner.NewGetLogsParams().
		WithBody(getLogsReq).WithHTTPClient(deployOpts.httpClient)
	var resp *service_runner.GetLogsOK
	var err error

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.GetLogs(getLogsParams, deployOpts)
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
	getLogsReply := resp.GetPayload()
	if getLogsReply.Result.Status != nil &&
		*getLogsReply.Result.Status != models.ServiceRunnerStatusSTATUSOK {
		return fmt.Errorf("GetLogs failure(%v): %v",
			*getLogsReply.Result.Status, getLogsReply.Result.StatusDetail)
	}

	for _, entry := range getLogsReply.Entries {
		var timeStr = "<unknown_time>"

		msecs, err := strconv.ParseInt(entry.Timestamp, 10, 64)
		if err == nil {
			timeStr = fmt.Sprintf("%v", time.UnixMilli(msecs).UTC())
		}

		fmt.Fprintf(deployOpts.output, "%v: %v", timeStr, entry.Message)
	}

	return nil
}
