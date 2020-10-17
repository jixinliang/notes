# Linux 常用命令排序版

## A

```bash
alias	-- Define or display aliases

awk -- pattern scanning and processing language
    -F  --field-separator
    $0 is the whole record
    $1 is the first record  
    $2 is the second record
    $NF	the last field of line
    NF  The number of fields in  the  current input record
    NR  The  total  number  of  input records seen so far
    /regular expression/

```

## B

```bash
basename    -- strip directory and suffix from filenames 
	basename OPTION... NAME...
    -a  --multiple	support multiple arguments and treat each as a NAME
    -s  --suffix=SUFFIX		remove a trailing SUFFIX
    
bc  -- An arbitrary precision calculator language
	-q  --quiet		Do not print the normal GNU bc welcome

blkid   -- locate/print block device attributes
```

## C

```bash
cd  -- Change the current directory to dir
    cd  [dir]
    .   -- current directory
    ..  -- parent directory
    -   -- previous directory; is equivalent to $OLDPWD 
    ~   -- /home/user directory  | -- /root directory
    /   -- / directory

cat  -- concatenate files and print on the standard output 
	cat [OPTION]... [FILE]...
	-A   --show-all  equivalent to -vET
	-b  --number-nonblank
	-E  --show-ends		display $ at end of each line
	-n  -- number
	-s  --squeeze-blank		suppress repeated empty output lines
	-T  --show-tabs	display TAB characters as ^I

cp   -- copy files and directories
    cp [OPTION]... SOURCE... DIRECTORY
    -a  -- archive      same as -dR --preserve=all
    -d  same as --no-dereference --preserve=links
    -f  -- force
    -i  -- interactive   prompt before overwrite
    -p  same as --preserve=mode,ownership,timestamps
    -r  -- recursive    copy directories recursively
    -u  -- update
    -a <=> -dpr
    -H ?

chkconfig    -- updates and queries runlevel information for system services
    --add name
    --del name
    --list name
    --level levels

cut  - remove sections from each line of files
    cut OPTION... [FILE]...
    -b  --bytes=LIST
    -c  --characters=LIST
    -d  --delimiter=DELIM
    -f  --fields=LIST

chmod - change file mode bits
    chmod [OPTION]... MODE[,MODE]... FILE...
    -R  --recursive    change files and directories recursively

chown - change file owner and group
    chown [OPTION]... [OWNER][:[GROUP]] FILE...
    -h  --no-dereference
    -R  --recursive  operate on files and directories recursively

chattr   -- change file attributes on a Linux file system The format of a 	 
	symbolic mode is +-=[aAcCdDeijsStTu]
    a   --append  only
    i   --immutable
    -R  --Recursively
    
crontab  -- maintains crontab files for individual users
    -l	--Displays the current crontab on standard output
    -e	--Edits the current crontab
    -r	--Removes the current crontab
    -u	--Appends the name of the user whose crontab is to  be modified

chage	-- change user password expiry information
    chage [options] LOGIN
    -d  --lastday LAST_DAY
    -E  --expiredate
    Use chage -E 0 user command instead for full account locking
    Passing the number -1 as the EXPIRE_DATE will remove an
    account expiration date
    -l  --list
    -m  --mindays MIN_DAYS
    -M  --maxdays MAX_DAYS

chpasswd	-- update passwords in batch mode

comm - compare two sorted files line by line
    -1	suppress column 1 (lines unique to FILE1)
    -2	suppress column 2 (lines unique to FILE2)
    -3	suppress column 3 (lines that appear in both files)

curl	-- transfer a URL
    -I  --head (HTTP/FTP/FILE) Fetch  the  HTTP-header  only
    -s  --silent
    -o  --output <file>
    -w  --write-out <format>
    --connect-timeout <seconds>
    -m  --max-time <seconds>    allow the  whole operation to take.
        
createrepo  -- Create repomd (xml-rpm-metadata) repository
    -d  --database
    -p?
    -o?
    --update
    --basedir
```

