package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Defina o diretório o qual você quer organizar
var dirOrigem string = "/Users/ions/Downloads" // Troque o diretório

func main() {
	// Verificar se o diretório existe, caso contrário, retornar erro
	if _, err := os.Stat(dirOrigem); os.IsNotExist(err) {
		fmt.Println("BAD DIR :( \nDiretório não encontrado: ", dirOrigem)
		return
	} else {
		// Imprimir mensagem caso o diretório exista
		fmt.Println("GOOD DIR :) \nDiretório encontrado: ", dirOrigem)
	}

	// Percorrer e listar os arquivos no diretório dirOrigem
	err := filepath.Walk(dirOrigem, organizarArquivos)
	if err != nil {
		fmt.Println("Erro ao percorrer o diretório: ", err)
	}
}

// Função que organiza os arquivos do diretório
func organizarArquivos(caminho string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// Ignorar diretórios e exibir apenas arquivos
	if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
		fmt.Println("Arquivo encontrado: ", info.Name())

		// Obter extensões em letras minúsculas
		extensao := strings.ToLower(filepath.Ext(info.Name()))
		// Criar nome das subpastas diretamente em dirOrigem
		subpasta := filepath.Join(dirOrigem, extensao[1:])

		// Verificar se a subpasta já existe; caso contrário, criar pastas
		if _, err := os.Stat(subpasta); os.IsNotExist(err) {
			err := os.Mkdir(subpasta, os.ModePerm)
			if err != nil {
				fmt.Println("Erro ao criar subpasta: ", err)
				return err
			}
		}

		// Caminho de destino para mover arquivos
		caminhoDestino := filepath.Join(subpasta, info.Name())

		// Verificar se o arquivo já está na subpasta
		if caminhoDestino == caminho {
			fmt.Printf("O arquivo %s já está na subpasta %s. Ignorando...\n", info.Name(), subpasta)
		} else {
			// Mover o arquivo para a subpasta
			err := os.Rename(caminho, caminhoDestino)
			if err != nil {
				fmt.Println("Erro ao mover arquivo: ", err)
				return err
			}
			fmt.Printf("Arquivo %s movido para %s\n", info.Name(), subpasta)
		}
	}
	return nil
}
