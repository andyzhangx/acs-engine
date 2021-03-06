package containerregistry

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// RegistriesClient is the client for the Registries methods of the
// Containerregistry service.
type RegistriesClient struct {
	ManagementClient
}

// NewRegistriesClient creates an instance of the RegistriesClient client.
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return NewRegistriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegistriesClientWithBaseURI creates an instance of the RegistriesClient
// client.
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return RegistriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks whether the container registry name is
// available for use. The name must contain only alphanumeric characters, be
// globally unique, and between 5 and 60 characters in length.
//
// registryNameCheckRequest is the object containing information for the
// availability request.
func (client RegistriesClient) CheckNameAvailability(registryNameCheckRequest RegistryNameCheckRequest) (result RegistryNameStatus, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryNameCheckRequest,
			Constraints: []validation.Constraint{{Target: "registryNameCheckRequest.Name", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "registryNameCheckRequest.Name", Name: validation.MaxLength, Rule: 50, Chain: nil},
					{Target: "registryNameCheckRequest.Name", Name: validation.MinLength, Rule: 5, Chain: nil},
					{Target: "registryNameCheckRequest.Name", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil},
				}},
				{Target: "registryNameCheckRequest.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "CheckNameAvailability")
	}

	req, err := client.CheckNameAvailabilityPreparer(registryNameCheckRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client RegistriesClient) CheckNameAvailabilityPreparer(registryNameCheckRequest RegistryNameCheckRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/checkNameAvailability", pathParameters),
		autorest.WithJSON(registryNameCheckRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client RegistriesClient) CheckNameAvailabilityResponder(resp *http.Response) (result RegistryNameStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a container registry with the specified parameters. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
// registryCreateParameters is the parameters for creating a container
// registry.
func (client RegistriesClient) Create(resourceGroupName string, registryName string, registryCreateParameters RegistryCreateParameters, cancel <-chan struct{}) (<-chan Registry, <-chan error) {
	resultChan := make(chan Registry, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: registryCreateParameters,
			Constraints: []validation.Constraint{{Target: "registryCreateParameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "registryCreateParameters.Sku", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "registryCreateParameters.Sku.Name", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "registryCreateParameters.RegistryPropertiesCreateParameters", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "registryCreateParameters.RegistryPropertiesCreateParameters.StorageAccount", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "registryCreateParameters.RegistryPropertiesCreateParameters.StorageAccount.Name", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "registryCreateParameters.RegistryPropertiesCreateParameters.StorageAccount.AccessKey", Name: validation.Null, Rule: true, Chain: nil},
						}},
					}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result Registry
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, registryName, registryCreateParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client RegistriesClient) CreatePreparer(resourceGroupName string, registryName string, registryCreateParameters RegistryCreateParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithJSON(registryCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RegistriesClient) CreateResponder(resp *http.Response) (result Registry, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) Delete(resourceGroupName string, registryName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegistriesClient) DeletePreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegistriesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) Get(resourceGroupName string, registryName string) (result Registry, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegistriesClient) GetPreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegistriesClient) GetResponder(resp *http.Response) (result Registry, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the container registries under the specified subscription.
func (client RegistriesClient) List() (result RegistryListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RegistriesClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegistriesClient) ListResponder(resp *http.Response) (result RegistryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client RegistriesClient) ListNextResults(lastResults RegistryListResult) (result RegistryListResult, err error) {
	req, err := lastResults.RegistryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroup lists all the container registries under the specified
// resource group.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs.
func (client RegistriesClient) ListByResourceGroup(resourceGroupName string) (result RegistryListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client RegistriesClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client RegistriesClient) ListByResourceGroupResponder(resp *http.Response) (result RegistryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client RegistriesClient) ListByResourceGroupNextResults(lastResults RegistryListResult) (result RegistryListResult, err error) {
	req, err := lastResults.RegistryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListCredentials lists the login credentials for the specified container
// registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) ListCredentials(resourceGroupName string, registryName string) (result RegistryListCredentialsResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "ListCredentials")
	}

	req, err := client.ListCredentialsPreparer(resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListCredentials", resp, "Failure responding to request")
	}

	return
}

// ListCredentialsPreparer prepares the ListCredentials request.
func (client RegistriesClient) ListCredentialsPreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/listCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListCredentialsSender sends the ListCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) ListCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListCredentialsResponder handles the response to the ListCredentials request. The method always
// closes the http.Response Body.
func (client RegistriesClient) ListCredentialsResponder(resp *http.Response) (result RegistryListCredentialsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateCredential regenerates one of the login credentials for the
// specified container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
// regenerateCredentialParameters is specifies name of the password which
// should be regenerated -- password or password2.
func (client RegistriesClient) RegenerateCredential(resourceGroupName string, registryName string, regenerateCredentialParameters RegenerateCredentialParameters) (result RegistryListCredentialsResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "RegenerateCredential")
	}

	req, err := client.RegenerateCredentialPreparer(resourceGroupName, registryName, regenerateCredentialParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "RegenerateCredential", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateCredentialSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "RegenerateCredential", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateCredentialResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "RegenerateCredential", resp, "Failure responding to request")
	}

	return
}

// RegenerateCredentialPreparer prepares the RegenerateCredential request.
func (client RegistriesClient) RegenerateCredentialPreparer(resourceGroupName string, registryName string, regenerateCredentialParameters RegenerateCredentialParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/regenerateCredential", pathParameters),
		autorest.WithJSON(regenerateCredentialParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RegenerateCredentialSender sends the RegenerateCredential request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) RegenerateCredentialSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegenerateCredentialResponder handles the response to the RegenerateCredential request. The method always
// closes the http.Response Body.
func (client RegistriesClient) RegenerateCredentialResponder(resp *http.Response) (result RegistryListCredentialsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a container registry with the specified parameters.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
// registryUpdateParameters is the parameters for updating a container
// registry.
func (client RegistriesClient) Update(resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters) (result Registry, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, registryName, registryUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RegistriesClient) UpdatePreparer(resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithJSON(registryUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RegistriesClient) UpdateResponder(resp *http.Response) (result Registry, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
