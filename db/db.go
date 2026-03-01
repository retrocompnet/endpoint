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
)

type Database struct {
	SchemaVersion int
	LatestMigration int

	connection *gorqlite.Connection
	migrations embed.FS
}
