class Solution {
public:
    long long wonderfulSubstrings(string word) {
        long long result = 0;
        map<int, int> mapa;

        int cur_xor = 0;
        mapa[0] = 1;
        for (auto c : word) {
           cur_xor ^= 1 << (c - 'a');

            result += mapa[cur_xor];
            for (int i = 0; i < 10; i++) {
                int to_find_xor = cur_xor ^ (1 << i);
                result += mapa[to_find_xor];
            }
            mapa[cur_xor]++;
        }
        return result;
    }
};
