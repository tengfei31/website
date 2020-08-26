###
 # @Author: wtf
 # @Date: 2020-08-26 14:54:02
 # @LastEditors: wtf
 # @LastEditTime: 2020-08-26 14:56:00
 # @Description: plase write Description
### 
#!/bin/bash
#运行MySQL的docker
`docker run --name mysql -p 3306:4406 -e MYSQL_ROOT_PASSWORD=tengfei31 -d mysql`


