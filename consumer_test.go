//go:build consumer

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"net/http"
	"testing"
)
func createPact() (pact *dsl.Pact, cleanUp func()) {
	pact = &dsl.Pact{
		Host:                     "localhost",
		Consumer:                 "foo",
		Provider:                 "bar",
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
		LogDir:                   "./pacts/logs",
	}

	cleanUp = func() { pact.Teardown() }

	return pact, cleanUp
}
func publishPact() error{
	p:= dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{"./pacts/foo-bar.json"},
		PactBroker:      "https://eneskzlcn.pactflow.io/",
		BrokerToken:     "L0IzB6WxiCRX7sEdAQoWlQ",
		ConsumerVersion: "1.0.0",
		Tags:            nil,
	})
	if err != nil {
		return err
	}
	return nil
}
func TestIGetPongsArrayAmountOfSentPingCount(t *testing.T){
	pact, cleanUp := createPact()
	defer cleanUp()

	testPingCount := PingCount{Times: 3}
	testResponsePongs := PingPongs{Pongs: []string{"pong","pong","pong"}}
	pact.
		AddInteraction().
		Given("I get pongs array amounf of sent ping count ").
		UponReceiving("A post request with ping count").
		WithRequest(dsl.Request{
			Method:  "POST",
			Path:    dsl.String("/ping"),
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body:  dsl.Like(testPingCount),

		}).
		WillRespondWith(dsl.Response{
			Status:  http.StatusOK,
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body: dsl.Like(testResponsePongs),
		})
		test := func() error {
			return MakePingRequest(pact.Server.Port, testPingCount)
		}
		err := pact.Verify(test)

		if err != nil {
			t.Fatal(err)
		}
		err = publishPact()
		if err!= nil{
			t.Fatal(err)
		}
}