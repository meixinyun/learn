package main

import (
	"net/http"
	"os"
	"reflect"
	"runtime"
	"time"

	stdlog "log"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

var logger = stdlog.New(os.Stdout, "[TEST] ", stdlog.Ldate|stdlog.Ltime)

func wrapFunc(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fn(w, req)
		logger.Printf("%s %s   |  Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func wrapFuncWithLogrus(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recovered from panic: %+v / %+v", err, req)
				// http.Error(w, http.StatusText(500), 500)
			}
		}()

		start := time.Now()
		fn(w, req) // execute the handler
		took := time.Since(start)

		// ip, err := getIP(req)
		// if err != nil {
		// 	log.Warnf("getIP error: %v", err)
		// }

		log.WithFields(log.Fields{
			"event_type": runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
			"referrer":   req.Referer(),
			"ua":         req.UserAgent(),
			"method":     req.Method,
			"path":       req.URL.Path,
			"ip":         req.RemoteAddr,
			"real_ip":    getRealIP(req),
			"uuid":       uuid.NewV4(),
		}).Debugf("took %s", took)
	}
}

func wrapHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		logger.Printf("%s %s   |  Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}
