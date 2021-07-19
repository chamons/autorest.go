// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// NetworkInterfaceLoadBalancersClient contains the methods for the NetworkInterfaceLoadBalancers group.
// Don't use this type directly, use NewNetworkInterfaceLoadBalancersClient() instead.
type NetworkInterfaceLoadBalancersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkInterfaceLoadBalancersClient creates a new instance of NetworkInterfaceLoadBalancersClient with the specified values.
func NewNetworkInterfaceLoadBalancersClient(con *armcore.Connection, subscriptionID string) *NetworkInterfaceLoadBalancersClient {
	return &NetworkInterfaceLoadBalancersClient{con: con, subscriptionID: subscriptionID}
}

// List - List all load balancers in a network interface.
// If the operation fails it returns the *CloudError error type.
func (client *NetworkInterfaceLoadBalancersClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceLoadBalancersListOptions) NetworkInterfaceLoadBalancersListPager {
	return &networkInterfaceLoadBalancersListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		advancer: func(ctx context.Context, resp NetworkInterfaceLoadBalancersListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceLoadBalancerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *NetworkInterfaceLoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceLoadBalancersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkInterfaceLoadBalancersClient) listHandleResponse(resp *azcore.Response) (NetworkInterfaceLoadBalancersListResponse, error) {
	result := NetworkInterfaceLoadBalancersListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NetworkInterfaceLoadBalancerListResult); err != nil {
		return NetworkInterfaceLoadBalancersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *NetworkInterfaceLoadBalancersClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
