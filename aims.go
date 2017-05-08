package main

type Aim int

const (
	A_START = 0
	A_FIT
	A_MUSCLE
	A_WEIGHT
	A_MONEY
	A_PIE
)

func (a Aim) String() string {
	switch a {
	case 0:
		return "Starting Out"
	case 1:
		return "Get Fit"
	case 2:
		return "Build Muscle"
	case 3:
		return "Lose Weight"
	case 4:
		return "To Lose Money"
	case 5:
		return "To eat my weight in Pie"
	}
	return "No Aim"
}
