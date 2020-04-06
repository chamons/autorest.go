// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ReadonlypropertyOperations contains the methods for the Readonlyproperty group.
type ReadonlypropertyOperations interface {
	// GetValid - Get complex types that have readonly properties
	GetValid(ctx context.Context) (*ReadonlyObjResponse, error)
	// PutValid - Put complex types that have readonly properties
	PutValid(ctx context.Context, complexBody ReadonlyObj) (*http.Response, error)
}

// readonlypropertyOperations implements the ReadonlypropertyOperations interface.
type readonlypropertyOperations struct {
	*Client
}

// GetValid - Get complex types that have readonly properties
func (client *readonlypropertyOperations) GetValid(ctx context.Context) (*ReadonlyObjResponse, error) {
	req, err := client.getValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getValidCreateRequest creates the GetValid request.
func (client *readonlypropertyOperations) getValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *readonlypropertyOperations) getValidHandleResponse(resp *azcore.Response) (*ReadonlyObjResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ReadonlyObjResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ReadonlyObj)
}

// PutValid - Put complex types that have readonly properties
func (client *readonlypropertyOperations) PutValid(ctx context.Context, complexBody ReadonlyObj) (*http.Response, error) {
	req, err := client.putValidCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *readonlypropertyOperations) putValidCreateRequest(complexBody ReadonlyObj) (*azcore.Request, error) {
	urlPath := "/complex/readonlyproperty/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleResponse handles the PutValid response.
func (client *readonlypropertyOperations) putValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}
