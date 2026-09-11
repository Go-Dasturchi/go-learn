# Sliding Window Median

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [1.0,-1.0,-1.0,3.0,5.0,6.0]
Tushuntirish:
Oynalar va ularning medianalari:
[1 3 -1] -> [-1, 1, 3] -> 1
[3 -1 -3] -> [-3, -1, 3] -> -1
[-1 -3 5] -> [-3, -1, 5] -> -1
[-3 5 3] -> [-3, 3, 5] -> 3
[5 3 6] -> [3, 5, 6] -> 5
[3 6 7] -> [3, 6, 7] -> 6
```

## TASK

Sizga `nums` butun sonlar massivi va `k` oynaning o'lchami (sliding window size) berilgan. Oyna massivning boshidan boshlab oxirigacha har safar bitta qadam (bir element) o'ngga siljiydi.

Siz har bir oynadagi elementlarning medianasini topishingiz va natijada barcha medianalardan iborat massivni qaytarishingiz kerak.
Agar elementlar soni juft bo'lsa, mediana — o'rtadagi ikkita elementning o'rtachasidir.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, k int) []float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Oynani har safar boshidan saralashning hojati yo'q — o'lchami k bo'lgan saralangan massiv (yoki tuzilma) ni saqlab, oyna siljiganda faqat bitta eski elementni chiqarib, bitta yangisini kiritishni o'ylab ko'ring.
2. Boshlang'ich k ta elementni saralab oling. Har bir siljishda: chiqib ketayotgan elementni binar qidiruv (sort.SearchInts) yordamida saralangan massivda topib o'chiring, keyin kirayotgan yangi elementni ham binar qidiruv bilan to'g'ri joyiga qo'shib qo'ying (slice ichiga joylash). Har bir oyna uchun saralangan massiv o'rtasidagi (yoki ikkita o'rta elementning o'rtachasi) qiymatni medianasi sifatida qaytaring.
