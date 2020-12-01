// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ImagesClient contains the methods for the Images group.
// Don't use this type directly, use NewImagesClient() instead.
type ImagesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewImagesClient creates a new instance of ImagesClient with the specified values.
func NewImagesClient(con *armcore.Connection, subscriptionID string) ImagesClient {
	return ImagesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client ImagesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Create or update an image.
func (client ImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesBeginCreateOrUpdateOptions) (ImagePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return ImagePollerResponse{}, err
	}
	result := ImagePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ImagesClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ImagePollerResponse{}, err
	}
	poller := &imagePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ImageResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ImagePoller from the specified resume token.
// token - The value must come from a previous call to ImagePoller.ResumeToken().
func (client ImagesClient) ResumeCreateOrUpdate(token string) (ImagePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ImagesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &imagePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update an image.
func (client ImagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client ImagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client ImagesClient) createOrUpdateHandleResponse(resp *azcore.Response) (ImageResponse, error) {
	result := ImageResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Image)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client ImagesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes an Image.
func (client ImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, imageName string, options *ImagesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ImagesClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client ImagesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ImagesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes an Image.
func (client ImagesClient) delete(ctx context.Context, resourceGroupName string, imageName string, options *ImagesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client ImagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, imageName string, options *ImagesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client ImagesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets an image.
func (client ImagesClient) Get(ctx context.Context, resourceGroupName string, imageName string, options *ImagesGetOptions) (ImageResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return ImageResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ImageResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ImageResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return ImageResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client ImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, imageName string, options *ImagesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client ImagesClient) getHandleResponse(resp *azcore.Response) (ImageResponse, error) {
	result := ImageResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Image)
	return result, err
}

// getHandleError handles the Get error response.
func (client ImagesClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Gets the list of Images in the subscription. Use nextLink property in the response to get the next page of Images. Do this till nextLink is null
// to fetch all the Images.
func (client ImagesClient) List(options *ImagesListOptions) ImageListResultPager {
	return &imageListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ImageListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ImageListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client ImagesClient) listCreateRequest(ctx context.Context, options *ImagesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client ImagesClient) listHandleResponse(resp *azcore.Response) (ImageListResultResponse, error) {
	result := ImageListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ImageListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client ImagesClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Gets the list of images under a resource group.
func (client ImagesClient) ListByResourceGroup(resourceGroupName string, options *ImagesListByResourceGroupOptions) ImageListResultPager {
	return &imageListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ImageListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ImageListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client ImagesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ImagesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client ImagesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ImageListResultResponse, error) {
	result := ImageListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ImageListResult)
	return result, err
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client ImagesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Update an image.
func (client ImagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesBeginUpdateOptions) (ImagePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return ImagePollerResponse{}, err
	}
	result := ImagePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ImagesClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return ImagePollerResponse{}, err
	}
	poller := &imagePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ImageResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ImagePoller from the specified resume token.
// token - The value must come from a previous call to ImagePoller.ResumeToken().
func (client ImagesClient) ResumeUpdate(token string) (ImagePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ImagesClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &imagePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update an image.
func (client ImagesClient) update(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client ImagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client ImagesClient) updateHandleResponse(resp *azcore.Response) (ImageResponse, error) {
	result := ImageResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Image)
	return result, err
}

// updateHandleError handles the Update error response.
func (client ImagesClient) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
