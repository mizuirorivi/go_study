package session

type Cookie interface {
	Get(name string) string
	Set(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	Delete(name string)
}
