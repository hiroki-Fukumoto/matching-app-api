package validator

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/hiroki-Fukumoto/matching-app-api/api/config"
	"github.com/hiroki-Fukumoto/matching-app-api/api/enum"
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"gopkg.in/go-playground/validator.v9"
	ja_translations "gopkg.in/go-playground/validator.v9/translations/ja"
	"gorm.io/gorm"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

const (
	emailExists = "email_exists"
	dateFormat  = "date_format"
	sex         = "sex"
	prefecture  = "prefecture"
)

func Init() {
	ja := ja.New()
	uni = ut.New(ja, ja)
	t, _ := uni.GetTranslator("ja")
	trans = t
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		fieldName := fld.Tag.Get("ja")
		if fieldName == "-" {
			return ""
		}
		return fieldName
	})
	ja_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterValidation(emailExists, emailExistsValidator)
	validate.RegisterValidation(dateFormat, dateFormatValidator)
	validate.RegisterValidation(sex, sexValidator)
	validate.RegisterValidation(prefecture, prefectureValidator)
}

func Validate(i interface{}) error {
	Init()
	return validate.Struct(i)
}

type ValidateError struct {
}

func GetErrorMessages(err error) []string {
	if err == nil {
		return []string{}
	}
	var messages []string
	for _, m := range err.(validator.ValidationErrors) {
		switch m.ActualTag() {
		case emailExists:
			message := "指定されたメールアドレスは既に使用されています"
			messages = append(messages, message)
		case dateFormat:
			message := fmt.Sprintf("%sはyyyy-mm-ddの形式にしてください", m.Field())
			messages = append(messages, message)
		case sex:
			message := "性別が正しくありません"
			messages = append(messages, message)
		case prefecture:
			message := "都道府県が正しくありません"
			messages = append(messages, message)
		default:
			messages = append(messages, m.Translate(trans))
		}

	}
	return messages
}

// メールアドレスが存在しないかチェック
func emailExistsValidator(field validator.FieldLevel) bool {
	value := field.Field().String()

	if value == "" {
		return true
	}

	db := config.Connect()
	var user model.User
	err := db.Where("email = ?", value).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

// 性別チェック
func sexValidator(field validator.FieldLevel) bool {
	value := field.Field().String()

	if value == "" {
		return true
	}

	var sex []string
	sexm, _ := util.StructToJsonTagMap(enum.SEX)
	for _, s := range sexm {
		sex = append(sex, s.(string))
	}

	if util.Contains(sex, value) {
		return true
	}

	return false
}

// 都道府県チェック
func prefectureValidator(field validator.FieldLevel) bool {
	value := field.Field().Int()

	prefs := util.GetPrefectureCodeList()

	if util.Contains(prefs, value) {
		return true
	}

	return false
}

// 日付がyyyy-mm-ddのフォーマットになっているかチェック
func dateFormatValidator(field validator.FieldLevel) bool {
	value := field.Field().String()

	if value == "" {
		return true
	}

	_, err := time.Parse("2006-01-02", value)

	return err == nil
}
