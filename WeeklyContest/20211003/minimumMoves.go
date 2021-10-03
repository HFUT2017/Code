package _0211003

func minimumMoves(s string) int {
	res:=0
	str:=[]byte(s)
	for i:=0;i<len(s);i++{
		if str[i]=='X'{
			res+=1
			for j:=i;j<=i+2&&j<len(s);j++{
				str[j]=0
			}
		}
	}
	return res
}
