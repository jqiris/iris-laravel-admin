package config

type CaptchaConfig struct {
	Width   int
	Height  int
	CaptchaLen int
	DotCount int
	MaxSkew float64
}
var DefaultCaptcha *CaptchaConfig

func initCaptcha(){
	DefaultCaptcha = &CaptchaConfig{
		CaptchaLen:  4,
		Width:   120,
		Height:  36,
		DotCount: 90,
		MaxSkew: 10,
	}
}
