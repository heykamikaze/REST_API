package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

//логирование с инфой о файле, где происходит логирование и фнкц, возвращает функцию и номер строки
//все логи в 1 файле с разными уровнями логирования чтоб это говно не дублировалось

//hook на запись, массив левлов и райтеров. в кафку райтер инфо и дебаг, в файлы трейсы и ошибки, в стандартный вывод варнинги и крит.ошибки
type writerHook struct {
	Writer    []io.Writer
	Loglevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

//метод на левлы логирования
func (hook *writerHook) Levels() []logrus.Level {
	return hook.Loglevels
}

var e *logrus.Entry

//eсли этот логгер отъебнет, в этой структурке заменить логрус на другой какой-нибудь логгер рабочий
type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithFiled(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//красивые логи
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	//dir for logs
	err := os.Mkdir("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, _ := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	//if err != nil {
	//	panic(err)
	//}

	l.SetOutput(io.Discard) //чтобы по умолчанию никуда не писалось

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		Loglevels: logrus.AllLevels, //все левлы трейсов в структурке собраны
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
