class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        v1 = list(map(int, version1.split('.')))
        v2 = list(map(int, version2.split('.')))

        size1 = len(v1)
        size2 = len(v2)
        maxSize = max(size1, size2)

        for i in range(maxSize):
            revision1 = 0
            if i < size1:
                revision1 = v1[i]

            revision2 = 0
            if i < size2:
                revision2 = v2[i]

            if revision1 < revision2:
                return -1
            elif revision1 > revision2:
                return 1

        return 0
