package src

import "log"

func init() {
	initLogging()
}

func initLogging() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {

}
