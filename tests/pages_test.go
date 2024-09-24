package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3000"

	var (
		resp *http.Response
		err  error
	)

	resp, err = http.Get(baseURL + "/")

	assert.NoError(t, err, "error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "statuscode")
}
