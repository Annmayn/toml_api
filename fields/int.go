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

//string struct
type StringStruct struct {
	Value      string
	Min_Length int
	Max_Length int
	Error_Min  string
	Error_Max  string
}

//---------------------------Function definition----------------------

//string validate function
func (s *StringStruct) Validate(errorMap map[string]string) {

	if len(s.Value) < s.MinLength {
		errorMap[s.Value] = s.Error_Min
	} else if len(s.Value) > s.Max_Length {
		errorMap[s.Value] = s.Max_Length
	}

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
