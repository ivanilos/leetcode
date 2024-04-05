class Solution:
    def makeGood(self, s: str) -> str:
        stack = []

        for c in s:
            if len(stack) != 0:
                if c != stack[-1] and c.lower() == stack[-1].lower():
                    stack.pop()
                else:
                    stack.append(c)
            else:
                stack.append(c)

        return ''.join(stack)
