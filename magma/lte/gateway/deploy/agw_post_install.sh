#!/bin/bash
# Setting up env variable, user and project path
ERROR=""
INFO=""
SUCCESS_MESSAGE="ok"

addError() {
    ERROR="$ERROR\n$1  to fix it: $2"
}

addInfo() {
    INFO="$INFO $1 \n"
}

if ! grep -q 'Debian' /etc/issue; then
  addError "Debian is not installed" "Restart installation following agw_install.sh, agw has to run on Debian"
fi

if ! grep -q "$MAGMA_USER ALL=(ALL) NOPASSWD:ALL" /etc/sudoers; then
    addError "Debian is not installed" "Restart installation following agw_install.sh, magma has to be sudoer"
fi

KVERS=$(uname -r)
if [ "$KVERS" != "4.9.0-9-amd64" ]; then
    addError "Kernel version is not 4.9.0-9-amd64" "Restart installation following agw_install.sh, KVERS has to be 4.9.0-9-amd64"
fi

interfaces=("eth1" "eth0" "gtp_br0")
for interface in "${interfaces[@]}"; do
    OPERSTATE_LOCATION="/sys/class/net/$interface/operstate"
    if test -f "$OPERSTATE_LOCATION"; then
        OPERSTATE=$(cat "$OPERSTATE_LOCATION")
        if [[ $OPERSTATE == 'down'  ]]; then
            addError "$interface is not configured" "Try to ifup $interface"
        fi
    else
        addError "$interface is not configured" "Check that /etc/network/interfaces.d/$interface has been set up"
    fi
done

PING_RESULT=$(ping -c 1 -I eth0 8.8.8.8 > /dev/null 2>&1 && echo "$SUCCESS_MESSAGE")
if [ "$PING_RESULT" != "$SUCCESS_MESSAGE" ]; then
    addError "eth0 is connected to the internet" "Make sure the hardware has been properly plugged in (eth0 to internet)"
fi

allServices=("control_proxy" "directoryd" "dnsd" "enodebd" "magmad" "mme" "mobilityd" "pipelined" "policydb" "redis" "sessiond" "state" "subscriberdb")
for service in "${allServices[@]}"; do
    if ! systemctl is-active --quiet "magma@$service"; then
        addError "$service is not running" "Please check our faq"
    fi
done

nonMagmadServices=("sctpd")
for service in "${nonMagmadServices[@]}"; do
    if ! systemctl is-active --quiet "$service"; then
        addError "$service is not running" "Please check our faq"
    fi
done

packages=("magma" "magma-cpp-redis" "magma-libfluid" "oai-gtp" "libopenvswitch" "openvswitch-datapath-dkms" "openvswitch-datapath-source" "openvswitch-common" "openvswitch-switch")
for package in "${packages[@]}"; do
    PACKAGE_INSTALLED=$(dpkg-query -W -f='${Status}' $package  > /dev/null 2>&1 && echo "$SUCCESS_MESSAGE")
    if [ "$PACKAGE_INSTALLED" != "$SUCCESS_MESSAGE" ]; then
        addError "$package hasn't been installed" "Rerun the agw_install.sh"
    fi
done

if [ -z "$ERROR" ]; then
    echo "Installation went smoothly, please let us know what went wrong/good on github"
else
    echo "There was a few errors during installation"
    printf "%s" "$ERROR"
fi

apt-get update > /dev/null
addInfo "$(apt list -qq --upgradable 2> /dev/null)"

if [ -n "$INFO" ]; then
    echo "INFO:"
    printf "%s" "$INFO"
fi

echo "- Check for Root Certificate"
CA=/var/opt/magma/tmp/certs/rootCA.pem
if [ -d "/var/opt/magma/tmp/certs/" ]; then
    if [ -f "$CA" ]; then
        echo "$CA exists"
    fi
else
    echo "Verify Root CA in /var/opt/magma/tmp/certs/"
    echo "Access Gateway configurations failed"
fi

echo "- Check for Control Proxy"
CP=/var/opt/magma/configs/control_proxy.yml
if [ -d "/var/opt/magma/configs/" ]; then
    if [ -f "$CP" ]; then
        echo "$CP exists"
        echo "- Check control proxy content"
        cp_content=("cloud_address" "cloud_port" "bootstrap_address" "bootstrap_port" "rootca_cert" "fluentd_address" "fluentd_port")
        for content in "${cp_content[@]}"; do
            if ! grep -q $content $CP; then
                echo "Missing $content in control proxy"
            fi
        done
    else
        echo "Control proxy file missing. Check magma installation docs"
    fi
else
    echo "Check Control Proxy Configs in /var/opt/magma/configs/"
    echo "Access Gateway configurations failed"
fi

echo "- Verifying Cloud check-in"
CLOUD=$(journalctl -n20 -u magma@magmad | grep -e 'Checkin Successful' -e 'Got heartBeat from cloud')
if [ "$CLOUD" ]; then
    echo "Cloud Check Success"
else
    echo "Cloud Check Failed"
    echo "Check control proxy content in $CP"
fi
