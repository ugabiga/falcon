package service

import (
	"fmt"
	"github.com/ugabiga/falcon/internal/ent"
)

func dbRollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: %v", err, rollbackErr)
	}
	return err
}
