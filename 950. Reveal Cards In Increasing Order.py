class Solution:
    def deckRevealedIncreasing(self, deck: List[int]) -> List[int]:
        deck = sorted(deck, reverse=True)

        result = deque()

        for newCard in deck:
            result.appendleft(newCard)
            oldCard = result.pop()
            result.appendleft(oldCard)

        oldCard = result.popleft()
        result.append(oldCard)

        return list(result)
