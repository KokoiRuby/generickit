/*
 * Copyright 2024 KoKoiRuby
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package option

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr string
	Port uint16
	// options
	Protocol string
	Timeout  time.Duration
	MaxConn  int
	TLS      *tls.Config
}

func NewServer(addr string, port uint16, opts ...Option[Server]) (*Server, error) {
	srv := Server{
		Addr: addr,
		Port: port,
		// options
		Protocol: "tcp",
		Timeout:  5 * time.Second,
		MaxConn:  100,
		TLS:      nil,
	}
	err := Apply[Server](&srv, opts...)
	if err != nil {
		// Error Handling
	}
	return &srv, nil
}

func WithProtocol(protocol string) Option[Server] {
	return func(s *Server) error {
		s.Protocol = protocol
		return nil
	}
}

func WithTimeout(timeout time.Duration) Option[Server] {
	return func(s *Server) error {
		s.Timeout = timeout
		return nil
	}
}

func WithMaxConn(conn int) Option[Server] {
	return func(s *Server) error {
		s.MaxConn = conn
		return nil
	}
}

func WithTLS(tls *tls.Config) Option[Server] {
	return func(s *Server) error {
		s.TLS = tls
		return nil
	}
}

func ExampleApply() {
	srv, _ := NewServer("127.0.0.1:8080", 8080)
	fmt.Printf("%+v\n", srv)
	srv, _ = NewServer("127.0.0.1:8080", 8080, WithProtocol("udp"), WithTimeout(10*time.Second), WithMaxConn(50))
	fmt.Printf("%+v\n", srv)
	// output:
	// &{Addr:127.0.0.1:8080 Port:8080 Protocol:tcp Timeout:5s MaxConn:100 TLS:<nil>}
	// &{Addr:127.0.0.1:8080 Port:8080 Protocol:udp Timeout:10s MaxConn:50 TLS:<nil>}
}
