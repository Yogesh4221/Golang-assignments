package main

import (
	"fmt"
)
var data map[int]Customer
var customer Customer

type data map[int]Customer

type customer  struct {
	CustomerID int
	Name       string
	Address      int
}

func init(){
	data =make(map[int]customer)
}
	func (c customer)add (k int , v customer){	

if _, check :=data [k];check {
	fmt.Println(k,"customer not added")

}else {
	data[k]=v
	fmt.println (k,"customer added..")
   }
}
func (c customer)update(k int , v customer){
if _, check :=data [k];check {
	fmt.Println(k,"customer updated")
}else{
	fmt.println("not updated.")
  }

}
func (c customer) delete(k int) {
	if _,check :=data[k]; check {
		delete(data,k)
	fmt.println(k,"customer deleted.")

}else {
	fmt.println("not deleted..")
   }
}
func (c customer)update(k int) , (customer,string){
	if_; check :=data [k];check {
		return data{K},"success.."
	}
	return customer{},"failed.."
}
func (c Customer)getall() data{
	return data
}
func main(){
	C1 :=customer {001,"Ajay", Andhra}
	C2 :=Customer {002,"vijay", Tamilnadu}
	fmt.Println(data)
}
	customer.add {001, c1)
    Customer.add {002, c2)
	Customer.add {003, Customer {003,"ajith",karnataka}
fmt.println(data)
customer ;= customer,getall()
fmt.println(customers)

