package main

import "fmt"

func main(){
  var num1, num2 int
  var operator string

  fmt.Println("enter the first number")
  fmt.Scanf("%d", &num1)
  fmt.Println("enter the second number")
  fmt.Scanf("%d", &num2)
  fmt.Println("enter the operator")
  fmt.Scanf("%s", &operator)

  result, errorMessage :=calc(num1, num2, operator)
  if len(errorMessage)>0{
    fmt.Println(errorMessage)
  }else{
    fmt.Println("The result is:", result)
  }
}

func calc(num1 int, num2 int, operator string) (int, string){
  var result int
  var errorMessage string

  switch operator {
  case "+":
    result = num1+num2
  case "-":
    if num1>num2 {
        result=num1-num2

    } else{
      errorMessage="We cannot process if num1 is less than num2"

    }
  case "*":
    result = num1*num2
  case "/":
    if num2 == 0{
      errorMessage="We cannot process if num2 is 0"
    }else if num1>num2{
      result =num1/num2
    }else {
      errorMessage="We cannot process if num1 is less than num2 "

    }

  }
  return result, errorMessage
}
