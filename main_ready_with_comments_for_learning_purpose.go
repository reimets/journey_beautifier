package main

import (
	"bufio" // Pakub puhverdatud sisend/väljund operatsioone. See ümbritseb nt io.Reader või io.Writer  objekti, luues uue objekti (Reader või Writer), mis samuti rakendab vastavat liidest, kuid pakub puhverdamist ja abi tekstiga töötamisel
	"encoding/csv" //Rakendab CSV (Comma Separated Values ehk komadega eraldatud väärtuste) failide lugemist ja genereerimist. Kasulik tabelikujul andmete lugemiseks ja kirjutamiseks.
	"flag" // Toetab lihtsat käsurea lipukeste (flags) analüüsi. Võimaldab kasutajate poolt antud käsurea argumente lihtsalt Go programmi muutujateks teisendada.
	"fmt" // Rakendab vormindatud sisend/väljund operatsioone funktsioonidega, mis on sarnased C keele printf ja scanf funktsioonidele. Kasutatakse stringide vormindamiseks, printimiseks ja sisendi lugemiseks.
	"os" // Pakub platvormist sõltumatut liidest operatsioonisüsteemi funktsionaalsusega. Sisaldab funktsioone faili- ja protsessioperatsioonideks, keskkonnamuutujateks ja kasutaja informatsiooniks.
	"strings" //  Sisaldab kasulikke funktsioone stringidega töötamiseks. See hõlmab stringide manipuleerimist, analüüsimist ja operatsioone nagu alamstring, asendamine ja tähemärkide muutmine.
	"regexp" // Pakub regulaaravaldiste abil stringide mustrituvastust. Oluline stringide analüüsimiseks ja manipuleerimiseks mustripõhiste definitsioonide järgi.
	"time" // Pakub funktsionaalsust aja mõõtmiseks ja kuvamiseks. See hõlmab ajavahemike arvutamist, aja vormindamist ja aja stringide analüüsimist.
)

	// Struktuuri Airport 
// kasutatakse lennujaama andmete salvestamiseks ja haldamiseks, võimaldades neid andmeid kergesti kasutada ja edastada programmi eri osade vahel.

// Struktuurid on ideaalsed, kui vajate keerukamat andmetüüpi, mis koondab erinevaid andmeid (nagu stringid, intid, muud struktuurid) ühte loogilisse üksusesse.
// Struktuure kasutatakse, defineerides kõigepealt struktuuri väljad, seejärel luues selle struktuuri eksemplare (instantsid) ja andes nendele väärtused.
// Neid saab kasutada funktsioonide argumentidena, tagastusväärtustena või andmestruktuuridena andmete salvestamiseks ja töötlemiseks.
/*
	Millal mitte kasutada?
Kui teil on vaja ainult lihtsaid, ühetüübilisi andmeid (nagu ainult integerid või stringid), võib mõnikord olla lihtsam ja otstarbekam kasutada põhilisi 
andmetüüpe või massiive/slice'e.
Samuti, kui teie andmestruktuur on väga keeruline või sisaldab palju omavahel seotud andmeid, võib olla parem kasutada klasside ja objektidega keeli 
nagu Java või C#, kus on rohkem funktsioone keerukate andmestruktuuride haldamiseks.
*/
// Airport represents the structure for airport data from CSV
type Airport struct {
	Name			string
	Country 		string
	Municipality	string
	ICAO			string
	IATA			string
	Coordinates		string
}

