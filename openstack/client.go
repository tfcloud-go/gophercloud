// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package openstack

import "github.com/gophercloud/gophercloud"

func initClientOpts(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}

func NewTradingV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "trading")
}

func NewMessagingV2(client *gophercloud.ProviderClient, clientID string, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "messaging")
	sc.MoreHeaders = map[string]string{"Client-ID": clientID}
	return sc, err
}
