

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type PingCount struct{
	Times int `json:"times"`
}
type PingPongs struct{
	Pongs []string `json:"pongs"`
}
func MakePingRequest(port int, pingCount PingCount)error{
	pingCountAsByte,_ := json.Marshal(pingCount)
	resp,err := http.Post(fmt.Sprintf("http://localhost:%d/ping",port),fiber.MIMEApplicationJSON,strings.NewReader(string(pingCountAsByte)))
	if err != nil {
		return err
	}
	var body []byte
	body ,err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	var response = PingPongs{}
	json.Unmarshal(body,&response)
	log.Printf("%v",response)
	return nil
}
func main(){

	//MakePingRequest(4000,"http://localhost",PingCount{Times: 3})
}