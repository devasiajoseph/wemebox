package website

import (
	"net/http"
	"strings"
)

func GetDomainName(r *http.Request) string {
	host := r.Host
	host = strings.Split(host, ":")[0]
	host = strings.Replace(host, "www.", "", -1)
	return host
}
