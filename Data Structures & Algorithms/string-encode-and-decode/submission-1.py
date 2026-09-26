class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            res = res + str(len(s))+"*"+s
        
        print(res)
        return res

    def decode(self, s: str) -> List[str]:
        res = []

        while s != "":
            l = ""
            i = 0
            while i < len(s) and s[i] != '*':
                l += s[i]
                i += 1

            s = s[i+1:]
            res.append(s[:int(l)])
            s = s[int(l):]
        return res
        