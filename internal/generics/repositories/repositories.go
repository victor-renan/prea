package repositories

import (
	"log"
	"prea/internal/common"
	"reflect"
	"strings"
	"time"
)

func GetLogger() *log.Logger {
	return common.MakeLogger("repositories")
}

type IModelInjectable interface {
	Table() string
	Pk() string
}

type IGenericRepository[T IModelInjectable] interface {
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

func ModelToInsert[T IModelInjectable](data T) (string, string) {
	refls := reflect.TypeOf(data)
	vals := reflect.ValueOf(data)

	var k = "("
	var v = "("
	var sep = ","

	for i := range refls.NumField() {
		if refls.Field(i).Name != data.Pk() && !vals.Field(i).IsZero() {
			val := vals.Field(i).String()
			if reflect.TypeOf(val).Kind() == reflect.String {
				val = `'` + val + `'`
			}

			k += strings.ToLower(refls.Field(i).Name) + sep
			v += val + sep
		}
	}

	return strings.TrimSuffix(k, sep) + ")", strings.TrimSuffix(v, sep) + ")"
}

func ModelToUpdate[T IModelInjectable](data T) string {
	refls := reflect.TypeOf(data)
	vals := reflect.ValueOf(data)

	var update = ""
	var sep = ","

	for i := range refls.NumField() {
		if refls.Field(i).Name != data.Pk() && !vals.Field(i).IsZero() {
			key := strings.ToLower(refls.Field(i).Name)
			val := vals.Field(i).String()

			if refls.Field(i).Type.Kind() == reflect.String {
				val = `'` + val + `'`
			}

			if _, err := time.Parse(time.DateTime, val); err == nil {
				val = `'` + val + `'`
			}

			update += key + "=" + val + sep
		}
	}

	return strings.TrimSuffix(update, sep)
}
