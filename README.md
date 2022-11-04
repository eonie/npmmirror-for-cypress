解决 `Cypress` 国内安装下载缓慢的问题
## 解决思路
Cypress 提供了 `CYPRESS_DOWNLOAD_MIRROR` 环境变量，Cypress 检测到此环境变量后将通过环境变量指定的服务地址下载安装包。

举例来说，在 MacOS 操作系统中， 指定环境变量为 `CYPRESS_DOWNLOAD_MIRROR=http://localhost:8080`, Cypress 将通过 `http://localhost:8080/desktop/8.7.0?platform=darwin&arch=x64` 地址下载安装包。

因此可以编写一个服务解析请求中的参数后重定向到国内的镜像源，比如 [npmmirror](https://registry.npmmirror.com), npmmirror 提供的 Cypress 下载地址形式如下：
```
https://registry.npmmirror.com/-/binary/cypress/8.7.0/darwin-x64/cypress.zip
```
## 如何使用
使用 docker 镜像部署，最好部署在服务器上，方便项目组内的其他人也可以使用。
```
docker run -d --name=npmmirror-for-cypress -p 8090:8090 --restart=always eonie/npmmirror-for-cypress
```

前端工程中创建 `.npmrc` 文件, 在文件中配置 `CYPRESS_DOWNLOAD_MIRROR` 环境变量

```
CYPRESS_DOWNLOAD_MIRROR=http://<server>:<port>
```
> 如果不想自己搭建服务，也可以使用 `https://www.eonie.tk` 作为镜像源，不过不保证永久可用。

接下来就可以享受闪电般的速度啦。


### 其他问题
使用 `pnpm` 安装时未显示下载进度与安装进度