package server

import (
	"context"

	appdb "github.com/wildberries-tech/sgroups/v2/internal/registry/sgroups"

	"github.com/pkg/errors"
)

func newInternalDB(ctx context.Context) (appdb.Registry, error) {
	m, e := appdb.NewMemDB(appdb.AllTables())
	if e != nil {
		return nil, errors.WithMessage(e, "create mem db")
	}
	return appdb.NewRegistryFromMemDB(m), nil
}
