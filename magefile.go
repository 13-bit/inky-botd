//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	if err := sh.Run("go", "build", "-o", "build/", "./cmd/inky-botd"); err != nil {
		return err
	}

	return nil
}

func Install() error {
	if err := sh.Run("go", "install", "./cmd/bulletinbird"); err != nil {
		return err
	}

	return nil
}
