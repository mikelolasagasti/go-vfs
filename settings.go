package xdg

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	settingsCmdName = "xdg-settings"

	// DefaultURLSchemeHandlerProperty is the default URL scheme handler property.
	DefaultURLSchemeHandlerProperty = "default-url-scheme-handler"

	// DefaultWebBrowserProperty is the default web browser property.
	DefaultWebBrowserProperty = "default-web-browser"
)

// CheckSetting checks that value of property.subProperty is value. See
// https://portland.freedesktop.org/doc/xdg-settings.html.
func CheckSetting(property, subProperty, value string) (bool, error) {
	args := []string{"check", property}
	if subProperty != "" {
		args = append(args, subProperty)
	}
	args = append(args, value)
	output, err := exec.Command(settingsCmdName, args...).Output()
	if err != nil {
		return false, err
	}
	s := strings.TrimSpace(string(output))
	switch s {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	default:
		return false, fmt.Errorf(`CheckSetting(%q, %q, %q): expected "yes" or "no", got %q`, property, subProperty, value, s)
	}
}

// SetSetting sets property.subProperty to value.
func SetSetting(property, subProperty, value string) error {
	args := []string{"set", property}
	if subProperty != "" {
		args = append(args, subProperty)
	}
	args = append(args, value)
	return exec.Command(settingsCmdName, args...).Run()
}

// Setting gets the value of property.subProperty.
func Setting(property, subProperty string) (string, error) {
	args := []string{"get", property}
	if subProperty != "" {
		args = append(args, subProperty)
	}
	output, err := exec.Command(settingsCmdName, args...).Output()
	return strings.TrimSpace(string(output)), err
}