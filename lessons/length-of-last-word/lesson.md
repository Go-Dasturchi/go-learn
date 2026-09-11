# Length of last word

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "Hello World"
Output: 5
Tushuntirish: Oxirgi so'z "World", uning uzunligi 5 ga teng.
```

## TASK

So'zlar va bo'sh joylardan tashkil topgan `s` satri berilgan. Satrdagi oxirgi so'zning uzunligini qaytaring.

So'z deb faqatgina bo'sh joy bo'lmagan belgilardan tashkil topgan maksimal ketma-ketlikka aytiladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Satrni oxiridan boshlab teskari tomondan o'qishni o'ylang — avval oxiridagi bo'sh joylarni, keyin so'zning o'zini hisoblang.
2. Satr oxiridan boshlab bo'sh joylarni o'tkazib yuboring, so'ng bo'sh joy bo'lmagan belgilarni sanashni davom ettiring — birinchi bo'sh joyga (yoki satr boshiga) yetguncha.
