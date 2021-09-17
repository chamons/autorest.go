//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// ODataClient contains the methods for the OData group.
// Don't use this type directly, use NewODataClient() instead.
type ODataClient struct {
	con *Connection
}

// NewODataClient creates a new instance of ODataClient with the specified values.
func NewODataClient(con *Connection) *ODataClient {
	return &ODataClient{con: con}
}

// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
// If the operation fails it returns the *Error error type.
func (client *ODataClient) GetWithFilter(ctx context.Context, options *ODataGetWithFilterOptions) (ODataGetWithFilterResponse, error) {
	req, err := client.getWithFilterCreateRequest(ctx, options)
	if err != nil {
		return ODataGetWithFilterResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ODataGetWithFilterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ODataGetWithFilterResponse{}, client.getWithFilterHandleError(resp)
	}
	return ODataGetWithFilterResponse{RawResponse: resp}, nil
}

// getWithFilterCreateRequest creates the GetWithFilter request.
func (client *ODataClient) getWithFilterCreateRequest(ctx context.Context, options *ODataGetWithFilterOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/odata/filter"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWithFilterHandleError handles the GetWithFilter error response.
func (client *ODataClient) getWithFilterHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
