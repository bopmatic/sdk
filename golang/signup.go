/* Copyright © 2024 Bopmatic, LLC. All Rights Reserved.
 *
 * See LICENSE file at the root of this package for license terms
 */
package golang

import (
	"fmt"
	"time"

	"github.com/bopmatic/sdk/golang/goswag_signup"
	"github.com/bopmatic/sdk/golang/goswag_signup/sign_up"
	"github.com/bopmatic/sdk/golang/models"
)

func RequestAccess(firstName string, lastName string, email string,
	programmingLang string, projectDesc string,
	opts ...DeployOption) error {

	deployOpts := fillDeployOptions(opts...)

	reqAccessReq := &models.RequestAccessRequest{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		ProgrammingLang: programmingLang,
		ProjectDesc:     projectDesc,
	}

	reqAccessParams := sign_up.NewRequestAccessParams().
		WithBody(reqAccessReq).WithHTTPClient(deployOpts.httpClient)
	config := goswag_signup.DefaultTransportConfig()
	client := goswag_signup.NewHTTPClientWithConfig(nil, config)

	var err error

	for retries := defaultRetries; retries > 0; retries-- {
		_, err = client.SignUp.RequestAccess(reqAccessParams)
		if err == nil {
			break
		}

		client = nil
		client = goswag_signup.NewHTTPClientWithConfig(nil, config)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("Client/HTTP failure: %v", err)
	}

	return nil
}
