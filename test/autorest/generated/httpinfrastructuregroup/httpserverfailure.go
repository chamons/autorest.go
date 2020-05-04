// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPServerFailureOperations contains the methods for the HTTPServerFailure group.
type HTTPServerFailureOperations interface {
	// Delete505 - Return 505 status code - should be represented in the client as an error
	Delete505(ctx context.Context) (*http.Response, error)
	// Get501 - Return 501 status code - should be represented in the client as an error
	Get501(ctx context.Context) (*http.Response, error)
	// Head501 - Return 501 status code - should be represented in the client as an error
	Head501(ctx context.Context) (*http.Response, error)
	// Post505 - Return 505 status code - should be represented in the client as an error
	Post505(ctx context.Context) (*http.Response, error)
}

// httpServerFailureOperations implements the HTTPServerFailureOperations interface.
type httpServerFailureOperations struct {
	*Client
}

// Delete505 - Return 505 status code - should be represented in the client as an error
func (client *httpServerFailureOperations) Delete505(ctx context.Context) (*http.Response, error) {
	req, err := client.delete505CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.delete505HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// delete505CreateRequest creates the Delete505 request.
func (client *httpServerFailureOperations) delete505CreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/server/505"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, req.MarshalAsJSON(true)
}

// delete505HandleResponse handles the Delete505 response.
func (client *httpServerFailureOperations) delete505HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.delete505HandleError(resp)
}

// delete505HandleError handles the Delete505 error response.
func (client *httpServerFailureOperations) delete505HandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get501 - Return 501 status code - should be represented in the client as an error
func (client *httpServerFailureOperations) Get501(ctx context.Context) (*http.Response, error) {
	req, err := client.get501CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.get501HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get501CreateRequest creates the Get501 request.
func (client *httpServerFailureOperations) get501CreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/server/501"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// get501HandleResponse handles the Get501 response.
func (client *httpServerFailureOperations) get501HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.get501HandleError(resp)
}

// get501HandleError handles the Get501 error response.
func (client *httpServerFailureOperations) get501HandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head501 - Return 501 status code - should be represented in the client as an error
func (client *httpServerFailureOperations) Head501(ctx context.Context) (*http.Response, error) {
	req, err := client.head501CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.head501HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// head501CreateRequest creates the Head501 request.
func (client *httpServerFailureOperations) head501CreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/server/501"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodHead, *u)
	return req, nil
}

// head501HandleResponse handles the Head501 response.
func (client *httpServerFailureOperations) head501HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.head501HandleError(resp)
}

// head501HandleError handles the Head501 error response.
func (client *httpServerFailureOperations) head501HandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post505 - Return 505 status code - should be represented in the client as an error
func (client *httpServerFailureOperations) Post505(ctx context.Context) (*http.Response, error) {
	req, err := client.post505CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.post505HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// post505CreateRequest creates the Post505 request.
func (client *httpServerFailureOperations) post505CreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/server/505"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(true)
}

// post505HandleResponse handles the Post505 response.
func (client *httpServerFailureOperations) post505HandleResponse(resp *azcore.Response) (*http.Response, error) {
	return nil, client.post505HandleError(resp)
}

// post505HandleError handles the Post505 error response.
func (client *httpServerFailureOperations) post505HandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
