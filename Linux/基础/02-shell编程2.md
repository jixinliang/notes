# 2. shell编程基础二

```bash
如何进行整数的计算?
    (()) let expr expr$[] bc typeset $[] awk

  AWK
    [jack@me ~]$ echo "1.8 1.6"|awk '{print($2-$1)}'
    -0.2
    [jack@me ~]$ echo "1.4 1.6"|awk '{print($2-$1)}'
    0.2

    [jack@me ~]$ echo `seq -s "+" 10` = `seq -s "+" 10|bc`
    1+2+3+4+5+6+7+8+9+10 = 55
    [jack@me ~]$ echo `seq -s "+" 10` = $(expr $[`seq -s "+" 10`])
    1+2+3+4+5+6+7+8+9+10 = 55
    [jack@me ~]$ echo `seq -s "+" 10` = $[`seq -s "+" 10`]
    1+2+3+4+5+6+7+8+9+10 = 55
    [jack@me ~]$ echo `seq -s "+" 10` = $((`seq -s "+" 10`))
    1+2+3+4+5+6+7+8+9+10 = 55

shell变量的输入
    Shell变量除了可以直接赋值或脚本传参外,还可以使用read命令从标准输入获得, 
    read为bash内置命令,可通过help read查看帮助。



伪代码
    - 需求分析
    - 需求设计

简单地说,函数的作用就是把程序里多次调用相同的代码部分定义成一份,
然后为这一份代码起个名字,其它所有的重复调用这部分代码就都只调用这个名字就可以了。
当需要修改这部分重复代码时,只需要改变函数体内的一份代码即可实现所有调用修改。

使用函数的优势: 
    1、把相同的程序段定义成函数,可以减少整个程序的代码量。
    2、可以让程序代码结构更清晰。
    3、增加程序的可读、易读性,以及可管理性。
    4、可以实现程序功能模块化,不同的程序使用函数模块化。
    强调:对于shell来说, linux系统的2000个命令都可以说是shell的函数。

funcName() {
    order
    return n
}

function funcName() {
    order
    return n
}
提示: shell的返回值是exit输出返回值,函数里用return输出返回值。

shell函数的执行
调用函数:
1)直接执行函数名即可(不带括号)。
注意:
a.执行函数时,函数后的小括号不要带了。
b.函数定义及函数体必须在要执行的函数名的前面定义,
shell的执行从上到下按行执行的。
2)带参数的函数执行方法:
函数名 参数1 参数2.
提示:函数的传参和脚本的传参类似,只是脚本名换成函数名即可。

shel1的位置参数($1 $2 $3 ... $# $* $?以及$@)都可以是函数的参数。
此时父脚本的参数临时地被函数参数所掩盖或隐藏。
$0比较特殊,它仍然是父脚本的名称。
当函数完成时,原来的命令行脚本的参数即恢复。
在shel1函数里面, return命令功能与shel1里的exit类似,作用是跳出函数。
在shel1函数体里使用exit会退出整个shel1脚本,而不是退出shel1函数。
return语句会返回一个退出值(返回值)给调用函数的程序。
函数的参数变量是在函数体里面定义,如果是普通变量一般会使用local i定义。





颜色  ## console_codes - Linux console escape and control sequences


restart 查看进程号有没有发生变化
reload 查看配置参数有没有生效

case语句小结
    1.case语句就相当于多分支的if语句,case语句优势更规范易读.
    2.语句适合变量的值少,且为固定的数字或字符串集合.
    3.系统服务启动脚本传参的判断多用case语句.参考rpcbind,nfs,crond脚本
    4.所有的case语句都可以用if实现,但是case更规范清晰些,
    5.case一般适合于服务的启动脚本.
    6.case的变量的值如果已固定的start,stop,restart元素的时候比较合适一些.

运维层面
    case只要是写启动脚本,范围更更窄
    if取值判断,比较,应用更广.


读懂functions


防止客户端脚本执行中断的方法:
1) sh while_01.sh &
2) nohup /server/scripts/uptime.sh &
3) screen 保持会话


[jack@me ~]$ while true;do uptime &>/dev/null;sleep 2;done &
[1] 1923
[jack@me ~]$ jobs
[1]+  Running                 while true; do
    uptime &> /dev/null; sleep 2;
done &
[jack@me ~]$ kill %1    ##指定杀死进程号
[jack@me ~]$ jobs
[1]+  Terminated              while true; do
    uptime &> /dev/null; sleep 2;
done


bg:后台运行
fg:挂起程序
jobs:显示后台程序
kill,killl,pkill 掉进程
crontab:设置定时
ps:查看进程
pstree:显示进程状态树
top:显示进程
nice:改变优先权
nohup:用户退出系统之后继续工作
pgrep:查找匹配条件的进程
strace:跟踪一个进程的系统调用情况
ltrace:跟进程调用库函数的情况
vmstat:报告虚拟内存统计信息

[jack@me ~]$ time sh -x sum_n.sh 100000
+ i=100000
+ sum=0
+ (( sum=i*(i+1)/2 ))
+ echo 5000050000
5000050000

real	0m0.004s
user	0m0.000s
sys	0m0.004s
[jack@me ~]$ 

[jack@me ~]$ time sh sum100.sh 
5050

real	0m0.005s
user	0m0.001s
sys	0m0.003s


while循环小结:
    1, while循环的特长是执行守护进程以及我们希望循环不退出持续执行的情况,用于频率小于1分钟循环处理(crond) ,其他的while循环几乎都可以被我们即将要讲 for循环替代。
    2, case语句可以if语句替换,一般在系统启动脚本传入少量固定规则字符串,用case语句,其他普通判断多用if.
    3、一句话, if, for语句最常用,其次while (守护进程) , case (服务启动脚本)。
    各个语句使用场景:
    条件表达式,简短的判断(文件是否存在,字符串是否为空等)
    if取值判断,不同值数量较少的情况。
    for 正常的循环处理,最常用!
    while守护进程、无限循环(sleep) case服务启动脚本、菜单。
    函数逻辑清晰,减少重复语句。


1)程序持续运行多用while,包括守护进程,还有配合read读入循环处理.
2)有限次循环多用for,工作中for使用更多。


break n   |n表示跳出循环的层数,如果省略n表示跳出整个循环。
continue n|n表示退到第n层继续循环,如果省略n表示跳过本次循环,忽略本次循环的剩余代码,
           进入循环的下一次循环。
exit n    |退出当前shell程序,n为返回值。n也可以省略,再下一个shell里通过$? 
           接收这个n的值。
return n  |用于在函数里,作为函数的返回值,用于判断函数执行是否正确。



[jack@me ~]$ li=(`ls`)
[jack@me ~]$ echo ${li[*]}
functions scripts shellScripts
[jack@me ~]$ echo ${li[@]}
functions scripts shellScripts
[jack@me ~]$ echo ${#li[@]}
3
[jack@me ~]$ echo ${!li[@]}
0 1 2


1、各类监控脚本,文件、内存、磁盘、端口, URL监控报警。
2、如何监控网站目录文件是否被篡改,以及站点目录批量被篡改后如何恢复。
3、如何开发各类服务rsync, nginx, mysql等的启动及停止专业脚本(使用chkconfig管理)
4、如何开发MySQL主从复制监控报警以及自动处理不复制的脚本。
5、一键配置mysa1多实例、一键配置mysql主从, N多一键部署脚本.
6、监控http/mysql/rsync/nfs/memcached等服务是否异常的生产脚本。
7、一键软件安装及优化, lanmp, linux一键优化,一键数据库安装,优化,配置主从。
8、MySQL多实例启动脚本,分库、分表自动备份脚本。
9、根据网络连接数以及根据web日志PV封IP的脚本。
10、监控网站的pv以及流量,并且对流量信息进行统计。
11、检查web服务器多个URL地址是否异常,要可以批量及通用。
12、系统的基础优化一键优化的脚本。
13、清理系统垃圾及文件(过期备份, clientmquene目录等)脚本。
14, tcp连接状态统计及报警。
15、批量创建用户并设置随机8位密码。
16、批量获取服务器信息、批量分发文件。

(生产实战案例)监控ysQL主从同步是否异常,如果异常,则发送短信或者邮件给管理员。
提示:如果没主从同步环境,可以用下面文本放到文件里读取来模拟：
阶段1:开发一个守护进程脚本每30秒实现检测一次。
阶段2:如果同步出现如下错误号(118, 119, 1008, 1007,106),则跳过错误。
阶段3,请使用数组技术实现上述脚本(获取主从判断及错误号部分)


autocmd BufNewFile *.sh Or ~/.vim/jack/jack.sh


UNIX shell by example

写一些好的博文



[jack@me ~]$ netstat -an|grep EST
tcp        0     36 10.0.0.13:22            10.0.0.1:11313          ESTABLISHED


[jack@me ~]$ help :
:: :
    Null command.
    
    No effect; the command does nothing.
    
    Exit Status:
    Always succeeds.



idea

创建log文件当Ctrl c是自动删除
#!/bin/bash

trap "find /tmp -type f -name "*.log" -delete && exit" INT
while :;do
	touch /tmp/a_$(date +%N).log
	sleep 0.5
done



1,首先跳板机禁止外网IP登录,只能内网IP登录。
ListenAddress 10.0.0.13
2,其他服务器有外网IP的也别忘了禁止外网IP登录,只能内网IP登录。同时禁止root登录,等做完ssh key认证,连密码登录也禁了,只能通过证书登录,而且只有跳板机有其他服务器的密钥。

修改/etc/ssh/sshd_config
PasswordAuthentication yes
改为
PasswordAuthentication no

3,先远程拨号登录VPN,然后登录跳板机,然后再从跳板机登录其他服务器。

md5sum|cut -c 1-8
21029299
00205d1c
a3da1677
1f6d12dd
890684b



[jack@me ~]$ sh md5Cut.sh
1346: 00205d1c
7041: 1f6d12dd
25345: a3da1677
25667: 21029299


计算字符串长度的方式
[jack@me ~]$ a="hello world"
[jack@me ~]$ echo ${#a}
11
[jack@me ~]$ echo ${a} |wc -L
11
[jack@me ~]$ expr length "${a}"
11


如何实现对mysql数据库进行分库备份,用脚本实现.
1.编程思想
    mysqldump db1 > db1.sql
    mysqldump db2 > db2.sql
2.拿到库名列表
3.循环db,dump

[jack@me ~]$ cat backmysql.sh 
#!/bin/bash
#[ -f /etc/init.d/functions ] . /etc/init.d/functions
backPath=/srv/backup
user=root
pass=666
socket=/var/lib/mysql/mysql.sock
#socket1=
myCmd="mysql -u$user -p$pass -S$socket"
myDump="mysqldump -u$user -p$pass -S$socket -x -B -F -R"
dbList=`$myCmd -e "show databases;"|sed 1d|egrep -v "_schema"`
[ ! -d $backPath ] && mkdir -p $backPath

for dbname in $dbList; do
	$myDump $dbname|gzip >$backPath/${dbname}_$(date +%F).sql.gz
done


如何实现对MySQL数据库进行分库加分表备份,请用脚本实现
1.编程思想
    mysqldump db1 t1 > db1_t1.sql
    mysqldump db1 t2 > db1_t2.sql

    mysqldump db2 t1 > db2_t1.sql
    mysqldump db2 t2 > db2_t2.sql

[jack@me ~]$ cat backmysqltable.sh 
#!/bin/bash
#[ -f /etc/init.d/functions ] . /etc/init.d/functions

backPath=/srv/backup
user=root
pass=666
socket=/var/lib/mysql/mysql.sock
#socket1=
myCmd="mysql -u$user -p$pass -S$socket"
myDump="mysqldump -u$user -p$pass -S$socket -x -F -R"   ## 备份表示不用B参数
dbList=`$myCmd -e "show databases;"|sed 1d|egrep -v "_schema"`
[ ! -d $backPath ] && mkdir -p $backPath
tbList=`$myCmd -e "show tables from mysql;"|sed 1d`

for dbname in $dbList; do
	for tbname in $tbList; do
		mkdir -p $backPath/${dbname}
		$myDump $dbname $tbname|gzip >$backPath/${dbname}/${dbname}_${tbname}_$(date +%F).sql.gz
	done
done



mariadb.service
MySQL多实例启动脚本
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


写一个脚本解决DOS攻击生产案例
提示:根据web日志或者或者网络连接数,监控当某个IP并发连接数或者短时内PV达到100,即调用防火墙命令封掉该IP,监控频率每隔3分钟。防火墙命令为: 
iptables -I INPUT -s 10.0.1.10 -j DROP.

[jack@me ~]$ echo `sudo awk -F '[ ]+' '{print $2}' /var/log/messages|tr "\n" "+"|sed 's/12+$/12/g'` |bc
53564

[jack@me ~]$ sudo awk -F '[ ]+' '{print $2}' /var/log/messages|sort|uniq -c
     78 10
   3112 11
   1546 12
[jack@me ~]$ sudo awk -F '[ ]+' '{print $2}' /var/log/messages|sort|uniq -c|sort -rn -k1
   3112 11
   1546 12
     78 10

思路:
1.统计大于100PV的ip
2.知道封ip的命令
[jack@me ~]$ cat web_dos.sh     ## 放定时任务
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


（生产实战案例）：监控MySQL主从同步是否异常，如果异常，则发送短信或者邮件给管理员。提示：如果没主从同步环境,可以用下面文本放到文件里读取来模拟：
阶段1：开发一个守护进程脚本每30秒实现检测一次。
阶段2：如果同步出现如下错误号（1158,1159,1008,1007,1062），则跳过错误。
阶段3：请使用数组技术实现上述脚本（获取主从判断及错误号部分）
mysql -uroot -p'oldboy' -S /data/3307/mysql.sock -e "show slavestatus\G;"
*************************** 1. row ***************************
               Slave_IO_State:Waiting for master to send event
                  Master_Host:10.0.0.179   #当前的mysql master服务器主机
                  Master_User: rep
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File:mysql-bin.000013
         Read_Master_Log_Pos: 502547
               Relay_Log_File:relay-bin.000013
                Relay_Log_Pos:251
        Relay_Master_Log_File:mysql-bin.000013
             Slave_IO_Running:Yes
           Slave_SQL_Running: Yes
              Replicate_Do_DB: 
         Replicate_Ignore_DB: mysql
          Replicate_Do_Table: 
      Replicate_Ignore_Table: 
     Replicate_Wild_Do_Table: 
 Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
         Exec_Master_Log_Pos: 502547
              Relay_Log_Space:502986
              Until_Condition:None
               Until_Log_File: 
                Until_Log_Pos: 0
          Master_SSL_Allowed: No
          Master_SSL_CA_File: 
          Master_SSL_CA_Path: 
              Master_SSL_Cert: 
           Master_SSL_Cipher: 
               Master_SSL_Key: 
       Seconds_Behind_Master: 0   #和主库比同步延迟的秒数，这个参数很重要
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error:


stop slave sql_thread;

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


禁ping (1 禁 0 放)
/proc/sys/net/ipv4/icmp_echo_ignore_all

sudo nmap -sS 10.0.0.13



```

