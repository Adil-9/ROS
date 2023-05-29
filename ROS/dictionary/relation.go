package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ros/structure"
	"strings"
)

const specials string = "!@#$%^&*()_-+={[}]|\\:;<,>.?/"

// 136192 variations of passwords in dictionary

func Depending(strength int, tosend structure.SENT) string {
	// dict := []string{}
	name := strings.ToLower(tosend.Name)
	surname := strings.ToLower(tosend.Surname)
	// email_full := strings.ToLower(tosend.Email)
	// index := strings.Index(email_full, "@")
	// var email_bgn string
	// if index != -1 {
	// 	email_bgn = email_full[:index]
	// }
	// login := strings.ToLower(tosend.Login)
	year := tosend.Year
	password := tosend.Password
	write_bf(year, name, surname)
	pass, alike := check_bf(password)
	if pass == 0 {
		return fmt.Sprintf("Dictionary contains password, %d passwords alike in dictionary, weak password", alike)
	} else if pass == -1 {
		return "Internal server error"
	} else if alike != 0 {
		return fmt.Sprintf("Dictionary contains %d passwords alike your password or vice verce, password not very good", alike)
	} else if pass == 1 {
		return "Password is not found in dictionary"
	}
	return "Some string"
}

func check_bf(password string) (int, int) {
	dict := []string{}
	// Open the file for reading
	file, err := os.Open("bf_dic.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1, -1
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		dict = append(dict, line)
	}

	alike := 0
	fail := false
	for i := range dict {
		if password == dict[i] {
			fail = true
		} else if strings.Contains(password, dict[i]) || strings.Contains(dict[i], password) {
			alike++
		}
	}
	if fail {
		return 0, alike
	}

	return 1, alike
}

