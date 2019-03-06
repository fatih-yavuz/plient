# Plient
Plient is a wrapper for `net/http` package for setting default proxy and headers for every request.
Imagine you have many proxies and you need to use the same `"User-Agent"` http header for every request made from the same proxy. Then plient is right for you.

## Features
**Default proxy:** While creating a plient, you pass proxy string as argument. Every request made on the plient will use the provided proxy.

**Default headers:** While creating a plient, you pass the http headers as argument. Every request made on the plient will have the provided headers.

**Tested:** Both proxy and user agent headers are tested on by http://httpbin.org/ip and http://httpbin.org/user-agent. So its guaranteed that your request is how you want it. You can go head and check out the test files.

## Usage

```
import (
	"github.com/fatihpy/plient"
)
HTTP_PROXY = "http://3.17.154.4:8080"
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
fmt.Println(body)
```
If you don't want to include headers, you can just pass `nil` instead



