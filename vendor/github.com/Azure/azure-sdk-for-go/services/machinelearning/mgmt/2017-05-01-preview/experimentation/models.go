package experimentation

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// Account an object that represents a machine learning team account.
type Account struct {
	autorest.Response `json:"-"`
	// ID - The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags *map[string]*string `json:"tags,omitempty"`
	// AccountProperties - The properties of the machine learning team account.
	*AccountProperties `json:"properties,omitempty"`
}

// AccountListResult the result of a request to list machine learning team accounts.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of machine learning team accounts. Since this list may be incomplete, the nextLink field should be used to request the next list of machine learning team accounts.
	Value *[]Account `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next list of machine learning team accounts.
	NextLink *string `json:"nextLink,omitempty"`
}

// AccountListResultIterator provides access to a complete listing of Account values.
type AccountListResultIterator struct {
	i    int
	page AccountListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountListResultIterator) Response() AccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountListResultIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AccountListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// accountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AccountListResult) accountListResultPreparer() (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AccountListResultPage contains a page of Account values.
type AccountListResultPage struct {
	fn  func(AccountListResult) (AccountListResult, error)
	alr AccountListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListResultPage) Next() error {
	next, err := page.fn(page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountListResultPage) Response() AccountListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountListResultPage) Values() []Account {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// AccountProperties the properties of a machine learning team account.
type AccountProperties struct {
	// VsoAccountID - The fully qualified arm id of the vso account to be used for this team account.
	VsoAccountID *string `json:"vsoAccountId,omitempty"`
	// AccountID - The immutable id associated with this team account.
	AccountID *string `json:"accountId,omitempty"`
	// Description - The description of this workspace.
	Description *string `json:"description,omitempty"`
	// FriendlyName - The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `json:"friendlyName,omitempty"`
	// KeyVaultID - The fully qualified arm id of the user key vault.
	KeyVaultID *string `json:"keyVaultId,omitempty"`
	// Seats - The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats *string `json:"seats,omitempty"`
	// DiscoveryURI - The uri for this machine learning team account.
	DiscoveryURI *string `json:"discoveryUri,omitempty"`
	// CreationDate - The creation date of the machine learning team account in ISO8601 format.
	CreationDate *date.Time `json:"creationDate,omitempty"`
	// StorageAccount - The properties of the storage account for the machine learning team account.
	StorageAccount *StorageAccountProperties `json:"storageAccount,omitempty"`
	// ProvisioningState - The current deployment state of team account resource. The provisioningState is to indicate states for resource provisioning. Possible values include: 'Creating', 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// AccountPropertiesUpdateParameters the parameters for updating the properties of a machine learning team account.
type AccountPropertiesUpdateParameters struct {
	// Description - The description of this workspace.
	Description *string `json:"description,omitempty"`
	// FriendlyName - The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `json:"friendlyName,omitempty"`
	// Seats - The no of users/seats who can access this team account. This property defines the charge on the team account.
	Seats *string `json:"seats,omitempty"`
	// StorageAccountKey - The key for storage account associated with this team account
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`
}

// AccountUpdateParameters the parameters for updating a machine learning team account.
type AccountUpdateParameters struct {
	// Tags - The resource tags for the machine learning team account.
	Tags *map[string]*string `json:"tags,omitempty"`
	// AccountPropertiesUpdateParameters - The properties that the machine learning team account will be updated with.
	*AccountPropertiesUpdateParameters `json:"properties,omitempty"`
}

// ErrorResponse the error response send when an operation fails.
type ErrorResponse struct {
	// Code - error code
	Code *string `json:"code,omitempty"`
	// Message - error message
	Message *string `json:"message,omitempty"`
}

