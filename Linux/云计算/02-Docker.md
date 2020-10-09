# 下载`Docker` 源

```bash
sudo wget -O /etc/yum.repos.d/docker-ce.repo https://mirror.tuna.tsinghua.edu.cn/docker-ce/linux/centos/docker-ce.repo

sudo curl https://mirror.tuna.tsinghua.edu.cn/docker-ce/linux/centos/docker-ce.repo -o /etc/yum.repos.d/docker-ce.repo
```

# 安装软件

```bash
dnf install docker-ce
```

# 开启 `Docker` 服务

```bash
systemctl enable --now docker.service

sudo usermod -aG docker change_to_your_user_name
```



# 官方仓库地址

## [docker Hub](https://hub.docker.com/)



# 加速器地址

```bash
# 官方源:	https://registry.docker-cn.com

https://reg-mirror.qiniu.com

cat > /etc/docker/daemon.json <<EOF
{
"registry-mirrors": ["http://hub-mirror.c.163.com"]

}
EOF

# 用于永久设置,重启docker后,自动重启container
	"live-restore": true

# 配置本地镜像服务
"insecure-registries": ["10.0.0.13:5000"]

systemctl daemon-reload
systemctl restart docker
```



# 常用命令

```bash
docker version
docker system info => docker info

docker search [names]

systemctl restart docker
systemctl daemon-reload

yum makecache fast

```



## 镜像`docker images`的基础操作

```bash
# 查看镜像
docker image ls

# 查看完整的容器完整ID
docker image ls --no-trunc
docker image inspect ubtuntu
docker image inspect [img id]

# 镜像的导入和导出
docker image save [img-name|id] > /path-name/img-name.tgz

docker iamge load -i /path-name/img-name.tgz
    load	--Load an image from a tar archive or STDIN
    -i 	--input string   Read from tar archive file, instead of STDIN

# 手动修改image的tag名
docker image tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
	Usage:	docker image tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

# 删除image
docker image rm [img_name|id]
	-f, --force      Force removal of the image

docker image rm -f `docker ps -aq`
```



## 容器`docker container`的操作

