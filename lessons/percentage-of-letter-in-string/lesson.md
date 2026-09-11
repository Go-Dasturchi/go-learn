# Percentage of Letter in String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: s = "foobar", letter = "o"
Output: 33
Tushuntirish:
Satrda 6 ta belgi bor. "o" harfi 2 marta qatnashgan.
Foiz: (2 / 6) * 100 = 33.33...
Pastga yaxlitlasak 33 bo'ladi.
```

## TASK

Sizga `s` satri va `letter` belgisi (character) berilgan.

Ushbu satrda `letter` belgisi qatnashgan foiz miqdorini toping va uni butun songa yaxlitlab (pastga qarab eng yaqin butun songa yaxlitlab) qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, letter string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
