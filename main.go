package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println("Digite oque deseja substituir:")
	var replace string
	fmt.Scanln(&replace)

	fmt.Println("Digite oque deseja substiuir por ", replace, " ou aperter ENTER para substituir por vazio (apagar)")
	var replaceTo string
	fmt.Scanln(&replaceTo)

	dir := "./arquivos"
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Define o novo nome para o arquivo
		newName := strings.Replace(info.Name(), replace, replaceTo, -1)
		newPath := filepath.Join(filepath.Dir(path), newName)

		fmt.Println(newPath)

		// Renomeia o arquivo
		if err := os.Rename(path, newPath); err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		panic(err)
	}

}
