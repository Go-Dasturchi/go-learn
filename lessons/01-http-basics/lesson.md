# 01 — HTTP Basics

## THEORY

Dunyoning istalgan burchagidagi pochta bo'limi bir xil qoidalar bo'yicha ishlaydi — konvert qanday yozilishi, markaning qayerga yopishtirilishi hammasi standart, shuning uchun har qanday pochtachi, har qanday xatni tushunadi. **HTTP (HyperText Transfer Protocol)** — internet uchun aynan shunday standart: brauzer, mobil ilova, yoki boshqa dastur — qaysi tilda yozilgan bo'lishidan qat'iy nazar, bir xil qoidalar bo'yicha "so'rov" yuborib, "javob" oladi. "HTTP Handlers" darsida bu haqda qisqacha to'xtalgan edik — endi butun rasmni ko'rib chiqamiz.

**So'rov (request) — uchta asosiy qismdan iborat:**
1. **Usul (method)** — nima qilish kerakligini bildiradi (pastda batafsil).
2. **Manzil (path)** — qaysi resurs bilan ishlash kerakligini bildiradi, masalan `/kitoblar/5`.
3. **Sarlavhalar (headers) va, ba'zida, tana (body)** — qo'shimcha ma'lumot.

**Eng ko'p ishlatiladigan HTTP usullari — CRUD bilan mos keladi:**

| Usul | Ma'nosi | CRUD amal |
|---|---|---|
| `GET` | Ma'lumotni **o'qish** | Read |
| `POST` | Yangi narsa **yaratish** | Create |
| `PUT` | Mavjud narsani **to'liq almashtirish** | Update |
| `PATCH` | Mavjud narsani **qisman yangilash** | Update |
| `DELETE` | Narsani **o'chirish** | Delete |

Bu jadval — "Building a REST API" darsida ko'rgan g'oyaning poydevori: **bitta manzil** (masalan, `/kitoblar`) turli usullar orqali **turli ma'no** anglatadi.

**HTTP — "holatga ega bo'lmagan" (stateless) protokol.** Har bir so'rov, oldingi so'rovlardan **mustaqil** — server, standart holatda, "bu mijoz kim, avval nima so'ragan edi" deb eslab qolmaydi. Har bir so'rov o'zi bilan kerakli barcha ma'lumotni olib keladi (masalan, sarlavhadagi autentifikatsiya tokeni orqali — "Authentication" darsida ko'ramiz). Bu — HTTP'ni juda oddiy va serverlarni osongina ko'paytirish (scale) imkonini beruvchi qiladi, lekin "foydalanuvchi holatini" saqlash uchun qo'shimcha mexanizmlar (masalan, tokenlar, cookie'lar) kerak bo'lishining ham sababi.

**Javob (response) — status kod + sarlavhalar + tana.** "HTTP Status Codes" darsida ko'rgan kodlar ("HTTP Handlers" bo'limida ham eslatilgan edi) — javobning **umumiy natijasini** bir necha raqamda ifodalaydi: `2xx` muvaffaqiyat, `4xx` mijoz xatosi, `5xx` server xatosi.

## EXAMPLE

```go
package main

import "fmt"

func httpUsulTanla(amal string) string {
	switch amal {
	case "yaratish":
		return "POST"
	case "olish":
		return "GET"
	case "yangilash":
		return "PUT"
	case "ochirish":
		return "DELETE"
	default:
		return ""
	}
}

func main() {
	fmt.Println(httpUsulTanla("yaratish"))
	fmt.Println(httpUsulTanla("olish"))
	fmt.Println(httpUsulTanla("ochirish"))
}
```

Natija:

```
POST
GET
DELETE
```

## TASK

`httpUsulTanla(amal string) string` funksiyasi berilgan. Uni `switch` yordamida shunday to'ldiringki:

- `"yaratish"` → `"POST"`
- `"olish"` → `"GET"`
- `"yangilash"` → `"PUT"`
- `"ochirish"` → `"DELETE"`
- boshqa har qanday qiymat → `""`

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `switch amal { case "yaratish": return "POST" ... }` — "Switch" darsida ko'rgan naqsh.
2. Har bir aniq holatdan keyin `default: return ""` qo'ying — mos kelmagan holatlar uchun.
