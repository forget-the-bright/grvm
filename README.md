# j
Gradle_Version_Manager




## 安装
  [下载](https://github.com/forget-the-bright/j/releases) 下载自己需要的版本, 到自己自定义的目录 修改可执行文件名称为j
  
  默认文件下载安装在用户目录下 ```.grvm```目录，目录下  ```versions```, ```downloads```, ```gradle```  分别是本地安装目录，安装包下载目录，当前使用的java版本目录 

  将 JAVA_HOME 配置为 ```USER_HOME\.grvm\gradle```  

  指定安装目录需要 添加环境变量 ```GRVM_HOME```
## 命令

### 列出

列出所有可安装版本
```
grvm ls-remote
```


列出本地安装版本
```
grvm ls
```

### 下载
```
grvm install 8
```


### 切换版本
```
grvm use 17
```





