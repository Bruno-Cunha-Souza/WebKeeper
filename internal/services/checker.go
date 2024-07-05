package services

import (
	"fmt"
	"net/http"
	"time"
)

const monitoramentos = 2
const delay = 5

func StartMonit() {
	fmt.Println("Monitorando...")
	sites := "http://httpstat.us/random/200,201,500-504"

	for i := 0; i < monitoramentos; i++ {
		for j, site := range sites {
			fmt.Println("Testando site", j, ":", site)
			testSites(sites)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}
func testSites(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

// func readSiteList() []string {
// 	var sites []string
// 	arquivo, err := os.Open("sites.txt")

// 	if err != nil {
// 		fmt.Println("Ocorreu um erro:", err)
// 	}

// 	leitor := bufio.NewReader(arquivo)
// 	for {
// 		linha, err := leitor.ReadString('\n')
// 		linha = strings.TrimSpace(linha)

// 		sites = append(sites, linha)

// 		if err == io.EOF {
// 			break
// 		}

// 	}

// 	arquivo.Close()
// 	return sites
// }
