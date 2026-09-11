# Keep Multiplying Found Values by Two

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [5,3,6,1,12], original = 3
Output: 24
Tushuntirish:
- 3 bor, uni 2*3 = 6 ga o'zgartiring.
- 6 bor, uni 2*6 = 12 ga o'zgartiring.
- 12 bor, uni 2*12 = 24 ga o'zgartiring.
- 24 massivda yo'q. Natija 24.
```

## TASK

Sizga butun sonlar massivi `nums` va butun son `original` berilgan.

Siz quyidagi amallarni bajarasiz:

1. Agar `original` massivda mavjud bo'lsa, `original` ni `2` ga ko'paytiring (ya'ni, `original = 2 * original`).
2. Aks holda, jarayonni to'xtating.
3. Jarayon to'xtaguncha 1-qadamni takrorlang.

Oxirida topilgan yakuniy `original` qiymatini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, original int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivda son bor-yo'qligini tez tekshirish uchun set qo'llashni o'ylang, so'ngra jarayonni takrorlang.
2. nums'ni set'ga aylantiring; original shu set'da mavjud ekan uni 2 ga ko'paytirib, set'da yo'q bo'lguncha davom eting.
