import socket

BYTES = 2

'''
Sends a message by first sending its length, then the message, encoded in UTF-8
'''
def send_message(sock: socket, msg: str):
    data = msg.encode("utf-8")
    sock.sendall(len(data).to_bytes(BYTES, "big"))
    sock.sendall(data)

'''
Receives a message by first receiving its length, then the message, which is decoded
'''
def receive_message(sock: socket) -> str:
    length = int.from_bytes(sock.recv(BYTES), byteorder="big")
    msg = b""
    while len(msg) < length:
        recv_data = sock.recv(length - len(msg))
        msg += recv_data
    return msg.decode("utf-8")
