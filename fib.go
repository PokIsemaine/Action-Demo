package Action_Demo

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	s := make([]int, u)
	_ = s
	return Fib(u-2) + Fib(u-1)
}
