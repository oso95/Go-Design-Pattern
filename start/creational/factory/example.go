package main

import "fmt"

func main(){
	mag1, _ := newPublication("magazine", "magazine1", 50, "publisher1")
	mag2, _ := newPublication("magazine", "magazine2", 40, "publisher2")
	news1, _ := newPublication("newspaper", "newspaper1", 60, "publisher3")
	news2, _ := newPublication("newspaper", "newspaper2", 30, "publisher4")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Printf("--------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())

}