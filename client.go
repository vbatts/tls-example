package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"
)

var (
	flCert    = flag.String("cert", "cert2.pem", "certificate")
	flKey     = flag.String("key", "cert2.key", "certificate")
	flMinutes = flag.Int("t", 6, "minutes to chat to the server")
)

func main() {
	flag.Parse()

	cert2_b, _ := ioutil.ReadFile(*flCert)
	priv2_b, _ := ioutil.ReadFile(*flKey)
	priv2, _ := x509.ParsePKCS1PrivateKey(priv2_b)

	cert := tls.Certificate{
		Certificate: [][]byte{cert2_b},
		PrivateKey:  priv2,
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "127.0.0.1:11111", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		fmt.Println(v.Subject)
	}
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	// ping pong timeout
	quit := time.After(time.Duration(*flMinutes) * time.Minute)
	for {
		message := "Hello\n"
		n, err := io.WriteString(conn, message)
		if err != nil {
			log.Fatalf("client: write: %s", err)
		}
		log.Printf("client: wrote %q (%d bytes)", message, n)

		reply := make([]byte, 256)
		n, err = conn.Read(reply)
		log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
		select {
		case <-quit:
			break
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	log.Print("client: exiting")
}
