# Shortest Palindrome

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s = "aacecaaa"
Output: "aaacecaaa"
Tushuntirish: O'zi "aacecaa" qismi palindrom, shuning uchun boshiga bitta "a" qo'shish kifoya.
```

## TASK

Sizga `s` satri berilgan. Satrning faqatgina boshiga istalgancha belgi qo'shish orqali uni palindromga (boshidan ham oxiridan ham bir xil o'qiladigan satr) aylantirishingiz mumkin.

Ushbu amallarni bajarish orqali hosil bo'lishi mumkin bo'lgan **eng qisqa** palindromni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
