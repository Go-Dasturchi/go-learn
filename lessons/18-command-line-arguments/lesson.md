# 18 — Command Line Arguments

## THEORY

Terminalda `go-learn lesson 05-for` deb yozganingizda, `lesson` va `05-for` — dasturga terminaldan uzatilgan **buyruq qatori argumentlari (command line arguments)**. Aslida, aynan shu kursning o'zi ham (`go-learn list`, `go-learn progress` va h.k.) shu mexanizm orqali ishlaydi!

**`os.Args` — barcha argumentlar slice'i.** Go'da dastur ishga tushganda unga uzatilgan barcha "so'z"lar `os.Args` nomli `[]string`da saqlanadi:

```go
import "os"

func main() {
	fmt.Println(os.Args) // masalan: [./dastur lesson 05-for]
}
```

**Muhim: `os.Args[0]` — dasturning o'zi.** Birinchi element (`os.Args[0]`) — har doim **dastur faylining nomi/yo'li**, haqiqiy argumentlar emas. Haqiqiy, foydalanuvchi bergan argumentlar — `os.Args[1:]` ("Slices" darsida ko'rgan slicing sintaksisi):

```go
argumentlar := os.Args[1:] // ["lesson", "05-for"]
```

**Nega bu kursda `os.Args`ni to'g'ridan-to'g'ri sinash mumkin emas.** `go test` dastur ishga tushganda haqiqiy `os.Args`ni **sizning testlaringiz uchun** boshqarish imkonini bermaydi — u har doim `go test`ning o'z ichki argumentlarini ko'rsatadi. Shuning uchun, professional Go kodida odatiy amaliyot — **argumentlarni parametr sifatida qabul qiluvchi**, alohida, **sinaladigan** funksiya yozish, va uni `main()` ichida `os.Args[1:]` bilan chaqirish:

```go
func argumentlarniTahlil(args []string) map[string]string {
	// ... args ustida ishlaydi, os.Args'ga bog'liq emas ...
}

func main() {
	natija := argumentlarniTahlil(os.Args[1:]) // haqiqiy ishlatilishda shunday chaqiriladi
	fmt.Println(natija)
}
```

Bu naqsh — funksiyaning o'zini **to'liq sinaladigan** qilib qoladi (chunki testda istalgan `[]string`ni to'g'ridan-to'g'ri berish mumkin), `main()` esa faqat uni haqiqiy `os.Args` bilan "ulaydi".

**`--kalit=qiymat` uslubidagi argumentlarni tahlil qilish.** Ko'p CLI dasturlar argumentlarni `--nomi=qiymat` shaklida qabul qiladi. Buni "Strings" darsida ko'rgan `strings.Split` va `strings.TrimPrefix` bilan qo'lda tahlil qilish mumkin:

```go
qism := strings.TrimPrefix("--ism=Ali", "--") // "ism=Ali"
boluklar := strings.SplitN(qism, "=", 2)       // ["ism", "Ali"]
```

(Standart kutubxonada bundan tashqari, oddiy `-nomi qiymat` uslubidagi argumentlarni avtomatik tahlil qiluvchi tayyor `flag` paketi ham bor — lekin `--kalit=qiymat` formatini qo'lda tahlil qilish, tahlil jarayonining o'zini chuqurroq tushunish uchun foydali.)

## EXAMPLE

```go
package main

import (
	"fmt"
	"strings"
)

func argumentlarniTahlil(args []string) map[string]string {
	natija := make(map[string]string)
	for _, arg := range args {
		qism := strings.TrimPrefix(arg, "--")
		boluklar := strings.SplitN(qism, "=", 2)
		if len(boluklar) == 2 {
			natija[boluklar[0]] = boluklar[1]
		}
	}
	return natija
}

func main() {
	// Haqiqiy ishlatilishda: argumentlarniTahlil(os.Args[1:])
	fmt.Println(argumentlarniTahlil([]string{"--ism=Ali", "--yosh=25"}))
}
```

Natija:

```
map[ism:Ali yosh:25]
```

## TASK

`argumentlarniTahlil(args []string) map[string]string` funksiyasi berilgan. Uni shunday to'ldiringki, u `"--kalit=qiymat"` ko'rinishidagi har bir argumentni tahlil qilib, `map[kalit]qiymat` ko'rinishida qaytarsin. Format mos kelmagan argumentlarni (masalan, `=` belgisi bo'lmagan) e'tiborsiz qoldiring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir `arg` uchun: `qism := strings.TrimPrefix(arg, "--")`, keyin `boluklar := strings.SplitN(qism, "=", 2)`.
2. `if len(boluklar) == 2 { natija[boluklar[0]] = boluklar[1] }` — faqat `=` bilan to'g'ri bo'lingan argumentlarni qo'shing.