## D

```bash
date -- print or set the system date and time
    -R  --rfc-2822  output date and time in RFC 2822  format.   
    Example: Mon, 07 Aug 2006 12:34:56 -0600
    -s  --set=STRING
    -d  --date=STRING
    %F     full date; same as %Y-%m-%d
    %T     time; same as %H:%M:%S

dmesg    -- print or control the kernel ring buffer
    dmesg [options]
    -C  --clear
    -c  --read-clear
    -H  --human
    -l  --level list
    -L  --color     Colorize important messages

dd   -- convert and copy a file
    if=FILE     --read from FILE instead of stdin
    of=FILE     --write to FILE instead of stdout
    bs=BYTES    --read and write up to BYTES bytes at a time
    count=N     --copy only N input blocks

df   -- report file system disk space usage
    df [OPTION]... [FILE]...
    -h  --human-readable
    -i  --inodes
    -T  --print-type
    
du   -- estimate file space usage
    Summarize disk usage of each FILE, recursively for directo-ries
    du [OPTION]... [FILE]...
    -a  --all
    -B  --block-size=SIZE
    -d  --max-depth=N
    -h  --human-readable
    -s  --summarize
    
dirname -- strip last component from file name

diff -- compare files line by line
    -c, -C NUM  --context[=NUM]
    -u, -U NUM  --unified[=NUM]
    -a  --text
    -r  --recursive
    -N  --new-file

dos2unix	-- DOS/Mac to Unix and vice versa text file format converter
    
dig	-- DNS lookup utility
```

## E

```bash
echo - display a line of text
    echo [SHORT-OPTION]... [STRING]...
    -n     do not output the trailing newline
    -e     enable interpretation of backslash escapes
    If -e is in effect, the following sequences are recognized:
    \\     backslash
    \a     alert (BEL)
    \b     backspace
    \f     form feed
    \n     new line
    \r     carriage return
    \t     horizontal tab
    \v     vertical tab

printf - format and print data

ethtool -- query or control network driver and hardware settings ethtool devname
    
exports -- NFS server export table
    
exportfs    -- maintain table of exported NFS file systems
    -a  --Export or unexport all directories
    -r  --Reexport all  directories
    -v  --Be verbose
    -o  --options
    
expr	-- evaluate expressions
	expr EXPRESSION
```

## F

```bash 
find -- search for files in a directory hierarchy
    -type c  File is of type c:
    b	block (buffered) special
    c	character (unbuffered) special
    d	directory
    p	named pipe (FIFO)
    f	regular file
    l	symbolic link
    s	socket
    D	door (Solaris)

    -mtime n
    File's data was last modified n*24 hours  ago.   See
    the  comments  for -atime to understand how rounding
    affects  the  interpretation  of  file  modification
    times.
    n   -- on n day
    +n  -- n days before
    -n  -- recent n days 

	-mmin n    File's data was last modified n minutes ago.

	-size n[cwbkMG]     
	--File uses n units of space
	-print  --print the full file name on the standard output,allowed by a 	
	newline
	-print0 --True; print the full file name on the standard  out-
	put,  followed  by  a null character
    -maxdepth levels
    -inum n     --File has inode number n
    -name pattern   --The  metacharacters (`*', `?', and `[]')
    -perm   --mode
    ! expr  --True  if  expr  is  false.  This character will also
     usually need protection from interpretation by the shell.
     -not expr  --Same as ! expr, but not POSIX compliant.
     -o  --or
     -a  --and
     -delete    --Delete files
     -newer file --File was modified more recently than file

file -- determine file type

free -- Display amount of free and used memory in the system
    free [options]
    -h  --human

fdisk   -- manipulate disk partition table
    -l  --List  the partition tables for the specified devices and then exit.
    -s partition... Print the size (in blocks) of each given partition

fsck    -- check and repair a Linux filesystem
	-A  --all

fold    -- wrap each input line to fit in specified width
    -w  --width=WIDTH
    -s  --spaces    break at spaces
    
fmt	-- simple optimal text formatter
    -c  --crown-margin   preserve indentation of first two lines
    -w  --width=WIDTH
    -p  --prefix=STRING
    reformat only lines beginning with STRING, reattaching the prefix to 
    reformatted lines

ftp -- Internet file transfer program
```

