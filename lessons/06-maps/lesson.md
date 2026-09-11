# 24 — Maps

## THEORY

Telefon kitobingizni tasavvur qiling: ism aytasiz, u sizga mos raqamni topib beradi — ismlar tartib bilan (1, 2, 3...) emas, balki **nom orqali** qidiriladi. Go'da bunday "nom → qiymat" tuzilmasi **map** (lug'at) deb ataladi.

**Yaratish usullari.** Tayyor qiymatlar bilan (literal):

```go
yoshlar := map[string]int{
	"Ali":  25,
	"Vali": 30,
}
```

Bo'sh map yaratish uchun `make` ishlatiladi (bu — slice'dagi `make([]int, 0)`ga o'xshaydi):

```go
narxlar := make(map[string]int)
```

**Qiymat qo'shish va yangilash — bir xil sintaksis:**

```go
narxlar["non"] = 5000
narxlar["non"] = 5500 // eski qiymat ustidan yoziladi
```

**Mavjud bo'lmagan kalitni o'qish — panik emas, standart qiymat.** Bu — slice'dan (chegaradan chiqsa darhol "panic" beradi) muhim farq: agar map'da mavjud bo'lmagan kalitni o'qisangiz, Go xatolik bermaydi, balki o'sha turning **standart (zero) qiymatini** qaytaradi:

```go
narxlar := map[string]int{"non": 5000}
fmt.Println(narxlar["sut"]) // 0 — "sut" yo'q, lekin xato ham yo'q, shunchaki 0
```

Bu xususiyat amalda juda qulay — masalan, sanoq yuritishda:

```go
soni := make(map[string]int)
soni["olma"]++ // "olma" hali mavjud emas edi, lekin 0+1=1 bo'lib, avtomatik qo'shildi
soni["olma"]++ // endi 2
```

**Muammo: "kalit yo'q" bilan "kalitning qiymati 0" ni qanday farqlash mumkin?** Yuqoridagi xususiyat ba'zan chalkashlik keltirib chiqarishi mumkin — agar `narxlar["sut"]` `0` qaytarsa, bu "sut" narxi rostdan ham 0 so'mligini bildiradimi, yoki "sut" umuman ro'yxatda yo'qligini bildiradimi? Buning uchun maxsus **"comma ok"** naqshi bor — map'dan o'qishda ikkinchi (ixtiyoriy) qiymat sifatida `bool` ham olish mumkin:

```go
narx, mavjud := narxlar["sut"]
if !mavjud {
	fmt.Println("Sut ro'yxatda yo'q")
} else {
	fmt.Println("Sut narxi:", narx)
}
```

`mavjud` — agar kalit chindan ham map ichida bo'lsa `true`, aks holda `false` bo'ladi. Bu naqsh "Multiple Return Values" darsida ko'rgan `strconv.Atoi` kabi funksiyalarning ishlash tamoyiliga juda o'xshaydi.

**Kalitni o'chirish — `delete`:**

```go
delete(narxlar, "non") // "non" kalitini butunlay olib tashlaydi
```

**Iteratsiya tartibi kafolatlanmagan.** "For Loop" darsida aytilganidek, `range` bilan map ustida aylanganda, kalitlarning chiqish tartibi **har safar boshqacha** bo'lishi mumkin. Agar aniq tartib kerak bo'lsa, kalitlarni alohida slice'ga yig'ib, keyin saralash kerak (buni Slices bo'limida ko'rgan `append`/`sort` bilan qilish mumkin).

**Muhim ogohlantirish: `nil` map'ga yozish — panic!** Slice'dan farqli (u yerda nil slice'ga `append` qilish ishlaydi), **nil map'ga to'g'ridan-to'g'ri yozishga urinish dastur ishlashini to'xtatadi**:

```go
var m map[string]int // nil map
fmt.Println(m["hech narsa"]) // 0 — O'QISH xavfsiz
m["kalit"] = 5                // PANIC: yozish uchun avval make() bilan joy ajratish kerak
```

Shuning uchun, map bilan ishlashni boshlaganda, uni har doim `make(map[K]V)` yoki literal bilan **initsializatsiya qilib** olish kerak.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	yoshlar := map[string]int{"Ali": 25, "Vali": 30}
	fmt.Println(yoshlar["Ali"]) // 25

	yoshlar["Guli"] = 22 // yangi kalit qo'shish
	fmt.Println(yoshlar["Guli"])

	narx, mavjud := yoshlar["Karim"]
	fmt.Println(narx, mavjud) // 0 false — Karim yo'q

	delete(yoshlar, "Vali")
	_, mavjud = yoshlar["Vali"]
	fmt.Println(mavjud) // false — endi o'chirilgan

	soni := make(map[string]int)
	soni["olma"]++
	soni["olma"]++
	soni["banan"]++
	fmt.Println(soni["olma"], soni["banan"]) // 2 1
}
```

Natija:

```
25
22
0 false
false
2 1
```

## TASK

`sozSoni(sozlar []string) map[string]int` funksiyasi berilgan. Uni shunday to'ldiringki, u berilgan `sozlar` slice'idagi har bir so'z **necha marta uchraganini** hisoblab, natijani map ko'rinishida qaytarsin (masalan, `sozSoni([]string{"olma", "banan", "olma"})` → `map[olma:2 banan:1]`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'sh map bilan boshlang: `natija := make(map[string]int)`.
2. `for _, s := range sozlar { natija[s]++ }` — mavjud bo'lmagan so'z uchun ham bu ishlaydi, chunki yo'q kalitni o'qish standart qiymat (0) beradi.
