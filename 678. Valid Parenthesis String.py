class Solution:
    def checkValidString(self, s: str) -> bool:
        stack = []
        extra = []

        for idx, c in enumerate(s):
            if c == '(':
                stack.append(idx)
            elif c == ')':
                if len(stack) > 0:
                    stack.pop()
                elif len(extra) > 0:
                    extra = extra[1::]
                else:
                    return False
            else:
                extra.append(idx)

        while len(stack) > 0:
            if len(extra) > 0 and extra[-1] > stack[-1]:
                stack.pop()
                extra.pop()
            else:
                break

        return len(stack) == 0
