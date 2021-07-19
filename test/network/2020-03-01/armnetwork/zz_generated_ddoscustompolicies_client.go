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
	"time"
)

// DdosCustomPoliciesClient contains the methods for the DdosCustomPolicies group.
// Don't use this type directly, use NewDdosCustomPoliciesClient() instead.
type DdosCustomPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDdosCustomPoliciesClient creates a new instance of DdosCustomPoliciesClient with the specified values.
func NewDdosCustomPoliciesClient(con *armcore.Connection, subscriptionID string) *DdosCustomPoliciesClient {
	return &DdosCustomPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a DDoS custom policy.
// If the operation fails it returns the *CloudError error type.
func (client *DdosCustomPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy, options *DdosCustomPoliciesBeginCreateOrUpdateOptions) (DdosCustomPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ddosCustomPolicyName, parameters, options)
	if err != nil {
		return DdosCustomPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := DdosCustomPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("DdosCustomPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return DdosCustomPoliciesCreateOrUpdatePollerResponse{}, err
	}
	poller := &ddosCustomPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DdosCustomPoliciesCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new DdosCustomPoliciesCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to DdosCustomPoliciesCreateOrUpdatePoller.ResumeToken().
func (client *DdosCustomPoliciesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (DdosCustomPoliciesCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("DdosCustomPoliciesClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return DdosCustomPoliciesCreateOrUpdatePollerResponse{}, err
	}
	poller := &ddosCustomPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DdosCustomPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := DdosCustomPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DdosCustomPoliciesCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a DDoS custom policy.
// If the operation fails it returns the *CloudError error type.
func (client *DdosCustomPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy, options *DdosCustomPoliciesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DdosCustomPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy, options *DdosCustomPoliciesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DdosCustomPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified DDoS custom policy.
// If the operation fails it returns the *CloudError error type.
func (client *DdosCustomPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesBeginDeleteOptions) (DdosCustomPoliciesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ddosCustomPolicyName, options)
	if err != nil {
		return DdosCustomPoliciesDeletePollerResponse{}, err
	}
	result := DdosCustomPoliciesDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("DdosCustomPoliciesClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return DdosCustomPoliciesDeletePollerResponse{}, err
	}
	poller := &ddosCustomPoliciesDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DdosCustomPoliciesDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new DdosCustomPoliciesDeletePoller from the specified resume token.
// token - The value must come from a previous call to DdosCustomPoliciesDeletePoller.ResumeToken().
func (client *DdosCustomPoliciesClient) ResumeDelete(ctx context.Context, token string) (DdosCustomPoliciesDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("DdosCustomPoliciesClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return DdosCustomPoliciesDeletePollerResponse{}, err
	}
	poller := &ddosCustomPoliciesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DdosCustomPoliciesDeletePollerResponse{}, err
	}
	result := DdosCustomPoliciesDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DdosCustomPoliciesDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified DDoS custom policy.
// If the operation fails it returns the *CloudError error type.
func (client *DdosCustomPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DdosCustomPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client *DdosCustomPoliciesClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets information about the specified DDoS custom policy.
// If the operation fails it returns the *CloudError error type.
func (client *DdosCustomPoliciesClient) Get(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesGetOptions) (DdosCustomPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, options)
	if err != nil {
		return DdosCustomPoliciesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DdosCustomPoliciesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DdosCustomPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DdosCustomPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, options *DdosCustomPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
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

// getHandleResponse handles the Get response.
func (client *DdosCustomPoliciesClient) getHandleResponse(resp *azcore.Response) (DdosCustomPoliciesGetResponse, error) {
	result := DdosCustomPoliciesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DdosCustomPolicy); err != nil {
		return DdosCustomPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DdosCustomPoliciesClient) getHandleError(resp *azcore.Response) error {
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

// UpdateTags - Update a DDoS custom policy tags.
// If the operation fails it returns the *CloudError error type.
func (client *DdosCustomPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject, options *DdosCustomPoliciesUpdateTagsOptions) (DdosCustomPoliciesUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ddosCustomPolicyName, parameters, options)
	if err != nil {
		return DdosCustomPoliciesUpdateTagsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DdosCustomPoliciesUpdateTagsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DdosCustomPoliciesUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DdosCustomPoliciesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject, options *DdosCustomPoliciesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosCustomPolicyName == "" {
		return nil, errors.New("parameter ddosCustomPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DdosCustomPoliciesClient) updateTagsHandleResponse(resp *azcore.Response) (DdosCustomPoliciesUpdateTagsResponse, error) {
	result := DdosCustomPoliciesUpdateTagsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DdosCustomPolicy); err != nil {
		return DdosCustomPoliciesUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *DdosCustomPoliciesClient) updateTagsHandleError(resp *azcore.Response) error {
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
