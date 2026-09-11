# 06 — String Formatting

## THEORY

"Strings" darsida `fmt.Sprintf`ni ko'rgan edingiz — `%s` va `%d` orqali qiymatlarni matn ichiga "quyish". Aslida `fmt` paketida bundan ko'ra ancha ko'p **joy egallovchi (verb)** bor, va ularni bilish — qiymatlarni terminalga chiroyli va aniq chiqarish uchun juda muhim. Buni turli shakldagi qoliplarga (mold) quyilgan xamir kabi tasavvur qiling — bir xil qiymat, turli qoliplarda turlicha ko'rinishga keladi.

**Eng ko'p ishlatiladigan verblar:**

```go
fmt.Sprintf("%s", "salom")     // "salom" — string
fmt.Sprintf("%d", 42)          // "42" — butun son
fmt.Sprintf("%f", 3.14)        // "3.140000" — kasr son, standart 6 xona
fmt.Sprintf("%.2f", 3.14159)   // "3.14" — kasr son, nuqtadan keyin aniq 2 xona
fmt.Sprintf("%t", true)        // "true" — bool
fmt.Sprintf("%v", 42)          // "42" — universal, istalgan turdagi qiymat uchun
fmt.Sprintf("%T", 42)          // "int" — qiymatning turini ko'rsatadi
fmt.Sprintf("%q", "salom")     // "\"salom\"" — qo'shtirnoq bilan o'ralgan string
```

**Kenglik (width) — matnni tekislash.** Raqamni belgilangan kenglikda, bo'sh joy bilan to'ldirib chiqarish mumkin — bu jadval kabi tekis ustunlar yasashda foydali:

```go
fmt.Sprintf("%5d", 42)   // "   42" — jami 5 belgi, chapdan bo'shliq bilan to'ldirilgan
fmt.Sprintf("%-5d|", 42) // "42   |" — chapga tekislangan, o'ngdan bo'shliq
fmt.Sprintf("%05d", 42)  // "00042" — nol bilan to'ldirilgan
```

**`%v` ning maxsus "kuchaytirilgan" shakllari — struct'lar uchun ayniqsa foydali.** "Structs" darsida `fmt.Println(odam)` chaqirilganda `{Ali 25}` kabi ko'rinishda chiqishini ko'rgan bo'lsangiz kerak — bu aslida `%v` verbining natijasi. Struct'ni **maydon nomlari bilan** ko'rish uchun `%+v`, hatto Go kodiga o'xshash "qayta yozib bo'ladigan" ko'rinishda ko'rish uchun `%#v` bor:

```go
type Odam struct {
	Ism  string
	Yosh int
}
ali := Odam{Ism: "Ali", Yosh: 25}

fmt.Sprintf("%v", ali)  // "{Ali 25}"
fmt.Sprintf("%+v", ali) // "{Ism:Ali Yosh:25}" — maydon nomlari bilan
fmt.Sprintf("%#v", ali) // "main.Odam{Ism:\"Ali\", Yosh:25}" — Go sintaksisiga o'xshash
```

**`Sprintf` va `Printf` farqi.** `fmt.Printf` — formatlangan matnni **darhol ekranga chiqaradi**, `fmt.Sprintf` esa uni **string sifatida qaytaradi** (ekranga chiqarmaydi) — keyin bu stringni boshqa joyda (masalan, xatolik matni sifatida, yoki faylga yozish uchun) ishlatish mumkin. Amalda, agar formatlangan matnni darhol ko'rsatish kerak bo'lsa `Printf`, keyinroq ishlatish uchun saqlab qo'yish kerak bo'lsa `Sprintf` ishlatiladi.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	fmt.Printf("%d - %s - %.2f\n", 5, "olma", 3.14159)

	type Odam struct {
		Ism  string
		Yosh int
	}
	ali := Odam{Ism: "Ali", Yosh: 25}
	fmt.Printf("%v\n", ali)
	fmt.Printf("%+v\n", ali)

	fmt.Println(fmt.Sprintf("%5d|", 7))
	fmt.Println(fmt.Sprintf("%-5d|", 7))
}
```

Natija:

```
5 - olma - 3.14
{Ali 25}
{Ism:Ali Yosh:25}
    7|
7    |
```

## TASK

`formatlash(ism string, yosh int, boy float64) string` funksiyasi berilgan. Uni shunday to'ldiringki, u aynan `"Ism: <ism>, Yosh: <yosh>, Bo'y: <boy>"` ko'rinishidagi matnni qaytarsin, `boy` qiymati **nuqtadan keyin aniq 2 xona** bilan ko'rsatilsin (masalan, `formatlash("Ali", 25, 1.7)` → `"Ism: Ali, Yosh: 25, Bo'y: 1.70"`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `fmt.Sprintf("Ism: %s, Yosh: %d, Bo'y: %.2f", ism, yosh, boy)` — uchta joy egallovchini shu tartibda ishlating.
2. `%.2f` — kasr sonni nuqtadan keyin aniq 2 xona bilan ko'rsatadi, `%f` esa standart 6 xona bilan chiqaradi — bu masalada aynan `%.2f` kerak.
