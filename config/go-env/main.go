package main

import (
	"github.com/subosito/gotenv"
	"log"
	"os"
	"strings"
)

func init() {
	//gotenv.Load()
	//the first value set for a variable will win.
	//gotenv.Load(".env.production", "credentials")
}

func main() {
	log.Println(os.Getenv("APP_ID"))     // "1234567"
	log.Println(os.Getenv("APP_SECRET")) // "abcdef"

	//Both gotenv.Load and gotenv.Apply DO NOT overrides existing environment variables
	gotenv.Apply(strings.NewReader("APP_ID=123456789"))
	log.Println(os.Getenv("APP_ID"))
	// Output: "123456789"

	os.Setenv("HELLO", "world")

	// NOTE: using Apply existing value will be reserved
	gotenv.Apply(strings.NewReader("HELLO=universe"))
	log.Println(os.Getenv("HELLO"))
	// Output: "world"

	// NOTE: using OverApply existing value will be overridden
	gotenv.OverApply(strings.NewReader("HELLO=universe"))
	log.Println(os.Getenv("HELLO"))
	// Output: "universe"

	err := gotenv.Load(".env-is-not-exist")
	log.Println("error", err)
	// error: open .env-is-not-exist: no such file or directory

	//gotenv.Must(gotenv.Load, ".env-is-not-exist")
	// it will throw a panic
	// panic: open .env-is-not-exist: no such file or directory

	// import "strings"

	pairs := gotenv.Parse(strings.NewReader("FOO=test\nBAR=$FOO"))
	log.Println(pairs)
	// gotenv.Env{"FOO": "test", "BAR": "test"}

	pairs, err = gotenv.StrictParse(strings.NewReader(`FOO="bar"`))
	log.Println(err)
	log.Println(pairs)
	// gotenv.Env{"FOO": "bar"}
}
