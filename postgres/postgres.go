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

package postgres

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres interface {
	Open() error
	Close()

	Migrate(model interface{}) error
	Create(model interface{})
	Read(model, cond, value interface{})
	Update(model interface{}, column string, value interface{})
	Delete(model, cond, value interface{})
}

type Config struct {
	Db   string
	Host string
	Pass string
	Port string
	User string
}

type _postgres struct {
	config   *Config
	database *gorm.DB
}

const (
	dsn = "sslmode=disable TimeZone=Asia/Shanghai"
)

func New(_ context.Context, config *Config) Postgres {
	return &_postgres{
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

func (p *_postgres) Open() error {
	host := "host=" + p.config.Host + " "
	port := "port=" + p.config.Port + " "
	user := "user=" + p.config.User + " "
	pass := "password=" + p.config.Pass + " "
	dbname := "dbname=" + p.config.Db + " "

	db, err := gorm.Open(postgres.Open(host+port+user+pass+dbname+dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to open")
	}

	p.database = db

	return nil
}

func (p *_postgres) Close() {
	// PASS
}

func (p *_postgres) Migrate(model interface{}) error {
	if err := p.database.AutoMigrate(model); err != nil {
		return errors.Wrap(err, "failed to migrate")
	}

	return nil
}

func (p *_postgres) Create(model interface{}) {
	p.database.Create(model)
}

func (p *_postgres) Read(model, cond, value interface{}) {
	p.database.First(model, cond, value)
}

func (p *_postgres) Update(model interface{}, column string, value interface{}) {
	p.database.Model(model).Update(column, value)
}

func (p *_postgres) Delete(model, cond, value interface{}) {
	p.database.Delete(model, cond, value)
}
