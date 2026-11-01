package main

import (
	"bufio"
	"fmt"
	"strings"
)

func parseRequestLine(linha string) (string, string, string, error) {
	partes := strings.Fields(linha)

	if len(partes) != 3 {
		return "", "", "", fmt.Errorf("linha de requisição inválida")
	}

	metodo := partes[0]
	caminho := partes[1]
	versao := partes[2]

	return metodo, caminho, versao, nil
}

func parseHeaders(reader *bufio.Reader) (map[string]string, error) {
	headers := make(map[string]string)

	for {
		linha, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		linha = strings.TrimRight(linha, "\r\n")

		// Linha vazia indica o fim dos cabeçalhos
		if linha == "" {
			break
		}

		partes := strings.SplitN(linha, ":", 2)

		if len(partes) != 2 {
			return nil, fmt.Errorf("cabeçalho inválido")
		}

		chave := strings.ToLower(strings.TrimSpace(partes[0]))
		valor := strings.TrimSpace(partes[1])

		headers[chave] = valor
	}

	return headers, nil
}