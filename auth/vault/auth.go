package vault

import (
	"context"

	"github.com/gotd/td/telegram/tgflow"
	"github.com/hashicorp/vault/api"
)

const (
	phoneKey    = "phone"
	passwordKey = "password"
)

// Auth is tgflow.UserAuthenticator implementation
type Auth struct {
	tgflow.CodeAuthenticator

	vault vaultClient
	path  string
}

var _ tgflow.UserAuthenticator = Auth{}

// NewAuth creates new Auth.
func NewAuth(code tgflow.CodeAuthenticator, client *api.Client, path string) Auth {
	return Auth{
		CodeAuthenticator: code,
		vault:             vaultClient{Client: client},
		path:              path,
	}
}

// SavePhone stores given phone to the Vault.
func (a Auth) SavePhone(ctx context.Context, phone string) error {
	return a.vault.put(ctx, a.path, phoneKey, phone)
}

// SavePassword stores given password to the Vault.
func (a Auth) SavePassword(ctx context.Context, password string) error {
	return a.vault.put(ctx, a.path, passwordKey, password)
}

// Phone loads phone from the Vault.
func (a Auth) Phone(ctx context.Context) (string, error) {
	return a.vault.get(ctx, a.path, phoneKey)
}

// Password loads password from the Vault.
func (a Auth) Password(ctx context.Context) (string, error) {
	return a.vault.get(ctx, a.path, passwordKey)
}