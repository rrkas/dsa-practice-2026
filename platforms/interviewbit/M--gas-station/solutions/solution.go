/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer
 */
func canCompleteCircuit(A []int, B []int) int {
	total := 0
	tank := 0
	start := 0

	for i := 0; i < len(A); i++ {
		diff := A[i] - B[i]
		total += diff
		tank += diff

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if total >= 0 {
		return start
	}
	return -1
}
