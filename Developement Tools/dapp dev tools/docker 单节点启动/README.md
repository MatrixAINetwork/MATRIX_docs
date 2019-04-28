### 基于docker的单节点启动 - 一键启动，跨平台，钱包，浏览器均可

第一步：下载并安装 docker，运行后会有提示：Docker Desktop is running

第二步：下载 docker-compose.yaml 

第三步：在命令行窗口执行 docker-compose -f docker-compose.yml up -d 启动

第四步：登录

- http://127.0.0.1/ -钱包

- http://127.0.0.1:3000 浏览器

注意：docker-compose.yaml 文件内置IP参数，linux mac修改为172.0.18.2；Windows保留为默认 10.0.75.2


