package logger;

import (
	"log"
)
type Logger struct{

}
func (logger Logger) GetInstance() *Logger {
	return new(Logger)
}
func (logger *Logger) Log(values ...interface{}){
	log.Print(values...)
}