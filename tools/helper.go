package tools

import "fmt"

func ParseAddress(host, port string) string {
	return fmt.Sprintf("http://%s:%s", host, port)
}
