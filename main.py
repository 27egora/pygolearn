def inverted(word):
    reversed_word = ''.join(reversed(word))
    return reversed_word, word == reversed_word

print(inverted("привет"))
print(inverted("топот"))

