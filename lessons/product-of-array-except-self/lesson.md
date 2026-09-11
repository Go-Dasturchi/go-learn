# Product of Array Except Self

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Tushuntirish:
answer[0] = 2 * 3 * 4 = 24
answer[1] = 1 * 3 * 4 = 12
answer[2] = 1 * 2 * 4 = 8
answer[3] = 1 * 2 * 3 = 6
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Siz `answer` massivini qaytarishingiz kerak bo'lib, bunda `answer[i]` qiymati `nums[i]` bundan tashqari massivdagi barcha elementlarning ko'paytmasiga teng bo'ladi.

Yechim `O(n)` vaqt murakkabligida ishlashi va bo'linish operatorini ishlatmagan holda hal etilishi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'lishsiz yechish uchun har bir indeksdagi javobni "chapdagi barcha elementlar ko'paytmasi" va "o'ngdagi barcha elementlar ko'paytmasi"ning ko'paytmasi sifatida ikki bosqichda hisoblashni o'ylab ko'ring.
2. Avval chapdan o'ngga yurib, har bir indeksga o'zidan oldingi barcha elementlar ko'paytmasini yozib chiqing (natija massivida); so'ng o'ngdan chapga yurib, o'zgaruvchida o'ngdagi elementlar ko'paytmasini saqlab, uni natijadagi mos elementga ko'paytirib boring.
