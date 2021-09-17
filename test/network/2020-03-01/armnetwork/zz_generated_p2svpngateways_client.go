//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// P2SVPNGatewaysClient contains the methods for the P2SVPNGateways group.
// Don't use this type directly, use NewP2SVPNGatewaysClient() instead.
type P2SVPNGatewaysClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewP2SVPNGatewaysClient creates a new instance of P2SVPNGatewaysClient with the specified values.
func NewP2SVPNGatewaysClient(con *arm.Connection, subscriptionID string) *P2SVPNGatewaysClient {
	return &P2SVPNGatewaysClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysBeginCreateOrUpdateOptions) (P2SVPNGatewaysCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return P2SVPNGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result := P2SVPNGatewaysCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return P2SVPNGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *P2SVPNGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters P2SVPNGateway, options *P2SVPNGatewaysBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, p2SVPNGatewayParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *P2SVPNGatewaysClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a virtual wan p2s vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginDeleteOptions) (P2SVPNGatewaysDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewaysDeletePollerResponse{}, err
	}
	result := P2SVPNGatewaysDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return P2SVPNGatewaysDeletePollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a virtual wan p2s vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *P2SVPNGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *P2SVPNGatewaysClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDisconnectP2SVPNConnections - Disconnect P2S vpn connections of the virtual wan P2SVpnGateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) BeginDisconnectP2SVPNConnections(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysBeginDisconnectP2SVPNConnectionsOptions) (P2SVPNGatewaysDisconnectP2SVPNConnectionsPollerResponse, error) {
	resp, err := client.disconnectP2SVPNConnections(ctx, resourceGroupName, p2SVPNGatewayName, request, options)
	if err != nil {
		return P2SVPNGatewaysDisconnectP2SVPNConnectionsPollerResponse{}, err
	}
	result := P2SVPNGatewaysDisconnectP2SVPNConnectionsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.DisconnectP2SVPNConnections", "location", resp, client.pl, client.disconnectP2SVPNConnectionsHandleError)
	if err != nil {
		return P2SVPNGatewaysDisconnectP2SVPNConnectionsPollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysDisconnectP2SVPNConnectionsPoller{
		pt: pt,
	}
	return result, nil
}

// DisconnectP2SVPNConnections - Disconnect P2S vpn connections of the virtual wan P2SVpnGateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnections(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysBeginDisconnectP2SVPNConnectionsOptions) (*http.Response, error) {
	req, err := client.disconnectP2SVPNConnectionsCreateRequest(ctx, resourceGroupName, p2SVPNGatewayName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.disconnectP2SVPNConnectionsHandleError(resp)
	}
	return resp, nil
}

// disconnectP2SVPNConnectionsCreateRequest creates the DisconnectP2SVPNConnections request.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnectionsCreateRequest(ctx context.Context, resourceGroupName string, p2SVPNGatewayName string, request P2SVPNConnectionRequest, options *P2SVPNGatewaysBeginDisconnectP2SVPNConnectionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{p2sVpnGatewayName}/disconnectP2sVpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if p2SVPNGatewayName == "" {
		return nil, errors.New("parameter p2SVPNGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{p2sVpnGatewayName}", url.PathEscape(p2SVPNGatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// disconnectP2SVPNConnectionsHandleError handles the DisconnectP2SVPNConnections error response.
func (client *P2SVPNGatewaysClient) disconnectP2SVPNConnectionsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginGenerateVPNProfile - Generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) BeginGenerateVPNProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysBeginGenerateVPNProfileOptions) (P2SVPNGatewaysGenerateVPNProfilePollerResponse, error) {
	resp, err := client.generateVPNProfile(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return P2SVPNGatewaysGenerateVPNProfilePollerResponse{}, err
	}
	result := P2SVPNGatewaysGenerateVPNProfilePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.GenerateVPNProfile", "location", resp, client.pl, client.generateVPNProfileHandleError)
	if err != nil {
		return P2SVPNGatewaysGenerateVPNProfilePollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysGenerateVPNProfilePoller{
		pt: pt,
	}
	return result, nil
}

// GenerateVPNProfile - Generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) generateVPNProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysBeginGenerateVPNProfileOptions) (*http.Response, error) {
	req, err := client.generateVPNProfileCreateRequest(ctx, resourceGroupName, gatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.generateVPNProfileHandleError(resp)
	}
	return resp, nil
}

