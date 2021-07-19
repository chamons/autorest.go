// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// HSMSecurityDomainDownloadPoller provides polling facilities until the operation reaches a terminal state.
type HSMSecurityDomainDownloadPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final HSMSecurityDomainDownloadResponse will be returned.
	FinalResponse(ctx context.Context) (HSMSecurityDomainDownloadResponse, error)
}

type hsmSecurityDomainDownloadPoller struct {
	pt *azcore.LROPoller
}

func (p *hsmSecurityDomainDownloadPoller) Done() bool {
	return p.pt.Done()
}

func (p *hsmSecurityDomainDownloadPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *hsmSecurityDomainDownloadPoller) FinalResponse(ctx context.Context) (HSMSecurityDomainDownloadResponse, error) {
	respType := HSMSecurityDomainDownloadResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SecurityDomainObject)
	if err != nil {
		return HSMSecurityDomainDownloadResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *hsmSecurityDomainDownloadPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *hsmSecurityDomainDownloadPoller) pollUntilDone(ctx context.Context, freq time.Duration) (HSMSecurityDomainDownloadResponse, error) {
	respType := HSMSecurityDomainDownloadResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.SecurityDomainObject)
	if err != nil {
		return HSMSecurityDomainDownloadResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// HSMSecurityDomainUploadPoller provides polling facilities until the operation reaches a terminal state.
type HSMSecurityDomainUploadPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final HSMSecurityDomainUploadResponse will be returned.
	FinalResponse(ctx context.Context) (HSMSecurityDomainUploadResponse, error)
}

type hsmSecurityDomainUploadPoller struct {
	pt *azcore.LROPoller
}

func (p *hsmSecurityDomainUploadPoller) Done() bool {
	return p.pt.Done()
}

func (p *hsmSecurityDomainUploadPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *hsmSecurityDomainUploadPoller) FinalResponse(ctx context.Context) (HSMSecurityDomainUploadResponse, error) {
	respType := HSMSecurityDomainUploadResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SecurityDomainOperationStatus)
	if err != nil {
		return HSMSecurityDomainUploadResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *hsmSecurityDomainUploadPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *hsmSecurityDomainUploadPoller) pollUntilDone(ctx context.Context, freq time.Duration) (HSMSecurityDomainUploadResponse, error) {
	respType := HSMSecurityDomainUploadResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.SecurityDomainOperationStatus)
	if err != nil {
		return HSMSecurityDomainUploadResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// KeyVaultClientFullBackupPoller provides polling facilities until the operation reaches a terminal state.
type KeyVaultClientFullBackupPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final KeyVaultClientFullBackupResponse will be returned.
	FinalResponse(ctx context.Context) (KeyVaultClientFullBackupResponse, error)
}

type keyVaultClientFullBackupPoller struct {
	pt *azcore.LROPoller
}

func (p *keyVaultClientFullBackupPoller) Done() bool {
	return p.pt.Done()
}

func (p *keyVaultClientFullBackupPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *keyVaultClientFullBackupPoller) FinalResponse(ctx context.Context) (KeyVaultClientFullBackupResponse, error) {
	respType := KeyVaultClientFullBackupResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.FullBackupOperation)
	if err != nil {
		return KeyVaultClientFullBackupResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *keyVaultClientFullBackupPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *keyVaultClientFullBackupPoller) pollUntilDone(ctx context.Context, freq time.Duration) (KeyVaultClientFullBackupResponse, error) {
	respType := KeyVaultClientFullBackupResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.FullBackupOperation)
	if err != nil {
		return KeyVaultClientFullBackupResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// KeyVaultClientFullRestoreOperationPoller provides polling facilities until the operation reaches a terminal state.
type KeyVaultClientFullRestoreOperationPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final KeyVaultClientFullRestoreOperationResponse will be returned.
	FinalResponse(ctx context.Context) (KeyVaultClientFullRestoreOperationResponse, error)
}

type keyVaultClientFullRestoreOperationPoller struct {
	pt *azcore.LROPoller
}

func (p *keyVaultClientFullRestoreOperationPoller) Done() bool {
	return p.pt.Done()
}

func (p *keyVaultClientFullRestoreOperationPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *keyVaultClientFullRestoreOperationPoller) FinalResponse(ctx context.Context) (KeyVaultClientFullRestoreOperationResponse, error) {
	respType := KeyVaultClientFullRestoreOperationResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.RestoreOperation)
	if err != nil {
		return KeyVaultClientFullRestoreOperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *keyVaultClientFullRestoreOperationPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *keyVaultClientFullRestoreOperationPoller) pollUntilDone(ctx context.Context, freq time.Duration) (KeyVaultClientFullRestoreOperationResponse, error) {
	respType := KeyVaultClientFullRestoreOperationResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.RestoreOperation)
	if err != nil {
		return KeyVaultClientFullRestoreOperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// KeyVaultClientSelectiveKeyRestoreOperationPoller provides polling facilities until the operation reaches a terminal state.
type KeyVaultClientSelectiveKeyRestoreOperationPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final KeyVaultClientSelectiveKeyRestoreOperationResponse will be returned.
	FinalResponse(ctx context.Context) (KeyVaultClientSelectiveKeyRestoreOperationResponse, error)
}

type keyVaultClientSelectiveKeyRestoreOperationPoller struct {
	pt *azcore.LROPoller
}

func (p *keyVaultClientSelectiveKeyRestoreOperationPoller) Done() bool {
	return p.pt.Done()
}

func (p *keyVaultClientSelectiveKeyRestoreOperationPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *keyVaultClientSelectiveKeyRestoreOperationPoller) FinalResponse(ctx context.Context) (KeyVaultClientSelectiveKeyRestoreOperationResponse, error) {
	respType := KeyVaultClientSelectiveKeyRestoreOperationResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SelectiveKeyRestoreOperation)
	if err != nil {
		return KeyVaultClientSelectiveKeyRestoreOperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *keyVaultClientSelectiveKeyRestoreOperationPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *keyVaultClientSelectiveKeyRestoreOperationPoller) pollUntilDone(ctx context.Context, freq time.Duration) (KeyVaultClientSelectiveKeyRestoreOperationResponse, error) {
	respType := KeyVaultClientSelectiveKeyRestoreOperationResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.SelectiveKeyRestoreOperation)
	if err != nil {
		return KeyVaultClientSelectiveKeyRestoreOperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
