fins
====
Inspect your Chef server

###Install
Clone repo Run 'make && sudo make install'

###Usage

####Outdated
fins outdated shows all cookbooks that are outdated in a Chef server
based on what is currently marked as the latest in the supermarket.

A future release will allow for the specification of git repos to
check for cookbooks.

```shell
# fins outdated
bin/fins -c ./fins.json outdated
consul:
  Chef Server: 0.9.0
  Supermarket: 0.9.1
aws:
  Chef Server: 2.5.0
  Supermarket: 2.7.0
yum:
  Chef Server: 3.5.4
  Supermarket: 3.6.0
ntp:
  Chef Server: 1.7.0
  Supermarket: 1.8.2
yum-repoforge:
  Chef Server: 0.5.0
  Supermarket: 0.5.1
runit:
  Chef Server: 1.5.18
  Supermarket: 1.6.0
database:
  Chef Server: 4.0.3
  Supermarket: 4.0.6
```

####Diff
Shows the difference between what is on the chef server and what
an environment allows through.  So if an environment named staging
has been pinned to a version older then it will show up in the diff.

If you pass two chef environments it will show you the difference between
them.

Note: This is will only return cookbooks that are different.

```shell
# Difference between the chef server and staging
bin/fins -c ./fins.json diff staging
git:
  Environment Constraint: = 4.1.0
  Latest Cookbook Version: 4.2.1

# Difference between staging and production
bin/fins -c ./fins.json diff staging production
iptables:
  staging Constraint: no version constraint
  production Constraint: = 0.14.1
  Latest Cookbook Version: 1.0.0
monit:
  staging Constraint: no version constraint
  production Constraint: = 0.7.5
  Latest Cookbook Version: 0.7.5
pg-multi:
  staging Constraint: no version constraint
  production Constraint: = 0.1.2
  Latest Cookbook Version: 0.1.2
python:
  staging Constraint: no version constraint
  production Constraint: = 1.4.6
  Latest Cookbook Version: 1.4.7
locale:
  staging Constraint: = 1.0.2
  production Constraint: no version constraint
  Latest Cookbook Version: 1.0.2
```
