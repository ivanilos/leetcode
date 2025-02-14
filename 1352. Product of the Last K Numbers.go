type ProductOfNumbers struct {
    prefixProd []int
    zeroQt []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        []int{},
        []int{},
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    sz := len(this.prefixProd)

    if sz == 0 {
        if num == 0 {
            this.prefixProd = []int{1}
            this.zeroQt = []int{1}
        } else {
            this.prefixProd = []int{num}
            this.zeroQt = []int{0}
        }
    } else {
        if num == 0 {
            this.prefixProd = append(this.prefixProd, 1)
            this.zeroQt = append(this.zeroQt, this.zeroQt[sz - 1] + 1)
        } else {
            this.prefixProd = append(this.prefixProd, this.prefixProd[sz - 1] * num)
            this.zeroQt = append(this.zeroQt, this.zeroQt[sz - 1])
        }
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    sz := len(this.prefixProd)

    
    zeroes := this.zeroQt[sz - 1]
    if sz - k - 1 >= 0 {
        zeroes -= this.zeroQt[sz - k - 1]
    }

    if zeroes > 0 {
        return 0
    }

    prod := this.prefixProd[sz - 1]
    if sz - k - 1 >= 0 {
        prod /= this.prefixProd[sz - k - 1]
    }
    return prod
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
