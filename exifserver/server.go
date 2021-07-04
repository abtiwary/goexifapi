package exifserver

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rwcarlsen/goexif/exif"
)

type Server struct {
	Router *chi.Mux
	IP     string
	Port   int
}

func NewServer(ip string, port int) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return &Server{
		Router: r,
		IP:     ip,
		Port:   port,
	}
}

func (s *Server) initAPI() {
	s.Router.Post("/api/v1/exif", s.ExifFromImage)
}

func (s *Server) StartHTTPServer() *http.Server {
	s.initAPI()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", s.IP, s.Port),
		Handler: s.Router,
	}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("ListenAndServe(): %v", err)
	}

	return srv
}

func (s *Server) ExifFromImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 * 1024 * 1024)
	if err != nil {
		fmt.Println("error parsing multipart form")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var buf bytes.Buffer
	memFileWriter := bufio.NewWriter(&buf)
	memFile := bufio.NewReader(&buf)
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("error getting file from form")
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer file.Close()
	fname := strings.Split(header.Filename, ".")
	fmt.Println("file name = %v", fname)
	_, err = io.Copy(memFileWriter, file)
	if err != nil {
		fmt.Println("error copying file")
		w.WriteHeader(http.StatusInternalServerError)
	}

	exifDat := NewExifData()
	x, err := exif.Decode(memFile)
	if err != nil {
		fmt.Println("error decoding file")
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = x.Walk(exifDat)

	for k, v := range exifDat.Data {
		fmt.Printf("%v -> %v", k, v)
	}

	j, err := json.Marshal(exifDat.Data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
