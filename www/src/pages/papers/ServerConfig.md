---
title: 配置服务器的操作的备忘（2023.2)
sections: [ 
  {id: "system", name: "适用系统"},
  {id: "environment", name: "更新环境"},
  {id: "disk", name: "挂载数据盘"},
  {id: "adduser", name: "增加用户"},
  {id: "install", name: "安装必要软件"},
  {id: "nginx", name: "配置 Nginx",sub: [
    {id: "single", name: "仅前端配置"},
    {id: "multiple", name: "前后端配置"},        
    {id: "enable", name: "启用配置"},      
  ]},
  {id: "supervisor", name: "配置 Supervisor",sub: [
    {id: "extension", name: "文件后缀"},      
    {id: "common", name: "常用命令"},      
  ]},
]
---

<script setup>
import {defineEmits,onMounted} from "vue";

const emits = defineEmits(["syncMeta"]);

onMounted(()=>{
    emits("syncMeta", frontmatter);
    
})
</script>

# 配置服务器的操作的备忘（2023.2) {#title .arco-typography}
  因为用 Docker 之于我的需要有些「杀鸡牛刀」，另外一些细节上的原因导致不适用。这样新服务器需要手动配置，故记录一下配置过程以备忘。
  
## 适用系统 {#system .arco-typography}
- Ubuntu v22.04

## 更新环境 {#environment .arco-typography}
  须先更新 Ubuntu 当前环境的默认依赖到最新。
```powershell {.line-numbers .match-braces .rainbow-braces}
  sudo apt update
  sudo apt upgrade -y
```

## 挂载数据盘 {#disk .arco-typography}
  容量小于 2TB 且控制台挂载。
```powershell {.line-numbers .match-braces .rainbow-braces}
  # 查看磁盘名称
  sudo fdisk -l 
  
  # 格式化数据盘
  sudo mkfs -t {文件系统格式} {磁盘名称}
  sudo mkfs -t ext4 /dev/vdb
  
  # 设置挂载点
  sudo mount {磁盘名称} {挂载点}
  sudo mount /dev/vdb /data
  # 查看挂载结果
  sudo df -TH
  
  # 设置开机自动挂载磁盘
  # 使用文件系统的 UUID 方式
  sudo blkid {磁盘名称}
  sudo blkid /dev/vdb
  # 编辑 fstab
  sudo vi /etc/fstab
  # 在 fstab 写入磁盘信息
  {设备信息} {挂载点} {文件系统格式} {文件系统安装选项} {文件系统转储频率} {启动时的文件系统检查顺序}
  UUID=d489ca1c-5057-4536-81cb-ceb2847f9954 /data  ext4 defaults     0   0
  # 执行以下命令，检查 /etc/fstab 文件是否写入成功。
  mount -a 
```

## 增加用户 {#adduser .arco-typography}
```powershell {.line-numbers .match-braces .rainbow-braces}
  # 增加用户
  sudo useradd {用户名}  -d {指定用户目录} -m -s /bin/bash
  sudo useradd www -d /data/web -m -s /bin/bash
  # 切到刚增用户继续执行，进入用户根目录
  # 配置密钥登录
  ssh-keygen -t rsa
  cd .ssh
  touch authorized_keys
  cat id_rsa.pub >> authorized_keys
  # 给密钥配置权限
  sudo chmod 600 authorized_keys
  sudo chmod 700 ~/.ssh
```

## 安装必要软件 {#install .arco-typography}
```powershell {.line-numbers .match-braces .rainbow-braces}
  sudo apt install nginx -y
  sudo apt install supervisor -y
```

