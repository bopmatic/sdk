/* Copyright © 2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package golang

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bopmatic/sdk/golang/goswag"
	"github.com/bopmatic/sdk/golang/goswag/service_runner"
	"github.com/bopmatic/sdk/golang/models"
)

type getLogsOptions struct {
	httpClient *http.Client
	output     io.Writer
}

type GetLogsOption func(*getLogsOptions) *getLogsOptions

func GetLogsOptHttpClient(httpClient *http.Client) GetLogsOption {
	return func(opts *getLogsOptions) *getLogsOptions {
		opts.httpClient = httpClient

		return opts
	}
}

func GetLogsOptOutput(output io.Writer) GetLogsOption {
	return func(opts *getLogsOptions) *getLogsOptions {
		opts.output = output

		return opts
	}
}

func fillGetLogsOptions(opts ...GetLogsOption) *getLogsOptions {
	options := &getLogsOptions{
		httpClient: http.DefaultClient,
		output:     os.Stdout,
	}
	for _, optApplyFunc := range opts {
		if optApplyFunc != nil {
			options = optApplyFunc(options)
		}
	}

	return options
}

// GetLogs() implemented using a client generated with go-swagger:
//
//	https://github.com/go-swagger/go-swagger
func GetLogs(projName string, svcName string, startTime time.Time,
	endTime time.Time, opts ...GetLogsOption) error {

	getLogsOpts := fillGetLogsOptions(opts...)

	config := goswag.DefaultTransportConfig()
	client := goswag.NewHTTPClientWithConfig(nil, config)

	getLogsReq := &models.GetLogsRequest{
		ProjectName: projName,
		ServiceName: svcName,
		StartTime:   strconv.FormatInt(startTime.Unix(), 10),
		EndTime:     strconv.FormatInt(endTime.Unix(), 10),
	}

	getLogsParams := service_runner.NewGetLogsParams().
		WithBody(getLogsReq).WithHTTPClient(getLogsOpts.httpClient)
	var resp *service_runner.GetLogsOK
	var err error

	for retries := defaultRetries; retries > 0; retries-- {
		resp, err = client.ServiceRunner.GetLogs(getLogsParams)
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

	for _, entry := range getLogsReply.Entries {
		var timeStr = "<unknown_time>"

		secs, err := strconv.ParseInt(entry.Timestamp, 10, 64)
		if err == nil {
			timeStr = fmt.Sprintf("%v", time.Unix(secs, 0).UTC())
		}

		fmt.Fprintf(getLogsOpts.output, "%v: %v", timeStr, entry.Message)
	}

	return nil
}
