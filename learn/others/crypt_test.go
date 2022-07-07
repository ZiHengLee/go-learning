package others

import (
	"fmt"
	"testing"
)

var (
	AppKey    = "460fc32bfd626d490592074338f6845b"
	RedisKey  = "7ba33770b8f25216d39fe11b8e487ba7"
	MongoKey  = "bba76d378f2511b8e4821709fe7ba337"
	SensorKey = "e7ba19784456fff6fa03cf121ff3612c"
	DidKey    = "f25a2fc72690b780"
)

func TestDecrypt(t *testing.T) {
	res, _ := Decrypt("crypt-23e1da19c0628c484786a2164658338e17ed2306a2b9c81884124b10", MongoKey)
	fmt.Printf("%s", res)
	//fmt.Println(threeSum(nums))
}
