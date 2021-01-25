// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/craftslab/metalflow/auth"
	"github.com/craftslab/metalflow/config"
	"github.com/craftslab/metalflow/controller"
	"github.com/craftslab/metalflow/util"
)

const (
	timeout = 5 * time.Second
)

type Router struct {
	auth   *auth.Auth
	config *config.Config
	engine *gin.Engine
}

func New() *Router {
	return &Router{}
}

func (r *Router) Run(addr string, cfg *config.Config) error {
	r.config = cfg

	if err := r.initAuth(); err != nil {
		return errors.Wrap(err, "failed to init auth")
	}

	if err := r.initRoute(); err != nil {
		return errors.Wrap(err, "failed to init route")
	}

	if err := r.setRoute(); err != nil {
		return errors.Wrap(err, "failed to set route")
	}

	return r.runRouter(addr)
}

func (r *Router) initAuth() error {
	r.auth = auth.New()
	if r.auth == nil {
		return errors.New("failed to new Auth")
	}

	if err := r.auth.Init(); err != nil {
		return errors.Wrap(err, "failed to init Auth")
	}

	return nil
}

func (r *Router) initRoute() error {
	r.engine = gin.New()
	if r.engine == nil {
		return errors.New("failed to new gin")
	}

	r.engine.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"DELETE", "GET", "PATCH", "POST", "PUT"},
		AllowOrigins:     []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		ExposeHeaders: []string{"Content-Type"},
		MaxAge:        24 * time.Hour,
	}))

	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())

	return nil
}

func (r *Router) setRoute() error {
	ctrl := controller.New()
	if ctrl == nil {
		return errors.New("failed to new controller")
	}

	_auth := r.engine.Group("/auth")
	_auth.POST("login", r.auth.Middleware.LoginHandler)
	_auth.GET("refresh", r.auth.Middleware.RefreshHandler)

	accounts := r.engine.Group("/accounts")
	accounts.Use(r.auth.Middleware.MiddlewareFunc())
	accounts.GET(":id", ctrl.GetAccount)
	accounts.GET("/", ctrl.QueryAccount)

	cfg := r.engine.Group("/config")
	cfg.Use(r.auth.Middleware.MiddlewareFunc())
	cfg.GET("server/version", ctrl.GetServerVersion)

	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.engine.NoRoute(r.auth.Middleware.MiddlewareFunc(), func(ctx *gin.Context) {
		util.NewError(ctx, http.StatusNotFound, errors.New("Page not found"))
	})

	return nil
}

func (r Router) runRouter(addr string) error {
	srv := &http.Server{
		Addr:           addr,
		Handler:        r.engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen and serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shutdown")
	}

	<-ctx.Done()

	return nil
}
