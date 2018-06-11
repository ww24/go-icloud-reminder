package datastore

import (
	"context"

	"google.golang.org/appengine/datastore"
)

const (
	secretKind = "Secret"
	secretID   = "secret"
)

type Secret struct {
	AppleWebauthUser  string `datastore:"apple_web_authtoken_user,noindex"`
	AppleWebauthToken string `datastore:"apple_web_authtoken_token,noindex"`
}

func (s *Secret) Save(ctx context.Context) error {
	k := datastore.NewKey(ctx, secretKind, secretID, 0, nil)
	if _, err := datastore.Put(ctx, k, s); err != nil {
		return err
	}
	return nil
}

func (s *Secret) Get(ctx context.Context) error {
	k := datastore.NewKey(ctx, secretKind, secretID, 0, nil)
	if err := datastore.Get(ctx, k, s); err != nil {
		return err
	}
	return nil
}

func (s *Secret) Valid() bool {
	return s.AppleWebauthUser != "" && s.AppleWebauthToken != ""
}
