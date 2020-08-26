###
 # @Author: wtf
 # @Date: 2020-08-26 14:56:41
 # @LastEditors: wtf
 # @LastEditTime: 2020-08-26 20:00:32
 # @Description: plase write Description
### 
#!/bin/bash

#构建docker环境
container_name="website-docker-scratch"
count=`docker images | grep "${container_name}" | wc -l`
echo $count
if [ $count -le 0 ]; then
    build=`docker build -t $container_name .`
    echo $build
else
    echo "已有镜像,无需执行docker build指令"
fi
#运行构建的docker
run=`docker run -p 8000:8088 --link mysql:mysql $container_name`
echo "${website-docker-scratch}容器ID=${run}"
