package command

import "fmt"

var SUCCESS = 1

func Execute() int {
	fmt.Println("This command has now run")
	return SUCCESS
}
