import re
from flask import Flask, request, jsonify, abort

app = Flask(__name__)

# Регулярное выражение для валидации
TEXT_REGEX = re.compile(r'^[а-яА-ЯёЁa-zA-Z]+$')

def inverted(word):
    """Переворот строки и проверка на палиндром"""
    reversed_word = word[::-1]
    return reversed_word, word.lower() == reversed_word.lower()

@app.route('/', methods=['PUT'])
def put_handler():
    """Обработка PUT-запроса"""
    input_text = request.get_json().get('text', '') #забираем значение ключа text из тела запроса
    reversed_text, is_palindrome = inverted(input_text) #инвертирование и проверка
    response = { # JSON ответ
        "text": reversed_text,
        "isPalindrome": is_palindrome
    }
    return jsonify(response)
    

@app.route('/', methods=['GET', 'POST'])
def web_handler():
    """Обработка GET и POST запросов"""
    text = request.args.get('text', '') if request.method == 'GET' else request.form.get('text', '')
    
    if not text:
        abort(400, "Текст не может быть пустым")
    if len(text) > 50:
        abort(400, "Текст слишком длинный (максимум 50 символов)")
    if not TEXT_REGEX.match(text):
        abort(400, "Текст может содержать только буквы и пробелы")
    
    reversed_text, is_palindrome = inverted(text)
    
    # Выбор цвета
    if request.method == 'GET':
        text_color = "red" if is_palindrome else "green"
    else:  # POST
        text_color = "blue" if is_palindrome else "red"
    
    return f'<h1 style="color: {text_color}">{reversed_text}</h1>'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)

