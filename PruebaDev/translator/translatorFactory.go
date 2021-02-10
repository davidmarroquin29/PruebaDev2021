package translator

import (
	"errors"
)

func GetTranslator (inFormat, outFormat string) (iTranslator, error){

	var formatIn format
	var formatOut format

	switch inFormat {
	case "TEXT":
		formatIn = NewFormat(nil)
	case "BINARY":
		formatIn = NewFormat(binaryChars)
	case "MORSE":
		formatIn = NewFormat(morseChars)
	default:
		return nil, errors.New("Invalid In Format")
	}

	switch outFormat {
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
