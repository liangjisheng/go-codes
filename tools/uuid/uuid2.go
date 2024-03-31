package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/sony/sonyflake"

	uuid "github.com/satori/go.uuid"
)

func getUUID() {
	// result of this command is unique.
	strCommd := "ls /dev/disk/by-uuid"
	cmd := exec.Command("sh", "-c", strCommd)

	if byteBuffer, err := cmd.Output(); err != nil {
		log.Printf("GetUUID Error in get uuid, err=%v", err)
	} else {
		strBuffer := string(byteBuffer)
		fmt.Println(strBuffer)
		arr := strings.Split(strBuffer, "\n")
		if len(arr) >= 0 {
			fmt.Println(arr[0])
		} else {
			log.Printf("GetUUID Error in get uuid, split len is 0")
		}
	}

	return
}

func getUUID1() string {
	res := uuid.NewV4().String()
	log.Println(res)
	return res
}

func getUUID2() {
	// Creating UUID Version 4
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Printf("UUID: %s\n", u1.String())

	// or err handling
	u2 := uuid.NewV4()
	fmt.Printf("UUID: %s\n", u2.String())

	//u2, err := uuid.FromString("9f2e0f4c-5b02-4dc2-ab89-74ea8e53dbec")
	//if err != nil {
	//	fmt.Printf("Something went wrong: %s", err)
	//	return
	//}
	//fmt.Printf("Successfully parsed: %s\n", u2)
}

func uuidgen() {
	// run in linux
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:          %s\n", id)
}

func genShortUUID() {
	id := shortuuid.New()
	fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
}

func genUUID() {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())
}

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:              %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:      %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
}
