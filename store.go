package scs

import (
	"context"
	"time"
)

// Store is the interface for session stores.
type Store interface {
	// Delete should remove the session token and corresponding data from the
	// session store. If the token does not exist then Delete should be a no-op
	// and return nil (not an error).
	Delete(token string) (err error)

	// Find should return the data for a session token from the store. If the
	// session token is not found or is expired, the found return value should
	// be false (and the err return value should be nil). Similarly, tampered
	// or malformed tokens should result in a found return value of false and a
	// nil err value. The err return value should be used for system errors only.
	Find(token string) (b []byte, found bool, err error)

	// Commit should add the session token and data to the store, with the given
	// expiry time. If the session token already exists, then the data and
	// expiry time should be overwritten.
	Commit(token string, b []byte, expiry time.Time) (err error)
}

// ContextStore is the interface for session stores that need a request
// context, e.g. Google Cloud Platform Firestore.
type ContextStore interface {
	// Delete should remove the session token and corresponding data from the
	// session store. If the token does not exist then Delete should be a no-op
	// and return nil (not an error).
	Delete(ctx context.Context, token string) (err error)

	// Find should return the data for a session token from the store. If the
	// session token is not found or is expired, the found return value should
	// be false (and the err return value should be nil). Similarly, tampered
	// or malformed tokens should result in a found return value of false and a
	// nil err value. The err return value should be used for system errors only.
	Find(ctx context.Context, token string) (b []byte, found bool, err error)

	// Commit should add the session token and data to the store, with the given
	// expiry time. If the session token already exists, then the data and
	// expiry time should be overwritten.
	Commit(ctx context.Context, token string, b []byte, expiry time.Time) (err error)
}

// StoreAdapter is used to for the scs version 2 store with a ContextStore,
// dropping the unused context argument.
type StoreAdapter struct {
	Store Store
}

func (sa *StoreAdapter) Delete(ctx context.Context, token string) (err error) {
	return sa.Store.Delete(token)
}

func (sa *StoreAdapter) Find(ctx context.Context, token string) (b []byte, found bool, err error) {
	return sa.Store.Find(token)
}

func (sa *StoreAdapter) Commit(ctx context.Context, token string, b []byte, expiry time.Time) (err error) {
	return sa.Store.Commit(token, b, expiry)
}
