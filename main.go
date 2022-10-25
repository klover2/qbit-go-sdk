package main

import (
	"fmt"

	"github.com/klover2/qbit-go-sdk/client"
)

func main() {
	// params := map[string]interface{}{
	// 	"id":                  "ee74c872-8173-4b67-81b1-5746e7d5ab88",
	// 	"accountId":           nil,
	// 	"holderId":            "d2bd6ab3-3c28-4ac7-a7c4-b7eed5eee367",
	// 	"currency":            "USD",
	// 	"settlementCurrency":  nil,
	// 	"counterparty":        "SAILINGWOOD;;US;1800948598;;091000019",
	// 	"transactionAmount":   11,
	// 	"fee":                 0,
	// 	"businessType":        "Inbound",
	// 	"status":              "Closed",
	// 	"transactionTime":     "2021-11-22T07:34:10.997Z",
	// 	"transactionId":       "124d3804-defa-4033-9f30-1d8b0468e506",
	// 	"clientTransactionId": nil,
	// 	"createTime":          "2021-11-22T07:34:10.997Z",
	// 	"appendFee":           0,
	// }
	// fmt.Println(utils.Sign(utils.JoinStr(params), "25d55ad283aa400af464c76d713c07ad"))
	// => 8287d5539c03918c9de51176162c2bf7065d5a8756b014e3293be1920c20d102

	client := client.NewClient("qbit1f6efee44ceb8ca2", "8f70d42a1393802aebf567be27a47879")

	client.SetBaseUrl("http://127.0.0.1:3000")

	// res, err := client.Get("http://127.0.0.1:3000/open-api/oauth/authorize?clientId=qbit1f6efee44ceb8ca2")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res.Content)
	body := &map[string]interface{}{"clientId": "qbit1f6efee44ceb8ca2", "clientSecret": "8f70d42a1393802aebf567be27a47879", "code": "921dcd768f48e6a109f85d367f027712"}

	res, err := client.Post("http://127.0.0.1:3000/open-api/oauth/access-token", body, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Content)

}
