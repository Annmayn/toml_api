package fields

//IntStruct : struct of int data type
type IntStruct struct {
	minValue   int
	maxValue   int
	rangeValue []int
	errMin     string
	errMax     string
	errRange   string
}

//Validate : validates int structure
func (in *IntStruct) Validate(n int, field string, errorMap map[string]string) {
	if n < in.minValue {
		errorMap[field] = in.errMin

	} else if n > in.maxValue {
		errorMap[field] = in.errMax

	} else if n < in.rangeValue[0] || n > in.rangeValue[1] {
		errorMap[field] = in.errRange
	}
}
