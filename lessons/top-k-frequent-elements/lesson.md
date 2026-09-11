# Top K Frequent Elements

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va `k` soni berilgan. Ushbu massivda eng ko'p takrorlanadigan `k` ta elementni qaytaring. Natijani istalgan tartibda qaytarish mumkin.

Yechimingiz `O(n log n)` dan tezroq ishlashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, k int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval har bir sonning necha marta takrorlanganini sanab chiqing, so'ng shu chastotalar bo'yicha noyob sonlarni tartiblashni o'ylab ko'ring (to'liq O(n log n) sort ham ishlaydi, tezroq yechim uchun heap yoki bucket sort ni ko'rib chiqing).
2. Har bir sonning chastotasini map orqali hisoblang; noyob sonlar ro'yxatini chastotasi bo'yicha kamayish tartibida saralang (yoki chastotani indeks sifatida ishlatib bucket'larga joylashtiring) va eng katta chastotali dastlabki k tasini natija sifatida qaytaring.
