package gen

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/sethvargo/go-password/password"
)

// pathPassword corresponds to POST gen/password.
func (b *backend) pathPassword(_ context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	if err := validateFields(req, d); err != nil {
		return nil, logical.CodedError(http.StatusUnprocessableEntity, err.Error())
	}

	length := d.Get("length").(int)
	digits := d.Get("digits").(int)
	symbols := d.Get("symbols").(int)
	allowUpper := d.Get("allow_uppercase").(bool)
	allowRepeat := d.Get("allow_repeat").(bool)

	pwd, err := password.Generate(length, digits, symbols, !allowUpper, allowRepeat)
	if err != nil {
		return nil, fmt.Errorf("failed to generate password: %w", err)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"value": pwd,
		},
	}, nil
}
