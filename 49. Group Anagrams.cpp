class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> ans;
        map<string, vector<string>> mapa;
        
        for (auto& s : strs) {
            string aux = s;
            sort(aux.begin(), aux.end());
            mapa[aux].push_back(s);
        }
        for (auto& anagrams : mapa) {
            ans.push_back(anagrams.second);
        }
        return ans;
    }
};