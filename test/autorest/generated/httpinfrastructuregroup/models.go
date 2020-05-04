// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"fmt"
	"net/http"
)

type B struct {
	StatusCode     *string `json:"statusCode,omitempty"`
	TextStatusCode *string `json:"textStatusCode,omitempty"`
}

// BoolResponse is the response envelope for operations that return a bool type.
type BoolResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// simple boolean
	Value *bool
}

type C struct {
	HTTPCode *string `json:"httpCode,omitempty"`
}

type D struct {
	HTTPStatusCode *string `json:"httpStatusCode,omitempty"`
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// HTTPRedirectsDelete307Response contains the response from method HTTPRedirects.Delete307.
type HTTPRedirectsDelete307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsGet300Response contains the response from method HTTPRedirects.Get300.
type HTTPRedirectsGet300Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsGet301Response contains the response from method HTTPRedirects.Get301.
type HTTPRedirectsGet301Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsGet302Response contains the response from method HTTPRedirects.Get302.
type HTTPRedirectsGet302Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsGet307Response contains the response from method HTTPRedirects.Get307.
type HTTPRedirectsGet307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsHead300Response contains the response from method HTTPRedirects.Head300.
type HTTPRedirectsHead300Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsHead301Response contains the response from method HTTPRedirects.Head301.
type HTTPRedirectsHead301Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsHead302Response contains the response from method HTTPRedirects.Head302.
type HTTPRedirectsHead302Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsHead307Response contains the response from method HTTPRedirects.Head307.
type HTTPRedirectsHead307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsOptions307Response contains the response from method HTTPRedirects.Options307.
type HTTPRedirectsOptions307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsPatch302Response contains the response from method HTTPRedirects.Patch302.
type HTTPRedirectsPatch302Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsPatch307Response contains the response from method HTTPRedirects.Patch307.
type HTTPRedirectsPatch307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsPost303Response contains the response from method HTTPRedirects.Post303.
type HTTPRedirectsPost303Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsPost307Response contains the response from method HTTPRedirects.Post307.
type HTTPRedirectsPost307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsPut301Response contains the response from method HTTPRedirects.Put301.
type HTTPRedirectsPut301Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPRedirectsPut307Response contains the response from method HTTPRedirects.Put307.
type HTTPRedirectsPut307Response struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type MyException struct {
	StatusCode *string `json:"statusCode,omitempty"`
}

func (e MyException) Error() string {
	msg := ""
	if e.StatusCode != nil {
		msg += fmt.Sprintf("StatusCode: %v\n", *e.StatusCode)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// MyExceptionResponse is the response envelope for operations that return a MyException type.
type MyExceptionResponse struct {
	MyException *MyException

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
