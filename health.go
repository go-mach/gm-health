// Copyright (c) 2019 Luca Stasio
// Use of this source code is governed by an MIT license that can be
// found in the LICENSE file at https://github.com/go-mach/gm-health/blob/master/LICENSE

// Package health defines the go-mach pluggable Health endpoint component
package health

import (
	"encoding/json"
	"log"
	"net/http"
)

// Status keeps the status to return
type status struct {
	Status string `json:"status"`
}

// ServeDefault serve the default health {"status": "UP"} endpoint
func ServeDefault() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		status := status{"UP"}
		jsonStatus, err := json.Marshal(status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonStatus)
	})

	log.Fatal(http.ListenAndServe(":3010", nil))
}
