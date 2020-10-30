# Mysql 集群部署

## 目标

- 支持主从复制的 Mysql 集群
- 一个 Master
- 多个 Slave
- 支持 Slave 水平扩展
- Master 负责写操作
- Slave 负责读操作

## 容器化

- Master 节点和 Slave 节点需要有不同的配置文件（即：不同的 my.cnf）
- Master 节点和 Slave 节点需要能够传输备份信息文件
- 在 Slave 节点第一次启动之前，需要执行一些初始化 SQL 操作

## ConfigMap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql
  labels:
    app: mysql
data:
  master.cnf: |
    # 主节点MySQL的配置文件
    [mysqld]
    log-bin
  slave.cnf: |
    # 从节点MySQL的配置文件
    [mysqld]
    super-read-only
```

## Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  ports:
    - name: mysql
      port: 3306
  clusterIP: None
  selector:
    app: mysql
---
apiVersion: v1
kind: Service
metadata:
  name: mysql-read
  labels:
    app: mysql
spec:
  ports:
    - name: mysql
      port: 3306
  selector:
    app: mysql
```

## 测试连接

```bash

kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never --mysql -h mysql-0.mysql <<EOF
CREATE DATABASE test;
CREATE TABLE test.messages (message VARCHAR(250));
INSERT INTO test.messages VALUES ('hello');
EOF


kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never -- mysql -h mysql-0.mysql <
```

## 扩展

```bash
kubectl scale statefulset mysql --replicas=4
```

