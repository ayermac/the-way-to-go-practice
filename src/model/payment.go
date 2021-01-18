package model

import "fmt"

//微信支付
//支付宝
//银联
//银行卡
type PaymentArgs struct {
    AppId string
    MchId string
    Key string
    CallbackUrl string
}

func (this *PaymentArgs) Info()  {
    fmt.Println(this.AppId)
}