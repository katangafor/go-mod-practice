package calc

import "testing"

func TestAdd(t *testing.T) {
	sum0 := Add(0)
	if sum0 != 0 {
		t.Errorf("Expected Add(0) to return 0, but got %v", sum0)
	}

	sum2 := Add(1, 2)
	if sum2 != 3 {
		t.Errorf("Expected Add(1, 2) to return 3, but got %v", sum2)
	}

	sum3 := Add(5, 10, 20)
	if sum3 != 35 {
		t.Errorf("Expected Add(5, 10, 20) to reurn 35, got %v", sum3)
	}
}
