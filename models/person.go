package models

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/appengine/datastore"
)

type Person struct {
	ID       int64 `datastore:"-"` // "-" means don't read/write from the datastore
	Name     string
	Address  Address
	JobTitle string
}

func (p *Person) Create(ctx context.Context) error {
	var err error
	key := datastore.NewIncompleteKey(ctx, "Person", nil)
	if key, err = datastore.Put(ctx, key, p); err != nil {
		return errors.WithStack(err)
	}
	p.ID = key.IntID()
	return nil
}
