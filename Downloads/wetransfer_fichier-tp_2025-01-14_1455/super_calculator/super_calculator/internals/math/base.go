package math

func Addition(x, y float64) float64 {
	return x + y
}

func Multiplication(x, y float64) float64 {
	return x * y
}

func Subtraction(x, y float64) float64 {
	return x - y
}

func Division(x, y float64) float64 {
	return x / y
}
func TestDivisionByZero(t *testing.T) {
    result, err := Divide(10, 0)
    if err == nil {
        t.Errorf("Expected an error for division by zero, got %v", result)
    }
}
