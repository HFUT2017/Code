package _0210911

func LengthOfLIS_1(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func lengthOfLIS(nums []int) int {
//	//O(nlogn)
//	m := len(nums)
//	tail := make([]int, m+1) //长度为i的最长子序列结尾字符
//	end := 0                 //tail 结尾
//	tail[0] = math.MinInt32  //长度0则没有
//	dp := make([]int, m)     //i字符结尾最长子序列值1
//	for i := 1; i < m; i++ {
//		dp[i] = 1
//	}
//
//	for i := 0; i < m; i++ {
//		if nums[i] > tail[end] {
//			end++
//			tail[end] = nums[i] //记录长度end的最小字典序的结尾
//			dp[i] = end
//		} else {
//			//二分找到大于等于元素下标
//			left, right := 1, end
//			for left < right {
//				mid := left + (right-left)/2
//				if tail[mid] < nums[i] {
//					left = mid + 1
//				} else {
//					right = mid
//				}
//			}
//			tail[left] = nums[i] //因为有更小的,所以要更新前面长度为i的结尾字符
//			dp[i] = left         //i结尾最大长度
//		}
//	}
//
//	//找具体字典序最小的序列
//	arr := make([]int, end) //长度end,实际需要end-1下标
//	size := end
//	for i := len(nums) - 1; i >= 0; i-- {
//		if dp[i] == size {
//			arr[size-1] = nums[i]
//			size--
//		}
//	}
//	fmt.Println(arr)
//	return end
//}
