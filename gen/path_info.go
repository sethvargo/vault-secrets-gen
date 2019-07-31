package gen

import (
	"context"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/sethvargo/vault-secrets-gen/version"
)

// pathInfo corresponds to READ gen/info.
func (b *backend) pathInfo(_ context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	return &logical.Response{
		Data: map[string]interface{}{
			"name":    version.Name,
			"commit":  version.GitCommit,
			"version": version.Version,
		},
	}, nil
}
