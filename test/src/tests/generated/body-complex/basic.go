package complexgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// BasicClient is the test Infrastructure for AutoRest
type BasicClient struct {
	BaseClient
}

// NewBasicClient creates an instance of the BasicClient client.
func NewBasicClient() BasicClient {
	return NewBasicClientWithBaseURI(DefaultBaseURI)
}

// NewBasicClientWithBaseURI creates an instance of the BasicClient client.
func NewBasicClientWithBaseURI(baseURI string) BasicClient {
	return BasicClient{NewWithBaseURI(baseURI)}
}

// GetEmpty get a basic complex type that is empty
func (client BasicClient) GetEmpty(ctx context.Context) (result Basic, err error) {
	req, err := client.GetEmptyPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client BasicClient) GetEmptyPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/basic/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client BasicClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client BasicClient) GetEmptyResponder(resp *http.Response) (result Basic, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInvalid get a basic complex type that is invalid for the local strong type
func (client BasicClient) GetInvalid(ctx context.Context) (result Basic, err error) {
	req, err := client.GetInvalidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client BasicClient) GetInvalidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/basic/invalid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client BasicClient) GetInvalidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client BasicClient) GetInvalidResponder(resp *http.Response) (result Basic, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNotProvided get a basic complex type while the server doesn't provide a response payload
func (client BasicClient) GetNotProvided(ctx context.Context) (result Basic, err error) {
	req, err := client.GetNotProvidedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetNotProvided", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotProvidedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetNotProvided", resp, "Failure sending request")
		return
	}

	result, err = client.GetNotProvidedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetNotProvided", resp, "Failure responding to request")
	}

	return
}

// GetNotProvidedPreparer prepares the GetNotProvided request.
func (client BasicClient) GetNotProvidedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/basic/notprovided"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotProvidedSender sends the GetNotProvided request. The method will close the
// http.Response Body if it receives an error.
func (client BasicClient) GetNotProvidedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotProvidedResponder handles the response to the GetNotProvided request. The method always
// closes the http.Response Body.
func (client BasicClient) GetNotProvidedResponder(resp *http.Response) (result Basic, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get a basic complex type whose properties are null
func (client BasicClient) GetNull(ctx context.Context) (result Basic, err error) {
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client BasicClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/basic/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client BasicClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client BasicClient) GetNullResponder(resp *http.Response) (result Basic, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetValid get complex type {id: 2, name: 'abc', color: 'YELLOW'}
func (client BasicClient) GetValid(ctx context.Context) (result Basic, err error) {
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "GetValid", resp, "Failure responding to request")
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client BasicClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/basic/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client BasicClient) GetValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client BasicClient) GetValidResponder(resp *http.Response) (result Basic, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutValid please put {id: 2, name: 'abc', color: 'Magenta'}
//
// complexBody is please put {id: 2, name: 'abc', color: 'Magenta'}
func (client BasicClient) PutValid(ctx context.Context, complexBody Basic) (result autorest.Response, err error) {
	req, err := client.PutValidPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "PutValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "PutValid", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.BasicClient", "PutValid", resp, "Failure responding to request")
	}

	return
}

// PutValidPreparer prepares the PutValid request.
func (client BasicClient) PutValidPreparer(ctx context.Context, complexBody Basic) (*http.Request, error) {
	const APIVersion = "2014-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/basic/valid"),
		autorest.WithJSON(complexBody),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidSender sends the PutValid request. The method will close the
// http.Response Body if it receives an error.
func (client BasicClient) PutValidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutValidResponder handles the response to the PutValid request. The method always
// closes the http.Response Body.
func (client BasicClient) PutValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
