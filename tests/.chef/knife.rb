# See http://docs.opscode.com/config_rb_knife.html for more information on knife configuration options

current_dir = File.dirname(__FILE__)

log_level                :info
log_location             STDOUT
client_key               "#{current_dir}/admin.pem"
validation_client_name   "testorg-validator"
validation_key           "#{current_dir}/validation.pem"
chef_server_url          "https://localhost/organizations/testorg"
node_name                'admin'
cache_type               'BasicFile'
cache_options( :path => "#{ENV['HOME']}/.chef/checksums" )
cookbook_path            ["#{current_dir}/../cookbooks"]
