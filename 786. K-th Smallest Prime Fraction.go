class Solution:
    def kthSmallestPrimeFraction(self, arr: List[int], k: int) -> List[int]:
        frac = []

        for i in range(len(arr)):
            for j in range(i + 1, len(arr)):
                frac.append((arr[i], arr[j]))

        frac.sort(key=lambda a: a[0] / a[1])

        return frac[k - 1]
