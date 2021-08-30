package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"net/http"
	"net/http/httptest"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

type User struct {
	Name string `json:"name"`
	Age  int 	`json:"age"`
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendrequestTo(method, endpoint string) (err error) {

	user := &User{
		Name: "testUser",
		Age: 18,
	}

	buff,err := json.Marshal(user)

	resp,err := http.Post(endpoint,"application/json",bytes.NewReader(buff))
	if err != nil {
		return
	}


	a.resetResponse(resp)


	json.NewDecoder(resp.Body).Decode(user)
	a.resp.WriteHeader(resp.StatusCode)

	buff,_  = json.Marshal(user)
	a.resp.Write(buff)


	return
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) error {
	var expected, actual []byte
	var data User
	var err error
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return err
	}
	if expected, err = json.Marshal(data); err != nil {
		return err
	}

	actual = a.resp.Body.Bytes()

	if !bytes.Equal(actual, expected) {
		err = fmt.Errorf("expected json, does not match actual: %s", string(actual))
	}
	return err
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &apiFeature{}


	s.Step(`^I send "(POST)" request to "([^"]*)"$`, api.iSendrequestTo)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, api.theResponseShouldMatchJSON)
}
