package netaddr

import "net/url"

func SplitHostPort(addr, defaultPort string) (host, port string, err error) {
	u, err := url.Parse("//" + addr)
	if err != nil {
		return "", "", err
	}

	host, port = u.Hostname(), u.Port()
	if port == "" {
		port = defaultPort
	}

	return host, port, err
}
