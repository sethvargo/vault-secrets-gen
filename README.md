# Password Generator for HashiCorp Vault

The Vault Password Generator is a [Vault](https://www.vaultproject.io) secrets
plugin for generating cryptographically secure passwords and passphrases.

This is both a real custom Vault secrets plugin, and an example of how to build,
install, and maintain your own Vault secrets plugin.

## Setup

The setup guide assumes some familiarity with Vault and Vault's plugin
ecosystem. You must have a Vault server already running, unsealed, and
authenticated.

1. Download and decompress the latest plugin binary from the Releases tab on
GitHub. Alternatively you can compile the plugin from source, if you're into
that kinda thing.

1. Move the compiled plugin into Vault's configured `plugin_directory`. You must
   set this value in the [server configuration][vault-config]:

    ```sh
    $ mv vault-secrets-gen /etc/vault/plugins/vault-secrets-gen
    ```

1. Enable mlock so the plugin can safely be enabled and disabled:

   ```sh
   setcap cap_ipc_lock=+ep /etc/vault/plugins/vault-secrets-gen
   ```
    > **_NOTE:_** For vault in kubernetes you don't need to do this step. If you want to enable mlock you need to edit the security context to run with privileges permissions and run container with root. Also you will need to set the /etc/vault/plugins/vault-secrets-gen with vault:vault permissions

1. Calculate the SHA256 of the plugin and register it in Vault's plugin catalog.
If you are downloading the pre-compiled binary, it is highly recommended that
you use the published checksums to verify integrity.

    ```sh
    $ export SHA256=$(shasum -a 256 "/etc/vault/plugins/vault-secrets-gen" | cut -d' ' -f1)

    $ vault plugin register \
        -sha256="${SHA256}" \
        -command="vault-secrets-gen" \
        secret secrets-gen
    ```

1. Mount the secrets engine:

    ```sh
    $ vault secrets enable \
        -path="gen" \
        -plugin-name="secrets-gen" \
        plugin
    ```

## Usage & API

### Policy requirements

The token used should have the following policy permissions to be able to generate passwords.

```hcl
path "gen/password" {
  capabilities = ["create", "update"]
}
```

### Generate Password

Generates a random, high-entropy password with the specified number of
characters, digits, symbols, and configurables.

| Method   | Path                         | Produces                 |
| :------- | :--------------------------- | :----------------------- |
| `POST`   | `/gen/password`              | `200 (application/json)` |

#### Parameters

- `length` `(int: 64)` - Specifies the total length of the password including
  all letters, digits, and symbols.

- `digits` `(int: 10)` - Specifies the number of digits to include in the
  password.

- `symbols` `(int: 10)` - Specifies the number of symbols to include in the
  password.

- `allow_uppercase` `(bool: true)` - Specifies whether to allow uppercase and
  lowercase letters in the password.

- `allow_repeat` `(bool: true)` - Specifies to allow duplicate characters in the
  password. If set to false, be conscious of password length as values cannot be
  re-used.

#### CLI

```
$ vault write gen/password length=36 symbols=0
Key  	Value
---  	-----
value	27f3L5zKCZS8DD6D2PEK1xm0ECNaImg1PJqg
```

### Generate Passphrase

Generates a random, high-entropy passphrase with the specified number of words
and separator using the diceware algorithm.

| Method   | Path                         | Produces                 |
| :------- | :--------------------------- | :----------------------- |
| `POST`   | `/gen/passphrase`            | `200 (application/json)` |

#### Parameters

- `words` `(int: 6)` - Specifies the total number of words to generate.

- `separator` `(string: "-")` - Specifies the string value to use as a separator
  between words.

#### CLI

```
$ vault write gen/passphrase words=4
Key  	Value
---  	-----
value	obstacle-sacrament-sizable-variably
```

## License

This code is licensed under the MIT license.

[vault-config]: https://www.vaultproject.io/docs/configuration#plugin_directory
