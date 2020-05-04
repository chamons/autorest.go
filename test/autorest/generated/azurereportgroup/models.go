// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import (
	"fmt"
	"net/http"
)

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

// MapOfInt32Response is the response envelope for operations that return a map[string]int32 type.
type MapOfInt32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Dictionary of <integer>
	Value *map[string]int32
}

// OperationsGetReportOptions contains the optional parameters for the Operations.GetReport method.
type OperationsGetReportOptions struct {
	// If specified, qualifies the generated report further (e.g. '2.7' vs '3.5' in for Python). The only effect is, that generators
	// that run all tests several times, can distinguish the generated reports.
	Qualifier *string
}
