package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	// agora a terceira coluna seria o mapeamento do JSON
	//Nome  tipo     mapeamento JSON
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1999.99, []string{"Promocao", "Eletronico"}}

	// funcao json.Marshal retorna o json e o erro
	p1Json, erro := json.Marshal(p1) // retorna um slice de bytes

	if erro != nil {
		panic(erro)
	}

	fmt.Println(string(p1Json))
	// {"id":1,"nome":"Notebook","preco":1999.99,"tags":["Promocao","Eletronico"]}

	var p2 produto
	jsonString := `{"id":2,"nome":"TV","preco":1999.99,"tags":["Promocao","Eletronico"]}`
	json.Unmarshal([]byte(jsonString), &p2) // passar slice de bytes e dps referencia
	// referencia pois será alterado o atributo do obj

	fmt.Println(p2) // {2 TV 1999.99 [Promocao Eletronico]}

	fmt.Println(p2.Tags[0]) // Promocao
}
