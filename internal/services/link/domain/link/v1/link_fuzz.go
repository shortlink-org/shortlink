//go:build gofuzz

package v1

func Fuzz(link []byte) int {
	data := &Link{Url: string(link)}
	err := NewURL(data)
	if err != nil {
		return -1
	}
	if len(newLink.Hash) != 7 {
		return -1
	}
	return 1
}
