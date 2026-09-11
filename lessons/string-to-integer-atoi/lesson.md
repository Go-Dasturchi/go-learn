# String to Integer (atoi)

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: s = "   -42"
Output: -42
```

## TASK

Boshlang'ich bo'sh joylarni e'tiborsiz qoldiradigan, ishorani tekshirib (agar mavjud bo'lsa), so'ngra barcha ketma-ket raqamlarni o'qib butun songa aylantiruvchi algoritmni yozing. Chegaralangan 32-bit oralig'ida (`[-2^31, 2^31 - 1]`) natijani saqlang.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Satrni belgidan-belgiga bosqichma-bosqich o'qing: avval bo'sh joylarni o'tkazib yuboring, keyin ishorani aniqlang, so'ng raqamlarni yig'ing — har bir bosqichda "keyingi belgi kutilgan turdami" deb tekshiring.
2. Raqamlarni o'qish jarayonida natijani `result*10 + raqam` tarzida yig'ib boring va har safar 32-bit chegaradan (2147483647 yoki -2147483648) oshib ketishni tekshirib, oshsa darhol chegara qiymatini qaytaring; raqam bo'lmagan belgiga yetganda o'qishni to'xtatib, ishora bilan ko'paytirilgan natijani qaytaring.
