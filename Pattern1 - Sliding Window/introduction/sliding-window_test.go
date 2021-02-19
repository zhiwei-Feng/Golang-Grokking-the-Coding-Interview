package introduction

import "testing"

func TestFindAveragesOfSubArraysBySlidingWindow(t *testing.T) {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	results := findAveragesOfSubArraysBySlidingWindow(5, arr)
	expected := []float64{2.2, 2.8, 2.4, 3.6, 2.8}
	if len(results) != len(expected) {
		t.Errorf("Averages of subarrays of size K(5): %d, answer is %f, expect is %f", arr, results, expected)
	} else {
		eq := true
		for i := 0; i < len(results); i++ {
			if results[i] != expected[i] {
				eq = false
				break
			}
		}
		if !eq {
			t.Errorf("Averages of subarrays of size K(5): %d, answer is %f, expect is %f", arr, results, expected)
		}
	}
}
