package automation

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ObjectDataTypesClient is the automation Client
type ObjectDataTypesClient struct {
	BaseClient
}

// NewObjectDataTypesClient creates an instance of the ObjectDataTypesClient client.
func NewObjectDataTypesClient(subscriptionID string, resourceGroupName string) ObjectDataTypesClient {
	return NewObjectDataTypesClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName)
}

// NewObjectDataTypesClientWithBaseURI creates an instance of the ObjectDataTypesClient client.
func NewObjectDataTypesClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) ObjectDataTypesClient {
	return ObjectDataTypesClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName)}
}

// ListFieldsByModuleAndType retrieve a list of fields of a given type identified by module name.
//
// automationAccountName is the automation account name. moduleName is the name of module. typeName is the name of
// type.
func (client ObjectDataTypesClient) ListFieldsByModuleAndType(ctx context.Context, automationAccountName string, moduleName string, typeName string) (result TypeFieldListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ObjectDataTypesClient", "ListFieldsByModuleAndType")
	}

	req, err := client.ListFieldsByModuleAndTypePreparer(ctx, automationAccountName, moduleName, typeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ObjectDataTypesClient", "ListFieldsByModuleAndType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListFieldsByModuleAndTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ObjectDataTypesClient", "ListFieldsByModuleAndType", resp, "Failure sending request")
		return
	}

	result, err = client.ListFieldsByModuleAndTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ObjectDataTypesClient", "ListFieldsByModuleAndType", resp, "Failure responding to request")
	}

	return
}

// ListFieldsByModuleAndTypePreparer prepares the ListFieldsByModuleAndType request.
func (client ObjectDataTypesClient) ListFieldsByModuleAndTypePreparer(ctx context.Context, automationAccountName string, moduleName string, typeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"typeName":              autorest.Encode("path", typeName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/objectDataTypes/{typeName}/fields", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListFieldsByModuleAndTypeSender sends the ListFieldsByModuleAndType request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectDataTypesClient) ListFieldsByModuleAndTypeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListFieldsByModuleAndTypeResponder handles the response to the ListFieldsByModuleAndType request. The method always
// closes the http.Response Body.
func (client ObjectDataTypesClient) ListFieldsByModuleAndTypeResponder(resp *http.Response) (result TypeFieldListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListFieldsByType retrieve a list of fields of a given type across all accessible modules.
//
// automationAccountName is the automation account name. typeName is the name of type.
func (client ObjectDataTypesClient) ListFieldsByType(ctx context.Context, automationAccountName string, typeName string) (result TypeFieldListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.ObjectDataTypesClient", "ListFieldsByType")
	}

	req, err := client.ListFieldsByTypePreparer(ctx, automationAccountName, typeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ObjectDataTypesClient", "ListFieldsByType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListFieldsByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.ObjectDataTypesClient", "ListFieldsByType", resp, "Failure sending request")
		return
	}

	result, err = client.ListFieldsByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.ObjectDataTypesClient", "ListFieldsByType", resp, "Failure responding to request")
	}

	return
}

// ListFieldsByTypePreparer prepares the ListFieldsByType request.
func (client ObjectDataTypesClient) ListFieldsByTypePreparer(ctx context.Context, automationAccountName string, typeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"typeName":              autorest.Encode("path", typeName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/objectDataTypes/{typeName}/fields", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListFieldsByTypeSender sends the ListFieldsByType request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectDataTypesClient) ListFieldsByTypeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListFieldsByTypeResponder handles the response to the ListFieldsByType request. The method always
// closes the http.Response Body.
func (client ObjectDataTypesClient) ListFieldsByTypeResponder(resp *http.Response) (result TypeFieldListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
