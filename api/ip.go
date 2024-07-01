package handler

import (
	"net"
	"github.com/ipinfo/go/v2/ipinfo"
)

func GetCityInfo(ip string) (*ipinfo.Core, error) {
	const token = "d861d3ff432d58"
	
	// params: httpClient, cache, token. `http.DefaultClient` and no cache will be used in case of `nil`.
	client := ipinfo.NewClient(nil, nil, token)


	info, err := client.GetIPInfo(net.ParseIP(ip))

	if err != nil {
		return nil, err
	}

	return info, nil

}