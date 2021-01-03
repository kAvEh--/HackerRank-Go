package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(abbreviation("cCccCoccocOOCCOccoccoCooCocoOoCcoCoooooOcococccOoocCoccOcoCcoooocCoooocCwcooowcocoocOococoocooOooCCooccooCCocooccoCoococccCccocoOoCcOCocccocOoOooOOooooccOcococcOOooCccooCoccOccoCcoCccOcccOoCCococCooOCoocccocoOocoOCCcoccOcOcccoOooooOOOoOcCcocCoCoOCOOcOcOOocooooocoCccocooocoooocccccooccccCCcoocococCcccCOcccccOoooOoooCcocccooocoCccOCCCccooccOwcCoccCcCcccocooOocCocccoOccocooccOocccooocooccOcccocoocoOOCOocOoococooOoOcocoocOcCcococcocCcoCoCOoOcoOOccoCcOoococoCooocccCooccCCcccCOooocoCOoOCcCccccocwcoCCOOcOoOccccCcocoCCococcCooOCcocccocOcocoocooooCoccccooOccCocoOOocococooOcccCocoOoccoCoocOccOoOOooooooocCoocCCcccococoooocCcoOooooOCcOccCooooOoocccccocoOocoCccCCcwOoOcocoocoOocccoOoCccocoocccccccooowccccOcCCoocooocOooococOOoooccoOwooOCccccoooocCooooooooooCwcooCcccoOcCoooOoOcwOoCoCcCocwoOOCcoocOOcCooocOoOooOoOccOcccocCoOcOcocoococcOoooccccccCoCooCcoooocCccOCccCooCcoOCcocoocOcoocoooOocCcCcoocoCOoOoocooCococoOccCoCoooocOoooOcoooCccocoocococOcOCccccoocccccccCooOoowoOcooOcCCOoCccCocooccoccoCCoccocOcccCo",
		"CCCOOCCOCOCOCOOCCCOOCOCOCOCOCCCOOCOOCOOWCOOOOOOOOOOCCOOCCCCCOCOCOOCOCOOOOOOOOOOOCCOCOOCCCCCOCOOCCCCOCOCCOOOOCCCOCOOCOOOOOOOCCCCOOCOOCOOOOOOCCOOOOCCCCCCCCOCCCCOOOOOCOCOCCOCCCCOCOCCCCCCCOOCCOCCOCCOOOCOOOOCOOOOOOCCOOCCOOCOCCCCOOOOOOOCOCOCCCCCCOOOCCOOCCCCCCOOCOOCCCCCOOCCCOCCCOCOOCOCCOCCOOOOOOOOOCCOOOCCOCOOOOOOOOCCCCOCOOCCOOOCOCCOOOOCOOOCCCWOOOCCOOCCOCCOCCOOCOCCCOCOOOOOCOOOCCCCCOCOCOOCOOWOOCCCOOOCCOOOCCOOCOOOOOOCOCCCOOCCCOOCOCOCOCCCCOCCCCOCCOOOOOCCOCCOOOOCOOCCCOOOCOOOOCCCOCOCCCOCCOOOCCOOCCCOOCCOCOOCC"))

	abbreviation("cCccCoccocOOCCOccoccoCooCocoOoCcoCoooooOcococccOoocCoccOcoCcoooocCoooocCwcooowcocoocOococoocooOooCCooccooCCocooccoCoococccCccocoOoCcOCocccocOoOooOOooooccOcococcOOooCccooCoccOccoCcoCccOcccOoCCococCooOCoocccocoOocoOCCcoccOcOcccoOooooOOOoOcCcocCoCoOCOOcOcOOocooooocoCccocooocoooocccccooccccCCcoocococCcccCOcccccOoooOoooCcocccooocoCccOCCCccooccOwcCoccCcCcccocooOocCocccoOccocooccOocccooocooccOcccocoocoOOCOocOoococooOoOcocoocOcCcococcocCcoCoCOoOcoOOccoCcOoococoCooocccCooccCCcccCOooocoCOoOCcCccccocwcoCCOOcOoOccccCcocoCCococcCooOCcocccocOcocoocooooCoccccooOccCocoOOocococooOcccCocoOoccoCoocOccOoOOooooooocCoocCCcccococoooocCcoOooooOCcOccCooooOoocccccocoOocoCccCCcwOoOcocoocoOocccoOoCccocoocccccccooowccccOcCCoocooocOooococOOoooccoOwooOCccccoooocCooooooooooCwcooCcccoOcCoooOoOcwOoCoCcCocwoOOCcoocOOcCooocOoOooOoOccOcccocCoOcOcocoococcOoooccccccCoCooCcoooocCccOCccCooCcoOCcocoocOcoocoooOocCcCcoocoCOoOoocooCococoOccCoCoooocOoooOcoooCccocoocococOcOCccccoocccccccCooOoowoOcooOcCCOoCccCocooccoccoCCoccocOcccCo",
		"CCCOOCCOCOCOCOOCCCOOCOCOCOCOCCCOOCOOCOOWCOOOOOOOOOOCCOOCCCCCOCOCOOCOCOOOOOOOOOOOCCOCOOCCCCCOCOOCCCCOCOCCOOOOCCCOCOOCOOOOOOOCCCCOOCOOCOOOOOOCCOOOOCCCCCCCCOCCCCOOOOOCOCOCCOCCCCOCOCCCCCCCOOCCOCCOCCOOOCOOOOCOOOOOOCCOOCCOOCOCCCCOOOOOOOCOCOCCCCCCOOOCCOOCCCCCCOOCOOCCCCCOOCCCOCCCOCOOCOCCOCCOOOOOOOOOCCOOOCCOCOOOOOOOOCCCCOCOOCCOOOCOCCOOOOCOOOCCCWOOOCCOOCCOCCOCCOOCOCCCOCOOOOOCOOOCCCCCOCOCOOCOOWOOCCCOOOCCOOOCCOOCOOOOOOCOCCCOOCCCOOCOCOCOCCCCOCCCCOCCOOOOOCCOCCOOOOCOOCCCOOOCOOOOCCCOCOCCCOCCOOOCCOOCCCOOCCOCOOCC")
}

