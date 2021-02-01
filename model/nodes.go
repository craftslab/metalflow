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

package model

import (
	"github.com/pkg/errors"
)

type Node struct {
	Address  string `json:"address"`
	Asset    string `json:"asset"`
	Comments string `json:"comments"`
	Health   string `json:"health"`
	Id       uint   `json:"id"`
	Info     string `json:"info"`
	Perf     string `json:"perf"`
	Region   string `json:"region"`
}

var nodes = []Node{
	{
		Address:  "127.0.0.1",
		Asset:    "0",
		Comments: "node 0",
		Health:   "running",
		Id:       0,
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
	},
	{
		Address:  "127.0.0.2",
		Asset:    "1",
		Comments: "node 1",
		Health:   "stop",
		Id:       1,
		Info: `"
			{
				"bare": {
					"cpu": "4 CPU",
					"disk": "49.0 GB (16.0 GB Used)",
					"io": "RD 11887928 KB WR 61067948 KB",
					"ip": "127.0.0.2",
					"kernel": "5.4.0-58-generic",
					"mac": "00:01:02:03:04:06",
					"network": "RX packets 8974179 TX packets 3124096",
					"os": "Ubuntu 18.04.5 LTS",
					"ram": "7692 MB (1345 MB Used)",
					"system": ""
				}
			}
		"`,
		Perf:   "Low",
		Region: "Xian",
	},
}

func GetNode(id uint) (Node, error) {
	var f bool
	var n Node

	f = false

	for _, v := range nodes {
		if id == v.Id {
			f = true
			n = v
			break
		}
	}

	if !f {
		return Node{}, errors.New("invalid id")
	}

	return n, nil
}

func GetHealth(id uint) (string, error) {
	var f bool
	var n Node

	f = false

	for _, v := range nodes {
		if id == v.Id {
			f = true
			n = v
			break
		}
	}

	if !f {
		return "", errors.New("invalid id")
	}

	return n.Health, nil
}

func GetInfo(id uint) (string, error) {
	var f bool
	var n Node

	f = false

	for _, v := range nodes {
		if id == v.Id {
			f = true
			n = v
			break
		}
	}

	if !f {
		return "", errors.New("invalid id")
	}

	return n.Info, nil
}

func GetPerf(id uint) (string, error) {
	var f bool
	var n Node

	f = false

	for _, v := range nodes {
		if id == v.Id {
			f = true
			n = v
			break
		}
	}

	if !f {
		return "", errors.New("invalid id")
	}

	return n.Perf, nil
}

func QueryNode(q string) (Node, error) {
	if q == "" {
		return Node{}, errors.New("invalid query")
	}

	var buf Node

	for k, v := range nodes {
		if q == v.Address {
			buf = nodes[k]
			break
		}
	}

	return buf, nil
}

func AddNode(id uint) (Node, error) {
	// TODO
	return Node{}, nil
}

func DelNode(id uint) (Node, error) {
	// TODO
	return Node{}, nil
}
