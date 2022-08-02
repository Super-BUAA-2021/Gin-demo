#!/bin/bash
# 设置为 release 生产模式
# export GIN_MODE=release
# 切换到路径下，这样才能够使用和开发时候一样的相对路径
# 启动 build 后的可执行文件
./Gin-demo
# 启动 gin 开发服务器

ps -ef | grep Gin-demo