package translator

import (
	"strings"
)

func encodeMessage(message, letterSplitter string, chars map[string]string) string {
	var output string
	message = strings.ToUpper(message)
	words := strings.Split(message, " ")
	for _, word := range words {
		word := encodeWord(word, letterSplitter, chars)
		if word != "" {
			output += word
		}
	}
	return output
}

func encodeWord(word, letterSplitter string, chars map[string]string) string {
	var res string
	for i := 0; i < len(word); i++ {
		code := chars[word[i:i+1]]
		if code != "" {
			res += code + letterSplitter
		}
	}
	return res
}

func decodeLetter(code string) string {
	return characterMap[code]
}

func decodeMessage(message, letterSplitter string) string {
	var decodedMessage string
	var letters = strings.Split(message, letterSplitter)
	for _, letter := range letters {
		decodedMessage += decodeLetter(strings.Trim(letter, " ")) + " "
	}
	return decodedMessage
}