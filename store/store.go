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

package store

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store interface {
	Open() error
	Close() error

	Migrate(model interface{}) error
}

type Config struct {
	Db   string
	Host string
	Pass string
	Port string
	User string
}

type store struct {
	config   *Config
	database *gorm.DB
}

const (
	dsn = "sslmode=disable TimeZone=Asia/Shanghai"
)

func New(_ context.Context, config *Config) Store {
	return &store{
		config:   config,
		database: nil,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Db:   "metalflow",
		Host: "127.0.0.1",
		Pass: "",
		Port: "5432",
		User: "",
	}
}

func (s *store) Open() error {
	host := "host=" + s.config.Host + " "
	port := "port=" + s.config.Port + " "
	user := "user=" + s.config.User + " "
	pass := "password=" + s.config.Pass + " "
	dbname := "dbname=" + s.config.Db + " "

	db, err := gorm.Open(postgres.Open(host+port+user+pass+dbname+dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to open")
	}

	s.database = db

	return nil
}

func (s *store) Close() error {
	// TODO
	return nil
}

func (s *store) Migrate(model interface{}) error {
	if err := s.database.AutoMigrate(model); err != nil {
		return errors.Wrap(err, "failed to migrate")
	}

	return nil
}
