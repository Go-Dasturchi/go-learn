# First Bad Version

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 5, bad = 4
Output: 4
Tushuntirish:
isBadVersion(3) -> false
isBadVersion(5) -> true
isBadVersion(4) -> true
Demak 4 - birinchi yomon versiya.
```

## TASK

Siz mahsulot menejerisiz (Product Manager) va jamoangiz bilan yangi mahsulot yaratyapsiz. Afsuski, mahsulotning oxirgi versiyasi sifat tekshiruvidan (quality check) o'ta olmadi. Har bir versiya o'zidan oldingisiga asoslanganligi sababli, bitta yomon versiyadan keyingi barcha versiyalar yomon (bad) bo'ladi.

Sizda `n` ta versiya bor `[1, 2, ..., n]`. Birinchi yomon (bad) versiyani topishingiz kerak, qaysiki o'zidan keyingi barchasini yomon qilgan.

Sizda ma'lum versiyaning yomon yoki yo'qligini tekshirib beradigan `isBadVersion(version)` funksiyasi tayyor berilgan deb tasavvur qiling. API chaqiriqlarini imkon qadar kamaytirish uchun, yechimingiz Binary Search algoritmidan foydalanishi kerak.

_Eslatma:_ Tizim testida biz `isBadVersion` funksiyasi o'rniga `bad` deb nomlangan versiyani parametrdan berib yuboramiz. Siz funksiyangiz ichida `bad` qaysi ekanligini izlashingiz kerak bo'ladi (ya'ni `version >= bad` bo'lsa u yomon versiya). Lekin `O(log n)` da ishlashi shart.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int, bad int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
