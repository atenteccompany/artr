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

package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"github.com/atenteccompany/artr/tlsutil"
)

func getConn(addr string) (*tls.Conn, error) {
	// Load CA cert to verify agent's identity
	// caCert, err := os.ReadFile("/mnt/repos/my_ca/certs/ca/ca.cert.pem")
	// if err != nil {
	// 	log.Fatal("Error reading CA cert:", err)
	// }

	// Load CA from embed
	caCert, err := tlsutil.GetCACert()
	if err != nil {
		log.Fatalf("failed to load embedded CA cert: %v", err)
		return nil, err
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	// Load client certificate
	cert, err := tls.LoadX509KeyPair("certs/client.cert.pem", "certs/client.key.pem")
	if err != nil {
		log.Fatal("Error loading client cert/key:", err)
		return nil, err
	}

	// TLS config
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caPool,
		InsecureSkipVerify: false,                // DO NOT set to true in production!
		ServerName:         "agent.anchor.local", // must match CN or SAN of agent cert
		MinVersion:         tls.VersionTLS13,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
