package httputils

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMyClientRequestGet(t *testing.T) {
	c := GetClient(time.Second * 2)

	header := map[string]string{
		"X-Forward-IP": "127.0.0.1",
		"X-Trace-ID":   "1234",
	}

	Query := map[string]interface{}{
		"a": fmt.Sprint(1),
		"b": "c",
	}

	resp, err := c.RequestGet(context.TODO(), "", header, Query)
	assert.NotNil(t, err)
	fmt.Println(string(resp))
}
