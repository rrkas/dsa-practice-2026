package kata

func Multiple3And5(number int) int {
	s := 0

	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			s += i
		}
	}

	return s
}
