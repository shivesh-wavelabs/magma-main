# Deploy Fresh Magma and Orc8r

Setting up Orc8r locally on Ubuntu

**Prerequisites -**
1. we need docker,docker compose,virtualbox and vagrant for creating magma VM 
     ansible,python3.5
2. after vagrant installation, install vagrant-disksize (0.1.3) and vagrant-vbguest (0.24.0) plugins.

**Steps -**

**a) Set AGW VM**

1. git clone https://github.com/magma/magma.git from master branch.
2. cd magma/lte/gateway
3. changes to do in magma/lte/gateway/Vagrantfile => 
     magma.vm.hostname = "magma-dev-demo"  --> change the VM name here
     magma.vm.network "private_network", ip: "192.168.60.142", nic_type: "82540EM" --> replace 142 with a number in range (2-255) which     shouldn’t be in use elsewhere.
     vb.name = "magma-dev-demo"  --> change the VM name here
4. changes to do in magma/lte/gateway/deploy/hosts =>
    under [focal_dev] --> replace the IP with the same IP that was assigned recently in VagrantFile.
5. cd magma/lte/gateway
6. vagrant up magma (if giving problems then use sudo with vagrant)
7. if vagrant giving any plugin errors, then uninstall and install the plugins
8. if vagrant gives some other problems then destroy it (vagrant destroy magma) and create it again

**Error or incompatibility with Virtualbox -**

1. Install the latest vagrant version by downloading .deb file from its website.

2. Install Virtualbox using apt manager and follow installation instructions (mokutil and secure boot options (give password as 12345678 for easiness)) and reboot.
3. On restarting, select mokutil option at boot screen and enter the asked password (12345678)
4. Come back and select simple boot option
5. Make sure virtualization is enabled in bios.

2. If Virtualbox is giving configuration errors, then do the following -
    1. Go to this site and download Virtual Box for your OS: https://www.virtualbox.org/wiki/Linux_Downloads
       
    2. sudo apt remove virtualbox or sudo apt remove virtualbox-6.1
       
       //In Downloads Folder:
    3. sudo dpkg -i virtualbox-6.1_6.1.22-144080~Ubuntu~eoan_amd64.deb
    4. //If there are errors:
    5. sudo apt install -f
       
       //virtualbox Configure:
    6. $sudo /sbin/vboxconfig
    7. $/sbin/vboxconfig
       
       //If Configure fails:
       
       //For mdprobe error:
    8. sudo modprobe vboxdrv
    9. sudo apt-get update
    10. sudo apt-get upgrade
    11. sudo apt-get install mokutil
    12. sudo /usr/src/linux-headers-$(uname -r)/scripts/sign-file sha256 ./MOK.priv ./MOK.der $(modinfo -n vboxdrv)
       
       //Permission Denied check OS settings error:
    13. sudo mokutil --disable-validation 
    14. reboot

**b) Build Orc8r**

1. cd magma/orc8r/cloud/docker
2. sudo PWD=$PWD ./build.py –all

**c) Build AGW**

1. cd magma/lte/gateway
2. vagrant ssh magma, this will login to magma VM (if giving problems then use sudo with vagrant)
3. cd magma/lte/gateway in MAGMA-VM
4. make run

**d) Start Orc8r**

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
