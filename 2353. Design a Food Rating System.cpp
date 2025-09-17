struct Food {
    string name;
    int rating;
};

bool operator <(const Food& a, const Food& b) {
    return a.rating > b.rating || (a.rating == b.rating && a.name < b.name);
}

class FoodRatings {
public:
    FoodRatings(vector<string>& foods, vector<string>& cuisines, vector<int>& ratings) {
        for (int i = 0; i < foods.size(); i++) {
            foodToRating[foods[i]] = ratings[i];
            foodToCuisine[foods[i]] = cuisines[i];
            cuisineToFood[cuisines[i]].insert(Food{foods[i], ratings[i]});
        }
    }
    
    void changeRating(string food, int newRating) {
        int curRating = foodToRating[food];
        string cuisine = foodToCuisine[food];

        auto it = cuisineToFood[cuisine].find(Food{food, curRating});
        cuisineToFood[cuisine].erase(it);

        foodToRating[food] = newRating;
        cuisineToFood[cuisine].insert(Food{food, newRating});
    }
    
    string highestRated(string cuisine) {
        return cuisineToFood[cuisine].begin()->name;
    }
private:
    map<string, int> foodToRating;
    map<string, string> foodToCuisine;
    map<string, set<Food>> cuisineToFood;
};

/**
 * Your FoodRatings object will be instantiated and called as such:
 * FoodRatings* obj = new FoodRatings(foods, cuisines, ratings);
 * obj->changeRating(food,newRating);
 * string param_2 = obj->highestRated(cuisine);
 */
