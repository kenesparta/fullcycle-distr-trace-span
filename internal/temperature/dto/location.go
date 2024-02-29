package dto

type LocationInput struct {
	CEP string `json:"cep"`
}

type LocationOut struct {
	CEP         string      `json:"cep"`
	Logradouro  string      `json:"logradouro"`
	Complemento string      `json:"complemento"`
	Bairro      string      `json:"bairro"`
	Localidade  string      `json:"localidade"`
	Uf          string      `json:"uf"`
	Ibge        string      `json:"ibge"`
	Gia         string      `json:"gia"`
	Ddd         string      `json:"ddd"`
	Siafi       string      `json:"siafi"`
	ErrorOut    interface{} `json:"erro"`
}
