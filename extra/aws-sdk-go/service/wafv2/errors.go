// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeWAFAssociatedItemException for service response error code
	// "WAFAssociatedItemException".
	//
	// WAF couldn’t perform the operation because your resource is being used
	// by another resource or it’s associated with another resource.
	ErrCodeWAFAssociatedItemException = "WAFAssociatedItemException"

	// ErrCodeWAFConfigurationWarningException for service response error code
	// "WAFConfigurationWarningException".
	//
	// The operation failed because you are inspecting the web request body, headers,
	// or cookies without specifying how to handle oversize components. Rules that
	// inspect the body must either provide an OversizeHandling configuration or
	// they must be preceded by a SizeConstraintStatement that blocks the body content
	// from being too large. Rules that inspect the headers or cookies must provide
	// an OversizeHandling configuration.
	//
	// Provide the handling configuration and retry your operation.
	//
	// Alternately, you can suppress this warning by adding the following tag to
	// the resource that you provide to this operation: Tag (key:WAF:OversizeFieldsHandlingConstraintOptOut,
	// value:true).
	ErrCodeWAFConfigurationWarningException = "WAFConfigurationWarningException"

	// ErrCodeWAFDuplicateItemException for service response error code
	// "WAFDuplicateItemException".
	//
	// WAF couldn’t perform the operation because the resource that you tried
	// to save is a duplicate of an existing one.
	ErrCodeWAFDuplicateItemException = "WAFDuplicateItemException"

	// ErrCodeWAFExpiredManagedRuleGroupVersionException for service response error code
	// "WAFExpiredManagedRuleGroupVersionException".
	//
	// The operation failed because the specified version for the managed rule group
	// has expired. You can retrieve the available versions for the managed rule
	// group by calling ListAvailableManagedRuleGroupVersions.
	ErrCodeWAFExpiredManagedRuleGroupVersionException = "WAFExpiredManagedRuleGroupVersionException"

	// ErrCodeWAFInternalErrorException for service response error code
	// "WAFInternalErrorException".
	//
	// Your request is valid, but WAF couldn’t perform the operation because of
	// a system problem. Retry your request.
	ErrCodeWAFInternalErrorException = "WAFInternalErrorException"

	// ErrCodeWAFInvalidOperationException for service response error code
	// "WAFInvalidOperationException".
	//
	// The operation isn't valid.
	ErrCodeWAFInvalidOperationException = "WAFInvalidOperationException"

	// ErrCodeWAFInvalidParameterException for service response error code
	// "WAFInvalidParameterException".
	//
	// The operation failed because WAF didn't recognize a parameter in the request.
	// For example:
	//
	//    * You specified a parameter name or value that isn't valid.
	//
	//    * Your nested statement isn't valid. You might have tried to nest a statement
	//    that can’t be nested.
	//
	//    * You tried to update a WebACL with a DefaultAction that isn't among the
	//    types available at DefaultAction.
	//
	//    * Your request references an ARN that is malformed, or corresponds to
	//    a resource with which a web ACL can't be associated.
	ErrCodeWAFInvalidParameterException = "WAFInvalidParameterException"

	// ErrCodeWAFInvalidPermissionPolicyException for service response error code
	// "WAFInvalidPermissionPolicyException".
	//
	// The operation failed because the specified policy isn't in the proper format.
	//
	// The policy specifications must conform to the following:
	//
	//    * The policy must be composed using IAM Policy version 2012-10-17.
	//
	//    * The policy must include specifications for Effect, Action, and Principal.
	//
	//    * Effect must specify Allow.
	//
	//    * Action must specify wafv2:CreateWebACL, wafv2:UpdateWebACL, and wafv2:PutFirewallManagerRuleGroups
	//    and may optionally specify wafv2:GetRuleGroup. WAF rejects any extra actions
	//    or wildcard actions in the policy.
	//
	//    * The policy must not include a Resource parameter.
	//
	// For more information, see IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).
	ErrCodeWAFInvalidPermissionPolicyException = "WAFInvalidPermissionPolicyException"

	// ErrCodeWAFInvalidResourceException for service response error code
	// "WAFInvalidResourceException".
	//
	// WAF couldn’t perform the operation because the resource that you requested
	// isn’t valid. Check the resource, and try again.
	ErrCodeWAFInvalidResourceException = "WAFInvalidResourceException"

	// ErrCodeWAFLimitsExceededException for service response error code
	// "WAFLimitsExceededException".
	//
	// WAF couldn’t perform the operation because you exceeded your resource limit.
	// For example, the maximum number of WebACL objects that you can create for
	// an Amazon Web Services account. For more information, see WAF quotas (https://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
	// in the WAF Developer Guide.
	ErrCodeWAFLimitsExceededException = "WAFLimitsExceededException"

	// ErrCodeWAFLogDestinationPermissionIssueException for service response error code
	// "WAFLogDestinationPermissionIssueException".
	//
	// The operation failed because you don't have the permissions that your logging
	// configuration requires. For information, see Logging web ACL traffic information
	// (https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) in the
	// WAF Developer Guide.
	ErrCodeWAFLogDestinationPermissionIssueException = "WAFLogDestinationPermissionIssueException"

	// ErrCodeWAFNonexistentItemException for service response error code
	// "WAFNonexistentItemException".
	//
	// WAF couldn’t perform the operation because your resource doesn't exist.
	// If you've just created a resource that you're using in this operation, you
	// might just need to wait a few minutes. It can take from a few seconds to
	// a number of minutes for changes to propagate.
	ErrCodeWAFNonexistentItemException = "WAFNonexistentItemException"

	// ErrCodeWAFOptimisticLockException for service response error code
	// "WAFOptimisticLockException".
	//
	// WAF couldn’t save your changes because you tried to update or delete a
	// resource that has changed since you last retrieved it. Get the resource again,
	// make any changes you need to make to the new copy, and retry your operation.
	ErrCodeWAFOptimisticLockException = "WAFOptimisticLockException"

	// ErrCodeWAFServiceLinkedRoleErrorException for service response error code
	// "WAFServiceLinkedRoleErrorException".
	//
	// WAF is not able to access the service linked role. This can be caused by
	// a previous PutLoggingConfiguration request, which can lock the service linked
	// role for about 20 seconds. Please try your request again. The service linked
	// role can also be locked by a previous DeleteServiceLinkedRole request, which
	// can lock the role for 15 minutes or more. If you recently made a call to
	// DeleteServiceLinkedRole, wait at least 15 minutes and try the request again.
	// If you receive this same exception again, you will have to wait additional
	// time until the role is unlocked.
	ErrCodeWAFServiceLinkedRoleErrorException = "WAFServiceLinkedRoleErrorException"

	// ErrCodeWAFSubscriptionNotFoundException for service response error code
	// "WAFSubscriptionNotFoundException".
	//
	// You tried to use a managed rule group that's available by subscription, but
	// you aren't subscribed to it yet.
	ErrCodeWAFSubscriptionNotFoundException = "WAFSubscriptionNotFoundException"

	// ErrCodeWAFTagOperationException for service response error code
	// "WAFTagOperationException".
	//
	// An error occurred during the tagging operation. Retry your request.
	ErrCodeWAFTagOperationException = "WAFTagOperationException"

	// ErrCodeWAFTagOperationInternalErrorException for service response error code
	// "WAFTagOperationInternalErrorException".
	//
	// WAF couldn’t perform your tagging operation because of an internal error.
	// Retry your request.
	ErrCodeWAFTagOperationInternalErrorException = "WAFTagOperationInternalErrorException"

	// ErrCodeWAFUnavailableEntityException for service response error code
	// "WAFUnavailableEntityException".
	//
	// WAF couldn’t retrieve a resource that you specified for this operation.
	// If you've just created a resource that you're using in this operation, you
	// might just need to wait a few minutes. It can take from a few seconds to
	// a number of minutes for changes to propagate. Verify the resources that you
	// are specifying in your request parameters and then retry the operation.
	ErrCodeWAFUnavailableEntityException = "WAFUnavailableEntityException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"WAFAssociatedItemException":                 newErrorWAFAssociatedItemException,
	"WAFConfigurationWarningException":           newErrorWAFConfigurationWarningException,
	"WAFDuplicateItemException":                  newErrorWAFDuplicateItemException,
	"WAFExpiredManagedRuleGroupVersionException": newErrorWAFExpiredManagedRuleGroupVersionException,
	"WAFInternalErrorException":                  newErrorWAFInternalErrorException,
	"WAFInvalidOperationException":               newErrorWAFInvalidOperationException,
	"WAFInvalidParameterException":               newErrorWAFInvalidParameterException,
	"WAFInvalidPermissionPolicyException":        newErrorWAFInvalidPermissionPolicyException,
	"WAFInvalidResourceException":                newErrorWAFInvalidResourceException,
	"WAFLimitsExceededException":                 newErrorWAFLimitsExceededException,
	"WAFLogDestinationPermissionIssueException":  newErrorWAFLogDestinationPermissionIssueException,
	"WAFNonexistentItemException":                newErrorWAFNonexistentItemException,
	"WAFOptimisticLockException":                 newErrorWAFOptimisticLockException,
	"WAFServiceLinkedRoleErrorException":         newErrorWAFServiceLinkedRoleErrorException,
	"WAFSubscriptionNotFoundException":           newErrorWAFSubscriptionNotFoundException,
	"WAFTagOperationException":                   newErrorWAFTagOperationException,
	"WAFTagOperationInternalErrorException":      newErrorWAFTagOperationInternalErrorException,
	"WAFUnavailableEntityException":              newErrorWAFUnavailableEntityException,
}
