# -*- coding: utf-8 -*-
import threading
import paramiko
import subprocess


def ssh_command(ip, user, password, command):
    client = paramiko.SSHClient()
    # client.load_host_keys('~/.ssh/known_hosts')
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(ip, username=user, password=password)
    ssh_session = client.get_transport().open_session()

    if ssh_session.active:
        ssh_session.exec_command(command)
        print ssh_session.recv(4096)

    client.close()


def ssh_command_rsa(ip, username, command):
    pkey = paramiko.RSAKey.from_private_key_file('/Users/lvtao/.ssh/id_rsa')
    ssh = paramiko.SSHClient()
    ssh.connect(hostname=ip,
                port=paramiko.client.SSH_PORT,
                username=username,
                pkey=pkey)
    stdin, stdout, stderr = ssh.exec_command(command)
    print stdout.read().decode()
    ssh.close()


def sftp(ip, username, password, local_path, remote_path):
    trans = paramiko.Transport((ip, paramiko.client.SSH_PORT))
    trans.connect(username=username, password=password)

    sftp = paramiko.SFTPClient.from_transport(trans)
    sftp.put(localpath=local_path, remotepath=remote_path)
    trans.close()


if __name__ == '__main__':
    pass
