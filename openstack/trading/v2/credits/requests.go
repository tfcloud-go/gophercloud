// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package credits

import (
	"net/http"

	"github.com/gophercloud/gophercloud"
)

type BindPayerOpts struct {
	ProjectID string
}

func BindPayer(client *gophercloud.ServiceClient, creditID string, opts BindPayerOpts) (r BindPayerResult) {
	_, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		r.Err = err
		return
	}

	resp, err := client.Put(bindPayerURL(client, creditID, opts.ProjectID), nil, nil, &gophercloud.RequestOpts{
		OkCodes: []int{http.StatusNoContent},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
