package main

import (
	"encoding/json"
	"fmt"
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
func MakePingRequest(port int, host string, pingCount PingCount){
	pingCountAsByte,_ := json.Marshal(pingCount)
	resp,err := http.Post(fmt.Sprintf("%s:%d",host,port),"",strings.NewReader(string(pingCountAsByte)))
	if err != nil {
		log.Fatal(err)
	}
	body ,_:= ioutil.ReadAll(resp.Body)
	var response = PingPongs{}
	json.Unmarshal(body,&response)
	log.Printf("%v",response)
}
func main(){

	//MakePingRequest(4000,"http://localhost",PingCount{Times: 3})
}