var memory map[string]int

// Complete the abbreviation function below.
func abbreviation(a string, b string) string {
	fmt.Println(":::", len(memory))
	memory = make(map[string]int)
	fmt.Println(":::>>>", len(memory))
	if isValid(a, b) {
		return "YES"
	}
	return "NO"
}

func isValid(a string, b string) bool {
	if len(b) < 1 {
		return strings.ToLower(a) == a
	}
	if len(a) < 1 {
		return false
	}
	if strings.ToLower(string(a[0])) == string(a[0]) {
		if strings.ToLower(string(b[0])) == string(a[0]) {
			if memory[a[1:]+":"+b] > 0 {
				if memory[a[1:]+":"+b[1:]] > 0 {
					return memory[a[1:]+":"+b] == 1 || memory[a[1:]+":"+b[1:]] == 1
				}
				tt := isValid(a[1:], b[1:])
				if tt {
					memory[a[1:]+":"+b[1:]] = 1
				} else {
					memory[a[1:]+":"+b[1:]] = 2
				}
				return memory[a[1:]+":"+b] == 1 || tt
			}
			if memory[a[1:]+":"+b[1:]] > 0 {
				tt := isValid(a[1:], b)
				if tt {
					memory[a[1:]+":"+b] = 1
				} else {
					memory[a[1:]+":"+b] = 2
				}
				return tt || memory[a[1:]+":"+b[1:]] == 1
			}
			tt, ttt := isValid(a[1:], b), isValid(a[1:], b[1:])
			if tt {
				memory[a[1:]+":"+b] = 1
			} else {
				memory[a[1:]+":"+b] = 2
			}
			if ttt {
				memory[a[1:]+":"+b[1:]] = 1
			} else {
				memory[a[1:]+":"+b[1:]] = 2
			}
			return tt || ttt
		}
		if memory[a[1:]+":"+b] > 0 {
			return memory[a[1:]+":"+b] == 1
		}
		tt := isValid(a[1:], b)
		if tt {
			memory[a[1:]+":"+b] = 1
		} else {
			memory[a[1:]+":"+b] = 2
		}
		return tt
	} else {
		if string(a[0]) == string(b[0]) {
			tt := isValid(a[1:], b[1:])
			if tt {
				memory[a[1:]+":"+b[1:]] = 1
			} else {
				memory[a[1:]+":"+b[1:]] = 2
			}
			return isValid(a[1:], b[1:])
		}
		return false
	}
}
