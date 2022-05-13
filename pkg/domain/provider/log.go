package provider

type LogProvider interface {
	Trace(value ...interface{})
	Debug(value ...interface{})
	Info(value ...interface{})
	Warn(value ...interface{})
	Error(value ...interface{})
	Fatal(value ...interface{})
	Panic(value ...interface{})
}
