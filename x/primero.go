package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// Verificar la instalación de Go
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Go version:", strings.TrimSpace(string(out)))

	// Verificar los privilegios de usuario
	cmd = exec.Command("id", "-u")
	out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	if strings.TrimSpace(string(out)) == "0" {
		fmt.Println("Estás ejecutando este programa como root.")
	} else {
		fmt.Println("Estás ejecutando este programa como un usuario normal.")
	}
}
