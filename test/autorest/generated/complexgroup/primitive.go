// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PrimitiveOperations contains the methods for the Primitive group.
type PrimitiveOperations interface {
	// GetBool - Get complex types with bool properties
	GetBool(ctx context.Context) (*BooleanWrapperResponse, error)
	// GetByte - Get complex types with byte properties
	GetByte(ctx context.Context) (*ByteWrapperResponse, error)
	// GetDate - Get complex types with date properties
	GetDate(ctx context.Context) (*DateWrapperResponse, error)
	// GetDateTime - Get complex types with datetime properties
	GetDateTime(ctx context.Context) (*DatetimeWrapperResponse, error)
	// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties
	GetDateTimeRFC1123(ctx context.Context) (*Datetimerfc1123WrapperResponse, error)
	// GetDouble - Get complex types with double properties
	GetDouble(ctx context.Context) (*DoubleWrapperResponse, error)
	// GetDuration - Get complex types with duration properties
	GetDuration(ctx context.Context) (*DurationWrapperResponse, error)
	// GetFloat - Get complex types with float properties
	GetFloat(ctx context.Context) (*FloatWrapperResponse, error)
	// GetInt - Get complex types with integer properties
	GetInt(ctx context.Context) (*IntWrapperResponse, error)
	// GetLong - Get complex types with long properties
	GetLong(ctx context.Context) (*LongWrapperResponse, error)
	// GetString - Get complex types with string properties
	GetString(ctx context.Context) (*StringWrapperResponse, error)
	// PutBool - Put complex types with bool properties
	PutBool(ctx context.Context, complexBody BooleanWrapper) (*http.Response, error)
	// PutByte - Put complex types with byte properties
	PutByte(ctx context.Context, complexBody ByteWrapper) (*http.Response, error)
	// PutDate - Put complex types with date properties
	PutDate(ctx context.Context, complexBody DateWrapper) (*http.Response, error)
	// PutDateTime - Put complex types with datetime properties
	PutDateTime(ctx context.Context, complexBody DatetimeWrapper) (*http.Response, error)
	// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties
	PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper) (*http.Response, error)
	// PutDouble - Put complex types with double properties
	PutDouble(ctx context.Context, complexBody DoubleWrapper) (*http.Response, error)
	// PutDuration - Put complex types with duration properties
	PutDuration(ctx context.Context, complexBody DurationWrapper) (*http.Response, error)
	// PutFloat - Put complex types with float properties
	PutFloat(ctx context.Context, complexBody FloatWrapper) (*http.Response, error)
	// PutInt - Put complex types with integer properties
	PutInt(ctx context.Context, complexBody IntWrapper) (*http.Response, error)
	// PutLong - Put complex types with long properties
	PutLong(ctx context.Context, complexBody LongWrapper) (*http.Response, error)
	// PutString - Put complex types with string properties
	PutString(ctx context.Context, complexBody StringWrapper) (*http.Response, error)
}

// primitiveOperations implements the PrimitiveOperations interface.
type primitiveOperations struct {
	*Client
}

// GetBool - Get complex types with bool properties
func (client *primitiveOperations) GetBool(ctx context.Context) (*BooleanWrapperResponse, error) {
	req, err := client.getBoolCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getBoolCreateRequest creates the GetBool request.
func (client *primitiveOperations) getBoolCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/bool"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getBoolHandleResponse handles the GetBool response.
func (client *primitiveOperations) getBoolHandleResponse(resp *azcore.Response) (*BooleanWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getBoolHandleError(resp)
	}
	result := BooleanWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// getBoolHandleError handles the GetBool error response.
func (client *primitiveOperations) getBoolHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetByte - Get complex types with byte properties
func (client *primitiveOperations) GetByte(ctx context.Context) (*ByteWrapperResponse, error) {
	req, err := client.getByteCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getByteCreateRequest creates the GetByte request.
func (client *primitiveOperations) getByteCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/byte"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getByteHandleResponse handles the GetByte response.
func (client *primitiveOperations) getByteHandleResponse(resp *azcore.Response) (*ByteWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getByteHandleError(resp)
	}
	result := ByteWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ByteWrapper)
}

