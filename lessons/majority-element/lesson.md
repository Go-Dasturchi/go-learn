# Majority element

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [3,2,3]
Output: 3
```

## TASK

`n` o'lchamli massiv `nums` berilgan. Eng ko'p marta uchragan (majority) elementni toping.

Majority element deb `n / 2` martadan ko'proq uchragan songa aytiladi. Massivda har doim majority element mavjud bo'ladi deb faraz qilishingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. n/2 dan ko'p uchragan element borligini alohida sanoq (map) saqlamasdan ham "ovoz berish" g'oyasi orqali topish mumkin — Boyer-Moore voting algoritmini eslang.
2. Nomzod va hisoblagich saqlang: hisoblagich 0 bo'lganda joriy elementni yangi nomzod qiling, keyingi elementlar nomzodga mos kelsa hisoblagichni oshiring, mos kelmasa kamaytiring — massiv oxirida qolgan nomzod javob bo'ladi.
