package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/alexedwards/scs/redisstore"
)

var sessionManager *scs.SessionManager

func main() {
	// Establish a redigo connection pool.
	redisPort := os.Getenv("REDIS_PORT")
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			redisURL := fmt.Sprintf("localhost:%v", redisPort)
			return redis.Dial("tcp", redisURL)
		},
	}

	// Initialize a new session manager and configure it to use redisstore as
	// the session store.
	sessionManager = scs.New()
	sessionManager.Store = redisstore.New(pool)

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	// Wrap your handlers with the LoadAndSave() middleware.
	http.ListenAndServe(":4000", sessionManager.LoadAndSave(mux))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	sessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := sessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}