// getByteHandleError handles the GetByte error response.
func (client *primitiveOperations) getByteHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDate - Get complex types with date properties
func (client *primitiveOperations) GetDate(ctx context.Context) (*DateWrapperResponse, error) {
	req, err := client.getDateCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getDateCreateRequest creates the GetDate request.
func (client *primitiveOperations) getDateCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/date"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getDateHandleResponse handles the GetDate response.
func (client *primitiveOperations) getDateHandleResponse(resp *azcore.Response) (*DateWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getDateHandleError(resp)
	}
	result := DateWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
}

// getDateHandleError handles the GetDate error response.
func (client *primitiveOperations) getDateHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDateTime - Get complex types with datetime properties
func (client *primitiveOperations) GetDateTime(ctx context.Context) (*DatetimeWrapperResponse, error) {
	req, err := client.getDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getDateTimeCreateRequest creates the GetDateTime request.
func (client *primitiveOperations) getDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetime"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getDateTimeHandleResponse handles the GetDateTime response.
func (client *primitiveOperations) getDateTimeHandleResponse(resp *azcore.Response) (*DatetimeWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getDateTimeHandleError(resp)
	}
	result := DatetimeWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DatetimeWrapper)
}

// getDateTimeHandleError handles the GetDateTime error response.
func (client *primitiveOperations) getDateTimeHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties
func (client *primitiveOperations) GetDateTimeRFC1123(ctx context.Context) (*Datetimerfc1123WrapperResponse, error) {
	req, err := client.getDateTimeRfc1123CreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getDateTimeRfc1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getDateTimeRfc1123CreateRequest creates the GetDateTimeRFC1123 request.
func (client *primitiveOperations) getDateTimeRfc1123CreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getDateTimeRfc1123HandleResponse handles the GetDateTimeRFC1123 response.
func (client *primitiveOperations) getDateTimeRfc1123HandleResponse(resp *azcore.Response) (*Datetimerfc1123WrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getDateTimeRfc1123HandleError(resp)
	}
	result := Datetimerfc1123WrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Datetimerfc1123Wrapper)
}

// getDateTimeRfc1123HandleError handles the GetDateTimeRFC1123 error response.
func (client *primitiveOperations) getDateTimeRfc1123HandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDouble - Get complex types with double properties
func (client *primitiveOperations) GetDouble(ctx context.Context) (*DoubleWrapperResponse, error) {
	req, err := client.getDoubleCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getDoubleCreateRequest creates the GetDouble request.
func (client *primitiveOperations) getDoubleCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/double"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getDoubleHandleResponse handles the GetDouble response.
func (client *primitiveOperations) getDoubleHandleResponse(resp *azcore.Response) (*DoubleWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getDoubleHandleError(resp)
	}
	result := DoubleWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// getDoubleHandleError handles the GetDouble error response.
func (client *primitiveOperations) getDoubleHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDuration - Get complex types with duration properties
func (client *primitiveOperations) GetDuration(ctx context.Context) (*DurationWrapperResponse, error) {
	req, err := client.getDurationCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getDurationCreateRequest creates the GetDuration request.
func (client *primitiveOperations) getDurationCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/duration"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getDurationHandleResponse handles the GetDuration response.
func (client *primitiveOperations) getDurationHandleResponse(resp *azcore.Response) (*DurationWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getDurationHandleError(resp)
	}
	result := DurationWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DurationWrapper)
}

