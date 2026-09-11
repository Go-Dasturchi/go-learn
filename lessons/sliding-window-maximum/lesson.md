# Sliding Window Maximum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Tushuntirish:
Oyna siljishi                Maksimum
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va oyna (window) o'lchamini bildiruvchi `k` soni berilgan. Oyna massivning eng chap tomonidan boshlab har safar bitta pog'ona o'ngga siljiydi. Siz har doim oynaning ichidagi `k` ta sonni ko'ra olasiz.

Har bir siljishda oyna ichidagi maksimal elementni qaytaruvchi massivni tuzing va qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, k int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir oynada maksimumni qaytadan qidirish o'rniga, elementlarning INDEKSLARINI saqlaydigan monotonik kamayuvchi ikki tomonlama navbat (deque) dan foydalanishni o'ylab ko'ring — bunda deque boshida doim joriy oynaning maksimal elementi indeksi turadi.
2. Har bir yangi elementni qo'shishdan oldin, deque oxiridagi barcha undan kichik yoki teng qiymatli indekslarni chiqarib tashlang (ular endi hech qachon maksimal bo'la olmaydi), so'ng joriy indeksni deque oxiriga qo'shing. Agar deque boshidagi indeks oynadan chiqib ketgan bo'lsa (indeks <= i-k), uni deque boshidan olib tashlang; i >= k-1 bo'lganda deque boshidagi qiymat joriy oynaning maksimumi bo'ladi.
