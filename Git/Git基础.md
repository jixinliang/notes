# 1. git的基础操作

```bash
git status

git confit --list

git config --global color.ui always
git config --global user.name string
git config --global user.email string

git init
git add .
git status
git rm --cached <file>...

# 同时删除文件和临时缓存中的文件
git rm -f a

# 从缓存区提交到本地仓库
git commit -m "add file a and b"

# 这样并没有把临时缓存区的a删除(还很麻烦)
mv a a.txt
git status

# 先删除缓存再添加文件
git rm --cache a
git add a.txt

# 利用git mv一步到位的修改文件
git mv a.txt a

# 文件对比
echo aa > a
git diff a

# 对比工作目录(不跟文件)
git diff

# 对比暂存区与工作目录
git diff --cached

# 这样工作目录和暂存区的数据就都相同了
git add .
git commit -m "add new data to a"

# 缩写形式(添加又提交)
git commit -am "add new data to a"

# 查看历史提交
git log
# 只看描述信息
git log --oneline
# 
git log --oneline -p

# 最近一次提交
git log -1

# 数据回滚
git reset --hard 431cc79

# 查看所有改动log(回滚后的前的log)
git reflog

```

# 2. git的分支

```bash
# 查看分支的指向
git log --oneline --decorate 

# 创建分支
git branch testing

# 查看分支
git branch

# 切换分支
git checkout testing 

# 创建并切换
git checkout -b testing

# 删除分支
git branch -d testing

# 合并分支
git merge testing -m "merge testing branch"

```



## 2.1 代码合并案例

```bash
案例
[root@docker01 data]# touch aa bb cc
[root@docker01 data]# git add aa
[root@docker01 data]# git commit -m "add aa"
[master 7c7ba9e] add aa
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 aa
[root@docker01 data]# git add bb
[root@docker01 data]# git commit -m "add bb"
[master 208d4da] add bb
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 bb
[root@docker01 data]# git add cc
[root@docker01 data]# git commit -m "add cc"
[master 1ef8187] add cc
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 cc
[root@docker01 data]# git log --oneline --decorate -1
1ef8187 (HEAD -> master) add cc
# 新分支
[root@docker01 data]# git checkout -b testing
Switched to a new branch 'testing'
[root@docker01 data]# git branch 
  master
* testing
[root@docker01 data]# git log --oneline --decorate -1
1ef8187 (HEAD -> testing, master) add cc	# 此处指针指向testing和master
[root@docker01 data]# touch test-dd
[root@docker01 data]# git add .
[root@docker01 data]# git commit -m "add new file test-dd"
[testing 3e53e8f] add new file test-dd
[root@docker01 data]# git log --oneline --decorate -1
3e53e8f (HEAD -> testing) add new file test-dd	# 此处指针指向testing
# 两边分支互补影响
[root@docker01 data]# git checkout master 
Switched to branch 'master'
[root@docker01 data]# git branch 
* master
  testing
[root@docker01 data]# ll
# 合并分支(一定要在主分支上合并)
[root@docker01 data]# git merge testing -m "merge testing branch"
Merge made by the 'recursive' strategy.
 test-dd | 0
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 test-dd
# 合并完数据后删除小分支就好(重新创建分支会更快些)
[root@docker01 data]# git branch -d testing 
Deleted branch testing (was 3e53e8f).

```

## 2.2 冲突合并案例

```bash
# 添加了testing分支,又返回主分支
[root@docker01 data]# git checkout -b testing
Switched to a new branch 'testing'
[root@docker01 data]# git checkout master 
Switched to branch 'master'
# 主分支添加数据
[root@docker01 data]# echo master >> aa
[root@docker01 data]# git commit -am "add data to file aa"
[master d012719] add data to file aa
 1 file changed, 1 insertion(+)
 # 切换到testing支添并加数据
[root@docker01 data]# git checkout testing 
Switched to branch 'testing'
[root@docker01 data]# cat aa	# testing 分支里的aa是没有数据的
[root@docker01 data]# echo testing >> aa
[root@docker01 data]# git commit -am "testing add data to aa"
[testing ebddb9d] testing add data to aa
 1 file changed, 1 insertion(+)
[root@docker01 data]# cat aa
testing
# 返回主分支
[root@docker01 data]# git checkout master 
Switched to branch 'master'
# 合并分支
[root@docker01 data]# git merge testing 
Auto-merging aa
CONFLICT (content): Merge conflict in aa
Automatic merge failed; fix conflicts and then commit the result.
[root@docker01 data]# cat aa 
<<<<<<< HEAD
master
=======
testing
>>>>>>> testing
# 冲突文件只能手动修改后再提交(使用: vim)
[root@docker01 data]# git commit -am "merged testing branch"
[master f88d003] merged testing
[root@docker01 data]# git status
On branch master
nothing to commit, working tree clean

```

# 3. git的标签

```bash
git tag
    -a, --annotate        annotated tag, needs a message
    -m, --message <message>

git tag -a v1.0 -m "add new tag"
git tag -a v1.1 7983d83 -m "tag v1.1 add data"

# 查看标签
git tag

# 查看该标签上的操作
git show v1.0

# 通过标签便于回滚数据
git reset --hard v1.0

git tag -d v1.0
```

