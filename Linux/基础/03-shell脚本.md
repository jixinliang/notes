# shell脚本

### 01.  批量创建文件, 修改文件名

```bash
#############################
## 批量创建文件,修改文件名 ##
#############################
dir= ~/folder
[ ! -d $dir ] && mkdir -p $dir
cd $dir
for ((i=1; i<=10; i++)); do 
    touch `echo jack-"$i"`
done

for i in `ls`; do 
    mv "$i" `echo ${i/jack/linux}`
done

rename jack linux *

########################################
批量创建文件
dir= ~/folder
[ ! -d $dir ] && mkdir -p $dir
cd $dir
for i in `seq 4`; do 
    touch `echo stu_999_"$i"_finished.jpg`
done

批量修改文件名
#!/bin/bash
for i in `ls *.jpg`
do
    #echo "$i"
    #echo ${i%finished*}.jpg
    mv "$i" `echo ${i%finished*}.jpg`
done

有三种方法才有得选
ls|awk -F '_' '{print "mv "$0" "$1".jay.HTML"}'

rename from to *.html

批量替换扩展名
#!/bin/bash
for i in `ls *.jpg`
do 
    #echo "$i" #echo ${i%finished*}.jpg
    mv "$i" `echo ${i/jpg/JPG}`
done

#########################################
[jack@me ~]$ cat createfile.sh 
#!/bin/bash

dir=~/folder
[ ! -d $dir ] && mkdir -p $dir
cd $dir

for ((i=0; i<10; i++)); do
	prefix=`echo $RANDOM|md5sum|tr "[0-9]" "[j-z]"|cut -c 2-11`
	touch ${prefix}_jack.html
done
```

### 02. 批量创建用户及密码

```bash
########################
## 批量创建用户及密码 ##
########################
for i in `seq -w 10`; do
    useradd jack-"$i" && \
    echo "$i" |passwd --stdin jack-"$i"
done

for i in `seq -w 10`; do
    userdel -r jack-"$i"
done

echo jack-{01..10}|tr "" "\n"|sed -r 's#(.*)#useradd \1;pass=$((RANDOM+10000000));echo "$pass"|passwd --stdin \1;echo -e "\1 \t `echo "$pass"`" >>/tmp/pass.log#g' |bash
```

### 03. 批量创建用户及密码Pro版

```bash
#############################
## 批量创建用户及密码Pro版 ##
#############################
[jack@me ~]$ cat muti_user_add.sh 
#!/bin/bash
[ -f /etc/init.d/functions ] && . /etc/init.d/functions
> /home/jack/user_pass.txt

for i in `seq -w 10`; do
    flag=`grep "\b$i\b" /etc/passwd|wc -l`  ## good idea
    if [ $flag -eq 1 ]; then
        action "user: $i already exists" /bin/false
        continue
    fi
    passwd=`echo $RANDOM |md5sum |cut -c 2-9`
    useradd jack-"$i" &>/dev/null && add_status=$?
    echo "$passwd" |passwd --stdin jack-"$i" &>/dev/null && pass_status=$?
    if [ $add_status -eq 0 -a $pass_status -eq 0 ]; then
        action "useradd jack-$i" /bin/true
        echo "user: jack-$i  pass: $passwd" >> /home/jack/user_pass.txt
    else
        action "useradd jack-$i" /bin/false
        echo "user: jack-$i  pass: $passwd" >> /home/jack/faild_user.txt
    fi
done

passwd=`echo $RANDOM |md5sum |cut -c 2-9`
useradd jack-"$i" && \
echo "jack-$i:$passwd"|tee >>user_pass.txt
放循环之外
chpasswd < user_pass.txt

批量删除用户
[jack@me ~]$ cat muti_user_del.sh 
#!/bin/bash
. /etc/init.d/functions

for i in `seq -w 10`; do
    userdel -r jack-"$i" &>/dev/null && del_status=$?
    if [ $del_status -eq 0 ]; then
        action "user jack-$i deleted" /bin/true
    else
        action "user jack-$I del faild" /bin/false
    fi
done
```

### 04.扫描局域网内可用 `ip`

