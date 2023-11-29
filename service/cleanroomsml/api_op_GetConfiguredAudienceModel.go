// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a specified configured audience model.
func (c *Client) GetConfiguredAudienceModel(ctx context.Context, params *GetConfiguredAudienceModelInput, optFns ...func(*Options)) (*GetConfiguredAudienceModelOutput, error) {
	if params == nil {
		params = &GetConfiguredAudienceModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfiguredAudienceModel", params, optFns, c.addOperationGetConfiguredAudienceModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfiguredAudienceModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfiguredAudienceModelInput struct {

	// The Amazon Resource Name (ARN) of the configured audience model that you are
	// interested in.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	noSmithyDocumentSerde
}

type GetConfiguredAudienceModelOutput struct {

	// The Amazon Resource Name (ARN) of the audience model used for this configured
	// audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The Amazon Resource Name (ARN) of the configured audience model.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The time at which the configured audience model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The name of the configured audience model.
	//
	// This member is required.
	Name *string

	// The output configuration of the configured audience model
	//
	// This member is required.
	OutputConfig *types.ConfiguredAudienceModelOutputConfig

	// Whether audience metrics are shared.
	//
	// This member is required.
	SharedAudienceMetrics []types.SharedAudienceMetrics

	// The status of the configured audience model.
	//
	// This member is required.
	Status types.ConfiguredAudienceModelStatus

	// The most recent time at which the configured audience model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The list of output sizes of audiences that can be created using this configured
	// audience model. A request to StartAudienceGenerationJob that uses this
	// configured audience model must have an audienceSize selected from this list.
	// You can use the ABSOLUTE AudienceSize to configure out audience sizes using the
	// count of identifiers in the output. You can use the Percentage AudienceSize to
	// configure sizes in the range 1-100 percent.
	AudienceSizeConfig *types.AudienceSizeConfig

	// Provides the childResourceTagOnCreatePolicy that was used for this configured
	// audience model.
	ChildResourceTagOnCreatePolicy types.TagOnCreatePolicy

	// The description of the configured audience model.
	Description *string

	// The minimum number of users from the seed audience that must match with users
	// in the training data of the audience model.
	MinMatchingSeedSize *int32

	// The tags that are associated to this configured audience model.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfiguredAudienceModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfiguredAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfiguredAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConfiguredAudienceModel"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetConfiguredAudienceModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfiguredAudienceModel(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetConfiguredAudienceModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConfiguredAudienceModel",
	}
}
