// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// Datetimerfc1123Client contains the methods for the Datetimerfc1123 group.
// Don't use this type directly, use NewDatetimerfc1123Client() instead.
type Datetimerfc1123Client struct {
	con *Connection
}

// NewDatetimerfc1123Client creates a new instance of Datetimerfc1123Client with the specified values.
func NewDatetimerfc1123Client(con *Connection) *Datetimerfc1123Client {
	return &Datetimerfc1123Client{con: con}
}

// GetInvalid - Get invalid datetime value
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetInvalid(ctx context.Context, options *Datetimerfc1123GetInvalidOptions) (Datetimerfc1123GetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetInvalidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetInvalidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetInvalidResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *Datetimerfc1123Client) getInvalidCreateRequest(ctx context.Context, options *Datetimerfc1123GetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *Datetimerfc1123Client) getInvalidHandleResponse(resp *azcore.Response) (Datetimerfc1123GetInvalidResponse, error) {
	result := Datetimerfc1123GetInvalidResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetInvalidResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *Datetimerfc1123Client) getInvalidHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetNull - Get null datetime value
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetNull(ctx context.Context, options *Datetimerfc1123GetNullOptions) (Datetimerfc1123GetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetNullResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetNullResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetNullResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *Datetimerfc1123Client) getNullCreateRequest(ctx context.Context, options *Datetimerfc1123GetNullOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *Datetimerfc1123Client) getNullHandleResponse(resp *azcore.Response) (Datetimerfc1123GetNullResponse, error) {
	result := Datetimerfc1123GetNullResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetNullResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getNullHandleError handles the GetNull error response.
func (client *Datetimerfc1123Client) getNullHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetOverflow - Get overflow datetime value
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetOverflow(ctx context.Context, options *Datetimerfc1123GetOverflowOptions) (Datetimerfc1123GetOverflowResponse, error) {
	req, err := client.getOverflowCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetOverflowResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetOverflowResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetOverflowResponse{}, client.getOverflowHandleError(resp)
	}
	return client.getOverflowHandleResponse(resp)
}

// getOverflowCreateRequest creates the GetOverflow request.
func (client *Datetimerfc1123Client) getOverflowCreateRequest(ctx context.Context, options *Datetimerfc1123GetOverflowOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/overflow"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getOverflowHandleResponse handles the GetOverflow response.
func (client *Datetimerfc1123Client) getOverflowHandleResponse(resp *azcore.Response) (Datetimerfc1123GetOverflowResponse, error) {
	result := Datetimerfc1123GetOverflowResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetOverflowResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getOverflowHandleError handles the GetOverflow error response.
func (client *Datetimerfc1123Client) getOverflowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetUTCLowercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions) (Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse, error) {
	req, err := client.getUTCLowercaseMaxDateTimeCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse{}, client.getUTCLowercaseMaxDateTimeHandleError(resp)
	}
	return client.getUTCLowercaseMaxDateTimeHandleResponse(resp)
}

// getUTCLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (client *Datetimerfc1123Client) getUTCLowercaseMaxDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/lowercase"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUTCLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (client *Datetimerfc1123Client) getUTCLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse, error) {
	result := Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getUTCLowercaseMaxDateTimeHandleError handles the GetUTCLowercaseMaxDateTime error response.
func (client *Datetimerfc1123Client) getUTCLowercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetUTCMinDateTime(ctx context.Context, options *Datetimerfc1123GetUTCMinDateTimeOptions) (Datetimerfc1123GetUTCMinDateTimeResponse, error) {
	req, err := client.getUTCMinDateTimeCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetUTCMinDateTimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetUTCMinDateTimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetUTCMinDateTimeResponse{}, client.getUTCMinDateTimeHandleError(resp)
	}
	return client.getUTCMinDateTimeHandleResponse(resp)
}

// getUTCMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (client *Datetimerfc1123Client) getUTCMinDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123GetUTCMinDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUTCMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (client *Datetimerfc1123Client) getUTCMinDateTimeHandleResponse(resp *azcore.Response) (Datetimerfc1123GetUTCMinDateTimeResponse, error) {
	result := Datetimerfc1123GetUTCMinDateTimeResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetUTCMinDateTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getUTCMinDateTimeHandleError handles the GetUTCMinDateTime error response.
