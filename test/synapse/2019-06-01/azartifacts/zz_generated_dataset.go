// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type datasetClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client datasetClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CreateOrUpdateDataset - Creates or updates a dataset.
func (client datasetClient) createOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetBeginCreateOrUpdateDatasetOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateDatasetCreateRequest(ctx, datasetName, dataset, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateDatasetHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateDatasetCreateRequest creates the CreateOrUpdateDataset request.
func (client datasetClient) createOrUpdateDatasetCreateRequest(ctx context.Context, datasetName string, dataset DatasetResource, options *DatasetBeginCreateOrUpdateDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dataset)
}

// createOrUpdateDatasetHandleResponse handles the CreateOrUpdateDataset response.
func (client datasetClient) createOrUpdateDatasetHandleResponse(resp *azcore.Response) (DatasetResourceResponse, error) {
	result := DatasetResourceResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.DatasetResource)
	return result, err
}

// createOrUpdateDatasetHandleError handles the CreateOrUpdateDataset error response.
func (client datasetClient) createOrUpdateDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteDataset - Deletes a dataset.
func (client datasetClient) deleteDataset(ctx context.Context, datasetName string, options *DatasetBeginDeleteDatasetOptions) (*azcore.Response, error) {
	req, err := client.deleteDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteDatasetHandleError(resp)
	}
	return resp, nil
}

// deleteDatasetCreateRequest creates the DeleteDataset request.
func (client datasetClient) deleteDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetBeginDeleteDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteDatasetHandleError handles the DeleteDataset error response.
func (client datasetClient) deleteDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDataset - Gets a dataset.
func (client datasetClient) GetDataset(ctx context.Context, datasetName string, options *DatasetGetDatasetOptions) (DatasetResourceResponse, error) {
	req, err := client.getDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return DatasetResourceResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return DatasetResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return DatasetResourceResponse{}, client.getDatasetHandleError(resp)
	}
	result, err := client.getDatasetHandleResponse(resp)
	if err != nil {
		return DatasetResourceResponse{}, err
	}
	return result, nil
}

// getDatasetCreateRequest creates the GetDataset request.
func (client datasetClient) getDatasetCreateRequest(ctx context.Context, datasetName string, options *DatasetGetDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDatasetHandleResponse handles the GetDataset response.
func (client datasetClient) getDatasetHandleResponse(resp *azcore.Response) (DatasetResourceResponse, error) {
	result := DatasetResourceResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.DatasetResource)
	return result, err
}

// getDatasetHandleError handles the GetDataset error response.
func (client datasetClient) getDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDatasetsByWorkspace - Lists datasets.
func (client datasetClient) GetDatasetsByWorkspace(options *DatasetGetDatasetsByWorkspaceOptions) DatasetListResponsePager {
	return &datasetListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getDatasetsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getDatasetsByWorkspaceHandleResponse,
		errorer:   client.getDatasetsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp DatasetListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DatasetListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getDatasetsByWorkspaceCreateRequest creates the GetDatasetsByWorkspace request.
func (client datasetClient) getDatasetsByWorkspaceCreateRequest(ctx context.Context, options *DatasetGetDatasetsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/datasets"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDatasetsByWorkspaceHandleResponse handles the GetDatasetsByWorkspace response.
func (client datasetClient) getDatasetsByWorkspaceHandleResponse(resp *azcore.Response) (DatasetListResponseResponse, error) {
	result := DatasetListResponseResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.DatasetListResponse)
	return result, err
}

// getDatasetsByWorkspaceHandleError handles the GetDatasetsByWorkspace error response.
func (client datasetClient) getDatasetsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameDataset - Renames a dataset.
func (client datasetClient) renameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetBeginRenameDatasetOptions) (*azcore.Response, error) {
	req, err := client.renameDatasetCreateRequest(ctx, datasetName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameDatasetHandleError(resp)
	}
	return resp, nil
}

// renameDatasetCreateRequest creates the RenameDataset request.
func (client datasetClient) renameDatasetCreateRequest(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetBeginRenameDatasetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameDatasetHandleError handles the RenameDataset error response.
func (client datasetClient) renameDatasetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