// getDurationHandleError handles the GetDuration error response.
func (client *primitiveOperations) getDurationHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetFloat - Get complex types with float properties
func (client *primitiveOperations) GetFloat(ctx context.Context) (*FloatWrapperResponse, error) {
	req, err := client.getFloatCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getFloatCreateRequest creates the GetFloat request.
func (client *primitiveOperations) getFloatCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/float"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getFloatHandleResponse handles the GetFloat response.
func (client *primitiveOperations) getFloatHandleResponse(resp *azcore.Response) (*FloatWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getFloatHandleError(resp)
	}
	result := FloatWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// getFloatHandleError handles the GetFloat error response.
func (client *primitiveOperations) getFloatHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetInt - Get complex types with integer properties
func (client *primitiveOperations) GetInt(ctx context.Context) (*IntWrapperResponse, error) {
	req, err := client.getIntCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getIntCreateRequest creates the GetInt request.
func (client *primitiveOperations) getIntCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/integer"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getIntHandleResponse handles the GetInt response.
func (client *primitiveOperations) getIntHandleResponse(resp *azcore.Response) (*IntWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getIntHandleError(resp)
	}
	result := IntWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// getIntHandleError handles the GetInt error response.
func (client *primitiveOperations) getIntHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLong - Get complex types with long properties
func (client *primitiveOperations) GetLong(ctx context.Context) (*LongWrapperResponse, error) {
	req, err := client.getLongCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLongCreateRequest creates the GetLong request.
func (client *primitiveOperations) getLongCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/long"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLongHandleResponse handles the GetLong response.
func (client *primitiveOperations) getLongHandleResponse(resp *azcore.Response) (*LongWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLongHandleError(resp)
	}
	result := LongWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// getLongHandleError handles the GetLong error response.
func (client *primitiveOperations) getLongHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetString - Get complex types with string properties
func (client *primitiveOperations) GetString(ctx context.Context) (*StringWrapperResponse, error) {
	req, err := client.getStringCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getStringCreateRequest creates the GetString request.
func (client *primitiveOperations) getStringCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/primitive/string"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getStringHandleResponse handles the GetString response.
func (client *primitiveOperations) getStringHandleResponse(resp *azcore.Response) (*StringWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getStringHandleError(resp)
	}
	result := StringWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// getStringHandleError handles the GetString error response.
func (client *primitiveOperations) getStringHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutBool - Put complex types with bool properties
func (client *primitiveOperations) PutBool(ctx context.Context, complexBody BooleanWrapper) (*http.Response, error) {
	req, err := client.putBoolCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putBoolCreateRequest creates the PutBool request.
func (client *primitiveOperations) putBoolCreateRequest(complexBody BooleanWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/bool"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putBoolHandleResponse handles the PutBool response.
func (client *primitiveOperations) putBoolHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putBoolHandleError(resp)
	}
	return resp.Response, nil
}

// putBoolHandleError handles the PutBool error response.
func (client *primitiveOperations) putBoolHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutByte - Put complex types with byte properties
func (client *primitiveOperations) PutByte(ctx context.Context, complexBody ByteWrapper) (*http.Response, error) {
	req, err := client.putByteCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putByteCreateRequest creates the PutByte request.
func (client *primitiveOperations) putByteCreateRequest(complexBody ByteWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/byte"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putByteHandleResponse handles the PutByte response.
func (client *primitiveOperations) putByteHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putByteHandleError(resp)
	}
	return resp.Response, nil
}

// putByteHandleError handles the PutByte error response.
func (client *primitiveOperations) putByteHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutDate - Put complex types with date properties
func (client *primitiveOperations) PutDate(ctx context.Context, complexBody DateWrapper) (*http.Response, error) {
	req, err := client.putDateCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putDateCreateRequest creates the PutDate request.
func (client *primitiveOperations) putDateCreateRequest(complexBody DateWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/date"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putDateHandleResponse handles the PutDate response.
func (client *primitiveOperations) putDateHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDateHandleError(resp)
	}
	return resp.Response, nil
}

// putDateHandleError handles the PutDate error response.
func (client *primitiveOperations) putDateHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutDateTime - Put complex types with datetime properties
func (client *primitiveOperations) PutDateTime(ctx context.Context, complexBody DatetimeWrapper) (*http.Response, error) {
	req, err := client.putDateTimeCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putDateTimeCreateRequest creates the PutDateTime request.
func (client *primitiveOperations) putDateTimeCreateRequest(complexBody DatetimeWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetime"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putDateTimeHandleResponse handles the PutDateTime response.
func (client *primitiveOperations) putDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putDateTimeHandleError handles the PutDateTime error response.
func (client *primitiveOperations) putDateTimeHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties
func (client *primitiveOperations) PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper) (*http.Response, error) {
	req, err := client.putDateTimeRfc1123CreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putDateTimeRfc1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putDateTimeRfc1123CreateRequest creates the PutDateTimeRFC1123 request.
func (client *primitiveOperations) putDateTimeRfc1123CreateRequest(complexBody Datetimerfc1123Wrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putDateTimeRfc1123HandleResponse handles the PutDateTimeRFC1123 response.
func (client *primitiveOperations) putDateTimeRfc1123HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDateTimeRfc1123HandleError(resp)
	}
	return resp.Response, nil
}

// putDateTimeRfc1123HandleError handles the PutDateTimeRFC1123 error response.
func (client *primitiveOperations) putDateTimeRfc1123HandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutDouble - Put complex types with double properties
func (client *primitiveOperations) PutDouble(ctx context.Context, complexBody DoubleWrapper) (*http.Response, error) {
	req, err := client.putDoubleCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putDoubleCreateRequest creates the PutDouble request.
func (client *primitiveOperations) putDoubleCreateRequest(complexBody DoubleWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/double"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putDoubleHandleResponse handles the PutDouble response.
func (client *primitiveOperations) putDoubleHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDoubleHandleError(resp)
	}
	return resp.Response, nil
}

// putDoubleHandleError handles the PutDouble error response.
func (client *primitiveOperations) putDoubleHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutDuration - Put complex types with duration properties
func (client *primitiveOperations) PutDuration(ctx context.Context, complexBody DurationWrapper) (*http.Response, error) {
	req, err := client.putDurationCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putDurationCreateRequest creates the PutDuration request.
func (client *primitiveOperations) putDurationCreateRequest(complexBody DurationWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/duration"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putDurationHandleResponse handles the PutDuration response.
func (client *primitiveOperations) putDurationHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDurationHandleError(resp)
	}
	return resp.Response, nil
}

// putDurationHandleError handles the PutDuration error response.
func (client *primitiveOperations) putDurationHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutFloat - Put complex types with float properties
func (client *primitiveOperations) PutFloat(ctx context.Context, complexBody FloatWrapper) (*http.Response, error) {
	req, err := client.putFloatCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putFloatCreateRequest creates the PutFloat request.
func (client *primitiveOperations) putFloatCreateRequest(complexBody FloatWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/float"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putFloatHandleResponse handles the PutFloat response.
func (client *primitiveOperations) putFloatHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putFloatHandleError(resp)
	}
	return resp.Response, nil
}

// putFloatHandleError handles the PutFloat error response.
func (client *primitiveOperations) putFloatHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutInt - Put complex types with integer properties
func (client *primitiveOperations) PutInt(ctx context.Context, complexBody IntWrapper) (*http.Response, error) {
	req, err := client.putIntCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putIntCreateRequest creates the PutInt request.
func (client *primitiveOperations) putIntCreateRequest(complexBody IntWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/integer"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putIntHandleResponse handles the PutInt response.
func (client *primitiveOperations) putIntHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putIntHandleError(resp)
	}
	return resp.Response, nil
}

// putIntHandleError handles the PutInt error response.
func (client *primitiveOperations) putIntHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutLong - Put complex types with long properties
func (client *primitiveOperations) PutLong(ctx context.Context, complexBody LongWrapper) (*http.Response, error) {
	req, err := client.putLongCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLongCreateRequest creates the PutLong request.
func (client *primitiveOperations) putLongCreateRequest(complexBody LongWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/long"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putLongHandleResponse handles the PutLong response.
func (client *primitiveOperations) putLongHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putLongHandleError(resp)
	}
	return resp.Response, nil
}

// putLongHandleError handles the PutLong error response.
func (client *primitiveOperations) putLongHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutString - Put complex types with string properties
func (client *primitiveOperations) PutString(ctx context.Context, complexBody StringWrapper) (*http.Response, error) {
	req, err := client.putStringCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putStringCreateRequest creates the PutString request.
func (client *primitiveOperations) putStringCreateRequest(complexBody StringWrapper) (*azcore.Request, error) {
	urlPath := "/complex/primitive/string"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putStringHandleResponse handles the PutString response.
func (client *primitiveOperations) putStringHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putStringHandleError(resp)
	}
	return resp.Response, nil
}

// putStringHandleError handles the PutString error response.
func (client *primitiveOperations) putStringHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
