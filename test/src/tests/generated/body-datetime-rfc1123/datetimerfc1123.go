package datetimerfc1123group

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// Datetimerfc1123Client is the test Infrastructure for AutoRest
type Datetimerfc1123Client struct {
	BaseClient
}

// NewDatetimerfc1123Client creates an instance of the Datetimerfc1123Client client.
func NewDatetimerfc1123Client() Datetimerfc1123Client {
	return NewDatetimerfc1123ClientWithBaseURI(DefaultBaseURI)
}

// NewDatetimerfc1123ClientWithBaseURI creates an instance of the Datetimerfc1123Client client.
func NewDatetimerfc1123ClientWithBaseURI(baseURI string) Datetimerfc1123Client {
	return Datetimerfc1123Client{NewWithBaseURI(baseURI)}
}

// GetInvalid get invalid datetime value
func (client Datetimerfc1123Client) GetInvalid(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetInvalidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client Datetimerfc1123Client) GetInvalidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/invalid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetInvalidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetInvalidResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null datetime value
func (client Datetimerfc1123Client) GetNull(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client Datetimerfc1123Client) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetNullResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflow get overflow datetime value
func (client Datetimerfc1123Client) GetOverflow(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetOverflowPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetOverflow", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetOverflow", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetOverflow", resp, "Failure responding to request")
	}

	return
}

// GetOverflowPreparer prepares the GetOverflow request.
func (client Datetimerfc1123Client) GetOverflowPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/overflow"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOverflowSender sends the GetOverflow request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetOverflowSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOverflowResponder handles the response to the GetOverflow request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetOverflowResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflow get underflow datetime value
func (client Datetimerfc1123Client) GetUnderflow(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetUnderflowPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUnderflow", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUnderflow", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUnderflow", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowPreparer prepares the GetUnderflow request.
func (client Datetimerfc1123Client) GetUnderflowPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/underflow"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUnderflowSender sends the GetUnderflow request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetUnderflowSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnderflowResponder handles the response to the GetUnderflow request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetUnderflowResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUtcLowercaseMaxDateTime get max datetime value fri, 31 dec 9999 23:59:59 gmt
func (client Datetimerfc1123Client) GetUtcLowercaseMaxDateTime(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetUtcLowercaseMaxDateTimePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcLowercaseMaxDateTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUtcLowercaseMaxDateTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcLowercaseMaxDateTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetUtcLowercaseMaxDateTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcLowercaseMaxDateTime", resp, "Failure responding to request")
	}

	return
}

// GetUtcLowercaseMaxDateTimePreparer prepares the GetUtcLowercaseMaxDateTime request.
func (client Datetimerfc1123Client) GetUtcLowercaseMaxDateTimePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/max/lowercase"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUtcLowercaseMaxDateTimeSender sends the GetUtcLowercaseMaxDateTime request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetUtcLowercaseMaxDateTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUtcLowercaseMaxDateTimeResponder handles the response to the GetUtcLowercaseMaxDateTime request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetUtcLowercaseMaxDateTimeResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUtcMinDateTime get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client Datetimerfc1123Client) GetUtcMinDateTime(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetUtcMinDateTimePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcMinDateTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUtcMinDateTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcMinDateTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetUtcMinDateTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcMinDateTime", resp, "Failure responding to request")
	}

	return
}

// GetUtcMinDateTimePreparer prepares the GetUtcMinDateTime request.
func (client Datetimerfc1123Client) GetUtcMinDateTimePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/min"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUtcMinDateTimeSender sends the GetUtcMinDateTime request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetUtcMinDateTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUtcMinDateTimeResponder handles the response to the GetUtcMinDateTime request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetUtcMinDateTimeResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUtcUppercaseMaxDateTime get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func (client Datetimerfc1123Client) GetUtcUppercaseMaxDateTime(ctx context.Context) (result DateTimeRfc1123, err error) {
	req, err := client.GetUtcUppercaseMaxDateTimePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcUppercaseMaxDateTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUtcUppercaseMaxDateTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcUppercaseMaxDateTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetUtcUppercaseMaxDateTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "GetUtcUppercaseMaxDateTime", resp, "Failure responding to request")
	}

	return
}

// GetUtcUppercaseMaxDateTimePreparer prepares the GetUtcUppercaseMaxDateTime request.
func (client Datetimerfc1123Client) GetUtcUppercaseMaxDateTimePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/max/uppercase"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUtcUppercaseMaxDateTimeSender sends the GetUtcUppercaseMaxDateTime request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) GetUtcUppercaseMaxDateTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUtcUppercaseMaxDateTimeResponder handles the response to the GetUtcUppercaseMaxDateTime request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) GetUtcUppercaseMaxDateTimeResponder(resp *http.Response) (result DateTimeRfc1123, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutUtcMaxDateTime put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
//
func (client Datetimerfc1123Client) PutUtcMaxDateTime(ctx context.Context, datetimeBody date.TimeRFC1123) (result autorest.Response, err error) {
	req, err := client.PutUtcMaxDateTimePreparer(ctx, datetimeBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "PutUtcMaxDateTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutUtcMaxDateTimeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "PutUtcMaxDateTime", resp, "Failure sending request")
		return
	}

	result, err = client.PutUtcMaxDateTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "PutUtcMaxDateTime", resp, "Failure responding to request")
	}

	return
}

// PutUtcMaxDateTimePreparer prepares the PutUtcMaxDateTime request.
func (client Datetimerfc1123Client) PutUtcMaxDateTimePreparer(ctx context.Context, datetimeBody date.TimeRFC1123) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/max"),
		autorest.WithJSON(datetimeBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutUtcMaxDateTimeSender sends the PutUtcMaxDateTime request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) PutUtcMaxDateTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutUtcMaxDateTimeResponder handles the response to the PutUtcMaxDateTime request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) PutUtcMaxDateTimeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutUtcMinDateTime put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
//
func (client Datetimerfc1123Client) PutUtcMinDateTime(ctx context.Context, datetimeBody date.TimeRFC1123) (result autorest.Response, err error) {
	req, err := client.PutUtcMinDateTimePreparer(ctx, datetimeBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "PutUtcMinDateTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutUtcMinDateTimeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "PutUtcMinDateTime", resp, "Failure sending request")
		return
	}

	result, err = client.PutUtcMinDateTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datetimerfc1123group.Datetimerfc1123Client", "PutUtcMinDateTime", resp, "Failure responding to request")
	}

	return
}

// PutUtcMinDateTimePreparer prepares the PutUtcMinDateTime request.
func (client Datetimerfc1123Client) PutUtcMinDateTimePreparer(ctx context.Context, datetimeBody date.TimeRFC1123) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/datetimerfc1123/min"),
		autorest.WithJSON(datetimeBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutUtcMinDateTimeSender sends the PutUtcMinDateTime request. The method will close the
// http.Response Body if it receives an error.
func (client Datetimerfc1123Client) PutUtcMinDateTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutUtcMinDateTimeResponder handles the response to the PutUtcMinDateTime request. The method always
// closes the http.Response Body.
func (client Datetimerfc1123Client) PutUtcMinDateTimeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
