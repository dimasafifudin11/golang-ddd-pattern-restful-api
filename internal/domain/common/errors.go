package common

import "errors"

var (
	ErrNotFound     = errors.New("data tidak ditemukan")
	ErrConflict     = errors.New("terjadi konflik data")
	ErrBadRequest   = errors.New("permintaan tidak valid")
	ErrUnauthorized = errors.New("tidak terautentikasi")
)
