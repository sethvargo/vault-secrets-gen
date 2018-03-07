package gen

import (
	"context"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
	"github.com/sethvargo/vault-secrets-gen/version"
)

// pathInfo corresponds to READ gen/info.
func (b *backend) pathInfo(_ context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	return &logical.Response{
		Data: map[string]interface{}{
			"commit":  version.GitCommit,
			"version": version.Version,
		},
	}, nil
}
