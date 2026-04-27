package definition

import (
	"context"

	"github.com/beckn-one/beckn-onix/pkg/model"
)

type RegistryLookup interface {
	// looks up Registry entry to obtain public keys to validate signature of the incoming message
	Lookup(ctx context.Context, req *model.Subscription) ([]model.Subscription, error)
	// Subscribe calls the registry's subscribe endpoint to register a keyset
	Subscribe(ctx context.Context, sub *model.Subscription) error
}

// RegistryLookupProvider initializes a new registry lookup instance.
type RegistryLookupProvider interface {
	New(context.Context, map[string]string) (RegistryLookup, func() error, error)
}
