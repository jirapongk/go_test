package toledo

import (
	"fmt"

	"github.com/jirapongk/go-test/toledo/internal/turboman"
)

func SayHelloToledo() {
	fmt.Println("Hello Toledo")
}

func genTextHello() {
	fmt.Println("Hello Toledo Private Function")
}

func SayHelloToledoPrivate() {
	genTextHello()
	genTextHello2()
	turboman.SayHelloTurboman()
}
