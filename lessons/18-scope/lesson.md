# 18 — Scope

## THEORY

Uyingizni tasavvur qiling: oshxonada aytilgan gap faqat oshxonada eshitiladi, mehmonxonada aytilgan gap esa butun uy bo'ylab eshitilishi mumkin (agar ovoz baland bo'lsa). **Scope (ko'lam)** — aynan shu narsa, lekin o'zgaruvchilar uchun: bu — o'zgaruvchining "ko'rinadigan", ya'ni undan foydalanish mumkin bo'lgan kod qismi. Go'da o'zgaruvchi faqat e'lon qilingan `{}` bloki ichida (va uning ichki bloklarida — xuddi oshxonaning ichidagi javon ichida ham eshitiladigandek) mavjud bo'ladi:

```go
func misol() {
	x := 10
	if true {
		y := 20
		fmt.Println(x, y)  // ikkalasi ham ko'rinadi — x tashqi "xona"dan, y shu "xona"dan
	}
	fmt.Println(y)  // XATO: y bu yerda mavjud emas!
}
```

`y` faqat `if` blokining ichida yaratilgan, shuning uchun undan tashqarida ishlatib bo'lmaydi — bu blokdan chiqishi bilan `y` "yo'qoladi" (aniqrog'i, unga endi hech qanday nom orqali murojaat qilib bo'lmaydi).

**Har bir blok — o'z alohida "xonasi".** Bu faqat `if`ga emas, `for`, `switch`, va hatto funksiyaning o'zi hosil qiladigan `{}` blokiga ham tegishli. Funksiyaning parametrlari ham xuddi shu funksiya blokining o'zida "yashaydigan" o'zgaruvchilar hisoblanadi — ular funksiyadan tashqarida ko'rinmaydi.

**Soyalash (shadowing) — bir xil nom, ikkita alohida o'zgaruvchi.** Agar ichki blokda tashqaridagi bilan **bir xil nomli** o'zgaruvchi yaratilsa (ya'ni yana `:=` bilan), bu **yangi, mustaqil** o'zgaruvchi hosil bo'ladi, va u ichki blok doirasida tashqi o'zgaruvchini "vaqtincha yashiradi":

```go
x := 1
if true {
	x := 2          // bu YANGI, alohida x — tashqisini yashiradi
	fmt.Println(x)  // 2
}
fmt.Println(x)      // 1 — tashqi x umuman o'zgarmagan
```

Bu — chalkashlik keltirib chiqarishi mumkin bo'lgan, ammo Go'da haqiqiy va tez-tez uchraydigan holat: ikkita bir xil nomli, lekin butunlay boshqa-boshqa "tortmalar". Agar tashqi o'zgaruvchining o'zini o'zgartirmoqchi bo'lsangiz (yangisini yaratish emas), ichkarida `:=` o'rniga oddiy `=` ishlatish kerak.

**Paket darajasi — butun uy bo'ylab eshitiladigan gap.** Funksiyadan tashqarida (paket darajasida) e'lon qilingan narsalar butun faylda, hatto boshqa funksiyalarda ham ko'rinadi — masalan, `const` (07-darsda ko'rgan edingiz) va paket darajasidagi `func`lar shunday ishlaydi. Bu — deyarli barcha darslarimizda `func main()` tashqarisida yozilgan yordamchi funksiyalarni `main()` ichidan chaqira olishimizning sababi.

**Nega bu muhim: yaxshi qoida.** O'zgaruvchini **iloji boricha eng tor** blokda e'lon qiling — agar u faqat `for` tsikli ichida kerak bo'lsa, uni tsikldan tashqarida e'lon qilmang. Bu ikki sababga ko'ra foydali:
1. Kodni o'qigan odam o'zgaruvchi qayerda ishlatilishini darhol ko'radi — uni butun funksiya bo'ylab "qidirishga" hojat qolmaydi.
2. Tasodifiy xatolarni kamaytiradi — masalan, eski qiymat qolib ketib, uni yangilashni unutib qo'yish kabi xatolar tor scope'da kamroq uchraydi.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	umumiy := 100
	if umumiy > 50 {
		xabar := "Katta qiymat"
		fmt.Println(xabar, umumiy)
	}
	fmt.Println(umumiy)

	x := 1
	if true {
		x := 2 // yangi, soyalovchi x
		fmt.Println("Ichkarida:", x)
	}
	fmt.Println("Tashqarida:", x)
}
```

Natija:

```
Katta qiymat 100
100
Ichkarida: 2
Tashqarida: 1
```

## TASK

`hisoblash()` funksiyasi berilgan. U ichida `if` bloki bor, lekin ichkaridagi o'zgaruvchi tashqarida ishlatilmoqchi bo'lib, kompilyatsiya xatosiga sabab bo'lmoqda. Muammoni **scope**ni to'g'ri tushunib, kodni qayta tuzib tuzating — kerakli o'zgaruvchini `if` blokidan tashqarida e'lon qiling, shunda ikkala joyda ham ishlatilishi mumkin bo'ladi.

Kutilgan natija: `"Natija: 42"`.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `natija` o'zgaruvchisini `if` blokidan **oldin**, `hisoblash()` funksiyasining boshida e'lon qiling (`var natija int` yoki `natija := 0`).
2. `if` bloki ichida endi `:=` o'rniga oddiy `=` bilan qiymat bering (chunki o'zgaruvchi allaqachon e'lon qilingan) — aks holda yana yangi, soyalovchi (shadowing) o'zgaruvchi hosil bo'ladi.
