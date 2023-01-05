// Package interfaces provides primitives to interact with the openapi HTTP API.
//
// Code generated by  version  DO NOT EDIT.
package interfaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 批量查询渠道信息
	// (GET /)
	Get(w http.ResponseWriter, r *http.Request, params GetParams)
	// 创建渠道
	// (POST /)
	Post(w http.ResponseWriter, r *http.Request)
	// 查询渠道属性
	// (GET /attributes)
	GetAttributes(w http.ResponseWriter, r *http.Request)
	// 创建渠道属性记录
	// (POST /attributes)
	PostAttributes(w http.ResponseWriter, r *http.Request)
	// 批量更新渠道属性
	// (PUT /attributes)
	PutAttributes(w http.ResponseWriter, r *http.Request)
	// 设置报价在国家级别是否生效
	// (PUT /costs)
	PutCosts(w http.ResponseWriter, r *http.Request)
	// 下载报价模版
	// (GET /costs/download-templates/{template_name})
	GetCostsDownloadTemplatesTemplateName(w http.ResponseWriter, r *http.Request, templateName string)
	// 上传报价信息
	// (POST /costs/{channel_cost_batch_id}/upload)
	PostCostsChannelCostBatchIdUpload(w http.ResponseWriter, r *http.Request, channelCostBatchId int64)
	// 获取客户运费配置
	// (GET /customer-configs)
	GetCustomerConfigs(w http.ResponseWriter, r *http.Request, params GetCustomerConfigsParams)
	// 创建客户运费配置
	// (POST /customer-configs)
	PostCustomerConfigs(w http.ResponseWriter, r *http.Request)
	// 更新客户运费配置
	// (PUT /customer-configs)
	PutCustomerConfigs(w http.ResponseWriter, r *http.Request)
	// 查询推荐渠道设置
	// (GET /recommends)
	GetRecommends(w http.ResponseWriter, r *http.Request, params GetRecommendsParams)
	// 批量查询渠道批次数据
	// (GET /{channel_id}/cost-batches)
	GetChannelIdCostBatches(w http.ResponseWriter, r *http.Request, channelId int64, params GetChannelIdCostBatchesParams)
	// 创建渠道批次
	// (POST /{channel_id}/cost-batches)
	PostChannelIdCostBatches(w http.ResponseWriter, r *http.Request, channelId int64)
	// 更新批次信息
	// (PUT /{channel_id}/cost-batches/{id})
	PutChannelIdCostBatchesId(w http.ResponseWriter, r *http.Request, channelId int64, id int64)
	// 批量查询报价信息
	// (GET /{channel_id}/costs/{channel_cost_batch_id})
	GetChannelIdCostsChannelCostBatchId(w http.ResponseWriter, r *http.Request, channelId int64, channelCostBatchId int64, params GetChannelIdCostsChannelCostBatchIdParams)
	// 更新渠道信息
	// (PUT /{id})
	PutId(w http.ResponseWriter, r *http.Request, id int64)
	// 创建推荐渠道设置
	// (POST /{id}/recommends)
	PostIdRecommends(w http.ResponseWriter, r *http.Request, id int64)
	// 更新推荐渠道推荐设置
	// (PUT /{id}/recommends)
	PutIdRecommends(w http.ResponseWriter, r *http.Request, id int64)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Get operation middleware
