// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package objectgroup

import "net/http"

// ObjectTypeClientGetResponse contains the response from method ObjectTypeClient.Get.
type ObjectTypeClientGetResponse struct {
	ObjectTypeClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectTypeClientGetResult contains the result from method ObjectTypeClient.Get.
type ObjectTypeClientGetResult struct {
	// Any object
	Object map[string]interface{}
}

// ObjectTypeClientPutResponse contains the response from method ObjectTypeClient.Put.
type ObjectTypeClientPutResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
