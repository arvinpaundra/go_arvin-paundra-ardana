package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(4, 3) != 7 {
		t.Error("Expected 4 + 3 equal to 7")
	}

	if Addition(10, 20) != 30 {
		t.Error("Expected 10 + 20 equal to 30")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(5, 3) != 2 {
		t.Error("Expected 5 - 3 equal to 2")
	}

	if Subtraction(10, 100) != -90 {
		t.Error("Expected 10 - 100 equal to -90")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 3) != 6 {
		t.Error("Expected 2 * 3 equal to 6")
	}
	if Multiplication(4, 4) != 16 {
		t.Error("Expected 4 * 4 equal to 16")
	}
}

func TestDivsion(t *testing.T) {
	if Division(4, 2) != 2 {
		t.Error("Expected 4 / 2 equal to 2")
	}

	if Division(0, 100) != 0 {
		t.Error("Expected 0 / 100 equal to 0")
	}
}
