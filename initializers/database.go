package initializers

import (
	"fmt"
	"os"
)

func ConnectDatabase() {
	fmt.Printf("Connect database: %s", os.Getenv("DATABASE_URL"))
}
