package modules

import "net/http"

const USERNAME = "marfs"
const PASSWORD = "12345"

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !basicAuth(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func basicAuth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Unauthorized!"))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("username or password isn't valid!"))
		return false
	}

	return true
}
