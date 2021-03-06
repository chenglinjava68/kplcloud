/**
 * @Time : 2019-07-02 17:15
 * @Author : solacowa@gmail.com
 * @File : endpoint
 * @Software: GoLand
 */

package project

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/kplcloud/kplcloud/src/repository/types"
	"github.com/kplcloud/kplcloud/src/util/encode"
)

type getRequest struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type pomFileRequest struct {
	PomFilePath string `json:"pom_file_path"`
}

type gitAddrRequest struct {
	GitAddr string `json:"git_addr"`
}

type postRequest struct {
	getRequest
	DisplayName string `json:"display_name"`
	Desc        string `json:"desc"`
}

type basicRequest struct {
	types.TemplateField
}

type listRequest struct {
	Name  string
	Group int64
	Page  int
	Limit int
}

type deleteRequest struct {
	getRequest
	Code string `json:"code"`
}

func makePostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postRequest)
		err = s.Post(ctx, req.Namespace, req.Name, req.DisplayName, req.DisplayName)
		return encode.Response{Err: err}, err
	}
}

func makeBasicPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(basicRequest)
		err = s.BasicPost(ctx, req.Name, req)
		return encode.Response{Err: err}, err
	}
}

func makeListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listRequest)
		res, err := s.List(ctx, req.Page, req.Limit, req.Name, req.Group)
		return encode.Response{Err: err, Data: res}, err
	}
}

func makeListByNsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := s.ListByNs(ctx)
		return encode.Response{Err: err, Data: res}, err
	}
}

func makePomFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pomFileRequest)
		err := s.PomFile(ctx, req.PomFilePath)
		return encode.Response{Err: err}, err
	}
}

func makeGitAddrEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(gitAddrRequest)
		err := s.GitAddr(ctx, req.GitAddr)
		return encode.Response{Err: err}, err
	}
}

func makeDetailEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := s.Detail(ctx)
		return encode.Response{Err: err, Data: res}, err
	}
}

func makeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRequest)
		err := s.Update(ctx, req.DisplayName, req.Desc)
		return encode.Response{Err: err}, err
	}
}

func makeWorkspaceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := s.Workspace(ctx)
		return encode.Response{Err: err, Data: res}, err
	}
}
func makeSyncEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.Sync(ctx)
		return encode.Response{Err: err}, err
	}
}

func makeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteRequest)
		err = s.Delete(ctx, req.Namespace, req.Name, req.Code)
		return encode.Response{Err: err}, err
	}
}

func makeConfigEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, err := s.Config(ctx)
		return encode.Response{Err: err, Data: data}, err
	}
}
