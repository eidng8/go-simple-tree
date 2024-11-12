package main

import (
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-simple-tree/ent"
)

type Server struct {
	EC *ent.Client
}

var _ StrictServerInterface = (*Server)(nil)

func newServer(entClient *ent.Client) Server {
	return Server{EC: entClient}
}

func newEngine(mode string, entClient *ent.Client) (*gin.Engine, error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil
	switch mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	server := newServer(entClient)
	handler := NewStrictHandler(server, []StrictMiddlewareFunc{})
	RegisterHandlersWithOptions(
		engine, handler, GinServerOptions{
			// ErrorHandler: func(ctx *gin.Context, err error, code int) {
			// // This doesn't work since the error is generated by fmt.Errorf().
			// // Such error cannot be converted to err.ValidationError.
			// 	if ent.IsValidationError(err) {
			// 		code = http.StatusUnprocessableEntity
			// 	}
			// 	ctx.JSON(code, gin.H{"error": err.Error()})
			// },
		},
	)
	return engine, nil
}
