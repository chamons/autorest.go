//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager provides operations for iterating over paged responses.
type DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager struct {
	client    *dataFlowDebugSessionClient
	current   DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.QueryDataFlowDebugSessionsResponse.NextLink == nil || len(*p.current.QueryDataFlowDebugSessionsResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.queryDataFlowDebugSessionsByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse page.
func (p *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager) PageResponse() DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse {
	return p.current
}

// DataFlowGetDataFlowsByWorkspacePager provides operations for iterating over paged responses.
type DataFlowGetDataFlowsByWorkspacePager struct {
	client    *dataFlowClient
	current   DataFlowGetDataFlowsByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataFlowGetDataFlowsByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataFlowGetDataFlowsByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataFlowGetDataFlowsByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataFlowListResponse.NextLink == nil || len(*p.current.DataFlowListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDataFlowsByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getDataFlowsByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataFlowGetDataFlowsByWorkspaceResponse page.
func (p *DataFlowGetDataFlowsByWorkspacePager) PageResponse() DataFlowGetDataFlowsByWorkspaceResponse {
	return p.current
}

// DatasetGetDatasetsByWorkspacePager provides operations for iterating over paged responses.
type DatasetGetDatasetsByWorkspacePager struct {
	client    *datasetClient
	current   DatasetGetDatasetsByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DatasetGetDatasetsByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DatasetGetDatasetsByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DatasetGetDatasetsByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatasetListResponse.NextLink == nil || len(*p.current.DatasetListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDatasetsByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getDatasetsByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DatasetGetDatasetsByWorkspaceResponse page.
func (p *DatasetGetDatasetsByWorkspacePager) PageResponse() DatasetGetDatasetsByWorkspaceResponse {
	return p.current
}

// LibraryListPager provides operations for iterating over paged responses.
type LibraryListPager struct {
	client    *libraryClient
	current   LibraryListResponseEnvelope
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LibraryListResponseEnvelope) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LibraryListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LibraryListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LibraryListResponse.NextLink == nil || len(*p.current.LibraryListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LibraryListResponseEnvelope page.
func (p *LibraryListPager) PageResponse() LibraryListResponseEnvelope {
	return p.current
}

// LinkedServiceGetLinkedServicesByWorkspacePager provides operations for iterating over paged responses.
type LinkedServiceGetLinkedServicesByWorkspacePager struct {
	client    *linkedServiceClient
	current   LinkedServiceGetLinkedServicesByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LinkedServiceGetLinkedServicesByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LinkedServiceGetLinkedServicesByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LinkedServiceGetLinkedServicesByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LinkedServiceListResponse.NextLink == nil || len(*p.current.LinkedServiceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getLinkedServicesByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getLinkedServicesByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LinkedServiceGetLinkedServicesByWorkspaceResponse page.
func (p *LinkedServiceGetLinkedServicesByWorkspacePager) PageResponse() LinkedServiceGetLinkedServicesByWorkspaceResponse {
	return p.current
}

// NotebookGetNotebookSummaryByWorkSpacePager provides operations for iterating over paged responses.
type NotebookGetNotebookSummaryByWorkSpacePager struct {
	client    *notebookClient
	current   NotebookGetNotebookSummaryByWorkSpaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NotebookGetNotebookSummaryByWorkSpaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NotebookGetNotebookSummaryByWorkSpacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NotebookGetNotebookSummaryByWorkSpacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NotebookListResponse.NextLink == nil || len(*p.current.NotebookListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getNotebookSummaryByWorkSpaceHandleError(resp)
		return false
	}
	result, err := p.client.getNotebookSummaryByWorkSpaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NotebookGetNotebookSummaryByWorkSpaceResponse page.
func (p *NotebookGetNotebookSummaryByWorkSpacePager) PageResponse() NotebookGetNotebookSummaryByWorkSpaceResponse {
	return p.current
}

// NotebookGetNotebooksByWorkspacePager provides operations for iterating over paged responses.
type NotebookGetNotebooksByWorkspacePager struct {
	client    *notebookClient
	current   NotebookGetNotebooksByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NotebookGetNotebooksByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NotebookGetNotebooksByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NotebookGetNotebooksByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NotebookListResponse.NextLink == nil || len(*p.current.NotebookListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getNotebooksByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getNotebooksByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NotebookGetNotebooksByWorkspaceResponse page.
func (p *NotebookGetNotebooksByWorkspacePager) PageResponse() NotebookGetNotebooksByWorkspaceResponse {
	return p.current
}

// PipelineGetPipelinesByWorkspacePager provides operations for iterating over paged responses.
type PipelineGetPipelinesByWorkspacePager struct {
	client    *pipelineClient
	current   PipelineGetPipelinesByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PipelineGetPipelinesByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PipelineGetPipelinesByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PipelineGetPipelinesByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PipelineListResponse.NextLink == nil || len(*p.current.PipelineListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getPipelinesByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getPipelinesByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PipelineGetPipelinesByWorkspaceResponse page.
func (p *PipelineGetPipelinesByWorkspacePager) PageResponse() PipelineGetPipelinesByWorkspaceResponse {
	return p.current
}

// SQLScriptGetSQLScriptsByWorkspacePager provides operations for iterating over paged responses.
type SQLScriptGetSQLScriptsByWorkspacePager struct {
	client    *sqlScriptClient
	current   SQLScriptGetSQLScriptsByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLScriptGetSQLScriptsByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLScriptGetSQLScriptsByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLScriptGetSQLScriptsByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLScriptsListResponse.NextLink == nil || len(*p.current.SQLScriptsListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getSQLScriptsByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getSQLScriptsByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLScriptGetSQLScriptsByWorkspaceResponse page.
func (p *SQLScriptGetSQLScriptsByWorkspacePager) PageResponse() SQLScriptGetSQLScriptsByWorkspaceResponse {
	return p.current
}

// SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager provides operations for iterating over paged responses.
type SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager struct {
	client    *sparkJobDefinitionClient
	current   SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SparkJobDefinitionsListResponse.NextLink == nil || len(*p.current.SparkJobDefinitionsListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getSparkJobDefinitionsByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getSparkJobDefinitionsByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse page.
func (p *SparkJobDefinitionGetSparkJobDefinitionsByWorkspacePager) PageResponse() SparkJobDefinitionGetSparkJobDefinitionsByWorkspaceResponse {
	return p.current
}

// TriggerGetTriggersByWorkspacePager provides operations for iterating over paged responses.
type TriggerGetTriggersByWorkspacePager struct {
	client    *triggerClient
	current   TriggerGetTriggersByWorkspaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TriggerGetTriggersByWorkspaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TriggerGetTriggersByWorkspacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TriggerGetTriggersByWorkspacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TriggerListResponse.NextLink == nil || len(*p.current.TriggerListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getTriggersByWorkspaceHandleError(resp)
		return false
	}
	result, err := p.client.getTriggersByWorkspaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TriggerGetTriggersByWorkspaceResponse page.
func (p *TriggerGetTriggersByWorkspacePager) PageResponse() TriggerGetTriggersByWorkspaceResponse {
	return p.current
}
