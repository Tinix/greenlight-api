//
// runtime.go
// Copyright (C) 2023 tinix <tinix@debian>
//
// Distributed under terms of the MIT license.
//

package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil

}
