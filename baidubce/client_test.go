package baidubce

import (
	"os"
	"testing"

	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"
	"github.com/stretchr/testify/assert"
)

func getTestClient() (*Client, error) {
	return New(WithTokenRequest(&baidubcev1.TokenRequest{
		GrantType:    "client_credentials",
		ClientId:     os.Getenv("TEST_BAIDU_CLIENT_ID"),
		ClientSecret: os.Getenv("TEST_BAIDU_CLIENT_SECRET"),
	}))
}

func TestNew(t *testing.T) {
	cli, err := getTestClient()
	assert.Nil(t, err)
	assert.NotNil(t, cli)
}
