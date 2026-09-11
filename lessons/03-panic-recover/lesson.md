# 03 — Panic and Recover

## THEORY

Binoda yong'in signalizatsiyasi ishga tushganda hamma narsa to'xtaydi va odamlar zudlik bilan chiqishga majbur bo'ladi. Go'da **panic** aynan shunday ishlaydi — bu, "Errors" darsida ko'rgan oddiy xatolikdan farqli, dasturning **normal ishlashini darhol to'xtatadigan** favqulodda holat.

**Panic qachon yuz beradi.** Ba'zi amallar avtomatik ravishda panic keltirib chiqaradi — masalan, slice'dan chegaradan tashqari indeksga murojaat qilish ("Slice Indexing" darsida ko'rgan edingiz), nolga bo'lish, yoki `nil` pointer orqali qiymatga murojaat qilish ("Pointers" darsida ko'rgan edingiz). Bulardan tashqari, o'zingiz ham ataylab `panic(...)` chaqirishingiz mumkin:

```go
func tekshir(yosh int) {
	if yosh < 0 {
		panic("yosh manfiy bo'lishi mumkin emas")
	}
}
```

Panic chaqirilganda, funksiya darhol to'xtaydi va bu "favqulodda holat" chaqiruvchi funksiyalar zanjiri bo'ylab yuqoriga ko'tarilaveradi — agar hech kim uni "ushlab qolmasa", butun dastur qulab tushadi va terminalga xato ma'lumoti (stack trace) chiqadi.

**`recover` — yong'in signalini o'chirib, ishni davom ettirish.** Agar panic'ni "ushlab qolib", dasturni qulatmasdan davom ettirmoqchi bo'lsangiz, `recover()` funksiyasi bor — lekin u faqat **`defer` ichida** chaqirilganda ishlaydi ("Defer" darsini eslang — `defer` funksiya qanday tugashidan qat'iy nazar ishga tushadi, hatto panic tufayli tugasa ham):

```go
func xavfsizIshla() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic ushlab qolindi:", r)
		}
	}()
	panic("nimadir noto'g'ri ketdi")
}
```

Bu yerda `recover()` panic'ning "sababi" (`r`)ni qaytaradi, agar panic yuz bergan bo'lsa; aks holda `nil` qaytaradi. `defer` ichida shu tarzda tekshirib, panic yuz berganini bilib olish va dasturni "qutqarib qolish" mumkin.

**Qachon panic, qachon error ishlatiladi?** Bu — muhim, amaliy qoida: **kutilgan, oddiy muvaffaqiyatsizliklar** uchun (masalan, foydalanuvchi noto'g'ri ma'lumot kiritdi, fayl topilmadi) — `error` qaytarilishi kerak, "Errors" darsidagi kabi. **Chindan ham kutilmagan, dasturchi xatosi bo'lgan holatlar** uchun (masalan, dastur ishga tushishi uchun zarur bo'lgan konfiguratsiya butunlay yo'q) — `panic` mos keladi. Amalda, ko'p Go dasturlarida `panic` juda kam, faqat chindan ham "davom etib bo'lmaydigan" holatlarda ishlatiladi — odatiy xatoliklar deyarli har doim `error` orqali qaytariladi.

**Named return qiymatini `recover` bilan o'zgartirish.** "Named Return Values" darsida ko'rgan "aha fakti"ni eslaysizmi — `defer` named return qiymatini o'zgartira olishini? Aynan shu naqsh `recover` bilan birga juda foydali: funksiya panic qilsa ham, `defer` ichida named return qiymatini o'rnatib, funksiyani "muvaffaqiyatli tugagandek" qaytarish mumkin:

```go
func xavfsizBolish(a, b int) (natija int) {
	defer func() {
		if recover() != nil {
			natija = 0 // panic (masalan, nolga bo'lish) yuz bersa, 0 qaytaramiz
		}
	}()
	natija = a / b // b == 0 bo'lsa, bu panic qiladi
	return
}
```

## EXAMPLE

```go
package main

import "fmt"

func xavfsizBolish(a, b int) (natija int) {
	defer func() {
		if recover() != nil {
			natija = 0
		}
	}()
	natija = a / b
	return
}

func main() {
	fmt.Println(xavfsizBolish(10, 2)) // 5
	fmt.Println(xavfsizBolish(10, 0)) // 0 — panic ushlab qolindi, dastur qulamadi
	fmt.Println("Dastur davom etyapti")
}
```

Natija:

```
5
0
Dastur davom etyapti
```

## TASK

`xavfsizBolish(a, b int) (natija int)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `defer` va `recover()` yordamida, agar `a / b` panic qilsa (masalan, `b == 0` bo'lganda), `natija`ni `0` ga tenglashtirsin va dastur qulamasin.
2. Aks holda, oddiygina `a / b` ni qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Funksiya boshida `defer func() { if recover() != nil { natija = 0 } }()` yozing — bu panic yuz berganda ishga tushadi.
2. `defer`dan keyin oddiygina `natija = a / b` deb yozing, so'ng `return` — agar `b == 0` bo'lsa, shu qator panic qiladi, va yuqoridagi `defer` uni ushlab, `natija`ni 0 ga o'rnatadi.
