//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ActionRulesListByResourceGroupPager provides operations for iterating over paged responses.
type ActionRulesListByResourceGroupPager struct {
	client    *ActionRulesClient
	current   ActionRulesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ActionRulesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ActionRulesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ActionRulesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ActionRulesList.NextLink == nil || len(*p.current.ActionRulesList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ActionRulesListByResourceGroupResponse page.
func (p *ActionRulesListByResourceGroupPager) PageResponse() ActionRulesListByResourceGroupResponse {
	return p.current
}

// ActionRulesListBySubscriptionPager provides operations for iterating over paged responses.
type ActionRulesListBySubscriptionPager struct {
	client    *ActionRulesClient
	current   ActionRulesListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ActionRulesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ActionRulesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ActionRulesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ActionRulesList.NextLink == nil || len(*p.current.ActionRulesList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ActionRulesListBySubscriptionResponse page.
func (p *ActionRulesListBySubscriptionPager) PageResponse() ActionRulesListBySubscriptionResponse {
	return p.current
}

// AlertsGetAllPager provides operations for iterating over paged responses.
type AlertsGetAllPager struct {
	client    *AlertsClient
	current   AlertsGetAllResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AlertsGetAllResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AlertsGetAllPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AlertsGetAllPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AlertsList.NextLink == nil || len(*p.current.AlertsList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getAllHandleError(resp)
		return false
	}
	result, err := p.client.getAllHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AlertsGetAllResponse page.
func (p *AlertsGetAllPager) PageResponse() AlertsGetAllResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationsList.NextLink == nil || len(*p.current.OperationsList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
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

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// SmartDetectorAlertRulesListByResourceGroupPager provides operations for iterating over paged responses.
type SmartDetectorAlertRulesListByResourceGroupPager struct {
	client    *SmartDetectorAlertRulesClient
	current   SmartDetectorAlertRulesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SmartDetectorAlertRulesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SmartDetectorAlertRulesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SmartDetectorAlertRulesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AlertRulesList.NextLink == nil || len(*p.current.AlertRulesList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SmartDetectorAlertRulesListByResourceGroupResponse page.
func (p *SmartDetectorAlertRulesListByResourceGroupPager) PageResponse() SmartDetectorAlertRulesListByResourceGroupResponse {
	return p.current
}

// SmartDetectorAlertRulesListPager provides operations for iterating over paged responses.
type SmartDetectorAlertRulesListPager struct {
	client    *SmartDetectorAlertRulesClient
	current   SmartDetectorAlertRulesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SmartDetectorAlertRulesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SmartDetectorAlertRulesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SmartDetectorAlertRulesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AlertRulesList.NextLink == nil || len(*p.current.AlertRulesList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
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

// PageResponse returns the current SmartDetectorAlertRulesListResponse page.
func (p *SmartDetectorAlertRulesListPager) PageResponse() SmartDetectorAlertRulesListResponse {
	return p.current
}

// SmartGroupsGetAllPager provides operations for iterating over paged responses.
type SmartGroupsGetAllPager struct {
	client    *SmartGroupsClient
	current   SmartGroupsGetAllResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SmartGroupsGetAllResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SmartGroupsGetAllPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SmartGroupsGetAllPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SmartGroupsList.NextLink == nil || len(*p.current.SmartGroupsList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getAllHandleError(resp)
		return false
	}
	result, err := p.client.getAllHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SmartGroupsGetAllResponse page.
func (p *SmartGroupsGetAllPager) PageResponse() SmartGroupsGetAllResponse {
	return p.current
}
