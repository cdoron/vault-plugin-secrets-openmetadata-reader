package omsecrets

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type KubernetesSecretsReader struct {
}

const Hello = "hello"
const World = "world"

// GetSecret returns the content of openmetadata secret.
func (s *KubernetesSecretsReader) GetSecret(ctx context.Context, secretName string, log hclog.Logger) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	data[Hello] = World

	return data, nil
}
