package main

import "testing"

type CharacterJumpTest struct {
	stepInput   int
	hurdleInput []int
	expected    bool
}

var tables = []CharacterJumpTest{
	{3, []int{0, 1, 0, 0, 0, 1, 0}, true},
	{4, []int{0, 1, 1, 0, 1, 0, 0, 0, 0}, false},
}

func TestCharacterJump(t *testing.T) {
	for _, test := range tables {
		min := CharacterJump(test.stepInput, test.hurdleInput)
		if min != test.expected {
			t.Errorf("CharacterJump was incorrect, got: %t, want: %t.", min, test.expected)
		}
	}
}
