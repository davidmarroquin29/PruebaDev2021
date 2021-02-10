package translator


type iTranslator interface{
	Translate(input string) (string, error)
}

