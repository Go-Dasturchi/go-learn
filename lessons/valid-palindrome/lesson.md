# Valid palindrome

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Tushuntirish: "amanaplanacanalpanama" - palindrom.
```

## TASK

Barcha katta harflarni kichigiga o'zgartirib va faqat alifbo-raqam (harf va raqam) belgilarini qoldirgandan so'ng satr oldinga ham, orqaga ham bir xil o'qilsa, u palindrom hisoblanadi.

Sizga `s` satri berilgan. U palindrom bo'lsa `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Yangi tozalangan satr yaratish shart emas — ikki ko'rsatkichni to'g'ridan-to'g'ri asl satrning ikki chetidan yurgizib, alifbo-raqam bo'lmagan belgilarni o'tkazib yuborish mumkin.
2. Bitta ko'rsatkichni boshidan, ikkinchisini oxiridan yurgizing; ikkalasi ham alifbo-raqam belgiga yetguncha suring, so'ng kichik harfga aylantirib ularni solishtiring — mos kelmasa `false`, ikkalasi kesishguncha davom etsa `true`.
