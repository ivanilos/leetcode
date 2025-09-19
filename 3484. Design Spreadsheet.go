type Spreadsheet struct {
    cells map[string]int
}


func Constructor(rows int) Spreadsheet {
    return Spreadsheet{
        map[string]int{},
    }
}


func (this *Spreadsheet) SetCell(cell string, value int)  {
    this.cells[cell] = value
}


func (this *Spreadsheet) ResetCell(cell string)  {
    this.cells[cell] = 0
}


func (this *Spreadsheet) GetValue(formula string) int {
    splitted := strings.Split(formula[1:], "+")

    result := 0
    for _, token := range splitted {
        val, err := strconv.Atoi(token)

        if err == nil {
            result += val
        } else {
            result += this.cells[token]
        }
    }

    return result
}


/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
