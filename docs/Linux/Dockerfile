FROM archlinux:latest

echo "Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
pacman -Syy
pacman -S vim perf sysstat sysbench stress tcpdump atop htop gcc go gdb bcc bcc-tools python-bcc git
export PATH=$PATH:/usr/share/bcc/tools
export GOPATH=~/go
export PATH=~/go/bin:$PATH

# go get golang.org/x/sys/unix$
# go get github.com/tobert/pcstat/pcstat
