package zhegengdi

import (
	"fmt"
	"net/url"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured: %s\r\n", err)
		os.Exit(1)
	}
}

// GetQueryMap get the field RawQuery of net/url URL struct,
//
// scheme://[userinfo@]netloc/path[;parameter][?query][#fragment]
//
// URL.RawQuery is a string type without '?'.
func GetQueryMap(rawurl string) (cond url.Values, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return url.ParseQuery(u.RawQuery)
}
