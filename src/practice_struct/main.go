package main

import (
    "errors"
    "fmt"
    "unsafe"
)

type userInfo struct {
    name string
    age int
    height float32
    hobby []string
    eduSchool string
    moreInfo map[string]interface{}
}

type Integer int

type role struct {
    user userInfo
    authorization Integer //1=超级管理员 2=管理员 3=普通用户
}

type calculate struct {
    left int
    right int
}

func (c *calculate) add() int {
    return c.left + c.right
}

func (c *calculate) sub() int {
    return c.left - c.right
}

func (c *calculate) multiply() int {
    return c.left * c.right
}

func (c *calculate) divide() (float32, error) {
    if c.right == 0 {
        return 0, errors.New("right num can not be empty")
    }

    return float32(c.left / c.right), nil
}

func (c *calculate) mod() (int, error) {
    if c.right == 0 {
        return 0, errors.New("right num can not be empty")
    }
    return c.left % c.right, nil
}

func main()  {
    // 结构体 struct
    // 定义自定义类型
    // 1.type
    type integer int
    var intVariables int
    var integerVariables integer
    fmt.Println(intVariables, integerVariables)

    // 2.struct
    //type 结构体名 struct {
    //    field1 type1
    //}

    var Jason userInfo
    Jason.name = "jason"
    Jason.age = 26
    Jason.height = 181
    Jason.eduSchool = "复旦大学"
    Jason.hobby = []string{"coding", "运动", "旅游"}
    Jason.moreInfo = map[string]interface{}{
        "work":"百度",
        "duty":"develop",
    }

    fmt.Printf("name=%s,hobby=%v\n", Jason.name, Jason.hobby)

    Amy := userInfo{
        name: "帅哥",
        age: 18,
        hobby: []string{"看电影","唱歌"},
        moreInfo: map[string]interface{}{
            "role":"演员",
            "earnmoney":3000,
        },
    }
    Amy.name = "test"
    fmt.Printf("name=%s,hobby=%v\n", Amy.name, Amy.hobby)
    fmt.Printf("name=%p,age=%p,hobby=%p\n", &Amy.name, &Amy.age, &Amy.hobby)
    fmt.Printf("name=%d,age=%d,hobby=%d\n", unsafe.Sizeof(Amy.name), unsafe.Sizeof(Amy.age), unsafe.Sizeof(Amy.hobby))

    superAdmin := role{
        user: userInfo{
            name:      "超级管理员",
            age:       0,
            height:    0,
            hobby:     nil,
            eduSchool: "",
            moreInfo:  nil,
        },
        authorization: 1,
    }

    admin := role{
        user: userInfo{
            name:      "管理员",
            age:       0,
            height:    0,
            hobby:     nil,
            eduSchool: "",
            moreInfo:  nil,
        },
        authorization: 2,
    }

    fmt.Println(superAdmin, admin)

    c := calculate{
        left:  4,
        right: 3,
    }

    fmt.Printf("add=%d\n", c.add())
    fmt.Printf("sub=%d\n", c.sub())
    fmt.Printf("multiply=%d\n", c.multiply())

    d, err := c.divide()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(d)

    m, err := c.mod()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("mod=%d\n", m)
}
