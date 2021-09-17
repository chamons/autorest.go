//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DedicatedHostGroupsClient contains the methods for the DedicatedHostGroups group.
// Don't use this type directly, use NewDedicatedHostGroupsClient() instead.
type DedicatedHostGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDedicatedHostGroupsClient creates a new instance of DedicatedHostGroupsClient with the specified values.
func NewDedicatedHostGroupsClient(con *arm.Connection, subscriptionID string) *DedicatedHostGroupsClient {
	return &DedicatedHostGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update a dedicated host group. For details of Dedicated Host and Dedicated Host Groups please see Dedicated Host Documentation
// [https://go.microsoft.com/fwlink/?linkid=2082596]
// If the operation fails it returns a generic error.
func (client *DedicatedHostGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroup, options *DedicatedHostGroupsCreateOrUpdateOptions) (DedicatedHostGroupsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, parameters, options)
	if err != nil {
		return DedicatedHostGroupsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DedicatedHostGroupsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedHostGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroup, options *DedicatedHostGroupsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DedicatedHostGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (DedicatedHostGroupsCreateOrUpdateResponse, error) {
	result := DedicatedHostGroupsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DedicatedHostGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Delete a dedicated host group.
// If the operation fails it returns a generic error.
func (client *DedicatedHostGroupsClient) Delete(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsDeleteOptions) (DedicatedHostGroupsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, options)
	if err != nil {
		return DedicatedHostGroupsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DedicatedHostGroupsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DedicatedHostGroupsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedHostGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DedicatedHostGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Retrieves information about a dedicated host group.
// If the operation fails it returns a generic error.
func (client *DedicatedHostGroupsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsGetOptions) (DedicatedHostGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, options)
	if err != nil {
		return DedicatedHostGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostGroupsClient) getHandleResponse(resp *http.Response) (DedicatedHostGroupsGetResponse, error) {
	result := DedicatedHostGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DedicatedHostGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - Lists all of the dedicated host groups in the specified resource group. Use the nextLink property in the response to get the next
// page of dedicated host groups.
// If the operation fails it returns a generic error.
func (client *DedicatedHostGroupsClient) ListByResourceGroup(resourceGroupName string, options *DedicatedHostGroupsListByResourceGroupOptions) *DedicatedHostGroupsListByResourceGroupPager {
	return &DedicatedHostGroupsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DedicatedHostGroupsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DedicatedHostGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedHostGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DedicatedHostGroupsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedHostGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedHostGroupsListByResourceGroupResponse, error) {
	result := DedicatedHostGroupsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroupListResult); err != nil {
		return DedicatedHostGroupsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DedicatedHostGroupsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListBySubscription - Lists all of the dedicated host groups in the subscription. Use the nextLink property in the response to get the next page of dedicated
// host groups.
// If the operation fails it returns a generic error.
func (client *DedicatedHostGroupsClient) ListBySubscription(options *DedicatedHostGroupsListBySubscriptionOptions) *DedicatedHostGroupsListBySubscriptionPager {
	return &DedicatedHostGroupsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DedicatedHostGroupsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DedicatedHostGroupListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedHostGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DedicatedHostGroupsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/hostGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedHostGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedHostGroupsListBySubscriptionResponse, error) {
	result := DedicatedHostGroupsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroupListResult); err != nil {
		return DedicatedHostGroupsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *DedicatedHostGroupsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Update an dedicated host group.
// If the operation fails it returns a generic error.
func (client *DedicatedHostGroupsClient) Update(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroupUpdate, options *DedicatedHostGroupsUpdateOptions) (DedicatedHostGroupsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, parameters, options)
	if err != nil {
		return DedicatedHostGroupsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostGroupsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroupUpdate, options *DedicatedHostGroupsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *DedicatedHostGroupsClient) updateHandleResponse(resp *http.Response) (DedicatedHostGroupsUpdateResponse, error) {
	result := DedicatedHostGroupsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *DedicatedHostGroupsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