```bash

# 交互式界面
docker run	- Run a command in a new container
	-d, --detach=true|false	Detached mode: run the container in the background and 
		print the new container ID. The default is false.
	-p, --publish ip:[hostPort]:containerPort | [hostPort:]containerPort
		Publish a container's port, or range of ports, to the host.
	-P, --publish-all=true|false
		Publish all exposed ports to random ports on the host interfaces.
       	The default is false.
    -i, --interactive=true|false
        Keep STDIN open even if not attached. The default is false.
    -t, --tty=true|false
         Allocate a pseudo-TTY. The default is false.
    -v|--volume[=[[HOST-DIR:]CONTAINER-DIR[:OPTIONS]]]
		Create a bind mount. If you specify, -v /HOST-DIR:/CONTAINER-DIR, Docker


	docker run -d -p 80:80 nginx
	
# 交互式启动
	docker run -it --name="centos1" centos
	
	# 一次性(测试场景,临时使用)的容器
	docker run -it --name="centos2" --rm centos2
	
	iptables ens33:0 10.0.0.8/24 up
	
# 非交互(守护进程)式启动
	# 端口映射
	1.指定ip/port
	docker run -d -p 10.0.0.8:80:80 nginx
	2.指定ip/随机宿主机port
	docker run -d -p 10.0.0.8::80 nginx # 随机宿主机端口
	3.随机ip/port
	docker run -d -P nginx # 随机端口(32768-60999)
	4.指定ip/port/protocol(指定协议要使用/)
	docker run -d -p 10.0.0.8:80:80/udp nginx
	5.使用HTTP/HTTPS多个端口
	docker run -d -p 80:80 -p 443:443 nginx
	


# 容器的连接
docker container attach
	docker container attach [OPTIONS] CONTAINER


# 子进程方式登录
docker container exec
	docker container exec [OPTIONS] CONTAINER COMMAND [ARG...]
	-i, --interactive	Keep STDIN open even if not attached
	-t, --tty	Allocate a pseudo-TTY

	docker container exec -it [id/name] /bin/bash


# 后台运行
1. ctrl +pq

2.死循环

3.让前台一直夯着

docker container ls -a == docker ps -a

# 容器的启动(交互式可选)
docker container start
	docker container start [OPTIONS] CONTAINER [CONTAINER...]
	-i, --interactive	Attach container's STDIN

# 关闭容器
docker container stop
docker container stop [OPTIONS] CONTAINER [CONTAINER...]


docker stop containerName
docker kill containerName


docker exec	- Run a command in a running container
	-d, --detach[=false]	Detached mode: run command in the background
	-i, --interactive[=false]	 Keep STDIN open even if not attached
	-t, --tty[=false]	Allocate a pseudo-TTY
	
	docker exec -it 34212f381ba5 /bin/bash

# 批量删除容器
docker rm -f `docker ps -aq`

# 批量删除已对退出的容器
docker rm -f `docker ps -qf status=exited`

docker rm 
	- Remove one or more containers
	-f, --force[=false]


iptables -t nat -L -n
sysctl -a|grep ipv4|grep range

# 数据卷操作
docker run -d --name="web01" -p 10.0.0.13:80:80 nginx

# 容器的数据复制
	docker container cp web01:/usr/share/nginx/html/index.html /srv/html/

# 数据持久化
	docker run -d --name="web01" -p 80:80 -v /srv/html:/usr/share/nginx/html nginx
	# 数据卷名: test docker volume ls 查看
	docker run -d -p 80:80 -v test:/usr/share/nginx/html nginx

# 在集中化管理集群中,大批量的容器都需要挂载相同的多个数据卷时,可以使用使用数据卷容器进行统一管理
1.宿主机魔衣数据目录
mkdir -p /srv/volumeX{web01,web02}
touch /srv/volumeX/web01/a.txt
touch /srv/volumeX/web02/b.txt

2.启动数据卷容器
docker run -it --name="nginx_volumes" -v /srv/volumeX/web01:/srv/web01 -v /srv/volumeX/web02:/srv/web02 centos

3.是用数据卷容器
docker run -d --name="web8080" -p 8080:80 --volumes-from nginx_volumes nginx
docker run -d --name="web8081" -p 8081:80 --volumes-from nginx_volumes nginx

# 搭建本地源
1.下载所需工具包
dnf install vsftpd

路径在/var/ftp/

测试ftp路径工具包
dnf install lftp ?
lftp 10.0.0.13
pwd

2.上传系统镜像
3.配置仓库路径
mkdir -p /var/ftp/contos8.2

cat >/etc/yum.repos.d/my-ftp.repo <<- EOF
[my-ftp]
name=Local Packages for Enterprise Linux
#baseurl=https://10.0.0.13/centos8.2
enabled=1
gpgcheck=0
EOP

dnf makecache fast

4.挂载镜像
mount -o loop /mnt/CentOS8.2.iso /ver/ftp/centos8.2

 /mnt/c/Users/jack/Documents/Linux\ Mirrors/CentOS-8.2.2004-x86_64-dvd1.iso

5.访问
ftp://10.0.0.13/centos8.2

# 过滤字段
docker container inspect -f "{{.Args}}" web01
```



# 基于容器制作镜像

```bash
docker container commit

Usage:	docker container commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

Create a new image from a container's changes

Options:
  -a, --author string    Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")
  -c, --change list      Apply Dockerfile instruction to the created image
  -m, --message string   Commit message
  -p, --pause            Pause container during commit (default true)


mkdir -p /srv/vol/mysql /srv/vol/html

docker run -it --name="web_bbs" -v /srv/vol/mysql:/var/lib/mysql -v /srv/vol/html:/usr/share/nginx/html ubuntu




```



# `Dockerfile`

```dockerfile
[root@me ubuntu_ssh]# cat Dockerfile 
# Nginx-v1.0
FROM nginx

# RUN 关键字后面跟的指令的两种格式
1.
RUN echo "<h1>Hello, Docker!</h1>" > /usr/share/nginx/html/index.html
2.
["mysqld", "--initialize-insecure", "--user=mysql", "--basedir=/usr/local/mysql", "--datadir=/data/mysql/data"]

# 暴露port
EXPOSE 22
EXPOSE 80

# 用于设置守护进程
CMD ["/usr/sbin/sshd", "-D"]

docker build -t nginx:v1 .


# ubuntu
FROM ubuntu

# 支持正则表达式
# 从dockerfile同级目录下拷贝文件到容器中; 如果拷贝目录,则只拷贝目录下的子文件
COPY init.sh /srv/
# 使用add命令会自动解压本地上传的文件
ADD bbs.tar.gz /tmp
# add命令也支持上传网络上的文件,但是不能解压
ADD https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/20.04/ubuntu-20.04-live-server-amd64.list /tmp

EXPOSE 22
EXPOSE 80
EXPOSE 3306
# 容器启动时运行的第一个进程;如果命令行有要执行的命令,如/bin/bash,该命令会被覆盖
CMD ["/bin/bash", "/srv/init.sh"] 
# 使用entrypoint当前命令就不会被覆盖了
ENTRYPOINT ["/bin/bash", "/srv/init.sh"]

vim init.sh
systemctl enable --now mysqld
systemctl enable --now nginx
/usr/sbin/sshd -D

chmod +x init.sh

VOLUME ["/srv/vol/html/", "/usr/share/nginx/html/"]

WORKDIR 相当于cd

ENV HTML_DIR="/usr/share/nginx/html/"
ENV DATA_DIR="/srv/data/mysql/data/"

COPY index.html ${HTML_DIR}



```