## G

```bash
grep, egrep, fgrep  -- print lines matching a pattern
    [OPTIONS] PATTERN [FILE...]

    Matching Control:
    -c  --count
    -i  --ignore-case
    -o  --only-matching
    -h  --no-filename
    -l  --files-with-matches
    -L  --files-without-match
    -v  --invert-match
    -w  --word-regexp
    Select only those lines containing matches that form
    whole  words
    -x  --line-regexp

    Context Line Control:
    -A  --after-context=NUM
    Print NUM lines of trailing context after matching lines.
    -B  --before-context=NUM
    Print NUM lines of trailing context before matching lines.
    -C  --context=NUM
    Print NUM lines of output context. 
    Places a linecontaining a group separator 
    between contiguous groups of matches

    [[:alpha:]]

gzip, gunzip, zcat - compress or expand files
    gzip [ -acdfhlLnNrtvV19 ] [-S suffix] [ name ...  ]
    gunzip [ -acfhlLnNrtvV ] [-S suffix] [ name ...  ]
    -c --stdout --to-stdout
    -d --decompress --uncompress
    -f --force
    -l --list
    -r --recursive
    -t --test   Test the integrity of a compressed file.
    -v --verbose
    
groupadd	-- create a new group
	-g  --gid GID
    
groups  -- print the groups a user is in
    
groff   -- front-end for the groff document formatting system
    
gpasswd -- administer /etc/group and /etc/gshadow
	-d  --delete user
```

## H

```bash
head -- output the first part of files Print the first 10 lines of each FILE 
	to standard output.
    head [OPTION]... [FILE]...
    -n  -- lines
    -q  --quiet, --silent
    never print headers giving file names
    -v  --verbose
    always print headers giving file names

hostname -- show or set the system's host name
    -I  --all-ip-addresses
    -s  --short     short host name

help -- Display information about builtin commands
    --help  -- Display Usage Information
    -m  --display usage in pseudo-manpage format

history  -- display the command history list with line numbers
    history -c
    history -d offset
    -c  --Clear the history list by deleting all the entries
    -d  --offset Delete the history entry at position offset
    
htpasswd    -- Manage user files for basic authentication
    -c  --Create  the  passwdfile
    -b  --Use batch mode; get the password from the command line
    rather than prompting for it
    -D  --Delete  user
```

## I

```bash
ip  -- show / manipulate routing, devices, policy routing and tunnels
    ip addr --Shows addresses assigned to all network interfaces
    ip route

ifup    -- bring a network interface up
	ifup CONFIG

ifdown  -- take a network interface down
	ifdown CONFIG
    
id - print real and effective user and group IDs
    
inotifywait -- wait for changes to files using inotify
    -m  --monitor
    -d  --daemon
    -r  --recursive
    -e <event>  --event <event>
    create  --file or directory created within watched directory
    delete	--file or directory deleted within watched directory
    close_write	file or directory closed, after being opened in
    writeable mode
    -q  --quiet    	Print less (only print events)
    --timefmt <fmt>
    --format <fmt>
```

## J

```bash
jobs    --Display status of jobs.
    -l  --lists process IDs in addition to the normal information
    -p	--lists process IDs only

join    -- join lines of two files on a common field
```

## K

```bash
kpartx - Create device maps from partition tables
    -a     Add partition mappings
    -d     Delete partition mappings
    -l     List partition mappings that would be added -a
```

## L

