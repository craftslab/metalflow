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
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Address  string `gorm:"uniqueIndex,sort:desc"`
	Asset    string `gorm:"uniqueIndex"`
	Comments string
	Health   string `gorm:"index"`
	Info     string
	Perf     string `gorm:"index"`
	Region   string `gorm:"index"`
}

const (
	pass = "postgres"
	user = "postgres"
)

var (
	model = Model{
		Address:  "127.0.0.1",
		Asset:    "0",
		Comments: "node 0",
		Health:   "running",
		Info: `"
			{
				"bare": {
					"cpu": "4 CPU",
					"disk": "49.0 GB (16.0 GB Used)",
					"io": "RD 11887928 KB WR 61067948 KB",
					"ip": "127.0.0.1",
					"kernel": "5.4.0-58-generic",
					"mac": "00:01:02:03:04:05",
					"network": "RX packets 8974179 TX packets 3124096",
					"os": "Ubuntu 18.04.5 LTS",
					"ram": "7692 MB (1345 MB Used)",
					"system": ""
				}
			}
		"`,
		Perf:   "High",
		Region: "Shanghai",
	}
)

func TestPostgres(t *testing.T) {
	config := DefaultConfig()
	config.User = user
	config.Pass = pass

	p := New(context.Background(), config)
	assert.NotEqual(t, nil, p)

	err := p.Open()
	assert.Equal(t, nil, err)

	err = p.Migrate(&Model{})
	assert.Equal(t, nil, err)

	p.Create(&model)

	m := Model{}
	p.Read(&m, "Address=?", "127.0.0.2")
	assert.NotEqual(t, "127.0.0.1", m.Address)

	m = Model{}
	p.Read(&m, "Address=?", "127.0.0.1")
	assert.Equal(t, "127.0.0.1", m.Address)

	p.Update(&m, "Address", "127.0.0.2")

	m = Model{}
	p.Read(&m, "Address=?", "127.0.0.2")
	assert.Equal(t, "127.0.0.2", m.Address)

	p.Delete(&m, "Address=?", "127.0.0.2")

	m = Model{}
	p.Read(&m, "Address=?", "127.0.0.2")
	assert.NotEqual(t, "127.0.0.2", m.Address)

	p.Close()
}