# 制作registry

```bash
docker pull registry
docker run -d -p 5000:5000 --restart always --name registry registry
-v /srv/vol/registry:/var/lib/registry

# 添加到配置文件 >> /etc/docker/daemon.json
"insecure-registries": ["10.0.0.13:5000"] 

# server
docker pull ubuntu
docker tag ubuntu:latest 10.0.0.13:5000/me/ubuntu:v1
docker push 10.0.0.13:5000/me/ubuntu:v1

# client
docker pull 10.0.0.13:5000/me/ubuntu:v1

# 本地仓库安全认证

dnf install httpd-tools -y
mkdir /srv/vol/registry-auth -p
htpasswd -Bbn jack 666 > /srv/vol/registry-auth/htpasswd

docker run -d -p 5000:5000 --restart always --name registry-auth -v /srv/vol/registry-auth/:/opt/auth/ -v /srv/vol/registry:/var/lib/registry -e "REGISTRY_AUTH=htpasswd" -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" -e "REGISTRY_AUTH_HTPASSWD_PATH=/opt/auth/htpasswd" registry



[root@docker01 ~]# docker push 10.0.0.13:5000/me/ubuntu
The push refers to repository [10.0.0.13:5000/me/ubuntu]
05f3b67ed530: Preparing 
ec1817c93e7c: Preparing 
9e97312b63ff: Preparing 
e1c75a5e0bfa: Preparing 
no basic auth credentials

# 认证
docker login 10.0.0.13:5000

# server
[root@docker01 ~]# docker login 10.0.0.13:5000
Username: jack
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

# client
[root@docker02 ~]# docker pull 10.0.0.13:5000/me/ubuntu:v1
Error response from daemon: Get http://10.0.0.13:5000/v2/me/ubuntu/manifests/v1: no basic auth credentials



wget http://storage.googleapis.com/harbor-releases/release-2.0.0/harbor-offline-installer-v2.0.1.tgz



docker tag SOURCE_IMAGE[:TAG] 10.0.0.13/me/REPOSITORY[:TAG]

docker push 10.0.0.13/me/REPOSITORY[:TAG]

# 推送是的授权验证
[root@docker01 ~]# docker push 10.0.0.13/me/ubuntu:v2 
The push refers to repository [10.0.0.13/me/ubuntu]
05f3b67ed530: Preparing 
ec1817c93e7c: Preparing 
9e97312b63ff: Preparing 
e1c75a5e0bfa: Preparing 
unauthorized: unauthorized to access repository: me/ubuntu, action: push: unauthorized 
[root@docker01 ~]# docker login 10.0.0.13
Username: admin
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store







```



# harbor

