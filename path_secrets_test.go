package omsecrets

import (
	"context"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/hashicorp/go-hclog"
)

func getTestBackend(t *testing.T) logical.Backend {
	b, _ := newBackend()

	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	err := b.Setup(context.Background(), c)
	if err != nil {
		t.Fatalf("unable to create backend: %v", err)
	}
	return b
}

func TestGetSecret(t *testing.T) {
	b := getTestBackend(t)

	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      secretsPrefix,
		Data:      make(map[string]interface{}),
	}

	resp, _ := b.HandleRequest(context.Background(), request)
	if resp.Error() != nil {
		t.Errorf("should not have gotten an error")
	}

	if len(resp.Data) != 1 {
		t.Errorf("secret should have single field")
	}

	value, ok := resp.Data[Hello]
	if !ok || value != World {
		t.Errorf("wrong secret")
	}
}
