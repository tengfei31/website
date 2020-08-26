###
 # @Author: wtf
 # @Date: 2020-08-26 14:54:02
 # @LastEditors: wtf
 # @LastEditTime: 2020-08-26 19:27:48
 # @Description: plase write Description
### 
#!/bin/bash
#运行MySQL的docker
echo `docker run --name mysql -p 4406:3306 -e MYSQL_ROOT_PASSWORD=tengfei31 -v ~/Documents/www/go/data/docker-mysql:/var/lib/mysql -d mysql`


