# 1. 生成ca key
openssl genrsa -out ca.key 2048
# 2. 生成ca csr证书请求文件
openssl req -new -key ca.key -out ca.csr
# 3. 生成自签名ca证书
openssl x509 -req -in ca.csr -out ca.crt -signkey ca.key -days 3650

# 4. 生成docker server key
openssl genrsa -out docker-server.key 2048
# 5. 生成docker server csr证书请求文件
## 复制k8s-master:/etc/ssl/openssl.cnf到本地目录
multipass mount /Users/lx1036/Code/lx1036/code/go/k8s-ui/backend/microk8s/mounts/k8s-master k8s-master
multipass copy-files k8s-master:/etc/ssl/openssl.cnf /Users/lx1036/Code/lx1036/code/go/k8s-ui/backend/microk8s/docker-server-openssl.cnf
openssl req -new -key docker-server.key -out docker-server.csr -config docker-server-openssl.cnf

# 6. 查看docker server csr证书请求文件
openssl req -text -noout -in docker-server.csr
# 7. 添加index和series文件，创建series, 生成docker server证书并用ca证书签名
mkdir -p ./demoCA/docker-server
touch ./demoCA/docker-server/index.txt
touch ./demoCA/docker-server/serial
echo 00 > ./demoCA/docker-server/serial
mkdir ./demoCA/docker-server/newcerts
openssl ca -in docker-server.csr -out docker-server.crt -cert ca.crt -keyfile ca.key -extensions v3_req -days 3650 -config docker-server-openssl.cnf

# 8. 制作docker client证书
openssl  genrsa -out docker-client.key  2048
openssl  req -new -key docker-client.key -out docker-client.csr -config docker-client-openssl.cnf
mkdir -p ./demoCA/docker-client
touch ./demoCA/docker-client/index.txt
touch ./demoCA/docker-client/serial
echo 01 > ./demoCA/docker-client/serial
mkdir ./demoCA/docker-client/newcerts
openssl ca -in docker-client.csr -out docker-client.crt -cert ca.crt -keyfile ca.key -extensions v3_req -days 3650 -config docker-client-openssl.cnf
