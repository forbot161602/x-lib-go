package preset

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/joho/godotenv"

	"github.com/forbot161602/pbc-golang-lib/source/entry/config"
)

func init() {
	if err := godotenv.Load(); err == nil {
		fmt.Println("[INFO] The .env file has been successfully loaded.")
	}
	config.SetConfig()
	rand.Seed(time.Now().UnixNano())
}
