package userservice

import (
	"errors"
	"strings"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"
)

func TestRegisterRequestValidate(t *testing.T) {
	t.Run("all required", func(t *testing.T) {
		err := (RegisterRequest{}).Validate()
		fields := fieldErrors(t, err)
		if got := fields["username"].Error(); got != errmsg.ErrorMsgUsernameRequired {
			t.Fatalf("username: %q", got)
		}
		if got := fields["phone_number"].Error(); got != errmsg.ErrorMsgPhoneNumberRequired {
			t.Fatalf("phone_number: %q", got)
		}
		if got := fields["password"].Error(); got != errmsg.ErrorMsgPasswordRequired {
			t.Fatalf("password: %q", got)
		}
	})

	t.Run("length and phone format", func(t *testing.T) {
		err := RegisterRequest{
			Name:        "ab",
			PhoneNumber: "123",
			Password:    "123",
		}.Validate()
		fields := fieldErrors(t, err)
		if got := fields["username"].Error(); got != errmsg.ErrorMsgUsernameLength {
			t.Fatalf("username: %q", got)
		}
		if got := fields["phone_number"].Error(); got != errmsg.ErrorMsgPhoneNumberIsNotValid {
			t.Fatalf("phone_number: %q", got)
		}
		if got := fields["password"].Error(); got != errmsg.ErrorMsgPasswordLength {
			t.Fatalf("password: %q", got)
		}
	})

	t.Run("username too long", func(t *testing.T) {
		err := RegisterRequest{
			Name:        strings.Repeat("a", 51),
			PhoneNumber: "09123456789",
			Password:    "1234",
		}.Validate()
		fields := fieldErrors(t, err)
		if _, ok := fields["phone_number"]; ok {
			t.Fatal("phone_number should be valid")
		}
		if got := fields["username"].Error(); got != errmsg.ErrorMsgUsernameLength {
			t.Fatalf("username: %q", got)
		}
	})

	t.Run("valid", func(t *testing.T) {
		err := RegisterRequest{
			Name:        "abc",
			PhoneNumber: "09123456789",
			Password:    "1234",
		}.Validate()
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestRegisterWrapsFieldErrors(t *testing.T) {
	_, err := Service{}.Register(RegisterRequest{Name: "ab"})
	if richerror.KindOf(err) != richerror.KindInvalid {
		t.Fatalf("kind = %v", richerror.KindOf(err))
	}
	if richerror.MessageOf(err) != errmsg.ErrorMsgInvalidFields {
		t.Fatalf("message = %q", richerror.MessageOf(err))
	}
	fields := fieldErrors(t, err)
	if fields["username"].Error() != errmsg.ErrorMsgUsernameLength {
		t.Fatalf("username: %v", fields["username"])
	}
}

func fieldErrors(t *testing.T, err error) validation.Errors {
	t.Helper()
	var fields validation.Errors
	if !errors.As(err, &fields) {
		t.Fatalf("got %T: %v", err, err)
	}
	return fields
}
