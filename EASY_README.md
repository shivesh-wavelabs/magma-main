#Already built Magma Files

**Note -** if any errors, refer README.md

**Prerequisites-**
1. we need docker,docker compose,virtualbox and vagrant for creating magma VM 
     ansible,python3.5
2. after vagrant installation, install vagrant-disksize (0.1.3) and vagrant-vbguest (0.24.0) plugins.

**a) Set AGW VM**

1. git clone https://github.com/ShiveshOjha/magma-main.git from main branch.
2. cd magma/lte/gateway
3. vagrant up magma (if giving problems then use sudo with vagrant)
4. if vagrant giving any plugin errors, then uninstall and install the plugins
5. if vagrant gives some other problems then destroy it (vagrant destroy magma) and create it again

**b) Build Orc8r**

1. cd magma/orc8r/cloud/docker
2. sudo PWD=$PWD ./build.py –all

**c) Build AGW**

1. cd magma/lte/gateway
2. vagrant ssh magma, this will login to magma VM (if giving problems then use sudo with vagrant)
3. cd magma/lte/gateway in MAGMA-VM
4. make run
**
d) Start Orc8r**

1. cd magma/orc8r/cloud/docker in HOST (your own login and not vagrant) 
2. sudo PWD=$PWD ./run.py --metrics
3. cd ../../.././cache/test_certs
4. cp admin_operator.pfx ../../../
5. add admin_operator.pfx to your browser certificate
6. open https://localhost:9443/swagger/v1/ui in your browser after adding admin_operator.pfx as a certificate

**e) If only want to use Orc8r**

1. exit from vagrant MAGMA-VM
2. cd magma/lte/gateway
3. vagrant halt magma (if giving problems then use sudo with vagrant)
