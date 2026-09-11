# 10 — Boolean

## THEORY

`bool` turi faqat ikkita qiymat qabul qiladi: `true` yoki `false` — xuddi lampochka o'chirilgan yoki yonganidek, uchinchi holat yo'q. Bu qiymatlar odatda solishtirish natijasida hosil bo'ladi:

```go
5 > 3          // true
5 == 5         // true
5 != 3         // true (teng emas)
```

**Mantiqiy amallar** bir nechta shartni birlashtirish uchun ishlatiladi — bularni bank seyfi va pult tugmalariga o'xshatish mumkin:

- `&&` (VA) — **ikkalasi ham** `true` bo'lsagina, natija `true`. Xuddi bank seyfini ochish uchun ikkita kalitning **ikkalasi ham** birga burilishi kerakligi kabi — bittasi yetmaydi.
- `||` (YOKI) — **kamida bittasi** `true` bo'lsa, natija `true`. Xuddi televizorni ikkita pultdan **istalgan biri** bilan yoqish mumkinligi kabi — ikkalasini bosish shart emas.
- `!` (EMAS) — qiymatni teskarisiga aylantiradi. Xuddi svet tugmasini bosish kabi — yonib turgan bo'lsa o'chadi, o'chgan bo'lsa yonadi.

```go
yosh := 20
haydovchilikGuvohi := true

if yosh >= 18 && haydovchilikGuvohi {
	fmt.Println("Haydashi mumkin")
}
```

**Qisqa tutashuv (short-circuit) — muhim, foydali xususiyat.** `&&` va `||` chapdan o'nga tekshiriladi, va agar natija chap tomonning o'ziyoq ma'lum bo'lib qolsa, Go o'ng tomonni **umuman tekshirmaydi**:

```go
// agar royxat bo'sh bo'lsa, len(royxat) > 0 yolg'on bo'lgani uchun
// Go ikkinchi shartni tekshirib ham o'tirmaydi — xatolikdan qutulamiz
if len(royxat) > 0 && royxat[0] == "olma" {
	fmt.Println("Birinchisi olma")
}
```

Bu yerda, agar `royxat` bo'sh bo'lsa, `royxat[0]`ga murojaat qilish dasturni "yiqitib" yuborardi (chunki bo'sh ro'yxatda nolinchi element yo'q) — lekin `&&` chap tomon (`len(royxat) > 0`) allaqachon `false` bo'lgani uchun, Go ikkinchi shartni tekshirishga hojat qolmaydi deb, uni sinab ko'rmaydi. Xuddi shunday, `||`da ham chap tomon `true` bo'lib qolsa, o'ng tomon tekshirilmaydi.

**Muhim qoida: Go'da "yolg'onga o'xshash" qiymatlar yo'q.** Ba'zi tillarda `0`, `""` (bo'sh matn), yoki bo'sh ro'yxat avtomatik `false` deb hisoblanadi. Go'da bunday emas — shart ichida faqat **aynan** `bool` turidagi qiymat ishlatilishi mumkin:

```go
son := 0
if son { ... }       // XATO: compile bo'lmaydi, son int, bool emas
if son == 0 { ... }  // TO'G'RI: solishtirish natijasi bool
```

Bu — kodni aniqroq va tushunarliroq qiladi: har doim shart ichida nima tekshirilayotgani ochiq-oydin ko'rinadi, "bo'sh string yolg'onmi yoki `false`ning o'zimi" kabi chalkashlik bo'lmaydi.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	yosh := 20
	balandlik := 180

	fmt.Println(yosh >= 18 && balandlik >= 150) // true — ikkalasi ham rost
	fmt.Println(yosh < 18 || balandlik < 150)   // false — ikkalasi ham yolg'on

	royxat := []string{}
	// agar bo'sh ro'yxat bo'lsa, ikkinchi shart tekshirilmaydi (short-circuit)
	fmt.Println(len(royxat) > 0 && royxat[0] == "olma")

	yoshEmas := !(yosh < 18)
	fmt.Println(yoshEmas) // true — chunki yosh 18dan kichik emas
}
```

Natija:

```
true
false
false
true
```

## TASK

`mosKeladimi(yosh, ball int) bool` funksiyasi berilgan. U quyidagi shart bajarilsa `true` qaytarishi kerak:

- `yosh` kamida `18` bo'lishi **VA**
- `ball` kamida `60` bo'lishi

Aks holda `false` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `&&` amalidan foydalaning: `yosh >= 18 && ball >= 60`.
2. Butun ifodani to'g'ridan-to'g'ri qaytarish mumkin: `return yosh >= 18 && ball >= 60` — alohida `if` yozish shart emas.
