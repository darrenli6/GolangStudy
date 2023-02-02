package test

import (
	"encoding/json"
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/define"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/helper"
	"testing"
)

var userServiceAddr = "http://127.0.0.1:8888"

func TestUserLogin(t *testing.T) {

	m := define.M{
		"username": "get",
		"password": "123",
	}

	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(helper.Md5("123"))
	fmt.Println(string(rep))
}
