package gen

import (
	"context"

	log "github.com/mgutz/logxi/v1"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
	"github.com/pkg/errors"
)

// Factory creates a new usable instance of this secrets engine.
func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
	b := Backend(c)
	if err := b.Setup(ctx, c); err != nil {
		return nil, errors.Wrap(err, "failed to create factory")
	}
	return b, nil
}

// backend is the actual backend.
type backend struct {
	*framework.Backend
	logger log.Logger
}

// Backend creates a new backend.
func Backend(c *logical.BackendConfig) *backend {
	var b backend

	b.logger = c.Logger

	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Help:        backendHelp,
		Paths: []*framework.Path{
			// gen/info
			&framework.Path{
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
			&framework.Path{
				Pattern:      "password",
				HelpSynopsis: "Generate and return a random password",
				HelpDescription: `

Generates a high-entropy password with the provided length and requirements,
returning it as part of the response. The generated password is not stored.

`,
				Fields: map[string]*framework.FieldSchema{
					"length": &framework.FieldSchema{
						Type:        framework.TypeInt,
						Description: "Number of characters for the password.",
						Default:     64,
					},
					"digits": &framework.FieldSchema{
						Type:        framework.TypeInt,
						Description: "Number of digits for the password.",
						Default:     10,
					},
					"symbols": &framework.FieldSchema{
						Type:        framework.TypeInt,
						Description: "Number of symbols for the password.",
						Default:     10,
					},
					"allow_uppercase": &framework.FieldSchema{
						Type:        framework.TypeBool,
						Description: "Allow uppercase letters in the generated password.",
						Default:     true,
					},
					"allow_repeat": &framework.FieldSchema{
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
			&framework.Path{
				Pattern:      "passphrase",
				HelpSynopsis: "Generate and return a random passphrase",
				HelpDescription: `

Generates a random passphrase with the provided number of words, returning it as
part of the response. The generated passphrase is not stored.

`,
				Fields: map[string]*framework.FieldSchema{
					"words": &framework.FieldSchema{
						Type:        framework.TypeInt,
						Description: "Number of words for the passphrase.",
						Default:     6,
					},
					"separator": &framework.FieldSchema{
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

const backendHelp = `
The gen secrets engine generates passwords and passphrases, and optionally
stores the resulting password in an accessor.
`
