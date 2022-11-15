# account账号服务

## 功能列表
* 用户登录
* 组织管理接口
* 用户管里接口



## 单元测试覆盖率查询
 https://zhuanlan.zhihu.com/p/365258413

1. 统计覆盖率,包含明细,错误信息
    
    go test -v -race $(go list ./... |grep -v /cmd |grep -v /vendor|grep -v /config|grep -v /api|grep -v /swagger|grep -v /third_party) -coverprofile=coverage.out

    * 查device_org 服务的测试情况& 覆盖率
    go test -v -race $(go list ./internal/module/device_org/... |grep -v /cmd |grep -v /vendor|grep -v /config|grep -v /api|grep -v /swagger|grep -v /third_party) -coverprofile=coverage.out
    
2. 统计覆盖率, 有测试文件的packagee的每一个方法统计和总计覆盖率. 需要先执行完1.
    
    go tool cover -func=coverage.out
