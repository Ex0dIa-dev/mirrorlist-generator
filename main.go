package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Ex0dIa-dev/mirrorlist-generator/helpers"
	"github.com/Ex0dIa-dev/mirrorlist-generator/helpers/colors"
	"github.com/PuerkitoBio/goquery"
)

func init() {

	flag.BoolVar(&interactiveModeBool, "i", false, "generate your mirrorlist in interactive mode")

	flag.BoolVar(&listCountries, "lc", false, "list all available countries")
	flag.BoolVar(&listProtocols, "lp", false, "list all available protocols")
	flag.BoolVar(&listIPVersions, "li", false, "list all available ip versions")

	flag.StringVar(&flagCountries, "countries", "", "enter countries acronym (ordered, and comma-separated)")
	flag.StringVar(&flagProtocols, "protocols", "", "enter protocols (ordered, and comma-separated)")
	flag.StringVar(&flagIPVersions, "ipversions", "", "enter ipversions (ordered, and comma-separated)")

}

// type arrayFlags []string

// func (i *arrayFlags) String() string {
// 	return "my string representation"
// }

// func (i *arrayFlags) Set(value string) error {
// 	*i = append(*i, value)
// 	return nil
// }

var (
	interactiveModeBool, listCountries, listProtocols, listIPVersions bool
	flagCountries, flagProtocols, flagIPVersions                      string
)

const errorMsg = "There was some error! Please retry, and enter all needed data."

func main() {

	flag.Parse()

	if interactiveModeBool {
		interactiveMode()
		os.Exit(0)
	}

	if listCountries {
		printAvailableCountries()
		os.Exit(0)
	}
	if listProtocols {
		printAvailableProtocols()
		os.Exit(0)
	}
	if listIPVersions {
		printAvailableIPVersions()
		os.Exit(0)
	}

	if len(flagCountries) == 0 || len(flagProtocols) == 0 || len(flagIPVersions) == 0 {
		fmt.Printf("%s%s%s\n", colors.Red, errorMsg, colors.Reset)
		os.Exit(1)
	}

	generateMirrorlist(helpers.ReturnAsArrays(flagCountries), helpers.ReturnAsArrays(flagProtocols), helpers.ReturnAsArrays(flagIPVersions))
}

// interactiveMode is a user-friendly method to generate a mirrorlist
func interactiveMode() {

	reader := bufio.NewReader(os.Stdin)

	//getting the countries
	printAvailableCountries()
	fmt.Print("\nEnter Countries (ordered and comma-separated) -> ")
	inputed_countries, err := reader.ReadString('\n')
	helpers.CheckErr(err)
	inputed_countries = strings.Replace(inputed_countries, "\n", "", -1)

	countries := strings.Split(inputed_countries, ",")
	var protocols []string

	fmt.Printf("\n%sPROTOCOLS%s\n", colors.Blue, colors.Reset)

	//getting the protocols
	fmt.Print("HTTP  -> [y/n] ")
	http, err := reader.ReadString('\n')
	helpers.CheckErr(err)
	http = strings.Replace(http, "\n", "", -1)
	if http == "Y" || http == "y" || http == "YES" || http == "yes" {
		protocols = append(protocols, "http")
	}

	fmt.Print("HTTPS -> [y/n] ")
	https, err := reader.ReadString('\n')
	helpers.CheckErr(err)
	https = strings.Replace(https, "\n", "", -1)
	if https == "Y" || https == "y" || https == "YES" || https == "yes" {
		protocols = append(protocols, "https")
	}

	fmt.Printf("%sIP-VERSION%s\n", colors.Blue, colors.Reset)

	var ipversions []string
	//getting the ip version
	fmt.Print("IPv4  -> [y/n] ")
	ipv4, err := reader.ReadString('\n')
	helpers.CheckErr(err)
	ipv4 = strings.Replace(ipv4, "\n", "", -1)
	if ipv4 == "Y" || ipv4 == "y" || ipv4 == "YES" || ipv4 == "yes" {
		ipversions = append(ipversions, "4")
	}

	fmt.Print("IPv6  -> [y/n] ")
	ipv6, err := reader.ReadString('\n')
	helpers.CheckErr(err)
	ipv6 = strings.Replace(ipv6, "\n", "", -1)
	if ipv6 == "Y" || ipv6 == "y" || ipv6 == "YES" || ipv6 == "yes" {
		ipversions = append(ipversions, "6")
	}

	generateMirrorlist(countries, protocols, ipversions)

}

// generateMirrorlist make a GET Request to "https://archlinux.org/mirrorlist/?" to get the mirrorlist
func generateMirrorlist(countries, protocols, ipversions []string) {

	base_url := "https://archlinux.org/mirrorlist/?"
	url := base_url
	for _, c := range countries {
		url += fmt.Sprintf("country=%s&", c)
	}

	for _, p := range protocols {
		url += fmt.Sprintf("protocol=%s&", p)
	}

	for _, i := range ipversions {
		url += fmt.Sprintf("ip_version=%s&", i)
	}

	url = strings.TrimRight(url, "&")
	if url == base_url {
		fmt.Printf("%s%s%s\n", colors.Red, errorMsg, colors.Reset)
		os.Exit(1)
	}
	fmt.Println(url)

	resp, err := http.Get(url)
	helpers.CheckErr(err)
	defer resp.Body.Close()

	fd, err := os.Create("./mirrorlist")
	helpers.CheckErr(err)
	defer fd.Close()

	io.Copy(fd, resp.Body)

	fmt.Printf("\n%sMirrorlist Successfully Generated!%s\n", colors.Green, colors.Reset)
}

// printAvailableCountries prints all available countries
func printAvailableCountries() {

	resp, err := http.Get("https://archlinux.org/mirrorlist/")
	helpers.CheckErr(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	helpers.CheckErr(err)

	doc.Find("#id_country").Each(func(i int, s *goquery.Selection) {
		s.Find("option").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("[%s%s%s] %s\n",
				colors.Green,
				s.AttrOr("value", ""),
				colors.Reset,
				s.Text())
		})
	})

}

// printAvailableProtocols prints all available protocols
func printAvailableProtocols() {

	resp, err := http.Get("https://archlinux.org/mirrorlist/")
	helpers.CheckErr(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	helpers.CheckErr(err)

	doc.Find("#id_protocol").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			s.Find("label").Each(func(i int, s *goquery.Selection) {
				s.Find("input").Each(func(i int, s *goquery.Selection) {
					fmt.Printf("[%s%s%s] %s\n",
						colors.Green,
						s.AttrOr("value", ""),
						colors.Reset,
						s.Text())
				})
			})
		})
	})

}

// printAvailableIPVersions prints all available ipversions
func printAvailableIPVersions() {

	resp, err := http.Get("https://archlinux.org/mirrorlist/")
	helpers.CheckErr(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	helpers.CheckErr(err)

	doc.Find("#id_ip_version").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			s.Find("label").Each(func(i int, s *goquery.Selection) {
				s.Find("input").Each(func(i int, s *goquery.Selection) {
					fmt.Printf("[%s%s%s] %s\n",
						colors.Green,
						s.AttrOr("value", ""),
						colors.Reset,
						s.Text())
				})
			})
		})
	})

}
