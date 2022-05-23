package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"strconv"
	"strings"
	"github.com/akamensky/argparse"
)

var target_ip string
var ports []string

func main() {
	parse_args()
	start_scanning()
}

func parse_args(){
	parser := argparse.NewParser("goscanner", "Start scanner")
	ip_arg := parser.String("","ip",&argparse.Options{Required: true, Help: "Target IP address"})
	ports_arg := parser.String(
		"", 
		"port", 
		&argparse.Options{Required: true,Help: "Ports to scan, example: 21 / 80,443 / 1-1024"})
	parser.Parse(os.Args)
	if *ip_arg == ""{
		println("Need target IP")
		println(os.Args[0] + " -h for Help")
		os.Exit(0)
	}
	if *ports_arg == ""{
		println("Need port")
		println(os.Args[0] + " -h for Help")
		os.Exit(0)
	}
	target_ip = *ip_arg
	ports = get_ports(*ports_arg)
}

func get_ports(port string,) ([]string,){
	if strings.Contains(port, ","){
		ports := strings.Split(port, ",")
		for _, s := range ports{
			_, err := strconv.Atoi(s)
			if err != nil{
				fmt.Printf("Please enter a valid ports")
				os.Exit(1)
			}
		}
		return ports
	} else if strings.Contains(port, "-"){
		ports_raw := strings.Split(port, "-")
		port_min, err := strconv.Atoi(ports_raw[0])
		if err != nil{
			println("Please enter a valid port range")
			os.Exit(1)
		}
		port_max, err := strconv.Atoi(ports_raw[1])
		if err != nil{
			println("Please enter a valid port range")
			os.Exit(1)
		}
		var ports []string
		if port_min > port_max {
			println("Please enter a valid port range")
		}
		for p_min := port_min; p_min <= port_max; p_min++ {
			port_str := strconv.Itoa(p_min)
			ports = append(ports, port_str)
		}
		return ports
	}
	_, err := strconv.Atoi(port)
	if err != nil{
		println("Please enter a valid port")
		os.Exit(1)
	}
	return []string{port}
}

func start_scanning(){
	for _, p:= range ports{
		go scan_port(p)
	}
	time.Sleep(1500 * time.Millisecond)
}

func scan_port(port string,){
	d := net.Dialer{Timeout: 500 * time.Millisecond}
	_, err := d.Dial("tcp", target_ip + ":" + port)
	if err != nil {
		return
	}
	fmt.Printf("[+] open port: %s\n", port)
}