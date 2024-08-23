class Solution:
    def fractionAddition(self, expression: str) -> str:
        fracs = self.parse(expression)

        result = self.solve(fracs)

        return self.getAnswer(result)

    def getAnswer(self, frac):
        return str(frac[0]) + '/' + str(frac[1])

    def solve(self, fracs):
        result = [0, 1]

        for frac in fracs:
            lcm = self.getLCM(result[1], frac[1])

            result[0] = result[0] * lcm // result[1] + frac[0] * lcm // frac[1]
            result[1] = lcm

            gcd = self.getGCD(result[0], result[1])
            result[0] //= gcd
            result[1] //= gcd

            if result[0] == 0:
                result[1] = 1

        return result

    def getLCM(self, a, b):
        return (a // self.getGCD(a, b)) * b

    def getGCD(self, a, b):
        if b == 0:
            return a
        return self.getGCD(b, a % b)

    def parse(self, expression: str):
        curSign = 1
        frac = [0, 0]
        fracs = []
        nextPart = 0

        for ch in expression:
            if ch == '+':
                if frac[0] != 0 and frac[1] != 0:
                    frac[0] *= curSign
                    fracs.append(frac)
                    frac = [0, 0]

                curSign = 1
                nextPart = 0
            elif ch == '-':
                if frac[0] != 0 and frac[1] != 0:
                    frac[0] *= curSign
                    fracs.append(frac)
                    frac = [0, 0]

                curSign = -1
                nextPart = 0
            elif ch == '/':
                nextPart = 1
            else:
                frac[nextPart] = 10 * frac[nextPart] + ord(ch) - ord('0')

        if frac[0] != 0 and frac[1] != 0:
            frac[0] *= curSign
            fracs.append(frac)
            frac = [0, 0]

        return fracs