```bash
########################
## 扫描局域网内可用ip ##
########################
[jack@me ~]$ cat ping_Ip.sh 
#!/bin/bash
[ -f /etc/init.d/functions ] && . /etc/init.d/functions

CMD="ping -W 2 -c 2"
IP="10.0.0."

for i in `seq 20`; do
	{
		$CMD $IP$i &>/dev/null
		if [ $? -eq 0 ]; then
			action $IP$i /bin/true
		fi
	} &
done

##############################################

[jack@me ~]$ cat nmap_ip.sh 
#!/bin/bash

IP="10.0.0.0/24"
CMD="nmap -sn" 

map() {
	$CMD $IP|grep "Nmap scan report"|awk -F '[ ]+' '{print $NF}'
}
map

[jack@me ~]$ nmap -sn 10.0.0.0/24|awk '/Nmap scan report/ {print $NF}'
10.0.0.2
10.0.0.13
```

### 05. 破解MD5值

```bash
###############
## 破解MD5值 ##
###############
91: 80ad60f9
1346: 00205d1c
7041: 1f6d12dd
25345: a3da1677
25667: 21029299

[jack@me ~]$ cat md5Cut.sh
#!/bin/bash
li=(80ad60f9 21029299 00205d1c a3da1677 1f6d12dd)
for i in {0..32767}; do
	#echo "$i `echo $i|md5sum|cut -c 1-8`"
	res=`echo $i|md5sum|cut -c 1-8`
	for ((j=0; j<${#li[*]}; j++)); do
		if [ "${li[j]}" = "$res" ]; then
			echo "$i: $res"
		fi
	done
done
```

### 06. 打印长度小于4的单词

```bash
#########################
## 打印长度小于4的单词 ##
#########################
[jack@me ~]$ cat less4word.sh 
#!/bin/bash
## string
str="This is jack speaking long time no see you"
for i in $str; do
	#echo $i
	if [ ${#i} -lt 4 ]; then
		echo $i
	fi
done

## array
li=(This is jack speaking long time no see you)
for i in ${li[@]}; do
    #echo $i
    if [ ${#i} -lt 4 ]; then
        echo $i
    fi
done


li=(This is jack speaking long time no see you)
for ((i=0; i<${#li[*]}; i++)); do
    #echo ${li[$i]}
    if [ ${#li[$i]} -lt 4 ]; then
        echo ${li[$i]}
    fi
done

## awk
str="This is jack speaking long time no see you"
echo ${str} |tr " " "\n"|awk '{if(length($1)<4) print $1}'
echo ${str} |tr " " "\n"|awk 'length < 4 {print $1}'

## wc -L
```

### 07. `wget` 检查网站

```bash
##################
## wget检查网站 ##
##################
[jack@me ~]$ cat wget_url.sh 
#!/bin/bash
[ -f /etc/init.d/functions ] && . /etc/init.d/functions

retVal=0
usage() {
	echo "USAGE: sh $0 www.example.com"
	exit 1
}

checkUrl() {
	wget -T 6 --spider -t 2 $1 &>/dev/null
	retVal=$?
	if [ $retVal -eq 0 ]; then
		action "Url $1" /bin/true
	else
		action "Url $1" /bin/false
	fi
	return $retVal
}

main() {
	if [ $# -ne 1 ]; then
		usage
	fi
	checkUrl $1
	retVal=$?
	return $retVal
}
main $*
```

### 08. 解决网站DOS攻击

```bash
#####################
## 解决网站DOS攻击 ##
#####################
[jack@me ~]$ cat web_dos.sh
#!/bin/bash

usage() {
	if [ $# -ne 1 ]; then
		echo "USAGE: $0 ARG"
		exit 1
	fi
}

iptalbe() {
	`awk '{print $1}' $1|sort|uniq -c|sort -rn -k1` >/tmp/multi_conn.log
	exec </tmp/multi_conn.log
	while read line; do
		ip=`echo $line|awk '{print $2}'`
		n=`echo $line|awk '{print $1}'`
		l=`iptables -L -n|grep "$ip"|wc -l`
		if [ $n -ge 10 -a $l -lt 1 ]; then
			iptalbes -I INPUT -s $ip -j DROP
			echo $ip >>/tmp/ipPool_$(date +%F).log
		fi
	done	
}

del() {
	exec </tmp/ipPool_$(date +%F -d "-1 day").log
	while read line; do
		l=`iptables -L -n|grep "$ip"|wc -l`
		if [ $l -ge 1 ]; then
			iptalbes -D INPUT -s $line -j DROP
		fi
	done
}

main() {
	usage	
	iptalbe $1
	sleep 3
	del
}
main $*
```

