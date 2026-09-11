# Restore IP Addresses

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: s = "25525511135"
Output: ["255.255.11.135","255.255.111.35"]
```

## TASK

Sizga faqat raqamlardan iborat `s` satri berilgan. Undan barcha mumkin bo'lgan, haqiqiy IPv4 manzillarini qaytaring. Natijani istalgan tartibda qaytarishingiz mumkin.

Haqiqiy IPv4 manzil `.` (nuqta) orqali ajratilgan 4 ta qismdan (har bir qism $0$ dan $255$ gacha bo'lgan butun son) iborat bo'ladi va ularda yetakchi nol (leading zero) bo'lishi mumkin emas. Masalan, `"0.1.2.201"` yoki `"192.168.1.1"` to'g'ri IPv4 manzillardir, ammo `"0.011.255.245"`, `"192.168.1.312"`, yoki `"192.168@1.1"` yaroqsiz hisoblanadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Satrni 4 ta segmentga bo'lish variantlarini sinab ko'radigan rekursiv qidiruv (backtracking) haqida o'ylab ko'ring — har bir segment 1 dan 3 belgigacha uzunlikda bo'lishi mumkin.
2. Joriy boshlanish pozitsiyasi va shu paytgacha yig'ilgan segmentlarni oluvchi rekursiv funksiya yozing: har chaqiriqda uzunligi 1, 2 yoki 3 bo'lgan keyingi segmentni sinab ko'ring — agar segment yetakchi noldan boshlansa (va uzunligi 1 dan katta) yoki qiymati 255 dan katta bo'lsa uni rad eting; 4 ta segment yig'ilib satr ham tugagan bo'lsa, ularni nuqta bilan birlashtirib natijaga qo'shing.
