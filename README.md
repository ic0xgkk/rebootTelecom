# rebootTelecom

## 背景

众所周知，电信装宽带送的光猫一向都特别垃圾，毕竟是消费级的产品。

在没办法更换其硬件的情况下，定时重启成了避免不稳定因素的唯一办法。

该工具主要用在公司里，公司使用了电信定制的华为HN8145V这款型号的光猫。这个机器在公司里每隔两三天就会抽风一次，以至于随机丢包并且非常严重，抽风时完全无法正常用网。为了避免突发疫情远程办公时，无法正常回连办公室的机器，只能写了这个工具定时重启一下光猫。

## 验证环境

手边只有电信定制的华为HN8145N的设备，固件版本号为21HWW41930010 。

其他机器请自行测试~

## 已知问题

执行重启成功后，由于光猫直接就给重启了，而并没有按照流程把所有进程关闭后再重启。因此，TCP连接会处于卡死状态。

如下：

```bash
# ./rebootTelecom -host "192.168.1.1" -username "***" -password "***" 
panic: dial patch page response failed: Post "http://192.168.1.1:8080/html/ssmp/devmanage/set.cgi?x=InternetGatewayDevice.X_HW_DEBUG.SMP.DM.ResetBoard&RequestFile=html/ssmp/devmanage/e8cdevicemanormal.asp": read tcp *.*.*.*:60386->192.168.1.1:8080: read: connection reset by peer

goroutine 1 [running]:
main.main()
        /data/project/rebootTelecom/main.go:34 +0x605
```

等待机器重启完成之后，重新发出的TCP包会在没有被跟踪的情况下直接被重置，就会报出上边这个错误。

但是！至少看到这个错误时，你可以非常确信重启成功了（笑

## 使用方法

如下：

```bas
# ./rebootTelecom -h
Usage of ./rebootTelecom:
  -host string
        The host address without port number, such as '192.168.1.1' or 'm.example.com' (default "localhost")
  -password string
        The password used to login your HN8145V (default "password")
  -username string
        The username used to login your HN8145V (default "username")
```

只需要三个控制台参数，对应分别是：

* 主机地址。请直接填地址**不要带端口号**，比如：192.168.1.1
* 密码
* 用户名

需要注意的是，这个用户名和密码需要是能登录进去半超管界面的。**请打开你的机器的8080端口的页面**，**而不是默认的80端口页面**，尝试使用凭据登录，如果能登录成功，再将凭据填写到此处即可。

执行实例：

```bash
./rebootTelecom -host "主机地址" -username "用户名" -password "密码" 
```

## 定时方式

可以运行在Linux下，通过Crontab定时触发。

但是！**请确保你的Linux系统本身能够通过NTP从互联网同步时间，如果系统时间错误，至于后果你也应该知道了。**

安装方式：

```bash
# 请使用root执行
crontab -e
# 此时会进入crontab任务的编辑页面

# 加入下方内容
0 6 * * * /bin/rebootTelecom -host "主机地址" -username "用户名" -password "密码" > /tmp/rebootTelecom.log
```

**请注意！**

* 下载二进制文件后，请给可执行权限~
* 执行路径请根据实际情况自行修改
* 最后的`>`将控制台输出重定向到了`/tmp/rebootTelecom.log`中，可以去查看输出。每次执行后会被覆盖

