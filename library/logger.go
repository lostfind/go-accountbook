package library

type logger interface {
	Error(v ...interface{})
	Info(v ...interface{})
	Fatal(v ...interface{})
}
