package logger

type Logger interface {
	Error(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Print(...interface{})
	Panic(...interface{})
}

type logger struct{}

func New() Logger {
	return logger{}
}

func (l logger) Error(args ...interface{}) {
}

func (l logger) Info(args ...interface{}) {

}

func (l logger) Warn(args ...interface{}) {

}

func (l logger) Print(args ...interface{}) {

}

func (l logger) Panic(args ...interface{}) {

}
