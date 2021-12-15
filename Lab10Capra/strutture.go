package main

import "fmt"

func main() {

 var d1 struct{ gg, mm, aa int } // d1 è una var di un tipo struttura che rappresenta delle date
 //all'interno delle graffe che indicano la struct ci sono i cosiddetti campi, in questo caso tutti dello
 //stesso tipo (intero).. possono essere anche di altri tipi comunque
 //ci permettono di creare gruppi di variabili logicamente collegate.
//

 fmt.Println("d1:", d1) //cosa stampa? //stampa d1: {0 0 0} --> sebbene non ci sia nessuna inizializzazione
 //specifica nella variabile d1, go gli assegna automaticamente e implicitamente il valore di default (che nel caso degli int è 0)


//se voglio accedere ai diversi membri della struttura?
 d1.gg = 15 // accesso ai membri con "dot notation" --> tipica della programmazione ad oggetti
 d1.mm = 12
 d1.aa = 2021

 fmt.Println("d1:", d1) //viene stampato d1: {15 12 2021}

 //che svantaggio ha var d1 struct{ gg, mm, aa int}? se io per esempio adesso volessi dichiarere altre
 //variabili qui dovrei aggiungere o ripetere la definizione dove mi interessa
 //QUINDI -> utilizziamo l'operatore Type che ci permette di definire degli alias per dei tipi !!!!!

 //MOLTO meglio definire un tipo di nome Data corrispondente alla struttura di cui sopra!

 type Data struct{ gg, mm, aa int } //<--- sintassi  //introduco un nome (data) e dico che da qui in poi è un nuovo tipo
 var d2, d3 Data //questo nuovo tipo ha come nome data e posso usarlo come tipo


//una volta che ho definito un tipo, posso usare dei letterali costanti di tipo struttura
 d1 = Data{15, 12, 2021} //questo è un letterale di tipo data, Data è un alias //cosa c'è di significativo in questa scrittura?
 // che sto chiamando i valori dei campi nello stesso ordine in cui si sono definiti i valori nella struttura
 //ancora più significativo è che n realtà in questo programma avvo già definito d1 (in var d1 struct)
//ma non è contradddetta nessuna regola di go percheè type Data struct è un alias di d1 quindi il compilatore non si lamenta
 fmt.Println("d1:", d1)

 d2 = Data{12, 15, 2021} // errore logico (data inesistente) ---> LOGICO(=semantico)! non da errore il compilatore, ma lo so io che 15 non è un mese
 fmt.Println("d2:", d2)

 d2 = Data{gg: 15, mm: 12, aa: 2021} // più sicuro: nominiamo i campi da inizializzare //così possiamo controllare che assegnamo
 //un certo valore a un certo campo. {campo: valore}
 //quando faccio un assegnamento come questo, vengono assegnati i diversi campi della nostra struttura
 fmt.Println("d2:", d2)

 if d1 == d2 { // possiamo confrontare due strutture (DELLO STESSO TIPO) con == !=
   // == questo operatore confronta il CONTENUTO  delle deu strutture. sono ovviamente variabili diverse, ma ne confronto il contenuto
   fmt.Println("d1 e d2 rappresentano la stessa data")
 }

 d3 = Data{gg: 1, mm: 1} // che data  è? //Data è fatta da tre parti, ma io posso anche usarne di argomento
 // mi stamperà 1 1 0
 fmt.Println("d3:", d3)

 d3 = d1 // possiamo copiare il contenuto di una struttura in un'altra usando l'assegnamento
 //d3 = d1 copia l'intero contenuto di d1 negli specifici campi di d3.
 fmt.Println("d3:", d3)

//spesso le strutture si utilizzano in congiunzione con i puntatori
//in generale quando si manipolano strutture è molto frequente manipolarle tramite puntatori perchè è più efficiente.

 ptd := &d1 // spesso si usano puntatori a strutture, ptd è un puntatore a Data //ottendo l'indirizzo della variabile d1 e lo assegno a ptd
 //la mia variabile ptd è impllicitamente di tipo *Data
 //Se p è un puntatore, *p è l'oggetto di p.
 fmt.Println("*ptd:", *ptd) //stampa 1*ptd: {15 12 2021}
 //e se volessi aumentare il giorno di 1?
 (*ptd).gg++ //<----- ++ incrementa il valore di 1
 fmt.Println("*ptd:", *ptd)
 ptd.gg++ // per riferirci ai campi della strutt. possiamo omettere l'operatore di dereferenzazione *
 fmt.Println("*ptd:", *ptd)

 // nuovo tipo "data" che include un flag per gli anni bisestili
 type DataB struct {
   gg, mm, aa int
   bis        bool //inserisco un valore booleano che mi dice se l'anno è bisestile o no
 }

 d4 := DataB{1, 1, 2020, true} //qui gli dico che l'uno gennaio 2020 è bisestile
 fmt.Println("d4:", d4)

 d5 := DataB{gg: 1, mm: 1, aa: 2020} //errore logico: qui nn ho indicato il valore booleano, e il default di bool è false
 //quindi qui mi direbbe che 2020 non è bisestile quando in realtà lo è!
 fmt.Println("d5:", d5) // inconsistente: è bisestile


 //in generale, meglio usare delle funzioni "builder" (v. seguito)


 //la funzione creaData è a riga 132
   d5, _ = CreaData(31, 4, 2020) // data inconsistente
   fmt.Println("d5:", d5)
   d5, _ = CreaData(30, 4, 2020) // data consistente
   fmt.Println("d5:", d5)

   d5, _ = xxx(100, 2021)
   fmt.Println("d5:", d5)

   //esempi d'uso di tipi composti/aggregati
   p1 := PersonaAnagrafica{nome: "paolo", cognome: "rossi", comunenascita: "vicenza",
     datanascita: DataB{gg: 21, mm: 1, aa: 1960, bis: true},
     nazione:     [2]byte{'I', 'T'}}
   fmt.Println("anno di nascita di p1:", p1.datanascita.aa, "cognome:", p1.cognome)

   stud1 := StudenteUnimi{anagrafica: &p1, cdl: []byte("F05Y"), matricola: 674769}
   fmt.Println("cognome stud1:", stud1.anagrafica.cognome)

}