```bash
ls  -- list directory contents
    ls [OPTION]... [FILE]...
    -a  -- all
    -d  -- directory
    -h  -- human-readable
    -i  -- inode    print the index number of each file
    -l  -- use a long listing format
    -F  -- classify  append indicator (one of */=>@|) to entries
    -p  -- indicator-style=slash  append / indicator to directories
    -r  -- reverse
    -t  -- sort by modification time, newest first
    --time-style=long-iso

less - opposite of more
	-N or --LINE-NUMBERS

lsblk    -- list block devices
    lsblk [options] [device...]
    -f  --fs    Output info about filesystems
    -m  --perms Output info about device owner, group and mode
    -i  --ascii Use ASCII characters for tree formatting
    -o  --output list

last, lastb  -- show listing of last logged in users
	-F     Print full login and logout times and dates

lastlog  -- reports the most recent login of all users or of
    a given user
    -u  --user LOGIN

lsof -- list open files
    -i [i]  selects the listing of files any of whose Internet
    address matches the address specified in i

ln - make links between files
    -s  --symbolic
    -d  -F, --directory
    allow the superuser to attempt to hard link directo-
    ries  (note:  will  probably  fail  due  to   system restrictions, even 
    for the superuser)

lsattr   --  list  file attributes on a Linux second extended file system
    -R  --Recursively
    
lsmod   -- Show the status of modules in the Linux Kernel
    
losetup -- set up and control loop devices
    -a  --all
    -f  --find
    -P  --partscan  force kernel to scan partition table on newly   
    created loop device
```

## M

```bash
mkdir -- make directories  
    mkdir [OPTION]... DIRECTORY...
    Create the DIRECTORY(ies), if they do not already exist
    -p  -- parents
    no error if existing, make parent directories as needed
    -v  --verbose

mv  -- move (rename) files
    mv [OPTION]... SOURCE... DIRECTORY
    -f  --force     do not prompt before overwriting
    -i  --interactive   prompt before overwrite
    -t  --target-directory=DIRECTORY
    move all SOURCE arguments into DIRECTORY
    -u  --update

man  -- an interface to the on-line reference manuals
    1   Executable programs or shell commands
    5   File formats and conventions eg /etc/passwd
    8   System administration commands (usually only for root)

md5sum  -- compute and check MD5 message digest
    -c  --check

mkfs    -- build a Linux filesystem
    -t  --type type

mount   -- mount a filesystem
    mount --target /mountpoint
    -a  --all
    -t  --types vfstype
    -o  --options opts
    remount  --Attempt to remount an already-mounted filesystem
    ro  --Mount the filesystem read-only.
    rw  --Mount the filesystem read-write.

mkswap  -- set up a Linux swap area

make	-- GNU make utility to maintain groups of programs
	-j [jobs]  --jobs[=jobs]

modinfo -- Show information about a Linux Kernel module

modprobe    -- Add and remove modules from the Linux Kernel
    -r  --remove
    -v  --verbose

mysqldump - a database backup program
    -A  --all-databases
    -B  --databases Dump several databases
    -F  --flush-logs
    -x  --lock-all-tables
    -R  --routines
    -E  --events    Include Event Scheduler events for the dumped   
    databases in theoutput.


mysql  -- the MySQL command-line client
    -e statement    --execute=statement
    -S socket
    -h host
    
mysqld_safe - MySQL server startup script
```

## N

```bash
netstat
    netstat -lntup |grep ssh
    -a  --all
    -e  --extend
    -i  --interfaces
    -l  --listening     Show  only  listening  sockets
    -n  --numeric
    -t  --tcp
    -u  --udp
    -p  --program   Show  the  PID
    -r  --route     Display the kernel routing tables
        

nl   -- number lines of files

ntpdate -- set the date and time via NTP
    ntpdate http://ntp4.aliyun.com

nslookup    -- query Internet name servers interactively

nmap    -- Network exploration tool and security / port scanner
    -p  port ranges (Only scan specified ports)
    -sn (No port scan)
    -PS port list (TCP SYN Ping)
    -sS (TCP SYN scan)

nc(ncat)    -- Concatenate and redirect sockets
    -w  --wait <time>

newusers    -- update and create new users in batch
    username:passwd:uid:gid:full name:home_dir:shell

nginx   -- HTTP and reverse proxy server, mail proxy server
    -s signal   Send a signal to the master process.  The argument
    signal can be one of: stop, quit, reopen, reload. 
    -t  --Do not run, just test the configuration file.
    
nohup   -- run a command immune to hangups, with output to a non-tty
    
nslookup - query Internet name servers interactively
    nslookup [-option] [name | -] [server]
```

