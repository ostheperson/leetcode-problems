def checkInclusion(s1: str, s2: str) -> bool:
    """
    create hash from s1 with char as key and count as value
    set left, right
    get a char: 
    if you have not seen char,
        add to hash with value 1
        set left to that index
    if it has been seen,
        if value of char in hash is at least one,
            then deduct hash[char] and progress right 
            return true if len(s2[left,right]) == len(s1)
        else:
            set left to index
            add char to hash
    return false
    """
    is_sub_set = False
    count = 0
    for i in s2:
        if i in s1: 
            if is_sub_set == True:
                count += 1
            else:
                is_sub_set = True
                count = 1
            if count == len(s1): return True
        else: is_sub_set = False
    return False

print(checkInclusion("hello", "ooleh"))