// Peamine sisenemispunkt programmile.
// Main entry point for the program.
func main() {
    // Kontrollib, kas käsurealt on antud õige arv argumente.
    // Checks if the correct number of arguments are provided from the command line.
    if len(os.Args) != 4 { // main.go input.txt output.txt airport-lookup.csv
        displayTheUsage()
        return
    }

    // Loeb käsurea argumendid.
	// See meetod on otsekohene ja kasutatakse siis, kui argumentide arv on fikseeritud või kui soovite neid argumendid otse töödelda.
	inputTxt := os.Args[1]
	outputTxt := os.Args[2]
	airportLookupCsv := os.Args[3]
	// või teine võimalus oleks kirjutada nii:
	// inputTxt, outputTxt, airportLookupCsv := os.Args[1], os.Args[2], os.Args[3]

    // Kontrollib vigu nagu sisendi olemasolu ja lipu '-h' kasutamist.
    // Checks for errors such as input existence and usage of '-h' flag.
/* 
	Eeldab, et olete juba eraldanud programmi nime ja ainult argumendid massiivist args, mis võib olla näiteks os.Args[1:] tulemus. 
	See on kasulik, kui teie programm kasutab flag paketti või muud argumentide töötlemise viisi, mis eraldab programmi nime ja argumendid.
	*/

	// if error found, then returning "return" and program stops
	if !checkingErrors(inputTxt, outputTxt, airportLookupCsv) {
		return
		/* 
kontrollib, kas funktsioon checkingErrors tagastab false. Eitusmärk ! ees tähendab, et kontrollitakse, kas funktsiooni tulemus ei ole true. 
Kui checkingErrors tagastab false, siis jõuab programm return avaldiseni, mis lõpetab main funktsiooni töö. Seega, kui checkingErrors leiab 
mõne vea (näiteks kui sisendfaili nimi kattub väljundfaili nimega), lõpetab programm oma töö.
*/
    }


	/* Funktsioon loeb CSV-faili sisu ja konverteerib selle lennujaamade andmeteks, mida hoitakse airports muutujas. Täpsemalt, see tagastab map tüüpi muutuja, 
	kus võtmeks on lennujaama kood (näiteks IATA või ICAO kood) ja väärtuseks on Airport tüüpi objekt, mis sisaldab lennujaama andmeid nagu nimi, riik, linn jne.
Funktsioon võib samuti tagastada vea (err), kui faili lugemisel või andmete töötlemisel tekib probleem. See võib juhtuda mitmel põhjusel, 
näiteks kui faili ei leita, fail on valesti vormindatud või kui faili lugemisel tekib muu tõrge.*/
    // main funktsioonis kutsutakse loadAirportData funktsiooni, et laadida lennujaamade andmed 
	// muutujasse airports. Seejärel kontrollitakse, kas laadimisprotsessis esines vigu:
    airports, err := loadAirportData(airportLookupCsv)
	/* Vea Kontroll: Kui loadAirportData funktsioonist tagastatakse viga (st err != nil), 
	siis prinditakse veateade konsooli ja programm lõpetab töö. See tagab, et programmi ei jätkata 
	vigaste või puuduvate andmetega.*/
    if err != nil {
        fmt.Println(err)
        return
		/* Edukas Andmete Laadimine: Kui vea muutuja on nil, tähendab see, et lennujaamade andmed on 
		edukalt laetud ja programm jätkab tööd airports map-iga, mis sisaldab nüüd lennujaamade andmeid.*/
    }

	// Avab sisendfaili lugemiseks.
    inputFile, err := os.Open(inputTxt)
    if err != nil {
        fmt.Println("\033[1m\033[31m-------------------\033[0m\033[22m\n"+
            "\033[31m| Input not found |\033[0m\n"+
            "\033[1m\033[31m-------------------\033[0m\033[22m")
        return
    }
    defer inputFile.Close()

    // Töötleb sisendfaili, teisendades lennujaama koodid ja kuupäevad.
    allCodesFound, output, err := processItinerary(inputFile, airports)
    if err != nil {
        if err.Error() == "Input file is empty" {
            fmt.Println(err) // Outputs an error message in the terminal.
            return // Exits the program without creating or overwriting the output.txt file.
        }
        fmt.Println(err) // Outputs another error message in the terminal.
        return
    }

        // Kui kõik koodid ei leitud, lõpetab programm töö veateatega.
    if !allCodesFound {
        fmt.Println("\033[1m\033[31m----------------------------------------------------------\033[0m\033[22m\n"+
        "\033[31m| Error: Not all airport codes were found in the lookup. |\033[0m\n"+
        "\033[1m\033[31m----------------------------------------------------------\033[0m\033[22m")
        return // Program stops and does not overwrite output.txt file
    }    
    
    fmt.Println(output) // prints output info to terminal

    output = trimColor(output) // removes colors from text, because of need to print info to putput.txt file

    // Kirjutab töödeldud andmed väljundfaili.
    err = os.WriteFile(outputTxt, []byte(output), 0644)
    if err != nil {
        fmt.Println("\033[31mError writing output:\033[0m", err)
        return
    }

    // Addition: Displays success message after successfully writing to the output file.
    fmt.Println("\033[1m\033[32m-------------------------------------\033[0m\033[22m\n"+
        "\033[1m\033[32m| Itinerary processed successfully. |\033[0m\033[22m\n"+
        "\033[1m\033[32m-------------------------------------\033[0m\033[22m")
}


