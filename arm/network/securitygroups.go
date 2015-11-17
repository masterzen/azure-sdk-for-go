package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// SecurityGroupsResourceProviderClient is the the Windows Azure Network
// management API provides a RESTful set of web services that interact with
// Windows Azure Networks service to manage your network resrources. The API
// has entities that capture the relationship between an end user and the
// Windows Azure Networks service.
type SecurityGroupsResourceProviderClient struct {
	ResourceProviderClient
}

// NewSecurityGroupsResourceProviderClient creates an instance of the
// SecurityGroupsResourceProviderClient client.
func NewSecurityGroupsResourceProviderClient(subscriptionID string) SecurityGroupsResourceProviderClient {
	return NewSecurityGroupsResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityGroupsResourceProviderClientWithBaseURI creates an instance of
// the SecurityGroupsResourceProviderClient client.
func NewSecurityGroupsResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) SecurityGroupsResourceProviderClient {
	return SecurityGroupsResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put NetworkSecurityGroup operation creates/updates a
// network security groupin the specified resource group.
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
// parameters is parameters supplied to the create/update Network Security
// Group operation
func (client SecurityGroupsResourceProviderClient) CreateOrUpdate(resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup) (result SecurityGroup, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, networkSecurityGroupName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityGroupsResourceProviderClient) CreateOrUpdatePreparer(resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsResourceProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityGroupsResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete NetworkSecurityGroup operation deletes the specifed
// network security group
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
func (client SecurityGroupsResourceProviderClient) Delete(resourceGroupName string, networkSecurityGroupName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, networkSecurityGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityGroupsResourceProviderClient) DeletePreparer(resourceGroupName string, networkSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsResourceProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecurityGroupsResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get NetworkSecurityGroups operation retrieves information about the
// specified network security group.
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
func (client SecurityGroupsResourceProviderClient) Get(resourceGroupName string, networkSecurityGroupName string) (result SecurityGroup, ae error) {
	req, err := client.GetPreparer(resourceGroupName, networkSecurityGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityGroupsResourceProviderClient) GetPreparer(resourceGroupName string, networkSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityGroupsResourceProviderClient) GetResponder(resp *http.Response) (result SecurityGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the list NetworkSecurityGroups returns all network security groups in
// a resource group
//
// resourceGroupName is the name of the resource group.
func (client SecurityGroupsResourceProviderClient) List(resourceGroupName string) (result SecurityGroupListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityGroupsResourceProviderClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityGroupsResourceProviderClient) ListResponder(resp *http.Response) (result SecurityGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client SecurityGroupsResourceProviderClient) ListNextResults(lastResults SecurityGroupListResult) (result SecurityGroupListResult, ae error) {
	req, err := lastResults.SecurityGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "List", "Failure responding to next results request request")
	}

	return
}

// ListAll the list NetworkSecurityGroups returns all network security groups
// in a subscription
func (client SecurityGroupsResourceProviderClient) ListAll() (result SecurityGroupListResult, ae error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "ListAll", "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "ListAll", "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "ListAll", "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client SecurityGroupsResourceProviderClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupsResourceProviderClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client SecurityGroupsResourceProviderClient) ListAllResponder(resp *http.Response) (result SecurityGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client SecurityGroupsResourceProviderClient) ListAllNextResults(lastResults SecurityGroupListResult) (result SecurityGroupListResult, ae error) {
	req, err := lastResults.SecurityGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "ListAll", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "ListAll", "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityGroupsResourceProviderClient", "ListAll", "Failure responding to next results request request")
	}

	return
}
