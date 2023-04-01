package mysql

import (
	"reflect"
	"testing"
)

func TestNewSqlHandler(t *testing.T) {
	want := "*mysql.SqlHandler"
	got, err := NewSqlHandler()

	if err != nil {
		t.Errorf("err: %v", err)
	}

	actual := reflect.TypeOf(got).String()
	if want != actual {
		t.Errorf("want: %v, actual: %v", want, actual)
	}

}
