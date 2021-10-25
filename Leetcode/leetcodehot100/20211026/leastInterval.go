package _0211026

func leastInterval(tasks []byte, n int) int {
	hash:=map[byte]int{}
	for _,value:=range tasks{
		hash[value]++
	}
	maxExec,maxExecCount:=0,0
	for _,count:=range hash{
		if count>maxExec{
			maxExec,maxExecCount=count,1
		}else if count==maxExec{
			maxExecCount++
		}
	}
	if time:=(maxExec-1)*(n+1)+maxExecCount;time>len(tasks){
		return time
	}
	return len(tasks)
}
