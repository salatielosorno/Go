package users

import (
	"fmt"
	"time"

	"github.com/salatielosorno/Go/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "John", time.Now(), true)
	fmt.Println(u)
}