func (client *Datetimerfc1123Client) getUTCMinDateTimeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetUTCUppercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions) (Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse, error) {
	req, err := client.getUTCUppercaseMaxDateTimeCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse{}, client.getUTCUppercaseMaxDateTimeHandleError(resp)
	}
	return client.getUTCUppercaseMaxDateTimeHandleResponse(resp)
}

// getUTCUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (client *Datetimerfc1123Client) getUTCUppercaseMaxDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/uppercase"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUTCUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (client *Datetimerfc1123Client) getUTCUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse, error) {
	result := Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getUTCUppercaseMaxDateTimeHandleError handles the GetUTCUppercaseMaxDateTime error response.
func (client *Datetimerfc1123Client) getUTCUppercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetUnderflow - Get underflow datetime value
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) GetUnderflow(ctx context.Context, options *Datetimerfc1123GetUnderflowOptions) (Datetimerfc1123GetUnderflowResponse, error) {
	req, err := client.getUnderflowCreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123GetUnderflowResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123GetUnderflowResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123GetUnderflowResponse{}, client.getUnderflowHandleError(resp)
	}
	return client.getUnderflowHandleResponse(resp)
}

// getUnderflowCreateRequest creates the GetUnderflow request.
func (client *Datetimerfc1123Client) getUnderflowCreateRequest(ctx context.Context, options *Datetimerfc1123GetUnderflowOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/underflow"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUnderflowHandleResponse handles the GetUnderflow response.
func (client *Datetimerfc1123Client) getUnderflowHandleResponse(resp *azcore.Response) (Datetimerfc1123GetUnderflowResponse, error) {
	result := Datetimerfc1123GetUnderflowResponse{RawResponse: resp.Response}
	var aux *timeRFC1123
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return Datetimerfc1123GetUnderflowResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getUnderflowHandleError handles the GetUnderflow error response.
func (client *Datetimerfc1123Client) getUnderflowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMaxDateTimeOptions) (Datetimerfc1123PutUTCMaxDateTimeResponse, error) {
	req, err := client.putUTCMaxDateTimeCreateRequest(ctx, datetimeBody, options)
	if err != nil {
		return Datetimerfc1123PutUTCMaxDateTimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123PutUTCMaxDateTimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123PutUTCMaxDateTimeResponse{}, client.putUTCMaxDateTimeHandleError(resp)
	}
	return Datetimerfc1123PutUTCMaxDateTimeResponse{RawResponse: resp.Response}, nil
}

// putUTCMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (client *Datetimerfc1123Client) putUTCMaxDateTimeCreateRequest(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMaxDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	aux := timeRFC1123(datetimeBody)
	return req, req.MarshalAsJSON(aux)
}

// putUTCMaxDateTimeHandleError handles the PutUTCMaxDateTime error response.
func (client *Datetimerfc1123Client) putUTCMaxDateTimeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
// If the operation fails it returns the *Error error type.
func (client *Datetimerfc1123Client) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMinDateTimeOptions) (Datetimerfc1123PutUTCMinDateTimeResponse, error) {
	req, err := client.putUTCMinDateTimeCreateRequest(ctx, datetimeBody, options)
	if err != nil {
		return Datetimerfc1123PutUTCMinDateTimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123PutUTCMinDateTimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123PutUTCMinDateTimeResponse{}, client.putUTCMinDateTimeHandleError(resp)
	}
	return Datetimerfc1123PutUTCMinDateTimeResponse{RawResponse: resp.Response}, nil
}

// putUTCMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (client *Datetimerfc1123Client) putUTCMinDateTimeCreateRequest(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMinDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	aux := timeRFC1123(datetimeBody)
	return req, req.MarshalAsJSON(aux)
}

// putUTCMinDateTimeHandleError handles the PutUTCMinDateTime error response.
func (client *Datetimerfc1123Client) putUTCMinDateTimeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
