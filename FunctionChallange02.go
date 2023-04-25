package main

import "fmt"

type Virtmach struct {
    ip      string
    hostname string
    diskgb  int
    ram     int
}

func (v *Virtmach) addDisk(disk int) {
    v.diskgb += disk
}

func (v *Virtmach) addRAM(ram int) {
    v.ram += ram
}

func main() {

    // Initializing the values
    vm := Virtmach{
        ip:       "255.255.255.1",
        hostname: "server01",
        diskgb:   200,
        ram:      16,
    }

    // Calling the methods
    vm.addDisk(50)
    vm.addRAM(4)

    fmt.Printf("IP address: %s\n", vm.ip)
    fmt.Printf("Hostname: %s\n", vm.hostname)
    fmt.Printf("Disk space: %dGB\n", vm.diskgb)
    fmt.Printf("RAM: %dGB\n", vm.ram)
}