## O

```bash
od   -- dump files in octal and other formats
    -a  --same as -t a, select named characters, ignoring high-order bit
    -x  --same as -t x2, select hexadecimal 2-byte units
    -t  --format=TYPE
```

## P

```bash
pwd -- Print the name of the current working directory
    -- Print the absolute pathname of the current working directory
    -L  --logical	use  PWD  from environment, even if it contains sym-links
    -P  --physical
    avoid all symlinks

ps -- report a snapshot of the current processes
    ps [options]
    ps -ef
    To see every process on the system using standard syntax
    -e  --Select all processes
    -f  --Do full-format listing.
    -a  --Select all processes except both session leaders
    -u  --userlist
    -C  --cmdlist   Select by command name
    --no-headers    Print no header line at all
    
passwd - update user's authentication tokens
    --stdin
    This  option  is used to indicate that passwd should read the new password 
    from standard input, which can be a pipe
    -d  --delete
    -l  --lock
    -u  --unlock
    -n  --minimum DAYS
    -x  --maximum DAYS
    -w  --warning DAYS
    -i  --inactive DAYS
    -S  --status

paste   -- merge lines of files
    -d  --delimiters=LIST	reuse characters from LIST instead of TABs
    -s  --serial	paste one file at a time instead of in parallel

partprobe   -- inform the OS of partition table changes
    -d  --Don't update the kernel
    -s  --Show a summary of devices and their partitions

parted  -- a partition manipulation program

pstree  -- display a tree of processes
	-h  --Highlight the current  process  and  its  ancestors

printenv -- print all or part of environment

patch - apply a diff file to an original

pr  -- convert text files for printing
    -l  --length=PAGE_LENGTH
    -w  --width=PAGE_WIDTH

printf  -- format and print data

ping    -- send ICMP ECHO_REQUEST to network hosts
    -c  --count
    -i  --interval
    -s  --packetsize
    -W  --timeout

pkill - look up or signal processes based on name and other attributes
```

## Q

## R

```bash
rm  -- remove files or directories
    rm [OPTION]... FILE...
    -f  -- force
    -r  -- recursive
    -I  -- prompt  once  before removing more than three files,
    or when removing recursively

rpm -- RPM Package Manager
    -a  --all
    -c  --configfiles
    -q  --query
    -i  --info
    -i  --install
    -e  --erase
    -U  --upgrade
    -l  --list
    -f  --file FILE
    -h  --hash
    --reinstall
    -v  --verbose
    -d  --docfiles
    -p, --package

route    -- show / manipulate the IP routing table
    -n  --numeric

readlink  - print resolved symbolic links or canonical file names

rename  --Rename file
    rename [options] expression replacement file...  
    -v  --verbose    explain what is being done
    -s  --symlink    act on symlink target

rev -- reverse lines of a file or files
    
rsync  - a fast, versatile, remote (and local) file-copying tool
    -a  --archive
    -z  --compress
    -v  --verbose
    -u  --update
    --delete
    -e  --rsh=COMMAND   specify the remote shell to use
    -r  --recursive     recurse into directories
    -t  --times         preserve modification times
    -o  --owner         preserve owner (super-user only)
    -p  --perms         preserve permissions
    -P                  same as --partial --progress
    -g  --group         preserve group
    -R  --relative      use relative path names
    --bwlimit=RATE      limit socket I/O bandwidth
    --password-file=FILE    read daemon-access password from FILE
    --config=FILE           specify alternate rsyncd.conf file
    --exclude=PATTERN       exclude files matching PATTERN
    --exclude-from=FILE     read exclude patterns from FILE
    --delete                delete extraneous files from dest dirs
    --partial               keep partially transferred files
    --timeout=SECONDS       set I/O timeout in seconds

read    --Read a line from the standard input and split it into fields
    -e  --use Readline to obtain the line in an interactive shell
    -i text --Use TEXT as the initial text for Readline
    -p prompt	output the string PROMPT without a trailing newline before attempting 
    to read
    -s  --Silent  mode. If  input  is coming from a terminal,characters are not 
    echoed.
    -t seconds Timeout. Terminate input after seconds. read returns 
    a non-zero exit status if an input times out. 
    
rpcbind -- universal addresses to RPC program number mapper

rpcinfo -- report RPC information
	-p  --Probe rpcbind on host
    
restorecon  -- restore file(s) default SELinux security contexts
    -R  --recursively
    -v  --verbosity
```

