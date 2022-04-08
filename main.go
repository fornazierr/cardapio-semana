package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
Este aplicativo irá gerar uma sugestão de cardápios para uma semana, todos os sete dias

O cardápio contará com os seguintes itens:

- Vegetal A
- Vegetal B
- Proteina
- Carboidrato

Exemplo:

- Vegetal A: alface, tomate, agrião, rucula
- Vegetal B: cenoura
- Proteina: cogumelo
- Carboidrato: arroz integral
**/

/////////////////////////
// Variables and Structs
/////////////////////////

type Alimento struct {
	Alimento       string
	UnidadeCaseira string
	Grama          string
}

var Cereais []Alimento
var Frutas []Alimento
var Laticinio []Alimento
var Leguminosas []Alimento
var Oleaginosas []Alimento
var VegetalA []Alimento
var VegetalB []Alimento

/////////////////////////
// FUNCTIONS
/////////////////////////

func main() {
	// Carregando dados dos cereais
	carregaDados()
	listaQuantidade()
}

//Carrega dados dos arquivos JSON e os coloca em variáveis em memória
func carregaDados() {
	fmt.Println("xxxxxxxxxx\nCARREGANDO DADOS\nxxxxxxxxxx")
	// Carregando dados dos cereais
	myFile, err := os.ReadFile("data/cereais.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &Cereais)
	checkErr(err)

	// Carregando dados das frutas
	myFile, err = os.ReadFile("data/frutas.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &Frutas)
	checkErr(err)

	// Carregando dados dos laticinios
	myFile, err = os.ReadFile("data/laticinio.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &Laticinio)
	checkErr(err)

	// Carregando dados das leguminoas
	myFile, err = os.ReadFile("data/leguminosas.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &Leguminosas)
	checkErr(err)

	// Carregando dados das oleaginosas
	myFile, err = os.ReadFile("data/oleaginosas.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &Oleaginosas)
	checkErr(err)

	// Carregando dados dos vegetais A
	myFile, err = os.ReadFile("data/vegetala.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &VegetalA)
	checkErr(err)

	// Carregando dados dos vegetais B
	myFile, err = os.ReadFile("data/vegetalb.json")
	checkErr(err)
	err = json.Unmarshal(myFile, &VegetalB)
	checkErr(err)

	fmt.Println("TODOS OS DADOS CARREGADOS\nxxxxxxxxxx")
}

//Lista a quantidades de itens carregados
func listaQuantidade() {
	fmt.Println("Quandidade de Cereais:", len(Cereais))
	fmt.Println("Quandidade de Frutas:", len(Frutas))
	fmt.Println("Quandidade de Laticinios:", len(Laticinio))
	fmt.Println("Quandidade de Leguminosas:", len(Leguminosas))
	fmt.Println("Quandidade de Oleaginosas:", len(Oleaginosas))
	fmt.Println("Quandidade de Vegetal A:", len(VegetalA))
	fmt.Println("Quandidade de VegetalB:", len(VegetalB), "xxxxxxxxxx")
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
