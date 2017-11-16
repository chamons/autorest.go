package stringgroup

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

// StringClient is the test Infrastructure for AutoRest Swagger BAT
type StringClient struct {
	BaseClient
}

// NewStringClient creates an instance of the StringClient client.
func NewStringClient() StringClient {
	return NewStringClientWithBaseURI(DefaultBaseURI)
}

// NewStringClientWithBaseURI creates an instance of the StringClient client.
func NewStringClientWithBaseURI(baseURI string) StringClient {
	return StringClient{NewWithBaseURI(baseURI)}
}

// GetBase64Encoded get value that is base64 encoded
func (client StringClient) GetBase64Encoded(ctx context.Context) (result Base64URL, err error) {
	req, err := client.GetBase64EncodedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetBase64Encoded", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBase64EncodedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetBase64Encoded", resp, "Failure sending request")
		return
	}

	result, err = client.GetBase64EncodedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetBase64Encoded", resp, "Failure responding to request")
	}

	return
}

// GetBase64EncodedPreparer prepares the GetBase64Encoded request.
func (client StringClient) GetBase64EncodedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/base64Encoding"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBase64EncodedSender sends the GetBase64Encoded request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetBase64EncodedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBase64EncodedResponder handles the response to the GetBase64Encoded request. The method always
// closes the http.Response Body.
func (client StringClient) GetBase64EncodedResponder(resp *http.Response) (result Base64URL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBase64URLEncoded get value that is base64url encoded
func (client StringClient) GetBase64URLEncoded(ctx context.Context) (result Base64URL, err error) {
	req, err := client.GetBase64URLEncodedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetBase64URLEncoded", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBase64URLEncodedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetBase64URLEncoded", resp, "Failure sending request")
		return
	}

	result, err = client.GetBase64URLEncodedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetBase64URLEncoded", resp, "Failure responding to request")
	}

	return
}

// GetBase64URLEncodedPreparer prepares the GetBase64URLEncoded request.
func (client StringClient) GetBase64URLEncodedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/base64UrlEncoding"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBase64URLEncodedSender sends the GetBase64URLEncoded request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetBase64URLEncodedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBase64URLEncodedResponder handles the response to the GetBase64URLEncoded request. The method always
// closes the http.Response Body.
func (client StringClient) GetBase64URLEncodedResponder(resp *http.Response) (result Base64URL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEmpty get empty string value value ''
func (client StringClient) GetEmpty(ctx context.Context) (result StringModel, err error) {
	req, err := client.GetEmptyPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client StringClient) GetEmptyPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client StringClient) GetEmptyResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMbcs get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ▇█〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€ '
func (client StringClient) GetMbcs(ctx context.Context) (result StringModel, err error) {
	req, err := client.GetMbcsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetMbcs", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMbcsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetMbcs", resp, "Failure sending request")
		return
	}

	result, err = client.GetMbcsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetMbcs", resp, "Failure responding to request")
	}

	return
}

// GetMbcsPreparer prepares the GetMbcs request.
func (client StringClient) GetMbcsPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/mbcs"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMbcsSender sends the GetMbcs request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetMbcsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetMbcsResponder handles the response to the GetMbcs request. The method always
// closes the http.Response Body.
func (client StringClient) GetMbcsResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNotProvided get String value when no string value is sent in response payload
func (client StringClient) GetNotProvided(ctx context.Context) (result StringModel, err error) {
	req, err := client.GetNotProvidedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNotProvided", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNotProvidedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNotProvided", resp, "Failure sending request")
		return
	}

	result, err = client.GetNotProvidedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNotProvided", resp, "Failure responding to request")
	}

	return
}

// GetNotProvidedPreparer prepares the GetNotProvided request.
func (client StringClient) GetNotProvidedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/notProvided"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNotProvidedSender sends the GetNotProvided request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetNotProvidedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNotProvidedResponder handles the response to the GetNotProvided request. The method always
// closes the http.Response Body.
func (client StringClient) GetNotProvidedResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null string value value
func (client StringClient) GetNull(ctx context.Context) (result StringModel, err error) {
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client StringClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client StringClient) GetNullResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNullBase64URLEncoded get null value that is expected to be base64url encoded
func (client StringClient) GetNullBase64URLEncoded(ctx context.Context) (result Base64URL, err error) {
	req, err := client.GetNullBase64URLEncodedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNullBase64URLEncoded", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullBase64URLEncodedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNullBase64URLEncoded", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullBase64URLEncodedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetNullBase64URLEncoded", resp, "Failure responding to request")
	}

	return
}

// GetNullBase64URLEncodedPreparer prepares the GetNullBase64URLEncoded request.
func (client StringClient) GetNullBase64URLEncodedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/nullBase64UrlEncoding"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullBase64URLEncodedSender sends the GetNullBase64URLEncoded request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetNullBase64URLEncodedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullBase64URLEncodedResponder handles the response to the GetNullBase64URLEncoded request. The method always
// closes the http.Response Body.
func (client StringClient) GetNullBase64URLEncodedResponder(resp *http.Response) (result Base64URL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetWhitespace get string value with leading and trailing whitespace '<tab><space><space>Now is the time for all good
// men to come to the aid of their country<tab><space><space>'
func (client StringClient) GetWhitespace(ctx context.Context) (result StringModel, err error) {
	req, err := client.GetWhitespacePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetWhitespace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetWhitespaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetWhitespace", resp, "Failure sending request")
		return
	}

	result, err = client.GetWhitespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "GetWhitespace", resp, "Failure responding to request")
	}

	return
}

