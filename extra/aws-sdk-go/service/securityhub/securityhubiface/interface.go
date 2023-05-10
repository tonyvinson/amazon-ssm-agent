// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package securityhubiface provides an interface to enable mocking the AWS SecurityHub service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package securityhubiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/securityhub"
)

// SecurityHubAPI provides an interface to enable mocking the
// securityhub.SecurityHub service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS SecurityHub.
//    func myFunc(svc securityhubiface.SecurityHubAPI) bool {
//        // Make svc.AcceptAdministratorInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := securityhub.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSecurityHubClient struct {
//        securityhubiface.SecurityHubAPI
//    }
//    func (m *mockSecurityHubClient) AcceptAdministratorInvitation(input *securityhub.AcceptAdministratorInvitationInput) (*securityhub.AcceptAdministratorInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSecurityHubClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SecurityHubAPI interface {
	AcceptAdministratorInvitation(*securityhub.AcceptAdministratorInvitationInput) (*securityhub.AcceptAdministratorInvitationOutput, error)
	AcceptAdministratorInvitationWithContext(aws.Context, *securityhub.AcceptAdministratorInvitationInput, ...request.Option) (*securityhub.AcceptAdministratorInvitationOutput, error)
	AcceptAdministratorInvitationRequest(*securityhub.AcceptAdministratorInvitationInput) (*request.Request, *securityhub.AcceptAdministratorInvitationOutput)

	AcceptInvitation(*securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error)
	AcceptInvitationWithContext(aws.Context, *securityhub.AcceptInvitationInput, ...request.Option) (*securityhub.AcceptInvitationOutput, error)
	AcceptInvitationRequest(*securityhub.AcceptInvitationInput) (*request.Request, *securityhub.AcceptInvitationOutput)

	BatchDisableStandards(*securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error)
	BatchDisableStandardsWithContext(aws.Context, *securityhub.BatchDisableStandardsInput, ...request.Option) (*securityhub.BatchDisableStandardsOutput, error)
	BatchDisableStandardsRequest(*securityhub.BatchDisableStandardsInput) (*request.Request, *securityhub.BatchDisableStandardsOutput)

	BatchEnableStandards(*securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error)
	BatchEnableStandardsWithContext(aws.Context, *securityhub.BatchEnableStandardsInput, ...request.Option) (*securityhub.BatchEnableStandardsOutput, error)
	BatchEnableStandardsRequest(*securityhub.BatchEnableStandardsInput) (*request.Request, *securityhub.BatchEnableStandardsOutput)

	BatchGetSecurityControls(*securityhub.BatchGetSecurityControlsInput) (*securityhub.BatchGetSecurityControlsOutput, error)
	BatchGetSecurityControlsWithContext(aws.Context, *securityhub.BatchGetSecurityControlsInput, ...request.Option) (*securityhub.BatchGetSecurityControlsOutput, error)
	BatchGetSecurityControlsRequest(*securityhub.BatchGetSecurityControlsInput) (*request.Request, *securityhub.BatchGetSecurityControlsOutput)

	BatchGetStandardsControlAssociations(*securityhub.BatchGetStandardsControlAssociationsInput) (*securityhub.BatchGetStandardsControlAssociationsOutput, error)
	BatchGetStandardsControlAssociationsWithContext(aws.Context, *securityhub.BatchGetStandardsControlAssociationsInput, ...request.Option) (*securityhub.BatchGetStandardsControlAssociationsOutput, error)
	BatchGetStandardsControlAssociationsRequest(*securityhub.BatchGetStandardsControlAssociationsInput) (*request.Request, *securityhub.BatchGetStandardsControlAssociationsOutput)

	BatchImportFindings(*securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error)
	BatchImportFindingsWithContext(aws.Context, *securityhub.BatchImportFindingsInput, ...request.Option) (*securityhub.BatchImportFindingsOutput, error)
	BatchImportFindingsRequest(*securityhub.BatchImportFindingsInput) (*request.Request, *securityhub.BatchImportFindingsOutput)

	BatchUpdateFindings(*securityhub.BatchUpdateFindingsInput) (*securityhub.BatchUpdateFindingsOutput, error)
	BatchUpdateFindingsWithContext(aws.Context, *securityhub.BatchUpdateFindingsInput, ...request.Option) (*securityhub.BatchUpdateFindingsOutput, error)
	BatchUpdateFindingsRequest(*securityhub.BatchUpdateFindingsInput) (*request.Request, *securityhub.BatchUpdateFindingsOutput)

	BatchUpdateStandardsControlAssociations(*securityhub.BatchUpdateStandardsControlAssociationsInput) (*securityhub.BatchUpdateStandardsControlAssociationsOutput, error)
	BatchUpdateStandardsControlAssociationsWithContext(aws.Context, *securityhub.BatchUpdateStandardsControlAssociationsInput, ...request.Option) (*securityhub.BatchUpdateStandardsControlAssociationsOutput, error)
	BatchUpdateStandardsControlAssociationsRequest(*securityhub.BatchUpdateStandardsControlAssociationsInput) (*request.Request, *securityhub.BatchUpdateStandardsControlAssociationsOutput)

	CreateActionTarget(*securityhub.CreateActionTargetInput) (*securityhub.CreateActionTargetOutput, error)
	CreateActionTargetWithContext(aws.Context, *securityhub.CreateActionTargetInput, ...request.Option) (*securityhub.CreateActionTargetOutput, error)
	CreateActionTargetRequest(*securityhub.CreateActionTargetInput) (*request.Request, *securityhub.CreateActionTargetOutput)

	CreateFindingAggregator(*securityhub.CreateFindingAggregatorInput) (*securityhub.CreateFindingAggregatorOutput, error)
	CreateFindingAggregatorWithContext(aws.Context, *securityhub.CreateFindingAggregatorInput, ...request.Option) (*securityhub.CreateFindingAggregatorOutput, error)
	CreateFindingAggregatorRequest(*securityhub.CreateFindingAggregatorInput) (*request.Request, *securityhub.CreateFindingAggregatorOutput)

	CreateInsight(*securityhub.CreateInsightInput) (*securityhub.CreateInsightOutput, error)
	CreateInsightWithContext(aws.Context, *securityhub.CreateInsightInput, ...request.Option) (*securityhub.CreateInsightOutput, error)
	CreateInsightRequest(*securityhub.CreateInsightInput) (*request.Request, *securityhub.CreateInsightOutput)

	CreateMembers(*securityhub.CreateMembersInput) (*securityhub.CreateMembersOutput, error)
	CreateMembersWithContext(aws.Context, *securityhub.CreateMembersInput, ...request.Option) (*securityhub.CreateMembersOutput, error)
	CreateMembersRequest(*securityhub.CreateMembersInput) (*request.Request, *securityhub.CreateMembersOutput)

	DeclineInvitations(*securityhub.DeclineInvitationsInput) (*securityhub.DeclineInvitationsOutput, error)
	DeclineInvitationsWithContext(aws.Context, *securityhub.DeclineInvitationsInput, ...request.Option) (*securityhub.DeclineInvitationsOutput, error)
	DeclineInvitationsRequest(*securityhub.DeclineInvitationsInput) (*request.Request, *securityhub.DeclineInvitationsOutput)

	DeleteActionTarget(*securityhub.DeleteActionTargetInput) (*securityhub.DeleteActionTargetOutput, error)
	DeleteActionTargetWithContext(aws.Context, *securityhub.DeleteActionTargetInput, ...request.Option) (*securityhub.DeleteActionTargetOutput, error)
	DeleteActionTargetRequest(*securityhub.DeleteActionTargetInput) (*request.Request, *securityhub.DeleteActionTargetOutput)

	DeleteFindingAggregator(*securityhub.DeleteFindingAggregatorInput) (*securityhub.DeleteFindingAggregatorOutput, error)
	DeleteFindingAggregatorWithContext(aws.Context, *securityhub.DeleteFindingAggregatorInput, ...request.Option) (*securityhub.DeleteFindingAggregatorOutput, error)
	DeleteFindingAggregatorRequest(*securityhub.DeleteFindingAggregatorInput) (*request.Request, *securityhub.DeleteFindingAggregatorOutput)

	DeleteInsight(*securityhub.DeleteInsightInput) (*securityhub.DeleteInsightOutput, error)
	DeleteInsightWithContext(aws.Context, *securityhub.DeleteInsightInput, ...request.Option) (*securityhub.DeleteInsightOutput, error)
	DeleteInsightRequest(*securityhub.DeleteInsightInput) (*request.Request, *securityhub.DeleteInsightOutput)

	DeleteInvitations(*securityhub.DeleteInvitationsInput) (*securityhub.DeleteInvitationsOutput, error)
	DeleteInvitationsWithContext(aws.Context, *securityhub.DeleteInvitationsInput, ...request.Option) (*securityhub.DeleteInvitationsOutput, error)
	DeleteInvitationsRequest(*securityhub.DeleteInvitationsInput) (*request.Request, *securityhub.DeleteInvitationsOutput)

	DeleteMembers(*securityhub.DeleteMembersInput) (*securityhub.DeleteMembersOutput, error)
	DeleteMembersWithContext(aws.Context, *securityhub.DeleteMembersInput, ...request.Option) (*securityhub.DeleteMembersOutput, error)
	DeleteMembersRequest(*securityhub.DeleteMembersInput) (*request.Request, *securityhub.DeleteMembersOutput)

	DescribeActionTargets(*securityhub.DescribeActionTargetsInput) (*securityhub.DescribeActionTargetsOutput, error)
	DescribeActionTargetsWithContext(aws.Context, *securityhub.DescribeActionTargetsInput, ...request.Option) (*securityhub.DescribeActionTargetsOutput, error)
	DescribeActionTargetsRequest(*securityhub.DescribeActionTargetsInput) (*request.Request, *securityhub.DescribeActionTargetsOutput)

	DescribeActionTargetsPages(*securityhub.DescribeActionTargetsInput, func(*securityhub.DescribeActionTargetsOutput, bool) bool) error
	DescribeActionTargetsPagesWithContext(aws.Context, *securityhub.DescribeActionTargetsInput, func(*securityhub.DescribeActionTargetsOutput, bool) bool, ...request.Option) error

	DescribeHub(*securityhub.DescribeHubInput) (*securityhub.DescribeHubOutput, error)
	DescribeHubWithContext(aws.Context, *securityhub.DescribeHubInput, ...request.Option) (*securityhub.DescribeHubOutput, error)
	DescribeHubRequest(*securityhub.DescribeHubInput) (*request.Request, *securityhub.DescribeHubOutput)

	DescribeOrganizationConfiguration(*securityhub.DescribeOrganizationConfigurationInput) (*securityhub.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationWithContext(aws.Context, *securityhub.DescribeOrganizationConfigurationInput, ...request.Option) (*securityhub.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationRequest(*securityhub.DescribeOrganizationConfigurationInput) (*request.Request, *securityhub.DescribeOrganizationConfigurationOutput)

	DescribeProducts(*securityhub.DescribeProductsInput) (*securityhub.DescribeProductsOutput, error)
	DescribeProductsWithContext(aws.Context, *securityhub.DescribeProductsInput, ...request.Option) (*securityhub.DescribeProductsOutput, error)
	DescribeProductsRequest(*securityhub.DescribeProductsInput) (*request.Request, *securityhub.DescribeProductsOutput)

	DescribeProductsPages(*securityhub.DescribeProductsInput, func(*securityhub.DescribeProductsOutput, bool) bool) error
	DescribeProductsPagesWithContext(aws.Context, *securityhub.DescribeProductsInput, func(*securityhub.DescribeProductsOutput, bool) bool, ...request.Option) error

	DescribeStandards(*securityhub.DescribeStandardsInput) (*securityhub.DescribeStandardsOutput, error)
	DescribeStandardsWithContext(aws.Context, *securityhub.DescribeStandardsInput, ...request.Option) (*securityhub.DescribeStandardsOutput, error)
	DescribeStandardsRequest(*securityhub.DescribeStandardsInput) (*request.Request, *securityhub.DescribeStandardsOutput)

	DescribeStandardsPages(*securityhub.DescribeStandardsInput, func(*securityhub.DescribeStandardsOutput, bool) bool) error
	DescribeStandardsPagesWithContext(aws.Context, *securityhub.DescribeStandardsInput, func(*securityhub.DescribeStandardsOutput, bool) bool, ...request.Option) error

	DescribeStandardsControls(*securityhub.DescribeStandardsControlsInput) (*securityhub.DescribeStandardsControlsOutput, error)
	DescribeStandardsControlsWithContext(aws.Context, *securityhub.DescribeStandardsControlsInput, ...request.Option) (*securityhub.DescribeStandardsControlsOutput, error)
	DescribeStandardsControlsRequest(*securityhub.DescribeStandardsControlsInput) (*request.Request, *securityhub.DescribeStandardsControlsOutput)

	DescribeStandardsControlsPages(*securityhub.DescribeStandardsControlsInput, func(*securityhub.DescribeStandardsControlsOutput, bool) bool) error
	DescribeStandardsControlsPagesWithContext(aws.Context, *securityhub.DescribeStandardsControlsInput, func(*securityhub.DescribeStandardsControlsOutput, bool) bool, ...request.Option) error

	DisableImportFindingsForProduct(*securityhub.DisableImportFindingsForProductInput) (*securityhub.DisableImportFindingsForProductOutput, error)
	DisableImportFindingsForProductWithContext(aws.Context, *securityhub.DisableImportFindingsForProductInput, ...request.Option) (*securityhub.DisableImportFindingsForProductOutput, error)
	DisableImportFindingsForProductRequest(*securityhub.DisableImportFindingsForProductInput) (*request.Request, *securityhub.DisableImportFindingsForProductOutput)

	DisableOrganizationAdminAccount(*securityhub.DisableOrganizationAdminAccountInput) (*securityhub.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountWithContext(aws.Context, *securityhub.DisableOrganizationAdminAccountInput, ...request.Option) (*securityhub.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountRequest(*securityhub.DisableOrganizationAdminAccountInput) (*request.Request, *securityhub.DisableOrganizationAdminAccountOutput)

	DisableSecurityHub(*securityhub.DisableSecurityHubInput) (*securityhub.DisableSecurityHubOutput, error)
	DisableSecurityHubWithContext(aws.Context, *securityhub.DisableSecurityHubInput, ...request.Option) (*securityhub.DisableSecurityHubOutput, error)
	DisableSecurityHubRequest(*securityhub.DisableSecurityHubInput) (*request.Request, *securityhub.DisableSecurityHubOutput)

	DisassociateFromAdministratorAccount(*securityhub.DisassociateFromAdministratorAccountInput) (*securityhub.DisassociateFromAdministratorAccountOutput, error)
	DisassociateFromAdministratorAccountWithContext(aws.Context, *securityhub.DisassociateFromAdministratorAccountInput, ...request.Option) (*securityhub.DisassociateFromAdministratorAccountOutput, error)
	DisassociateFromAdministratorAccountRequest(*securityhub.DisassociateFromAdministratorAccountInput) (*request.Request, *securityhub.DisassociateFromAdministratorAccountOutput)

	DisassociateFromMasterAccount(*securityhub.DisassociateFromMasterAccountInput) (*securityhub.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountWithContext(aws.Context, *securityhub.DisassociateFromMasterAccountInput, ...request.Option) (*securityhub.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountRequest(*securityhub.DisassociateFromMasterAccountInput) (*request.Request, *securityhub.DisassociateFromMasterAccountOutput)

	DisassociateMembers(*securityhub.DisassociateMembersInput) (*securityhub.DisassociateMembersOutput, error)
	DisassociateMembersWithContext(aws.Context, *securityhub.DisassociateMembersInput, ...request.Option) (*securityhub.DisassociateMembersOutput, error)
	DisassociateMembersRequest(*securityhub.DisassociateMembersInput) (*request.Request, *securityhub.DisassociateMembersOutput)

	EnableImportFindingsForProduct(*securityhub.EnableImportFindingsForProductInput) (*securityhub.EnableImportFindingsForProductOutput, error)
	EnableImportFindingsForProductWithContext(aws.Context, *securityhub.EnableImportFindingsForProductInput, ...request.Option) (*securityhub.EnableImportFindingsForProductOutput, error)
	EnableImportFindingsForProductRequest(*securityhub.EnableImportFindingsForProductInput) (*request.Request, *securityhub.EnableImportFindingsForProductOutput)

	EnableOrganizationAdminAccount(*securityhub.EnableOrganizationAdminAccountInput) (*securityhub.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountWithContext(aws.Context, *securityhub.EnableOrganizationAdminAccountInput, ...request.Option) (*securityhub.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountRequest(*securityhub.EnableOrganizationAdminAccountInput) (*request.Request, *securityhub.EnableOrganizationAdminAccountOutput)

	EnableSecurityHub(*securityhub.EnableSecurityHubInput) (*securityhub.EnableSecurityHubOutput, error)
	EnableSecurityHubWithContext(aws.Context, *securityhub.EnableSecurityHubInput, ...request.Option) (*securityhub.EnableSecurityHubOutput, error)
	EnableSecurityHubRequest(*securityhub.EnableSecurityHubInput) (*request.Request, *securityhub.EnableSecurityHubOutput)

	GetAdministratorAccount(*securityhub.GetAdministratorAccountInput) (*securityhub.GetAdministratorAccountOutput, error)
	GetAdministratorAccountWithContext(aws.Context, *securityhub.GetAdministratorAccountInput, ...request.Option) (*securityhub.GetAdministratorAccountOutput, error)
	GetAdministratorAccountRequest(*securityhub.GetAdministratorAccountInput) (*request.Request, *securityhub.GetAdministratorAccountOutput)

	GetEnabledStandards(*securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error)
	GetEnabledStandardsWithContext(aws.Context, *securityhub.GetEnabledStandardsInput, ...request.Option) (*securityhub.GetEnabledStandardsOutput, error)
	GetEnabledStandardsRequest(*securityhub.GetEnabledStandardsInput) (*request.Request, *securityhub.GetEnabledStandardsOutput)

	GetEnabledStandardsPages(*securityhub.GetEnabledStandardsInput, func(*securityhub.GetEnabledStandardsOutput, bool) bool) error
	GetEnabledStandardsPagesWithContext(aws.Context, *securityhub.GetEnabledStandardsInput, func(*securityhub.GetEnabledStandardsOutput, bool) bool, ...request.Option) error

	GetFindingAggregator(*securityhub.GetFindingAggregatorInput) (*securityhub.GetFindingAggregatorOutput, error)
	GetFindingAggregatorWithContext(aws.Context, *securityhub.GetFindingAggregatorInput, ...request.Option) (*securityhub.GetFindingAggregatorOutput, error)
	GetFindingAggregatorRequest(*securityhub.GetFindingAggregatorInput) (*request.Request, *securityhub.GetFindingAggregatorOutput)

	GetFindingHistory(*securityhub.GetFindingHistoryInput) (*securityhub.GetFindingHistoryOutput, error)
	GetFindingHistoryWithContext(aws.Context, *securityhub.GetFindingHistoryInput, ...request.Option) (*securityhub.GetFindingHistoryOutput, error)
	GetFindingHistoryRequest(*securityhub.GetFindingHistoryInput) (*request.Request, *securityhub.GetFindingHistoryOutput)

	GetFindingHistoryPages(*securityhub.GetFindingHistoryInput, func(*securityhub.GetFindingHistoryOutput, bool) bool) error
	GetFindingHistoryPagesWithContext(aws.Context, *securityhub.GetFindingHistoryInput, func(*securityhub.GetFindingHistoryOutput, bool) bool, ...request.Option) error

	GetFindings(*securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error)
	GetFindingsWithContext(aws.Context, *securityhub.GetFindingsInput, ...request.Option) (*securityhub.GetFindingsOutput, error)
	GetFindingsRequest(*securityhub.GetFindingsInput) (*request.Request, *securityhub.GetFindingsOutput)

	GetFindingsPages(*securityhub.GetFindingsInput, func(*securityhub.GetFindingsOutput, bool) bool) error
	GetFindingsPagesWithContext(aws.Context, *securityhub.GetFindingsInput, func(*securityhub.GetFindingsOutput, bool) bool, ...request.Option) error

	GetInsightResults(*securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error)
	GetInsightResultsWithContext(aws.Context, *securityhub.GetInsightResultsInput, ...request.Option) (*securityhub.GetInsightResultsOutput, error)
	GetInsightResultsRequest(*securityhub.GetInsightResultsInput) (*request.Request, *securityhub.GetInsightResultsOutput)

	GetInsights(*securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error)
	GetInsightsWithContext(aws.Context, *securityhub.GetInsightsInput, ...request.Option) (*securityhub.GetInsightsOutput, error)
	GetInsightsRequest(*securityhub.GetInsightsInput) (*request.Request, *securityhub.GetInsightsOutput)

	GetInsightsPages(*securityhub.GetInsightsInput, func(*securityhub.GetInsightsOutput, bool) bool) error
	GetInsightsPagesWithContext(aws.Context, *securityhub.GetInsightsInput, func(*securityhub.GetInsightsOutput, bool) bool, ...request.Option) error

	GetInvitationsCount(*securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error)
	GetInvitationsCountWithContext(aws.Context, *securityhub.GetInvitationsCountInput, ...request.Option) (*securityhub.GetInvitationsCountOutput, error)
	GetInvitationsCountRequest(*securityhub.GetInvitationsCountInput) (*request.Request, *securityhub.GetInvitationsCountOutput)

	GetMasterAccount(*securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error)
	GetMasterAccountWithContext(aws.Context, *securityhub.GetMasterAccountInput, ...request.Option) (*securityhub.GetMasterAccountOutput, error)
	GetMasterAccountRequest(*securityhub.GetMasterAccountInput) (*request.Request, *securityhub.GetMasterAccountOutput)

	GetMembers(*securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error)
	GetMembersWithContext(aws.Context, *securityhub.GetMembersInput, ...request.Option) (*securityhub.GetMembersOutput, error)
	GetMembersRequest(*securityhub.GetMembersInput) (*request.Request, *securityhub.GetMembersOutput)

	InviteMembers(*securityhub.InviteMembersInput) (*securityhub.InviteMembersOutput, error)
	InviteMembersWithContext(aws.Context, *securityhub.InviteMembersInput, ...request.Option) (*securityhub.InviteMembersOutput, error)
	InviteMembersRequest(*securityhub.InviteMembersInput) (*request.Request, *securityhub.InviteMembersOutput)

	ListEnabledProductsForImport(*securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error)
	ListEnabledProductsForImportWithContext(aws.Context, *securityhub.ListEnabledProductsForImportInput, ...request.Option) (*securityhub.ListEnabledProductsForImportOutput, error)
	ListEnabledProductsForImportRequest(*securityhub.ListEnabledProductsForImportInput) (*request.Request, *securityhub.ListEnabledProductsForImportOutput)

	ListEnabledProductsForImportPages(*securityhub.ListEnabledProductsForImportInput, func(*securityhub.ListEnabledProductsForImportOutput, bool) bool) error
	ListEnabledProductsForImportPagesWithContext(aws.Context, *securityhub.ListEnabledProductsForImportInput, func(*securityhub.ListEnabledProductsForImportOutput, bool) bool, ...request.Option) error

	ListFindingAggregators(*securityhub.ListFindingAggregatorsInput) (*securityhub.ListFindingAggregatorsOutput, error)
	ListFindingAggregatorsWithContext(aws.Context, *securityhub.ListFindingAggregatorsInput, ...request.Option) (*securityhub.ListFindingAggregatorsOutput, error)
	ListFindingAggregatorsRequest(*securityhub.ListFindingAggregatorsInput) (*request.Request, *securityhub.ListFindingAggregatorsOutput)

	ListFindingAggregatorsPages(*securityhub.ListFindingAggregatorsInput, func(*securityhub.ListFindingAggregatorsOutput, bool) bool) error
	ListFindingAggregatorsPagesWithContext(aws.Context, *securityhub.ListFindingAggregatorsInput, func(*securityhub.ListFindingAggregatorsOutput, bool) bool, ...request.Option) error

	ListInvitations(*securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error)
	ListInvitationsWithContext(aws.Context, *securityhub.ListInvitationsInput, ...request.Option) (*securityhub.ListInvitationsOutput, error)
	ListInvitationsRequest(*securityhub.ListInvitationsInput) (*request.Request, *securityhub.ListInvitationsOutput)

	ListInvitationsPages(*securityhub.ListInvitationsInput, func(*securityhub.ListInvitationsOutput, bool) bool) error
	ListInvitationsPagesWithContext(aws.Context, *securityhub.ListInvitationsInput, func(*securityhub.ListInvitationsOutput, bool) bool, ...request.Option) error

	ListMembers(*securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error)
	ListMembersWithContext(aws.Context, *securityhub.ListMembersInput, ...request.Option) (*securityhub.ListMembersOutput, error)
	ListMembersRequest(*securityhub.ListMembersInput) (*request.Request, *securityhub.ListMembersOutput)

	ListMembersPages(*securityhub.ListMembersInput, func(*securityhub.ListMembersOutput, bool) bool) error
	ListMembersPagesWithContext(aws.Context, *securityhub.ListMembersInput, func(*securityhub.ListMembersOutput, bool) bool, ...request.Option) error

	ListOrganizationAdminAccounts(*securityhub.ListOrganizationAdminAccountsInput) (*securityhub.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsWithContext(aws.Context, *securityhub.ListOrganizationAdminAccountsInput, ...request.Option) (*securityhub.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsRequest(*securityhub.ListOrganizationAdminAccountsInput) (*request.Request, *securityhub.ListOrganizationAdminAccountsOutput)

	ListOrganizationAdminAccountsPages(*securityhub.ListOrganizationAdminAccountsInput, func(*securityhub.ListOrganizationAdminAccountsOutput, bool) bool) error
	ListOrganizationAdminAccountsPagesWithContext(aws.Context, *securityhub.ListOrganizationAdminAccountsInput, func(*securityhub.ListOrganizationAdminAccountsOutput, bool) bool, ...request.Option) error

	ListSecurityControlDefinitions(*securityhub.ListSecurityControlDefinitionsInput) (*securityhub.ListSecurityControlDefinitionsOutput, error)
	ListSecurityControlDefinitionsWithContext(aws.Context, *securityhub.ListSecurityControlDefinitionsInput, ...request.Option) (*securityhub.ListSecurityControlDefinitionsOutput, error)
	ListSecurityControlDefinitionsRequest(*securityhub.ListSecurityControlDefinitionsInput) (*request.Request, *securityhub.ListSecurityControlDefinitionsOutput)

	ListSecurityControlDefinitionsPages(*securityhub.ListSecurityControlDefinitionsInput, func(*securityhub.ListSecurityControlDefinitionsOutput, bool) bool) error
	ListSecurityControlDefinitionsPagesWithContext(aws.Context, *securityhub.ListSecurityControlDefinitionsInput, func(*securityhub.ListSecurityControlDefinitionsOutput, bool) bool, ...request.Option) error

	ListStandardsControlAssociations(*securityhub.ListStandardsControlAssociationsInput) (*securityhub.ListStandardsControlAssociationsOutput, error)
	ListStandardsControlAssociationsWithContext(aws.Context, *securityhub.ListStandardsControlAssociationsInput, ...request.Option) (*securityhub.ListStandardsControlAssociationsOutput, error)
	ListStandardsControlAssociationsRequest(*securityhub.ListStandardsControlAssociationsInput) (*request.Request, *securityhub.ListStandardsControlAssociationsOutput)

	ListStandardsControlAssociationsPages(*securityhub.ListStandardsControlAssociationsInput, func(*securityhub.ListStandardsControlAssociationsOutput, bool) bool) error
	ListStandardsControlAssociationsPagesWithContext(aws.Context, *securityhub.ListStandardsControlAssociationsInput, func(*securityhub.ListStandardsControlAssociationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*securityhub.ListTagsForResourceInput) (*securityhub.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *securityhub.ListTagsForResourceInput, ...request.Option) (*securityhub.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*securityhub.ListTagsForResourceInput) (*request.Request, *securityhub.ListTagsForResourceOutput)

	TagResource(*securityhub.TagResourceInput) (*securityhub.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *securityhub.TagResourceInput, ...request.Option) (*securityhub.TagResourceOutput, error)
	TagResourceRequest(*securityhub.TagResourceInput) (*request.Request, *securityhub.TagResourceOutput)

	UntagResource(*securityhub.UntagResourceInput) (*securityhub.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *securityhub.UntagResourceInput, ...request.Option) (*securityhub.UntagResourceOutput, error)
	UntagResourceRequest(*securityhub.UntagResourceInput) (*request.Request, *securityhub.UntagResourceOutput)

	UpdateActionTarget(*securityhub.UpdateActionTargetInput) (*securityhub.UpdateActionTargetOutput, error)
	UpdateActionTargetWithContext(aws.Context, *securityhub.UpdateActionTargetInput, ...request.Option) (*securityhub.UpdateActionTargetOutput, error)
	UpdateActionTargetRequest(*securityhub.UpdateActionTargetInput) (*request.Request, *securityhub.UpdateActionTargetOutput)

	UpdateFindingAggregator(*securityhub.UpdateFindingAggregatorInput) (*securityhub.UpdateFindingAggregatorOutput, error)
	UpdateFindingAggregatorWithContext(aws.Context, *securityhub.UpdateFindingAggregatorInput, ...request.Option) (*securityhub.UpdateFindingAggregatorOutput, error)
	UpdateFindingAggregatorRequest(*securityhub.UpdateFindingAggregatorInput) (*request.Request, *securityhub.UpdateFindingAggregatorOutput)

	UpdateFindings(*securityhub.UpdateFindingsInput) (*securityhub.UpdateFindingsOutput, error)
	UpdateFindingsWithContext(aws.Context, *securityhub.UpdateFindingsInput, ...request.Option) (*securityhub.UpdateFindingsOutput, error)
	UpdateFindingsRequest(*securityhub.UpdateFindingsInput) (*request.Request, *securityhub.UpdateFindingsOutput)

	UpdateInsight(*securityhub.UpdateInsightInput) (*securityhub.UpdateInsightOutput, error)
	UpdateInsightWithContext(aws.Context, *securityhub.UpdateInsightInput, ...request.Option) (*securityhub.UpdateInsightOutput, error)
	UpdateInsightRequest(*securityhub.UpdateInsightInput) (*request.Request, *securityhub.UpdateInsightOutput)

	UpdateOrganizationConfiguration(*securityhub.UpdateOrganizationConfigurationInput) (*securityhub.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationWithContext(aws.Context, *securityhub.UpdateOrganizationConfigurationInput, ...request.Option) (*securityhub.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationRequest(*securityhub.UpdateOrganizationConfigurationInput) (*request.Request, *securityhub.UpdateOrganizationConfigurationOutput)

	UpdateSecurityHubConfiguration(*securityhub.UpdateSecurityHubConfigurationInput) (*securityhub.UpdateSecurityHubConfigurationOutput, error)
	UpdateSecurityHubConfigurationWithContext(aws.Context, *securityhub.UpdateSecurityHubConfigurationInput, ...request.Option) (*securityhub.UpdateSecurityHubConfigurationOutput, error)
	UpdateSecurityHubConfigurationRequest(*securityhub.UpdateSecurityHubConfigurationInput) (*request.Request, *securityhub.UpdateSecurityHubConfigurationOutput)

	UpdateStandardsControl(*securityhub.UpdateStandardsControlInput) (*securityhub.UpdateStandardsControlOutput, error)
	UpdateStandardsControlWithContext(aws.Context, *securityhub.UpdateStandardsControlInput, ...request.Option) (*securityhub.UpdateStandardsControlOutput, error)
	UpdateStandardsControlRequest(*securityhub.UpdateStandardsControlInput) (*request.Request, *securityhub.UpdateStandardsControlOutput)
}

var _ SecurityHubAPI = (*securityhub.SecurityHub)(nil)
