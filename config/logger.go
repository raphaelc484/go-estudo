package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

/***
Esse logger poder ser utilizado em outros projetos,caso seja necessário o uso
de um logger terceiro basta mudar neste arquivo e usar no retorno correspondente

obs. É possível estilizar mais o logger com cores e outros adendos, vale uma
busca depois
***/

func New_logger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) ErrorF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
