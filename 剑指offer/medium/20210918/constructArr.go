package _0210918

// ConstructArr 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法
func ConstructArr(a []int) []int {
	b := make([]int,len(a))
	if len(a) == 0 {
		return b
	}
	b[0] =1
	tmp :=1
	for i:=1;i<len(a);i++{
		b[i] = b[i-1] * a[i-1]
	}
	for j:= len(a)-2;j>=0;j--{
		tmp *= a[j + 1];
		b[j] *= tmp;
	}
	return b
}
