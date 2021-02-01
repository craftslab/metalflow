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
	"github.com/craftslab/metalflow/controller"
	"github.com/craftslab/metalflow/util"
)

const (
	timeout = 5 * time.Second
)

type Router interface {
	Init() error
	Run() error
}

type Config struct {
	Addr string
}

type router struct {
	auth   auth.Auth
	config *Config
	engine *gin.Engine
}

func New(config *Config) Router {
	return &router{
		auth:   nil,
		config: config,
		engine: nil,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Addr: ":9080",
	}
}

func (r *router) Init() error {
	if err := r.initAuth(); err != nil {
		return errors.Wrap(err, "failed to init auth")
	}

	if err := r.initRoute(); err != nil {
		return errors.Wrap(err, "failed to init route")
	}

	if err := r.setRoute(); err != nil {
		return errors.Wrap(err, "failed to set route")
	}

	return nil
}

func (r *router) initAuth() error {
	r.auth = auth.New(auth.DefaultConfig())
	if r.auth == nil {
		return errors.New("failed to new Auth")
	}

	if err := r.auth.Init(); err != nil {
		return errors.Wrap(err, "failed to init Auth")
	}

	return nil
}

func (r *router) initRoute() error {
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

func (r *router) setRoute() error {
	ctrl := controller.New(controller.DefaultConfig())
	if ctrl == nil {
		return errors.New("failed to new controller")
	}

	au := r.engine.Group("/auth")
	au.POST("login", r.auth.Middleware().LoginHandler)
	au.GET("refresh", r.auth.Middleware().RefreshHandler)

	ac := r.engine.Group("/accounts")
	ac.Use(r.auth.Middleware().MiddlewareFunc())
	ac.GET(":id", ctrl.GetAccount)
	ac.GET("/", ctrl.QueryAccount)

	c := r.engine.Group("/config")
	c.Use(r.auth.Middleware().MiddlewareFunc())
	c.GET("server/version", ctrl.GetServerVersion)

	n := r.engine.Group("/nodes")
	n.Use(r.auth.Middleware().MiddlewareFunc())
	n.GET(":id", ctrl.GetNode)
	n.GET(":id/health", ctrl.GetHealth)
	n.GET(":id/info", ctrl.GetInfo)
	n.GET(":id/perf", ctrl.GetPerf)
	n.GET("/", ctrl.QueryNode)
	n.PUT(":id", ctrl.AddNode)
	n.DELETE(":id", ctrl.DelNode)

	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.engine.NoRoute(r.auth.Middleware().MiddlewareFunc(), func(ctx *gin.Context) {
		util.NewError(ctx, http.StatusNotFound, errors.New("Page not found"))
	})

	return nil
}

func (r *router) Run() error {
	srv := &http.Server{
		Addr:           r.config.Addr,
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
