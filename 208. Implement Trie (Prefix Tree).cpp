class Trie {
public:
    /** Initialize your data structure here. */
    Trie() {
        trie.push_back(map<char, int>()); // root
        word_end.push_back(false);
    }
    
    /** Inserts a word into the trie. */
    void insert(string word) {
        int pos = 0;
        for (char c : word) {
            if (!trie[pos].count(c)) {
                trie[pos][c] = trie.size();
                trie.push_back(map<char,int>());
                word_end.push_back(false);
            }
            pos = trie[pos][c];
        }
        word_end[pos] = true;
    }
    
    /** Returns if the word is in the trie. */
    bool search(string word) {
        int pos = 0;
        for (char c : word) {
            if (!trie[pos].count(c)) { return false; }
            
            pos = trie[pos][c];
        }
        return word_end[pos]; 
    }

    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        int pos = 0;
        for (char c : prefix) {
            if (!trie[pos].count(c)) { return false; }
            
            pos = trie[pos][c];
        }
        return true; 
    }
    
private:
    vector<map<char, int>> trie;
    vector<bool> word_end;
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */