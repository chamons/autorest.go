// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type dataFlowClient struct {
	con *connection
}

// BeginCreateOrUpdateDataFlow - Creates or updates a data flow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) BeginCreateOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowBeginCreateOrUpdateDataFlowOptions) (DataFlowCreateOrUpdateDataFlowPollerResponse, error) {
	resp, err := client.createOrUpdateDataFlow(ctx, dataFlowName, dataFlow, options)
	if err != nil {
		return DataFlowCreateOrUpdateDataFlowPollerResponse{}, err
	}
	result := DataFlowCreateOrUpdateDataFlowPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("dataFlowClient.CreateOrUpdateDataFlow", resp, client.con.Pipeline(), client.createOrUpdateDataFlowHandleError)
	if err != nil {
		return DataFlowCreateOrUpdateDataFlowPollerResponse{}, err
	}
	poller := &dataFlowCreateOrUpdateDataFlowPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowCreateOrUpdateDataFlowResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdateDataFlow creates a new DataFlowCreateOrUpdateDataFlowPoller from the specified resume token.
// token - The value must come from a previous call to DataFlowCreateOrUpdateDataFlowPoller.ResumeToken().
func (client *dataFlowClient) ResumeCreateOrUpdateDataFlow(ctx context.Context, token string) (DataFlowCreateOrUpdateDataFlowPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("dataFlowClient.CreateOrUpdateDataFlow", token, client.con.Pipeline(), client.createOrUpdateDataFlowHandleError)
	if err != nil {
		return DataFlowCreateOrUpdateDataFlowPollerResponse{}, err
	}
	poller := &dataFlowCreateOrUpdateDataFlowPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DataFlowCreateOrUpdateDataFlowPollerResponse{}, err
	}
	result := DataFlowCreateOrUpdateDataFlowPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowCreateOrUpdateDataFlowResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdateDataFlow - Creates or updates a data flow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) createOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowBeginCreateOrUpdateDataFlowOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateDataFlowCreateRequest(ctx, dataFlowName, dataFlow, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateDataFlowHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateDataFlowCreateRequest creates the CreateOrUpdateDataFlow request.
func (client *dataFlowClient) createOrUpdateDataFlowCreateRequest(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowBeginCreateOrUpdateDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dataFlow)
}

// createOrUpdateDataFlowHandleError handles the CreateOrUpdateDataFlow error response.
func (client *dataFlowClient) createOrUpdateDataFlowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDeleteDataFlow - Deletes a data flow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) BeginDeleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowBeginDeleteDataFlowOptions) (DataFlowDeleteDataFlowPollerResponse, error) {
	resp, err := client.deleteDataFlow(ctx, dataFlowName, options)
	if err != nil {
		return DataFlowDeleteDataFlowPollerResponse{}, err
	}
	result := DataFlowDeleteDataFlowPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("dataFlowClient.DeleteDataFlow", resp, client.con.Pipeline(), client.deleteDataFlowHandleError)
	if err != nil {
		return DataFlowDeleteDataFlowPollerResponse{}, err
	}
	poller := &dataFlowDeleteDataFlowPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowDeleteDataFlowResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteDataFlow creates a new DataFlowDeleteDataFlowPoller from the specified resume token.
// token - The value must come from a previous call to DataFlowDeleteDataFlowPoller.ResumeToken().
func (client *dataFlowClient) ResumeDeleteDataFlow(ctx context.Context, token string) (DataFlowDeleteDataFlowPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("dataFlowClient.DeleteDataFlow", token, client.con.Pipeline(), client.deleteDataFlowHandleError)
	if err != nil {
		return DataFlowDeleteDataFlowPollerResponse{}, err
	}
	poller := &dataFlowDeleteDataFlowPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DataFlowDeleteDataFlowPollerResponse{}, err
	}
	result := DataFlowDeleteDataFlowPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowDeleteDataFlowResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// DeleteDataFlow - Deletes a data flow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) deleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowBeginDeleteDataFlowOptions) (*azcore.Response, error) {
	req, err := client.deleteDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteDataFlowHandleError(resp)
	}
	return resp, nil
}

