"""
- get index
- get shift right by k
- switch positon with item in shifted place
[1,2,3,4,5,6,7]

[7,1,2,3,4,5,6]

[6,7,1,2,3,4,5]
"""

def reverse(arr, lo, hi):
    while lo <= hi:
        arr[lo], arr[hi] = arr[hi], arr[lo]
        lo, hi = lo + 1, hi - 1


def rotate(nums, k: int):
    k = k%len(nums)
    lo, hi = 0, len(nums)-1
    reverse(nums, lo, hi)

    lo, hi = 0, k-1
    reverse(nums, lo, hi)

    lo, hi = k, len(nums)-1
    reverse(nums, lo, hi)

    return nums[-k:]


print (rotate([1,2,3,4,5,6,7], 3))