func (siw *ServerInterfaceWrapper) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetParams

	// ------------- Optional query parameter "channel_platform" -------------
	if paramValue := r.URL.Query().Get("channel_platform"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "channel_platform", r.URL.Query(), &params.ChannelPlatform)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_platform", Err: err})
		return
	}

	// ------------- Optional query parameter "channel_code" -------------
	if paramValue := r.URL.Query().Get("channel_code"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "channel_code", r.URL.Query(), &params.ChannelCode)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_code", Err: err})
		return
	}

	// ------------- Optional query parameter "channel_name" -------------
	if paramValue := r.URL.Query().Get("channel_name"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "channel_name", r.URL.Query(), &params.ChannelName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_name", Err: err})
		return
	}

	// ------------- Optional query parameter "sorter" -------------
	if paramValue := r.URL.Query().Get("sorter"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sorter", r.URL.Query(), &params.Sorter)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sorter", Err: err})
		return
	}

	// ------------- Optional query parameter "current" -------------
	if paramValue := r.URL.Query().Get("current"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "current", r.URL.Query(), &params.Current)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "current", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Get(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Post operation middleware
func (siw *ServerInterfaceWrapper) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Post(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAttributes operation middleware
func (siw *ServerInterfaceWrapper) GetAttributes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAttributes(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostAttributes operation middleware
func (siw *ServerInterfaceWrapper) PostAttributes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAttributes(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutAttributes operation middleware
func (siw *ServerInterfaceWrapper) PutAttributes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutAttributes(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutCosts operation middleware
func (siw *ServerInterfaceWrapper) PutCosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutCosts(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetCostsDownloadTemplatesTemplateName operation middleware
func (siw *ServerInterfaceWrapper) GetCostsDownloadTemplatesTemplateName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "template_name" -------------
	var templateName string

	err = runtime.BindStyledParameter("simple", false, "template_name", chi.URLParam(r, "template_name"), &templateName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "template_name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCostsDownloadTemplatesTemplateName(w, r, templateName)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostCostsChannelCostBatchIdUpload operation middleware
func (siw *ServerInterfaceWrapper) PostCostsChannelCostBatchIdUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "channel_cost_batch_id" -------------
	var channelCostBatchId int64

	err = runtime.BindStyledParameter("simple", false, "channel_cost_batch_id", chi.URLParam(r, "channel_cost_batch_id"), &channelCostBatchId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_cost_batch_id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostCostsChannelCostBatchIdUpload(w, r, channelCostBatchId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetCustomerConfigs operation middleware
func (siw *ServerInterfaceWrapper) GetCustomerConfigs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCustomerConfigsParams

	// ------------- Optional query parameter "channel_id" -------------
	if paramValue := r.URL.Query().Get("channel_id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "channel_id", r.URL.Query(), &params.ChannelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_id", Err: err})
		return
	}

	// ------------- Optional query parameter "sorter" -------------
	if paramValue := r.URL.Query().Get("sorter"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sorter", r.URL.Query(), &params.Sorter)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sorter", Err: err})
		return
	}

	// ------------- Optional query parameter "status" -------------
	if paramValue := r.URL.Query().Get("status"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "status", Err: err})
		return
	}

	// ------------- Optional query parameter "current" -------------
	if paramValue := r.URL.Query().Get("current"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "current", r.URL.Query(), &params.Current)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "current", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCustomerConfigs(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostCustomerConfigs operation middleware
func (siw *ServerInterfaceWrapper) PostCustomerConfigs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostCustomerConfigs(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutCustomerConfigs operation middleware
func (siw *ServerInterfaceWrapper) PutCustomerConfigs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutCustomerConfigs(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetRecommends operation middleware
func (siw *ServerInterfaceWrapper) GetRecommends(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRecommendsParams

	// ------------- Optional query parameter "country_code" -------------
	if paramValue := r.URL.Query().Get("country_code"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "country_code", r.URL.Query(), &params.CountryCode)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "country_code", Err: err})
		return
	}

	// ------------- Optional query parameter "channel_id" -------------
	if paramValue := r.URL.Query().Get("channel_id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "channel_id", r.URL.Query(), &params.ChannelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_id", Err: err})
		return
	}

	// ------------- Optional query parameter "is_recommended" -------------
	if paramValue := r.URL.Query().Get("is_recommended"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "is_recommended", r.URL.Query(), &params.IsRecommended)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "is_recommended", Err: err})
		return
	}

	// ------------- Optional query parameter "status" -------------
	if paramValue := r.URL.Query().Get("status"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "status", Err: err})
		return
	}

	// ------------- Optional query parameter "current" -------------
	if paramValue := r.URL.Query().Get("current"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "current", r.URL.Query(), &params.Current)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "current", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	// ------------- Optional query parameter "sorter" -------------
	if paramValue := r.URL.Query().Get("sorter"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sorter", r.URL.Query(), &params.Sorter)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sorter", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRecommends(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetChannelIdCostBatches operation middleware
func (siw *ServerInterfaceWrapper) GetChannelIdCostBatches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "channel_id" -------------
	var channelId int64

	err = runtime.BindStyledParameter("simple", false, "channel_id", chi.URLParam(r, "channel_id"), &channelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetChannelIdCostBatchesParams

	// ------------- Optional query parameter "id" -------------
	if paramValue := r.URL.Query().Get("id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "id", r.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	// ------------- Optional query parameter "effective_date" -------------
	if paramValue := r.URL.Query().Get("effective_date"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "effective_date", r.URL.Query(), &params.EffectiveDate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "effective_date", Err: err})
		return
	}

	// ------------- Optional query parameter "status" -------------
	if paramValue := r.URL.Query().Get("status"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "status", Err: err})
		return
	}

	// ------------- Optional query parameter "sorter" -------------
	if paramValue := r.URL.Query().Get("sorter"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sorter", r.URL.Query(), &params.Sorter)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sorter", Err: err})
		return
	}

	// ------------- Optional query parameter "current" -------------
	if paramValue := r.URL.Query().Get("current"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "current", r.URL.Query(), &params.Current)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "current", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetChannelIdCostBatches(w, r, channelId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostChannelIdCostBatches operation middleware
func (siw *ServerInterfaceWrapper) PostChannelIdCostBatches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "channel_id" -------------
	var channelId int64

	err = runtime.BindStyledParameter("simple", false, "channel_id", chi.URLParam(r, "channel_id"), &channelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostChannelIdCostBatches(w, r, channelId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutChannelIdCostBatchesId operation middleware
func (siw *ServerInterfaceWrapper) PutChannelIdCostBatchesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "channel_id" -------------
	var channelId int64

	err = runtime.BindStyledParameter("simple", false, "channel_id", chi.URLParam(r, "channel_id"), &channelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_id", Err: err})
		return
	}

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutChannelIdCostBatchesId(w, r, channelId, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetChannelIdCostsChannelCostBatchId operation middleware
func (siw *ServerInterfaceWrapper) GetChannelIdCostsChannelCostBatchId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "channel_id" -------------
	var channelId int64

	err = runtime.BindStyledParameter("simple", false, "channel_id", chi.URLParam(r, "channel_id"), &channelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_id", Err: err})
		return
	}

	// ------------- Path parameter "channel_cost_batch_id" -------------
	var channelCostBatchId int64

	err = runtime.BindStyledParameter("simple", false, "channel_cost_batch_id", chi.URLParam(r, "channel_cost_batch_id"), &channelCostBatchId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "channel_cost_batch_id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetChannelIdCostsChannelCostBatchIdParams

	// ------------- Optional query parameter "country_code" -------------
	if paramValue := r.URL.Query().Get("country_code"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "country_code", r.URL.Query(), &params.CountryCode)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "country_code", Err: err})
		return
	}

	// ------------- Optional query parameter "sorter" -------------
	if paramValue := r.URL.Query().Get("sorter"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sorter", r.URL.Query(), &params.Sorter)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sorter", Err: err})
		return
	}

	// ------------- Optional query parameter "current" -------------
	if paramValue := r.URL.Query().Get("current"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "current", r.URL.Query(), &params.Current)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "current", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	// ------------- Optional query parameter "status" -------------
	if paramValue := r.URL.Query().Get("status"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "status", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetChannelIdCostsChannelCostBatchId(w, r, channelId, channelCostBatchId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutId operation middleware
func (siw *ServerInterfaceWrapper) PutId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostIdRecommends operation middleware
func (siw *ServerInterfaceWrapper) PostIdRecommends(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostIdRecommends(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutIdRecommends operation middleware
func (siw *ServerInterfaceWrapper) PutIdRecommends(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Jwt_header_AuthorizationScopes, []string{"api:read", "api:write"})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutIdRecommends(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/", wrapper.Get)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/", wrapper.Post)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/attributes", wrapper.GetAttributes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/attributes", wrapper.PostAttributes)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/attributes", wrapper.PutAttributes)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/costs", wrapper.PutCosts)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/costs/download-templates/{template_name}", wrapper.GetCostsDownloadTemplatesTemplateName)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/costs/{channel_cost_batch_id}/upload", wrapper.PostCostsChannelCostBatchIdUpload)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/customer-configs", wrapper.GetCustomerConfigs)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer-configs", wrapper.PostCustomerConfigs)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/customer-configs", wrapper.PutCustomerConfigs)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/recommends", wrapper.GetRecommends)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/{channel_id}/cost-batches", wrapper.GetChannelIdCostBatches)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/{channel_id}/cost-batches", wrapper.PostChannelIdCostBatches)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/{channel_id}/cost-batches/{id}", wrapper.PutChannelIdCostBatchesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/{channel_id}/costs/{channel_cost_batch_id}", wrapper.GetChannelIdCostsChannelCostBatchId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/{id}", wrapper.PutId)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/{id}/recommends", wrapper.PostIdRecommends)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/{id}/recommends", wrapper.PutIdRecommends)
	})

	return r
}
