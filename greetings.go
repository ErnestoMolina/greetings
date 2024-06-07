package greetings

import (
	"errors"
	"fmt"
)

// Hello devuelve un saludo para la persona especifica.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf("Hola, %v! ¡Bienvenido!", name)
	return message, nil
}
