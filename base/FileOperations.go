package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FileOperations() {
	// 1. Criar diretório
	err := os.MkdirAll("temp/subdir", 0755)
	if err != nil {
		fmt.Println("Erro ao criar diretório:", err)
		return
	}

	// 2. Criar arquivo e escrever conteúdo
	content := []byte("Hello, Go!\nEste é um arquivo de teste.")
	err = os.WriteFile("temp/test.txt", content, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever arquivo:", err)
		return
	}

	// 3. Append em arquivo existente
	file, err := os.OpenFile("temp/test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("\nLinha adicionada com append"); err != nil {
		fmt.Println("Erro ao fazer append:", err)
		return
	}

	// 4. Ler arquivo completo
	data, err := os.ReadFile("temp/test.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}
	fmt.Printf("Conteúdo do arquivo:\n%s\n", string(data))

	// 5. Ler arquivo em chunks
	file, err = os.Open("temp/test.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	chunk := make([]byte, 8) // Tamanho do chunk 8 bytes
	for {
		n, err := file.Read(chunk)
		if err == io.EOF { // Fim do arquivo
			break
		}
		if err != nil {
			fmt.Println("Erro ao ler chunk:", err)
			return
		}
		fmt.Printf("Chunk lido (%d bytes): %s\n", n, chunk[:n])
	}

	// 6. Ler arquivo linha por linha
	file, err = os.Open("temp/test.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Linha:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler linhas:", err)
		return
	}

	// 7. Listar arquivos em um diretório
	files, err := os.ReadDir("temp")
	if err != nil {
		fmt.Println("Erro ao ler diretório:", err)
		return
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Erro ao obter info do arquivo:", err)
			continue
		}
		fmt.Printf("Nome: %s, Tamanho: %d bytes, Dir?: %v\n",
			file.Name(), info.Size(), file.IsDir())
	}

	// 8. Caminhar recursivamente pelo diretório
	err = filepath.Walk("temp", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Caminho: %s, Dir?: %v\n", path, info.IsDir())
		return nil
	})
	if err != nil {
		fmt.Println("Erro ao percorrer diretório:", err)
		return
	}

	// 9. Copiar arquivo
	source, err := os.Open("temp/test.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo fonte:", err)
		return
	}
	defer source.Close()

	destination, err := os.Create("temp/test_copy.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo destino:", err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println("Erro ao copiar arquivo:", err)
		return
	}

	// 10. Renomear/Mover arquivo
	err = os.Rename("temp/test_copy.txt", "temp/test_renamed.txt")
	if err != nil {
		fmt.Println("Erro ao renomear arquivo:", err)
		return
	}

	// 11. Remover arquivo
	err = os.Remove("temp/test_renamed.txt")
	if err != nil {
		fmt.Println("Erro ao remover arquivo:", err)
		return
	}

	// 12. Remover diretório e seu conteúdo
	err = os.RemoveAll("temp")
	if err != nil {
		fmt.Println("Erro ao remover diretório:", err)
		return
	}

	// 13. Verificar se arquivo/diretório existe
	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		fmt.Println("O diretório 'temp' não existe mais")
	}
}
