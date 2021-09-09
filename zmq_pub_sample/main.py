import time
import zmq


def main():
    context = zmq.Context()
    socket = context.socket(zmq.PUB)
    socket.bind(f'tcp://*:64230')

    while True:
        socket.send(b'pub')
        print('pub')
        time.sleep(1)

    socket.close()
    context.term()


if __name__ == '__main__':
    main()
