// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package morecustombaseurigroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PathsClient contains the methods for the Paths group.
// Don't use this type directly, use NewPathsClient() instead.
type PathsClient struct {
	con            *Connection
	subscriptionID string
}

// NewPathsClient creates a new instance of PathsClient with the specified values.
func NewPathsClient(con *Connection, subscriptionID string) *PathsClient {
	return &PathsClient{con: con, subscriptionID: subscriptionID}
}

// GetEmpty - Get a 200 to test a valid base uri
// If the operation fails it returns the *Error error type.
func (client *PathsClient) GetEmpty(ctx context.Context, vault string, secret string, keyName string, options *PathsGetEmptyOptions) (PathsGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, vault, secret, keyName, options)
	if err != nil {
		return PathsGetEmptyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PathsGetEmptyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PathsGetEmptyResponse{}, client.getEmptyHandleError(resp)
	}
	return PathsGetEmptyResponse{RawResponse: resp.Response}, nil
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *PathsClient) getEmptyCreateRequest(ctx context.Context, vault string, secret string, keyName string, options *PathsGetEmptyOptions) (*azcore.Request, error) {
	host := "{vault}{secret}{dnsSuffix}"
	host = strings.ReplaceAll(host, "{dnsSuffix}", client.con.DnsSuffix())
	host = strings.ReplaceAll(host, "{vault}", vault)
	host = strings.ReplaceAll(host, "{secret}", secret)
	urlPath := "/customuri/{subscriptionId}/{keyName}"
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.KeyVersion != nil {
		reqQP.Set("keyVersion", *options.KeyVersion)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *PathsClient) getEmptyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
