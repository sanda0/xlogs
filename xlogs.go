package xlogs

import "log"

func Info(msg ...any) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("💬INFO : ")
	log.Println(msg...)

}

func Warn(msg ...any) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("🚧WARN : ")
	log.Println(msg...)
}

func Debug(msg ...any) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("🛑DEBUG : ")
	log.Println(msg...)
}

func Error(msg ...any) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("❌ERROR : ")
	log.Println(msg...)
}
