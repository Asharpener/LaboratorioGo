

//ORDINE COMODO DI IMPLEMENTAZIONE
/*
CheckTipo
CheckCasella
StringPezzo
InizializzaScacchiera
leggiMossa
leggi
Metti
Svuota
Muovi
MossaValida
Check<Pezzo>
*/

package main

//Tipi e funzioni base per implementare un simulatore di scacchi.

type Colore bool // alias per il colore dei pezzi
type Tipo byte   // alias per il tipo dei pezzi

// Alias per le costanti di tipo Colore
const ( //costanti magiche -> il compilatore decide qual'è il modo più opportuno per usarle
 BIANCO = false
 NERO   = true
)

// Costanti che denotano i valori 1..6, corrispondenti ai diversi pezzi,
// secondo l'ordine ordine dei relativi simboli Unicode.
// Questi ultimi sono contigui, prima ci sono i bianchi poi i neri.
const (
 NULL   = iota // 0 (serve a inizializzare la sequenza)
 RE            // 1 (primo pezzo)
 REGINA        // 2
 TORRE         // ..
 ALFIERE
 CAVALLO
 PEDONE
)


// Il valore Unicode del "primo" pezzo, il RE BIANCO, è 0x2654
const UNIREBIANCO = 0x2654

//Tipi base per giocare a scacchi.
type (
 Pezzo struct {
   tipo   Tipo //tipo del Pezzo
   colore Colore //colore del pezzo
 }

 Casella struct {
   riga byte //riga: 1..8
   col  byte //colonna: 'a'..'h'
 }

 Scacchiera map[Casella]Pezzo //mappa che ha come chiavi le caselle e come valori i pezzi
 //Scacchiera [8][8]Pezzo  // possibile definizione alternativa del tipo Scacchiera
)

// ESERCIZIO: DEFINIRE LE SEGUENTI FUNZIONI

// CheckTipo verifica che un Tipo (di pezzo) sia nel range corretto.
func CheckTipo(p Tipo) bool {
  if p>=1 || p<= 6 {
    return true
  }
 return false
}

// CheckCasella verifica che le "coordinate" di una casella siano corrette
// (si veda la definizione del tipo Casella).
func CheckCasella(c Casella) bool {
  if (c.riga >= 1 || c.riga <= 8) && (c.col >= 'a' || c.col <= 'h'){
    return true
  }
  return false
}

// StringPezzo restituisce la rapresentazione di un Pezzo (una stringa formata da 1 rune).
// Restituisce "" se il pezzo è indefinito.
func StringPezzo(p Pezzo) string {
    if !CheckTipo(p.tipo) {
       return ""
  }
  if p.colore == Bianco {
      return string(UNIREBIANCO + rune(p.tipo) -1 )
  } else {
      return string (UNIREBIANCO)
  }

  }


}

// Leggi legge il contenuto della scacchiera, ritornando una coppia formata
// dal pezzo che si trova nella casella specificata e da true.
// Se la casella è "vuota" ritorna la coppia Pezzo{},false .
// Non esegue controlli.
func Leggi(s Scacchiera, c Casella) (p Pezzo, found bool) {
 return
}

// Metti colloca un pezzo su una casella di una scacchiera,
// senza eseguire alcun controllo.
func Metti(s Scacchiera, p Pezzo, c Casella) {
}

// Svuota svuota una casella di una scacchiera,
// senza eseguire alcun controllo.
func Svuota(s Scacchiera, c Casella) {
}

// Muovi esegue (se POSSIBILE, in base alla attuale disposizione dei pezzi) la mossa "pezzo in c1 va in c2"
// su una scacchiera, assumendo che sia il turno del colore passato come argomento.
// Se la mossa è valida, ritorna una tripla di valori (src, dest, true),
// dove src e dest sono i pezzi che si trovano sulle caselle di partenza e di arrivo prima dell'esecuzione della mossa
// (se la casella di arrivo è inizialmente vuota, dest sarà uguale al valore Pezzo{}).
// Ritorna (Pezzo{}, Pezzo{}, false) se per un QUALSIASI motivo la mossa non è valida.
// Si basa su MossaValida
func Muovi(s Scacchiera, turno Colore, primamossa bool, c1, c2 Casella) (src, dest Pezzo, ok bool) {
 return
}

