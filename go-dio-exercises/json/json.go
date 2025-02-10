package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Usuarios struct{
	usuarios []Usuario `json: "usuarios"`
}

type Usuario struct{
	nome string `json: "nome"`
	tipo string `json: "tipo"`
	idade int `json: "idade"`
	social Social `json: "social"`
}

type Social struct{
	facebook string `json: "facebook"`
	twitter string `json: "twitter"`
}

func main(){
	jsonFile, err := os.Open("usuarios.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var usuarios Usuarios

	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.usuarios); i++ {
		fmt.Println(usuarios.usuarios[i].tipo)
	}
}