## S

```bash
seq  -- print a sequence of numbers
    seq [OPTION]... LAST
    seq [OPTION]... FIRST LAST
    seq [OPTION]... FIRST INCREMENT LAST
    -f  --format=FORMAT
    FIRST, INCREMENT, and LAST
    %g
    -s  --separator=STRING
    use STRING to separate numbers (default: \n)
    -w  --equal-width
    equalize width by padding with leading zeroes
         

sed  -- stream editor for filtering and transforming text
    sed  [OPTION]...  {script-only-if-no-other-script}  [input-
    file]...

    -n, -- quiet, --silent suppress automatic printing of pattern space
    p   -- Print the current pattern space.
    l   -- List out line in a ``visually unambiguous'' form

    -e script   --expression=script
    --add the script to the commands to be executed

    -i  -- in-place      -- edit files in place
    -r  --regexp-extended
    -f  script-file, --file=script-file

    a text   --Append text, which has each embedded newline 
    preceded by a backslash
    i text   --Insert text, which has each  embedded  newline preceded by abackslash
    d   -- Delete pattern space.  Start next cycle
    c text   --Replace the selected lines with text, which has each embedded newline 
    preceded by a backslash.

    s/regexp/replacement/g      -- search-and-replace   -- global
    --Attempt to match regexp against the  pattern  space
    g   -- Copy/append hold space to pattern space.

    The replacement may contain the special character & to  refer to that 
    portion of the pattern space which matched, and the special 
    escapes \1 through \9 to refer to the 
    corresponding matching sub-expressions in the regexp y/source/dest/
    Transliterate  the  characters  in the pattern space which appear 
    in source to the corresponding  character in dest.

stat -- display file or file system status
    stat [OPTION]... FILE...
    -f  --file-system
    -c  --format=FORMAT
    %a     access rights in octal
    %B     the size in bytes of each block reported by %b
    %m     mount point
    %x     time of last access, human-readable

systemctl    -- Control the systemd system and service manager
    -t  --type=
    -a  --all
    -l  --full  :=--list
    Unit Commands
    list-units [PATTERN...]
    start PATTERN...
    stop PATTERN...
    reload PATTERN...
    restart PATTERN...
    try-restart PATTERN...
    kill PATTERN...
    is-active PATTERN...
    status [PATTERN...|PID...]]
    show [PATTERN...|JOB...]
    Unit File Commands
    list-unit-files [PATTERN...]
    enable NAME...
    disable NAME...
    reenable NAME...
    preset NAME...
    preset-all
    is-enabled NAME...
    Checks whether any of the specified unit files are
    enabled (as with enable).

su  -- run a command with substitute user and group ID
    -  -l  --login
    -c command  --command=command
    
sudo, sudoedit   -- execute a command as another user
    -l  --list
    -K  --remove-timestamp
    -u  user  --user=user
    -v  --validate

sort    -- sort lines of text files
    -h  --human-numeric-sort
    -f  --ignore-case
    -k  --key=KEYDEF
    -n  --numeric-sort
    -t  --field-separator=SEP
    -r  --reverse
    -u  --unique
    
ss  -- another utility to investigate sockets
    ss  is  used  to  dump socket statistics. It allows showing
    information similar to netstat.
    -l  --listening
    -n  --numeric
    -t  --tcp
    -u  --udp
    -p  --processes
    -i  --info
    
	swapon, swapoff  -- enable/disable devices and files for paging and swapping
    
split   -- split a file into pieces
    -b  --bytes=SIZE
    -l  --lines=NUMBER  put NUMBER lines per output file
    -d  --numeric-suffixes[=FROM]
    -a  --suffix-length=N
    
ssh -- OpenSSH SSH client (remote login program)
    -p  --port
    -t  --Force pseudo-terminal allocation

scp -- secure copy (remote file copy program)
    -P  --port
    -p  --Preserves
    -r  --Recursively
    -l  --limit

showmount   -- show mount information for an NFS server
	-e or --exports     Show the NFS server's export list
    
sftp    -- secure file transfer program
    -o  --ssh_option
    port
    
ssh-keygen -- authentication key generation, management and conver-sion
    -t dsa  --Specifies the type of key to create.
    -P  --passphrase
    -f  --filename
    -N  --new_passphrase
    
ssh-copy-id -- use locally available keys to authorise logins on a  
    remote machine
    -i  --identity_file
    -f  --Forced mode
```

