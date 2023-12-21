package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.Contains(info.Name(), "co") {
		fmt.Printf("%d ", len(info.Name())+1)
		return nil
	} else {
		return nil
	}
}

func main() {
	const root = "."
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
