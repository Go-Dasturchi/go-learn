# 05 — Higher-Order Functions

## THEORY

Oshxonada oshpazga "mana retsept, shunga qarab pishir" desangiz — siz unga tayyor ovqat emas, balki **qanday qilishni ko'rsatuvchi qog'oz** berayapsiz, va oshpaz shu qog'ozga qarab ishlaydi. Go'da funksiyalar ham xuddi shunday "qog'oz" — ularni oddiy qiymatlar kabi o'zgaruvchiga solish, boshqa funksiyaga argument sifatida berish, yoki funksiyadan natija sifatida qaytarish mumkin. Bunday, boshqa funksiyalarni qabul qiluvchi yoki qaytaruvchi funksiyalar — **higher-order functions** (yuqori darajali funksiyalar) deyiladi.

**Funksiyani o'zgaruvchiga solish.** Funksiya — bu ham bir tur qiymat, xuddi `int` yoki `string` kabi:

```go
var amal func(int, int) int
amal = func(a, b int) int {
	return a + b
}
fmt.Println(amal(3, 4)) // 7
```

**Funksiyani boshqa funksiyaga argument sifatida berish.** Bu — higher-order function'larning eng ko'p uchraydigan ishlatilishi. Masalan, ro'yxatdagi elementlarni **"biror shartga mos keladiganlarini tanlash"** (filter) — har xil shartlar uchun alohida-alohida funksiya yozish o'rniga, shartning o'zini parametr qilib berish mumkin:

```go
func filter(sonlar []int, shart func(int) bool) []int {
	natija := []int{}
	for _, son := range sonlar {
		if shart(son) {
			natija = append(natija, son)
		}
	}
	return natija
}
```

Endi shu bitta `filter` funksiyasidan **istalgan shart** bilan foydalanish mumkin — shartning o'zi funksiya sifatida uzatiladi:

```go
juft := func(x int) bool { return x%2 == 0 }
natija := filter([]int{1, 2, 3, 4, 5, 6}, juft)
fmt.Println(natija) // [2 4 6]

katta := func(x int) bool { return x > 3 }
fmt.Println(filter([]int{1, 2, 3, 4, 5, 6}, katta)) // [4 5 6]
```

**Nega bu foydali.** `filter` funksiyasi shartning **nima ekanligini bilishga hojat yo'q** — u faqat "menga har bir elementni ber, sen `true`/`false` qaytar" deb ishlaydi. Bu — bir marta yozilgan kodni son-sanoqsiz turli vaziyatda qayta ishlatish imkonini beradi, har safar yangi funksiya yozib o'tirmasdan.

**Anonim funksiyalar.** Yuqoridagi misollarda ko'rgan `func(x int) bool { ... }` kabi, nomi yo'q, joyida yozib qo'yiladigan funksiyalar — **anonim funksiyalar** deyiladi. Ular ayniqsa higher-order function'larga argument berishda qulay, chunki alohida nom o'ylab topib, alohida joyda e'lon qilishning hojati yo'q.

## EXAMPLE

```go
package main

import "fmt"

func filter(sonlar []int, shart func(int) bool) []int {
	natija := []int{}
	for _, son := range sonlar {
		if shart(son) {
			natija = append(natija, son)
		}
	}
	return natija
}

func main() {
	sonlar := []int{1, 2, 3, 4, 5, 6, 7, 8}

	juft := func(x int) bool { return x%2 == 0 }
	fmt.Println(filter(sonlar, juft))

	katta := func(x int) bool { return x > 5 }
	fmt.Println(filter(sonlar, katta))
}
```

Natija:

```
[2 4 6 8]
[6 7 8]
```

## TASK

`filter(sonlar []int, shart func(int) bool) []int` funksiyasi berilgan. Uni shunday to'ldiringki, u `sonlar` ichidan faqat `shart(son)` `true` qaytargan elementlarni o'z ichiga olgan yangi slice qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'sh slice bilan boshlang: `natija := []int{}`, keyin `for _, son := range sonlar { ... }` bilan aylaning.
2. Har bir `son` uchun `if shart(son) { natija = append(natija, son) }` qiling, tsikldan keyin `return natija`.
