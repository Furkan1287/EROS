package utils

import (
	"strings"
)

// Türkçe ve İngilizce tehlikeli kelimeler
var dangerousWords = map[string]bool{
	// Türkçe küfürler
	"amk": true, "aq": true, "orospu": true, "piç": true, "göt": true, "siktir": true,
	"amına": true, "ananı": true, "babani": true, "koyayım": true, "koyim": true,
	"yavşak": true, "pezevenk": true, "ibne": true, "gay": true, "lezbiyen": true,
	"kaltak": true, "sürtük": true, "fahişe": true, "kancık": true,

	// İngilizce küfürler
	"fuck": true, "shit": true, "bitch": true, "ass": true, "dick": true, "cock": true,
	"pussy": true, "cunt": true, "whore": true, "slut": true, "bastard": true,
	"motherfucker": true, "fucker": true, "dumbass": true, "idiot": true, "stupid": true,

	// Şiddet içeren kelimeler
	"öldür": true, "öldürmek": true, "katlet": true, "katliam": true, "bomba": true,
	"patlat": true, "vur": true, "çek": true, "silah": true, "bıçak": true, "bıçakla": true,
	"kill": true, "murder": true, "explode": true, "shoot": true, "gun": true,
	"knife": true, "weapon": true, "terrorist": true, "terörist": true,

	// Taciz/rahatsız etme
	"taciz": true, "rahatsız": true, "zorla": true, "tehdit": true, "şantaj": true,
	"harass": true, "threaten": true, "blackmail": true, "force": true, "rape": true,

	// Spam/reklam
	"satın al": true, "indirim": true, "kazan": true, "para": true, "dolar": true,
	"bitcoin": true, "kripto": true, "lottery": true, "winner": true, "prize": true,
	"buy now": true, "discount": true, "free": true, "money": true, "crypto": true,

	// Müstehcen içerik
	"porn": true, "sex": true, "seks": true, "çıplak": true, "nude": true, "naked": true,
	"penis": true, "vajina": true, "mastürbasyon": true, "masturbation": true,
}

// SecurityFilter - Keyword-based güvenlik filtresi
func SecurityFilter(message string) bool {
	// Mesajı küçük harfe çevir
	lowerMessage := strings.ToLower(message)

	// Kelimeleri ayır
	words := strings.Fields(lowerMessage)

	// Her kelimeyi kontrol et
	for _, word := range words {
		// Noktalama işaretlerini temizle
		cleanWord := strings.Trim(word, ".,!?;:()[]{}'\"")

		if dangerousWords[cleanWord] {
			return false // Güvensiz
		}
	}

	return true // Güvenli
}
