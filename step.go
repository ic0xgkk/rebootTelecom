package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"regexp"
	"strings"
)

var (
	// filtering non-alphanumeric
	reg1 *regexp.Regexp = nil

	// label match hwonttoken
	reg2 *regexp.Regexp = nil
)

func GetRandCount(url string) (string, error) {
	resp, err := http.Post(url, "text/html", nil)
	if err != nil {
		return "", fmt.Errorf("get randCount http response error: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("get randCount with wrong status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read from readCount response failed: %s", err)
	}

	if reg1 == nil {
		reg1, err = regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			return "", fmt.Errorf("init regex 1 failed: %s", err)
		}
	}

	return reg1.ReplaceAllString(string(body), ""), nil
}

func LoginAndGetCookie(url string, username, password, randCount string) (*http.Cookie, error) {
	resp, err := http.PostForm(url, neturl.Values{
		"UserName":     []string{username},
		"PassWord":     []string{base64.StdEncoding.EncodeToString([]byte(password))},
		"Language":     []string{"chinese"},
		"x.X_HW_Token": []string{randCount},
	})
	if err != nil {
		return nil, fmt.Errorf("post login form failed: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("get login cookie with wrong status code: %d", resp.StatusCode)
	}

	cookies := resp.Cookies()
	if len(cookies) == 0 {
		return nil, fmt.Errorf("login failed: credential may wrong")
	}

	for i, _ := range cookies {
		cookie := cookies[i]

		if cookie.Path == "/" {
			return cookie, nil
		}
	}

	return nil, fmt.Errorf("not found path / cookie")
}

func GetDevicePageRandCount(url string, cookie *http.Cookie) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("new device page http request failed: %s", err)
	}
	req.AddCookie(cookie)

	if reg2 == nil {
		reg2, err = regexp.Compile("name\\=\\\"onttoken\\\"\\sid\\=\\\"hwonttoken\\\"\\svalue=\\\"(?P<rc>.*?)\\\"")
		if err != nil {
			return "", fmt.Errorf("init regex 2 failed: %s", err)
		}
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("dial device manage page failed: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("get device manage page randCount with wrong status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body from device manage page response failed: %s", err)
	}

	matches := reg2.FindStringSubmatch(string(body))
	rcIndex := reg2.SubexpIndex("rc")

	return matches[rcIndex], nil
}

func Patch(url, token string, cookie *http.Cookie) error {
	postData := neturl.Values{
		"x.X_HW_Token": []string{token},
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(postData.Encode()))
	if err != nil {
		return fmt.Errorf("new patch request failed: %s", err)
	}
	req.AddCookie(cookie)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("dial patch page response failed: %s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("patch failed with wrong status code: %d", resp.StatusCode)
	}

	return nil
}