// generateVPNProfileCreateRequest creates the GenerateVPNProfile request.
func (client *P2SVPNGatewaysClient) generateVPNProfileCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVPNProfileParameters, options *P2SVPNGatewaysBeginGenerateVPNProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/generatevpnprofile"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// generateVPNProfileHandleError handles the GenerateVPNProfile error response.
func (client *P2SVPNGatewaysClient) generateVPNProfileHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the details of a virtual wan p2s vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysGetOptions) (P2SVPNGatewaysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewaysGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return P2SVPNGatewaysGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return P2SVPNGatewaysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *P2SVPNGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *P2SVPNGatewaysClient) getHandleResponse(resp *http.Response) (P2SVPNGatewaysGetResponse, error) {
	result := P2SVPNGatewaysGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.P2SVPNGateway); err != nil {
		return P2SVPNGatewaysGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *P2SVPNGatewaysClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginGetP2SVPNConnectionHealth - Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) BeginGetP2SVPNConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthOptions) (P2SVPNGatewaysGetP2SVPNConnectionHealthPollerResponse, error) {
	resp, err := client.getP2SVPNConnectionHealth(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return P2SVPNGatewaysGetP2SVPNConnectionHealthPollerResponse{}, err
	}
	result := P2SVPNGatewaysGetP2SVPNConnectionHealthPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.GetP2SVPNConnectionHealth", "location", resp, client.pl, client.getP2SVPNConnectionHealthHandleError)
	if err != nil {
		return P2SVPNGatewaysGetP2SVPNConnectionHealthPollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysGetP2SVPNConnectionHealthPoller{
		pt: pt,
	}
	return result, nil
}

// GetP2SVPNConnectionHealth - Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthOptions) (*http.Response, error) {
	req, err := client.getP2SVPNConnectionHealthCreateRequest(ctx, resourceGroupName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.getP2SVPNConnectionHealthHandleError(resp)
	}
	return resp, nil
}

// getP2SVPNConnectionHealthCreateRequest creates the GetP2SVPNConnectionHealth request.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealth"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getP2SVPNConnectionHealthHandleError handles the GetP2SVPNConnectionHealth error response.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginGetP2SVPNConnectionHealthDetailed - Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the
// specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) BeginGetP2SVPNConnectionHealthDetailed(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthDetailedOptions) (P2SVPNGatewaysGetP2SVPNConnectionHealthDetailedPollerResponse, error) {
	resp, err := client.getP2SVPNConnectionHealthDetailed(ctx, resourceGroupName, gatewayName, request, options)
	if err != nil {
		return P2SVPNGatewaysGetP2SVPNConnectionHealthDetailedPollerResponse{}, err
	}
	result := P2SVPNGatewaysGetP2SVPNConnectionHealthDetailedPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("P2SVPNGatewaysClient.GetP2SVPNConnectionHealthDetailed", "location", resp, client.pl, client.getP2SVPNConnectionHealthDetailedHandleError)
	if err != nil {
		return P2SVPNGatewaysGetP2SVPNConnectionHealthDetailedPollerResponse{}, err
	}
	result.Poller = &P2SVPNGatewaysGetP2SVPNConnectionHealthDetailedPoller{
		pt: pt,
	}
	return result, nil
}

// GetP2SVPNConnectionHealthDetailed - Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified
// resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailed(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthDetailedOptions) (*http.Response, error) {
	req, err := client.getP2SVPNConnectionHealthDetailedCreateRequest(ctx, resourceGroupName, gatewayName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.getP2SVPNConnectionHealthDetailedHandleError(resp)
	}
	return resp, nil
}

// getP2SVPNConnectionHealthDetailedCreateRequest creates the GetP2SVPNConnectionHealthDetailed request.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailedCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, request P2SVPNConnectionHealthRequest, options *P2SVPNGatewaysBeginGetP2SVPNConnectionHealthDetailedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealthDetailed"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// getP2SVPNConnectionHealthDetailedHandleError handles the GetP2SVPNConnectionHealthDetailed error response.
func (client *P2SVPNGatewaysClient) getP2SVPNConnectionHealthDetailedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all the P2SVpnGateways in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) List(options *P2SVPNGatewaysListOptions) *P2SVPNGatewaysListPager {
	return &P2SVPNGatewaysListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp P2SVPNGatewaysListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListP2SVPNGatewaysResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *P2SVPNGatewaysClient) listCreateRequest(ctx context.Context, options *P2SVPNGatewaysListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *P2SVPNGatewaysClient) listHandleResponse(resp *http.Response) (P2SVPNGatewaysListResponse, error) {
	result := P2SVPNGatewaysListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListP2SVPNGatewaysResult); err != nil {
		return P2SVPNGatewaysListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *P2SVPNGatewaysClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists all the P2SVpnGateways in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) ListByResourceGroup(resourceGroupName string, options *P2SVPNGatewaysListByResourceGroupOptions) *P2SVPNGatewaysListByResourceGroupPager {
	return &P2SVPNGatewaysListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp P2SVPNGatewaysListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListP2SVPNGatewaysResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *P2SVPNGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *P2SVPNGatewaysListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *P2SVPNGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (P2SVPNGatewaysListByResourceGroupResponse, error) {
	result := P2SVPNGatewaysListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListP2SVPNGatewaysResult); err != nil {
		return P2SVPNGatewaysListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *P2SVPNGatewaysClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Updates virtual wan p2s vpn gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *P2SVPNGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters TagsObject, options *P2SVPNGatewaysUpdateTagsOptions) (P2SVPNGatewaysUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, gatewayName, p2SVPNGatewayParameters, options)
	if err != nil {
		return P2SVPNGatewaysUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return P2SVPNGatewaysUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return P2SVPNGatewaysUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *P2SVPNGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, p2SVPNGatewayParameters TagsObject, options *P2SVPNGatewaysUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, p2SVPNGatewayParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *P2SVPNGatewaysClient) updateTagsHandleResponse(resp *http.Response) (P2SVPNGatewaysUpdateTagsResponse, error) {
	result := P2SVPNGatewaysUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.P2SVPNGateway); err != nil {
		return P2SVPNGatewaysUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *P2SVPNGatewaysClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
