package builder

var CHARACTERS = map[string]string{
	"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111": "█" ,
	"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000": " " ,
	"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000110000001100000000000": "." ,
	"00011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000": "|" ,
	"11000011110000110110011001100110001001000011110000011000000110000001100000011000001111000010010001100110011001101100001111000011": "X" ,
	"00000000000000000000000000000000000000000000000000000000111111111111111100000000000000000000000000000000000000000000000000000000": "-" ,
	"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001111111111111111": "_" ,
	"11111111111111110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000111111111111111111": "]" ,
	"11111111111111111100000011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001111111111111111": "[" ,
	"00011000000110000001100000011000000110000001100000011000000110000001100000011000000111111111111111111111111111110000001100000011": "A" ,
	"11111111111111110000001100000011000000110000001111111111111111111111111111111110000001100000011000000110000001111111111111111111": "B",
	"00011111111111110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011": "C",
	//"11111111111111110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000111111111111111111": "D",
	"11111111111111111100000011000000110000001100000011111111111111111100000011000000110000001100000011000000110000001100000011000000": "F",
	"11000000110000001110000001100000011100000011000000111000000110000001100000011100000011000000111000000110000001110000001100000011": "\\",
	"00000011000000110000011100000110000011100000110000011100000110000001100000111000001100000111000001100000111000001100000011000000": "/",
	"00000000011111110000000001111111000000000111111100000000011111110000000001111111000000000111111100000000011111110000000001111111": "▒",
}

var BLANK_CHARACTER = CHARACTERS["00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"]