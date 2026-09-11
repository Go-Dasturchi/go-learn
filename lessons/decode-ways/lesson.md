# Decode Ways

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: s = "12"
Output: 2
Tushuntirish: "12" ni "AB" (1 2) yoki "L" (12) ko'rinishida o'girish mumkin.
```

## TASK

A dan Z gacha bo'lgan harflar quyidagi tarzda raqamlarga kodlangan (encode):

```
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
```

Raqamlardan iborat satrni yana harflarga o'girish (decode) uchun yuqoridagi moslikdan foydalaniladi. Bitta raqamlar ketma-ketligi turli xil tarzda o'girilishi mumkin. Masalan, `"11106"` satri quyidagicha o'girilishi mumkin:

- `"AAJF"` guruhlarga ajratsak (1 1 10 6)
- `"KJF"` guruhlarga ajratsak (11 10 6)

E'tibor bering, `"06"` ni `'F'` deb o'girib bo'lmaydi, chunki faqat `"6"` yaroqli, qo'shimcha boshidagi `"0"` hisobga olinmaydi.

Sizga faqat raqamlardan iborat `s` satri berilgan. Uni dekodlashning nechta turli usullari mavjudligini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir pozitsiyagacha bo'lgan qism-satrni necha xil usulda dekodlash mumkinligini saqlaydigan dinamik dasturlash (DP) yondashuvini o'ylab ko'ring — javob oldingi bitta yoki ikkita belgidagi natijalarga bog'liq bo'ladi.
2. dp[0]=1 dan boshlab, har bir i pozitsiyasida: agar s[i-1] '0' bo'lmasa dp[i] ga dp[i-1] ni qo'shing (bitta harf sifatida o'qish), va agar oxirgi ikki belgi 10 dan 26 gacha bo'lgan (hamda birinchisi '0' bo'lmagan) sonni tashkil qilsa dp[i] ga dp[i-2] ni ham qo'shing (ikkita harf sifatida o'qish).
