package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especifica.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randonFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	mensages := make(map[string]string)
	for _, name := range names {
		mensage, err := Hello(name)
		if err != nil {
			return nil, err
		}

		mensages[name] = mensage
	}

	return mensages, nil
}

func randonFormat() string {
	format := []string{
		"Hola, %v !Bienvenido¡",
		"!Que bueno verte¡%v",
		"!Saludos, %v !Encantado de conocerte¡",
	}
	return format[rand.Intn(len(format))]
}