## T

```bash
touch -- change file timestamps
    touch [OPTION]... FILE...
    -a  --change only the access time
    -d  --date=STRING
    -m  --change only the modification time
    -r  --reference=FILE
    use this file's times instead of current time
    -t  --STAMP  use [[CC]YY]MMDDhhmm[.ss] instead of current time

tail -- output the last part of files
    Print the last 10 lines of each FILE to standard output.
    tail [OPTION]... [FILE]...
    -f   --follow[={name|descriptor}]  <=> tailf
    -F   --same as --follow=name --retry
    -n   -- lines

tree -- list contents of directories in a tree-like format
    -d  --List directories only
    -L  --level    Max display depth of the directory tree
    -p  --Print the file type and permissions  for  each  file
    (as per ls -l)
    -D  --Print  the  date of the last modification time
    -i  --Don't print indentation lines
    -f  --Prints the full path prefix for each file
    -F  Append a `/' for directories, a `=' for socket
    files, a `*' for executable files, a `>' for doors
    (Solaris) and a `|' for FIFO's, as per ls -F

tailf    -- follow the growth of a log file
    tailf [OPTION] file
    -n  --lines=N, -N

tar - manual page for tar 1.26
    -c  --create
    -z  --gzip
    -f  --file=ARCHIVE
    -t  --list
    -x  --extract, --get
    -v  --verbose
    -C  --directory=DIR
    -j  --bzip2
    -r  --append    append files to the end of an archive
    -X  --exclude-from=FILE
    -T  --files-from=FILE   get names to extract or create from FILE
    -N  --newer=DATE-OR-FILE
    -h  --dereference   follow  symlinks; archive and  dump the files they point to
    -p  --preserve-permissions, --same-permissions
    --exclude=PATTERN   exclude files, given as a PATTERN
    
type	-- Display information about command type

tr   - translate or delete characters
    -d  --delete
    -s  --squeeze-repeats
    
tac -- concatenate and print files in reverse

top -- display Linux processes

tee -- read from standard input and write to standard output and files
-a  --append    append to the given FILEs, do not overwrite
    
tbl -- format tables for troff

time    -- Time the execution of a script

traceroute  -- print the route packets trace to network host
	-n  Do not try to map IP addresses to host names when displaying them.

tcpdump -- dump traffic on a network
    -c  --count
    -i  --interface
    -S  --absolute-tcp-sequence-numbers
    -t  --Do not print a timestamp on each dump line.
    -n  --Do not convert host addresses to names.  This can be used  to
    avoid DNS lookups.
```

