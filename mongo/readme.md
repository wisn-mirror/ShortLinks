## 安装

`//安装复制到 /usr/local/`
`export PATH=${PATH}:/usr/local/MongoDB/bin`

`查看版本 mongod -version`

`//创建文件夹`
`sudo mkdir -p /data/db`
`启动`
`mongod`


 `打开浏览器，输入localhost:27017`
`显示:`
`It looks like you are trying to access MongoDB over HTTP on the native driver port.`


`重新打开命令行，输入`
`mongo`

`停止MongoDB的时候一定要正确的退出，不然下次再次连接数据库会出现问题，使用下面的两行代码可以完成这一操作。`
`use admin;`
`db.shutdownServer();`

## 使用

