简介
---
hdnslog 是一款基于DNSLog-GO二次开发的监控 DNS 解析记录的工具，自带多用户WEB界面


安装
---
# 1.修改配置文件 config.ini

```
HTTP:
  port: 8000 //http web监听端口
  #{"token":"用户对应子域名"}
  user: { "admin": "admin" } //用户admin 对应的dnslog子域名是 admin.demo.com
  consoleDisable: false  //是否关闭web页面
Dns:
  domain: demo.com //dnslog域名
```