// Operation azure Machine Learning team account REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display name of operation
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay display name of operation
type OperationDisplay struct {
	// Provider - The resource provider name: Microsoft.MachineLearningExperimentation
	Provider *string `json:"provider,omitempty"`
	// Resource - The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - The operation that users can perform.
	Operation *string `json:"operation,omitempty"`
	// Description - The description for the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult an array of operations supported by the resource provider.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of AML team account operations supported by the AML team account resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// Project an object that represents a machine learning project.
type Project struct {
	autorest.Response `json:"-"`
	// ID - The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags *map[string]*string `json:"tags,omitempty"`
	// ProjectProperties - The properties of the Project.
	*ProjectProperties `json:"properties,omitempty"`
}

// ProjectListResult the result of a request to list projects.
type ProjectListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of projects. Since this list may be incomplete, the nextLink field should be used to request the next list of projects.
	Value *[]Project `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next list of projects.
	NextLink *string `json:"nextLink,omitempty"`
}

// ProjectListResultIterator provides access to a complete listing of Project values.
type ProjectListResultIterator struct {
	i    int
	page ProjectListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ProjectListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ProjectListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ProjectListResultIterator) Response() ProjectListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ProjectListResultIterator) Value() Project {
	if !iter.page.NotDone() {
		return Project{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (plr ProjectListResult) IsEmpty() bool {
	return plr.Value == nil || len(*plr.Value) == 0
}

// projectListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (plr ProjectListResult) projectListResultPreparer() (*http.Request, error) {
	if plr.NextLink == nil || len(to.String(plr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(plr.NextLink)))
}

// ProjectListResultPage contains a page of Project values.
type ProjectListResultPage struct {
	fn  func(ProjectListResult) (ProjectListResult, error)
	plr ProjectListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ProjectListResultPage) Next() error {
	next, err := page.fn(page.plr)
	if err != nil {
		return err
	}
	page.plr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ProjectListResultPage) NotDone() bool {
	return !page.plr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ProjectListResultPage) Response() ProjectListResult {
	return page.plr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ProjectListResultPage) Values() []Project {
	if page.plr.IsEmpty() {
		return nil
	}
	return *page.plr.Value
}

// ProjectProperties the properties of a machine learning project.
type ProjectProperties struct {
	// Description - The description of this project.
	Description *string `json:"description,omitempty"`
	// AccountID - The immutable id of the team account which contains this project.
	AccountID *string `json:"accountId,omitempty"`
	// WorkspaceID - The immutable id of the workspace which contains this project.
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// ProjectID - The immutable id of this project.
	ProjectID *string `json:"projectId,omitempty"`
	// Gitrepo - The reference to git repo for this project.
	Gitrepo *string `json:"gitrepo,omitempty"`
	// FriendlyName - The friendly name for this project.
	FriendlyName *string `json:"friendlyName,omitempty"`
	// CreationDate - The creation date of the project in ISO8601 format.
	CreationDate *date.Time `json:"creationDate,omitempty"`
	// ProvisioningState - The current deployment state of project resource. The provisioningState is to indicate states for resource provisioning. Possible values include: 'Creating', 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// ProjectPropertiesUpdateParameters the parameters for updating the properties of a project.
type ProjectPropertiesUpdateParameters struct {
	// FriendlyName - The friendly name for this project.
	FriendlyName *string `json:"friendlyName,omitempty"`
	// Description - The description of this project.
	Description *string `json:"description,omitempty"`
	// Gitrepo - The reference to git repo for this project.
	Gitrepo *string `json:"gitrepo,omitempty"`
}

// ProjectUpdateParameters the parameters for updating a machine learning project.
type ProjectUpdateParameters struct {
	// Tags - The resource tags for the machine learning project.
	Tags *map[string]*string `json:"tags,omitempty"`
	// ProjectPropertiesUpdateParameters - The properties that the project will be updated with.
	*ProjectPropertiesUpdateParameters `json:"properties,omitempty"`
}

// Resource an Azure resource.
type Resource struct {
	// ID - The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags *map[string]*string `json:"tags,omitempty"`
}

// StorageAccountProperties the properties of a storage account for a machine learning team account.
type StorageAccountProperties struct {
	// StorageAccountID - The fully qualified arm Id of the storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
	// AccessKey - The access key to the storage account.
	AccessKey *string `json:"accessKey,omitempty"`
}

// Workspace an object that represents a machine learning team account workspace.
type Workspace struct {
	autorest.Response `json:"-"`
	// ID - The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags *map[string]*string `json:"tags,omitempty"`
	// WorkspaceProperties - The properties of the machine learning team account workspace.
	*WorkspaceProperties `json:"properties,omitempty"`
}

// WorkspaceListResult the result of a request to list machine learning team account workspaces.
type WorkspaceListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of machine learning team account workspaces. Since this list may be incomplete, the nextLink field should be used to request the next list of machine learning team accounts.
	Value *[]Workspace `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next list of machine learning workspaces.
	NextLink *string `json:"nextLink,omitempty"`
}

