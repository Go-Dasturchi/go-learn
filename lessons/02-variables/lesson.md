# 02 — Variables

## THEORY

O'zgaruvchini (variable) uydagi **yorliqlangan tortmaga** o'xshatish mumkin. Tortmaning ustiga yorliq yopishtirasiz (masalan, "ismlar"), ichiga narsa solasiz, va keyin har safar shu yorliqqa qarab, ichida nima borligini bilasiz — kerak bo'lsa, ichidagi narsani boshqasiga almashtirishingiz ham mumkin. Dasturlashda ham xuddi shunday: o'zgaruvchi — bu xotiradagi joyga qo'yilgan **nom**, uning ichiga qiymat "solinadi".

**To'liq (uzun) yozuv** — tortmaning nomini ham, unga nima solinishini ham, hatto tortma qanday shakldagi narsalar uchun mo'ljallanganini ham bir vaqtda aytib beradi:

```go
var name string = "Ali"
```

- `var` — "yangi tortma yasayapman" degan kalit so'z.
- `name` — tortmaning nomi.
- `string` — bu tortma faqat **matn** uchun mo'ljallangan, degani (tur haqida keyingi darsda batafsil gaplashamiz).
- `"Ali"` — hozircha ichiga solingan qiymat.

**Qisqa yozuv** — funksiya ichida ishlaganda, tur nomini yozmasdan ham o'zgaruvchi yaratsa bo'ladi:

```go
name := "Ali"
```

`:=` belgisi Go'ga: "yangi tortma yasa, va uning qanday shakldagi ekanligini menga qiymatga qarab o'zing top" deydi. `"Ali"` qo'shtirnoq ichida bo'lgani uchun, Go buni avtomatik `string` deb belgilaydi. Amalda, deyarli har doim aynan shu qisqa yozuv ishlatiladi — u qulayroq va tezroq yoziladi.

**Muhim qoida:** `:=` faqat funksiya **ichida** ishlaydi. Funksiyadan tashqarida (paket darajasida) o'zgaruvchi yaratmoqchi bo'lsangiz, albatta `var` kerak bo'ladi — `:=` u yerda ishlamaydi.

**Qiymatni keyin o'zgartirish.** O'zgaruvchi shunchaki bir marta to'ldirilib qo'yiladigan idish emas — tortmadagi narsani istalgan payt almashtirish mumkin, faqat endi `:=` emas, oddiy `=` ishlatiladi (chunki tortma allaqachon yaratilgan, uni qayta yaratish shart emas):

```go
pul := 1000
pul = 500       // eski qiymat o'chib, yangisi yoziladi
pul = pul + 100 // hozirgi qiymatga qarab yangi qiymat hisoblash ham mumkin
```

**Bir nechta o'zgaruvchini birga yaratish.** Bir qatorda bir nechta tortmani birdaniga yasash mumkin:

```go
ism, yosh := "Vali", 30
```

Bu yuqoridagi ikkita alohida qatorga aynan teng — shunchaki qisqaroq yozilgan.

**Bo'sh tortma — standart (zero) qiymat.** Agar `var` bilan o'zgaruvchi yaratib, unga darhol qiymat bermasangiz, u **bo'sh qolmaydi** — Go unga turi bo'yicha standart qiymat beradi:

```go
var son int       // 0
var matn string    // "" (bo'sh matn)
var flag bool      // false
var kasr float64   // 0
```

Bu — Go'ning muhim xususiyati: hech qachon "aniqlanmagan, tasodifiy" qiymatli o'zgaruvchi bo'lmaydi, har doim ma'lum bir boshlang'ich holatdan boshlanadi.

**Nomlash qoidalari.** O'zgaruvchi nomi harf yoki pastki chiziqcha (`_`) bilan boshlanishi kerak, raqam bilan emas (`1son` — xato, `son1` — to'g'ri). Katta-kichik harf farqlanadi (`yosh` va `Yosh` — ikkita butunlay boshqa o'zgaruvchi). Go'ning o'zi band qilib qo'ygan so'zlarni (masalan `func`, `var`, `if`) nom sifatida ishlatib bo'lmaydi.

**Muhim ogohlantirish: ishlatilmagan o'zgaruvchi.** Go juda qattiqqo'l: agar biror o'zgaruvchi yaratib, uni keyin hech qayerda ishlatmasangiz, dastur **umuman compile bo'lmaydi** — "declared and not used" (e'lon qilingan, lekin ishlatilmagan) degan xato chiqadi. Bu — boshqa ko'plab tillarda yo'q, lekin toza kod yozishga majburlaydigan Go'ga xos qat'iy qoida.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	name := "Shoxrux"
	age := 35

	fmt.Println(name)
	fmt.Println(age)

	age = age + 1 // qiymatni qayta yozish — endi := emas, = ishlatiladi
	fmt.Println(age)

	shahar, aholi := "Toshkent", 3000000 // bir qatorda ikkita o'zgaruvchi
	fmt.Println(shahar, aholi)
}
```

Natija:

```
Shoxrux
35
36
Toshkent 3000000
```

## TASK

1. `name` nomli o'zgaruvchi yarating, qiymati `"Ali"` bo'lsin.
2. `age` nomli o'zgaruvchi yarating, qiymati `25` bo'lsin.
3. Ikkalasini `fmt.Println` orqali, ikkita alohida qatorda ekranga chiqaring (avval `name`, keyin `age`).

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `:=` belgisidan foydalaning: `name := "Ali"`.
2. Ikkita alohida `fmt.Println` chaqiruvi kerak bo'ladi — bittasi `name` uchun, bittasi `age` uchun, aynan shu tartibda.
