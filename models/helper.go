package models

func GetExpectedTime(a, m, b int) int {
	return (a + (4 * m) + b) / 6
}
