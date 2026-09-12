# 19 — Environment Variables

## THEORY

Uyingizdagi termostatni tasavvur qiling — u xonaning harorati haqidagi ma'lumotni **atrof-muhitdan** oladi, uni dasturlashda "qattiq yozib qo'yish" shart emas. **Environment variable (muhit o'zgaruvchisi)** — dasturga, uning kodi ichiga yozilmagan, balki **operatsion tizim orqali, tashqaridan** beriladigan sozlama: masalan, ma'lumotlar bazasi manzili, API kaliti, yoki "debug rejimi yoqilganmi" degan bayroq.

**O'qish — `os.Getenv`:**

```go
import "os"

kalit := os.Getenv("API_KALIT")
```

Agar `API_KALIT` muhit o'zgaruvchisi o'rnatilmagan bo'lsa, `os.Getenv` xato bermaydi — "Maps" darsida ko'rgan xatti-harakatga o'xshab, shunchaki **bo'sh string** qaytaradi.

**Mavjudligini aniq bilish — `os.LookupEnv`.** Agar sizga "bo'sh string qiymat" bilan "umuman o'rnatilmagan" holatlarini **ajratish** kerak bo'lsa (masalan, kimdir ataylab bo'sh qiymat bergan bo'lishi mumkin), `os.LookupEnv` "comma ok" naqshini beradi ("Maps" va "Multiple Return Values" darslarini eslang):

```go
qiymat, mavjud := os.LookupEnv("API_KALIT")
if !mavjud {
	fmt.Println("API_KALIT o'rnatilmagan")
}
```

**Standart qiymat bilan o'qish — amaliy naqsh:**

```go
func muhitniOl(nomi, standart string) string {
	if qiymat, mavjud := os.LookupEnv(nomi); mavjud {
		return qiymat
	}
	return standart
}
```

Bu naqsh — "If/Else" darsida ko'rgan boshlang'ich amalli shartni ("If Else" darsidagi `if natija := f(); shart {}` naqshini eslang) ishlatadi: `os.LookupEnv`ning natijasi darhol shu `if` doirasida tekshiriladi.

**Test ichida muhit o'zgaruvchisini o'rnatish — `os.Setenv`.** Amaliy dasturlarda muhit o'zgaruvchisi odatda operatsion tizim orqali (masalan, terminalda `export API_KALIT=maxfiy`) o'rnatiladi. Lekin testlarda, kodni haqiqiy tashqi sozlamaga bog'liq qilmasdan sinash uchun, `os.Setenv` orqali **dastur ichidan** vaqtinchalik o'rnatish mumkin:

```go
os.Setenv("API_KALIT", "test-qiymat")
defer os.Unsetenv("API_KALIT") // "Defer" darsini eslang — test tugagach tozalab qo'yamiz
```

**Nega muhit o'zgaruvchilari foydali.** Ular kodni **o'zgartirmasdan**, turli muhitlarda (sizning kompyuteringizda, sinov serverida, ishlab chiqarish serverida) turlicha sozlash imkonini beradi — masalan, ma'lumotlar bazasi manzili har bir muhitda boshqacha bo'lishi mumkin, lekin dasturning **kodi** hamma joyda bir xil qoladi. Bu — maxfiy ma'lumotlarni (parollar, API kalitlari) kodning o'ziga "qattiq yozib qo'yish" o'rniga, xavfsizroq boshqarishning ham keng tarqalgan usuli.

## EXAMPLE

```go
package main

import (
	"fmt"
	"os"
)

func muhitniOl(nomi, standart string) string {
	if qiymat, mavjud := os.LookupEnv(nomi); mavjud {
		return qiymat
	}
	return standart
}

func main() {
	os.Setenv("SALOM_TILI", "uz")
	defer os.Unsetenv("SALOM_TILI")

	fmt.Println(muhitniOl("SALOM_TILI", "en"))
	fmt.Println(muhitniOl("MAVJUD_EMAS", "standart_qiymat"))
}
```

Natija:

```
uz
standart_qiymat
```

## TASK

`muhitniOl(nomi, standart string) string` funksiyasi berilgan. Uni `os.LookupEnv` yordamida shunday to'ldiringki, u `nomi` muhit o'zgaruvchisi o'rnatilgan bo'lsa uning qiymatini, aks holda `standart`ni qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if qiymat, mavjud := os.LookupEnv(nomi); mavjud { return qiymat }` — boshlang'ich amalli `if`.
2. `if` blokidan keyin `return standart` — agar muhit o'zgaruvchisi topilmagan bo'lsa.
