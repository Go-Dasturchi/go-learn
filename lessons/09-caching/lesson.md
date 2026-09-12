# 09 — Caching

## THEORY

"Redis" darsida, ma'lumotni **muddat bilan** saqlashni ko'rdik. Lekin savol qoladi: kesh **qachon** to'ldiriladi, va u bilan asosiy baza orasidagi aloqa qanday tashkil qilinadi? Bu — **caching strategy** (keshlash strategiyasi) deb ataladi.

**Eng keng tarqalgan strategiya — Cache-Aside (yoki "Lazy Loading"):**

1. Ma'lumot kerak bo'lganda, avval **keshdan** so'raladi.
2. Agar keshda **bor** bo'lsa (cache hit) — darhol qaytariladi, baza umuman **so'ralmaydi**.
3. Agar keshda **yo'q** bo'lsa (cache miss) — asosiy bazadan o'qiladi, natija **keshga yoziladi**, so'ng qaytariladi.

```go
func QiymatniOl(kesh map[string]string, baza map[string]string, kalit string) string {
	if qiymat, bor := kesh[kalit]; bor {
		fmt.Println("cache HIT")
		return qiymat
	}

	fmt.Println("cache MISS — bazadan o'qiyapmiz")
	qiymat := baza[kalit]
	kesh[kalit] = qiymat // keyingi safar uchun keshga yozamiz
	return qiymat
}
```

**Nega bu tez ishlaydi.** Birinchi so'rovda albatta "cache MISS" bo'ladi (baza so'raladi), lekin **ikkinchi va undan keyingi** so'rovlar — to'g'ridan-to'g'ri keshdan, bazaga umuman murojaat qilmasdan qaytadi. Agar bir xil ma'lumot **ko'p marta** so'ralsa (masalan, mashhur mahsulot sahifasi), bu **ulkan** tezlik farqini beradi.

**Kesh eskirishi muammosi (cache invalidation).** Agar asosiy bazadagi ma'lumot **o'zgarsa** (masalan, mahsulot narxi yangilansa), lekin keshdagi eski nusxa **o'chirilmasa** — foydalanuvchilar eski, noto'g'ri ma'lumotni ko'radi. Shuning uchun, ma'lumot yangilanganda, tegishli kesh yozuvini ham **o'chirish yoki yangilash** kerak:

```go
func QiymatniYangila(kesh map[string]string, baza map[string]string, kalit, yangiQiymat string) {
	baza[kalit] = yangiQiymat
	delete(kesh, kalit) // eski keshni bekor qilamiz — keyingi Ol() qayta yuklaydi
}
```

Dasturchilar orasida hazil bor: *"Kompyuter fanida faqat ikkita qiyin narsa bor: cache invalidation va nomlash."* Bu — hazilga o'xshab tuyulsa ham, **haqiqiy muammo**: eskirgan keshni to'g'ri boshqarish, ko'plab real xatolarning manbai hisoblanadi.

**Boshqa strategiyalar (qisqacha):**

- **Write-Through** — ma'lumot yozilganda, **darhol** ham bazaga, ham keshga yoziladi (Cache-Aside'dagi kabi "keyinroq to'ldirish" o'rniga). Bu, kesh doim **yangi** bo'lishini kafolatlaydi, lekin har bir yozish sekinroq bo'ladi.
- **Write-Back** — avval faqat keshga yoziladi, bazaga esa **keyinroq**, guruhlab yoziladi. Tezroq, lekin server qulab tushsa, yozilmagan ma'lumot yo'qolishi xavfi bor.

Bu kursda, **Cache-Aside** — eng ko'p ishlatiladigan va tushunish oson strategiya sifatida chuqurroq o'rganildi.

**Terminalda ko'rish uchun (Redis bilan):**

```bash
redis-cli GET mahsulot:1     # cache miss bo'lsa — (nil)
# dastur bazadan o'qib, Redis'ga SET qiladi
redis-cli GET mahsulot:1     # endi qiymat bor — cache hit
```

## EXAMPLE

```go
package main

import "fmt"

func QiymatniOl(kesh map[string]string, baza map[string]string, kalit string) string {
	if qiymat, bor := kesh[kalit]; bor {
		fmt.Println("cache HIT")
		return qiymat
	}
	fmt.Println("cache MISS")
	qiymat := baza[kalit]
	kesh[kalit] = qiymat
	return qiymat
}

func main() {
	baza := map[string]string{"mahsulot:1": "Noutbuk"}
	kesh := map[string]string{}

	fmt.Println(QiymatniOl(kesh, baza, "mahsulot:1")) // MISS, bazadan
	fmt.Println(QiymatniOl(kesh, baza, "mahsulot:1")) // HIT, keshdan
}
```

Natija:

```
cache MISS
Noutbuk
cache HIT
Noutbuk
```

## TASK

`QiymatniOl(kesh map[string]string, baza map[string]string, kalit string) string` funksiyasini to'ldiring:

1. Avval `kesh[kalit]` ni comma-ok bilan tekshiring. Agar mavjud bo'lsa (`"cache HIT"` chiqarib), uni qaytaring.
2. Agar mavjud bo'lmasa (`"cache MISS"` chiqarib), `baza[kalit]` ni oling, uni `kesh[kalit]` ga yozing, so'ng qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if qiymat, bor := kesh[kalit]; bor { fmt.Println("cache HIT"); return qiymat }`.
2. `fmt.Println("cache MISS")`, keyin `qiymat := baza[kalit]; kesh[kalit] = qiymat; return qiymat`.
