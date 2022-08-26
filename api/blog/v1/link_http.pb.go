// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/third_party/pagination"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type LinkServiceHTTPServer interface {
	CreateLink(context.Context, *CreateLinkRequest) (*Link, error)
	DeleteLink(context.Context, *DeleteLinkRequest) (*emptypb.Empty, error)
	GetLink(context.Context, *GetLinkRequest) (*Link, error)
	ListLink(context.Context, *pagination.PagingRequest) (*ListLinkResponse, error)
	UpdateLink(context.Context, *UpdateLinkRequest) (*Link, error)
}

func RegisterLinkServiceHTTPServer(s *http.Server, srv LinkServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/links", _LinkService_ListLink0_HTTP_Handler(srv))
	r.GET("/blog/v1/links/{id}", _LinkService_GetLink0_HTTP_Handler(srv))
	r.POST("/blog/v1/links", _LinkService_CreateLink0_HTTP_Handler(srv))
	r.PUT("/blog/v1/links/{id}", _LinkService_UpdateLink0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/links/{id}", _LinkService_DeleteLink0_HTTP_Handler(srv))
}

func _LinkService_ListLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.LinkService/ListLink")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListLink(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListLinkResponse)
		return ctx.Result(200, reply)
	}
}

func _LinkService_GetLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.LinkService/GetLink")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLink(ctx, req.(*GetLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_CreateLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.LinkService/CreateLink")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateLink(ctx, req.(*CreateLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_UpdateLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateLinkRequest
		if err := ctx.Bind(&in.Link); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.LinkService/UpdateLink")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateLink(ctx, req.(*UpdateLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_DeleteLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/blog.v1.LinkService/DeleteLink")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteLink(ctx, req.(*DeleteLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type LinkServiceHTTPClient interface {
	CreateLink(ctx context.Context, req *CreateLinkRequest, opts ...http.CallOption) (rsp *Link, err error)
	DeleteLink(ctx context.Context, req *DeleteLinkRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetLink(ctx context.Context, req *GetLinkRequest, opts ...http.CallOption) (rsp *Link, err error)
	ListLink(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *ListLinkResponse, err error)
	UpdateLink(ctx context.Context, req *UpdateLinkRequest, opts ...http.CallOption) (rsp *Link, err error)
}

type LinkServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLinkServiceHTTPClient(client *http.Client) LinkServiceHTTPClient {
	return &LinkServiceHTTPClientImpl{client}
}

func (c *LinkServiceHTTPClientImpl) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...http.CallOption) (*Link, error) {
	var out Link
	pattern := "/blog/v1/links"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/blog.v1.LinkService/CreateLink"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/blog.v1.LinkService/DeleteLink"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) GetLink(ctx context.Context, in *GetLinkRequest, opts ...http.CallOption) (*Link, error) {
	var out Link
	pattern := "/blog/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/blog.v1.LinkService/GetLink"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) ListLink(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*ListLinkResponse, error) {
	var out ListLinkResponse
	pattern := "/blog/v1/links"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/blog.v1.LinkService/ListLink"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...http.CallOption) (*Link, error) {
	var out Link
	pattern := "/blog/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/blog.v1.LinkService/UpdateLink"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Link, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
