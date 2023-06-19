package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Realiza aquí la verificación de autenticación
		// Por ejemplo, verifica si el usuario está autenticado o si se proporciona un token de acceso válido

		// Si la autenticación es exitosa, llama al siguiente manejador
		next(w, r)

		// Si la autenticación falla, puedes enviar una respuesta de error o redirigir al usuario a otra página
	}
}
