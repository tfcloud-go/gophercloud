// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package queues

import "github.com/gophercloud/gophercloud"

const (
	apiVersion = "v2"
	apiName    = "queues"
)

func getSubscriptionURL(client *gophercloud.ServiceClient, queue, subscriptionID string) string {
	return client.ServiceURL(apiVersion, apiName, queue, "subscriptions", subscriptionID)
}
