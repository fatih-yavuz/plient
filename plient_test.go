package plient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
)

var HTTP_PROXY = "http://3.17.154.4:8080"
var HTTPS_PROXY = "http://3.17.154.4:8080"

func TestHTTPProxy(t *testing.T) {
	type IpResponse struct {
		Origin string
	}
	var ipRes IpResponse

	// Send a request to server.
	client := create(HTTP_PROXY, nil)
	resp, err := client.Get("http://httpbin.org/ip")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Httpbin returned " + string(resp.StatusCode))
	}
	fmt.Println("Httpbin status code:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &ipRes)
	if err != nil {
		panic(err.Error())
	}

	ips := strings.Split(ipRes.Origin, ",")
	parts, _ := url.ParseRequestURI(HTTP_PROXY)
	ip := parts.Hostname()
	if strings.TrimSpace(ips[len(ips)-1]) != ip {
		msg := fmt.Sprintf("Http proxy is not being used. Expected %s Got %s\n", ip, strings.TrimSpace(ips[len(ips)-1]))
		panic(msg)
	}

}

func TestHTTPSProxy(t *testing.T) {
	type IpResponse struct {
		Origin string
	}
	var ipRes IpResponse

	client := create(HTTPS_PROXY, nil)
	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Httpbin returned " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &ipRes)
	if err != nil {
		panic(err.Error())
	}

	ips := strings.Split(ipRes.Origin, ",")
	parts, _ := url.ParseRequestURI(HTTPS_PROXY)
	ip := parts.Hostname()
	if strings.TrimSpace(ips[len(ips)-1]) != ip {
		msg := fmt.Sprintf("Https proxy is not being used. Expected %s Got %s\n", ip, ips[len(ips)-1])
		panic(msg)
	}
}

func TestUserAgent(t *testing.T) {

	type AgentResponse struct {
		UserAgent string `json:"user-agent"`
	}

	var agentResp AgentResponse

	plient := create(HTTP_PROXY, []Header{{
		key:   "User-Agent",
		value: "Custom Agent",
	}})

	resp, err := plient.Get("http://httpbin.org/user-agent")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &agentResp)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(agentResp.UserAgent)

	if agentResp.UserAgent != "Custom Agent" {
		t.Error("Expected Custom Agent got", agentResp.UserAgent)
	}
}
