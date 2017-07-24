package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "strconv"
)

// declare all sources of data retrieval

const (
    geoipdb = "http://download.maxmind.com/download/geoip/database/asnum/GeoIPASNum2.zip"
    cidrdb = "http://www.cidr-report.org/as2.0/autnums.html"
)

type Opensrc struct {
    ASN string
    Data map[string] string
}
// Argument Parsing , others to be done

func main() {

    OrganizationName: = os.Args[1]
    fmt.Println(OrganizationName)
    CheckSources()
}

// argument handling and parsing 

func Argparse() {


}

//To check sources switch case will be good

func CheckSources() {
    //Check the files in tmp and place them if they are not there :)
    if _, err: = os.Stat("/tmp/GeoIPASNum2.zip");
    os.IsNotExist(err) {
        command: = "wget http://download.maxmind.com/download/geoip/database/asnum/GeoIPASNum2.zip -O /tmp/GeoIPASNum2.zip"
        cmd: = exec.Command("sh", "-c", command)
        err: = cmd.Run()
        fmt.Println(err)
        fmt.Println("File Dumped in TMP Master")
    } else {

        fmt.Println("All Good")
    }

}

func ParseCidrReport(Orgname string) int {
    Orgname = strconv.Quote(Orgname)
        //Parse data from HTML Junk and Fetch ASn's
    return 0
}

//RCE possible, but lets not do it now

func ParseGeoipdb(Orgname string) int {

    Orgname = strconv.Quote(Orgname)
    command: = "curl -s " + geoipdb + "| gunzip | cut -d\",\" -f3 | sed 's/\"//g' | sort -u | grep -i " + Orgname + ""
    cmd: = exec.Command("sh", "-c", command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err: = cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return 0
}
