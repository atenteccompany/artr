// MIT License
//
// Copyright (c) 2025 AtenTEC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Runs server only and calls handleConnection to handle
// all incoming connections
package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/atenteccompany/amms-inframgmt/internal/logger"
	"github.com/atenteccompany/amms-inframgmt/tlsutil"
)

const (
	DEFUALT_PORT = 4000
)

var (
	scriptsDir string
)

func Run(port int, sd string) {
	scriptsDir = sd
	// Load CA certificate (to validate client certs)
	// caCert, err := os.ReadFile("/mnt/repos/my_ca/certs/ca/ca.cert.pem")
	// if err != nil {
	// 	log.Fatal("Error reading CA cert:", err)
	// }

	// Load CA from embed
	caCert, err := tlsutil.GetCACert()
	if err != nil {
		log.Fatalf("failed to load embedded CA cert: %v", err)
	}

	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	// Load server (agent) certificate
	cert, err := tls.LoadX509KeyPair("certs/agent.cert.pem", "certs/agent.key.pem")
	if err != nil {
		log.Fatal("Error loading agent cert/key:", err)
	}

	// TLS config that enforces client authentication
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caPool,
		MinVersion:   tls.VersionTLS13, // optional: enforce TLS 1.3
	}

	if port == 0 {
		port = DEFUALT_PORT
	}

	listener, err := tls.Listen("tcp", fmt.Sprintf(":%v", port), tlsConfig)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()
	logger.Info("Server started on port: ", fmt.Sprintf("%v", port))

	// this keeps server running and accepting connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		logger.Info("Client connected: ", conn.LocalAddr().String())

		// Handle connection synchronously
		// We don't depend on routines here because tasks may conflict
		// with each other if run in async because this is a server management tool.
		handleConnection(conn)
	}
}
