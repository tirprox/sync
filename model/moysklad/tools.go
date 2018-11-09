package moysklad

import "reflect"

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	//index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				//index = i
				exists = true
				return
			}
		}
	}

	return
}

type Filterable interface {
	GetName() string
}

func FilterSlice(filterable []Filterable, allowed []string) (filtered []Filterable) {

	for _, item := range filterable {
		if InArray(item.GetName(), allowed) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
