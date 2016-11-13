package ClientConn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ClientConn struct {
	username      string
	password      string
	loginServer   string
	msgServerIP   string
	msgServerPort string
}

func NewClientConn(username, password, loginServer string) *ClientConn {
	return &ClientConn{
		username:    username,
		password:    password,
		loginServer: loginServer,
	}

}

func (cn *ClientConn) GetMsgServerInfo() (err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", cn.loginServer, nil)
	checkerr(err)
	response, err := client.Do(request)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("%s", body)
	}
	return nil
}
