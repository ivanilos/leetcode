class Solution {
public:
    int reverse(int x) {
        string number = to_string(x);
        while(number.back() == '0' && number.size() != 1) {
            number.pop_back();
        }
        if (number[0] == '-') {
            std::reverse(number.begin() + 1, number.end());
        } else {
            std::reverse(number.begin(), number.end());
        }
        try {
            int val = stoi(number);
            return val;
        } catch (std::out_of_range) {
            return 0;
        }
    }
};