
FROM ubuntu:18.04
MAINTAINER lx1036@126.com

RUN mv /etc/apt/sources.list /etc/apt/sourses.list.backup
COPY ./etc/apt/sources.list /etc/apt/sources.list

# mtr(my traceroute)=ping+traceroute https://mp.weixin.qq.com/s/iCH8VBAtlNnDO-04eHbVZQ
RUN apt-get update \
    && apt-get install -y curl iputils-ping iproute2 telnet net-tools mtr \
    && apt-get install -y vim
