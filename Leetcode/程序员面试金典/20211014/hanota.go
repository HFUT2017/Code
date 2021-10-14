package _0211014

func hanota(A []int, B []int, C []int) []int {
	move(len(A), &A, &B, &C)
	return C
}

func move(n int, A *[]int, B *[]int, C *[]int) {
	if n <= 0 {
		return
	}
	move(n-1, A, C, B)
	*C = append(*C, (*A)[len(*A)-1])
	*A = (*A)[:len(*A)-1]
	move(n-1, B, A, C)
}
