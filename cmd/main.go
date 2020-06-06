package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"

	"github.com/shnupta/sycli"
)

func main() {
	argLen := len(os.Args)
	if argLen == 1 {
		fmt.Printf(HELP_MESSAGE)
		return
	}

	switch os.Args[1] {
	case C_RUN:
		if argLen != 3 {
			fmt.Printf(RUN_HELP_MESSAGE)
		} else {
			run(os.Args[2])
		}
	case C_QUIT:
		quit()
	case C_PAUSE:
		pause()
	case C_RESUME:
		resume()
	case C_VOLUME:
		if argLen != 3 {
			fmt.Printf(VOLUME_HELP_MESSAGE)
		} else {
			volume(os.Args[2])
		}
	default:
		fmt.Printf(HELP_MESSAGE)
	}
}

const (
	HELP_MESSAGE        = "Usage:\n\tsycli <command> [options...]\n\nCommands:\n- run [youtube link]\n"
	RUN_HELP_MESSAGE    = "Usage:\n\tsycli run [youtube link]\n"
	VOLUME_HELP_MESSAGE = "Usage:\n\tsycli volume [0-100]\n"
	C_RUN               = "run"
	C_QUIT              = "quit"
	C_PAUSE             = "pause"
	C_RESUME            = "resume"
	C_VOLUME            = "volume"
)

func connect() net.Conn {
	// addr := net.UnixAddr{Name: "/tmp/mpvsocket", Net: "unix"}
	conn, err := net.Dial("unix", "/tmp/mpvsocket")
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func run(url string) {
	cmd := exec.Command("streamlink", "--player", "mpv --no-video --input-ipc-server=/tmp/mpvsocket", url, "best")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func quit() {
	conn := connect()
	defer conn.Close()
	com := sycli.Command{CommandName: "quit"}
	if _, err := sycli.SendCommand(com, conn); err != nil {
		log.Fatal(err)
	}
}

func pause() {
	conn := connect()
	defer conn.Close()
	com := sycli.SetBoolPropertyCommand("pause", true)
	if _, err := sycli.SendCommand(com, conn); err != nil {
		log.Fatal(err)
	}
}

func resume() {
	conn := connect()
	defer conn.Close()
	com := sycli.SetBoolPropertyCommand("pause", false)
	if _, err := sycli.SendCommand(com, conn); err != nil {
		log.Fatal(err)
	}
}

func volume(val string) {
	conn := connect()
	defer conn.Close()
	level, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	com := sycli.SetIntPropertyCommand("volume", level)
	if _, err := sycli.SendCommand(com, conn); err != nil {
		log.Fatal(err)
	}
}
