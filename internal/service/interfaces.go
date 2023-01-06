package service

type CookieStorage interface {
	SetCookie(uri string, cookie string) error
	GetCookie(uri string) (cookie string, err error)
}
