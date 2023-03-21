B站书文练手项目

b站地址:https://www.bilibili.com/video/BV12r4y1w7Vr/?spm_id_from=333.788.recommend_more_video.0&vd_source=df3b7ca79ea2e3d34de8f4506bb98c45


2023年03月21日14:02:35 备忘录:
https://www.bilibili.com/video/BV12r4y1w7Vr?p=3&vd_source=df3b7ca79ea2e3d34de8f4506bb98c45

看到28分处;

# 庚中 安装说明
版本号:mysql-8.0.30-winx64
数据库密码:123456;
设置密码命令:
```js
use mysql;

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';

flush privileges;
quit;

//常用命令:
启动/关闭
net start/stop mysql

//进入mysql控制台命令密码填123456
mysql -uroot -p
//退出mysql
quit;


//显示所有数据库名
show databases;
//选择数据库
use 数据库名;
//显示表名
show tables;
//根据表名查询
select * from 表名;

```

