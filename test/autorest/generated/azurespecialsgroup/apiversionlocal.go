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

// APIVersionLocalOperations contains the methods for the APIVersionLocal group.
type APIVersionLocalOperations interface {
	// GetMethodLocalNull - Get method with api-version modeled in the method.  pass in api-version = null to succeed
	GetMethodLocalNull(ctx context.Context, apiVersionLocalGetMethodLocalNullOptions *APIVersionLocalGetMethodLocalNullOptions) (*http.Response, error)
	// GetMethodLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
	GetMethodLocalValid(ctx context.Context) (*http.Response, error)
	// GetPathLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
	GetPathLocalValid(ctx context.Context) (*http.Response, error)
	// GetSwaggerLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
	GetSwaggerLocalValid(ctx context.Context) (*http.Response, error)
}

// apiVersionLocalOperations implements the APIVersionLocalOperations interface.
type apiVersionLocalOperations struct {
	*Client
}

// GetMethodLocalNull - Get method with api-version modeled in the method.  pass in api-version = null to succeed
func (client *apiVersionLocalOperations) GetMethodLocalNull(ctx context.Context, apiVersionLocalGetMethodLocalNullOptions *APIVersionLocalGetMethodLocalNullOptions) (*http.Response, error) {
	req, err := client.getMethodLocalNullCreateRequest(apiVersionLocalGetMethodLocalNullOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getMethodLocalNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMethodLocalNullCreateRequest creates the GetMethodLocalNull request.
func (client *apiVersionLocalOperations) getMethodLocalNullCreateRequest(apiVersionLocalGetMethodLocalNullOptions *APIVersionLocalGetMethodLocalNullOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/null"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if apiVersionLocalGetMethodLocalNullOptions != nil && apiVersionLocalGetMethodLocalNullOptions.ApiVersion != nil {
		query.Set("api-version", *apiVersionLocalGetMethodLocalNullOptions.ApiVersion)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getMethodLocalNullHandleResponse handles the GetMethodLocalNull response.
func (client *apiVersionLocalOperations) getMethodLocalNullHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodLocalNullHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodLocalNullHandleError handles the GetMethodLocalNull error response.
func (client *apiVersionLocalOperations) getMethodLocalNullHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetMethodLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func (client *apiVersionLocalOperations) GetMethodLocalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getMethodLocalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getMethodLocalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMethodLocalValidCreateRequest creates the GetMethodLocalValid request.
func (client *apiVersionLocalOperations) getMethodLocalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/2.0"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2.0")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getMethodLocalValidHandleResponse handles the GetMethodLocalValid response.
func (client *apiVersionLocalOperations) getMethodLocalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodLocalValidHandleError handles the GetMethodLocalValid error response.
func (client *apiVersionLocalOperations) getMethodLocalValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetPathLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func (client *apiVersionLocalOperations) GetPathLocalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getPathLocalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getPathLocalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getPathLocalValidCreateRequest creates the GetPathLocalValid request.
func (client *apiVersionLocalOperations) getPathLocalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/local/2.0"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2.0")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getPathLocalValidHandleResponse handles the GetPathLocalValid response.
func (client *apiVersionLocalOperations) getPathLocalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getPathLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getPathLocalValidHandleError handles the GetPathLocalValid error response.
func (client *apiVersionLocalOperations) getPathLocalValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func (client *apiVersionLocalOperations) GetSwaggerLocalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.getSwaggerLocalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getSwaggerLocalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSwaggerLocalValidCreateRequest creates the GetSwaggerLocalValid request.
func (client *apiVersionLocalOperations) getSwaggerLocalValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/local/2.0"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2.0")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getSwaggerLocalValidHandleResponse handles the GetSwaggerLocalValid response.
func (client *apiVersionLocalOperations) getSwaggerLocalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getSwaggerLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getSwaggerLocalValidHandleError handles the GetSwaggerLocalValid error response.
func (client *apiVersionLocalOperations) getSwaggerLocalValidHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
