package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Depending()
	check_bf("here")
}

const specials string = "!@#$%^&*()_-+={[}]|\\:;<,>.?/"

func Depending() {
	// dict := []string{}
	name := strings.ToLower("adil")
	surname := strings.ToLower("maralbayev")
	// email_full := strings.ToLower(tosend.Email)
	// index := strings.Index(email_full, "@")
	// var email_bgn string
	// if index != -1 {
	// 	email_bgn = email_full[:index]
	// }
	// login := strings.ToLower(tosend.Login)
	year := "2003"
	// password := tosend.Password
	const_name, const_surname, const_year := name, surname, year

	f, err := os.Create("bf_dic.txt") // begin write
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, char := range specials {
		for i := 0; i < 4; i++ {
			switch i {
			case 0:
				year = string(char) + const_year + string(char)
			case 1:
				year = const_year + string(char)
			case 2:
				year = string(char) + const_year
			default:
				year = const_year
			}
			for j := 0; j < 4; j++ {
				switch j {
				case 0:
					name = string(char) + const_name + string(char)
				case 1:
					name = const_name + string(char)
				case 2:
					name = string(char) + const_name
				default:
					name = const_name
				}
				for k := 0; k < 4; k++ {
					switch k {
					case 0:
						surname = string(char) + const_surname + string(char)
					case 1:
						surname = const_surname + string(char)
					case 2:
						surname = string(char) + const_surname
					default:
						surname = const_surname
					}

					f.WriteString(fmt.Sprintln(name + surname))                                   //aass
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + surname))                  //AAss
					f.WriteString(fmt.Sprintln(name + strings.ToUpper(surname)))                  //aaSS
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + strings.ToUpper(surname))) //AASS
					f.WriteString(fmt.Sprintln(strings.Title(name) + surname))                    //Aass
					f.WriteString(fmt.Sprintln(name + strings.Title(surname)))                    //aaSs
					f.WriteString(fmt.Sprintln(strings.Title(name) + strings.Title(surname)))     //AaSs
					name, surname = surname, name
					f.WriteString(fmt.Sprintln(name + surname))                                   //ssaa
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + surname))                  //SSaa
					f.WriteString(fmt.Sprintln(name + strings.ToUpper(surname)))                  //ssAA
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + strings.ToUpper(surname))) //SSAA
					f.WriteString(fmt.Sprintln(strings.Title(name) + surname))                    //Ssaa
					f.WriteString(fmt.Sprintln(name + strings.Title(surname)))                    //ssAa
					f.WriteString(fmt.Sprintln(strings.Title(name) + strings.Title(surname)))     //SsAa
					name, surname = surname, name
					f.WriteString(fmt.Sprintln(year + name + surname))                                   //11ssaa
					f.WriteString(fmt.Sprintln(name + year + surname))                                   //ss11aa
					f.WriteString(fmt.Sprintln(name + surname + year))                                   //ssaa11
					f.WriteString(fmt.Sprintln(year + strings.ToUpper(name) + surname))                  //11SSaa
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + year + surname))                  //SS11aa
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + surname + year))                  //SSaa11
					f.WriteString(fmt.Sprintln(year + name + strings.ToUpper(surname)))                  //11ssAA
					f.WriteString(fmt.Sprintln(name + year + strings.ToUpper(surname)))                  //ss11AA
					f.WriteString(fmt.Sprintln(name + strings.ToUpper(surname) + year))                  //ssAA11
					f.WriteString(fmt.Sprintln(year + strings.ToUpper(name) + strings.ToUpper(surname))) //11SSAA
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + year + strings.ToUpper(surname))) //SS11AA
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + strings.ToUpper(surname) + year)) //SSAA11
					f.WriteString(fmt.Sprintln(year + strings.Title(name) + surname))                    //11Ssaa
					f.WriteString(fmt.Sprintln(strings.Title(name) + year + surname))                    //Ss11aa
					f.WriteString(fmt.Sprintln(strings.Title(name)+surname) + year)                      //Ssaa11
					f.WriteString(fmt.Sprintln(year + name + strings.Title(surname)))                    //11ssAa
					f.WriteString(fmt.Sprintln(name + year + strings.Title(surname)))                    //ss11Aa
					f.WriteString(fmt.Sprintln(name + strings.Title(surname) + year))                    //ssAa11
					f.WriteString(fmt.Sprintln(year + strings.Title(name) + strings.Title(surname)))     //11SsAa
					f.WriteString(fmt.Sprintln(strings.Title(name) + year + strings.Title(surname)))     //Ss11Aa
					f.WriteString(fmt.Sprintln(strings.Title(name) + strings.Title(surname) + year))     //SsAa11
				}
			}
		}
	}
}

func check_bf(password string) {
	dict := []string{}
	// Open the file for reading
	file, err := os.Open("bf_dic.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		dict = append(dict, line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	for i := range dict {
		fmt.Println(dict[i])
	}
}
