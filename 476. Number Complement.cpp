class Solution {
public:
    int findComplement(int num) {
        int mask = ~getMask(num);

        return (~num) ^ (mask);
    }

    int getMask(int num) {
        long long result = 1;
        while(result <= num) {
            result *= 2;
        }
        return int(result - 1);
    }
};