/* See funktsioon, processItinerary, on mõeldud reisiandmete töötlemiseks: see loeb sisendfailist 
ridadena teksti, teisendab lennujaama koodid nende täisnimedeks, formaadib kuupäevad ja ajad inimloetavaks 
ning eemaldab üleliigsed tühikud ja tühjad read. Funktsioon tagastab kolm väärtust: boolean, mis näitab, 
kas kõik lennujaama koodid leiti, töödeldud tekstistringi ja vea, kui see esineb.
*/
// Töötleb reisiandmed, asendades lennujaama koodid ja formaadib kuupäevad.
	/* Asterisk * enne os.File tüüpi 
funktsiooni processItinerary parameetris tähendab, 
et inputFile on osutaja (pointer) os.File objektile. Osutajate kasutamine Go keeles võimaldab 
funktsioonidel viidata otse muutujatele või objektidele mälus, mitte nende koopiatele. 
See tähendab, et inputFile parameeter ei ole tegelik os.File objekti koopia, vaid viit sellele, 
mis võimaldab funktsioonil lugeda ja manipuleerida faili, millele osutaja viitab.

os.File tüüpi kasutatakse failioperatsioonides, nagu lugemine või kirjutamine, ja 
*os.File annab funktsioonile processItinerary võime töötada otse sisendfailiga, 
lugedes selle sisu või käideldes faili muid aspekte, nagu vea kontrollimine või faili sulgemine.
*/
func processItinerary(inputFile *os.File, airports map[string]Airport) (bool, string, error) {
	// bufio.NewScanner(inputFile) loob uue skanneri, mis loeb teksti ridade kaupa sisendfailist. 
	// See on mugav viis tekstifailide järkjärguliseks lugemiseks.
	scanner := bufio.NewScanner(inputFile)

	/* strings.Builder on Go standardteegis olev struktuur, mida kasutatakse efektiivseks stringide koostamiseks või ehitamiseks.
		Efektiivsus: strings.Builder pakub efektiivset viisi stringide kokkupanekuks, eriti olukordades, 
		kus peate dünaamiliselt lisama palju väikeseid stringitükke, et luua suurem string. 
		Erinevalt stringide liitmise operatsioonist (+), mis võib olla kulukas suure hulga liitmiste 
		korral, kasutab strings.Builder sisemist puhvrit, et vähendada vajadust mälu uuesti jaotamise 
		järele, mis teeb stringide kokkupanemise kiiremaks ja mälusäästlikumaks.

	Meetodid: strings.Builder pakub mitmeid meetodeid stringi koostamiseks, sealhulgas Write, WriteString, 
	WriteRune ja WriteByte, mis võimaldavad lisada vastavalt baitide massiivi, stringi, üksiku rune'i 
	või baiti. Pärast kõigi vajalike osade lisamist saab lõpliku stringi kätte String meetodi abil.

	Kasutusnäide: Kui teil on vaja koostada string mitmest allikast (näiteks failist lugemisel, 
	kasutaja sisendist või andmetöötluse tulemustest), siis strings.Builder võimaldab teil lisada 
	need jupid järk-järgult ilma, et iga lisamise korral tekiks uus stringi koopia.

	Siin on lihtne näide strings.Builder kasutamisest:
		var builder strings.Builder

		builder.WriteString("Hello, ")
		builder.WriteString("world!")

		result := builder.String() // result on "Hello, world!"
	var output strings.Builder reas loob output muutuja, mida seejärel saab kasutada stringi 
	järkjärguliseks ehitamiseks programmi eri osades, ja lõpuks saab koostatud stringi kätte 
	output.String() meetodiga.
	*/
    var output strings.Builder // loob muutuja nimega output, mille tüüp on strings.Builder. 
    allCodesFound := true
    isFirstLine := true

		/* Muutuja isFirstLine kontrollimine: 
	See loogika kontrollib, kas failis on üldse sisu. Kui esimene rida on tühi ja see on ainuke rida, 
	siis programm tuvastab, et fail on tühi. See on oluline samm, et vältida tühja faili töötlemist.
	*/
    for scanner.Scan() {
        if isFirstLine {
            isFirstLine = false
            // Checks if the first line is empty
            if len(scanner.Text()) == 0 {
                continue // Continues to the next line if the first line is empty
            }
        }

		/* Töötlemistsükkel: For-tsükkel kasutab skannerit, et lugeda failist rida rea haaval. 
	Iga rea puhul:
Eemaldatakse üleliigsed tühikud ja tühjad read funktsiooniga cleanUpText.
Asendatakse lennujaama koodid täisnimedega funktsiooniga replaceAirportCodes. Kui mõnda koodi ei leita, 
märgitakse allCodesFound false-ks.
Formaaditakse kuupäevad ja kellaajad funktsiooniga formatDateAndTime.
Lisatakse töödeldud rida väljundisse output.
*/

        line := scanner.Text()
        line = cleanUpText(line) // Removes extra spaces and empty lines.
        processedLine, codesFound := replaceAirportCodes(line, airports)
        if !codesFound {
            allCodesFound = false
        }
        processedLine = formatDateAndTime(processedLine)
	/* WriteString meetod: See on strings.Builder tüüpi objekti meetod, mille ülesanne on lisada etteantud 
	string output puhvrisse. See tähendab, et processedLine sisu lisatakse output muutuja juba olemasoleva 
	sisu järele.
		processedLine + "\n": Siin processedLine on string, mis sisaldab töödeldud rida, millele on lisatud 
		uue rea märk (\n). Uue rea märk tagab, et järgmine WriteString käsuga lisatav sisu alustab 
		uuest reast, muutes lõpliku teksti loetavamaks, eriti kui seda kuvatakse konsoolis või 
		kirjutatakse faili.
*/
        output.WriteString(processedLine + "\n")
    }

	/* Tühja faili kontroll: Pärast tsükli lõppu kontrollitakse, kas isFirstLine on endiselt true, 
	mis tähendaks, et failis polnud ühtegi rida (või ainult tühje ridu). Sellisel juhul tagastatakse 
	veateade, et fail on tühi.
	*/
    // Checks if no lines were read (file may be empty)
    if isFirstLine {
        return false, "", fmt.Errorf("\033[1m\033[31m-----------------------\033[0m\033[22m\n"+
            "\033[31m| Input file is empty |\033[0m\n"+
            "\033[1m\033[31m-----------------------\033[0m\033[22m")
    }

	// Vea kontroll skannerilt: Kontrollitakse, kas skanner on lugemisel kohanud mingeid vigu. 
	// Kui jah, tagastatakse vastav veateade.
    if err := scanner.Err(); err != nil {
        return false, "", fmt.Errorf("\033[31mError reading input: %v \033[0m", err)
    }

	/* Lõpptulemuse tagastamine: Kui kõik read on töödeldud ja vigu ei esinenud, 
	tagastatakse lõplikult töödeldud tekst, info kõikide koodide leidmise kohta ja nil veana 
	(mis tähendab, et vigu ei esinenud).
	*/
    finalOutput := cleanUpText(output.String()) // Checks that the final result is cleaned up as well.
    return allCodesFound, finalOutput, nil
}