## U

```bash
unalias  -- Remove each name from the list of defined aliases

uname  -- print system information
    uname [OPTION]...
    -a  --all
    -s  --kernel-name
    -r  --kernel-release
    -m  --machine   print the machine hardware name
    -n  --nodename
    -o  --operating-system

useradd -- create a new user or update default new user
    information
    useradd [options] LOGIN
    -c  --comment COMMENT
    -d  --home-dir HOME_DIR
    -e  --expiredate EXPIRE_DATE
    -D  --defaults
    -f  --inactive INACTIVE
    -g  --gid GROUP
    -G  --groups GROUP1[,GROUP2,...[,GROUPN]]]
    -m  --create-home
    -M  --no-create-home
    -s  --shell SHELL
    -u  --uid UID

userdel -- delete a user account and related files
    userdel [options] LOGIN
    -f  --force
    -r  --remove

usermod - modify a user account
    -g  --gid GROUP
    -G  --groups
    -L  --lock
    -U  --unlock

users    -- print the user names of users currently logged in
    to the current host
    
uniq    -- report or omit repeated lines
    -c  --count
    -d  --repeated
    only print duplicate lines, one for each group
    -u  --unique
    
umask   --Display or set file mode mask.
    -S	makes the output symbolic; otherwise an octal number is output
    
umount  -- unmount file systems
    -f  --force
    -l   --lazy
```

## V

```bash
visudo  -- edit the sudoers file
	-c  --check     Enable check-only mode

vmstat  -- Report virtual memory statistics

vimdiff  -  edit two, three or four versions of a file with Vim and show differences
```

## W

```bash
which -- shows the full path of (shell) commands
	which [options] [--] programname [...]

whereis -- locate the binary, source, and manual page files for a command
    -b     Search only for binaries

wc -- print newline, word, and byte counts for each file
    wc [OPTION]... [FILE]...
    -c  --bytes     print the byte counts
    -m  --chars     print the character counts
    -l  --lines     print the newline counts
    -L  --max-line-length   print the length of the longest line
    -w  --words     print the word counts
    
w   -- Show who is logged on and what they are doing
    
who  -- show who is logged on

wget    -- The non-interactive network downloader
    --limit-rate=amount
    -i file --input-file=file
    -q  --quiet Turn off Wget's output
    -t number   --tries=number
    -T seconds  --timeout=seconds
    --spider    Wget will behave as a Web spider
    -F  --force-html
        
whatis  -- Display One-line Manual Page Descriptions

watch   -- execute a program periodically, showing output fullscreen
	-n  --interval seconds
```

## X

```bash
xargs -- build and execute command lines from standard input
    -n  -- max-args=max-args Use at most max-args arguments per command line.
    -i [replace-str]  --replace[=replace-str]
    -0, --null  Input items are terminated by a null  character  in-stead of by 
    whitespace, and the quotes and backslash are not special

xfs_info - expand an XFS filesystem
```

## Y

```bash
yum -- Yellowdog Updater Modified
    yum [options] [command] [package ...]
    yum  is  an interactive, rpm based, package manager.

    -y   --assume yes

    install package1 [package2] [...]
    update [package1] [package2] [...]
    update-to [package1] [package2] [...]
    check-update
    upgrade [package1] [package2] [...]
    remove | erase package1 [package2] [...]
    list [...]
    info [...]

    group  install
    group  update
    group  list
    group remove
```

## Z

```bash
zless   -- file perusal filter for crt viewing of compressed text

systemctl
localectl
timedatectl
hostnamectl
journalctl - Query the systemd journal

usage () {
    cat <<- EOF 
        Usage: sh $0 
    EOF
}

if (($# != 3)); then 
    usage 
    exit 1 
fi

for ((i = 0; i <= 20; ++i )); do

Error: This command has to be run under the root user.
```

