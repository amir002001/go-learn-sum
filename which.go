package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Which() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("boooo gimme exactly one arg")
		return
	}
	program_name := arguments[1]
	env_path := os.Getenv("PATH")
	env_path_array := strings.Split(env_path, ":")
	for _, path := range env_path_array {
		program_alleged_path := filepath.Join(path, program_name)
		path_status, err := os.Stat(program_alleged_path)
		if err != nil {
			continue
		}
		path_mode := path_status.Mode()
		if path_mode.IsRegular() && path_mode&0111 != 0 {
			fmt.Println(program_alleged_path)
			return
		}
	}
	fmt.Println("not found")
}
