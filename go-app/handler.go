package function

import (
	"encoding/json"
	"fmt"
	"github.com/openfaas/templates-sdk/go-http"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	user := &User{}
	err = json.Unmarshal([]byte(req.Body), user)
	if err != nil {
		fmt.Printf("Error Unmarshalling: %s\n", err.Error())
	}

	data, err := json.Marshal(&user)
	return handler.Response{
		Body:       data,
		StatusCode: http.StatusOK,
	}, err
}
