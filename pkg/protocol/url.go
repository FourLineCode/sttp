package protocol

import "fmt"

type Url struct {
	Host string
	Port uint16
}

func (u Url) String() string {
	return fmt.Sprintf("sttp://%v:%v", u.Host, u.Port)
}

func Parse(url string) (Url, error) {
	return Url{}, nil
}

func Validate(url string) (bool, error) {
	return false, nil
}
