package parse

import (
	"fmt"
	"os"
)

func OpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Erro ao abrir arquivo %v", err)
		return nil, err
	}
	return file, nil
}
