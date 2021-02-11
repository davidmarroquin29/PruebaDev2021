package translator

import (
	"errors"
	"strings"
)

func GetTranslator (inFormat, outFormat string) (iTranslator, error){

	var formatIn format
	var formatOut format

	if strings.ToUpper(inFormat) == strings.ToUpper(outFormat){
		return nil, errors.New("In and Out Format cant be the same")
	}

	switch strings.ToUpper(inFormat) {
	case "TEXT":
		formatIn = NewFormat(nil)
	case "BINARY":
		formatIn = NewFormat(binaryChars)
	case "MORSE":
		formatIn = NewFormat(morseChars)
	default:
		return nil, errors.New("Invalid In Format")
	}

	switch strings.ToUpper(outFormat) {
	case "TEXT":
		formatOut = NewFormat(nil)
	case "BINARY":
		formatOut = NewFormat(binaryChars)
	case "MORSE":
		formatOut = NewFormat(morseChars)
	default:
		return nil, errors.New("Invalid Out Format")
	}

	translatorRet := newTranslator(formatIn, formatOut)

	return translatorRet, nil
}
