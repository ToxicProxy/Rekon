package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strconv"
	   )

// declare all sources as constant
const geoipdb =  "http://download.maxmind.com/download/geoip/database/asnum/GeoIPASNum2.zip"

type Opensrc struct {
	ASN string
	Data map[string]string
}
// Argument Parsing , others to be done
func main(){
	
	OrganizationName := os.Args[1]
	fmt.Println(OrganizationName)
	Geoipdb(OrganizationName)
	}

// argument handling and parsing 
func Argparse(){


}
//RCE possible, but lets not do it now
func Geoipdb(Orgname string) int{

Orgname = strconv.Quote(Orgname)
command:= "curl -s "+geoipdb +"| gunzip | cut -d\",\" -f3 | sed 's/\"//g' | sort -u | grep -i "+ Orgname +""
cmd := exec.Command("sh","-c",command) 
cmd.Stdout = os.Stdout  
cmd.Stderr = os.Stderr 
err := cmd.Run()
if err != nil {
    log.Fatal(err)
}
return 0
}