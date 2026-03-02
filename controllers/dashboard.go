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

package controllers

import (
	"embed"
	"net/http"
	"html/template"
	"log"

	"gazelle/endpoint/db"
)

type DashboardHandler struct {
	t *template.Template
	db *db.Database
}

func (dh *DashboardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dh.t.ExecuteTemplate(w, "dashboard", nil)
}

func NewDashboardHandler(db *db.Database, template_fs embed.FS) (*DashboardHandler) {
	tmpl, err := template.New("dashboard").ParseFS(template_fs, "tmpl/*.tmpl")
	if (err != nil) {
		log.Fatal("Failed to instantiate templates: ", err)
	}
	
	return &DashboardHandler{
		t: tmpl,
		db: db,
	}
}
