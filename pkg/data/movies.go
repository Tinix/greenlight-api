//
// movies.go
// Copyright (C) 2023 tinix <tinix@debian>
//
// Distributed under terms of the MIT license.
//

package data

import (
	"time"
)

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string
	Version   int32
}
