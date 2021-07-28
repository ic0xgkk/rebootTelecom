package main

import "flag"

func main() {
	var (
		username string
		password string
		host     string
	)

	flag.StringVar(&username, "username", "username", "The username used to login your HN8145V")
	flag.StringVar(&password, "password", "password", "The password used to login your HN8145V")
	flag.StringVar(&host, "host", "localhost", "The host address without port number, such as '192.168.1.1' or 'm.example.com'")
	flag.Parse()

	initRandCount, err := GetRandCount(getRandCountUrl(host))
	if err != nil {
		panic(err)
	}

	cookie, err := LoginAndGetCookie(getLoginUrl(host), username, password, initRandCount)
	if err != nil {
		panic(err)
	}

	patchRandCount, err := GetDevicePageRandCount(getDevicePageUrl(host), cookie)
	if err != nil {
		panic(err)
	}

	err = Patch(getPatchPageUrl(host), patchRandCount, cookie)
	if err != nil {
		panic(err)
	}
}
