func findTheWinner(n int, k int) int {
    // taken from e-maxx josephus page
	result := 0;
	for i := 1; i <= n; i++ {
		result = (result + k) % i
    }
	return result + 1;
}
