package _0211024

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid+1) == true {
			right = mid
		} else if isBadVersion(mid+1) == false {
			left = mid + 1
		}
	}
	return left + 1
}
