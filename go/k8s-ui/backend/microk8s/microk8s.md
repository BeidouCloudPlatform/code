# MicroK8s is great for offline development, prototyping, and testing.(https://microk8s.io/)

## multipass(https://multipass.run/)
### Install microk8s on Mac
```shell script
brew cask install multipass
multipass launch microk8s-vm # launch an Ubuntu vm instance
multipass shell microk8s-vm # shell into vm
multipass exec microk8s-vm -- lsb_release --description
multipass ls # list vm instances
multipass info/stop microk8s-vm

# install the MicroK8s snap and configure the network
sudo snap install microk8s --classic --channel=1.17/stable
sudo iptables -P FORWARD ACCEPT
```

## Installing snap on Raspbian
```shell script
sudo vi /boot/firmware/nobtcmd.txt
cgroup_enable=memory cgroup_memory=1

# https://snapcraft.io/docs/installing-snap-on-raspbian
sudo apt update
sudo apt install snapd
sudo snap install multipass --classic --beta

```

## iptables
iptables 完成封包过滤、封包重定向和网络地址转换（NAT）等功能。

