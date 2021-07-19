// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"net/http"
	"time"
)

// Datetimerfc1123GetInvalidResponse contains the response from method Datetimerfc1123.GetInvalid.
type Datetimerfc1123GetInvalidResponse struct {
	Datetimerfc1123GetInvalidResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetInvalidResult contains the result from method Datetimerfc1123.GetInvalid.
type Datetimerfc1123GetInvalidResult struct {
	Value *time.Time
}

// Datetimerfc1123GetNullResponse contains the response from method Datetimerfc1123.GetNull.
type Datetimerfc1123GetNullResponse struct {
	Datetimerfc1123GetNullResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetNullResult contains the result from method Datetimerfc1123.GetNull.
type Datetimerfc1123GetNullResult struct {
	Value *time.Time
}

// Datetimerfc1123GetOverflowResponse contains the response from method Datetimerfc1123.GetOverflow.
type Datetimerfc1123GetOverflowResponse struct {
	Datetimerfc1123GetOverflowResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetOverflowResult contains the result from method Datetimerfc1123.GetOverflow.
type Datetimerfc1123GetOverflowResult struct {
	Value *time.Time
}

// Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse contains the response from method Datetimerfc1123.GetUTCLowercaseMaxDateTime.
type Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse struct {
	Datetimerfc1123GetUTCLowercaseMaxDateTimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetUTCLowercaseMaxDateTimeResult contains the result from method Datetimerfc1123.GetUTCLowercaseMaxDateTime.
type Datetimerfc1123GetUTCLowercaseMaxDateTimeResult struct {
	Value *time.Time
}

// Datetimerfc1123GetUTCMinDateTimeResponse contains the response from method Datetimerfc1123.GetUTCMinDateTime.
type Datetimerfc1123GetUTCMinDateTimeResponse struct {
	Datetimerfc1123GetUTCMinDateTimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetUTCMinDateTimeResult contains the result from method Datetimerfc1123.GetUTCMinDateTime.
type Datetimerfc1123GetUTCMinDateTimeResult struct {
	Value *time.Time
}

// Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse contains the response from method Datetimerfc1123.GetUTCUppercaseMaxDateTime.
type Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse struct {
	Datetimerfc1123GetUTCUppercaseMaxDateTimeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetUTCUppercaseMaxDateTimeResult contains the result from method Datetimerfc1123.GetUTCUppercaseMaxDateTime.
type Datetimerfc1123GetUTCUppercaseMaxDateTimeResult struct {
	Value *time.Time
}

// Datetimerfc1123GetUnderflowResponse contains the response from method Datetimerfc1123.GetUnderflow.
type Datetimerfc1123GetUnderflowResponse struct {
	Datetimerfc1123GetUnderflowResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123GetUnderflowResult contains the result from method Datetimerfc1123.GetUnderflow.
type Datetimerfc1123GetUnderflowResult struct {
	Value *time.Time
}

// Datetimerfc1123PutUTCMaxDateTimeResponse contains the response from method Datetimerfc1123.PutUTCMaxDateTime.
type Datetimerfc1123PutUTCMaxDateTimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Datetimerfc1123PutUTCMinDateTimeResponse contains the response from method Datetimerfc1123.PutUTCMinDateTime.
type Datetimerfc1123PutUTCMinDateTimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}