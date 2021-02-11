package main

import (
	translator2 "PruebaDev/translator"
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Translator")

	var inText string
	fmt.Println("Ingrese texto a traducir:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		inText = scanner.Text()
	}


	var inFormat string
	fmt.Println("Ingrese Formato Origen (TEXT, MORSE, BINARY):")
	if scanner.Scan(){
		inFormat = scanner.Text()
	}

	var outFormat string
	fmt.Println("Ingrese Formato Salida (TEXT, MORSE, BINARY):")
	if scanner.Scan(){
		outFormat = scanner.Text()
	}


	//TEST
	//inText := "test"
	//inFormat := "TEXT"
	//outFormat := "BINARY"

	translt, err := translator2.GetTranslator(inFormat, outFormat)
	if err != nil{
		fmt.Println("Ocurrió un error:")
		fmt.Println(err)
	}else{
		outText, err := translt.Translate(inText)
		fmt.Println("El resultado es:")
		fmt.Println(outText)

		if err != nil{
			fmt.Println("Ocurrió un error:")
			fmt.Println(err)
		}
	}


}