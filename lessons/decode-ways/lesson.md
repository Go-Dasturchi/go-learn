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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
