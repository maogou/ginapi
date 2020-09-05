package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

//设置日志等级
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return ""
	}
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)

	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

//设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

//设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx

	return ll
}

//设置当前某一层调用栈的信息(程序计数器,文件信息和行号)
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)

	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

//设置当前的整个调用栈的信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))

		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}

func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().Format("2006-01-02 15:04:05")
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).Output(LevelPanic, fmt.Sprintf(format, v...))
}
