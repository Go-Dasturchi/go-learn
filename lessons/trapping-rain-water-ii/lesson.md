# Trapping Rain Water II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
Output: 4
```

## TASK

Sizga `m x n` o'lchamdagi butun sonlardan iborat `heightMap` matritsasi berilgan. Unda har bir hujayra (cell) ma'lum balandlikda joylashgan ustunni bildiradi. Yomg'ir yog'gandan so'ng bu matritsada necha birlik suv jamg'arilishi mumkinligini hisoblang.

Quyidagi funksiyani to'ldiring:

```go
func solve(heightMap [][]int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu ikki o'lchamli masala — 1D "Trapping Rain Water" dagi ikki ko'rsatkich usuli endi ishlamaydi. Buning o'rniga chegara (chetdagi) kataklardan boshlab, minimal-heap (priority queue) yordamida ichkariga qarab "eng past devor"dan boshlab kengayib boruvchi algoritmni qo'llang.
2. Barcha chegaradagi kataklarni minimal-heap ga (balandligi bilan) joylab, hammasini "tashrif buyurilgan" deb belgilang. Heapdan eng kichik balandlikdagi katakni chiqarib, uning tashrif buyurilmagan qo'shnilarini ko'rib chiqing: agar qo'shni balandligi joriy katakning "suv sathi"dan past bo'lsa, farqni suv sifatida qo'shing va qo'shnini xuddi joriy sath balandligida heapga qo'shing (chunki suv shu devor bilan tutilib turadi); aks holda qo'shnini o'z balandligi bilan heapga qo'shing va tashrif buyurilgan deb belgilang.
