package main

import "fmt"

func main() {
		
		dem := 0
		opc := 0
		var n1, n2, rpd int
		
		
		for {
			fmt.Println("Elektu la operacion:")
			fmt.Println("1. Sumo \n2. Subtraho \n3. Multipliko \n4. Divido")
			fmt.Scanf("%d\n", &opc)
			fmt.Println("_________________________________________________")
			fmt.Println("Tajpu numeron: ")
			fmt.Scanf("%d\n", &n1)
			fmt.Println("Tajpu alian numeron: ")
			fmt.Scanf("%d\n", &n2)
			
			if opc == 1 {
				rpd = n1+n2
				fmt.Println("La sumo estas: \n", rpd)
				}else if opc == 2 {
					rpd = n1-n2
					fmt.Println("La subtraho estas: \n", rpd)
					}else if opc == 3 {
						rpd = n1*n2
						fmt.Println("La multipliko estas: \n", rpd)
						}else if opc == 4 {
							rpd = n1/n2
							fmt.Println("La divido estas: \n", rpd)
						}
			fmt.Println("______________________________")
			fmt.Println("Tajpu 1 se vi volas rekalkuli.")
			fmt.Println("______Tajpu 2 por eliri.______")
			fmt.Println("______________________________")
			fmt.Scanf("%d\n", &dem)
			
			if dem == 2 {
				break
				}else if dem != 1 {
					fmt.Println("Ne ebla!")
				}
		}	
}
