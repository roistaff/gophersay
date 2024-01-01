package main
import(
	"fmt"
	"os"
)

func main(){
	var say string
	if len(os.Args) == 1{
		say = "Hello World!"
	}else if os.Args[1] == "-h"{
		help :="Gophersay\n Use: gophersay 'string'\n Help: gophersay -h"
		fmt.Println(help)
		os.Exit(1)
	}else{
	        say = os.Args[1]
	}
	var under string
	var under2 string
	for i:=1; i<= len(say); i++{
		under += "-"
		under2 += "_"
	}
	var AA string = `
 _`+under2+`_
< `+say+` >
 -`+under+`-
        \   
         \  ʕ◔π◔ʔ
           -|   |-
            /---\
`
	fmt.Println(AA)
}
