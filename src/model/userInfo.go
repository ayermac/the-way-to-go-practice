package model

type userInfo struct {
    Name   string
    Age    int
    Height float32
    Hobby []string
    EduSchool string
    MoreInfo map[string]interface{}
}

//工厂模式：生成对象
func NewUserInfo(name string, age int, height float32, eduschool string, hobby []string, moreInfo map[string]interface{}) *userInfo  {
    return &userInfo{
        Name:      name,
        Age:       age,
        Height:    height,
        Hobby:     hobby,
        EduSchool: eduschool,
        MoreInfo:  moreInfo,
    }
}
