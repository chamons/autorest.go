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

// PolymorphicrecursiveOperations contains the methods for the Polymorphicrecursive group.
type PolymorphicrecursiveOperations interface {
	// GetValid - Get complex types that are polymorphic and have recursive references
	GetValid(ctx context.Context) (*FishResponse, error)
	// PutValid - Put complex types that are polymorphic and have recursive references
	PutValid(ctx context.Context, complexBody Fish) (*http.Response, error)
}

// polymorphicrecursiveOperations implements the PolymorphicrecursiveOperations interface.
type polymorphicrecursiveOperations struct {
	*Client
}

// GetValid - Get complex types that are polymorphic and have recursive references
func (client *polymorphicrecursiveOperations) GetValid(ctx context.Context) (*FishResponse, error) {
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
func (client *polymorphicrecursiveOperations) getValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *polymorphicrecursiveOperations) getValidHandleResponse(resp *azcore.Response) (*FishResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := FishResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Fish)
}

// PutValid - Put complex types that are polymorphic and have recursive references
func (client *polymorphicrecursiveOperations) PutValid(ctx context.Context, complexBody Fish) (*http.Response, error) {
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
func (client *polymorphicrecursiveOperations) putValidCreateRequest(complexBody Fish) (*azcore.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleResponse handles the PutValid response.
func (client *polymorphicrecursiveOperations) putValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}
