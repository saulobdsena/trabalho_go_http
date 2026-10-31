package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("[SERVIDOR] Escutando na porta 9000")

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("[Error] Não foi possivel iniciar o servidor: ", err)
		return
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("[Error] Não foi possivel aceitar a conexão: ", err)
			return
		}

		//fmt.Println("[CONEXÃO] Nova conexão de %s\n", conn.RemoteAddr())

		go tratarConexao(conn)

	}
}
func tratarConexao(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4069) //Espaço de memoria reservado para receber dados de entrada ou enviar dados

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("[Erro] Falha ao ler a requisição")
		return
	}

	req := string(buffer[:n])
	fmt.Println("----------------NOVA REQUISICAO RECEBIDA----------------")
	fmt.Println(req)
	fmt.Println("--------------------------------------------------------")

	body := "<html><head><title>My page</title></head></html>"

	header := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-type: text/html\r\nContent-length%d\r\n\r\n%s", len(body), body)
	conn.Write([]byte(header))

	fmt.Println("[RESPOSTA] ENviada com suceesso -> Servidor HTML (text/html)")
}
