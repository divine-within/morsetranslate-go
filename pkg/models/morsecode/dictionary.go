package models

var MorseCodeMap = map[string]string{
	"A":  ".-",
	"B":  "-...",
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	"0":  "-----",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	".":  ".-.-.-",
	",":  "--..--",
	"?":  "..--..",
	"'":  ".----.",
	"!":  "-.-.--",
	"/":  "-..-.",
	"(":  "-.--.",
	")":  "-.--.-",
	"&":  ".-...",
	":":  "---...",
	";":  "-.-.-.",
	"=":  "-...-",
	"+":  ".-.-.",
	"-":  "-....-",
	"_":  "..--.-",
	"\"": ".-..-.",
	"$":  "...-..-",
	"@":  ".--.-.",
	"¿":  "..-.-",
	"¡":  "--...-",
}

var ReverseMorseCodeMap = make(map[string]string)

func init() {
	ReverseMorseCodeMap = make(map[string]string)
	for key, value := range MorseCodeMap {
		ReverseMorseCodeMap[value] = key
	}
}
