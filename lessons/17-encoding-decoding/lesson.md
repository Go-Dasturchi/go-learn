# 17 — Encoding / Decoding

## THEORY

"JSON" darsida ma'lumotni JSON formatiga **kodlash (encode)** va undan qaytarib **dekodlash (decode)** ni ko'rgan edingiz. Aslida, bu — Go standart kutubxonasidagi **`encoding/*`** oilasidagi ko'plab paketlardan faqat bittasi. Umuman, **kodlash** — ma'lumotni bir shakldan (masalan, xotiradagi baytlar) boshqa, ma'lum qoidalar bo'yicha ifodalanadigan shaklga (masalan, matn) o'tkazish; **dekodlash** — teskari yo'nalish.

**Base64 — baytlarni "matn xavfsiz" qilib ko'rsatish.** Ba'zi tizimlar (masalan, eski email protokollari, yoki URL manzillari) faqat oddiy matn (harflar, raqamlar, bir nechta belgi) bilan ishonchli ishlaydi, ixtiyoriy baytlarni (masalan, rasm fayli tarkibini) to'g'ridan-to'g'ri qabul qila olmaydi. **Base64** — istalgan baytlar ketma-ketligini, faqat 64 ta "xavfsiz" belgidan (harflar, raqamlar, `+`, `/`) iborat matnga aylantiradi:

```go
import "encoding/base64"

matn := "Salom"
kodlangan := base64.StdEncoding.EncodeToString([]byte(matn))
fmt.Println(kodlangan) // "U2Fsb20="
```

Va teskarisiga, dekodlash:

```go
baytlar, err := base64.StdEncoding.DecodeString(kodlangan)
qayta := string(baytlar) // "Salom"
```

**Hex (o'n oltilik) kodlash — baytlarni o'qish uchun qulay ko'rinishda ko'rsatish.** `encoding/hex` — baytlarni har birini ikkita o'n oltilik raqam (`0`–`9`, `a`–`f`) bilan ifodalaydi — bu ko'pincha xesh qiymatlarini, yoki past darajadagi ma'lumotni odam o'qiy oladigan shaklda ko'rsatishda ishlatiladi:

```go
import "encoding/hex"

baytlar := []byte("AB")
kodlangan := hex.EncodeToString(baytlar)
fmt.Println(kodlangan) // "4142" — 'A' = 0x41, 'B' = 0x42
```

**Umumiy naqsh — barcha `encoding/*` paketlari bir xil "shakl"da ishlaydi.** E'tibor bering: `base64.StdEncoding.EncodeToString`, `hex.EncodeToString`, va "JSON" darsida ko'rgan `json.Marshal` — barchasi bir xil g'oyaga amal qiladi: **kodlash** — Go qiymatini (yoki baytlarni) boshqa formatdagi matn/baytga aylantiradi, **dekodlash** — buning teskarisi. Bu formatlarning har birini alohida-alohida yodlashdan ko'ra, ularning **umumiy naqshini** tushunish — yangi kodlash turini (masalan, kelajakda `encoding/xml` yoki `encoding/csv`) tez o'zlashtirish imkonini beradi.

## EXAMPLE

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	matn := "Salom, Go!"
	kodlangan := base64.StdEncoding.EncodeToString([]byte(matn))
	fmt.Println(kodlangan)

	baytlar, _ := base64.StdEncoding.DecodeString(kodlangan)
	fmt.Println(string(baytlar))
}
```

Natija:

```
U2Fsb20sIEdvIQ==
Salom, Go!
```

## TASK

`base64gaOtkaz(matn string) string` funksiyasi berilgan. Uni `encoding/base64` paketi yordamida shunday to'ldiringki, u `matn`ni Base64 formatiga kodlab qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `base64.StdEncoding.EncodeToString(...)` — `[]byte` qabul qiladi, shuning uchun `matn`ni avval `[]byte(matn)` orqali aylantiring.
2. `return base64.StdEncoding.EncodeToString([]byte(matn))` — funksiya tanasi shu bitta qatordan iborat bo'lishi mumkin.
