from flask import Flask, request, jsonify 

app = Flask(__name__)

def inverted(word): #переворот строки
    reversed_word = ''.join(reversed(word))
    return reversed_word, word == reversed_word

@app.route('/', methods=['PUT'])
def put_handler():
    input_text = request.get_json().get('text', '') #забираем значение ключа text из тела запроса
    reversed_text, is_palindrome = inverted(input_text) #инвертирование и проверка
    response = { # JSON ответ
        "text": reversed_text,
        "isPalindrome": is_palindrome
    }
    return jsonify(response)

@app.route('/', methods=['GET'])
def get_handler():
    text = request.form.get('text', '')
    reversed_text, is_palindrome = inverted(text)
    text_color = "red" if is_palindrome else "green"
    html_reponse = f'<h1 style = "color: {text_color}">{reversed_text}</h1>'
    return html_reponse

@app.route('/', methods=['POST'])
def post_handler():
    text = request.form.get('text', '')
    reversed_text, is_palindrome = inverted(text)
    text_color = "blue" if is_palindrome else "red"
    html_reponse = f'<h1 style = "color: {text_color}">{reversed_text}</h1>'
    return html_reponse

    

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)