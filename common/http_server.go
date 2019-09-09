/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file http_server.go
 * @package common
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 05/23/2019
 */

package common

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// List
var (
	_HTTPServerList = make(map[uint64]*HTTPServer)
	_HTTPServerID   uint64
)

// HTTPServer : HTTP server defination
/* {{{ [HTTPServer::struct] */
type HTTPServer struct {
	// addr : Server address listened to
	addr string

	// ID : Server ID
	id uint64

	// TLS : If TLS
	TLS bool

	// CertFile : SSL certificate file
	CertFile string

	// KeyFile : SSL certificate key file
	KeyFile string

	// Defaults
	ContentType string

	// server : Real server
	server fasthttp.Server

	// Router : Fasthttprouter
	Router *fasthttprouter.Router

	// Hooks
	OnRequest fasthttp.RequestHandler

	// Runtime
	Runtime *GlobalRuntime
}

/* }}} */

// HTTPServerRoute : Route defination of HTTP server
/* {{{ [HTTPServerRoute::struct] */
type HTTPServerRoute struct {
	// Method : HTTP method
	Method string

	// Path : Route path
	Path string

	// Handler : Route handler
	Handler fasthttp.RequestHandler
}

/* }}} */

// NewHTTPServer : Create new HTTP server
/* {{{ [NewHTTPServer] */
func NewHTTPServer(r *GlobalRuntime) *HTTPServer {
	var s = &HTTPServer{
		Router:  fasthttprouter.New(),
		Runtime: r,
	}

	_HTTPServerList[_HTTPServerID] = s
	_HTTPServerID++

	return s
}

/* }}} */

// GetHTTPServer : Get HTTP server from list
/* {{{ [GetHTTPServer] */
func GetHTTPServer(id uint64) *HTTPServer {
	return _HTTPServerList[id]
}

/* }}} */

// ID : Get ID of server
/* {{{ [HTTPServer:READONLY] */
func (s *HTTPServer) ID() uint64 {
	return s.id
}

/* }}} */

// Start : Startup HTTP server
/* {{{ [HTTPServer.Start] */
func (s *HTTPServer) Start() {
	s.addr = s.Runtime.Config.GetString("HTTP_Listen_Address")
	if s.OnRequest != nil {
		s.server.Handler = s.OnRequest
	} else {
		s.server.Handler = s.Router.Handler
	}

	s.server.Logger = s.Runtime.Logger
	s.Runtime.Waiter.Add(1)

	go func() {
		if s.TLS == true {
			s.server.Logger.Printf("HTTP server initialized at [%s] with TLS enabled", s.addr)
			s.server.ListenAndServeTLS(s.addr, s.CertFile, s.KeyFile)
		} else {
			s.server.Logger.Printf("HTTP server initialized at [%s]", s.addr)
			s.server.ListenAndServe(s.addr)
		}
	}()

	return
}

/* }}} */

// Stop : Stop HTTP server
/* {{{ [HTTPServer.Stop] */
func (s *HTTPServer) Stop() {
	s.Runtime.Waiter.Done()

	return
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
