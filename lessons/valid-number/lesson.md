# Valid number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "0"
Output: true
```

## TASK

`s` satri berilgan. Undagi yozuv matematikkada yaroqli bo'lgan raqamli qiymat ekanligini aniqlang.
Unda kasr, ishora (`+`, `-`) hamda eksponensial yozuvlar (`e`, `E`) qatnashishi mumkin, lekin ularning sintaksisi qat'iy qoidalarga mos kelishi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Regex yozishga urinmang — buning o'rniga bitta ko'rsatkich bilan satrni qismlarga (ixtiyoriy ishora, butun qism raqamlari, ixtiyoriy nuqta va kasr qismi, ixtiyoriy 'e'/'E' va daraja qismi) ketma-ket "iste'mol qilib" o'ting, har bir qismda nechta raqam uchraganini sanang.
2. Ko'rsatkichni: ixtiyoriy '+'/'-' dan o'tkazing, keyin nuqtagacha bo'lgan raqamlarni sanang, so'ng ixtiyoriy '.' dan keyin yana raqamlarni sanang — agar ikkala raqam sonlari ham nolga teng bo'lsa darhol false qaytaring (mantissa yo'q). Keyin ixtiyoriy 'e'/'E' va undan keyingi ixtiyoriy ishora hamda kamida bitta raqamdan iborat daraja qismini tekshiring (agar 'e' bor-u lekin daraja raqamlari nol bo'lsa false). Oxirida ko'rsatkich butun satrni to'liq bosib o'tganini (i == n) tekshiring.
