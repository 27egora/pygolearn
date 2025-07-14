from flask import Flask, request, jsonify 

app = Flask(__name__)

def inverted(word): #переворот строки
    uppercase = word.upper()
    reversed_word = uppercase[::-1] #переворот нормализ. строчки 
    return reversed_word, uppercase == reversed_word

@app.route('/', methods=['PUT'])
def put_handler():
    input_text = request.get_json().get('text', '') #забираем значение ключа text из тела запроса
    reversed_text, is_palindrome = inverted(input_text) #инвертирование и проверка
    response = { # JSON ответ
        "text": reversed_text,
        "isPalindrome": is_palindrome
    }
    status_code = 202 if is_palindrome else 200 #правильные коды статусов (статус http-ответа)
    return jsonify(response), status_code 

@app.route('/', methods=['GET'])
def get_handler():
    text = request.args.get('text', '')
    reversed_text, is_palindrome = inverted(text)
    text_color = "red" if is_palindrome else "green"
    html_reponse = f'<h1 style = "color: {text_color}">{reversed_text}</h1>'
    status_code = 202 if is_palindrome else 200 #правильные коды статусов (статус http-ответа)
    return html_reponse, status_code

@app.route('/', methods=['POST'])
def post_handler():
    text = request.form.get('text', '')
    reversed_text, is_palindrome = inverted(text)
    text_color = "blue" if is_palindrome else "red"
    html_reponse = f'<h1 style = "color: {text_color}">{reversed_text}</h1>'
    status_code = 202 if is_palindrome else 200 #правильные коды статусов (статус http-ответа)
    return html_reponse, status_code

    

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)