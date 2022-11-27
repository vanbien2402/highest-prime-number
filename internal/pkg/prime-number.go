package pkg

//HighestPrime calculate the highest prime number lower than N
func HighestPrime(N int) int {
	if N < 3 {
		return 0
	}
	isNotPrime := make([]bool, N)
	for i := 2; i*i < N; i++ {
		if !isNotPrime[i] {
			for j := i * i; j < N; j += i {
				isNotPrime[j] = true
			}
		}
	}
	var highestPrime int
	for i := N - 1; i > 1; i-- {
		if !isNotPrime[i] {
			highestPrime = i
			break
		}
	}
	return highestPrime
}
