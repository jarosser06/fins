fins
====
Inspect your Chef server for outdated versions

###Install

Clone repo Run 'make && sudo make install'

###Usage
fins outdated shows all cookbooks that are outdated in a Chef server

```shell
# fins outdated
[jim@wjarvis] ~/Go/src/github.com/jarosser06/fins/tests (master)$ ../bin/fins -c ./fins.json outdated
Name: consul Chef Server: 0.6.0 Supermarket: 0.9.1
Name: yum-repoforge Chef Server: 0.4.0 Supermarket: 0.5.1
Name: ntp Chef Server: 1.7.0 Supermarket: 1.8.2
Name: varnish Chef Server: 1.0.2 Supermarket: 2.1.1
Name: mysql-multi Chef Server: 1.4.1 Supermarket: 2.1.3
Name: database Chef Server: 2.3.1 Supermarket: 4.0.6
Name: ulimit Chef Server: 0.3.2 Supermarket: 0.3.3
Name: newrelic_meetme_plugin Chef Server: 0.1.1 Supermarket: 0.2.1
Name: logrotate Chef Server: 1.7.0 Supermarket: 1.9.1
Name: redisio Chef Server: 2.2.4 Supermarket: 2.3.0
Name: yum-epel Chef Server: 0.5.3 Supermarket: 0.6.0
Name: omnibus_updater Chef Server: 1.0.2 Supermarket: 1.0.4
Name: yum-mysql-community Chef Server: 0.1.11 Supermarket: 0.1.17
Name: application Chef Server: 4.1.4 Supermarket: 4.1.6
Name: elkstack Chef Server: 3.2.6 Supermarket: 4.2.3
Name: user Chef Server: 0.3.0 Supermarket: 0.4.2
Name: yum Chef Server: 3.5.2 Supermarket: 3.6.0
Name: postgresql Chef Server: 3.4.14 Supermarket: 3.4.18
Name: line Chef Server: 0.5.1 Supermarket: 0.6.1
Name: platformstack Chef Server: 1.5.3 Supermarket: 3.1.4
Name: aws Chef Server: 2.5.0 Supermarket: 2.7.0
Name: stack_commons Chef Server: 0.0.38 Supermarket: 0.0.50
Name: apt Chef Server: 2.6.1 Supermarket: 2.7.0
Name: logstash Chef Server: 0.10.2 Supermarket: 0.11.4
Name: rackspacecloud Chef Server: 0.0.6 Supermarket: 0.1.1
Name: nginx Chef Server: 2.7.4 Supermarket: 2.7.6
Name: ruby Chef Server: 0.9.2 Supermarket: 0.9.3
Name: chef-sugar Chef Server: 2.5.0 Supermarket: 3.1.0
Name: rsyslog Chef Server: 1.13.0 Supermarket: 1.15.0
Name: iptables Chef Server: 0.14.1 Supermarket: 1.0.0
Name: git Chef Server: 4.1.0 Supermarket: 4.2.2
Name: runit Chef Server: 1.5.12 Supermarket: 1.6.0
Name: java Chef Server: 1.29.0 Supermarket: 1.31.0
Name: mysql Chef Server: 5.6.1 Supermarket: 6.0.21
Name: rabbitmq Chef Server: 3.7.0 Supermarket: 4.0.0
Name: build-essential Chef Server: 2.1.3 Supermarket: 2.2.3
Name: newrelic Chef Server: 2.5.2 Supermarket: 2.11.2
Name: rackspace_cloudbackup Chef Server: 1.0.3 Supermarket: 1.0.4
Name: libarchive Chef Server: 0.4.1 Supermarket: 0.5.0
Name: apache2 Chef Server: 3.0.0 Supermarket: 3.0.1
Name: users Chef Server: 1.7.0 Supermarket: 1.8.2
Name: elasticsearch Chef Server: 0.3.10 Supermarket: 0.3.13
Name: chef-client Chef Server: 4.0.0 Supermarket: 4.3.0
Name: erlang Chef Server: 1.5.6 Supermarket: 1.5.8
Name: golang Chef Server: 1.3.0 Supermarket: 1.5.0
Name: openssl Chef Server: 2.0.2 Supermarket: 4.0.0
Name: windows Chef Server: 1.36.1 Supermarket: 1.36.6
Name: xml Chef Server: 1.2.9 Supermarket: 1.2.13
Name: mysql-chef_gem Chef Server: 0.0.5 Supermarket: 1.0.0
Name: iis Chef Server: 2.1.6 Supermarket: 4.1.0
Name: rackspace_gluster Chef Server: 0.3.0 Supermarket: 1.0.0
Name: cron Chef Server: 1.4.3 Supermarket: 1.6.1
```
