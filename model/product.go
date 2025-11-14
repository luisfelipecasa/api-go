package model

//model de produto
//o entre crase diz qual sera o nome retornado na requisição
type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
