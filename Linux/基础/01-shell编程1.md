# 1. shell编程基础一

```bash
1.必须root才能执行脚本,否则退出
2.成功切换目录(cd /var/log), 否则退出
3.清除日志(cat /dev/null >messges), 清理成功
4.echo 输出


[jack@me ~]$ mkdir tt
[jack@me ~]$ [ -f tt ] && echo 1 ||echo 0
0
[jack@me ~]$ [ -d tt ] && echo 1 ||echo 0
1
[jack@me ~]$ [ -e tt ] && echo 1 ||echo 0
1
[jack@me ~]$ rm -r tt
[jack@me ~]$ [ -e tt ] && echo 1 ||echo 0
0


[jack@me ~]$ touch ff
[jack@me ~]$ ll
total 0
-rw-rw-r--. 1 jack jack 0 Nov  3 10:15 ff
[jack@me ~]$ [ -r ff ] && echo 1 ||echo 0
1
[jack@me ~]$ [ -w ff ] && echo 1 ||echo 0
1
[jack@me ~]$ [ -x ff ] && echo 1 ||echo 0
0
[jack@me ~]$ chmod u+x ff 
[jack@me ~]$ [ -x ff ] && echo 1 ||echo 0
1
[jack@me ~]$ chmod 000 ff
[jack@me ~]$ ll
total 0
----------. 1 jack jack 0 Nov  3 10:15 ff
[jack@me ~]$ [ -r ff ] && echo 1 ||echo 0
0
[jack@me ~]$ sudo [ -r ff ] && echo 1 ||echo 0  ##root依然有可读权限.
1


[jack@me ~]$ [ -z "" ] && echo 1 ||echo 0
1
[jack@me ~]$ [ -z "haha" ] && echo 1 ||echo 0
0
[jack@me ~]$ [ ! -z "haha" ] && echo 1 ||echo 0
1
[jack@me ~]$ [ -n "haha" ] && echo 1 ||echo 0
1
[jack@me ~]$ [ -n "" ] && echo 1 ||echo 0
0


比较值两段,要有空格.


df -h |awk -F '[ ]+' 'NR==1;NR==6 {print $0}'

进程取行更高效
netstat -lnt |grep 22 |wc -l


以MySQL为例
    双多分支if条件句举例。
    范例1:用if双分支实现对Nginx或MysQL服务是否正常进行判断,
    使用进程数、端口、URL的方式判断,如果进程没起,把进程启动。
    课后作业:
    1、监控web服务是否正常,不低于5种思路.
    2、监控db服务是否正常,不低于5种思路.

    方法: web和db共同的方法:
    1.端口
      本地:netstat, ss, lsof
      远程: telnet, nmap, nc
    2.进程(本地)
    3.wget, curl(http方式,判断根据返回值或者返回内容)
    4.header(http) (http方式,根据状态码判断)
    5.数据库特有,通过mysql客户端连接,根据返回值或返回内容
    6.通过PHP,Java程序url方式监控mysql(此方式最接近用户访问,效果好)
    报警的最佳方式不是服务是否开启,而是网站的用户是否还访问正常.

    netstat -lntup |grep 3306
    ss -lntup |grep 3306
    lsof -i :3306 |grep 3306 |wc -l
    ps -ef |grep mysql |grep -v |wc -l
    mysql -uroot -p666 -e "select version();" &>/dev/null


    本地
    不要用 [ " " -eq 3306 ]
    要用   [ " " = 3306 ]
    用字符串进行比较,不要取值比较,而是wc -l 取行, 过滤, 计算
    if [ "netstat -lnt |grep 3306 |awk -F '[ :]+' '{print $5}'" -eq 3306 ]
    if [ "netstat -lnt |grep 3306 |awk -F '[ :]+' '{print $5}'" = "3306" ]

    if [ `ps -ef |grep mysql |grep -v grep |wc -l` -gt 0 ]
    if [ `netstat -lntup |grep mysqld |wc -l` -gt 0 ]
    if [ `lsof -i :3306 |wc -l` -gt 0 ]

    远程
    if [ `nc -w 2 10.0.0.7:3306 &>/dev/null && echo 1 |grep 1 |wc -l` -gt 0 ]
    if [ `nmap 10.0.0.7 -p 3306 2>/dev/null |grep open |wc -l` -gt 0 ]
      then
        echo "mysql is running"
    else
        echo "mysql is stopped"
        systemctl start mysqld


以web为例
if [ "`curl -I -s -o /dev/null -w "%{http_code}\n" http://10.0.0.7`" = "200" ]
if [ `curl -I http://10.0.0.7 2>/dev/null |head -1|egrep "200|301|302"|wc -l` -eq 1 ]

