// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headergroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

func newError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
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

// HeaderCustomRequestIDResponse contains the response from method Header.CustomRequestID.
type HeaderCustomRequestIDResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamBoolResponse contains the response from method Header.ParamBool.
type HeaderParamBoolResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamByteResponse contains the response from method Header.ParamByte.
type HeaderParamByteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamDateResponse contains the response from method Header.ParamDate.
type HeaderParamDateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamDatetimeRFC1123Options contains the optional parameters for the Header.ParamDatetimeRFC1123 method.
type HeaderParamDatetimeRFC1123Options struct {
	// Send a post request with header values "Wed, 01 Jan 2010 12:34:56 GMT" or "Mon, 01 Jan 0001 00:00:00 GMT"
	Value *time.Time
}

// HeaderParamDatetimeRFC1123Response contains the response from method Header.ParamDatetimeRFC1123.
type HeaderParamDatetimeRFC1123Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamDatetimeResponse contains the response from method Header.ParamDatetime.
type HeaderParamDatetimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamDoubleResponse contains the response from method Header.ParamDouble.
type HeaderParamDoubleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamDurationResponse contains the response from method Header.ParamDuration.
type HeaderParamDurationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamEnumOptions contains the optional parameters for the Header.ParamEnum method.
type HeaderParamEnumOptions struct {
	// Send a post request with header values 'GREY'
	Value *GreyscaleColors
}

// HeaderParamEnumResponse contains the response from method Header.ParamEnum.
type HeaderParamEnumResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamExistingKeyResponse contains the response from method Header.ParamExistingKey.
type HeaderParamExistingKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamFloatResponse contains the response from method Header.ParamFloat.
type HeaderParamFloatResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamIntegerResponse contains the response from method Header.ParamInteger.
type HeaderParamIntegerResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamLongResponse contains the response from method Header.ParamLong.
type HeaderParamLongResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamProtectedKeyResponse contains the response from method Header.ParamProtectedKey.
type HeaderParamProtectedKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderParamStringOptions contains the optional parameters for the Header.ParamString method.
type HeaderParamStringOptions struct {
	// Send a post request with header values "The quick brown fox jumps over the lazy dog" or null or ""
	Value *string
}

// HeaderParamStringResponse contains the response from method Header.ParamString.
type HeaderParamStringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderResponseBoolResponse contains the response from method Header.ResponseBool.
type HeaderResponseBoolResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *bool
}

// HeaderResponseByteResponse contains the response from method Header.ResponseByte.
type HeaderResponseByteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *[]byte
}

// HeaderResponseDateResponse contains the response from method Header.ResponseDate.
type HeaderResponseDateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *time.Time
}

// HeaderResponseDatetimeRFC1123Response contains the response from method Header.ResponseDatetimeRFC1123.
type HeaderResponseDatetimeRFC1123Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *time.Time
}

// HeaderResponseDatetimeResponse contains the response from method Header.ResponseDatetime.
type HeaderResponseDatetimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *time.Time
}

// HeaderResponseDoubleResponse contains the response from method Header.ResponseDouble.
type HeaderResponseDoubleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *float64
}

// HeaderResponseDurationResponse contains the response from method Header.ResponseDuration.
type HeaderResponseDurationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *time.Duration
}

// HeaderResponseEnumResponse contains the response from method Header.ResponseEnum.
type HeaderResponseEnumResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *GreyscaleColors
}

// HeaderResponseExistingKeyResponse contains the response from method Header.ResponseExistingKey.
type HeaderResponseExistingKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// UserAgent contains the information returned from the UserAgent header response.
	UserAgent *string
}

// HeaderResponseFloatResponse contains the response from method Header.ResponseFloat.
type HeaderResponseFloatResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *float32
}

// HeaderResponseIntegerResponse contains the response from method Header.ResponseInteger.
type HeaderResponseIntegerResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *int32
}

// HeaderResponseLongResponse contains the response from method Header.ResponseLong.
type HeaderResponseLongResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *int64
}

// HeaderResponseProtectedKeyResponse contains the response from method Header.ResponseProtectedKey.
type HeaderResponseProtectedKeyResponse struct {
	// ContentType contains the information returned from the ContentType header response.
	ContentType *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderResponseStringResponse contains the response from method Header.ResponseString.
type HeaderResponseStringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the Value header response.
	Value *string
}