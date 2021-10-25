package _0211026

func removeInvalidParentheses(s string) []string {
	var check func(s string) bool
	check=func(s string) bool{
		left:=0
		for _,value:=range s{
			if value=='('{
				left++
			}else if value==')'{
				left--
			}
			if left<0{
				return false
			}
		}
		return left==0
	}
	str:=map[string]bool{s:true}
	res:=[]string{}
	for len(str)>0{
		for value,_:=range str{
			if check(value)==true{
				res=append(res,value)
			}
		}
		if len(res)>0{
			return res
		}
		temp:=map[string]bool{}
		for value, _ := range str {
			for i := 0; i < len(value); i++ {
				if value[i] == '(' || value[i] == ')' {
					new := []byte{}
					new = append(new, []byte(value)[:i]...)
					new = append(new, []byte(value)[i+1:]...)
					temp[string(new)] = true
				}
			}
		}
		str=temp
	}
	return res
}
