// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

const (

	// ErrCodeArgumentException for service response error code
	// "ArgumentException".
	//
	// An invalid argument was specified.
	ErrCodeArgumentException = "ArgumentException"

	// ErrCodeIdempotencyException for service response error code
	// "IdempotencyException".
	//
	// An entity with the same name already exists.
	ErrCodeIdempotencyException = "IdempotencyException"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// There was an error with the update request, or you do not have sufficient
	// permissions to update this VPC endpoint configuration.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A limit was exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotEligibleException for service response error code
	// "NotEligibleException".
	//
	// Exception gets thrown when a user is not eligible to perform the specified
	// transaction.
	ErrCodeNotEligibleException = "NotEligibleException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The specified entity was not found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeServiceAccountException for service response error code
	// "ServiceAccountException".
	//
	// There was a problem with the service account.
	ErrCodeServiceAccountException = "ServiceAccountException"

	// ErrCodeTagOperationException for service response error code
	// "TagOperationException".
	//
	// The operation was not successful. Try again.
	ErrCodeTagOperationException = "TagOperationException"

	// ErrCodeTagPolicyException for service response error code
	// "TagPolicyException".
	//
	// The request doesn't comply with the AWS Identity and Access Management (IAM)
	// tag policy. Correct your request and then retry it.
	ErrCodeTagPolicyException = "TagPolicyException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The list of tags on the repository is over the limit. The maximum number
	// of tags that can be applied to a repository is 50.
	ErrCodeTooManyTagsException = "TooManyTagsException"
)