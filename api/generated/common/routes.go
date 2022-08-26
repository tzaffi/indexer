// Package common provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package common

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns 200 if healthy.
	// (GET /health)
	MakeHealthCheck(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// MakeHealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) MakeHealthCheck(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MakeHealthCheck(ctx)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/health", wrapper.MakeHealthCheck, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/8vyScb52AAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
