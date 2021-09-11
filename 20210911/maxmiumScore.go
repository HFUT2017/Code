package _0210911

// MaxmiumScore 给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
//func MaxmiumScore(cards []int, cnt int) int {
//	sort.Ints(cards)
//	sum := 0
//	oushu := make([]int, len(cards))
//	jishu := make([]int, len(cards))
//	index_o := 0
//	index_j := 0
//	for i := 0; i <len(cards); i++ {
//		if cards[i]%2 == 0 {
//			oushu[index_o] = cards[i]
//			index_o++
//		} else {
//			jishu[index_j] = cards[i]
//			index_j++
//		}
//	}
//	index_o--
//	index_j--
//	for cnt != 0 {
//		if cnt >= 2 {
//			if index_j >= 1 && index_o >= 1 {
//				if jishu[index_j]+jishu[index_j-1] >= oushu[index_o]+oushu[index_o-1] {
//					sum += jishu[index_j] + jishu[index_j-1]
//					index_j -= 2
//					cnt -= 2
//				} else {
//					sum += oushu[index_o]
//					index_o -= 1
//					cnt -= 1
//				}
//			} else if index_o >= 1 {
//				sum += oushu[index_o]
//				index_o--
//				cnt--
//			} else if index_j >= 1 {
//				sum += jishu[index_j] + jishu[index_j-1]
//				index_j -= 2
//				cnt -= 2
//			} else {
//				return 0
//			}
//		} else {
//			if index_o >= 0 {
//				sum += oushu[index_o]
//				return sum
//			} else {
//				return 0
//			}
//		}
//	}
//	return 0
//}
