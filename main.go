package main

import "fmt"

func main() {
  fmt.Println("Программа для расчета роста бамбука")
  
  fmt.Println("На бамбуковой плантации завелись гусеницы. Они спят днем и едят бамбук ночью. Бамбук генно-модифицированный, растет только при свете дня, зато  очень быстро.")

   fmt.Println("Введите высоту бамбука при посадке:")
  var bambooStart int
  fmt.Scan(&bambooStart)

  fmt.Println("Введите скорость роста бамбука:")
  var bambooHeight int
  fmt.Scan(&bambooHeight)

  fmt.Println("Введите скорость поедания бамбука гусеницами:")
  var bambooLoses int
  fmt.Scan(&bambooLoses)

  fmt.Println("Введите высоту бамбука, при которой его можно будет срубить:")
  var totalHeight int
  fmt.Scan(&totalHeight)

  var totalDay int
  //300=100+x(50-20) => x=(300 - 100)/(50-20)
  totalDay = (totalHeight - bambooStart) / (bambooHeight - bambooLoses)

  


  fmt.Println("Бамбук можно срубить через" , totalDay , "дней.")






}