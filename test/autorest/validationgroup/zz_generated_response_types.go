// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package validationgroup

import "net/http"

// AutoRestValidationTestGetWithConstantInPathResponse contains the response from method AutoRestValidationTest.GetWithConstantInPath.
type AutoRestValidationTestGetWithConstantInPathResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AutoRestValidationTestPostWithConstantInBodyResponse contains the response from method AutoRestValidationTest.PostWithConstantInBody.
type AutoRestValidationTestPostWithConstantInBodyResponse struct {
	AutoRestValidationTestPostWithConstantInBodyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AutoRestValidationTestPostWithConstantInBodyResult contains the result from method AutoRestValidationTest.PostWithConstantInBody.
type AutoRestValidationTestPostWithConstantInBodyResult struct {
	Product
}

// AutoRestValidationTestValidationOfBodyResponse contains the response from method AutoRestValidationTest.ValidationOfBody.
type AutoRestValidationTestValidationOfBodyResponse struct {
	AutoRestValidationTestValidationOfBodyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AutoRestValidationTestValidationOfBodyResult contains the result from method AutoRestValidationTest.ValidationOfBody.
type AutoRestValidationTestValidationOfBodyResult struct {
	Product
}

// AutoRestValidationTestValidationOfMethodParametersResponse contains the response from method AutoRestValidationTest.ValidationOfMethodParameters.
type AutoRestValidationTestValidationOfMethodParametersResponse struct {
	AutoRestValidationTestValidationOfMethodParametersResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AutoRestValidationTestValidationOfMethodParametersResult contains the result from method AutoRestValidationTest.ValidationOfMethodParameters.
type AutoRestValidationTestValidationOfMethodParametersResult struct {
	Product
}
