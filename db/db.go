// Copyright 2026, Retrocomp Networks Foundation, Inc.
//
// This file is part of RCN endpoint, a user-facing program to make
// management of an RCN endpoint easy.
//
// librcn-api is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see
// <https://www.gnu.org/licenses/>.

package db

import (
	"embed"
	"log"
	"os"
	"path"
	"io/fs"
	"sort"
	"strings"
	"strconv"

	"database/sql"

	_ "modernc.org/sqlite"
)

type Database struct {
	SchemaVersion int
	db *sql.DB
}

var (
	//go:embed migrations
	migrations embed.FS
)

func applyMigration(entry os.DirEntry, db *sql.DB) (error) {
	log.Print("Applying migration ", entry.Name())
	migration, err := fs.ReadFile(migrations, path.Join("migrations/", entry.Name()))
	if err != nil {
		return err
	}

	if tx, err := db.Begin(); err == nil {
		if _, err := tx.Exec(string(migration[:])); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}

			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}

func getSchemaVersion(db *sql.DB) (int) {
	version := -1
	result := db.QueryRow("SELECT db_version FROM Migration LIMIT 1")
	_ = result.Scan(&version)

	return version
}

func New(dbfile string) (*Database, error) {
	log.Print("Opening sqlite db ", dbfile)
	db, err := sql.Open("sqlite", dbfile)
	if err != nil {
		return nil, err
	}

	version := getSchemaVersion(db)

	if entries, err := migrations.ReadDir("migrations"); err == nil {
		sort.Slice(entries, func(i, j int) bool {
			return strings.Compare(entries[i].Name(), entries[j].Name()) == -1
		})

		for _, entry := range entries {
			migration_version, err := strconv.Atoi(strings.Split(entry.Name(), "_")[0])
			if err != nil {
				return nil, err
			}

			if migration_version > version {
				if err := applyMigration(entry, db); err != nil {
					return nil, err
				}
			}
		}
	}
	
	new_version := getSchemaVersion(db)

	if version != new_version {
		log.Printf("Database migration complete. Schema upgraded from %d to %d", version, new_version)
	} else {
		log.Print("Database ready.")
	}

	return &Database{
		SchemaVersion: version,
		db: db,
	}, nil
}
