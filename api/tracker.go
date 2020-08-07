package handler

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/crosbymichael/tracker/registry/inmem"
	"github.com/crosbymichael/tracker/server"
)

var s http.Handler

func init() {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	s = server.New(120, 30, inmem.New(), logger)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}
