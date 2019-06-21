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

func healthHandler(jsonStatus []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonStatus)
	})
}

// ServeDefault serve the default health {"status": "UP"} endpoint
func ServeDefault(addr string) error {
	return Serve(addr, status{"UP"})
}

// Serve starts the health endpoint listening on addr and serving the s status
func Serve(addr string, s interface{}) error {
	jsonStatus, err := json.Marshal(s)
	if err != nil {
		return err
	}
	http.Handle("/health", healthHandler(jsonStatus))
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
