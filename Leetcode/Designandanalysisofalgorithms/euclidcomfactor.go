package Designandanalysisofalgorithms

func Euclidcomfactor(m, n int) int{
	if m < n {
		m, n = n, m
	}
	for n > 0 {
		r := m % n
		m, n = n, r
	}
	return m
}
