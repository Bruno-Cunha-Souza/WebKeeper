package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
)

const delay = 2

func StartMonit() {
	for {
		var sites []models.Site
		result := database.DB.Find(&sites)

		if result.Error != nil {
			fmt.Println("Erro ao buscar sites:", result.Error)
			return
		}
		for _, site := range sites {
			fmt.Printf("Testando site: %s\n", site.Nome)
			testSite(site.URL)
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ocorreu um erro ao conectar ao URL:", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		fmt.Println("URL:", url, "\nSuccessful connecting. Status Code:", resp.StatusCode)
	case 400:
		fmt.Println("URL:", url, "\nBad Request. Status Code:", resp.StatusCode)
	case 401:
		fmt.Println("URL:", url, "\nUnauthorized. Status Code:", resp.StatusCode)
	case 404:
		fmt.Println("URL:", url, "\nNot Found. Status Code:", resp.StatusCode)
	case 500:
		fmt.Println("URL:", url, "\nInternal Server Error. Status Code:", resp.StatusCode)
	case 502:
		fmt.Println("URL:", url, "\nBad Gateway. Status Code:", resp.StatusCode)
	case 504:
		fmt.Println("URL:", url, "\nGateway Timeout. Status Code:", resp.StatusCode)
	default:
		fmt.Println("URL:", url, "\nError connecting. Status Code:", resp.StatusCode)
	}
}
