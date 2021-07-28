package main

import "fmt"

func getPatchPageUrl(host string) string {
	return fmt.Sprintf("http://%s:8080/html/ssmp/devmanage/set.cgi?x=InternetGatewayDevice.X_HW_DEBUG.SMP.DM.ResetBoard&RequestFile=html/ssmp/devmanage/e8cdevicemanormal.asp", host)
}

func getDevicePageUrl(host string) string {
	return fmt.Sprintf("http://%s:8080/html/ssmp/devmanage/e8cdevicemanormal.asp", host)
}

func getLoginUrl(host string) string {
	return fmt.Sprintf("http://%s:8080/login.cgi", host)
}

func getRandCountUrl(host string) string {
	return fmt.Sprintf("http://%s:8080/asp/GetRandCount.asp", host)
}
