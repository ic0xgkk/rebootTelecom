package main

import (
	"fmt"
	"testing"
)

var (
	TestHost     = "192.168.1.1"
	TestUsername = ""
	TestPassword = ""
)

func TestGetRandCount(t *testing.T) {
	rc, err := GetRandCount(getRandCountUrl(TestHost))
	if err != nil {
		panic(err)
	}

	fmt.Println(rc)
}

func TestLoginAndGetCookie(t *testing.T) {
	rc, err := GetRandCount(getRandCountUrl(TestHost))
	if err != nil {
		panic(err)
	}

	c, err := LoginAndGetCookie(getLoginUrl(TestHost), TestUsername, TestPassword, rc)
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Raw)
}

func TestGetDevicePageRandCount(t *testing.T) {
	rc, err := GetRandCount(getRandCountUrl(TestHost))
	if err != nil {
		panic(err)
	}

	c, err := LoginAndGetCookie(getLoginUrl(TestHost), TestUsername, TestPassword, rc)
	if err != nil {
		panic(err)
	}

	dpRc, err := GetDevicePageRandCount(getDevicePageUrl(TestHost), c)
	if err != nil {
		panic(err)
	}

	fmt.Println(dpRc)
}

func TestPatch(t *testing.T) {
	initRandCount, err := GetRandCount(getRandCountUrl(TestHost))
	if err != nil {
		panic(err)
	}

	cookie, err := LoginAndGetCookie(getLoginUrl(TestHost), TestUsername, TestPassword, initRandCount)
	if err != nil {
		panic(err)
	}

	patchRandCount, err := GetDevicePageRandCount(getDevicePageUrl(TestHost), cookie)
	if err != nil {
		panic(err)
	}

	err = Patch(getPatchPageUrl(TestHost), patchRandCount, cookie)
	if err != nil {
		panic(err)
	}
}
