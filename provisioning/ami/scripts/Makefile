# provisioning (run only in building ami)
#
# パッケージのインストール等はここでやる
#
.PHONY: all cloud-config yum td-agent-config app

all: cloud-config yum td-agent-config app

cloud-config: cloud.cfg.d/99_defaults.cfg
	cp cloud.cfg.d/99_defaults.cfg /etc/cloud/cloud.cfg.d/
	chmod 0644 /etc/cloud/cloud.cfg.d/99_defaults.cfg
	chown -R root:root /etc/cloud/cloud.cfg.d/99_defaults.cfg

td-agent-config: td-agent/td-agent.cfg
    cp -f td-agent/td-agent.cfg /etc/td-agent/td-agent.cfg

yum:
	yum -y update
	yum -y install https://dev.mysql.com/get/mysql-community-release-el7-5.noarch.rpm
	yum -y --enablerepo=mysql56-community install mysql-community-client
	yum -y install java-1.8.0-openjdk
	yum -y install https://s3.amazonaws.com/amazoncloudwatch-agent/amazon_linux/amd64/latest/amazon-cloudwatch-agent.rpm
	yum -y install http://packages.treasuredata.com.s3.amazonaws.com/3/amazon/1/2017.09/x86_64/td-agent-3.1.1-0.el2017_09.x86_64.rpm
	yum -y install mysql-community-devel
    yum -y install gcc
    td-agent-gem install fluent-plugin-mysql

/root/hakaru:
	mkdir -p /root/hakaru

/root/hakaru/team_name.txt: deploy/team_name.txt /root/hakaru
	cp deploy/team_name.txt /root/hakaru/team_name.txt

/root/hakaru/Makefile: deploy/Makefile /root/hakaru/team_name.txt /root/hakaru
	cp deploy/Makefile /root/hakaru/Makefile

app: /root/hakaru/Makefile
    /etc/init.d/td-agent start
