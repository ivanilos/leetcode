struct Movie {
    int shop;
    int id;
    int price;

    Movie() {}
    Movie(int shop, int id, int price):shop(shop), id(id), price(price) {}
};

bool operator< (const Movie& a, const Movie& b) {
    return a.price < b.price || (a.price == b.price && a.shop < b.shop) || 
    (a.price == b.price && a.shop == b.shop && a.id < b.id);
}

class MovieRentingSystem {
public:
    MovieRentingSystem(int n, vector<vector<int>>& entries) {
        for (auto entry : entries) {
            int shop = entry[0];
            int id = entry[1];
            int price = entry[2];

            Movie movie = Movie(shop, id, price);

            movieToMovieInShop[id].insert(movie);
            shopToMovieIdMap[shop][id] = movie;
        }
    }
    
    vector<int> search(int movie) {
        vector<int> result;

        for (auto it = movieToMovieInShop[movie].begin(); it != movieToMovieInShop[movie].end() && result.size() < 5; it++) {
            result.push_back(it->shop);
        }
        return result;
    }
    
    void rent(int shop, int movie) {
        auto it = shopToMovieIdMap[shop][movie];

        rented.insert(it);
        movieToMovieInShop[movie].erase(it);
    }
    
    void drop(int shop, int movie) {
        auto it = shopToMovieIdMap[shop][movie];

        rented.erase(it);
        movieToMovieInShop[movie].insert(it);
    }
    
    vector<vector<int>> report() {
        vector<vector<int>> result;

        for (auto it = rented.begin(); it != rented.end() && result.size() < 5; it++) {
            vector<int> elem = vector<int>{it->shop, it->id};
            result.push_back(elem);
        }

        return result;
    }

private:
    set<Movie> rented;
    map<int, set<Movie>> movieToMovieInShop;
    map<int, map<int, Movie>> shopToMovieIdMap;
};

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * MovieRentingSystem* obj = new MovieRentingSystem(n, entries);
 * vector<int> param_1 = obj->search(movie);
 * obj->rent(shop,movie);
 * obj->drop(shop,movie);
 * vector<vector<int>> param_4 = obj->report();
 */
