class Solution {
    public int distributeCandies(int[] candies) {
        Map<Integer, Integer> candyQuantity = new HashMap<Integer, Integer>();
        
        int unique = 0;
        for (int i = 0; i < candies.length; i++) {
            if (!candyQuantity.containsKey(candies[i])) {
                unique++;
            }
            candyQuantity.put(candies[i], 1);
        }
        return Math.min(unique, candies.length / 2);
    }
}