## 配置 Nginx {#nginx .arco-typography}
```powershell {.line-numbers .match-braces .rainbow-braces}
  cd /etc/nginx/sites-available
```
  仅前端配置 {#single}
```powershell {.line-numbers .match-braces .rainbow-braces}
  server {
        listen 443 ssl;
        server_name www.xxx.com xxx.com;

        access_log off;
        error_log /.../logs/nginx_error.log error;


        ssl_certificate "/.../ssl-bundle.crt";
        ssl_certificate_key "/...//xxx.com.key";
        ssl_session_cache shared:SSL:1m;
        ssl_session_timeout 10m;
        ssl_ciphers HIGH:!aNULL:!MD5;
        ssl_protocols TLSv1.2 TLSv1.3;
        ssl_prefer_server_ciphers on;

        #开启gzip
        gzip  on;
        #低于1kb的资源不压缩
        gzip_min_length 1k;
        #压缩级别1-9，越大压缩率越高，同时消耗cpu资源也越多，建议设置在5左右。
        gzip_comp_level 5;
        #需要压缩哪些响应类型的资源，多个空格隔开。不建议压缩图片.
        gzip_types text/plain application/javascript application/x-javascript text/javascript text/xml text/css;
        #是否添加“Vary: Accept-Encoding”响应头
        gzip_vary on;

        location / {
                root /.../dist;
                try_files $uri $uri/ /index.html;
                index  index.html index.htm;
                }
  }

  server {
        listen 80;
        server_name www.xxx.com xxx.com; 

        rewrite ^(.*)$ https://${server_name}$1 permanent;
  }
```
  前后端配置 {#multiple}
```powershell {.line-numbers .match-braces .rainbow-braces}
  server {
        listen 443 ssl;
        server_name www.xxx.com;

        access_log off;
        error_log /.../logs/nginx_error.log error;


        ssl_certificate "/.../ssl-bundle.crt";
        ssl_certificate_key "/.../xxx.com.key";
        ssl_session_cache shared:SSL:1m;
        ssl_session_timeout 10m;
        ssl_ciphers HIGH:!aNULL:!MD5;
        ssl_protocols TLSv1.2 TLSv1.3;
        ssl_prefer_server_ciphers on;

        #开启gzip
        gzip  on;
        #低于1kb的资源不压缩
        gzip_min_length 1k;
        #压缩级别1-9，越大压缩率越高，同时消耗cpu资源也越多，建议设置在5左右。
        gzip_comp_level 5;
        #需要压缩哪些响应类型的资源，多个空格隔开。不建议压缩图片.
        gzip_types text/plain application/javascript application/x-javascript text/javascript text/xml text/css;
        #是否添加“Vary: Accept-Encoding”响应头
        gzip_vary on;

        location / {
                root /.../dist;
                try_files $uri $uri/ /index.html;
                index  index.html index.htm;
                }
  }

  server {
        listen 443 ssl;
        server_name xxx.com;

        access_log off;
        error_log /.../logs/nginx_error.log error;


        ssl_certificate "/.../ssl-bundle.crt";
        ssl_certificate_key "/.../xxx.com.key";
        ssl_session_cache shared:SSL:1m;
        ssl_session_timeout 10m;
        ssl_ciphers HIGH:!aNULL:!MD5;
        ssl_protocols TLSv1.2 TLSv1.3;
        ssl_prefer_server_ciphers on;

        #开启gzip
        gzip  on;
        #低于1kb的资源不压缩
        gzip_min_length 1k;
        #压缩级别1-9，越大压缩率越高，同时消耗cpu资源也越多，建议设置在5左右。
        gzip_comp_level 5;
        #需要压缩哪些响应类型的资源，多个空格隔开。不建议压缩图片.
        gzip_types text/plain application/javascript application/x-javascript text/javascript text/xml text/css application/json;
        #是否添加“Vary: Accept-Encoding”响应头
        gzip_vary on;

        location / {
                proxy_pass http://localhost:8080;
                proxy_set_header Host $host;
                proxy_set_header X-Real_IP $remote_addr;
                proxy_set_header X-Forwarded-For  $proxy_add_x_forwarded_for;
                proxy_set_header Cookie $http_cookie;
        }
  }

  server {
        listen 80;
        server_name www.xxx.com xxx.com; 

        rewrite ^(.*)$ https://${server_name}$1 permanent;
  }
```
  启用配置 {#enable}
```powershell {.line-numbers .match-braces .rainbow-braces}
  cd /etc/nginx/sites-enabled
  # 建立软链接
  sudo ln -s {被引用目标文件} {软链接名}
  sudo ln -s /etc/nginx/sites-available/xxx.com xxx.com
  # 解决 nginx  failed (13: Permission denied)
  # The best solution in that case would be to add www-data to username group:
  sudo gpasswd -a www-data {目标用户名}
  # 测试配置
  sudo nginx -t
  # 生效配置
  sudo systemctl reload nginx
```
## 配置 Supervisor {#supervisor  .arco-typography}
  文件后缀 .conf {#extension}
```powershell {.line-numbers .match-braces .rainbow-braces}
  #项目名
  [program:name]
  #脚本目录
  directory=/.../test
  #脚本执行命令
  command=/.../test/go_build_core_linux
  
  #脚本运行的用户身份 
  user=www
  #supervisor启动的时候是否随着同时启动，默认True
  autostart=true
  #当程序exit的时候，这个program不会自动重启,默认unexpected，设置子进程挂掉后自动重启的情况，
  #有三个选项，false,unexpected和true。如果为false的时候，无论什么情况下，都不会被重新启动，
  #如果为unexpected，只有当进程的退出码不在下面的exitcodes里面定义的时候
  autorestart=true
  #这个选项是子进程启动多少秒之后，此时状态如果是running，则我们认为启动成功了。默认值为1
  startsecs=1
  #当进程启动失败后，最大尝试启动的次数，默认为3次
  startretries=3
  #这个是当我们向子进程发送stopsignal信号后，到系统返回信息给supervisord所等待的最大时间
  #默认为10秒
  stopwaitsecs=10
  #把stderr重定向到stdout，默认 false
  redirect_stderr=false
  #日志输出 
  stdout_logfile=/.../logs/supervisor_common.log
  #stdout日志文件大小，默认 50MB
  stdout_logfile_maxbytes=50MB
  #stdout日志文件备份数
  stdout_logfile_backups=10
  stderr_logfile=/.../logs/supervisor_error.log
  #stderr日志文件大小，默认 50MB
  stderr_logfile_maxbytes=50MB
  #stderr日志文件备份数
  stderr_logfile_backups=10
  #环境参数
  environment=MODE="LIVE",DEBUG="false",DB_SOURCE="host=uri port=5432 user=username password=xxx dbname=dbname sslmode=disable",
  DB_TYPE="postgres",REDIS_DB="",REDIS_PWD="",REDIS_URI="localhost:6370",SECRET_KEY=""
```
  常用命令 {#common}
```powershell {.line-numbers .match-braces .rainbow-braces}
  sudo systemctl enable supervisord # 开机自启动
  sudo systemctl start supervisord
  sudo systemctl restart supervisord
  sudo systemctl reload supervisord
  sudo systemctl update supervisord
  sudo systemctl status supervisord
  sudo systemctl stop supervisord
```
