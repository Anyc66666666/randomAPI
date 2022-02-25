randomAPI  接口说明

配置文件 ./workdir/config/config.yaml

mysql数据库 dsn: 账号root 密码root  要创建数据库名为 sentences

port 监听端口  eg:8468

host 域名   eg:http://127.0.0.1   



管理页面：  域名+端口/admin  增删改查句子 和 上传文件

随机音乐： 域名+端口/music   上传文件 进入管理页面或进入文件目录./music

随机视频： 域名+端口/video   上传文件 进入管理页面或进入文件目录./video

随机图片： 域名+端口/picture   上传文件 进入管理页面或进入文件目录./picture

随机句子：域名+端口/sentence   若要增删改句子 进入文件目录 ./sentence/st.go

随机句子（mysql）：域名+端口/sentences   若要增删改句子进入管理页面



```
/sentences/all    get 查询所有sentences  
/sentences/add    post 增添sentences    [json]  "sentences":"test"
/sentences/update  post  更新sentences  [json]  "id":1,"sentences":"test"
/sentences/delete  post 删除sentences   [json]  "id":1
```



目录 ./cache 为缓存文件，当上传的文件格式不是mp3,mp4,jpg,jpeg,png则会把文件保存在cache中

