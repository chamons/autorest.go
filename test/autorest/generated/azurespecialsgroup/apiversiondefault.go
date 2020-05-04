// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// APIVersionDefaultOperations contains the methods for the APIVersionDefault group.
type APIVersionDefaultOperations interface {
	// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
	GetMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error)
	// GetMethodGlobalValid - GET method with api-version modeled in global settings.
	GetMethodGlobalValid(ctx context.Context) (*http.Response, error)
	// GetPathGlobalValid - GET method with api-version modeled in global settings.
	GetPathGlobalValid(ctx context.Context) (*http.Response, error)
	// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
	GetSwaggerGlobalValid(ctx context.Context) (*http.Response, error)
}

// apiVersionDefaultOperations implements the APIVersionDefaultOperations interface.
type apiVersionDefaultOperations struct {
	*Client
}

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
func (client *apiVersionDefaultOperations) GetMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getMethodGlobalNotProvidedValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getMethodGlobalNotProvidedValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMethodGlobalNotProvidedValidCreateRequest creates the GetMethodGlobalNotProvidedValid request.
func (client *apiVersionDefaultOperations) getMethodGlobalNotProvidedValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/globalNotProvided/2015-07-01-preview"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getMethodGlobalNotProvidedValidHandleResponse handles the GetMethodGlobalNotProvidedValid response.
func (client *apiVersionDefaultOperations) getMethodGlobalNotProvidedValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodGlobalNotProvidedValidHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodGlobalNotProvidedValidHandleError handles the GetMethodGlobalNotProvidedValid error response.
func (client *apiVersionDefaultOperations) getMethodGlobalNotProvidedValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
func (client *apiVersionDefaultOperations) GetMethodGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getMethodGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getMethodGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMethodGlobalValidCreateRequest creates the GetMethodGlobalValid request.
func (client *apiVersionDefaultOperations) getMethodGlobalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/global/2015-07-01-preview"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getMethodGlobalValidHandleResponse handles the GetMethodGlobalValid response.
func (client *apiVersionDefaultOperations) getMethodGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodGlobalValidHandleError handles the GetMethodGlobalValid error response.
func (client *apiVersionDefaultOperations) getMethodGlobalValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
func (client *apiVersionDefaultOperations) GetPathGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getPathGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getPathGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getPathGlobalValidCreateRequest creates the GetPathGlobalValid request.
func (client *apiVersionDefaultOperations) getPathGlobalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/global/2015-07-01-preview"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getPathGlobalValidHandleResponse handles the GetPathGlobalValid response.
func (client *apiVersionDefaultOperations) getPathGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getPathGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getPathGlobalValidHandleError handles the GetPathGlobalValid error response.
func (client *apiVersionDefaultOperations) getPathGlobalValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
func (client *apiVersionDefaultOperations) GetSwaggerGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getSwaggerGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getSwaggerGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSwaggerGlobalValidCreateRequest creates the GetSwaggerGlobalValid request.
func (client *apiVersionDefaultOperations) getSwaggerGlobalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/global/2015-07-01-preview"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getSwaggerGlobalValidHandleResponse handles the GetSwaggerGlobalValid response.
func (client *apiVersionDefaultOperations) getSwaggerGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getSwaggerGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getSwaggerGlobalValidHandleError handles the GetSwaggerGlobalValid error response.
func (client *apiVersionDefaultOperations) getSwaggerGlobalValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