// GetWhitespacePreparer prepares the GetWhitespace request.
func (client StringClient) GetWhitespacePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/whitespace"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetWhitespaceSender sends the GetWhitespace request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) GetWhitespaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetWhitespaceResponder handles the response to the GetWhitespace request. The method always
// closes the http.Response Body.
func (client StringClient) GetWhitespaceResponder(resp *http.Response) (result StringModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutBase64URLEncoded put value that is base64url encoded
//
func (client StringClient) PutBase64URLEncoded(ctx context.Context, stringBody string) (result autorest.Response, err error) {
	req, err := client.PutBase64URLEncodedPreparer(ctx, stringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutBase64URLEncoded", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutBase64URLEncodedSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutBase64URLEncoded", resp, "Failure sending request")
		return
	}

	result, err = client.PutBase64URLEncodedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutBase64URLEncoded", resp, "Failure responding to request")
	}

	return
}

// PutBase64URLEncodedPreparer prepares the PutBase64URLEncoded request.
func (client StringClient) PutBase64URLEncodedPreparer(ctx context.Context, stringBody string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/base64UrlEncoding"),
		autorest.WithJSON(stringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutBase64URLEncodedSender sends the PutBase64URLEncoded request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) PutBase64URLEncodedSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutBase64URLEncodedResponder handles the response to the PutBase64URLEncoded request. The method always
// closes the http.Response Body.
func (client StringClient) PutBase64URLEncodedResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutEmpty set string value empty ''
//
// stringBody is
func (client StringClient) PutEmpty(ctx context.Context, stringBody string) (result autorest.Response, err error) {
	req, err := client.PutEmptyPreparer(ctx, stringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutEmptySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.PutEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutEmpty", resp, "Failure responding to request")
	}

	return
}

// PutEmptyPreparer prepares the PutEmpty request.
func (client StringClient) PutEmptyPreparer(ctx context.Context, stringBody string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/empty"),
		autorest.WithJSON(stringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutEmptySender sends the PutEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) PutEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutEmptyResponder handles the response to the PutEmpty request. The method always
// closes the http.Response Body.
func (client StringClient) PutEmptyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMbcs set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ▇█〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€ '
//
// stringBody is
func (client StringClient) PutMbcs(ctx context.Context, stringBody string) (result autorest.Response, err error) {
	req, err := client.PutMbcsPreparer(ctx, stringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutMbcs", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMbcsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutMbcs", resp, "Failure sending request")
		return
	}

	result, err = client.PutMbcsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutMbcs", resp, "Failure responding to request")
	}

	return
}

// PutMbcsPreparer prepares the PutMbcs request.
func (client StringClient) PutMbcsPreparer(ctx context.Context, stringBody string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/mbcs"),
		autorest.WithJSON(stringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMbcsSender sends the PutMbcs request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) PutMbcsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMbcsResponder handles the response to the PutMbcs request. The method always
// closes the http.Response Body.
func (client StringClient) PutMbcsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutNull set string value null
//
// stringBody is
func (client StringClient) PutNull(ctx context.Context, stringBody string) (result autorest.Response, err error) {
	req, err := client.PutNullPreparer(ctx, stringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutNullSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutNull", resp, "Failure sending request")
		return
	}

	result, err = client.PutNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutNull", resp, "Failure responding to request")
	}

	return
}

// PutNullPreparer prepares the PutNull request.
func (client StringClient) PutNullPreparer(ctx context.Context, stringBody string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/null"))
	if len(string(stringBody)) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(stringBody))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutNullSender sends the PutNull request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) PutNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutNullResponder handles the response to the PutNull request. The method always
// closes the http.Response Body.
func (client StringClient) PutNullResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutWhitespace set String value with leading and trailing whitespace '<tab><space><space>Now is the time for all good
// men to come to the aid of their country<tab><space><space>'
//
// stringBody is
func (client StringClient) PutWhitespace(ctx context.Context, stringBody string) (result autorest.Response, err error) {
	req, err := client.PutWhitespacePreparer(ctx, stringBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutWhitespace", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutWhitespaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutWhitespace", resp, "Failure sending request")
		return
	}

	result, err = client.PutWhitespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "stringgroup.StringClient", "PutWhitespace", resp, "Failure responding to request")
	}

	return
}

// PutWhitespacePreparer prepares the PutWhitespace request.
func (client StringClient) PutWhitespacePreparer(ctx context.Context, stringBody string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/string/whitespace"),
		autorest.WithJSON(stringBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutWhitespaceSender sends the PutWhitespace request. The method will close the
// http.Response Body if it receives an error.
func (client StringClient) PutWhitespaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutWhitespaceResponder handles the response to the PutWhitespace request. The method always
// closes the http.Response Body.
func (client StringClient) PutWhitespaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
