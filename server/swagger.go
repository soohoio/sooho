package server

import (
	"github.com/soohoio/stayking/v2/client/docs"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterSwaggerAPI provides a common function which registers swagger route with API Server
func RegisterSwaggerAPI(_ client.Context, rtr *mux.Router, swaggerEnabled bool) error {
	if !swaggerEnabled {
		return nil
	}

	root, err := fs.Sub(docs.SwaggerUI, "swagger-ui")
	if err != nil {
		return err
	}

	staticServer := http.FileServer(http.FS(root))
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))

	return nil
}
