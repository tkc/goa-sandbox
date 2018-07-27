// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa sample": Application User Types
//
// Command:
// $ goagen
// --design=github.com/tkc/goa-sandbox/design
// --out=$(GOPATH)/src/github.com/tkc/goa-sandbox
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// accountListPayload user type.
type accountListPayload struct {
	Ids    []string `form:"ids,omitempty" json:"ids,omitempty" yaml:"ids,omitempty" xml:"ids,omitempty"`
	Offset *int     `form:"offset,omitempty" json:"offset,omitempty" yaml:"offset,omitempty" xml:"offset,omitempty"`
}

// Validate validates the accountListPayload type instance.
func (ut *accountListPayload) Validate() (err error) {
	if ut.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "ids"))
	}
	if ut.Offset == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "offset"))
	}
	return
}

// Publicize creates AccountListPayload from accountListPayload
func (ut *accountListPayload) Publicize() *AccountListPayload {
	var pub AccountListPayload
	if ut.Ids != nil {
		pub.Ids = ut.Ids
	}
	if ut.Offset != nil {
		pub.Offset = *ut.Offset
	}
	return &pub
}

// AccountListPayload user type.
type AccountListPayload struct {
	Ids    []string `form:"ids" json:"ids" yaml:"ids" xml:"ids"`
	Offset int      `form:"offset" json:"offset" yaml:"offset" xml:"offset"`
}

// Validate validates the AccountListPayload type instance.
func (ut *AccountListPayload) Validate() (err error) {
	if ut.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "ids"))
	}

	return
}

// accountPayload user type.
type accountPayload struct {
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the accountPayload type instance.
func (ut *accountPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if ok := goa.ValidatePattern(`^[a-zA-Z0-9!-/:-@¥{}±§|[-{-~]+$`, *ut.Password); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.password`, *ut.Password, `^[a-zA-Z0-9!-/:-@¥{}±§|[-{-~]+$`))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, false))
		}
	}
	return
}

// Publicize creates AccountPayload from accountPayload
func (ut *accountPayload) Publicize() *AccountPayload {
	var pub AccountPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// AccountPayload user type.
type AccountPayload struct {
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// password
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
}

// Validate validates the AccountPayload type instance.
func (ut *AccountPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9!-/:-@¥{}±§|[-{-~]+$`, ut.Password); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.password`, ut.Password, `^[a-zA-Z0-9!-/:-@¥{}±§|[-{-~]+$`))
	}
	if utf8.RuneCountInString(ut.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 8, true))
	}
	if utf8.RuneCountInString(ut.Password) > 20 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 20, false))
	}
	return
}
