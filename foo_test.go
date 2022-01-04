package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Testing a post operation done correctly is enough for unit test


func TestIPerformPostOperationToAnApiSuccessfully(t *testing.T){
	type PingCount struct{
		Times int `json:"times"`
	}
	testData := PingCount{Times: 2}

	// here we created an api to test if our post operation is success.
	app := fiber.New()
	app.Post("/ping", func(ctx *fiber.Ctx) error {
		sentData := PingCount{}
		if err := ctx.BodyParser(&sentData) ; err != nil{
			return ctx.SendStatus(http.StatusNotFound)
		}
		return ctx.JSON(sentData)
	})

	testDataAsByte,_ := json.Marshal(testData)
	req:= httptest.NewRequest(http.MethodPost,"/ping",strings.NewReader(string(testDataAsByte)))
	req.Header.Set("Content-Type","application/json")
	resp,_ := app.Test(req,1)

	resentPingCountAsByte,_ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t,string(testDataAsByte),string(resentPingCountAsByte),"Posted ping count struct not go the api correctly.")
}