### 09. `mysql` 数据库-分库备份

```bash
#############################
## 对mysql数据库进行分库备份 ##
#############################
[jack@me ~]$ cat backmysql.sh 
#!/bin/bash
#[ -f /etc/init.d/functions ] . /etc/init.d/functions
backPath=/srv/backup
user=root
pass=666
socketPath=/var/lib/mysql/mysql.sock
#socket1=
myCmd="mysql -u$user -p$pass -S$socketPath"
myDump="mysqldump -u$user -p$pass -S$socketPath -x -B -F -R"
dbList=`$myCmd -e "show databases;"|sed 1d|egrep -v "_schema"`
[ ! -d $backPath ] && mkdir -p $backPath

for dbname in $dbList; do
	$myDump $dbname|gzip >$backPath/${dbname}_$(date +%F).sql.gz
done

```

### 10. `mysql` 数据库-分表备份

```bash
#############################
## 对mysql数据库进行分表备份 ##
#############################
[jack@me ~]$ cat backmysqltable.sh 
#!/bin/bash
#[ -f /etc/init.d/functions ] . /etc/init.d/functions

backPath=/srv/backup
user=root
pass=666
socketPath=/var/lib/mysql/mysql.sock
#socket1=
myCmd="mysql -u$user -p$pass -S$socketPath"
myDump="mysqldump -u$user -p$pass -S$socketPath -x -F -R"
dbList=`$myCmd -e "show databases;"|sed 1d|egrep -v "_schema"`

for dbname in $dbList; do
    tbList=`$myCmd -e "show tables from $dbname;"|sed 1d`
	for tbname in $tbList; do
		[ ! -d $backPath/${dbname} ] && mkdir -p $backPath/${dbname}
		$myDump $dbname $tbname|gzip >$backPath/${dbname}/${dbname}_${tbname}_$(date +%F).sql.gz
	done
done
```

### 11. `MySQL` 多实例启动脚本

```bash
#########################
## MySQL多实例启动脚本 ##
#########################
[jack@me ~]$ cat multi-case.sh 
#!/bin/bash
[ -f /etc/init.d/functions ] && . /etc/init.d/functions

user=root
pass=666
socket=/var/lib/mysql/mysql.sock
myCmd="mysql -u$user -p$pass -S$socket"
myAdmin="mysqladmin -u$user -p$pass -S$socket"

usage() {
	echo "USAGE: $0 {start|stop|restart}"
	exit 1
}

stat() {
retVal=$?
if [ $retVal -eq 0 ]; then
	action "mysql $1" /bin/true
else
	action "mysql $1" /bin/false
fi
return $retVal
}

start() {
	#mysqld_safe --defaults-file=/etc/my.cnf &>/dev/null &
	systemctl start mariadb.service
	stat start
}

stop() {
	$myAdmin shutdown &>/dev/null		
	stat stop
}

case "$1" in
  start)
	  start
	  ;;
  stop)
	  stop
	  ;;
  restart)
	  stop
	  sleep 2
	  start
	  ;;
  *)
	usage
esac
```

### 12. 监控`MySQL`主从同步是否异常

