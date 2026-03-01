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

package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var (
	//go:embed static
	statics embed.FS
	//go:embed tmpl
	templates embed.FS
	//  //go:embed db/migrations
	//  migrations embed.FS

	Flag_Port     = flag.Int("port", 80, "The port to listen on")
	Flag_DataPath = flag.String("data-path", "/home/rcn/data", "The path to persistent user data storage")
)

const RcnEndpointVersion = "0.0.1"

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, r.RequestURI)
	})
}

func main() {
	log.Printf("RCN Endpoint v%s starting...", RcnEndpointVersion)
	flag.Parse()

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(
		http.FileServer(http.FS(statics)))
	router.Handle("/favicon.ico",
		http.RedirectHandler("/static/images/favicon.ico", 302))
	router.Use(loggingMiddleware)

	server := &http.Server{
		Handler:      router,
		Addr:         ":" + strconv.Itoa(*Flag_Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// dbconn, err := db.New(migrations)
	// if err != nil {
	//   log.Fatal(err)
	// }
	// if err := dbconn.ApplyMigrations(); err != nil {
	//   log.Fatal(err)
	// }

	log.Fatal(server.ListenAndServe())
}
