package prompts

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func OTP() (string, error) {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("invalid otp code")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter your OTP",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", errors.New("otp prompt failed")
	}
	return result, nil
}

func Username() (string, error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("nvalid username")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", errors.New("username prompt failed")
	}
	return result, nil
}

func Password() (string, error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("invalid password")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	if err != nil {
		return "", errors.New("password prompt failed")
	}
	return result, nil
}
