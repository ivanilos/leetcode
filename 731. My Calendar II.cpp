class MyCalendarTwo {
public:
    MyCalendarTwo() {
        
    }
    
    bool book(int start, int end) {
        auto it = event.lower_bound({start, -INF});
        if (it != event.end() && it->first == start) {
            pair<int, int> toInsert = {it->first, it->second + 1};
            event.erase(it);
            event.insert(toInsert);
        } else {
            event.insert({start, 1});
        }

        it = event.lower_bound({end, -INF});
        if (it != event.end() && it->first == end) {
            pair<int, int> toInsert = {it->first, it->second - 1};
            event.erase(it);
            event.insert(toInsert);
        } else {
            event.insert({end, -1});
        }

        int counter = 0;
        for (auto it = event.begin(); it != event.end(); it++) {
            counter += it->second;

            if (counter >= 3) {
                auto it = event.lower_bound({start, -INF});
                pair<int, int> toInsert = {it->first, it->second - 1};
                event.erase(it);
                event.insert(toInsert);

                it = event.lower_bound({end, -INF});
                toInsert = {it->first, it->second + 1};
                event.erase(it);
                event.insert(toInsert);
                return false;
            }
        }
        return true;
    }

private:
    set<pair<int, int>> event;
    const int INF = (int)1e9;
};

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * MyCalendarTwo* obj = new MyCalendarTwo();
 * bool param_1 = obj->book(start,end);
 */
