package main

import (
    "fmt"
    "the-way-to-go-pratice/src/model"
)

func main()  {
    Jason := model.NewUserInfo(
        "Jason",
        18,
        0,
        "北京电影学院",
        []string{"运动","coding"},
        map[string]interface{}{
            "work":"Ali",
        },
    )

    fmt.Println(Jason.Name)

    product := &model.Product{}
    product.SetName("GoLand")
    product.SetPrice(100)

    fmt.Println(product.GetName())
    fmt.Println(product.GetPrice())
    
    alipay := model.AliPay{
        PaymentArgs:  model.PaymentArgs{
            AppId: "alipay",
            MchId: "alipay",
            Key: "alipay",
            CallbackUrl: "https://alipay",
        },
        AliPayOpenId: "aliPayOpenId",
    }

    fmt.Println(alipay.AppId)
    alipay.Info()

    weixinpay := model.WeixinPay{
        PaymentArgs:  model.PaymentArgs{
            AppId: "weixinpay",
            MchId: "weixinpay",
            Key: "weixinpay",
            CallbackUrl: "https://weixinpay",
        },
        WeixinOpenId: "weixinPayOpenId",
    }

    fmt.Println(weixinpay)
}


