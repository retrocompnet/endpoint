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
	
	"database/sql"
	_ "modernc.org/sqlite"
)

type Database struct {
	SchemaVersion int
	LatestMigration int

	db *sql.DB
}

var (
	//go:embed migrations
	migrations embed.FS
)

func New(dbfile string) (*Database, error) {
	log.Print("Opening sqlite db ", dbfile)
	db, err := sql.Open("sqlite", dbfile)
	if err != nil {
		return nil, err
	}

	version := -1
	result := db.QueryRow("SELECT db_version FROM Migration LIMIT 1")
	_ = result.Scan(&version)

	return &Database{
		SchemaVersion: version,
		LatestMigration: 0,
		db: db,
	}, nil
}
