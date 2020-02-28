// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type InheritanceOperations struct{}

// GetValidCreateRequest creates the GetValid request.
func (InheritanceOperations) GetValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/inheritance/valid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetValidHandleResponse handles the GetValid response.
func (InheritanceOperations) GetValidHandleResponse(resp *azcore.Response) (*InheritanceGetValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := InheritanceGetValidResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Siamese)
}

// PutValidCreateRequest creates the PutValid request.
func (InheritanceOperations) PutValidCreateRequest(u url.URL, complexBody Siamese) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/inheritance/valid")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutValidHandleResponse handles the PutValid response.
func (InheritanceOperations) PutValidHandleResponse(resp *azcore.Response) (*InheritancePutValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &InheritancePutValidResponse{StatusCode: resp.StatusCode}, nil
}
