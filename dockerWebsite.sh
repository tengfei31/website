###
 # @Author: wtf
 # @Date: 2020-08-26 14:56:41
 # @LastEditors: wtf
 # @LastEditTime: 2020-08-26 14:58:12
 # @Description: plase write Description
### 
#!/bin/bash

#构建docker环境
`docker build -t website-docker .`
#运行构建的docker
`docker run -p 8000:8088 --link mysql:mysql website-docker`