wget -T 10 -q --spider http://10.0.0.7 &> /dev/null

curl -s http://10.0.0.7 2>/dev/null
if [ $? -eq 0 ]; then
if [ "`curl -s http://10.0.0.7 2>/dev/null && echo $?`" = "0" ]; then
if [ "`curl -s http://10.0.0.7`" = "bbs" ]; then
    echo "httpd is running"
else
    echo "httpd is stopped"
    systemctl start nginx
fi

1. Linux中一定会有参数,取想要的值的;
2. 只有两种输出,标准输出和错误输出
    [jack@me ~]$ curl -I -s www.baidu.com |head -1
    HTTP/1.1 200 OK
    [jack@me ~]$ curl -I www.baidu.com 2>/dev/null |head -1 |grep 200
    HTTP/1.1 200 OK
    [jack@me ~]$ curl -I www.baidu.com 2>/dev/null |head -1 |grep 200 |wc -l
    1
    [jack@me ~]$ curl -I -s -o /dev/null -w "%{http_code}\n" www.baidu.com
    200

比较深度的监控是由开发写的,而不是运维,运维不懂业务.
运维只是监控服务器是否正常.


积极向上的人生态度,做事的态度,团队精神.
不需要什么都会,就因为不全都会才要12的,啥都会了就不会只要12了.


  三种清空数据的方法
    [jack@Jack ~]$ >test.log 
    [jack@Jack ~]$ echo >test.log 
    [jack@Jack ~]$ cat /dev/null >test.log 


    . 和 source 使脚本在父shell中执行,不会开启新的子shell
    sh 和 bash 使脚本在子shell中执行,会开启新的子shell

  取变量或字符串的长度方法
    echo $a | wc -L 
    echo ${#a} 取变量的子串

    [jack@me ~]$ a=haha
    [jack@me ~]$ echo $a
    haha
    [jack@me ~]$ echo ${#a}
    4
    [jack@me ~]$ echo $a |wc -L
    4
    [jack@me ~]$ expr length "$a"
    4



shell的特殊变量
    $0 文件名及路径
    $1 $2 ${n} 传的参数
    $# 传了几个参数
    $$ 取当前shell的进程号PID
    $? 一个脚本的退出状态码的值
        0 表示运行成功
        2 权限拒绝
        1~125 找到该命令了，但是无法执行
        126 未找到要运行的命令
        >128 命令被系统强制结束
    $* 将所有的命令行所有参数视为单个字符串; 例: "$1$2$3"
    $@ 将命令行每个参数视为单独的字符串; 例: "$1" "$2" "$3"


    [jack@me ~]$ set -- "it is" really ok.
    [jack@me ~]$ echo $#
    3

    [jack@me ~]$ for i in $*;do echo $i;done
    it
    is
    really
    ok.
    [jack@me ~]$ for i in "$*";do echo $i;done
    it is really ok.

    [jack@me ~]$ for i in $@;do echo $i;done
    it
    is
    really
    ok.
    [jack@me ~]$ for i in "$@";do echo $i;done
    it is
    really
    ok.
    [jack@me ~]$ for i;do echo $i;done      ## for i <=> for i in "$@"
    it is
    really
    ok.


bash内部变量
    有些内部命令在目录列表时是看不见的,它们由Shell本身提供,常用的内部命令有:
    echo, eval, exec, export, readonly, read, shift, wait,exit和点(.)

    echo
    变量名表将变量名表指定的变量显示到标准输出。
    evalargse
    读入参数args,并将它们组合成一个新的命令,然后执行。
    exec
    命令参数当Shell执行到exec语句时,不会去创建新的子进程,而是转去执行指定的命令,
    当指定的命令执行完时,该进程(也就是最初的Shell)就终止了,
    所以Shell程序中exec后面的语句将不再被执行。
    export变量名=value
    Shell可以用export把它的变量向下带入子Shell,从而让子进程继承父进程中的环境变量。但子Shell不能用export把它的变量向上带入父Shell

查看字符个数
    [jack@me ~]$ i="how are you"
    [jack@me ~]$ echo ${i}
    how are you
    [jack@me ~]$ echo ${#i}
    11
    [jack@me ~]$ echo $i |wc -m
    12
    [jack@me ~]$ echo $i |cat -E
    how are you$

切片
    [jack@me ~]$ echo ${i:4}
    are you
索引
    [jack@me ~]$ echo ${i:4:3}
    are
    [jack@me ~]$ echo $i |cut -c 5-7
    are
删除指定字符(从头开始)
    [jack@me ~]$ echo ${i#how}
    are you
删除指定字符(从尾开始)
    [jack@me ~]$ echo ${i%you}
    how are
替换
    [jack@me ~]$ echo ${i/how/who}
    who are you


[jack@me ~]$ for i in `seq 4`;do echo stu_999_"$i"_finished.jpg;done
stu_999_1_finished.jpg
stu_999_2_finished.jpg
stu_999_3_finished.jpg
stu_999_4_finished.jpg
[jack@me ~]$ for i in `seq 4`;do echo stu_999_"$i"_finished.jpg;done |awk -F '_finished' '{print $1$2}'
stu_999_1.jpg
stu_999_2.jpg
stu_999_3.jpg
stu_999_4.jpg

修改文件名
    ls |awk -F 'finished' '{print "mv "$0" "$1$2" "}' |bash

    rename "finished" "" *

    rename .jpg .JPG *.jpg

    rename 999 666 *.JPG


批量创建文件
    for i in `seq 4`;do echo stu_999_"$i"_finished.jpg;done |xargs touch

批量修改文件名
    #!/bin/bash
    for i in `ls *.jpg`
    do
        #echo "$i"
        #echo ${i%finished*}.jpg
        mv "$i" `echo ${i%finished*}.jpg`
    done

批量替换扩展名
    #!/bin/bash
    for i in `ls *.jpg`
    do 
        #echo "$i" #echo ${i%finished*}.jpg
        mv "$i" `echo ${i/jpg/JPG}`
    done



预定义后补变量
    ${value:-word}
    [jack@me ~]$ j=${test:-unset}
    [jack@me ~]$ echo $j
    unset
    [jack@me ~]$ echo $test

    [jack@me ~]$ test="aa"
    [jack@me ~]$ j=${test:-unset}
    [jack@me ~]$ echo $j
    aa


预定于变量
    ${value:=word}
    [jack@me ~]$ j=${test:=unset}
    [jack@me ~]$ echo $j
    unset
    [jack@me ~]$ echo $test
    unset
    [jack@me ~]$ test="stable"
    [jack@me ~]$ echo $j
    unset

用于捕捉由于变量未定义而导致的错误,并退出程序
    ${value:? "word"}
    [jack@me ~]$ echo ${value:? "not defined"}
    -bash: value:  not defined
    [jack@me ~]$ value=6
    [jack@me ~]$ echo ${value:? "not defined"}
    6
    [jack@me ~]$ unset value
    [jack@me ~]$ echo ${value:? "not defined"}
    -bash: value:  not defined


如果变量名存在且非空,则返回word,否则返回空
用于测试变量是否存在
    ${value:+word}
    [jack@me ~]$ a=${value:+9}
    [jack@me ~]$ echo $a

    [jack@me ~]$ value=tmp
    [jack@me ~]$ echo $a

    [jack@me ~]$ a=${value:+9}
    [jack@me ~]$ echo $a
    9


变量没定义就用-号后面的替代(删除操作利用这个功能好用)
    ${value-word}
    [jack@me ~]$ cat find_del.sh 
    #Path=/server/backup
    find ${Path-/tmp/} -type f
    [jack@me ~]$ sudo sh -x find_del.sh 
    + find /tmp -type f
    /tmp/.viminfo
    /tmp/ks-script-3c6qp1kl

    [jack@me ~]$ cat find_del.sh 
    #Path=/server/backup
    find ${Path:=/tmp/} -type f
    [jack@me ~]$ sudo sh -x find_del.sh 
    + find /tmp -type f
    /tmp/.viminfo
    /tmp/ks-script-3c6qp1kl


计算变量值字符串长度的方法
    chars=`seq -s " " 100`

    [jack@me ~]$ echo ${#chars}
    291
    [jack@me ~]$ echo $chars |wc -m
    292
    [jack@me ~]$ echo $(expr length "$chars")
    291

测试计算速度
    [jack@me ~]$ time for i in $(seq 10000);do count=${#chars};done
    real	0m0.035s
    user	0m0.033s
    sys	0m0.003s

    [jack@me ~]$ time for i in $(seq 10000);do count=`echo expr length "${chars}"`;done;
    real	0m9.867s
    user	0m2.810s
    sys	0m8.002s

    [jack@me ~]$ time for i in $(seq 10000);do count=`echo ${chars}|wc -m`;done; 
    real	0m29.775s
    user	0m9.708s
    sys	0m27.143s

返回了一个奇怪的结果
    [jack@me ~]$ cat who.sh 
    user=`whoami`
    [jack@me ~]$ sh who.sh 
    [jack@me ~]$ echo $user
                            ##这里输出了空
    [jack@me ~]$ 


  shift 每执行一下列表就向前移动一位


如何进行整数的计算?
    (()) let expr expr$[] bc typeset $[] awk

  变量在前,先输出变量值,变量在后,就是先运算后输出变量的值。
    [jack@me ~]$ ((a=5*6))
    [jack@me ~]$ echo $a
    30
    ## a在运算符之前,结果先出现a原来的值,但此时a已经递加1了
    [jack@me ~]$ echo $((a++))  
    30
    [jack@me ~]$ echo $a
    31
    [jack@me ~]$ echo $((a--))
    31
    [jack@me ~]$ echo $a
    30
    ## a在运算符之后,先递增在加a
    [jack@me ~]$ echo $((++a)) 
    31
    [jack@me ~]$ echo $a
    31
    [jack@me ~]$ echo $((--a))
    30


    [jack@me ~]$ b=$((5*8))
    [jack@me ~]$ echo $b
    40
    [jack@me ~]$ echo $((6*6))
    36

比较大小,真为1,假为0
    [jack@me ~]$ echo $((2>1))
    1
    [jack@me ~]$ echo $((2<1))
    0
    [jack@me ~]$ echo $((2==1))
    0
    [jack@me ~]$ echo $((2!=1))
    1

计算器v1.0
    [jack@me ~]$ cat cal.sh 
    #!/bin/bash
    a=$1
    b=$2
    echo "$a + $b = $(($a + $b))"
    echo "$a - $b = $(($a - $b))"
    echo "$a * $b = $(($a * $b))"
    echo "$a / $b = $(($a / $b))"
    echo "$a ** $b = $(($a ** $b))"
    echo "$a % $b = $(($a % $b))"


  let
    [jack@me ~]$ i=2
    [jack@me ~]$ let i=i+8
    [jack@me ~]$ echo $i
    10
  

  expr(运算符两侧都有空格,乘法号需要转义)
    [jack@me ~]$ expr 4 + 3
    7
    [jack@me ~]$ expr 4 - 3
    1
    [jack@me ~]$ expr 4 \* 3
    12
    [jack@me ~]$ expr 4 / 3
    1
    [jack@me ~]$ expr 4 % 3
    1

    [jack@me ~]$ expr $[5+2]
    7
    [jack@me ~]$ expr $[5-2]
    3
    [jack@me ~]$ expr $[5*2]
    10
    [jack@me ~]$ expr $[5/2]
    2
    [jack@me ~]$ expr $[5%2]
    1

  参考 ssh-copy-id
    if expr "$1" : ".*\.pub"; then

    输出字符
    [jack@me ~]$ expr "11.pub" :  ".*\.pub"
    6
    [jack@me ~]$ expr "11.txt" :  ".*\.txt"
    6

    判断是否已xx结尾(got you )
    [jack@me ~]$ expr "11.txt" :  ".*\.txt" &>/dev/null &&echo 1 || echo 0
    1
    [jack@me ~]$ expr "11.tx" :  ".*\.txt" &>/dev/null &&echo 1 || echo 0
    0
    [jack@me ~]$ expr "11" :  ".*\.txt" &>/dev/null &&echo 1 || echo 0
    0
    [jack@me ~]$ expr "txt" :  ".*\.txt" &>/dev/null &&echo 1 || echo 0
    0
    [jack@me ~]$ expr ".txt" :  ".*\.txt" &>/dev/null &&echo 1 || echo 0
    1

判断整数
    [jack@me ~]$ cat judge_int.sh 
    #!/bin/bash
    while true
    do
        read -p "Pls input: " a
        expr $a + 0 &>/dev/null
        [ $? -eq 0 ] && echo "int" ||echo "chars"

        [ "$a" = "exit" ] && {
            echo "bye !"
            exit 0
        }
    done

  BC
    [jack@me ~]$ seq -s "+" 100 |bc
    5050
    保留两位小数
    [jack@me ~]$ echo "scale=2;5/3" |bc
    1.66
    转换为二进制
    [jack@me ~]$ echo "obase=2;8" |bc
    1000

  typeset
    [jack@me ~]$ typeset -i a=1 b=2
    [jack@me ~]$ a=a+b
    [jack@me ~]$ echo $a
    3

  $[]
    [jack@me ~]$ echo $[2+2]
    4
    [jack@me ~]$ echo $[2**2]
    4
    [jack@me ~]$ echo $[2**3]
    8
    [jack@me ~]$ echo $[2*3]
    6




yangTriangle
    1.
    if (test -z $1); then
        read -p "Input a num: " num
    else
        num=$1
    fi

    # judge int
    [ -n "`echo $num|sed 's/[0-9]//g'`" ] && {
        echo the num must be int; bye !
        exit 1
    }

    # require < 10
    [ ! $num -lt 10 ] && {
        echo "the num you input must be in range 1 to 9; bey !"
        exit 1
    }

    # bagin
    a[0]=1
    for ((i=0;i<=num;i++))
    do
        for ((j=$i;j>0;j--))
        do
            ((a[$j]+=a[$j-1]))
        done
            for ((j=0;j<=$[num-$i];j++))
            do
                if [ $num -le 6 ]; then
                    echo -en "\t"
                else
                    echo -n "    "
                fi
            done
        for ((j=0;j<=$i;j++))
        do
            if [ $num -le 6 ]; then
                echo -en "\t\t"${a[$j]}
            else
                echo -n ${a[$j]}"        "
            fi
        done
    echo
    done

    2.
    #!/bin/bash
    if (test -z $1); then
        read -p "Input a num: " num
    else
        num=$1
    fi

    i=1
    while [ $i -le $num ]
    do
        j=1
        while [ $j -le $i ]
        do
            k=$[i-1]
            l=$[j-1]
            if [ $j -eq $i ] || [ $j -eq 1 ]; then
                declare SUM_${i}_${j}=1
            else
                declare A=$[SUM_${k}_$j]
                declare B=$[SUM_${k}_$l]
                declare SUM_${i}_$j=`expr $A + $B`
            fi
            echo -en $[SUM_${i}_$j]" "
            let j++
        done
        echo
        let i++
    done

    3.

    yangTriangle () {
        vector[0]=1
        echo ${vector[0]}
        for ((row=1;row<=${1};++row))
        do
            vector[row]=1
            for ((col=row-1;col>0;--col))
            do
                ((vector[col]+=vector[col-1]))
            done
            for ((col=0;col<=row;++col))
            do
                echo -n "${vector[col]}"
            done
            echo
        done

    }
    yangTriangle ${1}



test
    [jack@me ~]$ test -f who.sh && echo 1||echo 0
    1
    [jack@me ~]$ test -d who.sh && echo 1||echo 0
    0
    [jack@me ~]$ test ! -d who.sh && echo 1||echo 0
    1
    [jack@me ~]$ [ -f file ] && echo 1||echo 0
    0
    [jack@me ~]$ [ -f who.sh ] && echo 1||echo 0
    1
    [jack@me ~]$ [ ! -f who.sh ] && echo 1||echo 0
    0
    [jack@me ~]$ [ -d who.sh ] && echo 1||echo 0
    0
    [jack@me ~]$ [ ! -d who.sh ] && echo 1||echo 0
    1

    [jack@me ~]$ [ -f who.sh -o -d dir ] && echo 1||echo 0
    1
    [jack@me ~]$ [[ -f who.sh || -d dir ]] && echo 1||echo 0
    1
    [jack@me ~]$ [[ -f who.sh && -d dir ]] && echo 1||echo 0
    0
    [jack@me ~]$ [ -f who.sh -a -d dir ] && echo 1||echo 0
    0

    [jack@me ~]$ [ 2 -gt 1 ] && echo 1||echo 0
    1
    [jack@me ~]$ [ 2 -lt 1 ] && echo 1||echo 0
    0
    [jack@me ~]$ [ 2 -eq 1 ] && echo 1||echo 0
    0
    [jack@me ~]$ [ 2 -ge 1 ] && echo 1||echo 0
    1
    [jack@me ~]$ [ 2 -le 1 ] && echo 1||echo 0
    0
    [jack@me ~]$ [ 2 -ne 1 ] && echo 1||echo 0
    1
    ## 算术运算符要在bash里使用
    [jack@me ~]$ [[ 2 > 1 ]] && echo 1||echo 0
    1
    [jack@me ~]$ [[ 2 < 1 ]] && echo 1||echo 0
    0
    [jack@me ~]$ [[ 2 == 1 ]] && echo 1||echo 0
    0
    [jack@me ~]$ (( 2 >= 1 )) && echo 1||echo 0
    1
    [jack@me ~]$ (( 2 <= 1 )) && echo 1||echo 0
    0
    [jack@me ~]$ [[ 2 != 1 ]] && echo 1||echo 0
    1

    [jack@me ~]$ [ -r who.sh ] && echo 1||echo 0
    1
    [jack@me ~]$ [ -w who.sh ] && echo 1||echo 0
    1
    [jack@me ~]$ [ -x who.sh ] && echo 1||echo 0
    0


  字符串比较一定要加双引号""
    [jack@me ~]$ [ -z "" ] && echo 1||echo 0
    1
    [jack@me ~]$ [ -n "" ] && echo 1||echo 0
    0
    [jack@me ~]$ [ -n "dfs" ] && echo 1||echo 0
    1

  比较字符串长度
    [jack@me ~]$ f1=/etc/services
    [jack@me ~]$ f2=/etc/rc.local
    [jack@me ~]$ [ "$f1" = "$f2" ] && echo 1||echo 0
    0
    [jack@me ~]$ [ "${#f1}" = "${#f2}" ] && echo 1||echo 0
    1

  判断输入的是否位数字(把非数字部分替换为空)
    [jack@me ~]$ [ -z "`echo 11k |sed 's/[0-9]//g'`" ] && echo 1||echo 0
    0
    [jack@me ~]$ [ -n "`echo 11k |sed 's/[0-9]//g'`" ] && echo 1||echo 0
    1
    [jack@me ~]$ [ -z "`echo 11 |sed 's/[0-9]//g'`" ] && echo 1||echo 0
    1

    [jack@me ~]$ num=666
    [jack@me ~]$ num1=66d
    [jack@me ~]$ echo ${num/[^0-9]/}
    666
    [jack@me ~]$ echo ${num1/[^0-9]/}  ## 原理: 把非数字部分替换为空
    66
    [jack@me ~]$ echo ${num1/[^0-9]/9}
    669
    [jack@me ~]$ [ -n "$num" -a "$num" = "${num/[^0-9]/}" ] && echo 1||echo 0
    1
    [jack@me ~]$ [ -n "$num" -a "$num" = "${num1/[^0-9]/}" ] && echo 1||echo 0
    0

    expr $1 + 0 &>/dev/null
    [ $? -eq 0 ] && echo "it is int"


检查进程是否运行  ## 要用以下方式,简单易懂,不要取其端口,麻烦不实用
    #!/bin/bash
    ## PORT=`netstat -lnt|egrep "0:22" |awk -F '[ :]+' '{print $5}'` 

    PORT=`netstat -lnt|grep 22|wc -l`
    sshdProcessNum=`ps -ef |grep sshd |grep -v grep |wc -l`
    #if [ $PORT -ge 1 ]; then 
    if [ $PORT -ge 2 -a $sshdProcessNum -ge 3 ]; then
        echo "it is running"
    else
        echo "it is dead."
    fi


  要一直杀杀到进程不为0为止退出循环
    while true
    do
        killall mysqld &>/dev/null
        [ $? -ne 0 ] && break
        sleep
    done

    方式1
    mysql -uroot -p666 -e "select version();" &>/dev/null
    方式2
    mysqlStatus=`mysql -uroot -p666 -e "select version();" &>/dev/null`
    if [ $? -eq 0 ]; then
        echo "db is running"
    else
        systemctl start mysqld
    fi


    telnet 非交互式
    echo -e "\n" |telnet baidu.com 80 |grep Connected|wc -l


case结构条件
    case "var" in
      1)
        order
        ;;
      2)
        order
        ;;
      *)
        order
    esac

[jack@me ~]$ cat choose_color_case_plus.sh
    #!/bin/bash

    color() {

    RED='\E[1;31m'
    GREEN='\E[1;32m'
    YELLOW='\E[1;33m'
    BLUE='\E[1;34m'
    PINK='\E[1;35m'
    RES='\E[0m'

    if [ $# -ne 2 ]; then
        echo -e "${PINK}USAGE:$RES sh $0 content {red|green|yellow|blue}"
        exit 1
    fi
        
    case "$2" in
      red)
        echo -e "$RED$1$RES"
        ;;
      green)
        echo -e "$GREEN$1$RES"
        ;;
      yellow)
        echo -e "$YELLOW$1$RES"
        ;;
      blue)
        echo -e "$BLUE$1$RES"
    esac
    }

    color apple red
    color leaf green



  while循环
    while order
    do
        order
    done

    [jack@me ~]$ cat while_de.sh 
    #!/bin/bash
    i=10
    while [ $i -gt 0 ]
    do
        let j=i
        let i--
        sleep 1
        echo $j
    done

    i=2
    while [ $i -le 10 ]
    do
        let j=i
        let i++
        sleep 1
        echo $j
    done


