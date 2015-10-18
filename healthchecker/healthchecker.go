package healthchecker

// Healthchecker can check healths
type Healthchecker struct {
}

// Check checks healths
func (checker *Healthchecker) Check() (bool, error) {
	return true, nil
}
