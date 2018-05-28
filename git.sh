#!/bin/bash
#describe:用于git 下载 更新 上传ssh
#@2018-05-28
#@author
db_Server="git@github.com:huangchengwu/game_server.git"
case $1 in
download)
git $db_Server game_server
;;
update)
git fetch
git add .
git common -m "update $2"
git remote add game_admin $db_Server
git push -u game_admin master

;;
ssh_upload)
read -p "请输入邮箱: " email

if [ ! -z $email ]
then
	echo "正在生成ssh"

/usr/bin/expect << EOF
	set timeout -1
	spawn ssh-keygen -t rsa -C "$email"
	expect  {
	"Enter file in which to save the key (/root/.ssh/id_rsa): " { send "\r" }
	}
	expect "Overwrite (y/n)? "
	send "y\r"
	expect "Enter passphrase (empty for no passphrase):"
	send "\r"
	expect "Enter same passphrase again: "
	send "\r"
	expect eof
EOF
ssh -T git@github.com
fi
;;
*)
echo "download|update|ssh_upload"

;;
esac

