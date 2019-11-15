type IntStruct struct{
	value int
	min_value int
	max_value int
	range_value []int 
	err_min string
	err_max string
	err_range string
}
var Int Intstruct

func (*in IntStruct) Validate(errorMap map[string]string){
	if in.value < in.min_value{
		errorMap[in.value] = in.err_min
	}else if in.value > in.max_value{
		errorMap[in.value] = in.err_max
	}else if in.value < n.range_value[0] || n.value > n.range_value[1]{
		errorMap[in.value] = in.err_range
	}
}