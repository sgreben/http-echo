package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	defaultAddr          = ":8080"
	addrEnvVarName       = "ADDR"
	portEnvVarName       = "PORT"
	quietEnvVarName      = "QUIET"
	httpStatusEnvVarName = "HTTP_STATUS"
)

var (
	addrFlag       = os.Getenv(addrEnvVarName)
	portFlag64, _  = strconv.ParseInt(os.Getenv(portEnvVarName), 10, 64)
	portFlag       = int(portFlag64)
	quietFlag      bool
	httpStatusFlag = http.StatusOK
)

func init() {
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stderr)
	if addrFlag == "" {
		addrFlag = defaultAddr
	}
	flag.StringVar(&addrFlag, "addr", addrFlag, fmt.Sprintf("address to listen on (environment variable %q)", addrEnvVarName))
	flag.StringVar(&addrFlag, "a", addrFlag, "(alias for -addr)")
	flag.IntVar(&portFlag, "port", portFlag, fmt.Sprintf("port to listen on (overrides -addr port) (environment variable %q)", portEnvVarName))
	flag.IntVar(&portFlag, "p", portFlag, "(alias for -port)")
	flag.BoolVar(&quietFlag, "quiet", quietFlag, fmt.Sprintf("disable all log output (environment variable %q)", quietEnvVarName))
	flag.BoolVar(&quietFlag, "q", quietFlag, "(alias for -quiet)")
	flag.IntVar(&httpStatusFlag, "http-status", httpStatusFlag, fmt.Sprintf("use this HTTP status for responses (environment variable %q)", httpStatusEnvVarName))
	flag.Parse()
	if quietFlag {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	addr, err := addr()
	if err != nil {
		log.Fatalf("address/port: %v", err)
	}
	status, err := status()
	if err != nil {
		log.Fatal(err)
	}
	binaryPath, _ := os.Executable()
	if binaryPath == "" {
		binaryPath = "server"
	}
	statusText := http.StatusText(status)
	if statusText == "" {
		statusText = "unknown status code"
	}
	log.Printf("%s listening on %q, responding with HTTP %d (%s)", filepath.Base(binaryPath), addr, status, statusText)
	err = server(addr, status)
	if err != nil {
		log.Fatalf("start server: %v", err)
	}
}

func server(addr string, status int) error {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		err := r.Write(w)
		if err != nil {
			log.Printf("writing response: %v", err)
			return
		}
		if !quietFlag {
			var buf bytes.Buffer
			r.Write(&buf)
			log.Printf("request: %q", buf.String())
		}
	})
	return http.ListenAndServe(addr, mux)
}

func status() (int, error) {
	if httpStatusFlag < 100 || httpStatusFlag > 999 {
		return 0, fmt.Errorf("invalid HTTP status code: %d", httpStatusFlag)
	}
	return httpStatusFlag, nil
}

func addr() (string, error) {
	portSet := portFlag != 0
	addrSet := addrFlag != ""
	switch {
	case portSet && addrSet:
		a, err := net.ResolveTCPAddr("tcp", addrFlag)
		if err != nil {
			return "", err
		}
		a.Port = portFlag
		return a.String(), nil
	case !portSet && addrSet:
		a, err := net.ResolveTCPAddr("tcp", addrFlag)
		if err != nil {
			return "", err
		}
		return a.String(), nil
	case portSet && !addrSet:
		return fmt.Sprintf(":%d", portFlag), nil
	case !portSet && !addrSet:
		fallthrough
	default:
		return defaultAddr, nil
	}
}
