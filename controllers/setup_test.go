package controllers_test

import (
	"jessie_miniproject/config"
	"os"
	"testing"
)

// membuat func main untuk testing
func TestMain(m *testing.M) {
	config.InitDB()
	os.Exit(m.Run()) // jika tesnya gagal, maka langsung exit
}

// pendekatan saat ini blm bisa menggunakan mockdb jadi pake dbasli