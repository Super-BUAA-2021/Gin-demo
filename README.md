# Gin-demo

Golang Web Framework


# 使用说明
在根目录下创建`demo.yml` 文件，按照`configs/example.yml` 的格式填下你的信息方可运行

# 移植说明

1. 全局搜索并替换`github.com/Super-BUAA-2021/Gin-demo` 为你想要的包名。 本例为了符合大多项目对于module的设置，以及方便其他人的使用。将module设置为此。也在使用`service`包都需要`import "github.com/Super-BUAA-2021/Gin-demo/service"` 。
2. 全局搜索demo，并替换你想要的项目名。包括但不限于：日志文件名，配置文件名
3. `main.go/Swagger`注释更改为个人以及项目信息。
4. 在配置文件中更改为你自己的数据库以及jwt-secret等信息。



相信完成上述四步，运行main便可将此移植为你的项目了
