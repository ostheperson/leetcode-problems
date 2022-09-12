def reverseWords(s: str) -> str:
    # split
    string_arr = s.split(' ')

    for i in range(len(string_arr)):
        string_arr[i] = string_arr[i][::-1]

    return ' '.join(string_arr)

print (reverseWords('method returns a string'))
