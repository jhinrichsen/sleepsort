package sleepsort

import "testing"

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestSleepSort(t *testing.T) {
	want := []int{2, 4, 6, 8}
	got := SleepSort([]int{8, 4, 6, 2})
	if !equal(want, got) {
		t.Fatalf("Want %+v but got %+v\n", want, got)
	}
}
