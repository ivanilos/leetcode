struct Task {
    int id;
    int userId;
    int priority;
};

bool operator <(const Task& a, const Task& b) {
    return a.priority > b.priority || (a.priority == b.priority && a.id > b.id);
}

class TaskManager {
public:
    TaskManager(vector<vector<int>>& tasks) {
        for (auto task : tasks) {
            int userId = task[0];
            int taskId = task[1];
            int priority = task[2];

            taskIdToPriority[taskId] = priority;
            taskIdToUser[taskId] = userId;
            tasksSet.insert(Task{taskId, userId, priority});
        }
    }
    
    void add(int userId, int taskId, int priority) {
        taskIdToPriority[taskId] = priority;
        taskIdToUser[taskId] = userId;
        tasksSet.insert(Task{taskId, userId, priority});
    }
    
    void edit(int taskId, int newPriority) {
        int curPriority = taskIdToPriority[taskId];
        int userId = taskIdToUser[taskId];
        auto it = tasksSet.find(Task{taskId, userId, curPriority});

        tasksSet.erase(it);

        add(userId, taskId, newPriority);
    }
    
    void rmv(int taskId) {
        int curPriority = taskIdToPriority[taskId];
        int userId = taskIdToUser[taskId];

        taskIdToPriority.erase(taskId);
        taskIdToUser.erase(taskId);

        auto it = tasksSet.find(Task{taskId, userId, curPriority});

        tasksSet.erase(it);
    }
    
    int execTop() {
        if (tasksSet.size() == 0) {
            return -1;
        }

        auto it = tasksSet.begin();
        int userId = it->userId;

        tasksSet.erase(it);

        return userId;
    }
private:
    set<Task> tasksSet;
    map<int, int> taskIdToUser;
    map<int, int> taskIdToPriority;
};

/**
 * Your TaskManager object will be instantiated and called as such:
 * TaskManager* obj = new TaskManager(tasks);
 * obj->add(userId,taskId,priority);
 * obj->edit(taskId,newPriority);
 * obj->rmv(taskId);
 * int param_4 = obj->execTop();
 */
