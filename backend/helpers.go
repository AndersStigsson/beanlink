package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"beanlink/protos"

	"google.golang.org/protobuf/proto"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GetBean(text string) *protos.BeanProto {
	bean := &protos.BeanProto{}
	dec, _ := base64.StdEncoding.DecodeString(text)
	err := proto.Unmarshal(dec, bean)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return bean
}

func ParseUrlToGetBeanInfo(text string) string {
	u, err := url.Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	queryParams := u.Query()

	shareUserBeanZero := strings.ReplaceAll(queryParams.Get("shareUserBean0"), " ", "+")
	shareBeans := shareUserBeanZero
	if shareUserBeanOne := strings.ReplaceAll(queryParams.Get("shareUserBean1"), " ", "+"); len(shareUserBeanOne) > 0 {
		shareBeans = fmt.Sprintf("%v%v", shareBeans, shareUserBeanOne)
	}
	if shareUserBeanTwo := strings.ReplaceAll(queryParams.Get("shareUserBean2"), " ", "+"); len(shareUserBeanTwo) > 0 {
		shareBeans = fmt.Sprintf("%v%v", shareBeans, shareUserBeanTwo)
	}
	return shareBeans
}
