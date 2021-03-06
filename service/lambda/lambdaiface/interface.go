// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lambdaiface provides an interface for the AWS Lambda.
package lambdaiface

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaAPI is the interface type for lambda.Lambda.
type LambdaAPI interface {
	AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)

	CreateEventSourceMapping(*lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	CreateFunction(*lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)

	DeleteEventSourceMapping(*lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)

	GetEventSourceMapping(*lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)

	GetPolicy(*lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)

	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, error)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error)

	ListEventSourceMappings(*lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)

	RemovePermission(*lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)

	UpdateEventSourceMapping(*lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)

	UpdateFunctionCode(*lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
}
