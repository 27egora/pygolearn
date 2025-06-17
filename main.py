def inverted(word):
    is_palindrome = word == word[::-1]
    return ''.join(reversed(word)), is_palindrome

print(inverted("привет"))
print(inverted("топот"))
