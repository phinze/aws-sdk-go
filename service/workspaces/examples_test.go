// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package workspaces_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleWorkSpaces_CreateWorkspaces() {
	svc := workspaces.New(nil)

	params := &workspaces.CreateWorkspacesInput{
		Workspaces: []*workspaces.WorkspaceRequest{ // Required
			{ // Required
				BundleID:    aws.String("BundleId"),    // Required
				DirectoryID: aws.String("DirectoryId"), // Required
				UserName:    aws.String("UserName"),    // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreateWorkspaces(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleWorkSpaces_DescribeWorkspaceBundles() {
	svc := workspaces.New(nil)

	params := &workspaces.DescribeWorkspaceBundlesInput{
		BundleIDs: []*string{
			aws.String("BundleId"), // Required
			// More values...
		},
		NextToken: aws.String("PaginationToken"),
		Owner:     aws.String("BundleOwner"),
	}
	resp, err := svc.DescribeWorkspaceBundles(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleWorkSpaces_DescribeWorkspaceDirectories() {
	svc := workspaces.New(nil)

	params := &workspaces.DescribeWorkspaceDirectoriesInput{
		DirectoryIDs: []*string{
			aws.String("DirectoryId"), // Required
			// More values...
		},
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.DescribeWorkspaceDirectories(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleWorkSpaces_DescribeWorkspaces() {
	svc := workspaces.New(nil)

	params := &workspaces.DescribeWorkspacesInput{
		BundleID:    aws.String("BundleId"),
		DirectoryID: aws.String("DirectoryId"),
		Limit:       aws.Long(1),
		NextToken:   aws.String("PaginationToken"),
		UserName:    aws.String("UserName"),
		WorkspaceIDs: []*string{
			aws.String("WorkspaceId"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeWorkspaces(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleWorkSpaces_RebootWorkspaces() {
	svc := workspaces.New(nil)

	params := &workspaces.RebootWorkspacesInput{
		RebootWorkspaceRequests: []*workspaces.RebootRequest{ // Required
			{ // Required
				WorkspaceID: aws.String("WorkspaceId"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.RebootWorkspaces(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleWorkSpaces_RebuildWorkspaces() {
	svc := workspaces.New(nil)

	params := &workspaces.RebuildWorkspacesInput{
		RebuildWorkspaceRequests: []*workspaces.RebuildRequest{ // Required
			{ // Required
				WorkspaceID: aws.String("WorkspaceId"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.RebuildWorkspaces(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleWorkSpaces_TerminateWorkspaces() {
	svc := workspaces.New(nil)

	params := &workspaces.TerminateWorkspacesInput{
		TerminateWorkspaceRequests: []*workspaces.TerminateRequest{ // Required
			{ // Required
				WorkspaceID: aws.String("WorkspaceId"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.TerminateWorkspaces(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
