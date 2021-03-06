package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

const errMigrateHdr = "[migrate %s@%d] "

const (
	sqlCreateMigrations = `
	CREATE TABLE IF NOT EXISTS
	migrations (
		id SERIAL PRIMARY KEY,
		module varchar(255),
		version INTEGER,

		UNIQUE (module, version)
	)`
	// $1 = module, $2 = version
	sqlInsertMigrations = `INSERT INTO migrations (module, version) VALUES ($1, $2)`
	// $1 = module, $2 = version
	sqlSelectMigrations = `SELECT 1 FROM migrations WHERE module = $1 AND version = $2`
)

func (c *Conn) setupMigrate() (err error) {
	moduleIdentifier := "__core"
	version := 1478019678

	tx, err := c.Begin()
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"start transaction", moduleIdentifier, version)
	}
	defer func(tx *sql.Tx) {
		if err != nil {
			tx.Rollback()
		}
	}(tx)
	_, err = tx.Exec(sqlCreateMigrations)
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"exec %d", moduleIdentifier, version, 1)
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"commit")
	}
	return nil
}

// MustMigrate performs a SQL migration on the database, and panics if the migration fails.
// To perform multiple statements, provide them as multiple queries rather than a single query
// with semicolons; a single call to MustMigrate is wrapped in a single transaction (and is
// therefore atomic, as Postgres has transactional DDL).
//
// This should be called at module load time.
func (c *Conn) MustMigrate(moduleIdentifier string, version int, query ...string) {
	err := c.Migrate(moduleIdentifier, version, query...)
	if err != nil {
		panic(err)
	}
}

// Migrate performs a SQL migration on the database.
// The migration is applied if the migration has not succeeded before.
// Migrations are wrapped in a transaction.
func (c *Conn) Migrate(moduleIdentifier string, version int, query ...string) (err error) {
	if len(moduleIdentifier) > 255 {
		// TODO do this at load time
		panic(errors.Errorf("module identifier should be under 40 characters"))
	}

	ok, err := c.migrationExists(moduleIdentifier, version)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	tx, err := c.Begin()
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"start transaction", moduleIdentifier, version)
	}
	defer func(tx *sql.Tx) {
		if err != nil {
			tx.Rollback()
		}
	}(tx)

	for i := range query {
		_, err = tx.Exec(query[i])
		if err != nil {
			return errors.Wrapf(err, errMigrateHdr+"execute %d", moduleIdentifier, version, i)
		}
	}

	stmt, err := tx.Prepare(sqlInsertMigrations)
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"prepare record", moduleIdentifier, version)
	}
	defer stmt.Close()
	_, err = stmt.Exec(moduleIdentifier, version)
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"insert record", moduleIdentifier, version)
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, errMigrateHdr+"commit", moduleIdentifier, version)
	}
	return nil
}

func (c *Conn) migrationExists(moduleIdentifier string, version int) (bool, error) {
	stmt, err := c.Prepare(sqlSelectMigrations)
	if err != nil {
		return false, errors.Wrapf(err, errMigrateHdr+"prepare check", moduleIdentifier, version)
	}
	defer stmt.Close()

	row := stmt.QueryRow(moduleIdentifier, version)
	var found sql.NullInt64
	err = row.Scan(&found)
	if err == nil && found.Valid {
		return true, nil
	}
	if err != sql.ErrNoRows {
		return false, errors.Wrapf(err, errMigrateHdr+"exec check", moduleIdentifier, version)
	}
	return false, nil
}
