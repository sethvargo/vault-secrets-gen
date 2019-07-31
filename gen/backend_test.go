package gen

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/sethvargo/vault-secrets-gen/version"
)

func testBackend(tb testing.TB) (*backend, logical.Storage) {
	tb.Helper()

	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}

	b, err := Factory(context.Background(), config)
	if err != nil {
		tb.Fatal(err)
	}
	return b.(*backend), config.StorageView
}

func TestBackend(t *testing.T) {
	t.Run("info", func(t *testing.T) {
		t.Parallel()

		b, storage := testBackend(t)
		resp, err := b.HandleRequest(context.Background(), &logical.Request{
			Storage:   storage,
			Operation: logical.ReadOperation,
			Path:      "info",
		})
		if err != nil {
			t.Fatal(err)
		}

		if v, exp := resp.Data["version"].(string), version.Version; v != exp {
			t.Errorf("expected %q to be %q", v, exp)
		}

		if v, exp := resp.Data["commit"].(string), version.GitCommit; v != exp {
			t.Errorf("expected %q to be %q", v, exp)
		}
	})

	t.Run("password", func(t *testing.T) {
		t.Parallel()

		b, storage := testBackend(t)
		resp, err := b.HandleRequest(context.Background(), &logical.Request{
			Storage:   storage,
			Operation: logical.UpdateOperation,
			Path:      "password",
		})
		if err != nil {
			t.Fatal(err)
		}

		v := resp.Data["value"].(string)
		if len(v) != 64 {
			t.Errorf("expected %q to be 64 chars", v)
		}
	})

	t.Run("passphrase", func(t *testing.T) {
		t.Parallel()

		b, storage := testBackend(t)
		resp, err := b.HandleRequest(context.Background(), &logical.Request{
			Storage:   storage,
			Operation: logical.UpdateOperation,
			Path:      "passphrase",
		})
		if err != nil {
			t.Fatal(err)
		}

		v := resp.Data["value"].(string)
		parts := strings.Split(v, "-")
		if len(parts) != 6 {
			t.Errorf("expected %q to be 6 parts", v)
		}
	})
}
