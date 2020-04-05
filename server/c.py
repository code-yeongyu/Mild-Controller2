from socket import *
import json

port = 888

clientSock = socket(AF_INET, SOCK_STREAM)
clientSock.connect(('127.0.0.1', port))

print('접속 완료')

while True:
    data = clientSock.recv(1024)
    msg = json.loads(data.decode('utf-8'))
    if not msg['is_shutdown']:
        clientSock.send(
            (f"yeah kinda works, the '{msg['msg']}' command").encode('utf-8'))
