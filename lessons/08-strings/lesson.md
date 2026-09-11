# 08 — Strings

## THEORY

Stringni **munchoq shodasi** (marjon) deb tasavvur qiling — har bir munchoq o'z joyida qattiq bog'langan, va uni bir marta tizganingizdan keyin, ichkaridagi bitta munchoqni joyidan sug'urib, o'rniga boshqasini solib bo'lmaydi. Go'da string ham xuddi shunday — **o'zgarmas (immutable)**: bir marta yaratilgan stringning ichini "tuzatib" bo'lmaydi, faqat undan foydalanib **yangi** string yasash mumkin.

Go'da `string` — bu faqat matn emas, balki ichida ko'plab tayyor funksiyalar bo'lgan `strings` paketi bilan birga ishlaydigan tur. `strings` paketi eng ko'p ishlatiladigan matn amallarini o'z ichiga oladi:

```go
strings.ToUpper("salom")                  // "SALOM"
strings.ToLower("SALOM")                  // "salom"
strings.Contains("salom dunyo", "dunyo")  // true
strings.HasPrefix("salom.go", "salom")    // true — shu bilan boshlanadimi
strings.HasSuffix("salom.go", ".go")      // true — shu bilan tugaydimi
strings.Split("a,b,c", ",")               // []string{"a", "b", "c"}
strings.Join([]string{"a", "b"}, "-")     // "a-b"
strings.TrimSpace("  salom  ")            // "salom"
strings.Replace("mushuk", "u", "o", 1)    // "moshuk" — faqat 1-uchrashini almashtiradi
strings.ReplaceAll("mushuk", "u", "o")    // "moshok" — barcha uchrashini almashtiradi
strings.Repeat("ab", 3)                   // "ababab"
strings.Fields("bir  ikki   uch")         // []string{"bir", "ikki", "uch"} — bo'shliqlarga qarab bo'ladi
```

**Birlashtirish (concatenation).** Ikkita stringni birlashtirish uchun `+` belgisi ishlatiladi:

```go
ism := "Ali"
salom := "Salom, " + ism + "!"
```

Diqqat qiling: bu amal `salom`ni allaqachon bor stringlarni "tuzatib" emas, balki ularning nusxasidan **yangi** string yasab hosil qiladi — chunki, yuqorida aytilganidek, stringlar o'zgarmas.

**Bir nechta qiymatni birlashtirish — `fmt.Sprintf`.** Faqat matnlarni emas, son va boshqa turdagi qiymatlarni ham matn ichiga "quyish" kerak bo'lsa, `fmt.Sprintf` qulayroq (u ekranga chiqarmaydi, balki tayyor stringni **qaytaradi**):

```go
xabar := fmt.Sprintf("Salom, %s! Sen %d yoshdasan.", "Ali", 25)
// xabar == "Salom, Ali! Sen 25 yoshdasan."
```

Eng ko'p ishlatiladigan joy egallovchilar (verb):
- `%s` — string uchun
- `%d` — butun son uchun
- `%f` — kasr son uchun (masalan `%.2f` — verguldan keyin faqat 2 xona)
- `%v` — istalgan turdagi qiymat uchun (universal, "qanday bo'lsa shunday chiqar")
- `%T` — qiymatning turini chiqarish uchun (Types darsida ko'rgan edingiz)

**Uzunlikni o'lchash va bitta belgiga murojaat qilish.** `len(matn)` — stringning **bayt** sonini qaytaradi (oddiy lotin harf-raqamlar uchun bu harflar soniga teng bo'ladi). `matn[i]` esa `i`-o'rindagi bitta **bayt**ni qaytaradi — bu Go'da alohida "harf" (`char`) turi yo'qligi bilan bog'liq; Go'da bitta Unicode belgi `rune` deb ataladi (For Loop darsida ko'rgan edingiz), va ko'p belgilar (masalan lotin harf va raqamlar) aynan bitta baytga sig'adi:

```go
matn := "salom"
fmt.Println(len(matn))     // 5
fmt.Println(matn[0])       // 115 — 's' harfining bayt (raqamli) qiymati
fmt.Println(string(matn[0])) // "s" — uni qayta stringga aylantirib ko'rsatish
```

**Solishtirish.** Ikkita string `==` bilan to'g'ridan-to'g'ri solishtiriladi — ular harfma-harf bir xil bo'lsagina `true` qaytaradi:

```go
"salom" == "salom"  // true
"Salom" == "salom"  // false — katta-kichik harf farqlanadi
```

## EXAMPLE

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	ism := "shoxrux"
	fmt.Println(strings.ToUpper(ism))
	fmt.Println(strings.Contains(ism, "rux"))

	xabar := fmt.Sprintf("Salom, %s! Ismingiz %d harfdan iborat.", ism, len(ism))
	fmt.Println(xabar)

	qismlar := strings.Split("olma,banan,uzum", ",")
	fmt.Println(qismlar)
}
```

Natija:

```
SHOXRUX
true
Salom, shoxrux! Ismingiz 7 harfdan iborat.
[olma banan uzum]
```

## TASK

`salom` nomli o'zgaruvchi berilgan (`"salom dunyo"`). Quyidagilarni bajaring:

1. `strings.ToUpper` yordamida uni katta harflarga o'tkazib, ekranga chiqaring.
2. `strings.Contains` yordamida unda `"dunyo"` so'zi bor-yo'qligini tekshirib, natijani ekranga chiqaring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `strings` paketini `import` qilishni unutmang.
2. `fmt.Println(strings.ToUpper(salom))` va `fmt.Println(strings.Contains(salom, "dunyo"))` — ikkita alohida qatorda, shu tartibda.