// deleteDataFlowCreateRequest creates the DeleteDataFlow request.
func (client *dataFlowClient) deleteDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowBeginDeleteDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteDataFlowHandleError handles the DeleteDataFlow error response.
func (client *dataFlowClient) deleteDataFlowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetDataFlow - Gets a data flow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) GetDataFlow(ctx context.Context, dataFlowName string, options *DataFlowGetDataFlowOptions) (DataFlowGetDataFlowResponse, error) {
	req, err := client.getDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return DataFlowGetDataFlowResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataFlowGetDataFlowResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DataFlowGetDataFlowResponse{}, client.getDataFlowHandleError(resp)
	}
	return client.getDataFlowHandleResponse(resp)
}

// getDataFlowCreateRequest creates the GetDataFlow request.
func (client *dataFlowClient) getDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowGetDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDataFlowHandleResponse handles the GetDataFlow response.
func (client *dataFlowClient) getDataFlowHandleResponse(resp *azcore.Response) (DataFlowGetDataFlowResponse, error) {
	result := DataFlowGetDataFlowResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataFlowResource); err != nil {
		return DataFlowGetDataFlowResponse{}, err
	}
	return result, nil
}

// getDataFlowHandleError handles the GetDataFlow error response.
func (client *dataFlowClient) getDataFlowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetDataFlowsByWorkspace - Lists data flows.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) GetDataFlowsByWorkspace(options *DataFlowGetDataFlowsByWorkspaceOptions) DataFlowGetDataFlowsByWorkspacePager {
	return &dataFlowGetDataFlowsByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getDataFlowsByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DataFlowGetDataFlowsByWorkspaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataFlowListResponse.NextLink)
		},
	}
}

// getDataFlowsByWorkspaceCreateRequest creates the GetDataFlowsByWorkspace request.
func (client *dataFlowClient) getDataFlowsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowGetDataFlowsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/dataflows"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDataFlowsByWorkspaceHandleResponse handles the GetDataFlowsByWorkspace response.
func (client *dataFlowClient) getDataFlowsByWorkspaceHandleResponse(resp *azcore.Response) (DataFlowGetDataFlowsByWorkspaceResponse, error) {
	result := DataFlowGetDataFlowsByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataFlowListResponse); err != nil {
		return DataFlowGetDataFlowsByWorkspaceResponse{}, err
	}
	return result, nil
}

// getDataFlowsByWorkspaceHandleError handles the GetDataFlowsByWorkspace error response.
func (client *dataFlowClient) getDataFlowsByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRenameDataFlow - Renames a dataflow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) BeginRenameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowBeginRenameDataFlowOptions) (DataFlowRenameDataFlowPollerResponse, error) {
	resp, err := client.renameDataFlow(ctx, dataFlowName, request, options)
	if err != nil {
		return DataFlowRenameDataFlowPollerResponse{}, err
	}
	result := DataFlowRenameDataFlowPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("dataFlowClient.RenameDataFlow", resp, client.con.Pipeline(), client.renameDataFlowHandleError)
	if err != nil {
		return DataFlowRenameDataFlowPollerResponse{}, err
	}
	poller := &dataFlowRenameDataFlowPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowRenameDataFlowResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRenameDataFlow creates a new DataFlowRenameDataFlowPoller from the specified resume token.
// token - The value must come from a previous call to DataFlowRenameDataFlowPoller.ResumeToken().
func (client *dataFlowClient) ResumeRenameDataFlow(ctx context.Context, token string) (DataFlowRenameDataFlowPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("dataFlowClient.RenameDataFlow", token, client.con.Pipeline(), client.renameDataFlowHandleError)
	if err != nil {
		return DataFlowRenameDataFlowPollerResponse{}, err
	}
	poller := &dataFlowRenameDataFlowPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DataFlowRenameDataFlowPollerResponse{}, err
	}
	result := DataFlowRenameDataFlowPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowRenameDataFlowResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// RenameDataFlow - Renames a dataflow.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowClient) renameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowBeginRenameDataFlowOptions) (*azcore.Response, error) {
	req, err := client.renameDataFlowCreateRequest(ctx, dataFlowName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameDataFlowHandleError(resp)
	}
	return resp, nil
}

// renameDataFlowCreateRequest creates the RenameDataFlow request.
func (client *dataFlowClient) renameDataFlowCreateRequest(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowBeginRenameDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}/rename"
	if dataFlowName == "" {
		return nil, errors.New("parameter dataFlowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameDataFlowHandleError handles the RenameDataFlow error response.
func (client *dataFlowClient) renameDataFlowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
