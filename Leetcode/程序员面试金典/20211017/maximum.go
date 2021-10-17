package _0211017

func maximum(a int, b int) int {
	ret := int64(a - b)
	ret = int64(a) - ret&(ret>>63)
	return int(ret)
}
