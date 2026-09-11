# 05 — For Loop

## THEORY

Go'da tsikl (loop, ya'ni "takrorlash") uchun faqat **bitta** kalit so'z bor — `for`. Boshqa ko'p tillarda uchraydigan alohida `while` yoki `do-while` Go'da umuman yo'q — lekin xavotir olmang, ularning barcha vazifasini `for`ning turli ko'rinishlari bajaradi. Keling, ularning har birini ketma-ket ko'rib chiqamiz.

### 1. Klassik ko'rinish — sanoq bilan takrorlash

Bu — eng ko'p ishlatiladigan shakl, uchta qismdan iborat: **boshlang'ich qiymat; shart; qadam**. Buni sport maydonchasida yugurish doiralarini sanashga o'xshatish mumkin: "1-doiradan boshla; 5-doiragacha yugur; har doiradan keyin sonni bittaga oshir":

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

- `i := 1` — boshlang'ich qiymat, tsikl boshida **bir marta** bajariladi.
- `i <= 5` — shart, **har safar** tekshiriladi; `false` bo'lib qolsa, tsikl darrov to'xtaydi.
- `i++` — **har bir aylanishdan keyin** bajariladi (`i` ni 1 ga oshiradi).

### 2. Faqat shart bilan — "while" ko'rinishi

Agar sizga aniq sanoq kerak bo'lmasa, faqat "shu shart to'g'ri ekan, davom et" desangiz kifoya — boshlang'ich qiymat va qadam qismini butunlay tashlab yuborish mumkin. Bu — chойnakni qaynaguncha kutishga o'xshaydi: siz doiralarni sanamaysiz, faqat "hali qaynamadimi?" deb tekshirib turasiz:

```go
harorat := 20
for harorat < 100 {
	harorat += 10
	fmt.Println("Harorat:", harorat)
}
```

Boshqa tillarda bu xuddi `while (harorat < 100) { ... }` deb yozilar edi — Go'da esa `while` so'zining o'zi yo'q, shunchaki `for` dan keyin faqat bitta shart yoziladi.

### 3. Cheksiz tsikl va `break`

Agar hech qanday shart yozmasdan faqat `for { ... }` deb qoldirsangiz, tsikl **hech qachon o'z-o'zidan to'xtamaydi** — bu qorovul kabi: u darvoza oldida to'xtovsiz turadi, faqat kerakli voqea sodir bo'lganda o'z ishini tugatadi. Tsiklni to'xtatish uchun `break` kalit so'zi ishlatiladi:

```go
son := 1
for {
	fmt.Println(son)
	if son >= 3 {
		break // shart bajarilganda tsikldan butunlay chiqib ketamiz
	}
	son++
}
```

`break` — tsiklni **butunlay** to'xtatadi, undan keyingi kod tsikldan tashqarida davom etadi.

### 4. `continue` — joriy aylanishni tashlab, keyingisiga o'tish

`break`dan farqli, `continue` butun tsiklni to'xtatmaydi — faqat **joriy aylanishni** tashlab, darrov keyingi aylanishga o'tadi. Buni konveyer lentasida nosoz mahsulotni chetlab, qolganlarini tekshirishda davom etishga o'xshatish mumkin:

```go
for i := 1; i <= 10; i++ {
	if i%2 != 0 {
		continue // toq sonlarni chetlab o'tamiz
	}
	fmt.Println(i) // faqat juft sonlar chiqadi: 2, 4, 6, 8, 10
}
```

### 5. `range` — ro'yxat (slice) elementlarini birma-bir ko'rib chiqish

