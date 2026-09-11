# 08 — Quick Sort

## THEORY

O'quvchilar sinfini bo'yi bo'yicha saralashingiz kerak deb tasavvur qiling. Siz bitta o'quvchini **mo'ljal (pivot)** sifatida tanlaysiz, va qolgan hammani ikki guruhga bo'lasiz: undan **pastroq bo'yliklar** bir tomonga, **balandroqlar** ikkinchi tomonga. Endi mo'ljal o'z **to'g'ri joyida** turibdi — undan chapdagilar hammasi undan past, o'ngdagilar hammasi undan baland. Endi shu ikki guruhning har birida xuddi shu jarayonni **qaytaring** (rekursiv ravishda, kichikroq guruhlar bilan). Bu — **quick sort** algoritmi.

**Algoritm qadamlari:**
1. Ro'yxatdan bitta elementni **pivot** (mo'ljal) sifatida tanlang (masalan, oxirgisini).
2. Qolgan barcha elementlarni ikki guruhga ajrating: pivotdan **kichik yoki teng** bo'lganlar, va pivotdan **katta** bo'lganlar.
3. Ikkala guruhni **rekursiv** ravishda saralang, so'ng ularni pivot bilan birlashtiring: `kichiklar + [pivot] + kattalar`.

```go
func tezSaralash(sonlar []int) []int {
	if len(sonlar) <= 1 {
		return sonlar // bazaviy holat
	}
	pivot := sonlar[len(sonlar)-1]
	kichiklar := []int{}
	kattalar := []int{}
	for _, son := range sonlar[:len(sonlar)-1] {
		if son <= pivot {
			kichiklar = append(kichiklar, son)
		} else {
			kattalar = append(kattalar, son)
		}
	}
	natija := tezSaralash(kichiklar)
	natija = append(natija, pivot)
	natija = append(natija, tezSaralash(kattalar)...)
	return natija
}
```

**Merge Sort bilan solishtiring.** Ikkalasi ham "bo'lib tashla va yeng" (divide and conquer) strategiyasiga asoslangan, lekin farq shu: Merge Sort **avval bo'ladi, keyin (murakkabroq) birlashtiradi**; Quick Sort esa **avval (pivot orqali) "aqlli" ravishda bo'ladi, keyin (oddiy birlashtirish bilan) qo'shadi**. Ikkalasining o'rtacha murakkabligi ham `O(n log n)` — "Merge Sort" darsida ko'rgan tushuncha — lekin Quick Sort amalda ko'pincha tezroq ishlaydi, chunki qo'shimcha xotira kamroq talab qiladi (garchi eng yomon holatda, masalan pivot doim eng yomon tanlansa, `O(n²)`gacha sekinlashishi mumkin).

**Nega pivot tanlovi muhim.** Agar pivot har doim "o'rtacha" qiymatga yaqin bo'lsa, ikkala guruh taxminan teng hajmda bo'ladi va algoritm tez ishlaydi. Agar pivot har doim "eng yomon" (masalan, allaqachon saralangan ro'yxatda oxirgi element) bo'lib chiqsa, guruhlardan biri deyarli bo'sh, ikkinchisi deyarli to'liq bo'lib qoladi — bu algoritmni sekinlashtiradi. Amaliy kutubxonalar buni oldini olish uchun ko'pincha tasodifiy yoki "median of three" kabi aqlliroq pivot tanlash strategiyalarini ishlatadi.

## EXAMPLE

```go
package main

import "fmt"

func tezSaralash(sonlar []int) []int {
	if len(sonlar) <= 1 {
		return sonlar
	}
	pivot := sonlar[len(sonlar)-1]
	kichiklar := []int{}
	kattalar := []int{}
	for _, son := range sonlar[:len(sonlar)-1] {
		if son <= pivot {
			kichiklar = append(kichiklar, son)
		} else {
			kattalar = append(kattalar, son)
		}
	}
	natija := tezSaralash(kichiklar)
	natija = append(natija, pivot)
	natija = append(natija, tezSaralash(kattalar)...)
	return natija
}

func main() {
	fmt.Println(tezSaralash([]int{5, 2, 9, 1, 7, 3}))
}
```

Natija:

```
[1 2 3 5 7 9]
```

## TASK

`tezSaralash(sonlar []int) []int` funksiyasi berilgan. Uni **quick sort** algoritmi yordamida (oxirgi elementni pivot sifatida tanlab) shunday to'ldiringki, u `sonlar`ni o'sish tartibida saralab qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bazaviy holat: `if len(sonlar) <= 1 { return sonlar }`. Pivot: `pivot := sonlar[len(sonlar)-1]`, keyin qolgan elementlarni (`sonlar[:len(sonlar)-1]`) ikkiga ajrating.
2. `natija := tezSaralash(kichiklar)`, `natija = append(natija, pivot)`, `natija = append(natija, tezSaralash(kattalar)...)`, so'ng `return natija`.
