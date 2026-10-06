package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

type Protocol string

const (
	TCP Protocol = "tcp4"
	UDP Protocol = "udp4"
)

func send(ip string, port int, proto Protocol) {
	// On fait la connexion.
	conn, err := net.Dial(string(proto), fmt.Sprintf("%s:%d", ip, port))

	if err != nil {
		log.Fatalf("Connexion %s", string(proto))
	}
	defer conn.Close()

	conn.Write([]byte("Message de test"))
}

// Envoi d'un message TCP.
func tcp(ip string, port int) {
	send(ip, port, TCP)
}

// Envoi d'un message UDP.
func udp(ip string, port int) {
	send(ip, port, UDP)
}

func main() {
	// Flags cmd.
	ip := flag.String("ip", "localhost", "server address")
	port := flag.Int("p", 8000, "server port")

	flag.Parse()

	tcp(*ip, *port)
	udp(*ip, *port)
}
