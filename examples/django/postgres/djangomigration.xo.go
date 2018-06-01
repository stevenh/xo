// Package postgres contains the types for schema 'public'.
package postgres

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// DjangoMigration represents a row from 'public.django_migrations'.
type DjangoMigration struct {
	ID      int        `json:"id"`      // id
	App     string     `json:"app"`     // app
	Name    string     `json:"name"`    // name
	Applied *time.Time `json:"applied"` // applied

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoMigration exists in the database.
func (dm *DjangoMigration) Exists() bool {
	return dm._exists
}

// Deleted provides information if the DjangoMigration has been deleted from the database.
func (dm *DjangoMigration) Deleted() bool {
	return dm._deleted
}

// Insert inserts the DjangoMigration to the database.
func (dm *DjangoMigration) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.django_migrations (` +
		`app, name, applied` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, dm.App, dm.Name, dm.Applied)
	err = db.QueryRow(sqlstr, dm.App, dm.Name, dm.Applied).Scan(&dm.ID)
	if err != nil {
		return err
	}

	// set existence
	dm._exists = true

	return nil
}

// Update updates the DjangoMigration in the database.
func (dm *DjangoMigration) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.django_migrations SET (` +
		`app, name, applied` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, dm.App, dm.Name, dm.Applied, dm.ID)
	_, err = db.Exec(sqlstr, dm.App, dm.Name, dm.Applied, dm.ID)
	return err
}

// Save saves the DjangoMigration to the database.
func (dm *DjangoMigration) Save(db XODB) error {
	if dm.Exists() {
		return dm.Update(db)
	}

	return dm.Insert(db)
}

// Upsert performs an upsert for DjangoMigration.
//
// NOTE: PostgreSQL 9.5+ only
func (dm *DjangoMigration) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if dm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.django_migrations (` +
		`id, app, name, applied` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, app, name, applied` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.app, EXCLUDED.name, EXCLUDED.applied` +
		`)`

	// run query
	XOLog(sqlstr, dm.ID, dm.App, dm.Name, dm.Applied)
	_, err = db.Exec(sqlstr, dm.ID, dm.App, dm.Name, dm.Applied)
	if err != nil {
		return err
	}

	// set existence
	dm._exists = true

	return nil
}

// Delete deletes the DjangoMigration from the database.
func (dm *DjangoMigration) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dm._exists {
		return nil
	}

	// if deleted, bail
	if dm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.django_migrations WHERE id = $1`

	// run query
	XOLog(sqlstr, dm.ID)
	_, err = db.Exec(sqlstr, dm.ID)
	if err != nil {
		return err
	}

	// set deleted
	dm._deleted = true

	return nil
}

// DjangoMigrationByID retrieves a row from 'public.django_migrations' as a DjangoMigration.
//
// Generated from index 'django_migrations_pkey'.
func DjangoMigrationByID(db XODB, id int) (*DjangoMigration, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, app, name, applied ` +
		`FROM public.django_migrations ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	dm := DjangoMigration{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dm.ID, &dm.App, &dm.Name, &dm.Applied)
	if err != nil {
		return nil, err
	}

	return &dm, nil
}
