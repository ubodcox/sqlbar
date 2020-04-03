package logs

import (
	"log"
	"os"
	//"fmt"
	
	"time"
	color "gopkg.in/gookit/color.v1" 
	
)

const (
	LOG_LEVEL_NUM    = 6
	LOG_PACKAGE_NUM  = 11
	LOG_MODULE_NUM   = 11
	LOG_FUNCTION_NUM = 13
	LOG_METHOD_NUM   = 11
	LOG_SESSION_NUM  = 11
)

/*
* время
+ уровень
- пакет
- модуль
- функция
+ метод
- сессия
+ сообщение
*/

type (
	// LogWriter struct
	LogWriter struct {
		PackageName  []string
		ModuleName   []string
		FunctionName []string
		SessionName  string
	}
)

// Log writeer
var Log LogWriter

// Error func
/* func Error(err error) {
	if err != nil {
		log.Println(err)
	}
} */

//---------------------------------------------------------

// GetPackageName func
func (l *LogWriter) GetPackageName() string {
	if len(l.PackageName) == 0 {
		return ""
	}
	return l.PackageName[len(l.PackageName)-1]
}

// GetModuleName func
func (l *LogWriter) GetModuleName() string {
	if len(l.ModuleName) == 0 {
		return ""
	}
	return l.ModuleName[len(l.ModuleName)-1]
}

// GetFuncName func
func (l *LogWriter) GetFuncName() string {
	if len(l.FunctionName) == 0 {
		return ""
	}
	return l.FunctionName[len(l.FunctionName)-1]
}

//---------------------------------------------------------

// PushFuncName func
func (l *LogWriter) PushFuncName(pkg, mdl, fnc string) {
	l.PackageName  = append(l.PackageName, pkg)
	l.ModuleName   = append(l.ModuleName, mdl)
	l.FunctionName = append(l.FunctionName, fnc)
}

// PopFuncName func
func (l *LogWriter) PopFuncName() {
	l.PackageName  = l.PackageName[:len(l.PackageName)]
	l.ModuleName   = l.ModuleName[:len(l.ModuleName)]
	l.FunctionName = l.FunctionName[:len(l.FunctionName)]
}

//---------------------------------------------------------

// Log handler
func (l *LogWriter) Log(level string, v ...interface{}) {
	timeStr := time.Now().Format("02.01.2006 15:04:05")+" "
	levelStr := StrToLen(level, LOG_LEVEL_NUM)
	packageStr := StrToLen(l.GetPackageName(), LOG_PACKAGE_NUM)
	moduleStr := StrToLen(l.GetModuleName(), LOG_MODULE_NUM)
	functionStr := StrToLen(l.GetFuncName(), LOG_FUNCTION_NUM)
	sessionStr := StrToLen(l.SessionName, LOG_SESSION_NUM)

	//TODO: add log_level

	if (level == "INFO") {
		color.Success.Println(timeStr, levelStr, packageStr, moduleStr, functionStr, sessionStr, v)
	} else
	if (level == "DEBUG") {
		color.Secondary.Println(timeStr, levelStr, packageStr, moduleStr, functionStr, sessionStr, v)
	} else
	if (level == "ERROR") {
		color.Error.Println(timeStr, levelStr, packageStr, moduleStr, functionStr, sessionStr, v)
	} else
	if (level == "WARNG") {
		color.Warn.Println(timeStr, levelStr, packageStr, moduleStr, functionStr, sessionStr, v)
	}
}

// Info log
func (l *LogWriter) Info(v ...interface{}) {
	l.Log("INFO", v)
}

// Debug log
func (l *LogWriter) Debug(v ...interface{}) {
	l.Log("DEBUG", v)
}

// Warning log
func (l *LogWriter) Warning(v ...interface{}) {
	l.Log("WARNG", v)
}

// Error log
func (l *LogWriter) Error(v ...interface{}) {
	l.Log("ERROR", v)
}

var file *os.File

// Init func
func Init() bool {
	file, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.SetOutput(file)
	return true
}

// Deinit func
func Deinit() {
	file.Close()
}

// StrToLen func
func StrToLen(s string, leng int) string {
	temp := s
	for len(temp) < leng {
		temp = temp + " "
	}
	return temp
}
