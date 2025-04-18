package viacep

import (
	"github.com/AndreD23/mba-goexpert/labs/00-deploy-com-cloud-run/pkg/utils"
)

type CepData struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ViaCEP struct {
	CepData
}

func GetCityByZipCode(zipCode string) (string, error) {
	url := "https://viacep.com.br/ws/" + zipCode + "/json/"
	var data ViaCEP
	err := utils.FetchData(url, &data)
	if err != nil {
		panic(err)
	}

	return data.Localidade, nil
}
