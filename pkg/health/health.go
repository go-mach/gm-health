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

		w.Header().Set("Server", "sospediatra-video-poller")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonStatus)
	})

	log.Fatal(http.ListenAndServe(":3010", nil))
}
