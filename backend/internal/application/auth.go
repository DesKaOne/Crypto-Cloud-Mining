package application

import (
	"context"
	"errors"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("forbidden")
)

// Principal is the authenticated identity passed from transport/auth adapters.
type Principal struct {
	UserID    domain.UserID
	Provider  string
	SessionID string
}

// Authorizer keeps authorization decisions in the application boundary.
type Authorizer interface {
	CanStartMining(ctx context.Context, principal Principal, plan domain.MiningPlan) error
	CanViewWallet(ctx context.Context, principal Principal, userID domain.UserID, assetID domain.AssetID) error
}

// TelegramIdentity is the trusted result of backend validation.
type TelegramIdentity struct {
	ProviderSubject string
	Username        string
}

// TelegramAuthenticator validates Telegram signed initialization data.
type TelegramAuthenticator interface {
	Authenticate(ctx context.Context, initData string) (TelegramIdentity, error)
}

// IdentityStore resolves or provisions a User for a validated external identity.
type IdentityStore interface {
	FindOrCreateUser(ctx context.Context, provider string, providerSubject string) (domain.UserID, error)
}
