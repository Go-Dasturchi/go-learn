# 10 — File I/O

## THEORY

Daftaringizni tasavvur qiling — ba'zan uni **butunligicha** o'qiysiz (masalan, skanerlab, rasmga olib), ba'zan esa **sahifama-sahifa**, birma-bir varaqlab chiqasiz. Go'da fayllar bilan ishlash ham xuddi shunday ikki usulda bo'ladi.

**Butun faylni bir yo'la o'qish/yozish — `os` paketi.** Agar fayl unchalik katta bo'lmasa, uni to'liq xotiraga o'qib olish eng oddiy yo'l:

```go
import "os"

err := os.WriteFile("natija.txt", []byte("Salom, fayl!"), 0644)

baytlar, err := os.ReadFile("natija.txt")
matn := string(baytlar) // "Salom, fayl!"
```

`os.WriteFile` uchinchi argument (`0644`) — fayl ruxsatlari (kim o'qiy oladi, kim yoza oladi); hozircha buni "standart ruxsat" deb qabul qiling.

**Qatorma-qator o'qish — `bufio.Scanner`.** Katta fayllarni (yoki umuman, istalgan manbadan keladigan matnni) **butunligicha** emas, balki **qatorma-qator** o'qish kerak bo'lganda, `bufio.Scanner` ishlatiladi. Muhim jihati shundaki, `Scanner` fayldan ham, oddiy matndan ham (`strings.NewReader` orqali) bir xilda ishlaydi — chunki ikkalasi ham Go'ning umumiy "o'qish manbai" (`io.Reader`) tushunchasiga mos keladi:

```go
import (
	"bufio"
	"strings"
)

matn := "birinchi qator\nikkinchi qator\nuchinchi qator"
skanner := bufio.NewScanner(strings.NewReader(matn))

soni := 0
for skanner.Scan() {
	qator := skanner.Text() // joriy qatorning matni
	fmt.Println(qator)
	soni++
}
fmt.Println("Jami qatorlar:", soni)
```

`skanner.Scan()` — navbatdagi qatorni o'qiydi va, agar muvaffaqiyatli bo'lsa, `true` qaytaradi (fayl/matn tugasa `false`); `skanner.Text()` — joriy o'qilgan qatorning o'zini beradi. Bu naqsh — katta fayllarni, konfiguratsiya fayllarini, yoki foydalanuvchi kiritgan ko'p qatorli matnni qayta ishlashda juda keng qo'llaniladi.

**Nega `strings.NewReader` bilan mashq qilamiz.** Haqiqiy faylni o'qish uchun diskda fayl bo'lishi kerak — lekin `bufio.Scanner`ning o'zi faylga emas, balki **umumiy "matn manbai"** tushunchasiga asoslangani uchun, uni haqiqiy fayl o'rniga xotiradagi matn (`strings.NewReader`) bilan ham mashq qilish mumkin — natija va API bir xil, faqat manba boshqacha. Bu — kodni sinash uchun juda qulay, chunki haqiqiy fayl yaratish/o'chirishning hojati yo'q.

## EXAMPLE

```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	matn := "olma\nbanan\nuzum"
	skanner := bufio.NewScanner(strings.NewReader(matn))

	for skanner.Scan() {
		fmt.Println("Qator:", skanner.Text())
	}
}
```

Natija:

```
Qator: olma
Qator: banan
Qator: uzum
```

## TASK

`qatorlarSoni(matn string) int` funksiyasi berilgan. Uni `bufio.Scanner` yordamida shunday to'ldiringki, u berilgan `matn` ichidagi qatorlar sonini qaytarsin (masalan, `qatorlarSoni("a\nb\nc")` → `3`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `bufio.NewScanner(strings.NewReader(matn))` bilan skanner yarating, keyin `for skanner.Scan() { ... }` bilan aylaning.
2. Har bir aylanishda hisoblagichni bittaga oshiring (`soni++`), tsikldan keyin `return soni`.
