// Package utils is the package for runner utils.
package utils

import (
	"bytes"
	"context"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/component/dbfactory"

	"github.com/bytebase/bytebase/backend/plugin/db"
	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/transform"
	"github.com/bytebase/bytebase/backend/store"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

// ComputeDatabaseSchemaDiff computes the diff between current database schema
// and the given schema. It returns an empty string if there is no applicable
// diff.
func ComputeDatabaseSchemaDiff(ctx context.Context, instance *store.InstanceMessage, database *store.DatabaseMessage, dbFactory *dbfactory.DBFactory, newSchema string) (string, error) {
	driver, err := dbFactory.GetAdminDatabaseDriver(ctx, instance, database, db.ConnectionContext{})
	if err != nil {
		return "", errors.Wrap(err, "get admin driver")
	}
	defer func() {
		_ = driver.Close(ctx)
	}()

	// Use new driver to sync the schema to avoid the session state change, such as SET ROLE in PostgreSQL.
	syncDriver, err := dbFactory.GetAdminDatabaseDriver(ctx, instance, database, db.ConnectionContext{})
	if err != nil {
		return "", errors.Wrap(err, "get sync driver")
	}
	defer syncDriver.Close(ctx)
	dbSchema, err := syncDriver.SyncDBSchema(ctx)
	if err != nil {
		return "", errors.Wrap(err, "sync database schema")
	}

	var schema bytes.Buffer
	err = driver.Dump(ctx, &schema, dbSchema)
	if err != nil {
		return "", errors.Wrap(err, "dump old schema")
	}

	var engine storepb.Engine
	switch instance.Metadata.GetEngine() {
	case storepb.Engine_POSTGRES, storepb.Engine_RISINGWAVE:
		engine = storepb.Engine_POSTGRES
	case storepb.Engine_MYSQL, storepb.Engine_TIDB, storepb.Engine_MARIADB, storepb.Engine_OCEANBASE:
		engine = storepb.Engine_MYSQL
	default:
		return "", errors.Errorf("unsupported database engine %q", instance.Metadata.GetEngine())
	}

	sdlFormat, err := transform.SchemaTransform(engine, schema.String())
	if err != nil {
		return "", errors.Wrapf(err, "failed to transform SDL format")
	}
	diff, err := base.SchemaDiff(engine, base.DiffContext{
		IgnoreCaseSensitive: !store.IsObjectCaseSensitive(instance),
		StrictMode:          true,
	}, sdlFormat, newSchema)
	if err != nil {
		return "", errors.Wrapf(err, "compute schema diff")
	}
	return diff, nil
}
