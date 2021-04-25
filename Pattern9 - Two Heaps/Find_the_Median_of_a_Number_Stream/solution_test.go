package Find_the_Median_of_a_Number_Stream

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	medianOfAStream := Constructor()
	medianOfAStream.AddNum(3)
	medianOfAStream.AddNum(1)
	if medianOfAStream.FindMedian() != 2.0 {
		t.Errorf("want 2.0, but got %f", medianOfAStream.FindMedian())
	}
	medianOfAStream.AddNum(5)
	if medianOfAStream.FindMedian() != 3.0 {
		t.Errorf("want 3.0, but got %f", medianOfAStream.FindMedian())
	}
	medianOfAStream.AddNum(4)
	if medianOfAStream.FindMedian() != 3.5 {
		t.Errorf("want 3.5, but got %f", medianOfAStream.FindMedian())
	}
}
