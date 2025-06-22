package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rmhyde/fusion/internal/boards"
	"github.com/rs/zerolog"
)

func (o Options) StartWebServer(ctx context.Context) error {
	o.Ctx = ctx
	logger := zerolog.Ctx(ctx)
	addr := fmt.Sprintf("%s:%d", o.Host, o.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", o.getRoot)
	mux.HandleFunc("GET /api/boards/", o.getBoards)
	err := http.ListenAndServe(addr, mux) // Pass 'mux' instead of 'nil'
	if err != nil {
		logger.Err(err).Msgf("Server failed to start: %v", err)
	}
	return nil
}

func (o Options) getRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Boards can be found in /api/boards")
}

func (o Options) getBoards(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	logger := zerolog.Ctx(o.Ctx)
	opt := boards.Options{
		Folder:    o.Folder,
		Recursive: o.Recursive,
		Ctx:       o.Ctx,
	}

	b, err := opt.Combine()
	if err != nil {
		logger.Err(err).Msgf("Call to %s failed", r.URL.Path)
		http.Error(w, "Something went wrong", 500)
		return
	}

	// Since this is a webserver we shouldn't share anything about the file structure
	clear(b.Metadata.Errors.Files)

	response, err := json.Marshal(b)
	if err != nil {
		logger.Err(err).Msg("Marshal failed")
		http.Error(w, "Something went wrong", 500)
		return
	}
	w.Write(response)
}
