package translator

type iFormat interface {
	translateToText(input string) (string, error)
	translateFromText(input string)(string, error)
}

type format struct{
	charList map[string]string
}

func NewFormat (charList map[string]string) format{
	return format{charList: charList}
}

func (f format) translateToText(input string) (string, error){
	text := decodeMessage(input, "")
	return text, nil
}

func (f format) translateFromText(input string) (string, error){
	res := encodeMessage(input, "", f.charList)
	return res, nil
}