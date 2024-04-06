class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        toRemoveIdx = set()
        stack = []

        for idx, c in enumerate(s):
            if c == '(':
                stack.append((c, idx))
            elif c == ')':
                if len(stack) == 0:
                    toRemoveIdx.add(idx)
                else:
                    stack.pop()

        while len(stack) > 0:
            toRemoveIdx.add(stack[-1][1])
            stack.pop()

        result = ""
        for idx, c in enumerate(s):
            if idx not in toRemoveIdx:
                result += c

        return result
