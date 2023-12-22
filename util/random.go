package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "Thequickbrownfoxjumpsoverthelazydog"

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomFloat(min,max int64) float64{
	v := min + rand.Int63n(max-min+1)
	return float64(v)
}

func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)
	for i:=0;i<n;i++{
		c:= alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomOwner() string {
	return RandomString(6);
}
func RandomBalance() float64 { 
	v:= RandomFloat(1,10000)
	return v
}
func RandomCurrency() string{
	currency := []string{"USD","BDT","EURO","PESSO"}
	n:=len(currency)
	return currency[rand.Intn(n)]
}
