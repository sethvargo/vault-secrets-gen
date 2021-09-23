package gen

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/sethvargo/go-diceware/diceware"
)

// pathPassphrase corresponds to POST gen/passphrase.
func (b *backend) pathPassphrase(_ context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	if err := validateFields(req, d); err != nil {
		return nil, logical.CodedError(http.StatusUnprocessableEntity, err.Error())
	}

	words := d.Get("words").(int)
	if words == 0 {
		return nil, logical.CodedError(http.StatusUnprocessableEntity, "must generate at least 1 word")
	}

	sep := d.Get("separator").(string)

	list, err := diceware.Generate(words)
	if err != nil {
		return nil, fmt.Errorf("failed to generate passphrase: %w", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"value": strings.Join(list, sep),
		},
	}, nil
}
