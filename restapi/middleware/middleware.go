package middleware

import (
	"mime/multipart"
	"net/http"
)

// hacky way to get swagger generated code to work with slack-provided application/x-www-form-urlencoded content type
type FormContentTypeConversionMiddleware struct {
	Next http.Handler
}

func (m *FormContentTypeConversionMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		err := r.ParseForm()

		if err != nil {
			panic(err)
		}

		r.MultipartForm = &multipart.Form{
			Value: r.Form,
		}
	}

	m.Next.ServeHTTP(rw, r)
}
