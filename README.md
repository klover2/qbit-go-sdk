<p style="text-align: center;">
  <h1 align="center"><a href="javascript:void(0);">qbit-go-sdk</a></h1>
</p>

## Qbit 概念

开发者 API 旨在允许企业与 Qbit 系统集成，并轻松将其作为其工作流程的一部分。该 API 允许开发者使用【全球账户】、【量子卡】业务等。

## 项目状态

当前版本`1.0.0`为正式版本。暂时支持了 auth 相关的接口，其他接口带后续完善，同时也提供了 Qbit Api 所需的 Post、put、delete、get 请求，方便使用者更好调用其他接口，具体使用请看下面代码示例。

`注意`：请商户的专业技术人员在使用时注意系统和软件的正确性和兼容性，以及带来的风险。

## 环境要求

- go 1.16+

### 安装

#### 1、使用 Go Modules 管理你的项目

如果你的项目还不是使用 Go Modules 做依赖管理，在项目根目录下执行：

```shell
go mod init
```

#### 2、无需 clone 仓库中的代码，直接在项目目录中执行：

```shell
go get github.com/klover2/qbit-go-sdk

go mod vendor # 把代码下载到当前项目中
```

来添加依赖，完成 `go.mod` 修改与 SDK 下载。

## 名词解释

- Client，合作伙伴在 Qbit 我们称之为 Client。
- Account， 合作伙伴的客户在 Qbit 我们称之为 Account
- clientId，商户 id，请联系我们申请。
- clientSecret，商户密钥，用于签名，请联系我们申请。

## 开始

### 获取 access token

```go
package main

import (
	"fmt"

	qbit "github.com/klover2/qbit-go-sdk/services"
)

func main() {
	authService := qbit.NewAuthService("qbit1f6efee44ceb8ca2", "8f70d42a1393802aebf567be27a47879", "https://api-global.qbitnetwork.com")
	codeRes, err := authService.GetCode("123", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := authService.GetAccessToken(codeRes.Code)
	fmt.Println(err)
	fmt.Println(res)
	fmt.Println(res.AccessToken)
}
```

### 刷新 access token

```go
package main

import (
	"fmt"

	qbit "github.com/klover2/qbit-go-sdk/services"
)

func main() {
	authService := qbit.NewAuthService("qbit1f6efee44ceb8ca2", "8f70d42a1393802aebf567be27a47879", "https://api-global.qbitnetwork.com")
	res, err := authService.RefreshToken("0f8ef03f4f63e506bff12d8443b7c6c92d4c5c561d615e1819d7d6f3fac558c7")
	fmt.Println(err)
	fmt.Println(res)
}
```

### 调用其他接口示例

```go
package main

import (
	"fmt"

	qbit "github.com/klover2/qbit-go-sdk/services"
)

func main() {
	service := qbit.NewRequestService()
	service.SetAccessToken("c93bdb64d687a3fa9e13b9742d0fa9397e495245")
	data := map[string]interface{}{
		"name": "预算名",
		"cost": 10,
	}
	// res, err := service.GetRequest("https://api-global.qbitnetwork.com/open-api/v1/budget", data)
	res, err := service.PostRequest("https://api-global.qbitnetwork.com/open-api/v1/budget", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
```

## 敏感信息加解密

### 加密-HmacSHA256

```go
package main

import (
	"fmt"

	"github.com/klover2/qbit-go-sdk/utils"
)

func main() {
	params := map[string]interface{}{
		"id":                  "ee74c872-8173-4b67-81b1-5746e7d5ab88",
		"accountId":           nil,
		"holderId":            "d2bd6ab3-3c28-4ac7-a7c4-b7eed5eee367",
		"currency":            "USD",
		"settlementCurrency":  nil,
		"counterparty":        "SAILINGWOOD;;US;1800948598;;091000019",
		"transactionAmount":   11,
		"fee":                 0,
		"businessType":        "Inbound",
		"status":              "Closed",
		"transactionTime":     "2021-11-22T07:34:10.997Z",
		"transactionId":       "124d3804-defa-4033-9f30-1d8b0468e506",
		"clientTransactionId": nil,
		"createTime":          "2021-11-22T07:34:10.997Z",
		"appendFee":           0,
	}
	fmt.Println(utils.Sign(utils.JoinStr(params), "25d55ad283aa400af464c76d713c07ad"))
	// => 8287d5539c03918c9de51176162c2bf7065d5a8756b014e3293be1920c20d102
}
```

## 联系我们

如果你发现了**BUG**或者有任何疑问、建议，请通过 issue 进行反馈。

也欢迎访问我们的[官网](https://www.qbitnetwork.com/#/)。

## Keywords

QBIT 全球账户 量子卡
