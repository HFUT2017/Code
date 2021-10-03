package _0211003

func StoneGameIX(stones []int) bool {
	used:=make([]bool,len(stones))
	var dfs func(sum int,temp []int) bool
	temp:=[]int{}
	dfs= func(sum int, temp []int) bool {
		if len(temp)==len(stones){
			return true
		}
		for i:=0;i<len(stones);i++{
			if used[i]==true||(stones[i]+sum)%3==0{
				continue
			}
			used[i]=true
			temp=append(temp,stones[i])
			dfs(sum+stones[i],temp)
			used[i]=false
			temp=temp[:len(temp)-1]
		}
		return false
	}
	if dfs(0,temp){
		if len(temp)%2==0{
			return true
		}else{
			return false
		}
	}else{
		if len(temp)%2==0{
			return false
		}else{
			return true
		}
	}
}
