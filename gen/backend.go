// Package gen generates passwords and passphrases.
package gen

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// Factory creates a new usable instance of this secrets engine.
func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
	b := Backend(c)
	if err := b.Setup(ctx, c); err != nil {
		return nil, fmt.Errorf("failed to create factory: %w", err)
	}
	return b, nil
}

// backend is the actual backend.
type backend struct {
	*framework.Backend
}

// Backend creates a new backend.
func Backend(c *logical.BackendConfig) *backend {
	var b backend

	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Help: `

The gen secrets engine generates passwords and passphrases, and optionally
stores the resulting password in an accessor.

		`,
		Paths: []*framework.Path{
			// gen/info
			{
				Pattern:      "info",
				HelpSynopsis: "Display information about this plugin",
				HelpDescription: `

Displays information about the plugin, such as the plugin version and where to
get help.

`,
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.ReadOperation: b.pathInfo,
				},
			},

			// gen/password
			{
				Pattern:      "password",
				HelpSynopsis: "Generate and return a random password",
				HelpDescription: `

Generates a high-entropy password with the provided length and requirements,
returning it as part of the response. The generated password is not stored.

`,
				Fields: map[string]*framework.FieldSchema{
					"length": {
						Type:        framework.TypeInt,
						Description: "Number of characters for the password.",
						Default:     64,
					},
					"digits": {
						Type:        framework.TypeInt,
						Description: "Number of digits for the password.",
						Default:     10,
					},
					"symbols": {
						Type:        framework.TypeInt,
						Description: "Number of symbols for the password.",
						Default:     10,
					},
					"allow_uppercase": {
						Type:        framework.TypeBool,
						Description: "Allow uppercase letters in the generated password.",
						Default:     true,
					},
					"allow_repeat": {
						Type:        framework.TypeBool,
						Description: "Allow repeating characters, digits, and symbols in the generated password.",
						Default:     true,
					},
				},
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.UpdateOperation: b.pathPassword,
				},
			},

			// gen/passphrase
			{
				Pattern:      "passphrase",
				HelpSynopsis: "Generate and return a random passphrase",
				HelpDescription: `

Generates a random passphrase with the provided number of words, returning it as
part of the response. The generated passphrase is not stored.

`,
				Fields: map[string]*framework.FieldSchema{
					"words": {
						Type:        framework.TypeInt,
						Description: "Number of words for the passphrase.",
						Default:     6,
					},
					"separator": {
						Type:        framework.TypeString,
						Description: "Character to separate words in passphrase.",
						Default:     "-",
					},
				},
				Callbacks: map[logical.Operation]framework.OperationFunc{
					logical.UpdateOperation: b.pathPassphrase,
				},
			},
		},
	}

	return &b
}
