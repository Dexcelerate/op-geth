// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates an Amazon Virtual Private Cloud (Amazon VPC) from an Amazon Route
// 53 private hosted zone. Note the following:
//
//   - You can't disassociate the last Amazon VPC from a private hosted zone.
//
//   - You can't convert a private hosted zone into a public hosted zone.
//
//   - You can submit a DisassociateVPCFromHostedZone request using either the
//     account that created the hosted zone or the account that created the Amazon VPC.
//
//   - Some services, such as Cloud Map and Amazon Elastic File System (Amazon
//     EFS) automatically create hosted zones and associate VPCs with the hosted zones.
//     A service can create a hosted zone using your account or using its own account.
//     You can disassociate a VPC from a hosted zone only if the service created the
//     hosted zone using your account. When you run DisassociateVPCFromHostedZone (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListHostedZonesByVPC.html)
//     , if the hosted zone has a value for OwningAccount , you can use
//     DisassociateVPCFromHostedZone . If the hosted zone has a value for
//     OwningService , you can't use DisassociateVPCFromHostedZone .
//
// When revoking access, the hosted zone and the Amazon VPC must belong to the
// same partition. A partition is a group of Amazon Web Services Regions. Each
// Amazon Web Services account is scoped to one partition. The following are the
// supported partitions:
//   - aws - Amazon Web Services Regions
//   - aws-cn - China Regions
//   - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see Access Management (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// in the Amazon Web Services General Reference.
func (c *Client) DisassociateVPCFromHostedZone(ctx context.Context, params *DisassociateVPCFromHostedZoneInput, optFns ...func(*Options)) (*DisassociateVPCFromHostedZoneOutput, error) {
	if params == nil {
		params = &DisassociateVPCFromHostedZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateVPCFromHostedZone", params, optFns, c.addOperationDisassociateVPCFromHostedZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateVPCFromHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the VPC that you want to
// disassociate from a specified private hosted zone.
type DisassociateVPCFromHostedZoneInput struct {

	// The ID of the private hosted zone that you want to disassociate a VPC from.
	//
	// This member is required.
	HostedZoneId *string

	// A complex type that contains information about the VPC that you're
	// disassociating from the specified hosted zone.
	//
	// This member is required.
	VPC *types.VPC

	// Optional: A comment about the disassociation request.
	Comment *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the disassociate
// request.
type DisassociateVPCFromHostedZoneOutput struct {

	// A complex type that describes the changes made to the specified private hosted
	// zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateVPCFromHostedZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDisassociateVPCFromHostedZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDisassociateVPCFromHostedZone{}, middleware.After)
	if err != nil {
		return err
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addDisassociateVPCFromHostedZoneResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDisassociateVPCFromHostedZoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateVPCFromHostedZone(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisassociateVPCFromHostedZone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DisassociateVPCFromHostedZone",
	}
}

type opDisassociateVPCFromHostedZoneResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDisassociateVPCFromHostedZoneResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDisassociateVPCFromHostedZoneResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "route53"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "route53"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("route53")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDisassociateVPCFromHostedZoneResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDisassociateVPCFromHostedZoneResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
