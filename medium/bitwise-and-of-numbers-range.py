def rangeBitwiseAnd(left: int, right: int) -> int:
    counter = 0
    while left != right:
        left >>= 1
        right >>= 1
        counter += 1
    left <<= counter
    return left

print(rangeBitwiseAnd(4, 7))
