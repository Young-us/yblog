# yblog
一个用于记录和分享编程笔记的博客

## 项目概要

1.用户注册登录认证，支持多用户记录分享

2.在线编辑文章，实时文章分页预览，支持文章分类

3.HTML静态资源加载，页面加载迅速

4.kubernetes服务部署

## 目录架构

app：服务响应接口

bootstrap：项目启动初始化入口

config：环境变量配置

pkg：工具函数

public：本地资源

resources：本地资源

routes：HTTP路由

blogKube：kubernetes服务部署

## 项目启动

一.本地部署：
1.启动MySQL数据库，创建所需database。
2.go run main.go

二.kubernetes部署
1.docker build -t your-iamge-name
2.docker push your-iamge-name
3.cd blogKube
4.Run.bat
.
