// Package smoke is a tiny Go module whose only purpose is to exercise the
// reusable Go workflow in this repository on every push.
package smoke

// Greet returns a greeting for name, or for the world when name is empty.
func Greet(name string) string {
	if name == "" {
		name = "world"
	}

	return "hello, " + name
}