// removes colors from text, because of need to print info to putput.txt file
// without that func output.txt file lines were with strange marks
// func trimColor(text string) string {
// 	// text color
// 	text = strings.ReplaceAll(text, "\033[0m", "")
// 	text = strings.ReplaceAll(text, "\033[30m", "")
// 	text = strings.ReplaceAll(text, "\033[31m", "")
// 	text = strings.ReplaceAll(text, "\033[32m", "")
// 	text = strings.ReplaceAll(text, "\033[33m", "")
// 	text = strings.ReplaceAll(text, "\033[34m", "")
// 	text = strings.ReplaceAll(text, "\033[35m", "")
// 	text = strings.ReplaceAll(text, "\033[36m", "")
// 	text = strings.ReplaceAll(text, "\033[37m", "")
// 	// text background color
// 	text = strings.ReplaceAll(text, "\033[40m", "")
// 	text = strings.ReplaceAll(text, "\033[41m", "")
// 	text = strings.ReplaceAll(text, "\033[42m", "")
// 	text = strings.ReplaceAll(text, "\033[43m", "")
// 	text = strings.ReplaceAll(text, "\033[44m", "")
// 	text = strings.ReplaceAll(text, "\033[45m", "")
// 	text = strings.ReplaceAll(text, "\033[46m", "")
// 	text = strings.ReplaceAll(text, "\033[47m", "")
// 	// bold
// 	text = strings.ReplaceAll(text, "\033[1m", "") // to bold
// 	text = strings.ReplaceAll(text, "\033[22m", "") // back to normal
// 	return text
// }
// Selle eelmise pika variandi lühem variant regexiga:
// removes colors from text, because of need to print info to putput.txt file
// without that func output.txt file lines were with strange marks
func trimColor(text string) string {
	// Defineerib regulaaaravaldis, mis sobitub ANSI värvikoodidega.
	colorCodeRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	// Asendab kõik leitud värvikoodid tühja stringiga.
	return colorCodeRegex.ReplaceAllString(text, "")
}


// replaceAirportCodes replaces airport codes (both IATA and ICAO) in the text with
// the corresponding airport names, using a provided map of airport data.
// This function employs regular expressions to identify and replace IATA and ICAO codes.
// Replaces airport codes with their full names.
// Asendab lennujaama koodid nende täisnimedega.
func replaceAirportCodes(line string, airports map[string]Airport) (string, bool) {
		/* 
	airports map[string]Airport deklareerib funktsiooni parameetri nimega airports, mis on kaarditüüp (map) Go keeles. 
	Selles kaardis on võtmed string tüüpi ja väärtused on Airport tüüpi, kus Airport on kasutaja defineeritud struktuur, mis esindab lennujaama andmeid. 
	See tähendab, et airports kaardi abil saab lennujaama koodi (string) põhjal leida vastava lennujaama andmed (Airport struktuur).
	*/
	// Asendab ICAO ja IATA koodid lennujaamade nimedega. Järjestus on oluline!
	// Regular expressions to identify IATA and ICAO codes
    allCodesFound := true
    reIATA := regexp.MustCompile(`#\w{3}`)
    reICAO := regexp.MustCompile(`##\w{4}`)

	// Replaces ICAO and IATA codes with airport names. Order matters!
    line = reICAO.ReplaceAllStringFunc(line, func(code string) string {
        if airport, exists := airports[code[2:]]; exists {
            return "\033[1m\033[33m"+airport.Name+"\033[0m\033[22m"
        } else {
            allCodesFound = false
            return code // The code is left unchanged if not found.
        }
    })

    line = reIATA.ReplaceAllStringFunc(line, func(code string) string {
        if airport, exists := airports[code[1:]]; exists {
            return "\033[1m\033[33m"+airport.Name+"\033[0m\033[22m"
        } else {
            allCodesFound = false
            return code // The code is left unchanged if not found.
        }
    })
    return line, allCodesFound
}


