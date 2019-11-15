package fields

//string struct
type StringStruct struct {
	Value      string
	Min_Length int
	Max_Length int
	Error_Min  string
	Error_Max  string
}

var String StringStruct

//string validate function
func (s *StringStruct) Validate(errorMap map[string]string) {

	if len(s.Value) < s.MinLength {
		errorMap[s.Value] = s.Error_Min
	} else if len(s.Value) > s.Max_Length {
		errorMap[s.Value] = s.Max_Length
	}

}
