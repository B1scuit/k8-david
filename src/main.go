package main

import (
	configDAO "./contextDAO"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	. "github.com/tylerwince/godbg"
	"log"
	"net/http"
	"os"
	"time"
)

var errLog = log.New(os.Stderr, "", 0) // Initalize an error log handler

func main() {
	fmt.Println("Starting HTTP Server...")

	r := mux.NewRouter()

	// All API routes
	r.HandleFunc("/", statusOK).Methods("GET")

	// All API endpoints for config
	r.HandleFunc("/config", contextCurrentContext).Methods("GET")
	r.HandleFunc("/config/view", contextConfigView).Methods("GET")
	r.HandleFunc("/config", contextPutContext).Methods("PUT")

	srv := &http.Server{
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	errLog.Fatal(srv.ListenAndServe())
}

func contextConfigView(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, configDAO.ConfigView())
}

func contextPutContext(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("context")

	if name == "" {
		respondWithError(w, 422, "Context name missing")
		return
	}

	res, err := configDAO.SetContext(name)

	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}

	respondWithJson(w, 200, res)
}

func contextCurrentContext(w http.ResponseWriter, r *http.Request) {
	cmd, err := configDAO.CurrentContext()

	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	respondWithJson(w, 200, cmd)
}

// Collects any responses from the response channel and JSON encodes + sends the
// response back to the client
//
// Peter Holt <peter.holt@dochq.co.uk>
func responseHandle(conn *websocket.Conn, respChan <-chan interface{}) {
	repl := <-respChan

	Dbg(repl)

	_ = conn.WriteJSON(repl)
}

// Should an error occur, respond with an error code and a JSON message
// This will also send the error to the STDERR output, allowing GCE to pick it up
//
// Peter Holt <peter.holt@dochq.co.uk>
func respondWithError(w http.ResponseWriter, code int, msg string) {
	errLog.Println(msg)
	respondWithJson(w, code, map[string]string{"error": msg})
}

// Response handler, takes an interface and converts it into JSON and
// sets the response with the cuorrect headers
//
// Peter Holt <peter.holt@dochq.co.uk>
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// In order to use Google stackdriver for uptime checking
// the service needs top provide an endpoint to ping against
// this is designed to just return a status 200 so stackdriver
// can tell weather the APi is up or not
//
// Peter Holt <peter.holt@dochq.co.uk>
func statusOK(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Status string `json:"status"`
	}{
		"OK",
	}

	respondWithJson(w, http.StatusOK, res)
}
