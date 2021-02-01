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
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Address  string
	Asset    string
	Comments string
	Health   string
	Id       int
	Info     string
	Perf     string
	Region   string
}

const (
	pass = "postgres"
	user = "postgres"
)

func TestStore(t *testing.T) {
	config := DefaultConfig()
	config.User = user
	config.Pass = pass

	s := New(context.Background(), config)
	assert.NotEqual(t, nil, s)

	err := s.Open()
	assert.Equal(t, nil, err)

	err = s.Migrate(&Model{})
	assert.Equal(t, nil, err)

	err = s.Close()
	assert.Equal(t, nil, err)
}
