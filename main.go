package main

import (
	"fmt" 
	"net" 
	
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar servidor.", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("[SERVIDOR] Escutando na porta 8080.")

		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao se conectar", err)
			continue
		}
		

		go handleClient(connection)
	}

}

