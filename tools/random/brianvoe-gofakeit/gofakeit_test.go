package brianvoe_gofakeit

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func TestDemo(t *testing.T) {
	t.Log(gofakeit.Name())                // Markus Moen
	t.Log(gofakeit.Email())               // alaynawuckert@kozey.biz
	t.Log(gofakeit.Phone())               // (570)245-7485
	t.Log(gofakeit.BS())                  // front-end
	t.Log(gofakeit.BeerName())            // Duvel
	t.Log(gofakeit.Color())               // MediumOrchid
	t.Log(gofakeit.Company())             // Moen, Pagac and Wuckert
	t.Log(gofakeit.CreditCardNumber(nil)) // 4287271570245748
	t.Log(gofakeit.HackerPhrase())        // Connecting the array won't do anything, we need to generate the haptic COM driver!
	t.Log(gofakeit.JobTitle())            // Director
	t.Log(gofakeit.CurrencyShort())       // USD
}

func TestDemoFile(t *testing.T) {
	//Passing nil to CSV, JSON or XML will auto generate data using default values.

	t.Log(gofakeit.CSV(nil))
	t.Log(gofakeit.JSON(nil))
	t.Log(gofakeit.XML(nil))
	t.Log(gofakeit.FileExtension())
	t.Log(gofakeit.FileMimeType())
}

func TestDemoInternet(t *testing.T) {
	t.Log(gofakeit.URL())
	t.Log(gofakeit.DomainName())
	t.Log(gofakeit.DomainSuffix())
	t.Log(gofakeit.IPv4Address())
	t.Log(gofakeit.IPv6Address())
	t.Log(gofakeit.MacAddress())
	t.Log(gofakeit.HTTPStatusCode())
	t.Log(gofakeit.HTTPStatusCodeSimple())
	t.Log(gofakeit.LogLevel(""))
	t.Log(gofakeit.HTTPMethod())
	t.Log(gofakeit.HTTPVersion())
	t.Log(gofakeit.UserAgent())
	t.Log(gofakeit.ChromeUserAgent())
	t.Log(gofakeit.FirefoxUserAgent())
	t.Log(gofakeit.OperaUserAgent())
	t.Log(gofakeit.SafariUserAgent())
}

func TestDemoNumber(t *testing.T) {
	t.Log(gofakeit.Number(0, 10))
	t.Log(gofakeit.Int8())
	t.Log(gofakeit.Int16())
	t.Log(gofakeit.Int32())
	t.Log(gofakeit.Int64())
	t.Log(gofakeit.Uint8())
	t.Log(gofakeit.Uint16())
	t.Log(gofakeit.Uint32())
	t.Log(gofakeit.Uint64())
	t.Log(gofakeit.Float32())
	t.Log(gofakeit.Float32Range(0, 1))
	t.Log(gofakeit.Float64())
	t.Log(gofakeit.Float64Range(0, 1))

	s := []int{1, 2, 3}
	gofakeit.ShuffleInts(s)
	t.Log(s)
	t.Log(gofakeit.RandomInt(s))

	t.Log(gofakeit.HexUint8())
	t.Log(gofakeit.HexUint16())
	t.Log(gofakeit.HexUint32())
	t.Log(gofakeit.HexUint64())
	t.Log(gofakeit.HexUint128())
	t.Log(gofakeit.HexUint256())
}

func TestDemoString(t *testing.T) {
	t.Log(gofakeit.Digit())
	t.Log(gofakeit.DigitN(6))
	t.Log(gofakeit.Letter())
	t.Log(gofakeit.LetterN(6))
	t.Log(gofakeit.Lexify("hello"))
	t.Log(gofakeit.Numerify("123"))

	s := []string{"alice", "bob", "clare"}
	gofakeit.ShuffleStrings(s)
	t.Log(s)
	t.Log(gofakeit.RandomString(s))
}

func TestDemoEmoji(t *testing.T) {
	t.Log(gofakeit.Emoji())
	t.Log(gofakeit.EmojiDescription())
	t.Log(gofakeit.EmojiCategory())
	t.Log(gofakeit.EmojiAlias())
	t.Log(gofakeit.EmojiTag())
}

func TestDemo1(t *testing.T) {
	// 设置 seed
	//gofakeit.Seed(0) // If 0 will use crypto/rand to generate a number
	// or
	//gofakeit.Seed(8675309) // Set it to whatever number you want

	//切换Random源Gofakeit有多个rand源，默认是math.Rand，并且使用互斥锁实现并发安全。
	// Uses math/rand(Pseudo) with mutex locking
	faker := gofakeit.New(0)
	t.Log(faker.Name())

	// Uses math/rand(Pseudo) with NO mutex locking
	// More performant but not goroutine safe.
	faker = gofakeit.NewUnlocked(0)
	t.Log(faker.Name())

	// Uses crypto/rand(cryptographically secure) with mutex locking
	faker = gofakeit.NewCrypto()
	t.Log(faker.Name())

	// Pass in your own random source
	//faker := gofakeit.NewCustom()

	//全局设置rand 如果你需要全局替换rand源
	faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)
}

