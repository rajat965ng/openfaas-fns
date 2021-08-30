package function

import (
	"encoding/json"
	"github.com/openfaas/templates-sdk/go-http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerBody(t *testing.T) {

	req := handler.Request{
		Body: []byte("{\"name\":\"abc\"}"),
	}

	res, _ := Handle(req)
	user := &User{}

	json.Unmarshal(res.Body, user)

	assert.Equal(t, "abc", user.Name)
}
