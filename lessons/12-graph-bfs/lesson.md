# 12 — Graphs and BFS

## THEORY

Shaharlar xaritasini tasavvur qiling — shaharlar (tugunlar) yo'llar (bog'lanishlar) orqali ulangan, va bitta shahardan boshqasiga borish uchun turli yo'llar bo'lishi mumkin, "daraxt"dagidek qat'iy ota-bola tartibisiz. **Graph (graf)** — aynan shunday, eng umumiy ma'lumotlar tuzilmasi: tugunlar va ular orasidagi (istalgan tartibdagi) bog'lanishlar to'plami. Aslida, "Binary Tree" va "Linked List" darslarida ko'rgan tuzilmalar — grafning maxsus, cheklangan holatlari.

**Grafni ifodalash — qo'shnilik ro'yxati (adjacency list).** Go'da graf ko'pincha `map[int][]int` orqali ifodalanadi: har bir tugun (kalit) o'zining **qo'shnilari** (bevosita ulangan tugunlar) ro'yxatiga ega:

```go
graf := map[int][]int{
	1: {2, 3},
	2: {1, 4},
	3: {1, 4},
	4: {2, 3},
}
```

Bu — "1-shahar 2 va 3-shaharlarga bevosita ulangan", "2-shahar 1 va 4-shaharlarga ulangan" va hokazo, degani.

**BFS (Breadth-First Search) — kenglik bo'yicha qidiruv.** Suvga tosh tashlaganingizda hosil bo'lgan **to'lqinlarni** tasavvur qiling — ular boshlang'ich nuqtadan barcha yo'nalishlarga, **teng masofada**, qatlam-qatlam kengayib boradi. BFS aynan shunday ishlaydi: boshlang'ich tugundan boshlab, avval **bevosita qo'shnilarni**, keyin **ularning qo'shnilarini**, va hokazo, "qatlam-qatlam" aylanib chiqadi.

**BFS uchun `Queue` kerak.** "Queue" darsida ko'rgan FIFO tuzilma — BFS uchun aynan mos, chunki u "birinchi topilgan tugun birinchi qayta ishlanishi" tartibini ta'minlaydi:

```go
func BFS(graf map[int][]int, boshlanish int) []int {
	tashrifBuyurilgan := map[int]bool{boshlanish: true}
	navbat := []int{boshlanish}
	natija := []int{}

	for len(navbat) > 0 {
		joriy := navbat[0]
		navbat = navbat[1:] // "Queue" darsidagi Dequeue mantig'i

		natija = append(natija, joriy)

		for _, qoshni := range graf[joriy] {
			if !tashrifBuyurilgan[qoshni] {
				tashrifBuyurilgan[qoshni] = true
				navbat = append(navbat, qoshni)
			}
		}
	}
	return natija
}
```

**Nega `tashrifBuyurilgan` (ziyorat qilingan tugunlar) map kerak.** Grafda, daraxtdan farqli, **davra (cycle)** bo'lishi mumkin — ya'ni bir tugundan boshlab yurib, aylanib yana o'sha tugunga qaytib kelish mumkin. Agar allaqachon ko'rilgan tugunlarni "belgilab" bormasangiz, algoritm **cheksiz aylanib** qolishi mumkin. `tashrifBuyurilgan` — har bir tugunni faqat **bir marta** navbatga qo'shilishini kafolatlaydi, "Maps" darsida ko'rgan comma-ok tekshiruviga o'xshab.

**DFS bilan farqi (qisqacha).** BFS — **Queue** (FIFO) yordamida "kenglik bo'yicha" (yaqin tugunlar avval) yuradi; DFS (Depth-First Search, chuqurlik bo'yicha qidiruv) esa **Stack** (LIFO) yoki rekursiya yordamida, bitta yo'lni **oxirigacha** ketib, keyin orqaga qaytib boshqa yo'lni sinaydi. Ikkalasi ham grafni "to'liq aylanib chiqish" uchun ishlatiladi, lekin tartib va ba'zi amaliy xususiyatlari (masalan, "eng qisqa yo'l"ni topishda BFS afzalroq) boshqacha.

## EXAMPLE

```go
package main

import "fmt"

func BFS(graf map[int][]int, boshlanish int) []int {
	tashrifBuyurilgan := map[int]bool{boshlanish: true}
	navbat := []int{boshlanish}
	natija := []int{}

	for len(navbat) > 0 {
		joriy := navbat[0]
		navbat = navbat[1:]
		natija = append(natija, joriy)

		for _, qoshni := range graf[joriy] {
			if !tashrifBuyurilgan[qoshni] {
				tashrifBuyurilgan[qoshni] = true
				navbat = append(navbat, qoshni)
			}
		}
	}
	return natija
}

func main() {
	graf := map[int][]int{
		1: {2, 3},
		2: {1, 4},
		3: {1, 4},
		4: {2, 3},
	}
	fmt.Println(BFS(graf, 1))
}
```

Natija:

```
[1 2 3 4]
```

## TASK

`BFS(graf map[int][]int, boshlanish int) []int` funksiyasi berilgan. Uni **Queue** (navbat) yordamida shunday to'ldiringki, u `boshlanish` tugunidan boshlab, grafni **kenglik bo'yicha (BFS)** aylanib chiqib, tashrif buyurilgan tugunlar tartibini `[]int` qilib qaytarsin. Har bir tugun faqat bir marta natijaga qo'shilishi kerak.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `tashrifBuyurilgan := map[int]bool{boshlanish: true}`, `navbat := []int{boshlanish}`, `natija := []int{}` bilan boshlang, keyin `for len(navbat) > 0 { ... }` tsiklini yozing.
2. Har aylanishda: `joriy := navbat[0]`, `navbat = navbat[1:]`, `natija = append(natija, joriy)`, so'ng `graf[joriy]` bo'ylab yurib, hali ziyorat qilinmagan qo'shnilarni `tashrifBuyurilgan`ga belgilab, `navbat`ga qo'shing.
