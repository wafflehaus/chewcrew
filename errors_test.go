package main

import (
	"errors"
	"fmt"
	"testing"
)

func AssertContextNil(err Error, t *testing.T) {
	if found := err.Context; found != "" {
		t.Error("did not expect context, but found %s", found)
	}
}

func AssertMessage(expect string, err Error, t *testing.T) {
	if found := err.E.Error(); found != expect {
		t.Errorf("expected error message %s, but found %s", expect, found)
	}
}

func AssertContext(expect string, err Error, t *testing.T) {
	if found := err.Context; found != expect {
		t.Errorf("expected context message %s, but found %s", expect, found)
	}
}

func AssertStack(err Error, t *testing.T) {
	if err.Stack == "" {
		t.Errorf("expected stack, but did not find one")
	}
}

func TestNewError(t *testing.T) {
	msg := "test error"
	err := NewError(msg)

	AssertMessage(msg, err, t)
	AssertStack(err, t)
	AssertContextNil(err, t)
}

func TestNewErrorf(t *testing.T) {
	one := "1"
	two := "4"
	format := "%s23%s"
	err := NewErrorf(format, one, two)
	fmtdMessage := fmt.Sprintf(format, one, two)

	AssertMessage(fmtdMessage, err, t)
	AssertStack(err, t)
	AssertContextNil(err, t)
}

func TestNewMaskedError(t *testing.T) {
	msg := "underlying"
	underlying := errors.New(msg)
	err := NewMaskedError(underlying)

	AssertMessage(msg, err, t)
	AssertStack(err, t)
	AssertContextNil(err, t)
}

func TestNewMaskedErrorWithContext(t *testing.T) {
	msg := "test error"
	underlying := errors.New(msg)
	context := "my context"
	err := NewMaskedErrorWithContext(underlying, context)

	AssertMessage(msg, err, t)
	AssertStack(err, t)
	AssertContext(context, err, t)
}

func TestNewMaskedErrorWithContextf(t *testing.T) {
	msg := "test error"
	underlying := errors.New(msg)

	one := "1"
	two := "4"
	format := "%s23%s"
	err := NewMaskedErrorWithContextf(underlying, format, one, two)
	fmtdContext := fmt.Sprintf(format, one, two)

	AssertMessage(msg, err, t)
	AssertStack(err, t)
	AssertContext(fmtdContext, err, t)
}