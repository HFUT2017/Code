package _0211012

func merge(A []int, m int, B []int, n int) {
	l1, l2, tail := m-1, n-1, m+n-1
	for l1 >= 0 || l2 >= 0 {
		if l1 < 0 {
			A[tail] = B[l2]
			l2--
		} else if l2 < 0 {
			A[tail] = A[l1]
			l1--
		} else if A[l1] < B[l2] {
			A[tail] = B[l2]
			l2--
		} else {
			A[tail] = A[l1]
			l1--
		}
		tail--
	}
}
