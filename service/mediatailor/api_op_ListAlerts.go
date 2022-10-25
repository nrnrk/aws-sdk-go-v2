// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the alerts that are associated with a MediaTailor channel assembly
// resource.
func (c *Client) ListAlerts(ctx context.Context, params *ListAlertsInput, optFns ...func(*Options)) (*ListAlertsOutput, error) {
	if params == nil {
		params = &ListAlertsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAlerts", params, optFns, c.addOperationListAlertsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAlertsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAlertsInput struct {

	// The Amazon Resource Name (ARN) of the resource.
	//
	// This member is required.
	ResourceArn *string

	// The maximum number of alerts that you want MediaTailor to return in response to
	// the current request. If there are more than MaxResults alerts, use the value of
	// NextToken in the response to get the next page of results.
	MaxResults int32

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAlertsOutput struct {

	// A list of alerts that are associated with this resource.
	Items []types.Alert

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAlertsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAlerts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAlerts{}, middleware.After)
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
	if err = addOpListAlertsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAlerts(options.Region), middleware.Before); err != nil {
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

// ListAlertsAPIClient is a client that implements the ListAlerts operation.
type ListAlertsAPIClient interface {
	ListAlerts(context.Context, *ListAlertsInput, ...func(*Options)) (*ListAlertsOutput, error)
}

var _ ListAlertsAPIClient = (*Client)(nil)

// ListAlertsPaginatorOptions is the paginator options for ListAlerts
type ListAlertsPaginatorOptions struct {
	// The maximum number of alerts that you want MediaTailor to return in response to
	// the current request. If there are more than MaxResults alerts, use the value of
	// NextToken in the response to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAlertsPaginator is a paginator for ListAlerts
type ListAlertsPaginator struct {
	options   ListAlertsPaginatorOptions
	client    ListAlertsAPIClient
	params    *ListAlertsInput
	nextToken *string
	firstPage bool
}

// NewListAlertsPaginator returns a new ListAlertsPaginator
func NewListAlertsPaginator(client ListAlertsAPIClient, params *ListAlertsInput, optFns ...func(*ListAlertsPaginatorOptions)) *ListAlertsPaginator {
	if params == nil {
		params = &ListAlertsInput{}
	}

	options := ListAlertsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAlertsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAlertsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAlerts page.
func (p *ListAlertsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAlertsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAlerts(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAlerts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListAlerts",
	}
}
