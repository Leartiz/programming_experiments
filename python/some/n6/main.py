from flask import Flask, render_template
from flask_socketio import SocketIO, emit
    
app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret!'
socketio = SocketIO(app)

@app.route('/')
def index():
    return render_template('index.html')

req_count = 0
@socketio.event
def my_event(message):
    print("GOT IT!")
    global req_count
    req_count += 1
    emit('my_response', {'data': message['data'], 'count': req_count})

if __name__ == '__main__':
    socketio.run(app)