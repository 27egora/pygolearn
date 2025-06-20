import sys

def inverted(word):
    reversed_word = ''.join(reversed(word))
    return reversed_word, word == reversed_word



if __name__ == "__main__":
    if len(sys.argv) > 1:
        word = sys.argv[1]
        reversed_word, is_palindrome = inverted(word)
        print(f"Исходное слово: {word}") 
        print(f"Перевернутое: {reversed_word}")
        print(f"Палиндром: {is_palindrome}")
    else:
        print("Указать ОДНО слово в качестве аргумента!")