# Linked List Cycle

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: next = [1,2,0,3], start = 0
Output: true
Tushuntirish: 0 -> 1 -> 2 -> 0 (sikl)
```

## TASK

_Eslatma: Bu masala odatda bog'liq ro'yxat (linked list) uchun, ammo bu yerda biz uni massiv orqali tavsif etamiz. Agar har bir indeks keyingi uzellning indeksini bildirsa (oxirgi element va sikl bo'lmasa `-1` ko'rsatadi), sikl bor-yo'qligini aniqlaymiz._

Sizga butun sonlardan iborat `next` massivi va boshlang'ich `start` indeksi berilgan. `next[i]` qiymati `i`-chi tugundan keyingi tugunga ko'rsatadi. Agar `next[i] == -1` bo'lsa u ro'yxat oxiri. Siz sikl borligini aniqlashingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(next []int, start int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Tugunlarni ketma-ket 'yurib' borib, avval tashrif buyurilgan tugunga qaytib kelinsa sikl bor deb bilish mumkin — buni saqlash uchun set ishlating.
2. start'dan boshlab next[] orqali yurib boring, har bir tashrif buyurilgan indeksni set'ga qo'shing; agar joriy indeks avval ko'rilgan bo'lsa true, -1 ga yetsangiz false qaytaring.
