# Deploy Fresh Magma (AGW + Orc8r + NMS)


**Prerequisites -**
1. we need docker,docker compose,virtualbox and vagrant for creating magma VM 
     ansible,python3.5,go, pyenv(for global python environment).
2. after vagrant installation, install vagrant-disksize (0.1.3) and vagrant-vbguest (0.24.0) plugins.

**Steps -**

1. HOST [magma]$ `export MAGMA_ROOT= path of magma folder`

**a) Setup AGW VM**

1. HOST [magma]$ `git clone https://github.com/magma/magma.git` from master branch.
2. HOST [magma]$ `cd magma/lte/gateway`
3. changes to do in magma/lte/gateway/Vagrantfile (not necessary, only when more than 1 AGW VM is there on same system) => 
     vb.name = "magma-dev-demo"  --> change the VM name here
     magma.vm.network "private_network", ip: "192.168.60.142", nic_type: "82540EM" --> replace 142 with a number in range (2-255) which shouldn’t be in use elsewhere.
4. changes to do in magma/lte/gateway/deploy/hosts =>
    under [focal_dev] --> replace the IP with the same IP that was assigned recently in VagrantFile.
5. HOST [magma]$ `cd magma/lte/gateway`
6. HOST [magma]$ `vagrant up magma` (if giving problems then use sudo with vagrant)
7. if vagrant giving any plugin errors, then uninstall and install the plugins
8. if vagrant gives some other problems then destroy it (vagrant destroy magma) and create it again

**b) Build Orc8r**
     
0. HOST [magma]$ `export MAGMA_ROOT= path of magma folder`
1. HOST [magma]$ `cd magma/orc8r/cloud/docker`
2. HOST [magma]$ `sudo PWD=$PWD ./build.py –-all`

**c) Build AGW**
     
1. HOST [magma]$ `cd magma/lte/gateway`
2. HOST [magma]$ `vagrant ssh magma`, this will login to magma VM (if giving problems then use sudo with vagrant)
3. AGW-VM$ `cd magma/lte/gateway`
4. AGW-VM$ `make run`

**d) Start Orc8r**

1. HOST [magma]$ `cd magma/orc8r/cloud/docker` in HOST (your own login and not vagrant) 
2. HOST [magma]$ `sudo PWD=$PWD ./run.py --metrics`
3. HOST [magma]$ `cd ../../../.cache/test_certs`
4. HOST [magma]$ `cp admin_operator.pfx ../../../`
5. add admin_operator.pfx to your browser certificate
6. open https://localhost:9443/swagger/v1/ui in your browser after adding admin_operator.pfx as a certificate

**e) If only want to use Orc8r**

1. exit from vagrant MAGMA-VM
2. HOST [magma]$ `cd magma/lte/gateway`
3. HOST [magma]$ `vagrant halt magma` (if giving problems then use sudo with vagrant)

 **f) Connect AGW to Orc8r**
     
1. HOST [magma]$ `cd lte/gateway`
2. HOST [magma/lte/gateway]$ `fab -f dev_tools.py register_vm`
3. HOST [magma/lte/gateway]$ `vagrant ssh magma`

4. AGW-VM$ `sudo service magma@* stop`
5. AGW-VM$ `sudo service magma@magmad restart`
6. AGW-VM$ `sudo tail -f /var/log/syslog`

     After a minute or 2 you should see these messages:
     
     Sep 27 22:57:35 magma-dev magmad[6226]: [2018-09-27 22:57:35,550 INFO root] Checkin Successful!
     Sep 27 22:57:55 magma-dev magmad[6226]: [2018-09-27 22:57:55,684 INFO root] Processing config update g1
     Sep 27 22:57:55 magma-dev control_proxy[6418]: 2018-09-27T22:57:55.683Z [127.0.0.1 -> streamer-controller.magma.test,8443] "POST /magma.Streamer/GetUpdates        HTTP/2" 200 7bytes 0.009s

**f) Setup NMS UI**
1. HOST [magma]$ cd `nms/packages/magmalte`
2. HOST [magma/nms/packages/magmalte] $ `docker-compose build magmalte`
3. HOST [magma/nms/packages/magmalte] $ `docker-compose up -d`
4. HOST [magma/nms/packages/magmalte] $ `./scripts/dev_setup.sh` 
