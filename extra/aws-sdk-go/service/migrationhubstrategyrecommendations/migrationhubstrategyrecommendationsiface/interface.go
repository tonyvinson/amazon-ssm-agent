// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package migrationhubstrategyrecommendationsiface provides an interface to enable mocking the Migration Hub Strategy Recommendations service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package migrationhubstrategyrecommendationsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhubstrategyrecommendations"
)

// MigrationHubStrategyRecommendationsAPI provides an interface to enable mocking the
// migrationhubstrategyrecommendations.MigrationHubStrategyRecommendations service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Migration Hub Strategy Recommendations.
//    func myFunc(svc migrationhubstrategyrecommendationsiface.MigrationHubStrategyRecommendationsAPI) bool {
//        // Make svc.GetApplicationComponentDetails request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := migrationhubstrategyrecommendations.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMigrationHubStrategyRecommendationsClient struct {
//        migrationhubstrategyrecommendationsiface.MigrationHubStrategyRecommendationsAPI
//    }
//    func (m *mockMigrationHubStrategyRecommendationsClient) GetApplicationComponentDetails(input *migrationhubstrategyrecommendations.GetApplicationComponentDetailsInput) (*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMigrationHubStrategyRecommendationsClient{}
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
type MigrationHubStrategyRecommendationsAPI interface {
	GetApplicationComponentDetails(*migrationhubstrategyrecommendations.GetApplicationComponentDetailsInput) (*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput, error)
	GetApplicationComponentDetailsWithContext(aws.Context, *migrationhubstrategyrecommendations.GetApplicationComponentDetailsInput, ...request.Option) (*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput, error)
	GetApplicationComponentDetailsRequest(*migrationhubstrategyrecommendations.GetApplicationComponentDetailsInput) (*request.Request, *migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput)

	GetApplicationComponentStrategies(*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesInput) (*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput, error)
	GetApplicationComponentStrategiesWithContext(aws.Context, *migrationhubstrategyrecommendations.GetApplicationComponentStrategiesInput, ...request.Option) (*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput, error)
	GetApplicationComponentStrategiesRequest(*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesInput) (*request.Request, *migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput)

	GetAssessment(*migrationhubstrategyrecommendations.GetAssessmentInput) (*migrationhubstrategyrecommendations.GetAssessmentOutput, error)
	GetAssessmentWithContext(aws.Context, *migrationhubstrategyrecommendations.GetAssessmentInput, ...request.Option) (*migrationhubstrategyrecommendations.GetAssessmentOutput, error)
	GetAssessmentRequest(*migrationhubstrategyrecommendations.GetAssessmentInput) (*request.Request, *migrationhubstrategyrecommendations.GetAssessmentOutput)

	GetImportFileTask(*migrationhubstrategyrecommendations.GetImportFileTaskInput) (*migrationhubstrategyrecommendations.GetImportFileTaskOutput, error)
	GetImportFileTaskWithContext(aws.Context, *migrationhubstrategyrecommendations.GetImportFileTaskInput, ...request.Option) (*migrationhubstrategyrecommendations.GetImportFileTaskOutput, error)
	GetImportFileTaskRequest(*migrationhubstrategyrecommendations.GetImportFileTaskInput) (*request.Request, *migrationhubstrategyrecommendations.GetImportFileTaskOutput)

	GetLatestAssessmentId(*migrationhubstrategyrecommendations.GetLatestAssessmentIdInput) (*migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput, error)
	GetLatestAssessmentIdWithContext(aws.Context, *migrationhubstrategyrecommendations.GetLatestAssessmentIdInput, ...request.Option) (*migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput, error)
	GetLatestAssessmentIdRequest(*migrationhubstrategyrecommendations.GetLatestAssessmentIdInput) (*request.Request, *migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput)

	GetPortfolioPreferences(*migrationhubstrategyrecommendations.GetPortfolioPreferencesInput) (*migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput, error)
	GetPortfolioPreferencesWithContext(aws.Context, *migrationhubstrategyrecommendations.GetPortfolioPreferencesInput, ...request.Option) (*migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput, error)
	GetPortfolioPreferencesRequest(*migrationhubstrategyrecommendations.GetPortfolioPreferencesInput) (*request.Request, *migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput)

	GetPortfolioSummary(*migrationhubstrategyrecommendations.GetPortfolioSummaryInput) (*migrationhubstrategyrecommendations.GetPortfolioSummaryOutput, error)
	GetPortfolioSummaryWithContext(aws.Context, *migrationhubstrategyrecommendations.GetPortfolioSummaryInput, ...request.Option) (*migrationhubstrategyrecommendations.GetPortfolioSummaryOutput, error)
	GetPortfolioSummaryRequest(*migrationhubstrategyrecommendations.GetPortfolioSummaryInput) (*request.Request, *migrationhubstrategyrecommendations.GetPortfolioSummaryOutput)

	GetRecommendationReportDetails(*migrationhubstrategyrecommendations.GetRecommendationReportDetailsInput) (*migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput, error)
	GetRecommendationReportDetailsWithContext(aws.Context, *migrationhubstrategyrecommendations.GetRecommendationReportDetailsInput, ...request.Option) (*migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput, error)
	GetRecommendationReportDetailsRequest(*migrationhubstrategyrecommendations.GetRecommendationReportDetailsInput) (*request.Request, *migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput)

	GetServerDetails(*migrationhubstrategyrecommendations.GetServerDetailsInput) (*migrationhubstrategyrecommendations.GetServerDetailsOutput, error)
	GetServerDetailsWithContext(aws.Context, *migrationhubstrategyrecommendations.GetServerDetailsInput, ...request.Option) (*migrationhubstrategyrecommendations.GetServerDetailsOutput, error)
	GetServerDetailsRequest(*migrationhubstrategyrecommendations.GetServerDetailsInput) (*request.Request, *migrationhubstrategyrecommendations.GetServerDetailsOutput)

	GetServerDetailsPages(*migrationhubstrategyrecommendations.GetServerDetailsInput, func(*migrationhubstrategyrecommendations.GetServerDetailsOutput, bool) bool) error
	GetServerDetailsPagesWithContext(aws.Context, *migrationhubstrategyrecommendations.GetServerDetailsInput, func(*migrationhubstrategyrecommendations.GetServerDetailsOutput, bool) bool, ...request.Option) error

	GetServerStrategies(*migrationhubstrategyrecommendations.GetServerStrategiesInput) (*migrationhubstrategyrecommendations.GetServerStrategiesOutput, error)
	GetServerStrategiesWithContext(aws.Context, *migrationhubstrategyrecommendations.GetServerStrategiesInput, ...request.Option) (*migrationhubstrategyrecommendations.GetServerStrategiesOutput, error)
	GetServerStrategiesRequest(*migrationhubstrategyrecommendations.GetServerStrategiesInput) (*request.Request, *migrationhubstrategyrecommendations.GetServerStrategiesOutput)

	ListApplicationComponents(*migrationhubstrategyrecommendations.ListApplicationComponentsInput) (*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, error)
	ListApplicationComponentsWithContext(aws.Context, *migrationhubstrategyrecommendations.ListApplicationComponentsInput, ...request.Option) (*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, error)
	ListApplicationComponentsRequest(*migrationhubstrategyrecommendations.ListApplicationComponentsInput) (*request.Request, *migrationhubstrategyrecommendations.ListApplicationComponentsOutput)

	ListApplicationComponentsPages(*migrationhubstrategyrecommendations.ListApplicationComponentsInput, func(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, bool) bool) error
	ListApplicationComponentsPagesWithContext(aws.Context, *migrationhubstrategyrecommendations.ListApplicationComponentsInput, func(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, bool) bool, ...request.Option) error

	ListCollectors(*migrationhubstrategyrecommendations.ListCollectorsInput) (*migrationhubstrategyrecommendations.ListCollectorsOutput, error)
	ListCollectorsWithContext(aws.Context, *migrationhubstrategyrecommendations.ListCollectorsInput, ...request.Option) (*migrationhubstrategyrecommendations.ListCollectorsOutput, error)
	ListCollectorsRequest(*migrationhubstrategyrecommendations.ListCollectorsInput) (*request.Request, *migrationhubstrategyrecommendations.ListCollectorsOutput)

	ListCollectorsPages(*migrationhubstrategyrecommendations.ListCollectorsInput, func(*migrationhubstrategyrecommendations.ListCollectorsOutput, bool) bool) error
	ListCollectorsPagesWithContext(aws.Context, *migrationhubstrategyrecommendations.ListCollectorsInput, func(*migrationhubstrategyrecommendations.ListCollectorsOutput, bool) bool, ...request.Option) error

	ListImportFileTask(*migrationhubstrategyrecommendations.ListImportFileTaskInput) (*migrationhubstrategyrecommendations.ListImportFileTaskOutput, error)
	ListImportFileTaskWithContext(aws.Context, *migrationhubstrategyrecommendations.ListImportFileTaskInput, ...request.Option) (*migrationhubstrategyrecommendations.ListImportFileTaskOutput, error)
	ListImportFileTaskRequest(*migrationhubstrategyrecommendations.ListImportFileTaskInput) (*request.Request, *migrationhubstrategyrecommendations.ListImportFileTaskOutput)

	ListImportFileTaskPages(*migrationhubstrategyrecommendations.ListImportFileTaskInput, func(*migrationhubstrategyrecommendations.ListImportFileTaskOutput, bool) bool) error
	ListImportFileTaskPagesWithContext(aws.Context, *migrationhubstrategyrecommendations.ListImportFileTaskInput, func(*migrationhubstrategyrecommendations.ListImportFileTaskOutput, bool) bool, ...request.Option) error

	ListServers(*migrationhubstrategyrecommendations.ListServersInput) (*migrationhubstrategyrecommendations.ListServersOutput, error)
	ListServersWithContext(aws.Context, *migrationhubstrategyrecommendations.ListServersInput, ...request.Option) (*migrationhubstrategyrecommendations.ListServersOutput, error)
	ListServersRequest(*migrationhubstrategyrecommendations.ListServersInput) (*request.Request, *migrationhubstrategyrecommendations.ListServersOutput)

	ListServersPages(*migrationhubstrategyrecommendations.ListServersInput, func(*migrationhubstrategyrecommendations.ListServersOutput, bool) bool) error
	ListServersPagesWithContext(aws.Context, *migrationhubstrategyrecommendations.ListServersInput, func(*migrationhubstrategyrecommendations.ListServersOutput, bool) bool, ...request.Option) error

	PutPortfolioPreferences(*migrationhubstrategyrecommendations.PutPortfolioPreferencesInput) (*migrationhubstrategyrecommendations.PutPortfolioPreferencesOutput, error)
	PutPortfolioPreferencesWithContext(aws.Context, *migrationhubstrategyrecommendations.PutPortfolioPreferencesInput, ...request.Option) (*migrationhubstrategyrecommendations.PutPortfolioPreferencesOutput, error)
	PutPortfolioPreferencesRequest(*migrationhubstrategyrecommendations.PutPortfolioPreferencesInput) (*request.Request, *migrationhubstrategyrecommendations.PutPortfolioPreferencesOutput)

	StartAssessment(*migrationhubstrategyrecommendations.StartAssessmentInput) (*migrationhubstrategyrecommendations.StartAssessmentOutput, error)
	StartAssessmentWithContext(aws.Context, *migrationhubstrategyrecommendations.StartAssessmentInput, ...request.Option) (*migrationhubstrategyrecommendations.StartAssessmentOutput, error)
	StartAssessmentRequest(*migrationhubstrategyrecommendations.StartAssessmentInput) (*request.Request, *migrationhubstrategyrecommendations.StartAssessmentOutput)

	StartImportFileTask(*migrationhubstrategyrecommendations.StartImportFileTaskInput) (*migrationhubstrategyrecommendations.StartImportFileTaskOutput, error)
	StartImportFileTaskWithContext(aws.Context, *migrationhubstrategyrecommendations.StartImportFileTaskInput, ...request.Option) (*migrationhubstrategyrecommendations.StartImportFileTaskOutput, error)
	StartImportFileTaskRequest(*migrationhubstrategyrecommendations.StartImportFileTaskInput) (*request.Request, *migrationhubstrategyrecommendations.StartImportFileTaskOutput)

	StartRecommendationReportGeneration(*migrationhubstrategyrecommendations.StartRecommendationReportGenerationInput) (*migrationhubstrategyrecommendations.StartRecommendationReportGenerationOutput, error)
	StartRecommendationReportGenerationWithContext(aws.Context, *migrationhubstrategyrecommendations.StartRecommendationReportGenerationInput, ...request.Option) (*migrationhubstrategyrecommendations.StartRecommendationReportGenerationOutput, error)
	StartRecommendationReportGenerationRequest(*migrationhubstrategyrecommendations.StartRecommendationReportGenerationInput) (*request.Request, *migrationhubstrategyrecommendations.StartRecommendationReportGenerationOutput)

	StopAssessment(*migrationhubstrategyrecommendations.StopAssessmentInput) (*migrationhubstrategyrecommendations.StopAssessmentOutput, error)
	StopAssessmentWithContext(aws.Context, *migrationhubstrategyrecommendations.StopAssessmentInput, ...request.Option) (*migrationhubstrategyrecommendations.StopAssessmentOutput, error)
	StopAssessmentRequest(*migrationhubstrategyrecommendations.StopAssessmentInput) (*request.Request, *migrationhubstrategyrecommendations.StopAssessmentOutput)

	UpdateApplicationComponentConfig(*migrationhubstrategyrecommendations.UpdateApplicationComponentConfigInput) (*migrationhubstrategyrecommendations.UpdateApplicationComponentConfigOutput, error)
	UpdateApplicationComponentConfigWithContext(aws.Context, *migrationhubstrategyrecommendations.UpdateApplicationComponentConfigInput, ...request.Option) (*migrationhubstrategyrecommendations.UpdateApplicationComponentConfigOutput, error)
	UpdateApplicationComponentConfigRequest(*migrationhubstrategyrecommendations.UpdateApplicationComponentConfigInput) (*request.Request, *migrationhubstrategyrecommendations.UpdateApplicationComponentConfigOutput)

	UpdateServerConfig(*migrationhubstrategyrecommendations.UpdateServerConfigInput) (*migrationhubstrategyrecommendations.UpdateServerConfigOutput, error)
	UpdateServerConfigWithContext(aws.Context, *migrationhubstrategyrecommendations.UpdateServerConfigInput, ...request.Option) (*migrationhubstrategyrecommendations.UpdateServerConfigOutput, error)
	UpdateServerConfigRequest(*migrationhubstrategyrecommendations.UpdateServerConfigInput) (*request.Request, *migrationhubstrategyrecommendations.UpdateServerConfigOutput)
}

var _ MigrationHubStrategyRecommendationsAPI = (*migrationhubstrategyrecommendations.MigrationHubStrategyRecommendations)(nil)
