// Create struct Person with (Name and ID). update name of one person
package main 

import "fmt"

type Person struct{
	Name string
	id int
}

func update(p *Person, id int){
	p.id = id;
}


func main(){
	akhilesh := Person{Name: "Akhilesh", id: 20}

	fmt.Println(akhilesh)

	update(&akhilesh, 30)

	fmt.Println(akhilesh)
}

