# IPMI转速自动控制工具

> 适用于无转速曲线且支持ipmitools的服务器的风扇转速根据温度梯度自动控制工具 
> / Automatic fan speed control tool based on temperature gradient for servers without speed profiles and with ipmitools support

> 理论上来说适用与所有使用AST2500同类BMC的服务器，控制基于BMC WEB页面，而非ipmitool命令
 
> 测试服务器： 联泰集群 G2232G2V2 双GPU服务器
> 
> AST2500 BMC

## 使用
**Release编译好的二进制文件**

```
./serverTemperature -f ./config.json
```
-f为指定配置文件，可不填，默认读取./config.json


**从源码启动**

```
go run main.go -f ./config.json
```
-f为指定配置文件，同上

