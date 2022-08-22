// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package personalizeiface provides an interface to enable mocking the Amazon Personalize service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package personalizeiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/personalize"
)

// PersonalizeAPI provides an interface to enable mocking the
// personalize.Personalize service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Personalize.
//    func myFunc(svc personalizeiface.PersonalizeAPI) bool {
//        // Make svc.CreateBatchInferenceJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := personalize.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPersonalizeClient struct {
//        personalizeiface.PersonalizeAPI
//    }
//    func (m *mockPersonalizeClient) CreateBatchInferenceJob(input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPersonalizeClient{}
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
type PersonalizeAPI interface {
	CreateBatchInferenceJob(*personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error)
	CreateBatchInferenceJobWithContext(aws.Context, *personalize.CreateBatchInferenceJobInput, ...request.Option) (*personalize.CreateBatchInferenceJobOutput, error)
	CreateBatchInferenceJobRequest(*personalize.CreateBatchInferenceJobInput) (*request.Request, *personalize.CreateBatchInferenceJobOutput)

	CreateBatchSegmentJob(*personalize.CreateBatchSegmentJobInput) (*personalize.CreateBatchSegmentJobOutput, error)
	CreateBatchSegmentJobWithContext(aws.Context, *personalize.CreateBatchSegmentJobInput, ...request.Option) (*personalize.CreateBatchSegmentJobOutput, error)
	CreateBatchSegmentJobRequest(*personalize.CreateBatchSegmentJobInput) (*request.Request, *personalize.CreateBatchSegmentJobOutput)

	CreateCampaign(*personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error)
	CreateCampaignWithContext(aws.Context, *personalize.CreateCampaignInput, ...request.Option) (*personalize.CreateCampaignOutput, error)
	CreateCampaignRequest(*personalize.CreateCampaignInput) (*request.Request, *personalize.CreateCampaignOutput)

	CreateDataset(*personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error)
	CreateDatasetWithContext(aws.Context, *personalize.CreateDatasetInput, ...request.Option) (*personalize.CreateDatasetOutput, error)
	CreateDatasetRequest(*personalize.CreateDatasetInput) (*request.Request, *personalize.CreateDatasetOutput)

	CreateDatasetExportJob(*personalize.CreateDatasetExportJobInput) (*personalize.CreateDatasetExportJobOutput, error)
	CreateDatasetExportJobWithContext(aws.Context, *personalize.CreateDatasetExportJobInput, ...request.Option) (*personalize.CreateDatasetExportJobOutput, error)
	CreateDatasetExportJobRequest(*personalize.CreateDatasetExportJobInput) (*request.Request, *personalize.CreateDatasetExportJobOutput)

	CreateDatasetGroup(*personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error)
	CreateDatasetGroupWithContext(aws.Context, *personalize.CreateDatasetGroupInput, ...request.Option) (*personalize.CreateDatasetGroupOutput, error)
	CreateDatasetGroupRequest(*personalize.CreateDatasetGroupInput) (*request.Request, *personalize.CreateDatasetGroupOutput)

	CreateDatasetImportJob(*personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobWithContext(aws.Context, *personalize.CreateDatasetImportJobInput, ...request.Option) (*personalize.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobRequest(*personalize.CreateDatasetImportJobInput) (*request.Request, *personalize.CreateDatasetImportJobOutput)

	CreateEventTracker(*personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error)
	CreateEventTrackerWithContext(aws.Context, *personalize.CreateEventTrackerInput, ...request.Option) (*personalize.CreateEventTrackerOutput, error)
	CreateEventTrackerRequest(*personalize.CreateEventTrackerInput) (*request.Request, *personalize.CreateEventTrackerOutput)

	CreateFilter(*personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error)
	CreateFilterWithContext(aws.Context, *personalize.CreateFilterInput, ...request.Option) (*personalize.CreateFilterOutput, error)
	CreateFilterRequest(*personalize.CreateFilterInput) (*request.Request, *personalize.CreateFilterOutput)

	CreateRecommender(*personalize.CreateRecommenderInput) (*personalize.CreateRecommenderOutput, error)
	CreateRecommenderWithContext(aws.Context, *personalize.CreateRecommenderInput, ...request.Option) (*personalize.CreateRecommenderOutput, error)
	CreateRecommenderRequest(*personalize.CreateRecommenderInput) (*request.Request, *personalize.CreateRecommenderOutput)

	CreateSchema(*personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error)
	CreateSchemaWithContext(aws.Context, *personalize.CreateSchemaInput, ...request.Option) (*personalize.CreateSchemaOutput, error)
	CreateSchemaRequest(*personalize.CreateSchemaInput) (*request.Request, *personalize.CreateSchemaOutput)

	CreateSolution(*personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error)
	CreateSolutionWithContext(aws.Context, *personalize.CreateSolutionInput, ...request.Option) (*personalize.CreateSolutionOutput, error)
	CreateSolutionRequest(*personalize.CreateSolutionInput) (*request.Request, *personalize.CreateSolutionOutput)

	CreateSolutionVersion(*personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error)
	CreateSolutionVersionWithContext(aws.Context, *personalize.CreateSolutionVersionInput, ...request.Option) (*personalize.CreateSolutionVersionOutput, error)
	CreateSolutionVersionRequest(*personalize.CreateSolutionVersionInput) (*request.Request, *personalize.CreateSolutionVersionOutput)

	DeleteCampaign(*personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error)
	DeleteCampaignWithContext(aws.Context, *personalize.DeleteCampaignInput, ...request.Option) (*personalize.DeleteCampaignOutput, error)
	DeleteCampaignRequest(*personalize.DeleteCampaignInput) (*request.Request, *personalize.DeleteCampaignOutput)

	DeleteDataset(*personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *personalize.DeleteDatasetInput, ...request.Option) (*personalize.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*personalize.DeleteDatasetInput) (*request.Request, *personalize.DeleteDatasetOutput)

	DeleteDatasetGroup(*personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupWithContext(aws.Context, *personalize.DeleteDatasetGroupInput, ...request.Option) (*personalize.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupRequest(*personalize.DeleteDatasetGroupInput) (*request.Request, *personalize.DeleteDatasetGroupOutput)

	DeleteEventTracker(*personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error)
	DeleteEventTrackerWithContext(aws.Context, *personalize.DeleteEventTrackerInput, ...request.Option) (*personalize.DeleteEventTrackerOutput, error)
	DeleteEventTrackerRequest(*personalize.DeleteEventTrackerInput) (*request.Request, *personalize.DeleteEventTrackerOutput)

	DeleteFilter(*personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error)
	DeleteFilterWithContext(aws.Context, *personalize.DeleteFilterInput, ...request.Option) (*personalize.DeleteFilterOutput, error)
	DeleteFilterRequest(*personalize.DeleteFilterInput) (*request.Request, *personalize.DeleteFilterOutput)

	DeleteRecommender(*personalize.DeleteRecommenderInput) (*personalize.DeleteRecommenderOutput, error)
	DeleteRecommenderWithContext(aws.Context, *personalize.DeleteRecommenderInput, ...request.Option) (*personalize.DeleteRecommenderOutput, error)
	DeleteRecommenderRequest(*personalize.DeleteRecommenderInput) (*request.Request, *personalize.DeleteRecommenderOutput)

	DeleteSchema(*personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error)
	DeleteSchemaWithContext(aws.Context, *personalize.DeleteSchemaInput, ...request.Option) (*personalize.DeleteSchemaOutput, error)
	DeleteSchemaRequest(*personalize.DeleteSchemaInput) (*request.Request, *personalize.DeleteSchemaOutput)

	DeleteSolution(*personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error)
	DeleteSolutionWithContext(aws.Context, *personalize.DeleteSolutionInput, ...request.Option) (*personalize.DeleteSolutionOutput, error)
	DeleteSolutionRequest(*personalize.DeleteSolutionInput) (*request.Request, *personalize.DeleteSolutionOutput)

	DescribeAlgorithm(*personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error)
	DescribeAlgorithmWithContext(aws.Context, *personalize.DescribeAlgorithmInput, ...request.Option) (*personalize.DescribeAlgorithmOutput, error)
	DescribeAlgorithmRequest(*personalize.DescribeAlgorithmInput) (*request.Request, *personalize.DescribeAlgorithmOutput)

	DescribeBatchInferenceJob(*personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error)
	DescribeBatchInferenceJobWithContext(aws.Context, *personalize.DescribeBatchInferenceJobInput, ...request.Option) (*personalize.DescribeBatchInferenceJobOutput, error)
	DescribeBatchInferenceJobRequest(*personalize.DescribeBatchInferenceJobInput) (*request.Request, *personalize.DescribeBatchInferenceJobOutput)

	DescribeBatchSegmentJob(*personalize.DescribeBatchSegmentJobInput) (*personalize.DescribeBatchSegmentJobOutput, error)
	DescribeBatchSegmentJobWithContext(aws.Context, *personalize.DescribeBatchSegmentJobInput, ...request.Option) (*personalize.DescribeBatchSegmentJobOutput, error)
	DescribeBatchSegmentJobRequest(*personalize.DescribeBatchSegmentJobInput) (*request.Request, *personalize.DescribeBatchSegmentJobOutput)

	DescribeCampaign(*personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error)
	DescribeCampaignWithContext(aws.Context, *personalize.DescribeCampaignInput, ...request.Option) (*personalize.DescribeCampaignOutput, error)
	DescribeCampaignRequest(*personalize.DescribeCampaignInput) (*request.Request, *personalize.DescribeCampaignOutput)

	DescribeDataset(*personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error)
	DescribeDatasetWithContext(aws.Context, *personalize.DescribeDatasetInput, ...request.Option) (*personalize.DescribeDatasetOutput, error)
	DescribeDatasetRequest(*personalize.DescribeDatasetInput) (*request.Request, *personalize.DescribeDatasetOutput)

	DescribeDatasetExportJob(*personalize.DescribeDatasetExportJobInput) (*personalize.DescribeDatasetExportJobOutput, error)
	DescribeDatasetExportJobWithContext(aws.Context, *personalize.DescribeDatasetExportJobInput, ...request.Option) (*personalize.DescribeDatasetExportJobOutput, error)
	DescribeDatasetExportJobRequest(*personalize.DescribeDatasetExportJobInput) (*request.Request, *personalize.DescribeDatasetExportJobOutput)

	DescribeDatasetGroup(*personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupWithContext(aws.Context, *personalize.DescribeDatasetGroupInput, ...request.Option) (*personalize.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupRequest(*personalize.DescribeDatasetGroupInput) (*request.Request, *personalize.DescribeDatasetGroupOutput)

	DescribeDatasetImportJob(*personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobWithContext(aws.Context, *personalize.DescribeDatasetImportJobInput, ...request.Option) (*personalize.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobRequest(*personalize.DescribeDatasetImportJobInput) (*request.Request, *personalize.DescribeDatasetImportJobOutput)

	DescribeEventTracker(*personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error)
	DescribeEventTrackerWithContext(aws.Context, *personalize.DescribeEventTrackerInput, ...request.Option) (*personalize.DescribeEventTrackerOutput, error)
	DescribeEventTrackerRequest(*personalize.DescribeEventTrackerInput) (*request.Request, *personalize.DescribeEventTrackerOutput)

	DescribeFeatureTransformation(*personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error)
	DescribeFeatureTransformationWithContext(aws.Context, *personalize.DescribeFeatureTransformationInput, ...request.Option) (*personalize.DescribeFeatureTransformationOutput, error)
	DescribeFeatureTransformationRequest(*personalize.DescribeFeatureTransformationInput) (*request.Request, *personalize.DescribeFeatureTransformationOutput)

	DescribeFilter(*personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error)
	DescribeFilterWithContext(aws.Context, *personalize.DescribeFilterInput, ...request.Option) (*personalize.DescribeFilterOutput, error)
	DescribeFilterRequest(*personalize.DescribeFilterInput) (*request.Request, *personalize.DescribeFilterOutput)

	DescribeRecipe(*personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error)
	DescribeRecipeWithContext(aws.Context, *personalize.DescribeRecipeInput, ...request.Option) (*personalize.DescribeRecipeOutput, error)
	DescribeRecipeRequest(*personalize.DescribeRecipeInput) (*request.Request, *personalize.DescribeRecipeOutput)

	DescribeRecommender(*personalize.DescribeRecommenderInput) (*personalize.DescribeRecommenderOutput, error)
	DescribeRecommenderWithContext(aws.Context, *personalize.DescribeRecommenderInput, ...request.Option) (*personalize.DescribeRecommenderOutput, error)
	DescribeRecommenderRequest(*personalize.DescribeRecommenderInput) (*request.Request, *personalize.DescribeRecommenderOutput)

	DescribeSchema(*personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error)
	DescribeSchemaWithContext(aws.Context, *personalize.DescribeSchemaInput, ...request.Option) (*personalize.DescribeSchemaOutput, error)
	DescribeSchemaRequest(*personalize.DescribeSchemaInput) (*request.Request, *personalize.DescribeSchemaOutput)

	DescribeSolution(*personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error)
	DescribeSolutionWithContext(aws.Context, *personalize.DescribeSolutionInput, ...request.Option) (*personalize.DescribeSolutionOutput, error)
	DescribeSolutionRequest(*personalize.DescribeSolutionInput) (*request.Request, *personalize.DescribeSolutionOutput)

	DescribeSolutionVersion(*personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error)
	DescribeSolutionVersionWithContext(aws.Context, *personalize.DescribeSolutionVersionInput, ...request.Option) (*personalize.DescribeSolutionVersionOutput, error)
	DescribeSolutionVersionRequest(*personalize.DescribeSolutionVersionInput) (*request.Request, *personalize.DescribeSolutionVersionOutput)

	GetSolutionMetrics(*personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error)
	GetSolutionMetricsWithContext(aws.Context, *personalize.GetSolutionMetricsInput, ...request.Option) (*personalize.GetSolutionMetricsOutput, error)
	GetSolutionMetricsRequest(*personalize.GetSolutionMetricsInput) (*request.Request, *personalize.GetSolutionMetricsOutput)

	ListBatchInferenceJobs(*personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error)
	ListBatchInferenceJobsWithContext(aws.Context, *personalize.ListBatchInferenceJobsInput, ...request.Option) (*personalize.ListBatchInferenceJobsOutput, error)
	ListBatchInferenceJobsRequest(*personalize.ListBatchInferenceJobsInput) (*request.Request, *personalize.ListBatchInferenceJobsOutput)

	ListBatchInferenceJobsPages(*personalize.ListBatchInferenceJobsInput, func(*personalize.ListBatchInferenceJobsOutput, bool) bool) error
	ListBatchInferenceJobsPagesWithContext(aws.Context, *personalize.ListBatchInferenceJobsInput, func(*personalize.ListBatchInferenceJobsOutput, bool) bool, ...request.Option) error

	ListBatchSegmentJobs(*personalize.ListBatchSegmentJobsInput) (*personalize.ListBatchSegmentJobsOutput, error)
	ListBatchSegmentJobsWithContext(aws.Context, *personalize.ListBatchSegmentJobsInput, ...request.Option) (*personalize.ListBatchSegmentJobsOutput, error)
	ListBatchSegmentJobsRequest(*personalize.ListBatchSegmentJobsInput) (*request.Request, *personalize.ListBatchSegmentJobsOutput)

	ListBatchSegmentJobsPages(*personalize.ListBatchSegmentJobsInput, func(*personalize.ListBatchSegmentJobsOutput, bool) bool) error
	ListBatchSegmentJobsPagesWithContext(aws.Context, *personalize.ListBatchSegmentJobsInput, func(*personalize.ListBatchSegmentJobsOutput, bool) bool, ...request.Option) error

	ListCampaigns(*personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error)
	ListCampaignsWithContext(aws.Context, *personalize.ListCampaignsInput, ...request.Option) (*personalize.ListCampaignsOutput, error)
	ListCampaignsRequest(*personalize.ListCampaignsInput) (*request.Request, *personalize.ListCampaignsOutput)

	ListCampaignsPages(*personalize.ListCampaignsInput, func(*personalize.ListCampaignsOutput, bool) bool) error
	ListCampaignsPagesWithContext(aws.Context, *personalize.ListCampaignsInput, func(*personalize.ListCampaignsOutput, bool) bool, ...request.Option) error

	ListDatasetExportJobs(*personalize.ListDatasetExportJobsInput) (*personalize.ListDatasetExportJobsOutput, error)
	ListDatasetExportJobsWithContext(aws.Context, *personalize.ListDatasetExportJobsInput, ...request.Option) (*personalize.ListDatasetExportJobsOutput, error)
	ListDatasetExportJobsRequest(*personalize.ListDatasetExportJobsInput) (*request.Request, *personalize.ListDatasetExportJobsOutput)

	ListDatasetExportJobsPages(*personalize.ListDatasetExportJobsInput, func(*personalize.ListDatasetExportJobsOutput, bool) bool) error
	ListDatasetExportJobsPagesWithContext(aws.Context, *personalize.ListDatasetExportJobsInput, func(*personalize.ListDatasetExportJobsOutput, bool) bool, ...request.Option) error

	ListDatasetGroups(*personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error)
	ListDatasetGroupsWithContext(aws.Context, *personalize.ListDatasetGroupsInput, ...request.Option) (*personalize.ListDatasetGroupsOutput, error)
	ListDatasetGroupsRequest(*personalize.ListDatasetGroupsInput) (*request.Request, *personalize.ListDatasetGroupsOutput)

	ListDatasetGroupsPages(*personalize.ListDatasetGroupsInput, func(*personalize.ListDatasetGroupsOutput, bool) bool) error
	ListDatasetGroupsPagesWithContext(aws.Context, *personalize.ListDatasetGroupsInput, func(*personalize.ListDatasetGroupsOutput, bool) bool, ...request.Option) error

	ListDatasetImportJobs(*personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsWithContext(aws.Context, *personalize.ListDatasetImportJobsInput, ...request.Option) (*personalize.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsRequest(*personalize.ListDatasetImportJobsInput) (*request.Request, *personalize.ListDatasetImportJobsOutput)

	ListDatasetImportJobsPages(*personalize.ListDatasetImportJobsInput, func(*personalize.ListDatasetImportJobsOutput, bool) bool) error
	ListDatasetImportJobsPagesWithContext(aws.Context, *personalize.ListDatasetImportJobsInput, func(*personalize.ListDatasetImportJobsOutput, bool) bool, ...request.Option) error

	ListDatasets(*personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *personalize.ListDatasetsInput, ...request.Option) (*personalize.ListDatasetsOutput, error)
	ListDatasetsRequest(*personalize.ListDatasetsInput) (*request.Request, *personalize.ListDatasetsOutput)

	ListDatasetsPages(*personalize.ListDatasetsInput, func(*personalize.ListDatasetsOutput, bool) bool) error
	ListDatasetsPagesWithContext(aws.Context, *personalize.ListDatasetsInput, func(*personalize.ListDatasetsOutput, bool) bool, ...request.Option) error

	ListEventTrackers(*personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error)
	ListEventTrackersWithContext(aws.Context, *personalize.ListEventTrackersInput, ...request.Option) (*personalize.ListEventTrackersOutput, error)
	ListEventTrackersRequest(*personalize.ListEventTrackersInput) (*request.Request, *personalize.ListEventTrackersOutput)

	ListEventTrackersPages(*personalize.ListEventTrackersInput, func(*personalize.ListEventTrackersOutput, bool) bool) error
	ListEventTrackersPagesWithContext(aws.Context, *personalize.ListEventTrackersInput, func(*personalize.ListEventTrackersOutput, bool) bool, ...request.Option) error

	ListFilters(*personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error)
	ListFiltersWithContext(aws.Context, *personalize.ListFiltersInput, ...request.Option) (*personalize.ListFiltersOutput, error)
	ListFiltersRequest(*personalize.ListFiltersInput) (*request.Request, *personalize.ListFiltersOutput)

	ListFiltersPages(*personalize.ListFiltersInput, func(*personalize.ListFiltersOutput, bool) bool) error
	ListFiltersPagesWithContext(aws.Context, *personalize.ListFiltersInput, func(*personalize.ListFiltersOutput, bool) bool, ...request.Option) error

	ListRecipes(*personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error)
	ListRecipesWithContext(aws.Context, *personalize.ListRecipesInput, ...request.Option) (*personalize.ListRecipesOutput, error)
	ListRecipesRequest(*personalize.ListRecipesInput) (*request.Request, *personalize.ListRecipesOutput)

	ListRecipesPages(*personalize.ListRecipesInput, func(*personalize.ListRecipesOutput, bool) bool) error
	ListRecipesPagesWithContext(aws.Context, *personalize.ListRecipesInput, func(*personalize.ListRecipesOutput, bool) bool, ...request.Option) error

	ListRecommenders(*personalize.ListRecommendersInput) (*personalize.ListRecommendersOutput, error)
	ListRecommendersWithContext(aws.Context, *personalize.ListRecommendersInput, ...request.Option) (*personalize.ListRecommendersOutput, error)
	ListRecommendersRequest(*personalize.ListRecommendersInput) (*request.Request, *personalize.ListRecommendersOutput)

	ListRecommendersPages(*personalize.ListRecommendersInput, func(*personalize.ListRecommendersOutput, bool) bool) error
	ListRecommendersPagesWithContext(aws.Context, *personalize.ListRecommendersInput, func(*personalize.ListRecommendersOutput, bool) bool, ...request.Option) error

	ListSchemas(*personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error)
	ListSchemasWithContext(aws.Context, *personalize.ListSchemasInput, ...request.Option) (*personalize.ListSchemasOutput, error)
	ListSchemasRequest(*personalize.ListSchemasInput) (*request.Request, *personalize.ListSchemasOutput)

	ListSchemasPages(*personalize.ListSchemasInput, func(*personalize.ListSchemasOutput, bool) bool) error
	ListSchemasPagesWithContext(aws.Context, *personalize.ListSchemasInput, func(*personalize.ListSchemasOutput, bool) bool, ...request.Option) error

	ListSolutionVersions(*personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error)
	ListSolutionVersionsWithContext(aws.Context, *personalize.ListSolutionVersionsInput, ...request.Option) (*personalize.ListSolutionVersionsOutput, error)
	ListSolutionVersionsRequest(*personalize.ListSolutionVersionsInput) (*request.Request, *personalize.ListSolutionVersionsOutput)

	ListSolutionVersionsPages(*personalize.ListSolutionVersionsInput, func(*personalize.ListSolutionVersionsOutput, bool) bool) error
	ListSolutionVersionsPagesWithContext(aws.Context, *personalize.ListSolutionVersionsInput, func(*personalize.ListSolutionVersionsOutput, bool) bool, ...request.Option) error

	ListSolutions(*personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error)
	ListSolutionsWithContext(aws.Context, *personalize.ListSolutionsInput, ...request.Option) (*personalize.ListSolutionsOutput, error)
	ListSolutionsRequest(*personalize.ListSolutionsInput) (*request.Request, *personalize.ListSolutionsOutput)

	ListSolutionsPages(*personalize.ListSolutionsInput, func(*personalize.ListSolutionsOutput, bool) bool) error
	ListSolutionsPagesWithContext(aws.Context, *personalize.ListSolutionsInput, func(*personalize.ListSolutionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*personalize.ListTagsForResourceInput) (*personalize.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *personalize.ListTagsForResourceInput, ...request.Option) (*personalize.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*personalize.ListTagsForResourceInput) (*request.Request, *personalize.ListTagsForResourceOutput)

	StartRecommender(*personalize.StartRecommenderInput) (*personalize.StartRecommenderOutput, error)
	StartRecommenderWithContext(aws.Context, *personalize.StartRecommenderInput, ...request.Option) (*personalize.StartRecommenderOutput, error)
	StartRecommenderRequest(*personalize.StartRecommenderInput) (*request.Request, *personalize.StartRecommenderOutput)

	StopRecommender(*personalize.StopRecommenderInput) (*personalize.StopRecommenderOutput, error)
	StopRecommenderWithContext(aws.Context, *personalize.StopRecommenderInput, ...request.Option) (*personalize.StopRecommenderOutput, error)
	StopRecommenderRequest(*personalize.StopRecommenderInput) (*request.Request, *personalize.StopRecommenderOutput)

	StopSolutionVersionCreation(*personalize.StopSolutionVersionCreationInput) (*personalize.StopSolutionVersionCreationOutput, error)
	StopSolutionVersionCreationWithContext(aws.Context, *personalize.StopSolutionVersionCreationInput, ...request.Option) (*personalize.StopSolutionVersionCreationOutput, error)
	StopSolutionVersionCreationRequest(*personalize.StopSolutionVersionCreationInput) (*request.Request, *personalize.StopSolutionVersionCreationOutput)

	TagResource(*personalize.TagResourceInput) (*personalize.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *personalize.TagResourceInput, ...request.Option) (*personalize.TagResourceOutput, error)
	TagResourceRequest(*personalize.TagResourceInput) (*request.Request, *personalize.TagResourceOutput)

	UntagResource(*personalize.UntagResourceInput) (*personalize.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *personalize.UntagResourceInput, ...request.Option) (*personalize.UntagResourceOutput, error)
	UntagResourceRequest(*personalize.UntagResourceInput) (*request.Request, *personalize.UntagResourceOutput)

	UpdateCampaign(*personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error)
	UpdateCampaignWithContext(aws.Context, *personalize.UpdateCampaignInput, ...request.Option) (*personalize.UpdateCampaignOutput, error)
	UpdateCampaignRequest(*personalize.UpdateCampaignInput) (*request.Request, *personalize.UpdateCampaignOutput)

	UpdateRecommender(*personalize.UpdateRecommenderInput) (*personalize.UpdateRecommenderOutput, error)
	UpdateRecommenderWithContext(aws.Context, *personalize.UpdateRecommenderInput, ...request.Option) (*personalize.UpdateRecommenderOutput, error)
	UpdateRecommenderRequest(*personalize.UpdateRecommenderInput) (*request.Request, *personalize.UpdateRecommenderOutput)
}

var _ PersonalizeAPI = (*personalize.Personalize)(nil)
