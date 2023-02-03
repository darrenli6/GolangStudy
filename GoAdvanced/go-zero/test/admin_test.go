package test

import (
	"encoding/json"
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"

var m1 = map[string]string{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIiLCJuYW1lIjoiZ2V0IiwiZXhwIjoxNjc3ODYxOTUxfQ.pDtEdPB3dEhTtR08XTy0sBBhGyaa3zre53XTvHZVi40",
}

var header []byte

func init() {
	header, _ = json.Marshal(m1)
}

func TestDeviceList(t *testing.T) {

	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

func TestProductList(t *testing.T) {
	rep, err := helper.HttpGet(adminServiceAddr+"/product/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