// Loads airport data from a CSV file.
// Faili Avamine: os.Open(filePath) proovib avada CSV-faili, mille asukoht on antud filePath muutuja abil. 
// Kui faili ei leita või seda ei saa mingil põhjusel avada, tagastatakse err muutuja veaga ning 
// funktsioon lõpetab töö, tagastades nil ja veateate. 
// Veateade on siin vormindatud ANSI värvikoodidega, et teha see konsooli väljundis visuaalselt märgatavamaks.
	// See on kaart (map), kus võtmed (key) on tüüpi string ja väärtused (value) on Airport tüüpi.
	func loadAirportData(filePath string) (map[string]Airport, error) {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("\033[1m\033[31m----------------------------\033[0m\033[22m\n"+
				"\033[31m| Airport lookup not found |\033[0m\n"+
				"\033[1m\033[31m----------------------------\033[0m\033[22m")
		}
		/* defer avaldis lükkab funktsiooni või meetodi käivitamise edasi kuni ümbritseva funktsiooni töö 
		lõpetamiseni. defer kasutatakse tavaliselt ressursside, nagu failide või võrguühenduste, korrektseks 
		vabastamiseks isegi siis, kui funktsioonis ilmneb viga või kui see lõpetab töö enneaegselt 
		(näiteks return avaldise tõttu).
	
	Avaldise defer file.Close() kasutamine Go programmis tagab, et fail, mille viit on file, 
	suletakse kindlasti enne funktsiooni töö lõppu. See on oluline, et vältida ressursside lekkeid, 
	mis võivad tekkida, kui avatud failid jäävad sulgemata.
	*/
		defer file.Close()
	
	
	/* CSV Andmete Lugemine: Kasutades csv.NewReader, loetakse kogu faili sisu records muutujasse. 
	Iga record on ridade kaupa esitatud stringide massiiv, mis esindab ühe lennujaama andmeid. */
		reader := csv.NewReader(bufio.NewReader(file))
		/* kasutab csv paketi Reader tüüpi objekti, et lugeda kõik andmed CSV-failist korraga. 
		Siin reader on csv.Reader tüüpi objekt, mis on seadistatud lugema andmeid mingist sisendvoost, 
		antud juhul avatud failist. ReadAll() meetod loeb kõik järelejäänud rekordid failist ja tagastab 
		need kahe mõõtmelise stringide massiivina ([][]string)([rida][veerg]), kus iga alammassiiv esindab ühte rida 
		CSV-failist.
		*/
		records, err := reader.ReadAll()
		if err != nil {
			return nil, fmt.Errorf("\033[1m\033[31m--------------------------------\033[0m\033[22m\n"+
				"\033[31m| Error reading airport lookup |\033[0m\n"+
				"\033[1m\033[31m--------------------------------\033[0m\033[22m")
			/*     reader := csv.NewReader(bufio.NewReader(file))
	See rida kasutab kahte erinevat NewReader funktsiooni, et luua lugemisahel, mis võimaldab CSV-andmete 
	efektiivset lugemist failist. Iga NewReader funktsioon kuulub erinevasse paketti ja teenib erinevat eesmärki:
	
	Pakett: bufio
	Eesmärk: See funktsioon loob puhverdatud lugeja (bufio.Reader) antud io.Reader liidesest, 
	sel juhul avatud faili käepidemest (file). Puhverdatud lugeja eelis on see, et see loeb sisendvoogust 
	andmeid suuremate blokkidena, vähendades seeläbi IO operatsioonide arvu. 
	See muudab failist lugemise efektiivsemaks, eriti suurte failide puhul, kuna vähendab korduvaid 
	süsteemikõnesid failisüsteemi.
	
	Pakett: encoding/csv
	Eesmärk: See funktsioon loob uue CSV-lugeja (csv.Reader) antud io.Reader liidesest, mis on siin 
	puhverdatud lugeja, mille lõi eelmine samm. csv.Reader kasutab seda lugejat, et lugeda ja analüüsida 
	CSV-formaadis andmeid, teisendades need Go programmile töötlemiseks kättesaadavaks struktuuriks. 
	CSV-lugeja tegeleb CSV-formaadile iseloomulike aspektidega, nagu eraldajate tuvastamine, ridade lõppude 
	käsitlemine ja tsitaatide haldamine.
	
	Kokkuvõttes, bufio.NewReader kasutamine enne csv.NewReader kasutamist on optimeerimisvõte. 
	See võimaldab csv.Reader-il töötada puhverdatud sisendiga, parandades jõudlust ja efektiivsust, 
	kuna vähendab failisüsteemiga suhtlemise vajadust.
	*/
		}
	
	
		/* Andmete Töötlemine: Funktsioon seejärel itereerib läbi kõikide records ridade, 
		luues iga rea põhjal Airport tüüpi objekti. Need objektid lisatakse airports map-i, 
		kus võtmeks on lennujaama IATA või ICAO kood ja väärtuseks on Airport objekt.*/
	// make on sisseehitatud funktsioon Go keeles, mida kasutatakse uute andmestruktuuride loomiseks. 
	// Seda saab kasutada kaartide (maps), viilude (slices) ja kanalite (channels) loomiseks.
			/* map[string]Airport määratleb kaardi tüübi, kus:
			string on kaardi võtmete tüüp. Iga võti kaardil on unikaalne string.
			Airport on kaardi väärtuste tüüp. Iga võti kaardil on seotud Airport tüüpi väärtusega.
			*/
	/* See konkreetne rida loob uue tühja kaardi nimega airports, mida saab kasutada lennujaamade 
	salvestamiseks, kus iga lennujaama unikaalne kood (näiteks IATA või ICAO kood) on kaardi võtmeks ja 
	vastav Airport struktuur on selle võtme väärtuseks. See võimaldab kiiret juurdepääsu lennujaama andmetele 
	nende koodide järgi.
	*/
		airports := make(map[string]Airport)
		for i, record := range records {
			if i == 0 { // Checks header
				for _, header := range record {
					/* Andmete Kontroll: Iga rea puhul kontrollitakse, kas andmed on korrektselt vormindatud ja 
					kas kõik vajalikud väljad on olemas. Kui andmed on vigased, tagastatakse veateade.*/
					if header == "" {
						return nil, fmt.Errorf("\033[1m\033[31m----------------------------\033[0m\033[22m\n"+
						"\033[31m| Airport lookup malformed |\033[0m\n"+
						"\033[1m\033[31m----------------------------\033[0m\033[22m")
						/* Errorf kasutamine on kasulik olukordades, kus veateates on vaja kuvada muutuvat informatsiooni, 
						näiteks failinime, kus viga esines, või konkreetset väärtust, mis põhjustas vea. 
						See võimaldab arendajatel luua informatiivsemaid veateateid, mis aitavad vigade põhjust kiiremini 
						tuvastada ja parandada.
						*/
					}
				}
				/* Selle lõigu lõpus olev sõna 
					continue 
				kasutatakse siin selleks, et jätkata järgmise 
				iteratsiooniga for tsüklis, kui praegune iteratsioon on lõpetatud. Antud kontekstis tähendab 
				see, et kui tsükli käesolev iteratsioon töötleb esimest rida (mis on päise rida records 
					massiivist), siis pärast päise rea kontrollimist (et veenduda, kas kõik päised on olemas 
						ja ükski neist pole tühi), ütleb continue käsk programmil minna järgmisele 
						iteratsioonile for tsüklis ilma järgnevaid käske käesolevas iteratsioonis täitmata.
	
	Antud juhul kasutatakse continue-i, et vältida päiserida (i == 0) käsitlemast samamoodi nagu andmeridu. 
	Päiserida sisaldab veergude nimesid, mitte andmeid, seega ei peaks seda töötlema samamoodi nagu andmeridasid. 
	continue-i kasutamine siin tagab, et kui päiserida on edukalt kontrollitud, liigub programm edasi 
	järgmisele reale (tõeline andmerida), ilma et püüaks päiserida kui lennujaama andmet kirja panna.
	*/
				continue
			}
	
				// Airport struktuuri määratlus alguses ja selle kasutamine loadAirportData funktsioonis on 
				// omavahel tihedalt seotud: struktuur määratleb andmete vormi ja funktsioon kasutab seda vormi, 
				// et luua konkreetseid lennujaama andmete eksemplare.
			// Creates an airport data structure.
			airport := Airport {
				Name:         record[0],
				Country:      record[1],
				Municipality: record[2],
				ICAO:         record[3],
				IATA:         record[4],
				Coordinates:  record[5],
				/* Kui loadAirportData funktsioonis loetakse CSV-failist lennujaama andmed, kasutatakse Airport 
				struktuuri iga loetud rea andmete salvestamiseks. Seejärel lisatakse loodud Airport eksemplarid 
				airports kaardile (map), kasutades võtmena kas ICAO või IATA koodi. See struktuur võimaldab 
				programmil efektiivselt juurdepääsu lennujaama andmetele, otsides neid IATA või ICAO koodi alusel, 
				et teisendada koodid lennujaama nimedeks või muuks vajalikuks teabeks.
				*/
			}
	/* read on vajalikud lennujaama andmete salvestamiseks mappi (map), kus võtmeteks on lennujaama 
	ICAO ja IATA koodid ning väärtuseks on Airport tüüpi objekt, mis sisaldab lennujaama kohta käivat 
	informatsiooni (nimi, riik, linn, ICAO kood, IATA kood ja koordinaadid).
	*/
			airports[airport.ICAO] = airport
			airports[airport.IATA] = airport
	
			// Checks data integrity.
			if len(record) != 6 || airport.IATA == "" || airport.ICAO == "" || airport.Name == "" || airport.Municipality == "" || airport.Country == "" || airport.Coordinates == "" {
				return nil, fmt.Errorf("\033[1m\033[31m----------------------------\033[0m\033[22m\n"+
				"\033[31m| Airport lookup malformed |\033[0m\n"+
				"\033[1m\033[31m----------------------------\033[0m\033[22m")
			}
		}
		/* Tagastamine: Kui kõik andmed on edukalt töödeldud, tagastatakse airports map ja nil viga, 
		mis näitab, et protsess läks edukalt.*/
		return airports, nil
		/* Kokkuvõttes, loadAirportData funktsioon vastutab lennujaamade andmete lugemise ja valideerimise 
		eest CSV-failist, ning main funktsioon kasutab neid andmeid, tagades samal ajal, et andmete 
		laadimisel ei tekiks vigu, mis võiksid programmi tööd häirida.*/
	}


