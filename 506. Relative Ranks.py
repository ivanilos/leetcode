class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        scorePlayer = [(s, i) for i, s in enumerate(score)]

        scorePlayer.sort(reverse=True)

        result = ["" for _ in score]
        print(scorePlayer)

        for pos, entry in enumerate(scorePlayer):
            idx = entry[1]

            if pos == 0:
                result[idx] = "Gold Medal"
            elif pos == 1:
                result[idx] = "Silver Medal"
            elif pos == 2:
                result[idx] = "Bronze Medal"
            else:
                result[idx] = str(pos + 1)

        return result
