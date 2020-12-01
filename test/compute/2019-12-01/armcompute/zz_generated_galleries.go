// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GalleriesClient contains the methods for the Galleries group.
// Don't use this type directly, use NewGalleriesClient() instead.
type GalleriesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGalleriesClient creates a new instance of GalleriesClient with the specified values.
func NewGalleriesClient(con *armcore.Connection, subscriptionID string) GalleriesClient {
	return GalleriesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client GalleriesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Create or update a Shared Image Gallery.
func (client GalleriesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery, options *GalleriesBeginCreateOrUpdateOptions) (GalleryPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, gallery, options)
	if err != nil {
		return GalleryPollerResponse{}, err
	}
	result := GalleryPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleriesClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return GalleryPollerResponse{}, err
	}
	poller := &galleryPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new GalleryPoller from the specified resume token.
// token - The value must come from a previous call to GalleryPoller.ResumeToken().
func (client GalleriesClient) ResumeCreateOrUpdate(token string) (GalleryPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleriesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a Shared Image Gallery.
func (client GalleriesClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery, options *GalleriesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, gallery, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client GalleriesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery, options *GalleriesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(gallery)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client GalleriesClient) createOrUpdateHandleResponse(resp *azcore.Response) (GalleryResponse, error) {
	result := GalleryResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Gallery)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client GalleriesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Delete a Shared Image Gallery.
func (client GalleriesClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, galleryName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleriesClient.Delete", "", resp, client.deleteHandleError)
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
func (client GalleriesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleriesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Delete a Shared Image Gallery.
func (client GalleriesClient) delete(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, options)
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
func (client GalleriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client GalleriesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves information about a Shared Image Gallery.
func (client GalleriesClient) Get(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesGetOptions) (GalleryResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, options)
	if err != nil {
		return GalleryResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return GalleryResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return GalleryResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return GalleryResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client GalleriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
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

// getHandleResponse handles the Get response.
func (client GalleriesClient) getHandleResponse(resp *azcore.Response) (GalleryResponse, error) {
	result := GalleryResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Gallery)
	return result, err
}

// getHandleError handles the Get error response.
func (client GalleriesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List galleries under a subscription.
func (client GalleriesClient) List(options *GalleriesListOptions) GalleryListPager {
	return &galleryListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp GalleryListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client GalleriesClient) listCreateRequest(ctx context.Context, options *GalleriesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries"
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
func (client GalleriesClient) listHandleResponse(resp *azcore.Response) (GalleryListResponse, error) {
	result := GalleryListResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.GalleryList)
	return result, err
}

// listHandleError handles the List error response.
func (client GalleriesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - List galleries under a resource group.
func (client GalleriesClient) ListByResourceGroup(resourceGroupName string, options *GalleriesListByResourceGroupOptions) GalleryListPager {
	return &galleryListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp GalleryListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client GalleriesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GalleriesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client GalleriesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (GalleryListResponse, error) {
	result := GalleryListResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.GalleryList)
	return result, err
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client GalleriesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdate - Update a Shared Image Gallery.
func (client GalleriesClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate, options *GalleriesBeginUpdateOptions) (GalleryPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, galleryName, gallery, options)
	if err != nil {
		return GalleryPollerResponse{}, err
	}
	result := GalleryPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleriesClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return GalleryPollerResponse{}, err
	}
	poller := &galleryPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new GalleryPoller from the specified resume token.
// token - The value must come from a previous call to GalleryPoller.ResumeToken().
func (client GalleriesClient) ResumeUpdate(token string) (GalleryPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleriesClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update a Shared Image Gallery.
func (client GalleriesClient) update(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate, options *GalleriesBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, gallery, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client GalleriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate, options *GalleriesBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(gallery)
}

// updateHandleResponse handles the Update response.
func (client GalleriesClient) updateHandleResponse(resp *azcore.Response) (GalleryResponse, error) {
	result := GalleryResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Gallery)
	return result, err
}

// updateHandleError handles the Update error response.
func (client GalleriesClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
