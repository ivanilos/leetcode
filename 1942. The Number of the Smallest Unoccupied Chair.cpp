struct Event {
    int id;
    int time;
    bool isArrival;
};

bool operator <(const Event& a, const Event& b) {
    return a.time < b.time || a.time == b.time && a.isArrival < b.isArrival;
}

class Solution {
public:
    int smallestChair(vector<vector<int>>& times, int targetFriend) {
        vector<Event> events;

        for (int i = 0; i < times.size(); i++) {
            events.push_back(Event{i, times[i][0], true});
            events.push_back(Event{i, times[i][1], false});
        }

        sort(events.begin(), events.end());

        set<int> freeSeat;
        map<int, int> idToSeat;
        for (int i = 0; i < times.size(); i++) {
            freeSeat.insert(i);
        }

        for (int i = 0; i < events.size(); i++) {
            if (events[i].id == targetFriend) {
                return *(freeSeat.begin());
            } else if (events[i].isArrival) {
                auto it = freeSeat.begin();
                idToSeat[events[i].id] = *it;
                freeSeat.erase(it);
            } else {
                int availableSeat = idToSeat[events[i].id];
                freeSeat.insert(availableSeat);
            }
        }

        return -1;
    }
};