// Create structs with random injected data
type Foo struct {
	Str           string
	Int           int
	Pointer       *int
	Name          string         `fake:"{firstname}"`  // Any available function all lowercase
	Sentence      string         `fake:"{sentence:3}"` // Can call with parameters
	RandStr       string         `fake:"{randomstring:[hello,world]}"`
	Number        string         `fake:"{number:1,10}"`       // Comma separated for multiple values
	Regex         string         `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
	Map           map[string]int `fakesize:"2"`
	Array         []string       `fakesize:"2"`
	ArrayRange    []string       `fakesize:"2,6"`
	Bar           Bar
	Skip          *string   `fake:"skip"` // Set to "skip" to not generate data for
	SkipAlt       *string   `fake:"-"`    // Set to "-" to not generate data for
	Created       time.Time // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

type Bar struct {
	Name   string
	Number int
	Float  float32
}

func TestDemoStruct(t *testing.T) {
	// Pass your struct as a pointer
	var f Foo
	gofakeit.Struct(&f)

	fmt.Println(f.Str)              // hrukpttuezptneuvunh
	fmt.Println(f.Int)              // -7825289004089916589
	fmt.Println(*f.Pointer)         // -343806609094473732
	fmt.Println(f.Name)             // fred
	fmt.Println(f.Sentence)         // Record river mind.
	fmt.Println(f.RandStr)          // world
	fmt.Println(f.Number)           // 4
	fmt.Println(f.Regex)            // cbdfc
	fmt.Println(f.Map)              // map[PxLIo:52 lxwnqhqc:846]
	fmt.Println(f.Array)            // cbdfc
	fmt.Printf("%+v", f.Bar)        // {Name:QFpZ Number:-2882647639396178786 Float:1.7636692e+37}
	fmt.Println(f.Skip)             // <nil>
	fmt.Println(f.Created.String()) // 1908-12-07 04:14:25.685339029 +0000 UTC

	// Supported formats
	// int, int8, int16, int32, int64,
	// uint, uint8, uint16, uint32, uint64,
	// float32, float64,
	// bool, string,
	// array, pointers, map
	// time.Time // If setting time you can also set a format tag
	// Nested Struct Fields and Embedded Fields
}

// Custom string that you want to generate your own data for
// or just return a static value
type CustomString string

func (c *CustomString) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomString("my custom string"), nil
}

// Imagine a CustomTime type that is needed to support a custom JSON Marshaller
type CustomTime time.Time

func (c *CustomTime) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomTime(time.Now()), nil
}

func (c *CustomTime) MarshalJSON() ([]byte, error) {
	//...
	return nil, nil
}

func TestCustomFake(t *testing.T) {
	// This is the struct that we cannot modify to add struct tags
	type NotModifiable struct {
		Token    string
		Value    CustomString
		Creation *CustomTime
	}

	var f NotModifiable
	gofakeit.Struct(&f)
	fmt.Printf("%s\n", f.Token)     // yvqqdH
	fmt.Printf("%s\n", f.Value)     // my custom string
	fmt.Printf("%+v\n", f.Creation) // 2023-04-02 23:00:00 +0000 UTC m=+0.000000001
}

func TestCustomFunctions(t *testing.T) {
	//自定义函数 你可以实现自己的生成函数
	//extend the usage of struct tags, generate function, available usages in the gofakeit

	// Simple
	gofakeit.AddFuncLookup("friendname", gofakeit.Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return gofakeit.RandomString([]string{"bill", "bob", "sally"}), nil
		},
	})

	// With Params
	gofakeit.AddFuncLookup("jumbleword", gofakeit.Info{
		Category:    "jumbleword",
		Description: "Take a word and jumble it up",
		Example:     "loredlowlh",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "word", Type: "string", Description: "Word you want to jumble"},
		},
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}

			split := strings.Split(word, "")
			gofakeit.ShuffleStrings(split)
			return strings.Join(split, ""), nil
		},
	})

	type Foo struct {
		FriendName string `fake:"{friendname}"`
		JumbleWord string `fake:"{jumbleword:helloworld}"`
	}

	var f Foo
	gofakeit.Struct(&f)
	fmt.Printf("%s\n", f.FriendName) // bill
	fmt.Printf("%s\n", f.JumbleWord) // loredlowlh
}
