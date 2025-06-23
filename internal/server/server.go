package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rmhyde/fusion/internal/boards"
	"github.com/rs/zerolog"
)

func (o Options) StartWebServer() error {
	logger := zerolog.Ctx(o.Ctx)
	addr := fmt.Sprintf("%s:%d", o.Host, o.Port)
	mux := o.newRouter()
	logger.Info().Msgf("Starting server on %s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Err(err).Msgf("Server failed to start: %v", err)
	}

	return err
}

func (o Options) newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", o.getRoot)
	mux.HandleFunc("GET /api/boards/", o.getBoards)
	return mux
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
	opt := NewCombineOptions(o)

	boards, err := opt.Combine()
	if err != nil {
		logger.Err(err).Msgf("Call to %s failed", r.URL.Path)
		http.Error(w, "Something went wrong", 500)
		return
	}

	// Since this is a webserver we shouldn't share anything about the file structure
	clear(boards.Metadata.Errors.Files)

	response, err := json.Marshal(boards)
	if err != nil {
		logger.Err(err).Msg("Marshal failed")
		http.Error(w, "Something went wrong", 500)
		return
	}
	w.Write(response)
}

func NewCombineOptions(o Options) boards.Options {
	return boards.Options{
		Folder:    o.Folder,
		Recursive: o.Recursive,
		Ctx:       o.Ctx,
	}
}
