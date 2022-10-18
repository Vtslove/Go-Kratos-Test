// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.1
// - protoc             v3.19.1
// source: i_link.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-blog/api/content/service/v1"
	pagination "kratos-blog/third_party/pagination"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLinkServiceCreateLink = "/admin.service.v1.LinkService/CreateLink"
const OperationLinkServiceDeleteLink = "/admin.service.v1.LinkService/DeleteLink"
const OperationLinkServiceGetLink = "/admin.service.v1.LinkService/GetLink"
const OperationLinkServiceListLink = "/admin.service.v1.LinkService/ListLink"
const OperationLinkServiceUpdateLink = "/admin.service.v1.LinkService/UpdateLink"

type LinkServiceHTTPServer interface {
	CreateLink(context.Context, *v1.CreateLinkRequest) (*v1.Link, error)
	DeleteLink(context.Context, *v1.DeleteLinkRequest) (*emptypb.Empty, error)
	GetLink(context.Context, *v1.GetLinkRequest) (*v1.Link, error)
	ListLink(context.Context, *pagination.PagingRequest) (*v1.ListLinkResponse, error)
	UpdateLink(context.Context, *v1.UpdateLinkRequest) (*v1.Link, error)
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
		http.SetOperation(ctx, OperationLinkServiceListLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListLink(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListLinkResponse)
		return ctx.Result(200, reply)
	}
}

func _LinkService_GetLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceGetLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLink(ctx, req.(*v1.GetLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_CreateLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceCreateLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateLink(ctx, req.(*v1.CreateLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_UpdateLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateLinkRequest
		if err := ctx.Bind(&in.Link); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceUpdateLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateLink(ctx, req.(*v1.UpdateLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_DeleteLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceDeleteLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteLink(ctx, req.(*v1.DeleteLinkRequest))
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
	CreateLink(ctx context.Context, req *v1.CreateLinkRequest, opts ...http.CallOption) (rsp *v1.Link, err error)
	DeleteLink(ctx context.Context, req *v1.DeleteLinkRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetLink(ctx context.Context, req *v1.GetLinkRequest, opts ...http.CallOption) (rsp *v1.Link, err error)
	ListLink(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListLinkResponse, err error)
	UpdateLink(ctx context.Context, req *v1.UpdateLinkRequest, opts ...http.CallOption) (rsp *v1.Link, err error)
}

type LinkServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLinkServiceHTTPClient(client *http.Client) LinkServiceHTTPClient {
	return &LinkServiceHTTPClientImpl{client}
}

func (c *LinkServiceHTTPClientImpl) CreateLink(ctx context.Context, in *v1.CreateLinkRequest, opts ...http.CallOption) (*v1.Link, error) {
	var out v1.Link
	pattern := "/blog/v1/links"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLinkServiceCreateLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) DeleteLink(ctx context.Context, in *v1.DeleteLinkRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkServiceDeleteLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) GetLink(ctx context.Context, in *v1.GetLinkRequest, opts ...http.CallOption) (*v1.Link, error) {
	var out v1.Link
	pattern := "/blog/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkServiceGetLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) ListLink(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListLinkResponse, error) {
	var out v1.ListLinkResponse
	pattern := "/blog/v1/links"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkServiceListLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) UpdateLink(ctx context.Context, in *v1.UpdateLinkRequest, opts ...http.CallOption) (*v1.Link, error) {
	var out v1.Link
	pattern := "/blog/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLinkServiceUpdateLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Link, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
