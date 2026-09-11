# 19 — Arrays

## THEORY

Tuxum solinadigan **karton qutini** tasavvur qiling — unda aniq belgilangan sonli katak bor, masalan 10 tasi. Siz bu qutiga 11-tuxumni qo'shib sig'dira olmaysiz, va 3 tasini olib tashlab, "endi qutim 7 katakli" deb ham bo'lmaydi — qutining shakli ishlab chiqarilganda qat'iy belgilangan. Go'dagi **array (massiv)** aynan shunday ishlaydi: uning **o'lchami** e'lon qilingan zahoti qat'iy belgilanadi va keyin hech qachon o'zgarmaydi.

**E'lon qilish va literal (tayyor qiymatlar bilan yaratish):**

```go
var sonlar [5]int              // 5 ta int, hammasi standart qiymat — 0
narxlar := [3]int{10, 20, 30}   // 3 ta int, tayyor qiymatlar bilan
```

Diqqat qiling: `[5]int` va `[3]int` — bu **ikkita butunlay boshqa tur**, xuddi 5 katakli quti bilan 3 katakli qutining bir xil idish emasligi kabi. `[5]int` turidagi o'zgaruvchiga `[3]int` qiymat berib bo'lmaydi.

Agar elementlar sonini o'zingiz sanashni istamasangiz, `...` yordamida Go'ning o'zi sanashini so'rasa bo'ladi:

```go
kunlar := [...]string{"Dush", "Sesh", "Chor"} // uzunligi avtomatik 3 deb belgilanadi
```

**Elementga murojaat va uzunlik.** Har bir elementga indeks orqali (0dan boshlab) murojaat qilinadi, `len()` esa umumiy uzunlikni qaytaradi:

```go
fmt.Println(narxlar[0])   // 10 — birinchi element
fmt.Println(len(narxlar)) // 3
narxlar[1] = 99            // ikkinchi elementni almashtirish
```

**Muhim, keyingi darsda katta rol o'ynaydigan farq: array — bu "qiymat turi" (value type).** Bu degani, arrayni boshqa o'zgaruvchiga tenglashtirganingizda yoki funksiyaga argument sifatida berganingizda, Go uning **butun nusxasini** yaratadi — xuddi kartondagi barcha tuxumlarni boshqa bir qutiga bittalab qayta terib chiqqandek:

```go
a := [3]int{1, 2, 3}
b := a          // b — a ning TO'LIQ, mustaqil nusxasi
b[0] = 99
fmt.Println(a)  // [1 2 3] — a o'zgarmagan!
fmt.Println(b)  // [99 2 3] — faqat b o'zgargan
```

Bu — ba'zi boshqa tillardagi massivlardan farqli, muhim xususiyat. Amalda esa, aynan shu "har safar to'liq nusxa olish" xarajati tufayli, va o'lchamning qattiq belgilanganligi sababli, Go dasturchilari kundalik kodda arraydan ko'ra **slice**dan (keyingi darsda ko'ramiz) ko'proq foydalanishadi — slice xuddi shunday ro'yxat, lekin o'lchami moslashuvchan va funksiyaga berilganda to'liq nusxalanmaydi. Shunga qaramay, array — Go'ning slice tushunchasi qurilgan **poydevor**, shuning uchun uni tushunish muhim.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	narxlar := [3]int{10, 20, 30}
	fmt.Println(narxlar)      // [10 20 30]
	fmt.Println(len(narxlar)) // 3

	nusxa := narxlar
	nusxa[0] = 999
	fmt.Println(narxlar) // [10 20 30] — o'zgarmagan
	fmt.Println(nusxa)   // [999 20 30] — faqat nusxa o'zgargan

	kunlar := [...]string{"Dush", "Sesh", "Chor"}
	fmt.Println(kunlar, len(kunlar)) // [Dush Sesh Chor] 3
}
```

Natija:

```
[10 20 30]
3
[10 20 30]
[999 20 30]
[Dush Sesh Chor] 3
```

## TASK

`yigindiArray(sonlar [5]int) int` funksiyasi berilgan. Uni shunday to'ldiringki, u 5 ta elementli massivdagi barcha sonlarning yig'indisini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massiv ustida ham xuddi slice kabi `range` bilan aylanish mumkin: `for _, son := range sonlar { ... }`.
2. `natija := 0` bilan boshlang, tsikl ichida `natija += son` qiling, oxirida `return natija`.
