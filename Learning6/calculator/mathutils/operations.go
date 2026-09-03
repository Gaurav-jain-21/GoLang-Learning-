package mathutils

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) (int, string) {
	if b == 0 {
		return 0, "Error: cannot divide by Zero"
	}
	return a / b, ""
}
func secretFormula(x int) int {
	return x * x * x
}