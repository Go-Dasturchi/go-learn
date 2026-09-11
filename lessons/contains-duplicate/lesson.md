# Contains duplicate

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3,1]
Output: true
```

## TASK

Butun sonlar massivi `nums` berilgan. Agar massivda kamida bitta son ikki yoki undan ko'p marta uchrasa, `true` qaytaring. Agar barcha elementlar har xil bo'lsa, `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir sonni ko'rgan-ko'rmaganingizni tezda bilish uchun to'plam (set/map) dan foydalanishni o'ylang.
2. Massivni bir marta aylanib, har bir sonni map'ga qo'shishdan oldin u allaqachon map'da bor-yo'qligini tekshiring — bo'lsa darhol true qaytaring.
