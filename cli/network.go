package cli

import "net/url"

const defaultPort = "1338"

func splitHostPort(addr string) (hostname, port string, err error) {
	u, err := url.Parse("//" + addr)
	if err != nil {
		return
	}

	hostname = u.Hostname()
	port = u.Port()
	if port == "" {
		port = defaultPort
	}

	return
}
