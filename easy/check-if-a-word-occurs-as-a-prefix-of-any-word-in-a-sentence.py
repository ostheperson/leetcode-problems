# Check If a Word Occurs As a Prefix of Any Word in a Sentence
def isPrefixOfWord(sentence: str, searchWord: str) -> int:
    sentence = sentence.split(' ')
    for i in range(len(sentence)):
        if len(sentence[i]) >= len(searchWord):
            print(sentence[i][0:(len(searchWord) - 1)])
            if sentence[i][0:(len(searchWord))] == searchWord:
                return i
    return -1
# print (isPrefixOfWord('love errichto jonathan dumb', 'dumb'))