package eventago

import "fmt"

type DefaultCommand struct {
	data []string
}

func (df *DefaultCommand) init(data string) {
	for key, value := range df.data {
		if value == data {
			part := df
			fmt.Sprintln("Property %s is not a valid property on command %s", key, df)
			return
		}
	}
}
