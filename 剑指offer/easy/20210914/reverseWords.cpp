class Solution {
public:
    string reverseWords(string s) {
        vector<string> ans;
        string str;
        for(int i=0;i<s.size();i++){
            if(s[i]!=' '){
                str+=s[i];
            }else{
                if(str!=""){
                    ans.push_back(str);
                }
                str="";
            }
        }
        if(str!=""){
            ans.push_back(str);
            str="";
        }
        for(int i=ans.size()-1;i>=0;i--){
            str+=ans[i];
            if(i!=0){
                str+=' ';
            }
        }
        return str;
    }
};
