package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc gin.HandlerFunc
}

type ApiGroup struct {
	Routes []Route
	Prefix string
}

type Response[T any] struct {
	Code   uint   `json:"code"`
	Msg    string `json:"msg"`
	Result T      `json:"result"`
}

func GenerateSuccessResponse[T any](entity T) Response[T] {
	return Response[T]{Result: entity}
}

type HandlerFunc[T any] func(ctx *gin.Context) (T, error)

func (apis ApiGroup) RegisterToEngine(engine *gin.Engine) {
	g := engine.Group(apis.Prefix)
	for _, r := range apis.Routes {
		switch r.Method {
		case http.MethodGet:
			g.GET(r.Path, r.HandlerFunc)
		case http.MethodPost:
			g.POST(r.Path, r.HandlerFunc)
		case http.MethodPut:
			g.PUT(r.Path, r.HandlerFunc)
		case http.MethodDelete:
			g.DELETE(r.Path, r.HandlerFunc)
		case http.MethodHead:
			g.HEAD(r.Path, r.HandlerFunc)
		case http.MethodOptions:
			g.OPTIONS(r.Path, r.HandlerFunc)
		case http.MethodPatch:
			g.PATCH(r.Path, r.HandlerFunc)
		default:
			log.Panicf("method %s not allowed", r.Method)
		}
	}
}
