package puppy

import (
	"github.com/BotManu/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() string {
	return "I'm from v1.1.0"
}

func From12() string {
	return "I'm from v1.2.0"
}

func From13() string {
	return "I'm from v1.3.0"
}
