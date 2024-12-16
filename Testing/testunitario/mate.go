package testunitario

func Suma(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonnacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnacci(n-1) + Fibonnacci(n-2)
}
