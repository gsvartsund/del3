package main

import(
	"fmt"
)
//Array med fire ting inni seg av type string, fire karakterer. Positions er samme bare INT. Om båten er venstre eller høyre. True er ventsre
var Character_Names[4] string
var Character_Positions[4] int
var BoatLeftSide bool

//pws print men kortere. Her runner alt.
func main(){
Init()
	pws()
	PutIn("Mann")
	pws()
	CrossRiver()
	pws()
}

func pws(){
	fmt.Println(MakeWorldState())
}
//Begynner å sette boat ls true, legger inn alle navn med character names 0-3. Alle er på venstre side, tilbak linje 9.
func Init(){
	BoatLeftSide = true
	Character_Names[0] = "Mann"
	Character_Names[1] = "Korn"
	Character_Names[2] = "Kylling"
	Character_Names[3] = "Rev"
	//Mann
	Character_Positions[0] = 0
	//Korn
	Character_Positions[1] = 0
	//Kylling
	Character_Positions[2] = 0
	//Rev
	Character_Positions[3] = 0
}
//Dette er så go test skal fungere riktig
func ReturnWorldState() string{
	Init()
	return MakeWorldState()
}
//Kaller lag høyre, venstre båten og setter sammen til en streng. Setter sammen de forskjellige posisjonene. 
func MakeWorldState() string{
	rightSide := MakeRightSide()
	boat := MakeBoat()
	leftSide := MakeLeftSide()
	res := leftSide + boat + rightSide
	return res
}
/*Array med fire tekst strenger, variabel i = 0. Så lenge "i" er mindre en mindre en len(char). Søker gjennom og sjekker om den er 0 1 eller 2.
*Scanner hvor char er, hvis char = 2 så er det høyre siden.
*/ 
func MakeRightSide() string{
	var endResult[4]string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 2){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	res := fmt.Sprintf("%s-%s-%s-%s-", endResult[0], endResult[1], endResult[2], endResult[3])
	return res
}
//Sjekker hvor char er.
func MakeLeftSide() string{
	var endResult[4]string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 0){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	res := fmt.Sprintf("%s-%s-%s-%s-", endResult[0], endResult[1], endResult[2], endResult[3])
	return res
}
//tar inn navn på ting, oppdaterer posisjonen til 1 som er i båten. return makeworld state, returnerer en streng med vs, hs og båt. Gir beskjed hvor char er.
func PutIn(item string) string{
	for i := 0; i < len(Character_Names); i++{
		if(Character_Names[i] == item){
			Character_Positions[i] = 1
		}
	}
	return MakeWorldState()
}
//Samme som left og right. IF boat left side = true da tegner den båten på venstre siden. 99-102 printer status grafisk. res returnert
func MakeBoat() string{
	var endResult[4]string
	var res string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 1){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	if(BoatLeftSide){
		res = fmt.Sprintf(" \\_%s_%s_/________", endResult[0], endResult[1])
	}else{
		res = fmt.Sprintf("________\\_%s_%s_/  ", endResult[0], endResult[1])
	}
	
	return res
}
//Endrer båt verdien, hvis boatleftside er true. Høyre er false. makeworldstate, oppdaterer status.
func CrossRiver() string{
	BoatLeftSide = !BoatLeftSide
	
	return MakeWorldState()
}