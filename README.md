# grvm
Gradle_Version_Manager

![image](https://user-images.githubusercontent.com/56473277/229162036-e5c44612-d358-47a3-8a56-b05e954b63a8.png)

## 安装
  [下载](https://github.com/forget-the-bright/grvm/releases) 下载自己需要的版本, 到自己自定义的目录 修改可执行文件名称为j
  
  默认文件下载安装在用户目录下 ```.grvm```目录，目录下  ```versions```, ```downloads```, ```gradle```  分别是本地安装目录，安装包下载目录，当前使用的gradle版本目录 

  将 GRADLE_HOME 配置为 ```USER_HOME\.grvm\gradle```

  指定安装目录需要 添加环境变量 ```GRVM_HOME``` 将 GRADLE_HOME 配置为 ```GRVM_HOME\gradle```
## 命令

### 列出

列出所有可安装版本
```
grvm ls-remote
```
![image](https://user-images.githubusercontent.com/56473277/229162310-b2c66fdc-bd7f-4c56-bd12-c02586aedc27.png)



列出本地安装版本
```
grvm ls
```
![image](https://user-images.githubusercontent.com/56473277/229162099-a74e00df-38ea-432d-9871-5f31848bf1bc.png)

### 下载
```
grvm install 7.6.1
```
![image](https://user-images.githubusercontent.com/56473277/229162742-f8875d31-1977-47ac-9956-c45249144b80.png)


### 切换版本
```
grvm use 6.9.4
```
![image](https://user-images.githubusercontent.com/56473277/229162893-3f63ce42-8dd2-47ef-b1c3-eb3af9e37524.png)





