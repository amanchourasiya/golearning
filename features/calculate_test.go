package features

import "testing"

func TestAddWithPositiveValues(t *testing.T) {
	sum := Add(12, 23)

	if sum == 35 {
		t.Logf("Add success with positive values")
	} else {
		t.Errorf("Add failure with positive values")
	}
}

func TestAddWithNegativeValues(t *testing.T) {
	sum := Add(-23, -12)

	if sum == -35 {
		t.Log("Add success with negative values")
	} else {
		t.Error("Add failure with positive values")
	}
}
