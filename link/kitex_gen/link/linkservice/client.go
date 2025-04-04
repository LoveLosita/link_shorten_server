// Code generated by Kitex v0.12.3. DO NOT EDIT.

package linkservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	link "link_shorten_server/link/kitex_gen/link"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GenerateLink(ctx context.Context, req *link.GenerateLinkRequest, callOptions ...callopt.Option) (r *link.GenerateLinkResponse, err error)
	DeleteLink(ctx context.Context, req *link.DeleteLinkRequest, callOptions ...callopt.Option) (r *link.DeleteLinkResponse, err error)
	ChangeLink(ctx context.Context, req *link.ChangeLinkRequest, callOptions ...callopt.Option) (r *link.ChangeLinkResponse, err error)
	SeeLinkRanking(ctx context.Context, req *link.SeeLinkRankingRequest, callOptions ...callopt.Option) (r *link.SeeLinkRankingResponse, err error)
	SeeUserLink(ctx context.Context, req *link.SeeUserLinkRequest, callOptions ...callopt.Option) (r *link.SeeUserLinkResponse, err error)
	LinkRedirect(ctx context.Context, req *link.LinkRedirectRequest, callOptions ...callopt.Option) (r *link.LinkRedirectResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kLinkServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kLinkServiceClient struct {
	*kClient
}

func (p *kLinkServiceClient) GenerateLink(ctx context.Context, req *link.GenerateLinkRequest, callOptions ...callopt.Option) (r *link.GenerateLinkResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GenerateLink(ctx, req)
}

func (p *kLinkServiceClient) DeleteLink(ctx context.Context, req *link.DeleteLinkRequest, callOptions ...callopt.Option) (r *link.DeleteLinkResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteLink(ctx, req)
}

func (p *kLinkServiceClient) ChangeLink(ctx context.Context, req *link.ChangeLinkRequest, callOptions ...callopt.Option) (r *link.ChangeLinkResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeLink(ctx, req)
}

func (p *kLinkServiceClient) SeeLinkRanking(ctx context.Context, req *link.SeeLinkRankingRequest, callOptions ...callopt.Option) (r *link.SeeLinkRankingResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SeeLinkRanking(ctx, req)
}

func (p *kLinkServiceClient) SeeUserLink(ctx context.Context, req *link.SeeUserLinkRequest, callOptions ...callopt.Option) (r *link.SeeUserLinkResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SeeUserLink(ctx, req)
}

func (p *kLinkServiceClient) LinkRedirect(ctx context.Context, req *link.LinkRedirectRequest, callOptions ...callopt.Option) (r *link.LinkRedirectResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LinkRedirect(ctx, req)
}
