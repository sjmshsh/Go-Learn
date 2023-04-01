package gomock

import (
	"fmt"
	"io"
	"net/http"
)

// GetUser 它依赖了一个server，所以我们需要fake一个server起来
func GetUser(addr string) {
	resp, _ := http.Get(addr)
	defer resp.Body.Close()
	all, _ := io.ReadAll(resp.Body)
	fmt.Println(string(all))
}
