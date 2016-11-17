package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/szqh97/goTTClient/IM_BaseDefine"
	"github.com/szqh97/goTTClient/IM_Login"
	"testing"
)

func TestLogin(t *testing.T) {
	loginreq := IM_Login.IMLoginReq{}
	username := "dj352801"
	password := "111111"
	version := "1.0.0"
	client_type := IM_BaseDefine.ClientType_CLIENT_TYPE_ANDROID

	loginreq.UserName = &username
	loginreq.Password = &password
	loginreq.ClientVersion = &version
	loginreq.ClientType = &client_type
	fmt.Println(loginreq.String())
	fmt.Printf("%s", loginreq)
	data, _ := proto.Marshal(&loginreq)
	fmt.Printf("%s", data)

}