```yml
# Configuration file of Harbor

# The IP address or hostname to access admin UI and registry service.
# DO NOT use localhost or 127.0.0.1, because Harbor needs to be accessed by external clients.
hostname: xx.xx.xx.xx  # 对外开发访问的 ip或域名

# http related config
http:
  # port for http, default is 80. If https enabled, this port will redirect to https port
  port: 80
# 注释掉https,当然你也可以通过设置证书，开放https，建议生产使用https
# https related config
#https:
  # https port for harbor, default is 443
#  port: 443
  # The path of cert and key files for nginx
#  certificate: /your/certificate/path
#  private_key: /your/private/key/path

# # Uncomment following will enable tls communication between all harbor components
# internal_tls:
#   # set enabled to true means internal tls is enabled
#   enabled: true
#   # put your cert and key files on dir
#   dir: /etc/harbor/tls/internal

# Uncomment external_url if you want to enable external proxy
# And when it enabled the hostname will no longer used
# external_url: https://reg.mydomain.com:8433

# The initial password of Harbor admin
# It only works in first time to install harbor
# Remember Change the admin password from UI after launching Harbor.
harbor_admin_password: Harbor12345

# Harbor DB configuration
database:
  # The password for the root user of Harbor DB. Change this before any production use.
  password: root123
  # The maximum number of connections in the idle connection pool. If it <=0, no idle connections are retained.
  max_idle_conns: 50
  # The maximum number of open connections to the database. If it <= 0, then there is no limit on the number of open connections.
  # Note: the default number of connections is 100 for postgres.
  max_open_conns: 100

# The default data volume
data_volume: /data/harbor

# Harbor Storage settings by default is using /data dir on local filesystem
# Uncomment storage_service setting If you want to using external storage
# storage_service:
#   # ca_bundle is the path to the custom root ca certificate, which will be injected into the truststore
#   # of registry's and chart repository's containers.  This is usually needed when the user hosts a internal storage with self signed certificate.
#   ca_bundle:

#   # storage backend, default is filesystem, options include filesystem, azure, gcs, s3, swift and oss
#   # for more info about this configuration please refer https://docs.docker.com/registry/configuration/
#   filesystem:
#     maxthreads: 100
#   # set disable to true when you want to disable registry redirect
#   redirect:
#     disabled: false

# Clair configuration
clair:
  # The interval of clair updaters, the unit is hour, set to 0 to disable the updaters.
  updaters_interval: 12

# Trivy configuration
trivy:
  # ignoreUnfixed The flag to display only fixed vulnerabilities
  ignore_unfixed: false
  # skipUpdate The flag to enable or disable Trivy DB downloads from GitHub
  #
  # You might want to enable this flag in test or CI/CD environments to avoid GitHub rate limiting issues.
  # If the flag is enabled you have to manually download the `trivy.db` file and mount it in the
  # /home/scanner/.cache/trivy/db/trivy.db path.
  skip_update: false
  #
  # insecure The flag to skip verifying registry certificate
  insecure: false
  # github_token The GitHub access token to download Trivy DB
  #
  # Trivy DB contains vulnerability information from NVD, Red Hat, and many other upstream vulnerability databases.
  # It is downloaded by Trivy from the GitHub release page https://github.com/aquasecurity/trivy-db/releases and cached
  # in the local file system (/home/scanner/.cache/trivy/db/trivy.db). In addition, the database contains the update
  # timestamp so Trivy can detect whether it should download a newer version from the Internet or use the cached one.
  # Currently, the database is updated every 12 hours and published as a new release to GitHub.
  #
  # Anonymous downloads from GitHub are subject to the limit of 60 requests per hour. Normally such rate limit is enough
  # for production operations. If, for any reason, it's not enough, you could increase the rate limit to 5000
  # requests per hour by specifying the GitHub access token. For more details on GitHub rate limiting please consult
  # https://developer.github.com/v3/#rate-limiting
  #
  # You can create a GitHub token by following the instuctions in
  # https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line
  #
  # github_token: xxx

jobservice:
  # Maximum number of job workers in job service
  max_job_workers: 10

notification:
  # Maximum retry count for webhook job
  webhook_job_max_retry: 10

chart:
  # Change the value of absolute_url to enabled can enable absolute url in chart
  absolute_url: disabled

# Log configurations
log:
  # options are debug, info, warning, error, fatal
  level: info
  # configs for logs in local storage
  local:
    # Log files are rotated log_rotate_count times before being removed. If count is 0, old versions are removed rather than rotated.
    rotate_count: 50
    # Log files are rotated only if they grow bigger than log_rotate_size bytes. If size is followed by k, the size is assumed to be in kilobytes.
    # If the M is used, the size is in megabytes, and if G is used, the size is in gigabytes. So size 100, size 100k, size 100M and size 100G
    # are all valid.
    rotate_size: 200M
    # The directory on your host that store log
    location: /var/log/harbor

  # Uncomment following lines to enable external syslog endpoint.
  # external_endpoint:
  #   # protocol used to transmit log to external endpoint, options is tcp or udp
  #   protocol: tcp
  #   # The host of external endpoint
  #   host: localhost
  #   # Port of external endpoint
  #   port: 5140

#This attribute is for migrator to detect the version of the .cfg file, DO NOT MODIFY!
_version: 2.0.0

# Uncomment external_database if using external database.
# external_database:
#   harbor:
#     host: harbor_db_host
#     port: harbor_db_port
#     db_name: harbor_db_name
#     username: harbor_db_username
#     password: harbor_db_password
#     ssl_mode: disable
#     max_idle_conns: 2
#     max_open_conns: 0
#   clair:
#     host: clair_db_host
#     port: clair_db_port
#     db_name: clair_db_name
#     username: clair_db_username
#     password: clair_db_password
#     ssl_mode: disable
#   notary_signer:
#     host: notary_signer_db_host
#     port: notary_signer_db_port
#     db_name: notary_signer_db_name
#     username: notary_signer_db_username
#     password: notary_signer_db_password
#     ssl_mode: disable
#   notary_server:
#     host: notary_server_db_host
#     port: notary_server_db_port
#     db_name: notary_server_db_name
#     username: notary_server_db_username
#     password: notary_server_db_password
#     ssl_mode: disable

# Uncomment external_redis if using external Redis server
# external_redis:
#   host: redis
#   port: 6379
#   password:
#   # db_index 0 is for core, it's unchangeable
#   registry_db_index: 1
#   jobservice_db_index: 2
#   chartmuseum_db_index: 3
#   clair_db_index: 4
#   trivy_db_index: 5
#   idle_timeout_seconds: 30

# Uncomment uaa for trusting the certificate of uaa instance that is hosted via self-signed cert.
# uaa:
#   ca_file: /path/to/ca

# Global proxy
# Config http proxy for components, e.g. http://my.proxy.com:3128
# Components doesn't need to connect to each others via http proxy.
# Remove component from `components` array if want disable proxy
# for it. If you want use proxy for replication, MUST enable proxy
# for core and jobservice, and set `http_proxy` and `https_proxy`.
# Add domain to the `no_proxy` field, when you want disable proxy
# for some special registry.
proxy:
  http_proxy:
  https_proxy:
  no_proxy:
  components:
    - core
    - jobservice
    - clair
    - trivy



ssl_cert = /data/cert/server.crt           ##若没有此目录则需要手动建立
self_registration = on     ##是否开启自注册
token_expiration = 30      ##Token有效时间，默认30分钟
project_creation_restriction = everyone     ##用户创建项目权限控制，默认是everyone（所有人），也可以设置为adminonly（只能管理员）



cd /home/admin/harbor/
./install.sh


docker ps

# 需要harbor的配置文件,要在harbor目录上执行
docker-compose ps
```



