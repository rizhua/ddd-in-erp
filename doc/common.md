# 网关服务

## 约定说明

正式环境
```url
http://gw.rizhua.com/router/rest
```

测试环境
```url
http://dev-gw.rizhua.com/router/rest
```

全局编码

> 归类编码(10) + 业务编码(00) = 返回码(1000)

返回码|返回描述|英文描述|解决方案
:---:|---|---|---
1000|响应成功|Success|
|||
2000|接口异常：2000~2999||
2010|系统繁忙/超时|System busy or Request timeout|重新发起请求
2020|方法异常|-|
|||
3000|数据问题：3000~3999||
3010|数据为空|Is empty|
3020|无权访问|-|联系管理取得权限
3050|存储异常|-|
3051|插入失败|-|
3052|更新失败|-|
3053|删除失败|-|
3054|查询失败|-|
3060|缓存异常|-|
3070|令牌异常|Token|
|||
4000|用户错误：4000~4999||
4010|非法参数|Invalid parameter|
4020|解析异常|-|
4030|登录异常|-|
4040|未找到/不存在|not found|


公共参数

参数|类型|是否必填|描述|示例值
---|---|:---:|---|---
method|string(256)|是|接口名称|auth.login

## 接口示例

请求参数

参数|类型|必填|描述|示例值
---|---|:---:|---|---
userName|string(16)|是|登录账号|tom
password|string(32)|是|登录密码|e10adc3949ba59abbe56e057f20f883e
veriCode|string(4)|是|验证码|-

请求示例
```json
{
    "method": "auth.login",
    "userName": "tom",
    "password": "e10adc3949ba59abbe56e057f20f883e",
    "veriCode": "h3T4"
}
```

响应参数

参数|类型|必填|描述|示例值
---|---|:---:|---|---
code|int(5)|是|返回码|1000
desc|string|是|返回描述|-
data|any|否|返回数据|-

响应示例
```json
{
    "code": 1000,
    "desc": "Success",
    "data": "eyJleHAiOjE1Mjg4MjMwMzEsInVpZCI6MiwibGV2IjoxfQ==.cc1b705664d22d3dedcc311b34d32eea"
}
```