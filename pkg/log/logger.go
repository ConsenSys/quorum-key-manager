package log

//go:generate mockgen -source=logger.go -destination=mock/logger.go -package=mocks

type Logger interface {
	WithComponent(component string) Logger
	Debug(msg string, keysAndValues ...interface{}) Logger
	Warn(msg string, keysAndValues ...interface{}) Logger
	Info(msg string, keysAndValues ...interface{}) Logger
	Error(msg string, keysAndValues ...interface{}) Logger
	Panic(msg string, keysAndValues ...interface{}) Logger
	Fatal(msg string, keysAndValues ...interface{}) Logger
	WithError(err error) Logger
	With(args ...interface{}) Logger
	Write(p []byte) (n int, err error)
}
