package categories_test

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"fmt"
	"math/rand"

	"gotest.tools/v3/assert"
)

func post(resource string, body []byte) (*http.Response, error) {
	uri := fmt.Sprintf("http://localhost:3000/%s", resource)

	bodyReader := bytes.NewReader(body)

	return http.Post(uri, "application/json", bodyReader)
}

func TestCreate(t *testing.T) {
	name := fmt.Sprintf("Category #%d", rand.Intn(1_000_000))
	body := []byte(fmt.Sprintf(`{"name":"Category #%s"}`, name))

	t.Run("when Category is uniq", func(t *testing.T) {
		response, err := post("categories", body)
		assert.NilError(t, err)

		assert.Equal(t, response.Header.Get("content-type"), "application/json")
		assert.Equal(t, response.Status, "201 Created")

		defer response.Body.Close()
	})

	t.Run("when Category already exists", func(t *testing.T) {
		response, err := post("categories", body)
		assert.NilError(t, err)

		responseBody, err := io.ReadAll(response.Body)
		assert.NilError(t, err)

		assert.Equal(t, response.Header.Get("content-type"), "application/json")
		assert.Equal(t, response.Status, "422 Unprocessable Entity")
		assert.Equal(
			t,
			strings.TrimSpace(string(responseBody)),
			`{"api":{"application":{"database":{"query":"UNIQUE constraint failed: `+
				`categories.name, categories.currency"}}}}`,
		)

		defer response.Body.Close()
	})
}
