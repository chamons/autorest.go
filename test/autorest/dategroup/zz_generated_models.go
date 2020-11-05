// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dategroup

import (
	"fmt"
	"net/http"
	"time"
)

// DateGetInvalidDateOptions contains the optional parameters for the Date.GetInvalidDate method.
type DateGetInvalidDateOptions struct {
	// placeholder for future optional parameters
}

// DateGetMaxDateOptions contains the optional parameters for the Date.GetMaxDate method.
type DateGetMaxDateOptions struct {
	// placeholder for future optional parameters
}

// DateGetMinDateOptions contains the optional parameters for the Date.GetMinDate method.
type DateGetMinDateOptions struct {
	// placeholder for future optional parameters
}

// DateGetNullOptions contains the optional parameters for the Date.GetNull method.
type DateGetNullOptions struct {
	// placeholder for future optional parameters
}

// DateGetOverflowDateOptions contains the optional parameters for the Date.GetOverflowDate method.
type DateGetOverflowDateOptions struct {
	// placeholder for future optional parameters
}

// DateGetUnderflowDateOptions contains the optional parameters for the Date.GetUnderflowDate method.
type DateGetUnderflowDateOptions struct {
	// placeholder for future optional parameters
}

// DatePutMaxDateOptions contains the optional parameters for the Date.PutMaxDate method.
type DatePutMaxDateOptions struct {
	// placeholder for future optional parameters
}

// DatePutMinDateOptions contains the optional parameters for the Date.PutMinDate method.
type DatePutMinDateOptions struct {
	// placeholder for future optional parameters
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
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

// TimeResponse is the response envelope for operations that return a time.Time type.
type TimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *time.Time
}