package main

import (
	"fmt"
	translator2 "pruebadev/translator"
)

func main(){
	fmt.Println("Translator")

	var inText string
	fmt.Println("Ingrese texto a traducir:")
	fmt.Scan(&inText)

	var inFormat string
	fmt.Println("Ingrese Formato Origen:")
	fmt.Scan(&inFormat)

	var outFormat string
	fmt.Println("Ingrese Formato Salida:")
	fmt.Scan(&outFormat)

	translator, err := translator2.GetTranslator(inFormat, outFormat)
	if err != nil{
		fmt.Println("Ocurrió un error:")
		fmt.Println(err)
	}else{
		outText, err := translator.Translate(inText)
		fmt.Println("El resultado es:")
		fmt.Println(outText)

		if err != nil{
			fmt.Println("Ocurrió un error:")
			fmt.Println(err)
		}
	}


}