// WorkspaceListResultIterator provides access to a complete listing of Workspace values.
type WorkspaceListResultIterator struct {
	i    int
	page WorkspaceListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *WorkspaceListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter WorkspaceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter WorkspaceListResultIterator) Response() WorkspaceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter WorkspaceListResultIterator) Value() Workspace {
	if !iter.page.NotDone() {
		return Workspace{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (wlr WorkspaceListResult) IsEmpty() bool {
	return wlr.Value == nil || len(*wlr.Value) == 0
}

// workspaceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (wlr WorkspaceListResult) workspaceListResultPreparer() (*http.Request, error) {
	if wlr.NextLink == nil || len(to.String(wlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(wlr.NextLink)))
}

// WorkspaceListResultPage contains a page of Workspace values.
type WorkspaceListResultPage struct {
	fn  func(WorkspaceListResult) (WorkspaceListResult, error)
	wlr WorkspaceListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *WorkspaceListResultPage) Next() error {
	next, err := page.fn(page.wlr)
	if err != nil {
		return err
	}
	page.wlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page WorkspaceListResultPage) NotDone() bool {
	return !page.wlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page WorkspaceListResultPage) Response() WorkspaceListResult {
	return page.wlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page WorkspaceListResultPage) Values() []Workspace {
	if page.wlr.IsEmpty() {
		return nil
	}
	return *page.wlr.Value
}

// WorkspaceProperties the properties of a machine learning team account workspace.
type WorkspaceProperties struct {
	// Description - The description of this workspace.
	Description *string `json:"description,omitempty"`
	// AccountID - The immutable id of the team account which contains this workspace.
	AccountID *string `json:"accountId,omitempty"`
	// WorkspaceID - The immutable id of this workspace.
	WorkspaceID *string `json:"workspaceId,omitempty"`
	// FriendlyName - The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName *string `json:"friendlyName,omitempty"`
	// CreationDate - The creation date of the machine learning workspace in ISO8601 format.
	CreationDate *date.Time `json:"creationDate,omitempty"`
	// ProvisioningState - The current deployment state of team account workspace resource. The provisioningState is to indicate states for resource provisioning. Possible values include: 'Creating', 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// WorkspacePropertiesUpdateParameters the parameters for updating the properties of a machine learning team account
// workspace.
type WorkspacePropertiesUpdateParameters struct {
	// FriendlyName - Friendly name of this workspace.
	FriendlyName *string `json:"friendlyName,omitempty"`
	// Description - Description for this workspace.
	Description *string `json:"description,omitempty"`
}

// WorkspaceUpdateParameters the parameters for updating a machine learning team account workspace.
type WorkspaceUpdateParameters struct {
	// Tags - The resource tags for the machine learning team account workspace.
	Tags *map[string]*string `json:"tags,omitempty"`
	// WorkspacePropertiesUpdateParameters - The properties that the machine learning workspace will be updated with.
	*WorkspacePropertiesUpdateParameters `json:"properties,omitempty"`
}
