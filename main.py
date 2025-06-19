import sys

def inverted(word):
    reversed_word = ''.join(reversed(word))
    return reversed_word, word == reversed_word


if __name__ == "__main__":
    word = sys.argv[1]
    print(inverted(word))