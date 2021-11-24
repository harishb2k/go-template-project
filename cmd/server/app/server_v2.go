package app

import (
	"github.com/gin-gonic/gin"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"golang.org/x/net/context"
	"net/http"
)

type ServerImpl struct {
	fx.In
	server.Server
	RandomApiHandler http.HandlerFunc `name:"RandomApiHandler"`
	UserHandler      *handler.UserHandler
}

func (s *ServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.GetRouter().ServeHTTP(w, r)
}

func (s *ServerImpl) routes() {
	serviceName := "srv"
	publicRouter := s.GetRouter().Group(serviceName)
	internalRouter := s.GetRouter().Group(serviceName + "/internal")

	// We must add the MW with all the routers
	s.GetRouter().Use(ginContextToContextMiddleware())
	publicRouter.Use(ginContextToContextMiddleware())
	internalRouter.Use(ginContextToContextMiddleware())

	v1Apis := publicRouter.Group("/v1")
	{
		usersApi := v1Apis.Group("users")
		usersApi.POST("", s.handleAddUser())
		usersApi.GET("", s.handleGetUser())
	}
	{
		randomApi := v1Apis.Group("random")
		randomApi.GET("", s.handleRandomApi())
	}
}

// internal MW to add
func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *ServerImpl) handleAddUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		server.EnsureGinContextWrapper(s.UserHandler.Adduser()).ServeHTTP(c.Writer, c.Request)
	}
}

func (s *ServerImpl) handleGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		server.EnsureGinContextWrapper(s.UserHandler.GetUser()).ServeHTTP(c.Writer, c.Request)
	}
}

func (s *ServerImpl) handleRandomApi() func(c *gin.Context) {
	return func(c *gin.Context) {
		server.EnsureGinContextWrapper(s.RandomApiHandler).ServeHTTP(c.Writer, c.Request)
	}
}
