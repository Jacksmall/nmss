# nginx-multi-servers-spike
使用nginx upstream 本机启动多个端口转发服务到后端服务器,nginx配置如下
```
upstream load_rule {
	server 127.0.0.1:3001 weight=1;
	server 127.0.0.1:3002 weight=2;
	server 127.0.0.1:3003 weight=3;
	server 127.0.0.1:3004 weight=4;
}

server {
	listen 80;
	server_name load_rule.com www.load_rule.com;

	access_log  /usr/local/var/log/nginx/www.load_rule.com.access.log;
    error_log   /usr/local/var/log/nginx/www.load_rule.com.error.log;

	location / {
		proxy_pass http://load_rule;
	}
}
```
命令端运行：
$sudo nginx -s reload

修改/etc/hosts,新增:
127.0.0.1 www.load_rule.com

后端服务器启动:
go run main.go


命令端运行:
$ab -n 5000 -c 100 http://www.load_rule.com/buy/ticket

项目下会新建stat.log日志,可以查看日志，4个端口的秒杀按照配置均匀处理！
