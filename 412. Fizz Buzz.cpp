class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> output;
        
        for (int i = 1; i <= n; i++) {
            string next_out;
            if (i % 3 == 0) next_out += "Fizz";
            if (i % 5 == 0) next_out += "Buzz";
            if (next_out == "") next_out = to_string(i);
            
            output.push_back(next_out);
        }
        return output;
    }
};