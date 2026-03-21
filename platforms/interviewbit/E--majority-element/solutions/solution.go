/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func majorityElement(A []int) int {
	n := len(A)
	freq := make(map[int]int)
	for _, e := range A {
		freq[e]++
		if freq[e] > n/2 {
			return e
		}
	}
	return 0
}