func bis(aa int) bool {
 return aa%100 != 0 && aa%4 == 0 || aa%400 == 0
}

// giorniMese ritorna i giorni del mese indicato (0 se il mese non è consistente)
func giorniMese(mm, aa int) int { //dato un mese e un anno mi controlla che i valori siano utilizzabili e sensati
 if mm < 1 || mm > 12 {
   return 0 // codice di errore
 }
 giornimese := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
 bis := bis(aa)
 if bis {
   giornimese[2] = 29
 }
 return giornimese[mm]
}

// CreaData crea e ritorna una DataB effettuando tutti i controlli (l'anno può essere un qualsiasi valore)
// Ritorna DataB{} se la i valori specificati NON sono consistenti.
// Il secondo valore ritornato sarà true o false a seconda che la data specificata sia consistente.
func CreaData(gg, mm, aa int) (d DataB, ok bool) { //restituisce in realtà una copia di valori e un controllo del bisestile
 gmese := giorniMese(mm, aa)
 if gmese > 0 && gg > 0 && gg <= gmese {
   d = DataB{gg: gg, aa: aa, mm: mm, bis: bis(aa)}
   ok = true
 }
 return
}

func giorniAnno(aa int) int {
 if bis(aa) {
   return 366
 }
 return 365
}

// ESERCIZIO: cosa fa questa funzione?
func xxx(x, y int) (d DataB, ok bool) {
 if x > 0 && x <= giorniAnno(y) {
   var z int = 1
   for {
     gm := giorniMese(z, y)
     if x <= gm {
       break
     }
     x -= gm
     z++
   }
   d = DataB{gg: x, mm: z, aa: y, bis: bis(y)}
   ok = true

 }
 return
} //SOLUZIONE: dato un certo numero mi restituisce la data. tipo input: 100, 2021 che giorno/mese è? output: 10 aprile 2021

// esempi di tipi di dati "complessi" che usano Data
/* compito per le vacanze di natale:
type PersonaAnagrafica struct {
 nome, cognome, comunenascita string
 datanascita                  DataB
 nazione                      [2]byte
}

type StudenteUnimi struct {
 anagrafica *PersonaAnagrafica 
 matricola  int
 cdl        []byte
}
*/
