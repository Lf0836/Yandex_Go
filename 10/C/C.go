package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	if s.solvedProblems*int(s.scoreForOneTask) >= int(s.passingScore) {
		return true
	} else {
		return false
	}
}
