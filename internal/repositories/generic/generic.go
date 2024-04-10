package generic

import (
	"log"
	"prea/internal/common"
	"prea/internal/domain"
	"reflect"
	"strings"
)

func GetLogger() *log.Logger {
	return common.MakeLogger("repositories/generic")
}

type IGenericRepository[T domain.IModel] interface {
	GetAll() ([]T, error)
	GetById(id int64) (T, error)
	Create(data T) (T, error)
	Update(id int64, partial T) (T, error)
	Delete(id int64) error
	GetLast() (T, error)
}

func ModelToDest[T any](data *T) []any {
	addrs := reflect.ValueOf(data)

	var vals []any
	elem := addrs.Elem()
	for i := range elem.NumField() {
		vals = append(vals, elem.Field(i).Addr().Interface())
	}

	return vals
}

func ModelToInsert[T domain.IModel](data T) (string, string) {
	refls := reflect.TypeOf(data)
	vals := reflect.ValueOf(data)

	var k = "("
	var v = "("
	var sep = ","

	for i := range refls.NumField() {
		if refls.Field(i).Name != data.Pk() && !vals.Field(i).IsZero() {
			k += strings.ToLower(refls.Field(i).Name) + sep
			v += vals.Field(i).String() + sep
		}
	}

	return strings.TrimSuffix(k, sep) + ")", strings.TrimSuffix(v, sep) + ")"
}

func ModelToUpdate[T domain.IModel](data T) string {
	refls := reflect.TypeOf(data)
	vals := reflect.ValueOf(data)

	var update = ""
	var sep = ","

	for i := range refls.NumField() {
		if refls.Field(i).Name != data.Pk() && !vals.Field(i).IsZero() {
			key := strings.ToLower(refls.Field(i).Name)
			val := vals.Field(i).String()

			if reflect.TypeOf(val).Kind() == reflect.String {
				val = `'` + val + `'`
			}

			update += key + "=" + val + sep
		}
	}

	return strings.TrimSuffix(update, sep)
}
