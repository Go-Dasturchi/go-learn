# Reverse Nodes in k-Group

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
Tushuntirish: Guruhlar: [1,2], [3,4], [5]. Ularni teskari qilsak: [2,1], [4,3], [5]
```

## TASK

_Eslatma: Ushbu masalada bog'langan ro'yxat o'rniga massiv (array) beriladi._

Sizga butun sonlardan iborat `head` massivi va musbat butun son `k` berilgan.
Siz massivni har `k` ta elementdan iborat guruhlarga bo'lib, har bir guruh elementlarini teskari tartibda joylashtirishingiz kerak.
Agar eng oxirida qolgan elementlar soni `k` dan kam bo'lsa, ularni o'zgarishsiz qoldiring.

Siz ro'yxatdagi qiymatlarni shunchaki o'zgartirmasdan, ularni in-place teskari qilib (dastlabki xotiradan) ishlatishingiz mumkin (massiv uchun in-place yoki yangi massiv qaytaring).

Quyidagi funksiyani to'ldiring:

```go
func solve(head []int, k int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivni k tadan bo'laklarga bo'lib, har bir to'liq (aynan k ta elementdan iborat) bo'lakni alohida teskari qiling; oxirida k tadan kam qolgan bo'lak bo'lsa, uni tegmasdan qoldiring.
2. `start` ni 0 dan boshlab har safar k ga oshirib boring, lekin faqat `start+k <= n` bo'lgandagina o'sha bo'lakni ikkita ko'rsatkich (bo'lak boshi va oxiri) bilan ichiga qarab yurib almashtiring (start+k > n bo'lgan oxirgi qoldiqni tegmasdan qoldiring).
