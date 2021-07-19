// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import "net/http"

// BoolGetFalseResponse contains the response from method Bool.GetFalse.
type BoolGetFalseResponse struct {
	BoolGetFalseResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BoolGetFalseResult contains the result from method Bool.GetFalse.
type BoolGetFalseResult struct {
	// simple boolean
	Value *bool
}

// BoolGetInvalidResponse contains the response from method Bool.GetInvalid.
type BoolGetInvalidResponse struct {
	BoolGetInvalidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BoolGetInvalidResult contains the result from method Bool.GetInvalid.
type BoolGetInvalidResult struct {
	Value *bool
}

// BoolGetNullResponse contains the response from method Bool.GetNull.
type BoolGetNullResponse struct {
	BoolGetNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BoolGetNullResult contains the result from method Bool.GetNull.
type BoolGetNullResult struct {
	Value *bool
}

// BoolGetTrueResponse contains the response from method Bool.GetTrue.
type BoolGetTrueResponse struct {
	BoolGetTrueResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BoolGetTrueResult contains the result from method Bool.GetTrue.
type BoolGetTrueResult struct {
	// simple boolean
	Value *bool
}

// BoolPutFalseResponse contains the response from method Bool.PutFalse.
type BoolPutFalseResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BoolPutTrueResponse contains the response from method Bool.PutTrue.
type BoolPutTrueResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
