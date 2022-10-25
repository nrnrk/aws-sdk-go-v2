// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified WorkSpace properties. For important information about how
// to modify the size of the root and user volumes, see  Modify a WorkSpace
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html).
// The MANUAL running mode value is only supported by Amazon WorkSpaces Core.
// Contact your account team to be allow-listed to use this value. For more
// information, see Amazon WorkSpaces Core
// (http://aws.amazon.com/workspaces/core/).
func (c *Client) ModifyWorkspaceProperties(ctx context.Context, params *ModifyWorkspacePropertiesInput, optFns ...func(*Options)) (*ModifyWorkspacePropertiesOutput, error) {
	if params == nil {
		params = &ModifyWorkspacePropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyWorkspaceProperties", params, optFns, c.addOperationModifyWorkspacePropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyWorkspacePropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyWorkspacePropertiesInput struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string

	// The properties of the WorkSpace.
	//
	// This member is required.
	WorkspaceProperties *types.WorkspaceProperties

	noSmithyDocumentSerde
}

type ModifyWorkspacePropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyWorkspacePropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyWorkspaceProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyWorkspaceProperties{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyWorkspacePropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyWorkspaceProperties(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyWorkspaceProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyWorkspaceProperties",
	}
}
