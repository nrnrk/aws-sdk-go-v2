// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The live source configuration.
func (c *Client) CreateLiveSource(ctx context.Context, params *CreateLiveSourceInput, optFns ...func(*Options)) (*CreateLiveSourceOutput, error) {
	if params == nil {
		params = &CreateLiveSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLiveSource", params, optFns, c.addOperationCreateLiveSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLiveSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLiveSourceInput struct {

	// A list of HTTP package configuration parameters for this live source.
	//
	// This member is required.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The name of the live source.
	//
	// This member is required.
	LiveSourceName *string

	// The name of the source location.
	//
	// This member is required.
	SourceLocationName *string

	// The tags to assign to the live source. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLiveSourceOutput struct {

	// The ARN to assign to the live source.
	Arn *string

	// The time the live source was created.
	CreationTime *time.Time

	// A list of HTTP package configuration parameters for this live source.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The time the live source was last modified.
	LastModifiedTime *time.Time

	// The name to assign to the live source.
	LiveSourceName *string

	// The name to assign to the source location of the live source.
	SourceLocationName *string

	// The tags to assign to the live source. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLiveSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLiveSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLiveSource{}, middleware.After)
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
	if err = addOpCreateLiveSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLiveSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLiveSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "CreateLiveSource",
	}
}