func write_bf(year string, name string, surname string) {
	const_name, const_surname, const_year := name, surname, year

	f, err := os.Create("bf_dic.txt") // begin write
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, char := range specials {
		for i := 0; i < 4; i++ {
			switch i {
			case 1:
				year = string(char) + const_year + string(char)
			case 2:
				year = const_year + string(char)
			case 3:
				year = string(char) + const_year
			default:
				year = const_year
			}
			for j := 0; j < 4; j++ {
				switch j {
				case 1:
					name = string(char) + const_name + string(char)
				case 2:
					name = const_name + string(char)
				case 3:
					name = string(char) + const_name
				default:
					name = const_name
				}
				for k := 0; k < 4; k++ {
					switch k {
					case 1:
						surname = string(char) + const_surname + string(char)
					case 2:
						surname = const_surname + string(char)
					case 3:
						surname = string(char) + const_surname
					default:
						surname = const_surname
					}
					f.WriteString(fmt.Sprintln(name))                     //aa
					f.WriteString(fmt.Sprintln(strings.ToUpper(name)))    //aa
					f.WriteString(fmt.Sprintln(strings.Title(name)))      //aa
					f.WriteString(fmt.Sprintln(surname))                  //ss
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname))) //ss
					f.WriteString(fmt.Sprintln(strings.Title(surname)))   //ss
					f.WriteString(fmt.Sprintln(name + name))              //aaaa
					f.WriteString(fmt.Sprintln(surname + surname))        //ssss

					f.WriteString(fmt.Sprintln(name + surname))                                   //aass
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + surname))                  //AAss
					f.WriteString(fmt.Sprintln(name + strings.ToUpper(surname)))                  //aaSS
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + strings.ToUpper(surname))) //AASS
					f.WriteString(fmt.Sprintln(strings.Title(name) + surname))                    //Aass
					f.WriteString(fmt.Sprintln(name + strings.Title(surname)))                    //aaSs
					f.WriteString(fmt.Sprintln(strings.Title(name) + strings.Title(surname)))     //AaSs

					f.WriteString(fmt.Sprintln(year + name))                     //11aa
					f.WriteString(fmt.Sprintln(year + strings.ToUpper(name)))    //11aa
					f.WriteString(fmt.Sprintln(year + strings.Title(name)))      //11aa
					f.WriteString(fmt.Sprintln(name + year))                     //aa11
					f.WriteString(fmt.Sprintln(strings.ToUpper(name) + year))    //aa11
					f.WriteString(fmt.Sprintln(strings.Title(name) + year))      //aa11
					f.WriteString(fmt.Sprintln(year + surname))                  //11ss
					f.WriteString(fmt.Sprintln(year + strings.ToUpper(surname))) //11ss
					f.WriteString(fmt.Sprintln(year + strings.Title(surname)))   //11ss
					f.WriteString(fmt.Sprintln(surname + year))                  //ss11
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + year)) //ss11
					f.WriteString(fmt.Sprintln(strings.Title(surname) + year))   //ss11

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
					f.WriteString(fmt.Sprintln(strings.Title(name) + surname + year))                    //Ssaa11
					f.WriteString(fmt.Sprintln(year + name + strings.Title(surname)))                    //11ssAa
					f.WriteString(fmt.Sprintln(name + year + strings.Title(surname)))                    //ss11Aa
					f.WriteString(fmt.Sprintln(name + strings.Title(surname) + year))                    //ssAa11
					f.WriteString(fmt.Sprintln(year + strings.Title(name) + strings.Title(surname)))     //11SsAa
					f.WriteString(fmt.Sprintln(strings.Title(name) + year + strings.Title(surname)))     //Ss11Aa
					f.WriteString(fmt.Sprintln(strings.Title(name) + strings.Title(surname) + year))     //SsAa11

					f.WriteString(fmt.Sprintln(surname + name))                                   //ssaa
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + name))                  //SSaa
					f.WriteString(fmt.Sprintln(surname + strings.ToUpper(name)))                  //ssAA
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + strings.ToUpper(name))) //SSAA
					f.WriteString(fmt.Sprintln(strings.Title(surname) + name))                    //Ssaa
					f.WriteString(fmt.Sprintln(surname + strings.Title(name)))                    //ssAa
					f.WriteString(fmt.Sprintln(strings.Title(surname) + strings.Title(name)))     //SsAa

					f.WriteString(fmt.Sprintln(year + surname + name))                                   //11ssaa
					f.WriteString(fmt.Sprintln(surname + year + name))                                   //ss11aa
					f.WriteString(fmt.Sprintln(surname + name + year))                                   //ssaa11
					f.WriteString(fmt.Sprintln(year + strings.ToUpper(surname) + name))                  //11SSaa
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + year + name))                  //SS11aa
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + name + year))                  //SSaa11
					f.WriteString(fmt.Sprintln(year + surname + strings.ToUpper(name)))                  //11ssAA
					f.WriteString(fmt.Sprintln(surname + year + strings.ToUpper(name)))                  //ss11AA
					f.WriteString(fmt.Sprintln(surname + strings.ToUpper(name) + year))                  //ssAA11
					f.WriteString(fmt.Sprintln(year + strings.ToUpper(surname) + strings.ToUpper(name))) //11SSAA
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + year + strings.ToUpper(name))) //SS11AA
					f.WriteString(fmt.Sprintln(strings.ToUpper(surname) + strings.ToUpper(name) + year)) //SSAA11
					f.WriteString(fmt.Sprintln(year + strings.Title(surname) + name))                    //11Ssaa
					f.WriteString(fmt.Sprintln(strings.Title(surname) + year + name))                    //Ss11aa
					f.WriteString(fmt.Sprintln(strings.Title(surname) + name + year))                    //Ssaa11
					f.WriteString(fmt.Sprintln(year + surname + strings.Title(name)))                    //11ssAa
					f.WriteString(fmt.Sprintln(surname + year + strings.Title(name)))                    //ss11Aa
					f.WriteString(fmt.Sprintln(surname + strings.Title(name) + year))                    //ssAa11
					f.WriteString(fmt.Sprintln(year + strings.Title(surname) + strings.Title(name)))     //11SsAa
					f.WriteString(fmt.Sprintln(strings.Title(surname) + year + strings.Title(name)))     //Ss11Aa
					f.WriteString(fmt.Sprintln(strings.Title(surname) + strings.Title(name) + year))     //SsAa11
				}
			}
		}
	}
}
