class Solution {
public:
    char firstUniqChar(string s) {
    int c[26] = {0};
	int size = s.size();
	if (!size)return ' ';
	for (int i = 0; i < size; i++){c[s[i] - 'a']++;}
	for (int i = 0; i < size; i++)
	{
		if (c[s[i]-'a'] == 1)return s[i];
	}
	return ' ';
    }
};