# docker network 

```bash
Usage:	docker network COMMAND

[root@docker01 ~]# docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
8f7cd27b6a66        bridge              bridge              local 默认模式,相当于NAT
4ca748ac73b1        host                host       local 共用宿主机 network namespace
68c0b31975a2        none                null                local 无网络模式

# 高可用必备,一台机器宕机后,另一台机器使用其网络资源
container 与其他容器共用network namespace

docker run -it --name ubuntu --network container ubuntu


# macvlan
docker 跨主机访问macvlan实现(只能局域网内,无法上外网,可用于堡垒机)
docker network create --driver macvlan --sebnet 100.0.0.0/24 --gateway 10.0.0.254 -o parent=eth0 macvlan_1

# ubuntu(系列需要)
ip link set eth0 promsic on

# docker01
docker run -it --network macvlan_1 --ip 100.0.0.10 ubuntu

# docker 02
docker run -it --network macvlan_1 --ip 100.0.0.11 ubuntu

# 用于打开Ipv4
echo 1 > /proc/sys/net/ipv4/ip_forward


# overlay = bridge + macvlan +统一配置网络
每个容器有两块网卡,eth1实现外网访问,eth0实现内网访问

docker pull progrium/consul
docker run -d -p 8500:8500 -h consul --name consul progrium/consul -server -bootstrap

# 1.编译配置文件(*每台机器都需要)
vim /etc/docker/daemon.json
"hosts": ["tcp://0.0.0.0:2376", "unix:///var/run/docker.sock"],
"cluster-store": "consul://10.0.0.13:8500",
"cluster-advertise": "10.0.0.13:2376"

systemctl daemon-reload
systemctl restart docker

# 有的版本需要更改
vim /usr/lib/systemd/system/docker.socket

2.创建overlay网络驱动
docker network create -d overlay --subnet 172.16.0.0/24 --gateway 172.16.0.254 overlay01

3.测试
docker run -it --network overlay01 --name ubuntu01 ubuntu

docker pull busybox








```

