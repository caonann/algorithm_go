package arrary

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}

	for _, test := range tests {
		if got := twoSum(test.numbers, test.target); got[0] != test.want[0] || got[1] != test.want[1] {
			t.Errorf("6.twoSum(%v ,%v)=%v", test.numbers, test.target, test.want)
		}
		if got := twoSum2(test.numbers, test.target); got[0] != test.want[0] || got[1] != test.want[1] {
			t.Errorf("6.twoSum2(%v ,%v)=%v", test.numbers, test.target, test.want)
		}
	}

}
