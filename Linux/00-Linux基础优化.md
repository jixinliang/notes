## 1. 修改网卡配置文件

```bash
TYPE="Ethernet"
BOOTPROTO="none"
NAME="ens33"
DEVICE="ens33"
ONBOOT="yes"
IPADDR="10.0.0.13"
PREFIX="24"
GATEWAY="10.0.0.2"
DNS1="114.114.114.114"
```

## 2. 系统的优化

#### 防火墙的优化

```bash
systemctl stop firewalld.service
systemctl disable firewalld.service
systemctl status firewalld.service
```

#### `selinux` 的优化

```bash
# close selinux
cp /etc/selinux/config{,.ori}
sed -i 's/^SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
setenforce 0
cat /etc/selinux/config
```

#### `ssh` 的优化

```bash
cp /etc/ssh/sshd_config{,.ori}
vim /etc/ssh/sshd_config
87: GSSAPIAuthentication no
124: UseDNS no
```

#### `hosts` 优化

```bash
10.0.0.11 controller
10.0.0.31 computer
```

#### 修改主机名

#### `yum` 源优化

```bash
umount /mnt
cd /etc/yum.repos.d
mkdir test -p
\mv *.repo test
cat > local.repo <<EOF
[local]
name=local
baseurl=file:///mnt
gpgcheck=0
EOF
mount /dev/cdrom /mnt
yum makecache
```

#### 其他优化(CentOS 8 关闭上不了网)

```bash
systemctl stop NetworkManager.service 
systemctl disable NetworkManager.service 

# 清除内核显示
> /etc/issue
> /etc/issue.net

# sudo提权
echo "jack    ALL=(ALL)       NOPASSWD: ALL" >> /etc/sudoers

# 添加alias
echo "alias cl='clear'" >> /etc/bashrc
. /etc/bashrc
```

#### 常用的包

```bash
dnf install net-tools lrzsz wget tree screen lsof tcpdump bash-completion.noarch
```

## `epel8源`

```bash
# dnf upgrade https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm
```

