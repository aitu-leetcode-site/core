package types

import (
	"database/sql/driver"
	"github.com/aitu-leetcode-site/core/errors"
	httpcodes "github.com/aitu-leetcode-site/core/http/codes"
	"github.com/google/uuid"
)

type UserID uuid.UUID

func (ui UserID) String() string {
	return uuid.UUID(ui).String()
}

func (ui UserID) Value() (driver.Value, error) {
	return uuid.UUID(ui).Value()
}

func (ui *UserID) Scan(src interface{}) error {
	u := uuid.UUID(*ui)
	err := u.Scan(src)
	*ui = UserID(u)
	return err
}

func (ui UserID) UUID() uuid.UUID {
	return uuid.UUID(ui)
}

func NewUserID() UserID {
	outUUID := uuid.New()
	return UserID(outUUID)
}

var ErrEmptyUserID = &errors.ApiError{
	HTTPCode: httpcodes.StatusUnprocessableEntity,
	Message:  "Empty User ID",
}

func (ui UserID) Validate() error {
	if uuid.UUID(ui) == uuid.Nil {
		return ErrEmptyUserID
	}
	return nil
}
