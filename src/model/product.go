package model

type Product struct {
    name string
    price float32
}

// getter setter
func (this *Product) SetName(name string)  {
    this.name = name
}

func (this *Product) SetPrice(price float32)  {
    this.price = price
}

func (this *Product) GetName() string  {
    return this.name
}

func (this *Product) GetPrice() float32  {
    return this.price
}
