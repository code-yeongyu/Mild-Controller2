import socket
from threading import Thread
import json


def menu(client):
    msg = {'is_shutdown': False, 'msg': ""}
    value = input('''"t": Turn off the computer
"r": Run specific command
''')
    if value == 't':
        msg['is_shutdown'] = True
        send_command(client, msg)
    elif value == 'r':
        msg['msg'] = input("Type your command.\n")
        send_command(client, msg)
        data = client.recv(65000)
        print(data.decode("utf-8"))


def send_command(client_socket, command):
    client_socket.send(json.dumps(command).encode('utf-8'))


def main():
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind(('', 1818))
    server_socket.listen(0)

    print("Server ON")

    while True:
        try:
            menu(client_socket)
        except:
            client_socket, addr = server_socket.accept()
            print(f"{addr}의 접속")


main()