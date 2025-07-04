class Solution:
    def kthCharacter(self, k: int, operations: List[int]) -> str:
        sz = 1
        op = -1
        while sz < k:
            op += 1
            sz *= 2


        print(op)
        changes = 0
        while op >= 0:
            if k > sz // 2:
                k -= sz // 2
                if operations[op] == 1:
                    changes += 1

            sz //= 2
            op -= 1

        return chr(changes % 26 + ord('a'))
