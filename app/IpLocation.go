package app

import (
	"github.com/kayon/iploc"
	"log"
	"regexp"
)
var loc *iploc.Locator
func init(){
	path := StoragePath("/qqwry.dat")
	locs, err := iploc.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	loc = locs
}


func GetAddressByIP(ip string) string{
	if ip == ""{
		return "未知"
	}
	matched, _ := regexp.MatchString("(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}", ip)
	if !matched{
		return "未知!"
	}
	detail := loc.Find(ip)
	return detail.String()
}
