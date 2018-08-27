/*
Copyright 2018 Iguazio Systems Ltd.

Licensed under the Apache License, Version 2.0 (the "License") with
an addition restriction as set forth herein. You may not use this
file except in compliance with the License. You may obtain a copy of
the License at http://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing
permissions and limitations under the License.

In addition, you may not use the software for any purposes that are
illegal under applicable law, and the grant of the foregoing license
under the Apache 2.0 license is conditioned upon your compliance with
such restriction.
*/

package utils

import (
	"fmt"
	"time"

	"github.com/v3io/frames"
)

// AppendValue appends a value to data
func AppendValue(data interface{}, value interface{}) (interface{}, error) {
	switch data.(type) {
	case []int:
		ival, ok := value.(int)
		if !ok {
			return nil, fmt.Errorf("append type mismatch data is %T while value is %T", data, value)
		}
		idata := data.([]int)
		idata = append(idata, ival)
		return idata, nil
	case []float64:
		fval, ok := value.(float64)
		if !ok {
			return nil, fmt.Errorf("append type mismatch data is %T while value is %T", data, value)
		}
		fdata := data.([]float64)
		fdata = append(fdata, fval)
		return fdata, nil
	case []string:
		sval, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("append type mismatch data is %T while value is %T", data, value)
		}
		sdata := data.([]string)
		sdata = append(sdata, sval)
		return sdata, nil
	case []time.Time:
		tval, ok := value.(time.Time)
		if !ok {
			return nil, fmt.Errorf("append type mismatch data is %T while value is %T", data, value)
		}
		tdata := data.([]time.Time)
		tdata = append(tdata, tval)
		return tdata, nil
	}

	return nil, fmt.Errorf("unsupported data type - %T", data)
}

// NewColumn creates a new column from type of value
func NewColumn(value interface{}, size int) (interface{}, error) {
	switch value.(type) {
	case int:
		return make([]int, size), nil
	case float64:
		return make([]float64, size), nil
	case string:
		return make([]string, size), nil
	case time.Time:
		return make([]time.Time, size), nil
	}

	return nil, fmt.Errorf("Unknown type - %T", value)
}

// AppendNil appends an empty value to data
func AppendNil(col frames.Column) error {
	switch col.DType() {
	case frames.IntType:
		return col.Append(0)
	case frames.FloatType:
		return col.Append(0.0)
	case frames.StringType:
		return col.Append("")
	case frames.TimeType:
		return col.Append(time.Unix(0, 0))
	}

	return fmt.Errorf("unsupported data type - %s", col.DType())
}

// ColAt return value at index i in column as interface{}
func ColAt(col frames.Column, i int) (interface{}, error) {
	switch col.DType() {
	case frames.IntType:
		return col.IntAt(i), nil
	case frames.FloatType:
		return col.FloatAt(i), nil
	case frames.StringType:
		return col.StringAt(i), nil
	case frames.TimeType: // TODO: Does v3io support time.Time?
		return col.TimeAt(i), nil
	default:
		return nil, fmt.Errorf("unknown column type - %s", col.DType())
	}
}