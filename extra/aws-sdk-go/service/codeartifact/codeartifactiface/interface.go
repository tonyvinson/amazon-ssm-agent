// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codeartifactiface provides an interface to enable mocking the CodeArtifact service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codeartifactiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codeartifact"
)

// CodeArtifactAPI provides an interface to enable mocking the
// codeartifact.CodeArtifact service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // CodeArtifact.
//    func myFunc(svc codeartifactiface.CodeArtifactAPI) bool {
//        // Make svc.AssociateExternalConnection request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codeartifact.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeArtifactClient struct {
//        codeartifactiface.CodeArtifactAPI
//    }
//    func (m *mockCodeArtifactClient) AssociateExternalConnection(input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeArtifactClient{}
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
type CodeArtifactAPI interface {
	AssociateExternalConnection(*codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error)
	AssociateExternalConnectionWithContext(aws.Context, *codeartifact.AssociateExternalConnectionInput, ...request.Option) (*codeartifact.AssociateExternalConnectionOutput, error)
	AssociateExternalConnectionRequest(*codeartifact.AssociateExternalConnectionInput) (*request.Request, *codeartifact.AssociateExternalConnectionOutput)

	CopyPackageVersions(*codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error)
	CopyPackageVersionsWithContext(aws.Context, *codeartifact.CopyPackageVersionsInput, ...request.Option) (*codeartifact.CopyPackageVersionsOutput, error)
	CopyPackageVersionsRequest(*codeartifact.CopyPackageVersionsInput) (*request.Request, *codeartifact.CopyPackageVersionsOutput)

	CreateDomain(*codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *codeartifact.CreateDomainInput, ...request.Option) (*codeartifact.CreateDomainOutput, error)
	CreateDomainRequest(*codeartifact.CreateDomainInput) (*request.Request, *codeartifact.CreateDomainOutput)

	CreateRepository(*codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error)
	CreateRepositoryWithContext(aws.Context, *codeartifact.CreateRepositoryInput, ...request.Option) (*codeartifact.CreateRepositoryOutput, error)
	CreateRepositoryRequest(*codeartifact.CreateRepositoryInput) (*request.Request, *codeartifact.CreateRepositoryOutput)

	DeleteDomain(*codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *codeartifact.DeleteDomainInput, ...request.Option) (*codeartifact.DeleteDomainOutput, error)
	DeleteDomainRequest(*codeartifact.DeleteDomainInput) (*request.Request, *codeartifact.DeleteDomainOutput)

	DeleteDomainPermissionsPolicy(*codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error)
	DeleteDomainPermissionsPolicyWithContext(aws.Context, *codeartifact.DeleteDomainPermissionsPolicyInput, ...request.Option) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error)
	DeleteDomainPermissionsPolicyRequest(*codeartifact.DeleteDomainPermissionsPolicyInput) (*request.Request, *codeartifact.DeleteDomainPermissionsPolicyOutput)

	DeletePackageVersions(*codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error)
	DeletePackageVersionsWithContext(aws.Context, *codeartifact.DeletePackageVersionsInput, ...request.Option) (*codeartifact.DeletePackageVersionsOutput, error)
	DeletePackageVersionsRequest(*codeartifact.DeletePackageVersionsInput) (*request.Request, *codeartifact.DeletePackageVersionsOutput)

	DeleteRepository(*codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error)
	DeleteRepositoryWithContext(aws.Context, *codeartifact.DeleteRepositoryInput, ...request.Option) (*codeartifact.DeleteRepositoryOutput, error)
	DeleteRepositoryRequest(*codeartifact.DeleteRepositoryInput) (*request.Request, *codeartifact.DeleteRepositoryOutput)

	DeleteRepositoryPermissionsPolicy(*codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error)
	DeleteRepositoryPermissionsPolicyWithContext(aws.Context, *codeartifact.DeleteRepositoryPermissionsPolicyInput, ...request.Option) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error)
	DeleteRepositoryPermissionsPolicyRequest(*codeartifact.DeleteRepositoryPermissionsPolicyInput) (*request.Request, *codeartifact.DeleteRepositoryPermissionsPolicyOutput)

	DescribeDomain(*codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error)
	DescribeDomainWithContext(aws.Context, *codeartifact.DescribeDomainInput, ...request.Option) (*codeartifact.DescribeDomainOutput, error)
	DescribeDomainRequest(*codeartifact.DescribeDomainInput) (*request.Request, *codeartifact.DescribeDomainOutput)

	DescribePackage(*codeartifact.DescribePackageInput) (*codeartifact.DescribePackageOutput, error)
	DescribePackageWithContext(aws.Context, *codeartifact.DescribePackageInput, ...request.Option) (*codeartifact.DescribePackageOutput, error)
	DescribePackageRequest(*codeartifact.DescribePackageInput) (*request.Request, *codeartifact.DescribePackageOutput)

	DescribePackageVersion(*codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error)
	DescribePackageVersionWithContext(aws.Context, *codeartifact.DescribePackageVersionInput, ...request.Option) (*codeartifact.DescribePackageVersionOutput, error)
	DescribePackageVersionRequest(*codeartifact.DescribePackageVersionInput) (*request.Request, *codeartifact.DescribePackageVersionOutput)

	DescribeRepository(*codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error)
	DescribeRepositoryWithContext(aws.Context, *codeartifact.DescribeRepositoryInput, ...request.Option) (*codeartifact.DescribeRepositoryOutput, error)
	DescribeRepositoryRequest(*codeartifact.DescribeRepositoryInput) (*request.Request, *codeartifact.DescribeRepositoryOutput)

	DisassociateExternalConnection(*codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error)
	DisassociateExternalConnectionWithContext(aws.Context, *codeartifact.DisassociateExternalConnectionInput, ...request.Option) (*codeartifact.DisassociateExternalConnectionOutput, error)
	DisassociateExternalConnectionRequest(*codeartifact.DisassociateExternalConnectionInput) (*request.Request, *codeartifact.DisassociateExternalConnectionOutput)

	DisposePackageVersions(*codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error)
	DisposePackageVersionsWithContext(aws.Context, *codeartifact.DisposePackageVersionsInput, ...request.Option) (*codeartifact.DisposePackageVersionsOutput, error)
	DisposePackageVersionsRequest(*codeartifact.DisposePackageVersionsInput) (*request.Request, *codeartifact.DisposePackageVersionsOutput)

	GetAuthorizationToken(*codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenWithContext(aws.Context, *codeartifact.GetAuthorizationTokenInput, ...request.Option) (*codeartifact.GetAuthorizationTokenOutput, error)
	GetAuthorizationTokenRequest(*codeartifact.GetAuthorizationTokenInput) (*request.Request, *codeartifact.GetAuthorizationTokenOutput)

	GetDomainPermissionsPolicy(*codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error)
	GetDomainPermissionsPolicyWithContext(aws.Context, *codeartifact.GetDomainPermissionsPolicyInput, ...request.Option) (*codeartifact.GetDomainPermissionsPolicyOutput, error)
	GetDomainPermissionsPolicyRequest(*codeartifact.GetDomainPermissionsPolicyInput) (*request.Request, *codeartifact.GetDomainPermissionsPolicyOutput)

	GetPackageVersionAsset(*codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error)
	GetPackageVersionAssetWithContext(aws.Context, *codeartifact.GetPackageVersionAssetInput, ...request.Option) (*codeartifact.GetPackageVersionAssetOutput, error)
	GetPackageVersionAssetRequest(*codeartifact.GetPackageVersionAssetInput) (*request.Request, *codeartifact.GetPackageVersionAssetOutput)

	GetPackageVersionReadme(*codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error)
	GetPackageVersionReadmeWithContext(aws.Context, *codeartifact.GetPackageVersionReadmeInput, ...request.Option) (*codeartifact.GetPackageVersionReadmeOutput, error)
	GetPackageVersionReadmeRequest(*codeartifact.GetPackageVersionReadmeInput) (*request.Request, *codeartifact.GetPackageVersionReadmeOutput)

	GetRepositoryEndpoint(*codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error)
	GetRepositoryEndpointWithContext(aws.Context, *codeartifact.GetRepositoryEndpointInput, ...request.Option) (*codeartifact.GetRepositoryEndpointOutput, error)
	GetRepositoryEndpointRequest(*codeartifact.GetRepositoryEndpointInput) (*request.Request, *codeartifact.GetRepositoryEndpointOutput)

	GetRepositoryPermissionsPolicy(*codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error)
	GetRepositoryPermissionsPolicyWithContext(aws.Context, *codeartifact.GetRepositoryPermissionsPolicyInput, ...request.Option) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error)
	GetRepositoryPermissionsPolicyRequest(*codeartifact.GetRepositoryPermissionsPolicyInput) (*request.Request, *codeartifact.GetRepositoryPermissionsPolicyOutput)

	ListDomains(*codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error)
	ListDomainsWithContext(aws.Context, *codeartifact.ListDomainsInput, ...request.Option) (*codeartifact.ListDomainsOutput, error)
	ListDomainsRequest(*codeartifact.ListDomainsInput) (*request.Request, *codeartifact.ListDomainsOutput)

	ListDomainsPages(*codeartifact.ListDomainsInput, func(*codeartifact.ListDomainsOutput, bool) bool) error
	ListDomainsPagesWithContext(aws.Context, *codeartifact.ListDomainsInput, func(*codeartifact.ListDomainsOutput, bool) bool, ...request.Option) error

	ListPackageVersionAssets(*codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error)
	ListPackageVersionAssetsWithContext(aws.Context, *codeartifact.ListPackageVersionAssetsInput, ...request.Option) (*codeartifact.ListPackageVersionAssetsOutput, error)
	ListPackageVersionAssetsRequest(*codeartifact.ListPackageVersionAssetsInput) (*request.Request, *codeartifact.ListPackageVersionAssetsOutput)

	ListPackageVersionAssetsPages(*codeartifact.ListPackageVersionAssetsInput, func(*codeartifact.ListPackageVersionAssetsOutput, bool) bool) error
	ListPackageVersionAssetsPagesWithContext(aws.Context, *codeartifact.ListPackageVersionAssetsInput, func(*codeartifact.ListPackageVersionAssetsOutput, bool) bool, ...request.Option) error

	ListPackageVersionDependencies(*codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error)
	ListPackageVersionDependenciesWithContext(aws.Context, *codeartifact.ListPackageVersionDependenciesInput, ...request.Option) (*codeartifact.ListPackageVersionDependenciesOutput, error)
	ListPackageVersionDependenciesRequest(*codeartifact.ListPackageVersionDependenciesInput) (*request.Request, *codeartifact.ListPackageVersionDependenciesOutput)

	ListPackageVersions(*codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error)
	ListPackageVersionsWithContext(aws.Context, *codeartifact.ListPackageVersionsInput, ...request.Option) (*codeartifact.ListPackageVersionsOutput, error)
	ListPackageVersionsRequest(*codeartifact.ListPackageVersionsInput) (*request.Request, *codeartifact.ListPackageVersionsOutput)

	ListPackageVersionsPages(*codeartifact.ListPackageVersionsInput, func(*codeartifact.ListPackageVersionsOutput, bool) bool) error
	ListPackageVersionsPagesWithContext(aws.Context, *codeartifact.ListPackageVersionsInput, func(*codeartifact.ListPackageVersionsOutput, bool) bool, ...request.Option) error

	ListPackages(*codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error)
	ListPackagesWithContext(aws.Context, *codeartifact.ListPackagesInput, ...request.Option) (*codeartifact.ListPackagesOutput, error)
	ListPackagesRequest(*codeartifact.ListPackagesInput) (*request.Request, *codeartifact.ListPackagesOutput)

	ListPackagesPages(*codeartifact.ListPackagesInput, func(*codeartifact.ListPackagesOutput, bool) bool) error
	ListPackagesPagesWithContext(aws.Context, *codeartifact.ListPackagesInput, func(*codeartifact.ListPackagesOutput, bool) bool, ...request.Option) error

	ListRepositories(*codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error)
	ListRepositoriesWithContext(aws.Context, *codeartifact.ListRepositoriesInput, ...request.Option) (*codeartifact.ListRepositoriesOutput, error)
	ListRepositoriesRequest(*codeartifact.ListRepositoriesInput) (*request.Request, *codeartifact.ListRepositoriesOutput)

	ListRepositoriesPages(*codeartifact.ListRepositoriesInput, func(*codeartifact.ListRepositoriesOutput, bool) bool) error
	ListRepositoriesPagesWithContext(aws.Context, *codeartifact.ListRepositoriesInput, func(*codeartifact.ListRepositoriesOutput, bool) bool, ...request.Option) error

	ListRepositoriesInDomain(*codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error)
	ListRepositoriesInDomainWithContext(aws.Context, *codeartifact.ListRepositoriesInDomainInput, ...request.Option) (*codeartifact.ListRepositoriesInDomainOutput, error)
	ListRepositoriesInDomainRequest(*codeartifact.ListRepositoriesInDomainInput) (*request.Request, *codeartifact.ListRepositoriesInDomainOutput)

	ListRepositoriesInDomainPages(*codeartifact.ListRepositoriesInDomainInput, func(*codeartifact.ListRepositoriesInDomainOutput, bool) bool) error
	ListRepositoriesInDomainPagesWithContext(aws.Context, *codeartifact.ListRepositoriesInDomainInput, func(*codeartifact.ListRepositoriesInDomainOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codeartifact.ListTagsForResourceInput) (*codeartifact.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codeartifact.ListTagsForResourceInput, ...request.Option) (*codeartifact.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codeartifact.ListTagsForResourceInput) (*request.Request, *codeartifact.ListTagsForResourceOutput)

	PutDomainPermissionsPolicy(*codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error)
	PutDomainPermissionsPolicyWithContext(aws.Context, *codeartifact.PutDomainPermissionsPolicyInput, ...request.Option) (*codeartifact.PutDomainPermissionsPolicyOutput, error)
	PutDomainPermissionsPolicyRequest(*codeartifact.PutDomainPermissionsPolicyInput) (*request.Request, *codeartifact.PutDomainPermissionsPolicyOutput)

	PutPackageOriginConfiguration(*codeartifact.PutPackageOriginConfigurationInput) (*codeartifact.PutPackageOriginConfigurationOutput, error)
	PutPackageOriginConfigurationWithContext(aws.Context, *codeartifact.PutPackageOriginConfigurationInput, ...request.Option) (*codeartifact.PutPackageOriginConfigurationOutput, error)
	PutPackageOriginConfigurationRequest(*codeartifact.PutPackageOriginConfigurationInput) (*request.Request, *codeartifact.PutPackageOriginConfigurationOutput)

	PutRepositoryPermissionsPolicy(*codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error)
	PutRepositoryPermissionsPolicyWithContext(aws.Context, *codeartifact.PutRepositoryPermissionsPolicyInput, ...request.Option) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error)
	PutRepositoryPermissionsPolicyRequest(*codeartifact.PutRepositoryPermissionsPolicyInput) (*request.Request, *codeartifact.PutRepositoryPermissionsPolicyOutput)

	TagResource(*codeartifact.TagResourceInput) (*codeartifact.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codeartifact.TagResourceInput, ...request.Option) (*codeartifact.TagResourceOutput, error)
	TagResourceRequest(*codeartifact.TagResourceInput) (*request.Request, *codeartifact.TagResourceOutput)

	UntagResource(*codeartifact.UntagResourceInput) (*codeartifact.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codeartifact.UntagResourceInput, ...request.Option) (*codeartifact.UntagResourceOutput, error)
	UntagResourceRequest(*codeartifact.UntagResourceInput) (*request.Request, *codeartifact.UntagResourceOutput)

	UpdatePackageVersionsStatus(*codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error)
	UpdatePackageVersionsStatusWithContext(aws.Context, *codeartifact.UpdatePackageVersionsStatusInput, ...request.Option) (*codeartifact.UpdatePackageVersionsStatusOutput, error)
	UpdatePackageVersionsStatusRequest(*codeartifact.UpdatePackageVersionsStatusInput) (*request.Request, *codeartifact.UpdatePackageVersionsStatusOutput)

	UpdateRepository(*codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error)
	UpdateRepositoryWithContext(aws.Context, *codeartifact.UpdateRepositoryInput, ...request.Option) (*codeartifact.UpdateRepositoryOutput, error)
	UpdateRepositoryRequest(*codeartifact.UpdateRepositoryInput) (*request.Request, *codeartifact.UpdateRepositoryOutput)
}

var _ CodeArtifactAPI = (*codeartifact.CodeArtifact)(nil)
