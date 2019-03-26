package captcha

import (
	"fmt"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/mojocn/base64Captcha"
	"html/template"
)
const (
	TemplateField  = "captchaImg"
	CurrentCaptcha  = "currentCaptcha"
)
type Captcha struct {
	TemplateField,
	Idkey ,
	Base64String string
}

func New(args ...string) *Captcha{
	templdateField := TemplateField
	if len(args) > 0 && args[0] != ""{
		templdateField = args[0]
	}
	configD := base64Captcha.ConfigDigit{
		Height:     config.DefaultCaptcha.Height,
		Width:      config.DefaultCaptcha.Width,
		MaxSkew:    config.DefaultCaptcha.MaxSkew,
		DotCount:   config.DefaultCaptcha.DotCount,
		CaptchaLen: config.DefaultCaptcha.CaptchaLen,
	}

	//create a digits captcha.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//write to base64 string.
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return &Captcha{
		TemplateField:templdateField,
		Idkey:idKeyD,
		Base64String:base64stringD,
	}
}

func CaptchaImg(args ...string) context.Handler{
	cp := New(args...)
	return cp.Serve
}


func (c *Captcha) Serve(ctx iris.Context){
	ctx.ViewData(c.TemplateField, c.generateHtml())
	ctx.ViewData(c.TemplateField+"Script", c.generateHtmlScript())
	ctx.Next()
}

func Verify(idkey,verifyValue string) bool{
	return  base64Captcha.VerifyCaptcha(idkey, verifyValue)
}

func (c *Captcha) generateHtml() template.HTML{
	fragment := fmt.Sprintf(`<input type="hidden" name="%s" value="%s"/><a href="javascript:void(0);" id="%s"><img src="%s" alt="captcha">刷新验证码</a>`,
		c.TemplateField, c.Idkey, c.TemplateField, c.Base64String)
	return template.HTML(fragment)
}

func (c *Captcha) generateHtmlScript() template.HTML{
	fragment := fmt.Sprintf(`$('#%s').on('click', function(e){
        $.ajax({
            url:"/captcha/%s",
            success:function(data){
                if (data.ret == 0){
                    $('#captchaImg > img').attr('src', data.img);
                }
            }
        });
    })`,
		c.TemplateField, c.Idkey)
	return template.HTML(fragment)
}