# 4. GitHub的使用

```bash
添加remote
git remote add origin git@github.com:....

# 查看
git remote

git push -u origin master
	-u, --set-upstream    set upstream for git pull/status


git clone git@github.com:....
 

```

# gitlab

```bash

cat >/etc/yum.repos.d/gitlab-ce.repo<<EOF
[gitlab-ce]
name=Gitlab CE Repository
baseurl=https://mirrors.tuna.tsinghua.edu.cn/gitlab-ce/yum/el$releasever/
gpgcheck=0
enabled=1
EOF

sudo dnf install -y curl policycoreutils openssh-server


Thank you for installing GitLab!
GitLab was unable to detect a valid hostname for your instance.
Please configure a URL for your GitLab instance by setting `external_url`
configuration in /etc/gitlab/gitlab.rb file.
Then, you can start your GitLab instance by running the following command:
  sudo gitlab-ctl reconfigure

For a comprehensive list of configuration options please see the Omnibus GitLab readme
https://gitlab.com/gitlab-org/omnibus-gitlab/blob/master/README.md


# 安装目录
/opt/gitlab/
# 存放数据的目录
/var/opt/gitlab
# 存放仓库数据
/var/opt/gitlab/git-data

gitlab-ctl reconfigure
gitlab-ctl status
gitlab-ctl stop
gitlab-ctl stop nginx
gitlab-ctl tail

git remote rename origin old-origin
git remote add origin ...
git push -u origin master
git push -u origin --all
git push -u origin --tags


# 切换并上传dev组
git checkout -b dev
git add .
git commit -m "dev add new file"
git push -u origin dev

# 拉取最新数据
git pull

# 写新项目时一定要先删除支分支,再创建新分支后,再写,更新最新数据
git branck -d dev
git pull
git checkout -b dev

```

# jenkins的使用

```bash
docker run ^
  -u root ^
  --rm ^
  -d ^
  -p 8080:8080 ^
  -p 50000:50000 ^
  -v jenkins-data:/var/jenkins_home ^
  -v /var/run/docker.sock:/var/run/docker.sock ^
  jenkinsci/blueocean


# Jenkins需要JDK作为runtime
dnf install java-1.8.0-openjdk.x86_64
dnf localinstall jenkins-2.235.1-1.1.noarch.rpm

# 就该Jenkins权限
vim /etc/sysconfig/jenkins
JENKINS_USER="root" 
chown -R root. /var/lib/jenkins
chown -R root. /var/cache/jenkins
chown -R root. /var/log/jenkins


# 将plugins.tar.gz移动到该目录下解压
/var/lib/jenkins


/var/lib/jenkins/users


#!/bin/bash
CODE_DIR="/srv/"  # /var/lib/jenkins/workspace/freestyle-job
WEB_DIR="/srv/vol"
IP=10.0.0.8
TIME=`date +%F-%H-%M-%S`
DATA_DIR="web-${TIME}"
TAR_DATA="${DATA_DIR}.tar.gz"

cd $CODE_DIR && tar czf /tmp/${TAR_DATA} ./
scp /tmp/${TAR_DATA} $IP:${WEB_DIR}
ssh root@$IP "cd ${WEB_DIR} && mkdir ${DATA_DIR}"
ssh root@$IP "cd ${WEB_DIR} && tar -C ${DATA_DIR}/ -xf ${TAR_DATA} && rm -f ${TAR_DATA}"
ssh root@$IP "cd ${WEB_DIR} && rm -rf html && ln -s ${DATA_DIR} html"


ssh-copy-id -i .ssh/ root@10.0.0.8

# jenkins项目设置
1. 丢弃的构建
    保持天数5-7天
    构建的最大数5-7个

2. 源码管理
git
repository url: git@10.0.0.13:test/dzp.git
把本机的秘钥对配置到Credentials中
cat .ssh/id_rsa

3. # 设置gitlab触发器
Build when a change is pushed to GitLab. GitLab webhook URL: http://10.0.0.8:8080/project/freestyle-job

产生token,拿着该token到gitlab下的相关项目下配置web hook

4. 构建后操作
publish build status to gitlab

# gitlab配置
1. 获取gitlab设置中的access token,选中api,拿到产生的token到Jenkins设置中进行配置
以达到双向调用

# 脚本测试(在Jenkins服务器本地用docker开Nginx)
#!/bin/bash
CODE_DIR="/var/lib/jenkins/workspace/freestyle-job"
WEB_DIR="/srv/vol"
TIME=`date +%F-%H-%M-%S`
DATA_DIR="web-${TIME}"
TAR_DATA="${DATA_DIR}.tar.gz"

cd $CODE_DIR && tar czf ${WEB_DIR}/${TAR_DATA} ./
cd ${WEB_DIR} && mkdir ${DATA_DIR}
cd ${WEB_DIR} && tar -C ${DATA_DIR}/ -xf ${TAR_DATA} && rm -f ${TAR_DATA}
cd ${WEB_DIR} && rm -rf html && ln -s ${DATA_DIR} html
docker container restart nginx

docker run -d --name nginx -p 80:80 -v /srv/vol/html:/usr/share/nginx/html/ nginx

626 18:00



```

