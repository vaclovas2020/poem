/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package install

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
