#!/bin/bash

apt-get update
apt-get install -y curl vim wget git

CHEFDK_PKG=chefdk_0.3.5-1_amd64.deb
CHEFDK_URI=https://opscode-omnibus-packages.s3.amazonaws.com/ubuntu/12.04/x86_64
CHEF_SERVER_PACKAGE=chef-server-core_12.0.1-1_amd64.deb

## Caching this package b/c its huge
if ! [ -a /home/vagrant/.cache/${CHEFDK_PKG} ]; then
  pushd /home/vagrant/.cache &> /dev/null
  wget ${CHEFDK_URI}/${CHEFDK_PKG}
  popd &> /dev/null
fi
dpkg -i /home/vagrant/.cache/${CHEFDK_PKG}

mkdir -p /home/vagrant/.berkshelf
cat <<'EOF' > /home/vagrant/.berkshelf/config.json
{
  "ssl": {
    "verify": false
  }
}
EOF
chown -R vagrant:vagrant /home/vagrant/.berkshelf

## Caching this package b/c its huge
if ! [ -a /home/vagrant/.cache/${CHEF_SERVER_PACKAGE} ]; then
  pushd /home/vagrant/.cache &> /dev/null
  wget https://web-dl.packagecloud.io/chef/stable/packages/ubuntu/trusty/${CHEF_SERVER_PACKAGE}
  popd &> /dev/null
fi
dpkg -i /home/vagrant/.cache/${CHEF_SERVER_PACKAGE}

chef-server-ctl reconfigure

## Take a nap chef server doesn't become ready right away
sleep 20

## Make sure we get the newly generated credentials
rm -f /home/vagrant/.chef/admin.pem
rm -f /home/vagrant/.chef/validation.pem

chef-server-ctl user-create admin Fins Test admin@whocares.com password --filename /home/vagrant/.chef/admin.pem
chef-server-ctl org-create testorg Fins Test Org --association_user admin --filename /home/vagrant/.chef/validation.pem

sleep 2
pushd /home/vagrant/.chef &> /dev/null
su - vagrant -c 'cd .chef && berks install && berks upload'
su - vagrant -c 'cd .chef && knife environment from file environments/staging.json'
su - vagrant -c 'cd .chef && knife environment from file environments/production.json'
popd &> /dev/null
exit
