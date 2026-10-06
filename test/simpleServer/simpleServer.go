package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func handleTcp(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		msg := buf[:n]
		fmt.Printf("Reçu TCP de \"%s\" de la part de %s\n", string(msg), conn.RemoteAddr())

		reponse := bytes.ToUpper(msg)
		conn.Write(reponse)
	}

	fmt.Printf("Fin de connexion TCP pour %s\n", conn.RemoteAddr())
}

func tcpServer(listen string, port int) {
	serv, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", listen, port))

	if err != nil {
		log.Fatalf("Erreur serveur TCP")
	}
	defer serv.Close()

	fmt.Printf("Serveur TCP en écoute sur %s:%d\n", listen, port)

	for {
		conn, err := serv.Accept()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Erreur réception TCP")
			continue
		}

		fmt.Printf("Réception d'une connexion de %s\n", conn.RemoteAddr())
		go handleTcp(conn)
	}
}

func udpServer(listen string, port int) {
	conn, err := net.ListenPacket("udp4", fmt.Sprintf("%s:%d", listen, port))

	if err != nil {
		log.Fatalf("Erreur serveur UDP")
	}
	defer conn.Close()

	fmt.Printf("Serveur UDP en écoute sur %s:%d\n", listen, port)

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}

		msg := buf[:n]
		fmt.Printf("Reçu UDP de \"%s\" de la part de %s\n", string(msg), addr.String())

		reponse := bytes.ToUpper(msg)
		conn.WriteTo(reponse, addr)
	}
}

func main() {
	ip := flag.String("ip", "0.0.0.0", "IP to listen to")
	port := flag.Int("p", 8000, "Server port")

	flag.Parse()

	go udpServer(*ip, *port)
	tcpServer(*ip, *port)
}
