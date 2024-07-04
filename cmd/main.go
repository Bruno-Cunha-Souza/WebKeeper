package main

import (
	"fmt"
	"os"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/services"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/utils"
)

func main() {
	utils.ShowIntro()

	for {
		utils.ShowMenu()

		comando := utils.ReadComand()

		switch comando {
		case 1:
			services.StartMonit()
		case 2:
			fmt.Println("Exibindo Logs...")
			services.PrintLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}
