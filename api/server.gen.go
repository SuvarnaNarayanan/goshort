// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.5.0 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

const (
	ApiKeyAuthScopes  = "ApiKeyAuth.Scopes"
	Google_authScopes = "google_auth.Scopes"
)

// PostShortJSONBody defines parameters for PostShort.
type PostShortJSONBody struct {
	Campaign *string `json:"campaign,omitempty"`
	Url      *string `json:"url,omitempty"`
}

// PostShortCodeReportJSONBody defines parameters for PostShortCodeReport.
type PostShortCodeReportJSONBody struct {
	Explanation *string `json:"explanation,omitempty"`
}

// PostShortJSONRequestBody defines body for PostShort for application/json ContentType.
type PostShortJSONRequestBody PostShortJSONBody

// PostShortCodeReportJSONRequestBody defines body for PostShortCodeReport for application/json ContentType.
type PostShortCodeReportJSONRequestBody PostShortCodeReportJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (OPTIONS /key/{path*})
	OptionsKeyPath(ctx echo.Context, path string) error
	// List all campaigns
	// (GET /list/campaign)
	GetListCampaign(ctx echo.Context) error
	// List all shortened URLs
	// (GET /list/short)
	GetListShort(ctx echo.Context) error
	// Shorten a URL
	// (POST /short)
	PostShort(ctx echo.Context) error
	// Redirect to the original URL
	// (GET /{shortCode})
	GetShortCode(ctx echo.Context, shortCode string) error
	// Get URL statistics
	// (GET /{shortCode}/info)
	GetShortCodeInfo(ctx echo.Context, shortCode string) error
	// Report a URL
	// (POST /{shortCode}/report)
	PostShortCodeReport(ctx echo.Context, shortCode string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// OptionsKeyPath converts echo context to params.
func (w *ServerInterfaceWrapper) OptionsKeyPath(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "path" -------------
	var path string

	err = runtime.BindStyledParameterWithOptions("simple", "path", ctx.Param("path"), &path, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OptionsKeyPath(ctx, path)
	return err
}

// GetListCampaign converts echo context to params.
func (w *ServerInterfaceWrapper) GetListCampaign(ctx echo.Context) error {
	var err error

	ctx.Set(Google_authScopes, []string{"https://www.googleapis.com/auth/userinfo.email"})

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetListCampaign(ctx)
	return err
}

// GetListShort converts echo context to params.
func (w *ServerInterfaceWrapper) GetListShort(ctx echo.Context) error {
	var err error

	ctx.Set(Google_authScopes, []string{"https://www.googleapis.com/auth/userinfo.email"})

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetListShort(ctx)
	return err
}

// PostShort converts echo context to params.
func (w *ServerInterfaceWrapper) PostShort(ctx echo.Context) error {
	var err error

	ctx.Set(Google_authScopes, []string{"https://www.googleapis.com/auth/userinfo.email"})

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostShort(ctx)
	return err
}

// GetShortCode converts echo context to params.
func (w *ServerInterfaceWrapper) GetShortCode(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shortCode" -------------
	var shortCode string

	err = runtime.BindStyledParameterWithOptions("simple", "shortCode", ctx.Param("shortCode"), &shortCode, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shortCode: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetShortCode(ctx, shortCode)
	return err
}

// GetShortCodeInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetShortCodeInfo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shortCode" -------------
	var shortCode string

	err = runtime.BindStyledParameterWithOptions("simple", "shortCode", ctx.Param("shortCode"), &shortCode, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shortCode: %s", err))
	}

	ctx.Set(Google_authScopes, []string{"https://www.googleapis.com/auth/userinfo.email"})

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetShortCodeInfo(ctx, shortCode)
	return err
}

// PostShortCodeReport converts echo context to params.
func (w *ServerInterfaceWrapper) PostShortCodeReport(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shortCode" -------------
	var shortCode string

	err = runtime.BindStyledParameterWithOptions("simple", "shortCode", ctx.Param("shortCode"), &shortCode, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shortCode: %s", err))
	}

	ctx.Set(Google_authScopes, []string{"https://www.googleapis.com/auth/userinfo.email"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostShortCodeReport(ctx, shortCode)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.OPTIONS(baseURL+"/key/:path", wrapper.OptionsKeyPath)
	router.GET(baseURL+"/list/campaign", wrapper.GetListCampaign)
	router.GET(baseURL+"/list/short", wrapper.GetListShort)
	router.POST(baseURL+"/short", wrapper.PostShort)
	router.GET(baseURL+"/:shortCode", wrapper.GetShortCode)
	router.GET(baseURL+"/:shortCode/info", wrapper.GetShortCodeInfo)
	router.POST(baseURL+"/:shortCode/report", wrapper.PostShortCodeReport)

}

type OptionsKeyPathRequestObject struct {
	Path string `json:"path"`
}

type OptionsKeyPathResponseObject interface {
	VisitOptionsKeyPathResponse(w http.ResponseWriter) error
}

type OptionsKeyPath200Response struct {
}

func (response OptionsKeyPath200Response) VisitOptionsKeyPathResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetListCampaignRequestObject struct {
}

type GetListCampaignResponseObject interface {
	VisitGetListCampaignResponse(w http.ResponseWriter) error
}

type GetListCampaign200JSONResponse []struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
}

func (response GetListCampaign200JSONResponse) VisitGetListCampaignResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetListShortRequestObject struct {
}

type GetListShortResponseObject interface {
	VisitGetListShortResponse(w http.ResponseWriter) error
}

type GetListShort200JSONResponse []struct {
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortUrl    *string `json:"short_url,omitempty"`
}

func (response GetListShort200JSONResponse) VisitGetListShortResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostShortRequestObject struct {
	Body *PostShortJSONRequestBody
}

type PostShortResponseObject interface {
	VisitPostShortResponse(w http.ResponseWriter) error
}

type PostShort201JSONResponse struct {
	ShortUrl *string `json:"short_url,omitempty"`
}

func (response PostShort201JSONResponse) VisitPostShortResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostShort400Response struct {
}

func (response PostShort400Response) VisitPostShortResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetShortCodeRequestObject struct {
	ShortCode string `json:"shortCode"`
}

type GetShortCodeResponseObject interface {
	VisitGetShortCodeResponse(w http.ResponseWriter) error
}

type GetShortCode302ResponseHeaders struct {
	Location string
}

type GetShortCode302Response struct {
	Headers GetShortCode302ResponseHeaders
}

func (response GetShortCode302Response) VisitGetShortCodeResponse(w http.ResponseWriter) error {
	w.Header().Set("Location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(302)
	return nil
}

type GetShortCode404Response struct {
}

func (response GetShortCode404Response) VisitGetShortCodeResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetShortCodeInfoRequestObject struct {
	ShortCode string `json:"shortCode"`
}

type GetShortCodeInfoResponseObject interface {
	VisitGetShortCodeInfoResponse(w http.ResponseWriter) error
}

type GetShortCodeInfo200JSONResponse struct {
	Clicks *[]struct {
		IpAddress *string    `json:"ip_address,omitempty"`
		Timestamp *time.Time `json:"timestamp,omitempty"`
	} `json:"clicks,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	OriginalUrl *string    `json:"original_url,omitempty"`
}

func (response GetShortCodeInfo200JSONResponse) VisitGetShortCodeInfoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetShortCodeInfo404Response struct {
}

func (response GetShortCodeInfo404Response) VisitGetShortCodeInfoResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PostShortCodeReportRequestObject struct {
	ShortCode string `json:"shortCode"`
	Body      *PostShortCodeReportJSONRequestBody
}

type PostShortCodeReportResponseObject interface {
	VisitPostShortCodeReportResponse(w http.ResponseWriter) error
}

type PostShortCodeReport200Response struct {
}

func (response PostShortCodeReport200Response) VisitPostShortCodeReportResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostShortCodeReport404Response struct {
}

func (response PostShortCodeReport404Response) VisitPostShortCodeReportResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (OPTIONS /key/{path*})
	OptionsKeyPath(ctx context.Context, request OptionsKeyPathRequestObject) (OptionsKeyPathResponseObject, error)
	// List all campaigns
	// (GET /list/campaign)
	GetListCampaign(ctx context.Context, request GetListCampaignRequestObject) (GetListCampaignResponseObject, error)
	// List all shortened URLs
	// (GET /list/short)
	GetListShort(ctx context.Context, request GetListShortRequestObject) (GetListShortResponseObject, error)
	// Shorten a URL
	// (POST /short)
	PostShort(ctx context.Context, request PostShortRequestObject) (PostShortResponseObject, error)
	// Redirect to the original URL
	// (GET /{shortCode})
	GetShortCode(ctx context.Context, request GetShortCodeRequestObject) (GetShortCodeResponseObject, error)
	// Get URL statistics
	// (GET /{shortCode}/info)
	GetShortCodeInfo(ctx context.Context, request GetShortCodeInfoRequestObject) (GetShortCodeInfoResponseObject, error)
	// Report a URL
	// (POST /{shortCode}/report)
	PostShortCodeReport(ctx context.Context, request PostShortCodeReportRequestObject) (PostShortCodeReportResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// OptionsKeyPath operation middleware
func (sh *strictHandler) OptionsKeyPath(ctx echo.Context, path string) error {
	var request OptionsKeyPathRequestObject

	request.Path = path

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.OptionsKeyPath(ctx.Request().Context(), request.(OptionsKeyPathRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "OptionsKeyPath")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(OptionsKeyPathResponseObject); ok {
		return validResponse.VisitOptionsKeyPathResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetListCampaign operation middleware
func (sh *strictHandler) GetListCampaign(ctx echo.Context) error {
	var request GetListCampaignRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetListCampaign(ctx.Request().Context(), request.(GetListCampaignRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetListCampaign")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetListCampaignResponseObject); ok {
		return validResponse.VisitGetListCampaignResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetListShort operation middleware
func (sh *strictHandler) GetListShort(ctx echo.Context) error {
	var request GetListShortRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetListShort(ctx.Request().Context(), request.(GetListShortRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetListShort")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetListShortResponseObject); ok {
		return validResponse.VisitGetListShortResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostShort operation middleware
func (sh *strictHandler) PostShort(ctx echo.Context) error {
	var request PostShortRequestObject

	var body PostShortJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostShort(ctx.Request().Context(), request.(PostShortRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostShort")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostShortResponseObject); ok {
		return validResponse.VisitPostShortResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetShortCode operation middleware
func (sh *strictHandler) GetShortCode(ctx echo.Context, shortCode string) error {
	var request GetShortCodeRequestObject

	request.ShortCode = shortCode

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetShortCode(ctx.Request().Context(), request.(GetShortCodeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetShortCode")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetShortCodeResponseObject); ok {
		return validResponse.VisitGetShortCodeResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetShortCodeInfo operation middleware
func (sh *strictHandler) GetShortCodeInfo(ctx echo.Context, shortCode string) error {
	var request GetShortCodeInfoRequestObject

	request.ShortCode = shortCode

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetShortCodeInfo(ctx.Request().Context(), request.(GetShortCodeInfoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetShortCodeInfo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetShortCodeInfoResponseObject); ok {
		return validResponse.VisitGetShortCodeInfoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostShortCodeReport operation middleware
func (sh *strictHandler) PostShortCodeReport(ctx echo.Context, shortCode string) error {
	var request PostShortCodeReportRequestObject

	request.ShortCode = shortCode

	var body PostShortCodeReportJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostShortCodeReport(ctx.Request().Context(), request.(PostShortCodeReportRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostShortCodeReport")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostShortCodeReportResponseObject); ok {
		return validResponse.VisitPostShortCodeReportResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
