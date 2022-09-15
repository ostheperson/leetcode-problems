
def bagOfTokensScore(tokens, power) -> int:
    score, res = 0, 0
    left, right = 0, len(tokens) - 1
    while left <= right:
        if power >= tokens[left]:
            power -= tokens[left]
            score += 1
            res = max(score, res)
            left += 1
        elif score > 0:
            power += tokens[right]
            score -= 1
            right -= 1
        else: break
    return res

print (bagOfTokensScore([1,4,5,6], 4))



