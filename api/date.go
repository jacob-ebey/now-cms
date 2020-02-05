package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler is the main entrypoint for this lambda
func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