until循环
    #!/bin/bash

    if [ $# -ne 1 ]; then
        echo "USAGE: sh $0 num"
        exit 1
    fi

    i=$1
    while [ $i -gt 0 ]
    do
        let j=i
        let i--
        sleep 1
        echo $j
    done

    i=2
    until [ $i -gt 5 ]
    do
        let j=i
        let i++
        sleep 1
        echo $j
    done


while循环读入文件的方式(一行一行的读,速度相当的慢)
    [jack@me ~]$ cat size.sh
    #!/bin/bash
    FILE=~/mesg.log
    exec < $FILE    ##方式1
    sum=0

    ##cat ${FILE} |while read line; do  ##方式2
    while read line; do
        size=`echo $line|awk -F '[ ]+' '{print $2}'`
        [ "$size" = "-" ] && continue
        let sum=sum+$size
        #((sum = sum + $size))
    done
    ##done < $FILE  ##方式3
    [ -n $sum ] && echo $sum


for循环
    for var in [取值列表]
    do
        order
    done

    for ((i=1; i<n; i++))
    do
        order
    done


    for ((i=1; i<=5; i++)); do
    #for i in `seq 5`; do
        echo 1
        sleep 1
    done

    [jack@me ~]$ cat for.sh 
    #!/bin/bash

    size=`awk -F '[ ]+' '{print $2}' $1`
    sum=0
    for num in $size; do
        [ -n "$num" -a "$num" = "${num/[^0-9]/}" ] || continue
        let sum=$num+sum
    done

    echo "this file is $sum bytes;$(($sum/1024))KB"




随机数
    [jack@me ~]$ echo $RANDOM
    10016680

    [jack@me ~]$ echo $(date +%N)
    668018320

    [jack@me ~]$ openssl rand -base64 8
    20TfvhGRfJE=

    [jack@me ~]$ head /dev/urandom |cksum 
    2204713332 2640

    [jack@me ~]$ head /dev/random |cksum 
    1288731354 1888

    [jack@me ~]$ cat /proc/sys/kernel/random/uuid 
    00637bf1-6451-4f35-bed2-9f7814bf0842

    dnf install except
    mkpasswd -l 8

检查随机数重复性
    for i in `seq 20`; do echo $RANDOM|md5sum |cut -c 2-8;done|sort|uniq -c|sort -rn -k 1


#########################
## continue break exit ##
#########################
    [jack@me ~]$ cat break_continue_exit.sh 
    #!/bin/bash
    for ((i=0; i<=5; i++)); do
        if [ $i -eq 3 ]; then
            #continue
            #break
            exit
        fi
        echo "$i"
    done
    echo "finish"

    ##跳出当前for循环后,继续执行下次循环 参数n表示退到第n层继续循环
    continue 
        0
        1
        2
        4
        5
        finish

    ## 跳出整个for循环,执行循环之外的语句
    ## 参数n表示跳出循环的层数,如果省略n表示跳出整个循环。
    break 
        0
        1
        2
        finish

    ## 退出整个脚本, 并返回n, n也可以省略
    exit 
        0
        1
        2


函数,函数一定要在调用前定义,调用函数不用加括号


函数的传参
    [jack@me ~]$ cat check_url.sh 
    #!/bin/bash
    usage() {
        cat <<-EOF
            USAGE: sh $0 domain name
        EOF
    }

    if [ $# -ne 1 ]; then
        usage && exit 1
    fi

    check_url() {
        curl -I -s $1|head -1 && return 0||return 1
    }

    #check_url www.baidu.com
    check_url $1



##########
## 数组 ##
##########
    [jack@me ~]$ li=(1 2 3)
    [jack@me ~]$ echo ${#li[@]}
    3
    [jack@me ~]$ echo ${#li[*]}
    3
    [jack@me ~]$ echo ${li[0]}
    1
    [jack@me ~]$ echo ${li[1]}
    2
    [jack@me ~]$ echo ${li[2]}
    3
    [jack@me ~]$ echo ${li[@]}
    1 2 3
    [jack@me ~]$ echo ${li[@]}
    1 2 3
增加
    [jack@me ~]$ li[3]=4
    [jack@me ~]$ echo ${li[3]}
    4
    [jack@me ~]$ echo ${li[@]}
    1 2 3 4
修改
    [jack@me ~]$ li[0]=me
    [jack@me ~]$ echo ${li[@]}
    me 2 3 4
删除全部
    [jack@me ~]$ unset li
    [jack@me ~]$ echo ${li[@]}
删除单个元素
    [jack@me ~]$ unset li[0]
    [jack@me ~]$ echo ${li[@]}
    2 3
切片
    [jack@me ~]$ li=(1 2 3 4 5 6)
    [jack@me ~]$ echo ${li[@]}
    1 2 3 4 5 6
    从0开始取3个元素
    [jack@me ~]$ echo ${li[@]:0:3}
    1 2 3
替换
    li1=${li[@]/1/me}   ## 替换后要重新赋值给新的数组
    [jack@me ~]$ echo ${li1[@]}
    me 2 3 4 5

[jack@me ~]$ echo ${li[@]#1}
2 3 4 5

[jack@me ~]$ echo ${li[@]%5}
1 2 3 4
从前向后删
    [jack@me ~]$ li=(apple orange banana)
    [jack@me ~]$ echo ${li[@]}
    apple orange banana
    [jack@me ~]$ echo ${li[@]#a}
    pple orange banana
    [jack@me ~]$ echo ${li[@]#a*e}
    orange banana
从后向前删
    [jack@me ~]$ echo ${li[@]%a}
    apple orange banan
    [jack@me ~]$ echo ${li[@]%b*a}
    apple orange

查看key的值
    [jack@me ~]$ li=(1 2 3)
    [jack@me ~]$ echo ${!li[@]}
    0 1 2

key-value定义数组
    [jack@me ~]$ li=([0]=0 [1]=1 [2]=2 [3]=3)
    [jack@me ~]$ echo ${#li[@]}
    4
    [jack@me ~]$ echo ${li[@]}
    0 1 2 3

    [jack@me ~]$ li=(apple orange banana)
    [jack@me ~]$ for i in ${li[@]};do echo $i;done
    apple
    orange
    banana

    [jack@me ~]$ li[0]=a
    [jack@me ~]$ li[1]=b
    [jack@me ~]$ li[2]=c
    [jack@me ~]$ echo ${li[@]}
    a b c

li=($(ls))
[jack@me ~]$ for i in ${li[*]};do echo $i;done
add_net1.sh
add_net.sh
check_url1.sh
check_url.sh
continue_break_exit.sh
folder
func.sh
get_url_status.sh
muti_user_add.sh
muti_user_del.sh
scripts
user_pass.txt
[jack@me ~]$ echo "len: ${#li[*]}"
len: 12



dos2unix

export PS4='+${LINENO}'


信号
kill -l
trap -l

trap命令用于指定在接收到信号后将要采取的行动,信号的信息前面已经提到。
trap命令的一种常见用途是在脚本程序被中断时完成清理工作。历史上, 
shell总是用数字来代表信号,而新的脚本程序应该使用信号的名字,它们保存在用
##include命令包含进来的signal.h头文件中,在使用信号名时需要省略SIG前缀。
你可以在命令提示符下输入命令trap -l来查看信号编号及其关联的名称。
对于那些不熟悉信号的人们来说, "信号"是指那些被异步发送到一个程序的事件。
默认情况下,它们通常会终止一个程序的运行。
请记住,脚本程序通常是以从上到下的顺序解释执行的,所以必须在你想保护的那部分
代码以前指定trap命令。
如果要重置某个信号的处理条件到其默认值,只需简单的将command设置为-。如果要忽略某个信号,就把 command设置为空字符串""。


1) SIGHUP	 |挂起,通常因终端掉线或用户退出而引发
2) SIGINT    |中断,通常因按下Ctrl+C组合键而引发
3) SIGQUIT   |退出,通常因按下Ctrl+/组合键而引发
6) SIGABRT   |中止,通常因某些严重的执行错误而引发
14) SIGALRM  |报警,通常用来处理超时
15) SIGTERM  |终止,通常在系统关机时发送
20) SIGTSTP  |停止进程的运行,但该信号可以被处理和忽略,用户键入SUSP字符时
             |(通常是Ctrl-Z)发出这个信号


trap "" 2
trap ":" 2

stty -a

[jack@me ~]$ trap "" 2  ## 此时按Crtl C没反应(被屏蔽了)
[jack@me ~]$ trap -p
trap -- '' SIGINT
[jack@me ~]$ trap 2     ## 恢复信号
[jack@me ~]$ trap -p
[jack@me ~]$ ^C
[jack@me ~]$ trap "echo -n ' typing Crtl+C'" 2
[jack@me ~]$ ^C typing Crtl+C
[jack@me ~]$ ^C typing Crtl+C


## 取消
trap "" 1 2 3 15 20
## 恢复
trap 1 2 3 15 20
trap ":" 1 2 3 15 20





案例2 : shell跳板机(触发信号后屏蔽信号)
方法1:
    1)首先做好ssh key验证,见前面的ssh key免登陆验证。
    2)实现传统的远程连接菜单选择脚本 
    3)利用linux信号防止用户在跳板机上操作。
    4)用户登录后即调用脚本。

方法2: root连接服务器, expect每次重新建立ssh key.


## 声明局部变量
fun() {
    local i 
}

脚本名
start_xx.sh
stop_xx.sh
xx_mon.sh
xx_ctl.sh

目录
bin
conf
func


*************************** 1. row ***************************
               Slave_IO_State:Waiting for master to send event
                  Master_Host:10.0.0.69
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
       Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error:







```

