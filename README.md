Mongo-replication-example
=============
Mongo replication 的範例專案，此份專案將會運行 mongo-replication 叢集，叢集包括一個primary container, 兩個secondary container。primary container 負責mongo 的讀寫，secondary container 為備份，並在primary 死掉時遞補成為新的primary。另外在用golang 對mongo-replication-cluster 做連線操作

## 1. 建立docker-compose

```
$ docker-compose up -d
```
建立一個 golang-server Container, 與三個mongo Container。

其中mongo-secret 為 mongo 的金鑰，唯有持有相同金鑰的mongo 才能夠成為同個replication-cluster。

金鑰要求base64 格式，大小在 6 ~ 1024 個字元間，官方範例：
```
openssl rand -base64 756 > <path-to-keyfile>
```
## 2. Mongo Container Entrypoint
在docker-compose.yml 中，每個mongo container 會執行 entrypoint.sh
```
command: sh /usr/entrypoint.sh
```
在entrypoint.sh 中：
```
chmod 600 /usr/mongo-secret
```
將secret 的權限嚴格，只允許user 讀寫。以原本寬鬆的權限 (666) 會因為權限太寬鬆而無法成為mongo 的 keyfile。

```
/usr/bin/mongod --keyFile /usr/mongo-secret --bind_ip_all --replSet rs0
```
啟動mongo 其中：
--keyFile: 為金鑰所在位置
--bind_ip_all 接受的ip 這裡寫接受所有ip, 實際應用需一一天上其他mongo 實例的ip 位置
--replSet: 設定的 replica id。同一個replicat set 的 replSet 要設的一樣。

## 3. 設定Mongo Replication
做初步的mongo reclica 設定，首先進入golang-mongo-1 
```
$ docker exec -it golang-mongo-1 bash
```
進到container 後，將現有的三個mongo container 連結起來成為一個replica-set-cluster
```
$ mongo --eval "rs.initiate({_id:'rs0', members: [{_id: 0, host: 'golang-mongo-1:27017'}, { _id: 1, host: 'golang-mongo-2:27017'}, { _id: 2, host: 'golang-mongo-3:27017'}]})"
```
or
```
$ echo "rs.initiate({_id:'rs0', members: [{_id: 0, host: 'golang-mongo-1:27017'}, { _id: 1, host: 'golang-mongo-2:27017'}, { _id: 2, host: 'golang-mongo-3:27017'}]})" | mongo
```
整個過程大概需要 5 ~ 10 秒。
接著，創建super user。若cluster 還沒創建完成，沒辦法創建user。

```
$ mongo admin --eval "db.createUser({user: 'root', pwd: '1234', roles: ['root']})"
```
or
```
$ echo "db.createUser({user: 'root', pwd: '1234', roles: ['root']})" | mongo admin
```
在mongo 中，當第一個使用者建立完後，單純開啟mongo shell 便沒有任何權限，所以第一個使用者的權限最低要有admin 創建使用者的權限。

接著退出container
```
$ exit
```

## 4. 透過go server 對 mongo replicaset 建立 collection。
```
$ docker exec -it golang-server sh
$ go run mongotest.go
$ exit
```
go 與 mongo replica-set cluster 連結，有root 權限並在test db新建一個collection

```
$ docker exec -it golang-mongo-1 bash
$ mongo -u root -p 1234
show collections
```
可以看到新增了collection

進到其他mongo container 

```
$ docker exewc -it golang-mongo-2 bash
$ mongo -u root -p 1234
rs.secondaryOk()
show collections
```
可以看到同樣也建立了lala collection

## 5. Reference
https://docs.mongodb.com/manual/tutorial/deploy-replica-set-with-keyfile-access-control/

https://www.youtube.com/watch?v=Q2lJH156SUQ&list=PL34sAs7_26wPvZJqUJhjyNtm7UedWR8Ps&index=6
