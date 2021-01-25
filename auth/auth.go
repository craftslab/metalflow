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

package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/craftslab/metalflow/model"
	"github.com/craftslab/metalflow/util"
)

const (
	identityKey = "id"
	key         = "metalflow"
	realm       = "metalflow"
	maxRefresh  = time.Hour
	timeout     = time.Hour
)

type Auth struct {
	Middleware *jwt.GinJWTMiddleware
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	username string
}

func New() *Auth {
	return &Auth{}
}

func (a *Auth) Init() error {
	var err error

	a.Middleware, err = jwt.New(&jwt.GinJWTMiddleware{
		IdentityKey: identityKey,
		Key:         []byte(key),
		MaxRefresh:  maxRefresh,
		Realm:       realm,
		Timeout:     timeout,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var l Login
			if e := c.ShouldBind(&l); e != nil {
				return "", jwt.ErrMissingLoginValues
			}
			a, e := model.QueryAccount(l.Username)
			if e != nil {
				return "", jwt.ErrFailedAuthentication
			}
			if a.Username == l.Username && a.Password == l.Password {
				return &User{
					username: l.Username,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.username == "admin" {
				return true
			}
			return false
		},
		IdentityHandler: func(ctx *gin.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)
			return &User{
				username: claims[identityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.username,
				}
			}
			return jwt.MapClaims{}
		},
		TimeFunc:      time.Now,
		TokenHeadName: "Bearer",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		Unauthorized: func(ctx *gin.Context, code int, msg string) {
			util.NewError(ctx, code, errors.New(msg))
		},
	})

	if err != nil {
		return errors.Wrap(err, "failed to new jwt")
	}

	if err = a.Middleware.MiddlewareInit(); err != nil {
		return errors.Wrap(err, "failed to init middleware")
	}

	return nil
}