// Puhastab teksti, eemaldades üleliigsed tühikud ja mitu järjestikust reavahetust.
func cleanUpText(text string) string {
    // Eemaldab üleliigsed tühikud ja tabulaatorid.
    spacePattern := regexp.MustCompile(`[ \t]+`)
    text = spacePattern.ReplaceAllString(text, " ")

    // Eemaldab mitu järjestikust reavahetust.
    newlinePattern := regexp.MustCompile(`\n\n+`)
    text = newlinePattern.ReplaceAllString(text, "\n\n")

    // Eemaldab tühjad read teksti algusest ja lõpust.
    text = strings.TrimSpace(text)

    return text
}


// formatDateAndTime identifies and formats date and time strings in the text.
// This function uses regular expressions to detect patterns of dates and times
// and converts them into a human-readable format. If a date and time pattern is detected
// in the text, it is replaced with the appropriate format. If no pattern is found or it is
// invalid, the original text is returned.
func formatDateAndTime(line string) string {
	// Checks if the line starts with the specified text and leaves it in the original color
	if strings.HasPrefix(line, "Your flight departs from") {
		return line // returns line without changing color (here)
	}
	// working with patterns using regular expressions
	regs := regexp.MustCompile(`\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}([Z+-]\d{2}:\d{2}|Z)?)\)`) 
	// regexp.MustCompile: See funktsioon võtab argumendina regulaaravaldise stringina ja tagastab Regexp objekti, mida saab kasutada stringide vastavuse kontrollimiseks, otsimiseks või asendamiseks selle regulaaravaldise alusel.
	matches := regs.FindStringSubmatch(line)
// regs on regulaaravaldise objekt, mis on loodud regexp.MustCompile abil. See määrab kindla mustri, mida otsitakse stringist.
	// FindStringSubmatch on meetod, mis otsib vastet sellele mustrile antud stringis (line).
	// Kui muster leitakse, tagastab FindStringSubmatch viitaja, mis sisaldab vastet ja kõiki alamvasteid (submatches). 
	// Need alamvasted on osad stringist, mis vastavad regulaaravaldise gruppidel määratud osadele.
		// Näiteks, kui line on "2021-03-15" ja regulaaravaldis on (\d{4})-(\d{2})-(\d{2}), 
		// siis FindStringSubmatch tagastab viitaja ["2021-03-15", "2021", "03", "15"], 
		// kus esimene element on kogu vaste ja ülejäänud on alamvasted.
	if len(matches) > 1 {
		parsedTime, err := time.Parse("2006-01-02T15:04Z07:00", matches[1]) // Formaat "2006-01-02T15:04Z07:00" on Go keeles eeldefineeritud kuupäeva ja kellaaja esitus, 
																	// mida kasutatakse kui näidist, et määrata, kuidas sisendstringi kuupäev ja kellaaeg tuleks tõlgendada.
		// matches[1] viitab sellele, et kasutatakse esimest alamvastet (mitte kogu vastet, mis on matches[0]), mis saadi FindStringSubmatch funktsiooniga. 
		// Valik sõltub regulaaravaldise gruppide arvust ja nende tähtsusest teie loogikas.
		if err != nil {
			return "\033[31m"+line+"\033[0m" // Return original line if date/time is invalid (in red)
		}

		var parsedDate string
		// Formats the date or time according to the prefix or suffix.
		if strings.HasPrefix(line, "D") {
			parsedDate = "\033[1m\033[44m"+parsedTime.Format(" 02 Jan 2006 ")+"\033[0m\033[22m"
		} else if strings.HasPrefix(line, "T12") {
			parsedDate = "\033[32m"+parsedTime.Format("03:04PM (-07:00)")+"\033[0m"
		} else if strings.HasPrefix(line, "T24") {
			parsedDate = "\033[32m"+parsedTime.Format("15:04 (-07:00)")+"\033[0m"
		} else {
			return "\033[31m"+line+"\033[0m" // Return original line if wrong prefix
		}
		if strings.HasSuffix(line, "Z" ) {
			parsedDate += " (+00:00)"
		} 
		return parsedDate 
	} 
	return line // Returns original line if no pattern found
		// see variant toob outputi anomaaliad:
		// return "\033[31m"+line+"\033[0m" // Returns original line if no pattern found
		/* 
	Funktsiooni formatDateAndTime kasutamine toimub selles koodis funktsiooni processItinerary raames. Siin on üldine protsess:

1 - Funktsioon processItinerary loeb faili rea haaval, kasutades bufio.Scanner objekti.
2 - Iga loetud rea puhul kutsutakse välja formatDateAndTime funktsioon, edastades sellele praeguse rea (line) sisu.
3 - formatDateAndTime töötleb antud rea, tuvastades ja teisendades seal olevad kuupäeva ja kellaaja mustrid. Seejärel tagastab funktsioon töödeldud rea.
4 - processItinerary kogub kõik töödeldud read kokku ja loob neist väljundi, mida kasutatakse edasistes sammudes.

Seega, formatDateAndTime ei loe ise faili, vaid seda kasutatakse igale reale, mis on loetud processItinerary funktsioonis.
*/	
}


