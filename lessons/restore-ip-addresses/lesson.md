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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
