package translator

type translator struct {
	inFormat format
	outFormat format
}

func newTranslator(inFormat format, outFormat format) translator{
	return translator{
		inFormat: inFormat,
		outFormat: outFormat,
	}
}

func (t translator) Translate(input string)(string, error){

	text, _ := t.inFormat.translateToText(input)

	result, _ := t.outFormat.translateFromText(text)

	return result, nil
}