```bash
###############################
## 监控MySQL主从同步是否异常 ##
###############################
[jack@me ~]$ cat moni-mysql.sh 
#!/bin/bash
backPath=/srv/backup
user=root
pass=666
socket=/var/lib/mysql/mysql.sock
myCmd="mysql -u$user -p$pass -S$socket"
key="_Running|Seconds_Behind|Last_SQL_Errno"
stat=($($myCmd -e "show slave status\G;"|egrep $key|awk '{print $NF}'))
errno=(1158 1159 1008 1007 1062)

checkStatus() {
if [ "${stat[0]}" = "yes" -a "${stat[1]}" = "yes" -a "${stat[2]}" = "0" ]; then
	echo "mysql slave is OK."
	retVal=0
	return $retVal
else
	retVal=1
	return $retVal
fi
}

checkErrno() {
	checkStatus
	if [ $? -eq 1 ]; then
		for ((i=0; i<${#errno[*]}; i++)); do
			if [ "${errno[i]}" = "${li[3]}" ]; then
				$myCmd -e "stop slave;"
				$myCmd -e "set global sql_slave_skip_counter = 1;"
				$myCmd -e "start slave;"
			fi
		done
	fi
}

reCheckStatus() {
	stat=($($myCmd -e "show slave status\G;"|egrep $key|awk '{print $NF}'))
	checkStatus &>/dev/null
	if [ $? -eq 1 ]; then
		echo "mysql slave was faild `date +%F-%T`" >/tmp/myslave.log
		mail -s "mysql slave was faild `date +%F-%T`" </tmp/myslave.log
	fi
}

main() {
	while true; do
		checkErrno
		reCheckStatus
		sleep 30
	done
}
main
```

### 13. 网站文件指纹

```bash
##################
## 网站文件指纹 ##
##################
creat_md5_fingerprint.sh

Path=/var/www/html/bbs/
find $path -type f |xargs md5sum > /server/scripts/007/md5sum.db
find $path -type -f > /server/scripts/007/check_site.log

md5sum.sh

Path=/server/scripts/007
F=`cat $path/check_site.log`

[ ! -f "$F" ] && touch $Path/check_site.log
while ture
do
    n=`md5sum -c $Path/md5sum.db 2>/dev/null |grep FAILED |wc -l`
    find /var/www/html/bbs -type f  > $Path/new_site.log
    log=/tmp/check.log
    [ ! -f $log ] && touch $log
    if [ $n -ne 0 ] || [ `cat $Path/new_site.log |wc -l` -ne `cat $F|wc -l` ]; then
        echo "`md5sum -c md5sum.db 2 >/dev/null|grep FAILED`" > $log
        diff $Path/check_site.log $Path/new_site.log >> $log 2>&1
        mail -s "Attention, site was changed at $(date)" souvenir9527@qq.com <$log
    fi
sleep 3
done
```

### 14. 乘法口诀

```bash
##############
## 乘法口诀 ##
##############
[jack@me ~]$ cat muti_for.sh 
#!/bin/bash

for i in `seq 9`; do
    for j in `seq 9`; do
        if [ $i -ge $j ]; then
            echo -en "  $j x $i = $(expr $i \* $j)"
        fi
    done			
echo " "
done


for i in `seq 9`; do
    for j in `seq 9`; do
        [ $i -ge $j ] && echo -en "  $j x $i = $(expr $i \* $j)"
    done
echo " "
done
```

### 15. web监控程序

```bash
#################
## web监控程序 ##
#################

serverMonitor () {
timeout=10
fails=0
success=0
while true
do
    /usr/bin/wget -q -O /dev/null --timeout=$timeout http://10.0.0.8 
    if [ $? -ne 0 ]; then
        let fails=fails+1
        success=0
    else
        fails=0
        let success=success+1
    fi
    
    if [ $success -ge 1 ]; then
        exit 0
    fi
    
    if [ $fails -ge 2 ]; then
        Critical="there is something wrong with WEB server, ATTENTION !"
        echo $Critical |mail -s "enmergency" souvenir9527@qq.com
        exit
    fi
done
}
```

### 16. 一键增加删除虚拟 `ip`

```bash
########################
## 一键增加删除虚拟ip ##
########################
[jack@me ~]$ cat add_net1.sh 
#!/bin/bash
## for add virtual ip
start() {
    for ((i=0; i<=15; i++)); do
        if [ "$i" -eq 10 ]; then
            continue
        else
            ifconfig ens32:"$i" 10.0.1."$i" netmask 255.255.240.0 $1
        fi
    done
}

stop() {
    for ((i=0; i<=15; i++)); do
        if [ "$i" -eq 10 ]; then
            continue
        else
            ifconfig ens32:"$i" 10.0.1."$i" netmask 255.255.240.0 $1
        fi
    done
}

usage() {
    cat <<-EOF
        USAGE: sh $0 {up|down}
    EOF
}

case "$1" in
  up)
    start	
    ;;
  down)
    stop
    ;;
  *)
    usage
    exit 1
esac
exit 0
```

