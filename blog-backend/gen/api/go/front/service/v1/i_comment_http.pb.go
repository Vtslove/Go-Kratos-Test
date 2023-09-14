// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             (unknown)
// source: front/service/v1/i_comment.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-cms/gen/api/go/comment/service/v1"
	pagination "kratos-cms/gen/api/go/common/pagination"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCommentServiceCreateComment = "/front.service.v1.CommentService/CreateComment"
const OperationCommentServiceDeleteComment = "/front.service.v1.CommentService/DeleteComment"
const OperationCommentServiceGetComment = "/front.service.v1.CommentService/GetComment"
const OperationCommentServiceListComment = "/front.service.v1.CommentService/ListComment"
const OperationCommentServiceUpdateComment = "/front.service.v1.CommentService/UpdateComment"

type CommentServiceHTTPServer interface {
	// CreateComment 创建评论
	CreateComment(context.Context, *v1.CreateCommentRequest) (*v1.Comment, error)
	// DeleteComment 删除评论
	DeleteComment(context.Context, *v1.DeleteCommentRequest) (*emptypb.Empty, error)
	// GetComment 获取评论数据
	GetComment(context.Context, *v1.GetCommentRequest) (*v1.Comment, error)
	// ListComment 获取评论列表
	ListComment(context.Context, *pagination.PagingRequest) (*v1.ListCommentResponse, error)
	// UpdateComment 更新评论
	UpdateComment(context.Context, *v1.UpdateCommentRequest) (*v1.Comment, error)
}

func RegisterCommentServiceHTTPServer(s *http.Server, srv CommentServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/comments", _CommentService_ListComment0_HTTP_Handler(srv))
	r.GET("/blog/v1/comments/{id}", _CommentService_GetComment0_HTTP_Handler(srv))
	r.POST("/blog/v1/comments", _CommentService_CreateComment0_HTTP_Handler(srv))
	r.PUT("/blog/v1/comments/{id}", _CommentService_UpdateComment0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/comments/{id}", _CommentService_DeleteComment0_HTTP_Handler(srv))
}

func _CommentService_ListComment0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceListComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListComment(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListCommentResponse)
		return ctx.Result(200, reply)
	}
}

func _CommentService_GetComment0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceGetComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetComment(ctx, req.(*v1.GetCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Comment)
		return ctx.Result(200, reply)
	}
}

func _CommentService_CreateComment0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceCreateComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateComment(ctx, req.(*v1.CreateCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Comment)
		return ctx.Result(200, reply)
	}
}

func _CommentService_UpdateComment0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateCommentRequest
		if err := ctx.Bind(&in.Comment); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceUpdateComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateComment(ctx, req.(*v1.UpdateCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Comment)
		return ctx.Result(200, reply)
	}
}

func _CommentService_DeleteComment0_HTTP_Handler(srv CommentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentServiceDeleteComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*v1.DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type CommentServiceHTTPClient interface {
	CreateComment(ctx context.Context, req *v1.CreateCommentRequest, opts ...http.CallOption) (rsp *v1.Comment, err error)
	DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetComment(ctx context.Context, req *v1.GetCommentRequest, opts ...http.CallOption) (rsp *v1.Comment, err error)
	ListComment(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListCommentResponse, err error)
	UpdateComment(ctx context.Context, req *v1.UpdateCommentRequest, opts ...http.CallOption) (rsp *v1.Comment, err error)
}

type CommentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCommentServiceHTTPClient(client *http.Client) CommentServiceHTTPClient {
	return &CommentServiceHTTPClientImpl{client}
}

func (c *CommentServiceHTTPClientImpl) CreateComment(ctx context.Context, in *v1.CreateCommentRequest, opts ...http.CallOption) (*v1.Comment, error) {
	var out v1.Comment
	pattern := "/blog/v1/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCommentServiceCreateComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentServiceHTTPClientImpl) DeleteComment(ctx context.Context, in *v1.DeleteCommentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/comments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentServiceDeleteComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentServiceHTTPClientImpl) GetComment(ctx context.Context, in *v1.GetCommentRequest, opts ...http.CallOption) (*v1.Comment, error) {
	var out v1.Comment
	pattern := "/blog/v1/comments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentServiceGetComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentServiceHTTPClientImpl) ListComment(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListCommentResponse, error) {
	var out v1.ListCommentResponse
	pattern := "/blog/v1/comments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentServiceListComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentServiceHTTPClientImpl) UpdateComment(ctx context.Context, in *v1.UpdateCommentRequest, opts ...http.CallOption) (*v1.Comment, error) {
	var out v1.Comment
	pattern := "/blog/v1/comments/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCommentServiceUpdateComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Comment, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