// MossaValida verifica che la mossa "pezzo in c1 in va in c2" su una scacchiera sia valida.
// Controlla 1) che le coordinate delle caselle siano corrette, 2) che la casella iniziale contenga
// un pezzo dello stesso colore di quello passato come argomento alla funzione,
// 3) che la casella finale sia vuota o contenga un pezzo di colore diverso dal pezzo sulla casella iniziale,
// 4) che la mossa sia valida per il pezzo specificato, in base alle regole degli scacchi
// (si faccia attenzione alle regole del pedone e del cavallo: il parametro primamossa indica se si tratta della mossa iniziale per un colore...in questo caso il cavallo non può muoversi).
// Se la mossa è valida ritorna una tripla di valori (src, dest, true), dove src e dest sono i pezzi inizialmente
// sulle caselle di partenza e di arrivo (quest'ultimo sarà uguale a Pezzo{} se la casella di arrivo è vuota).
// Ritorna (Pezzo{}, Pezzo{}, false) se la mossa non è valida.
// Delega i controlli sulla validità della mossa ai vari metodi Check<Pezzo> definiti in seguito.
func MossaValida(s Scacchiera, turno Colore, primamossa bool, c1, c2 Casella) (src, dest Pezzo, ok bool) {
 return
}

// CheckAlfiere verifica che la mossa di un alfiere da una casella iniziale a una finale di una scacchiera
// sia valida. Non controlla lo stato delle caselle iniziale e finale, assumendo che sia corretto.
func CheckAlfiere(s Scacchiera, c1, c2 Casella) (ok bool) {
 return
}

// CheckTorre verifica che la mossa di una torre da una casella iniziale a una finale di una scacchiera
// sia valida. Non controlla lo stato delle caselle iniziale e finale, assumendo che sia corretto.
func CheckTorre(s Scacchiera, c1, c2 Casella) bool {
 return false
}

// CheckRe verifica che la mossa di un Re da una casella iniziale a una finale di una scacchiera
// sia valida. Non controlla lo stato delle caselle iniziale e finale, assumendo che sia corretto.
func CheckRe(c1, c2 Casella) bool {
 return false
}

// CheckCavallo verifica che la mossa di un Cavallo da una casella iniziale a una finale di una scacchiera
// sia valida. Non controlla lo stato delle caselle iniziale e finale, assumendo che sia corretto.
func CheckCavallo(c1, c2 Casella) bool {
 return false
}

//Regina = Torre + Alfiere ...

// CheckPedone verifica che la mossa di un Pedone da una casella iniziale a una finale di una scacchiera
// sia valida. Non controlla lo stato della casella iniziale, ed il colore dell'eventuale pezzo sulla casella finale, assumendo che siano corretti.
func CheckPedone(s Scacchiera, c1, c2 Casella) bool {
 return false
}

// InizializzaScacchiera inizializza una scacchiera secondo la tradizionale disposizione dei pezzi.
// Ritorna una scacchiera inizializzata.
func InizializzaScacchiera() Scacchiera {
 //può essere utile usare l'array seguente
 //primariga := []Tipo{TORRE, CAVALLO, ALFIERE, REGINA, RE, ALFIERE, CAVALLO, TORRE}

 return Scacchiera{}
}

// leggiMossa legge da stdin una mossa che si assume sia nella forma:
// riga1 col1 riga2 col2
// ad esempio:
// 2 a 3 a
// e ritorna le corrispondenti caselle (di partenza e di arrivo).
// Non esegue controlli sull'input.
func LeggiMossa() (c1, c2 Casella) {
 return
}
