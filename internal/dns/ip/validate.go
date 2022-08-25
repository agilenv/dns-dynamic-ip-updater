package ip

import (
	"fmt"
	"regexp"
)

func validate(ip string) error {
	r, err := regexp.Compile("\\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){4}\\b")
	if err != nil {
		return err
	}
	if r.Match([]byte(ip)) {
		return nil
	}
	return fmt.Errorf("ip %s seems to be invalid", ip)
}
