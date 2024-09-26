class MyCalendar {
public:
    MyCalendar() {
        events.clear();
    }
    
    bool book(int start, int end) {
        auto it = events.lower_bound({start + 1, -1});
        if (it != events.end() && intersect(it->second, it->first, start, end)) {
            return false;
        }
        events.insert({end, start});
        return true;
    }

private:
    set<pair<int, int>> events;

    bool intersect(int s1, int e1, int s2, int e2) {
        e1--;
        e2--;
        return s2 <= s1 && s1 <= e2 || s2 <= e1 && e1 <= e2 ||
                s1 <= s2 && s2 <= e1 || s1 <= e2 && e2 <= e1;
    }
};

/**
 * Your MyCalendar object will be instantiated and called as such:
 * MyCalendar* obj = new MyCalendar();
 * bool param_1 = obj->book(start,end);
 */
