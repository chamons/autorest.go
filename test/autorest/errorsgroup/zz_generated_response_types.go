//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import "net/http"

// PetDoSomethingResponse contains the response from method Pet.DoSomething.
type PetDoSomethingResponse struct {
	PetDoSomethingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetDoSomethingResult contains the result from method Pet.DoSomething.
type PetDoSomethingResult struct {
	PetAction
}

// PetGetPetByIDResponse contains the response from method Pet.GetPetByID.
type PetGetPetByIDResponse struct {
	PetGetPetByIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetGetPetByIDResult contains the result from method Pet.GetPetByID.
type PetGetPetByIDResult struct {
	Pet
}

// PetHasModelsParamResponse contains the response from method Pet.HasModelsParam.
type PetHasModelsParamResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
