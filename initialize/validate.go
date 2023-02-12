package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"pay/global"
)

func InitValidateTran(locale string) error {
	// 修改 validator 引擎属性， 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		//
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni GetTranslatir(%s)", locale)
		}

		switch locale {
		case "en":
			// enTranslations "github.com/go-playground/validator/v10/translations/en"
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			// zhTranslations "github.com/go-playground/validator/v10/translations/zh"
			_ = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
	}

	return nil
}
