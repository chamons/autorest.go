// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package migroup

import "net/http"

// MultipleInheritanceServiceClientGetCatResponse contains the response from method MultipleInheritanceServiceClient.GetCat.
type MultipleInheritanceServiceClientGetCatResponse struct {
	MultipleInheritanceServiceClientGetCatResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientGetCatResult contains the result from method MultipleInheritanceServiceClient.GetCat.
type MultipleInheritanceServiceClientGetCatResult struct {
	Cat
}

// MultipleInheritanceServiceClientGetFelineResponse contains the response from method MultipleInheritanceServiceClient.GetFeline.
type MultipleInheritanceServiceClientGetFelineResponse struct {
	MultipleInheritanceServiceClientGetFelineResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientGetFelineResult contains the result from method MultipleInheritanceServiceClient.GetFeline.
type MultipleInheritanceServiceClientGetFelineResult struct {
	Feline
}

// MultipleInheritanceServiceClientGetHorseResponse contains the response from method MultipleInheritanceServiceClient.GetHorse.
type MultipleInheritanceServiceClientGetHorseResponse struct {
	MultipleInheritanceServiceClientGetHorseResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientGetHorseResult contains the result from method MultipleInheritanceServiceClient.GetHorse.
type MultipleInheritanceServiceClientGetHorseResult struct {
	Horse
}

// MultipleInheritanceServiceClientGetKittenResponse contains the response from method MultipleInheritanceServiceClient.GetKitten.
type MultipleInheritanceServiceClientGetKittenResponse struct {
	MultipleInheritanceServiceClientGetKittenResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientGetKittenResult contains the result from method MultipleInheritanceServiceClient.GetKitten.
type MultipleInheritanceServiceClientGetKittenResult struct {
	Kitten
}

// MultipleInheritanceServiceClientGetPetResponse contains the response from method MultipleInheritanceServiceClient.GetPet.
type MultipleInheritanceServiceClientGetPetResponse struct {
	MultipleInheritanceServiceClientGetPetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientGetPetResult contains the result from method MultipleInheritanceServiceClient.GetPet.
type MultipleInheritanceServiceClientGetPetResult struct {
	Pet
}

// MultipleInheritanceServiceClientPutCatResponse contains the response from method MultipleInheritanceServiceClient.PutCat.
type MultipleInheritanceServiceClientPutCatResponse struct {
	MultipleInheritanceServiceClientPutCatResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientPutCatResult contains the result from method MultipleInheritanceServiceClient.PutCat.
type MultipleInheritanceServiceClientPutCatResult struct {
	Value *string
}

// MultipleInheritanceServiceClientPutFelineResponse contains the response from method MultipleInheritanceServiceClient.PutFeline.
type MultipleInheritanceServiceClientPutFelineResponse struct {
	MultipleInheritanceServiceClientPutFelineResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientPutFelineResult contains the result from method MultipleInheritanceServiceClient.PutFeline.
type MultipleInheritanceServiceClientPutFelineResult struct {
	Value *string
}

// MultipleInheritanceServiceClientPutHorseResponse contains the response from method MultipleInheritanceServiceClient.PutHorse.
type MultipleInheritanceServiceClientPutHorseResponse struct {
	MultipleInheritanceServiceClientPutHorseResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientPutHorseResult contains the result from method MultipleInheritanceServiceClient.PutHorse.
type MultipleInheritanceServiceClientPutHorseResult struct {
	Value *string
}

// MultipleInheritanceServiceClientPutKittenResponse contains the response from method MultipleInheritanceServiceClient.PutKitten.
type MultipleInheritanceServiceClientPutKittenResponse struct {
	MultipleInheritanceServiceClientPutKittenResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientPutKittenResult contains the result from method MultipleInheritanceServiceClient.PutKitten.
type MultipleInheritanceServiceClientPutKittenResult struct {
	Value *string
}

// MultipleInheritanceServiceClientPutPetResponse contains the response from method MultipleInheritanceServiceClient.PutPet.
type MultipleInheritanceServiceClientPutPetResponse struct {
	MultipleInheritanceServiceClientPutPetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MultipleInheritanceServiceClientPutPetResult contains the result from method MultipleInheritanceServiceClient.PutPet.
type MultipleInheritanceServiceClientPutPetResult struct {
	Value *string
}