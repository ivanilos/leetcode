class MyCalendarThree {
private:
    vector <pair <int, int> > bookings;
    
public:
    MyCalendarThree() {
         
    }
    
    int book(int start, int end) {
        bookings.push_back(make_pair(start, 1));
        bookings.push_back(make_pair(end, -1));
        
        sort(bookings.begin(), bookings.end());
        
        int ans = 0;
        int intersectionsNow = 0;
        for (int i = 0; i < (int)bookings.size(); i++) {
            intersectionsNow += bookings[i].second;
            ans = max(ans, intersectionsNow);
        }
        return ans;
    }
};

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * MyCalendarThree obj = new MyCalendarThree();
 * int param_1 = obj.book(start,end);
 */