/* Copyright © 2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package util

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	grpcb "google.golang.org/grpc/backoff"
	grpcc "google.golang.org/grpc/connectivity"

	dapr "github.com/dapr/go-sdk/client"
)

// NewDaprClient instantiates a Dapr client connected to the default dapr port
// on localhost. This is an alternative to dapr.NewClient() which has a bug
// that will DoS callers if the server isn't up yet. Additionally this version
// replaces the default GRPC exponential backoff strategy that defaults to
// O(1s) delays with a linear backoff strategy with O(10ms) delays. For our use
// case we're only ever connecting to localhost and don't want to induce
// unnecessary latency as this can potentially translate into end-user latency
// when invoked during an AWS Lambda cold start
func NewDaprClient(ctx context.Context) (dapr.Client, error) {

	const DefaultDaprHost = "127.0.0.1"
	const DefaultDaprPort = "50001"
	target := fmt.Sprintf("%v:%v", DefaultDaprHost, DefaultDaprPort)

	var grpcConn *grpc.ClientConn
	var err error

	connectParams := grpc.ConnectParams{
		Backoff: grpcb.Config{
			BaseDelay:  10.0 * time.Millisecond,
			Multiplier: 1.0,
			Jitter:     0.0,
			MaxDelay:   10.0 * time.Millisecond,
		},
		MinConnectTimeout: 10 * time.Second,
	}

	// https not needed because we are connecting to localhost
	grpcConn, err = grpc.NewClient(target, grpc.WithInsecure(),
		grpc.WithIdleTimeout(0), grpc.WithConnectParams(connectParams))
	if err != nil {
		return nil, fmt.Errorf("Failed to create dapr client for %v: %w",
			target, err)
	}
	grpcConn.Connect()
	err = waitForGrpcConnect(ctx, grpcConn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect dapr client for %v: %w",
			target, err)
	}

	return dapr.NewClientWithConnection(grpcConn), nil
}

func waitForGrpcConnect(ctx context.Context, grpcConn *grpc.ClientConn) error {
	for {
		curState := grpcConn.GetState()
		if curState == grpcc.Ready {
			break
		}

		select {
		case <-ctx.Done():
			return context.DeadlineExceeded
		default:
			grpcConn.WaitForStateChange(ctx, curState)
		}
	}

	return nil
}
