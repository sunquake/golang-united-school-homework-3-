package homework

func reverse(input []int64) (result []int64) {
	l := len(input)
	if l == 0 {return}
	for i := l-1; i >= 0; i-- {
		result = append(result, input[i])
	}
	return
}
