# IPMI转速自动控制工具 IPMI RPM Automation Tool

> 适用于无转速曲线且支持ipmitools的服务器的风扇转速根据温度梯度自动控制工具 
> 
> Automatic fan speed control tool based on temperature gradient for servers without speed profiles and with ipmitools support

> 理论上来说适用与所有使用AST2500同类BMC的服务器，控制基于BMC WEB页面的API，而非ipmitool命令
> 
> Theoretically applicable to all servers using AST2500-like BMC, the control is based on the BMC WEB api, not the ipmitool command
> 
> (其实是因为ipmitool别人操纵的地址我这都不对，AST2500的手册我也看不懂，所以只能用WebAPI
 
> 测试服务器： 联泰集群 G2232G2V2 双GPU服务器
> 
> Test Server: LTHPC G2232G2V2 Dual GPU Server
> 
> AST2500 BMC

## 使用
> 建议单开一个管理员账号供该程序使用，而非经常登录的账号，防止意外未能登出导致BMC页面无法登录。
>
> It is recommended that a separate administrator account be opened for use with this program, rather than a frequently logged in account, to prevent accidental failure to log out from preventing the BMC page from being logged in.

**Release编译好的二进制文件 Release compiled binaries**

```
./serverTemperature -f ./config.yaml
```
-f为指定配置文件，可不填，默认读取./config.yaml

-f for the specified configuration file, can not be filled, the default read . /config.yaml

**从源码启动 Launch from source**

开发环境: golang 1.23.2

Development environment: golang 1.23.2

不保证对过低版本的兼容

Compatibility with lower versions is not guaranteed

```
go run main.go -f ./config.yaml
```
-f为指定配置文件，同上

-f to specify the configuration file, as above
