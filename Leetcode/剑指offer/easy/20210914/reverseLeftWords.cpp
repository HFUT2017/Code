class Solution {
public:
    string reverseLeftWords(string s, int n) {
        string str1,str2;
        str1=s.substr(0,n);
        str2=s.substr(n,s.size());
        return str2+str1;
    }
};
