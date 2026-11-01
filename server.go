
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func handleClient(connection net.Conn) {
	fmt.Println("[CONEXÃO]: Nova conexão de ", connection.RemoteAddr())

	defer func() {
    connection.Close()
    fmt.Println("[CONEXÃO]: Encerrada ", connection.RemoteAddr())
	}()

	reader := bufio.NewReader(connection)

	
	linha, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a requisição:", err)
		return
	}

	requestLine := strings.TrimSpace(linha)
	partes := strings.Fields(requestLine)

	if len(partes) != 3 {
		fmt.Println("Erro na linha da requisição")
		return
	}

	var caminho string
	metodo := partes[0]
	if partes[1] == "/"{
		caminho = "/index.html"
	} else{caminho = partes[1]}
	

	
	var contentStr string

	for {
		linha, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler os headers:", err)
			return
		}

		
		if linha == "\r\n" {
			break
		}

		headersPartes := strings.SplitN(linha, ":", 2)

		nome := strings.TrimSpace(headersPartes[0])
		valor := strings.TrimSpace(headersPartes[1])

		if nome == "Content-Length" {
			contentStr = valor
		}
	}

	
	contentInt := 0

	if contentStr != "" {
		contentInt, err = strconv.Atoi(contentStr)

		if err != nil {
			fmt.Println("Erro na conversao do content")
			return
		}
	}

	
	bodyBytes := make([]byte, contentInt)

	if contentInt > 0 {
		_, err = io.ReadFull(reader, bodyBytes)

		if err != nil {
			fmt.Println("Erro ao ler o body")
			return
		}
	}

	// acho q saulo vai usar essa parte
	// body := string(bodyBytes)

	
	fmt.Printf("[REQUEST] Método: %s | Path: %s | Content-Lengh: %d\n",
		metodo, caminho, contentInt)
}

func getContentType(caminho string) string {
    extensao := strings.ToLower(filepath.Ext(caminho))

    switch extensao {
    case ".html":
        return "text/html"
    case ".css":
        return "text/css"
    case ".png":
        return "image/png"
    case ".jpg":
        return "image/jpeg"
    default:
        return "application/octet-stream"
    }
}