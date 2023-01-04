package log

type Logger interface {
	Debugf(tmp string, args ...interface{})
	Infof(tmp string, args ...interface{})
	Warnf(tmp string, args ...interface{})
	Errorf(tmp string, args ...interface{})
	Fatalf(tmp string, args ...interface{})
	Panicf(tmp string, args ...interface{})
}
