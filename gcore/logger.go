package gcore

// GenericLogger represents common interface for different loggers.
type GenericLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Println(v ...interface{})
}

// VoidLogger represents nil implementation of GenericLogger.
type VoidLogger struct{}

func (l *VoidLogger) Debug(args ...interface{})                 {}
func (l *VoidLogger) Debugf(format string, args ...interface{}) {}
func (l *VoidLogger) Info(args ...interface{})                  {}
func (l *VoidLogger) Infof(format string, args ...interface{})  {}
func (l *VoidLogger) Warn(args ...interface{})                  {}
func (l *VoidLogger) Warnf(format string, args ...interface{})  {}
func (l *VoidLogger) Error(args ...interface{})                 {}
func (l *VoidLogger) Errorf(format string, args ...interface{}) {}
func (l VoidLogger) Println(v ...interface{})                   {}

// SelectLogger returns one logger from given.
func SelectLogger(logger ...GenericLogger) GenericLogger {
	var log GenericLogger
	switch len(logger) {
	case 0:
		// No logger given, use VoidLogger
		log = &VoidLogger{}
	case 1:
		// We got logger to use
		log = logger[0]
	default:
		panic("Only one logger is supported")
	}
	return log
}
