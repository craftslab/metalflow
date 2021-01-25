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
	"errors"
)

type Account struct {
	Avatar      string `json:"avatar"`
	Displayname string `json:"displayname"`
	Email       string `json:"email"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

var accounts = []Account{
	{
		Avatar:      "https://user-images.githubusercontent.com/334891/29999089-2837c968-9009-11e7-92c1-6a7540a594d5.png",
		Displayname: "Administrator",
		Email:       "",
		Id:          0,
		Name:        "Administrator",
		Password:    "admin",
		Username:    "admin",
	},
	{
		Avatar:      "https://user-images.githubusercontent.com/334891/29999089-2837c968-9009-11e7-92c1-6a7540a594d5.png",
		Displayname: "Super John",
		Email:       "john.doe@example.com",
		Id:          1,
		Name:        "John Doe",
		Password:    "john",
		Username:    "john",
	},
}

var selfId = 0

func GetAccount(id int) (Account, error) {
	for _, v := range accounts {
		if id == v.Id {
			return v, nil
		}
	}

	return Account{}, errors.New("invalid id")
}

func GetSelfAccount() (Account, error) {
	for _, v := range accounts {
		if selfId == v.Id {
			return v, nil
		}
	}

	return Account{}, errors.New("invalid id")
}

func QueryAccount(q string) (Account, error) {
	if q == "" {
		return Account{}, errors.New("invalid query")
	}

	var buf Account

	for k, v := range accounts {
		if q == v.Username {
			buf = accounts[k]
			break
		}
	}

	return buf, nil
}