### 17. 用于web服务重启后快速检查

```bash
###############################
## 用于web服务重启后快速检查 ##
###############################
[jack@me ~]$ cat get_url_status.sh
#!/bin/bash
. /etc/init.d/functions
retVal=0
failCount=0
mailGroup="souvenir9527@qq.com souvenir9527@163.com"
logFile=/tmp/web_check.log
scriptsPath=/srv/scripts
webUrl=$1
getUrlStatus() {
	for ((i=0; i<3; i++)); do
		wget -T 6 --tries=1 --spider http://$1 &>/dev/null
		[ $? -ne 0 ] && let failCount+=1;
	done

	if [ $failCount -gt 1 ]; then
		retVal=1
		nowTime=`date +%F-%T`
		subject="http://$url service is error. ${nowTime}"
		echo "send to $mailUser, title: $subject" >$logFile
		for mailUser in $mailGroup; do
			#mail -s "$subject" $mailUser <$logFile
			continue
		done
	else
		retVal=0
	fi
	return $retVal
}

[ ! -d "$scriptsPath" ] && mkdir -p $scriptsPath
[ ! -f "$scriptsPath/domainList.txt" ] && {
cat > $scriptsPath/domainList.txt <<EOF
www.hehe.com
blog.hehe.com
bbs.hehe.com
www.baidu.com
www.qq.com
EOF
}
for url in `cat $scriptsPath/domainList.txt`; do
	#echo -n "checking for $url: "
	#getUrlStatus $url && action "successful" /bin/true|| action "faild" /bin/false
    getUrlStatus $url  ## for定时任务,不需要输出,也不需要发邮件
done
```

### 18. 跳板机框

```bash
##############
## 跳板机框 ##
##############
#!/bin/bash

trapper() {
    trap '' HUP INT QUIT TERM TSTP
}

usage() {
    cat <<- EOF
    ========== machine list =========
        1) web01
        2) web02
        3) exit
    =================================
    EOF
    echo "Pls select a num:"
}

menu() {
    read -p ">>> " num
    case "$num" in
      1)
        echo "web01"
        ;;
      2)
        echo "web02"
        ;;
      3)
        sleep 5
        echo "hello there"
        ;;
      exit)
        exit
    esac
}

main() {
    while :; do
        clear
        usage
        trapper
        menu
        sleep 2
    done
}
main   


root to do 
cat >> /etc/profile.d/springboard.sh <<EOF
[ $UID -ne 0 -o $USER != "jack" ] && /bin/sh /srv/springboard.sh
EOF
```

### 19. 颜色

```bash
##########
## 颜色 ##
##########
[jack@me ~]$ cat case_color.sh 
#!/bin/bash
red='\E[1;31m'
green='\E[1;32m'
yellow='\E[1;33m'
blue='\E[1;34m'
pink='\E[1;35m'
res='\E[0m'
flash='\E[31;5m'

usage() {
    echo -e "${flash}Pls select a num follow the menu !!${res}"
    #exit 1
}

menu() {
    cat <<- EOF
    1.apple
    2.leaf
    3.pinapple
    4.sky
    5.cherry
    6.exit
EOF
echo "Pls select a num:" 
read -p ">>> " num
}

sele() {
    case "$num" in
      1)
          echo -e "${red}apple${res}"
          ;;
      2)
          echo -e "${green}leaf${res}"
          ;;
      3)
          echo -e "${yellow}pinapple${res}"
          ;;
      4)
          echo -e "${blue}sky${res}"
          ;;
      5)
          echo -e "${pink}cherry${res}"
          ;;
      6)
          exit
          ;;
      *)
          usage
    esac
}

main() {
    while true; do
        menu
        sele
        sleep 1 
    done
}
main

```

