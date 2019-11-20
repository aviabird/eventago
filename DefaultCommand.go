package eventago

import "fmt"

//DefaultCommand DefaultCommand
type DefaultCommand struct {
	data []string
}

func (df *DefaultCommand) init(data string) {
	for key, value := range df.data {
		if value == data {
			part := df
			fmt.Printf("Property %d is not a valid property on command %s", key, part.data)
			return
		}
	}
}
