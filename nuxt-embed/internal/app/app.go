package app

import (
	"net/http"

	"example/internal/ui"
)

func Start() error {
	return http.ListenAndServe(":8080", ui.WebHandler())
}
