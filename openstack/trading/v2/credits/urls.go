// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package credits

import "github.com/gophercloud/gophercloud"

const (
	creditPath = "credits"
)

func bindPayerURL(client *gophercloud.ServiceClient, creditID, projectID string) string {
	return client.ServiceURL(creditPath, creditID, "payers", projectID)
}
