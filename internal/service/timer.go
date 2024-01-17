package service

import "time"

type Timer interface {
	NoSeconds() time.Time
}
