// healthcheck.go
// Copyright (C) 2023 tinix <tinix@debian>
//
// Distributed under terms of the MIT license.

package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Set the "Content-Type: application/json"
	w.Header().Set("Content-Type", "application/json")
	// Write the Json as the Http response body.
	w.Write([]byte(js))
}
