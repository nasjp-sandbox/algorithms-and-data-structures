package part4

import (
	"strconv"
)

func Stack(in []string) string {
	s := newStack()

	for _, v := range in {
		if isNum(v) {
			s.push(v)

			continue
		}

		last := s.pop()

		beforeLast := s.pop()

		result := evaluate(beforeLast, last, v)

		s.push(result)
	}

	return s.pop()
}

func isNum(v string) bool {
	if _, err := strconv.Atoi(v); err != nil {
		return false
	}

	return true
}

func evaluate(first string, second string, operand string) string {
	fi, _ := strconv.Atoi(first)

	si, _ := strconv.Atoi(second)

	switch operand {
	case "+":
		return strconv.Itoa(fi + si)
	case "-":
		return strconv.Itoa(fi - si)
	case "*":
		return strconv.Itoa(fi * si)
	default:
		return ""
	}
}

type stack struct {
	s []string
}

func newStack() *stack {
	return &stack{s: make([]string, 0)}
}

func (s *stack) push(x string) {
	s.s = append(s.s, x)
}

func (s *stack) pop() string {
	var pop string

	pop, s.s = s.s[len(s.s)-1], s.s[:len(s.s)-1]

	return pop
}