Ko'p hollarda sizga sanoq emas, balki ro'yxatdagi (Go'da bunga **slice** deyiladi, masalan `[]int{10, 20, 30}`) har bir elementni birma-bir ko'rib chiqish kerak bo'ladi — xuddi bozor ro'yxatidagi mahsulotlarni birin-ketin belgilab chiqqandek. Buning uchun `range` ishlatiladi:

```go
mevalar := []string{"olma", "banan", "uzum"}
for index, qiymat := range mevalar {
	fmt.Println(index, qiymat)
}
```

Natija:

```
0 olma
1 banan
2 uzum
```

`range` har bir aylanishda ikkita narsa beradi: **indeks** (element ro'yxatning nechinchi o'rnida turgani, `0`dan boshlanadi) va **qiymat** (elementning o'zi).

Agar sizga faqat qiymat kerak bo'lib, indeks kerak bo'lmasa, uning o'rniga pastki chiziqcha `_` qo'yiladi (bu Go'ga "bu qiymatni ataylab e'tiborsiz qoldiryapman" deyishning rasmiy usuli):

```go
for _, qiymat := range mevalar {
	fmt.Println(qiymat) // faqat: olma, banan, uzum
}
```

Aksincha, faqat indeks kerak bo'lsa, ikkinchi qiymatni umuman yozmasa ham bo'ladi:

```go
for index := range mevalar {
	fmt.Println(index) // faqat: 0, 1, 2
}
```

### 6. `range` matn (string) ustida

`range` matn ustida ham ishlaydi — har bir aylanishda harfning **pozitsiyasi** (bayt bo'yicha) va o'sha harfning o'zi (`rune` turida, ya'ni bitta Unicode belgi) qaytariladi:

```go
for i, harf := range "GO" {
	fmt.Println(i, string(harf))
}
```

Natija:

```
0 G
1 O
```

Kichik, lekin muhim bir nozik joy: agar matn ichida lotin alifbosidan tashqari belgilar (masalan, kirill harflari yoki emoji) bo'lsa, ular xotirada bir nechta bayt egallashi mumkin — shuning uchun pozitsiyalar `0, 1, 2, 3...` deb ketma-ket kelmasligi mumkin (masalan `0, 2, 4`). Bu — kelajakda chuqurroq o'rganiladigan mavzu, hozircha shuni bilib qo'ying: oddiy lotin harf va raqamlar bilan ishlaganda bu haqda umuman xavotir olish shart emas.

### 7. `range` lug'at (map) ustida

Map — bu "kalit → qiymat" juftliklarini saqlaydigan tuzilma, xuddi telefon kitobidagi "ism → raqam" yozuvlariga o'xshaydi. `range` map ustida ham ishlaydi, har safar kalit va qiymatni qaytaradi:

```go
yoshlar := map[string]int{"Ali": 25, "Vali": 30}
for ism, yosh := range yoshlar {
	fmt.Println(ism, yosh)
}
```

**Muhim ogohlantirish:** map ustida `range` yurganda, elementlarning chiqish **tartibi kafolatlanmagan** — har dasturni ishga tushirganingizda tartib boshqacha bo'lishi mumkin. Agar sizga aniq tartib kerak bo'lsa (masalan, ismlar alifbo tartibida chiqishi kerak bo'lsa), kalitlarni alohida ro'yxatga yig'ib, keyin saralash kerak bo'ladi — bu keyingi darslarda ko'rib chiqiladigan mavzu.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	// 1. Klassik — sanoq bilan
	for i := 1; i <= 3; i++ {
		fmt.Println("Klassik:", i)
	}

	// 2. Faqat shart — "while" ko'rinishi
	harorat := 70
	for harorat < 100 {
		harorat += 15
	}
	fmt.Println("Yakuniy harorat:", harorat)

	// 3. Cheksiz tsikl + break
	son := 0
	for {
		son++
		if son == 3 {
			break
		}
	}
	fmt.Println("Break bilan to'xtagan son:", son)

	// 4. range — slice ustida
	mevalar := []string{"olma", "banan", "uzum"}
	for i, m := range mevalar {
		fmt.Println(i, m)
	}
}
```

Natija:

```
Klassik: 1
Klassik: 2
Klassik: 3
Yakuniy harorat: 115
Break bilan to'xtagan son: 3
0 olma
1 banan
2 uzum
```

## TASK

Quyida `yigindi(n int) int` funksiyasi berilgan. Uni `for` tsikli yordamida shunday to'ldiringki, u `1` dan `n` gacha bo'lgan barcha sonlarning yig'indisini qaytarsin (masalan, `yigindi(5)` → `1+2+3+4+5` → `15`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Natijani saqlash uchun `for` tsiklidan oldin `natija := 0` deb boshlang'ich qiymat kiriting.
2. `for i := 1; i <= n; i++ { natija += i }` — har bir aylanishda `i` ni `natija`ga qo'shib boring, keyin `return natija` qiling.
