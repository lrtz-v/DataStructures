FROM archlinux:latest

echo "Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
pacman -Syy
pacman -S vim perf sysstat sysbench stress tcpdump atop htop gcc go gdb