// Checks the validity of input, output, and airport lookup file as well as usage of '-h' flag.
// This function checks that the input and output files, as well as the airport lookup file,
// do not overlap and are correctly provided. It also handles the '-h' flag to display
// usage instructions if requested by the user. The function returns 'true' if all checks
// are successful, otherwise 'false', and the usage instructions are displayed.
func checkingErrors(inputTxt, outputTxt, airportLookupCsv string) bool { // eeldab 3 string tüüpi parameetrit ja tagastab bool-tüüpi väärtuse (true, false)
	// Funktsiooni ülesanne on tavaliselt kontrollida, kas etteantud sisendparameetrid on kehtivad või vastavad teatud tingimustele. 
	// Näiteks võib see kontrollida, et sisendi- ja väljundfaili nimed ei oleks samad või et kõik vajalikud failinimed on antud. 
	// Kui kõik kontrollid on edukad (st ei esine vigu), tagastab funktsioon tõenäoliselt true; vastasel juhul false.
	/*
	flag.Bool loob uue bool-tüüpi lipukese nimega "h".
Teine argument (false) on lipukese vaikimisi väärtus, mis on kasutusel siis, kui lipukest käsurealt ei määrata.
Kolmas argument ("Displays the usage") on lipukese kirjeldus, mis kuvatakse, näiteks kui käivitate programmi käsurealipuga --help
 */
	// checking errors
	helpFlag := flag.Bool("h", false, "Displays the usage")
	// Seda lähenemist kasutatakse sageli, kui soovitakse programmile käsurealt paindlikult argumente edastada. See on eriti kasulik, kui programm vajab erinevaid sisendparameetreid või 
	// kui on vajalik teatud käsurea lipukeste olemasolu kontrollimine.
	// Seda ei pruugi olla hea kasutada lihtsamate programmide puhul, kus käsurealt argumentide edastamine pole vajalik või kui programm on piisavalt lihtne, 
	// et mitte vajada käsurealipukeste kasutamist.
	flag.Parse() // See rida käivitab käsurealipukeste tegeliku töötlemise. See analüüsib käsurealt sisestatud argumente ja määrab lipukestele väärtused.

	if *helpFlag {
		displayTheUsage()
		return false
	}
	/*
	kontrollitakse, kas käsurea lipuke helpFlag on seatud true-ks. Lipuke helpFlag on defineeritud kui viide 
		(*flag.Bool), 
	seega kasutatakse selle väärtuse kontrollimiseks tärni (*), et saada viidatava muutuja tegelik väärtus.

	Kui 
		*helpFlag 
	on true (st käsureal on määratud lipp -h), siis kutsutakse välja funktsioon displayTheUsage(), mis tavaliselt kuvab programmi kasutamise juhiseid. 
	Pärast selle funktsiooni käivitamist tagastab checkingErrors funktsioon väärtuse false, mis näitab, et programm peaks oma töö lõpetama 
	(tavaliselt kuna kasutaja soovis näha ainult abiinfot, mitte käivitada programmi peamist funktsionaalsust).
 */
	if inputTxt == outputTxt || inputTxt == airportLookupCsv || outputTxt == airportLookupCsv {
		displayTheUsage()
		return false
	}
	return true // return true ütleb, et kõik eelnevad kontrollid on edukad ja programm võib jätkata oma põhifunktsioonidega.
}

// displayTheUsage displays the usage instructions for correctly running the program.
// This function outputs explanatory examples and guidelines that help users understand
// how to use the program, including the purpose and required format of each input parameter.
func displayTheUsage() {
	fmt.Println("\033[31mitinerary usage:\033[0m")
	fmt.Println("\033[41m go run . ./input.txt ./output.txt ./airport-lookup.csv \033[0m")
}
