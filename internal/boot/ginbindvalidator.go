package boot

import (
	"bytes"
	"github.com/gin-gonic/gin/binding"
	localeszh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translationszh "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

/*
gin框架数据验证调整, 增加翻译
*/
type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    *ut.Translator
}

var _ binding.StructValidator = new(defaultValidator)

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		zhs := localeszh.New()
		uni := ut.New(zhs, zhs)
		trans, found := uni.GetTranslator("zh")
		if !found {
			log.Warnf("translator zh not found\n")
		}

		validate := validator.New()
		// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
		// 注册中文翻译
		err := translationszh.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			log.Warnf("register translation err: %v\n", err)
		}
		validate.SetTagName("binding")

		v.validate = validate
		v.trans = &trans
	})
}

func GinBindingValidatorAddTranslator() {
	binding.Validator = new(defaultValidator)
}

func TranslateError(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	va, ok := binding.Validator.(*defaultValidator)
	if !ok {
		return err
	}

	buf := new(bytes.Buffer)
	for _, e := range errs {
		buf.WriteString(e.Translate(*va.trans) + " \n")
	}
	s := buf.String()
	return errors.New(s[:len(s)-2])
}
