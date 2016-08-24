package calc

var workSteps []string

// AddStep adds a step in the formula computation process to the overall step list
func AddStep(step string) {
	workSteps = append(workSteps, step)
}

// GetSteps adds a step in the formula computation process to the overall step list
func GetSteps() []string {
	return workSteps
}

// ClearSteps gets step to nil
func ClearSteps() {
	workSteps = nil
}
