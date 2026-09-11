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

1. Foizni hisoblash uchun avval `letter` nechta marta uchraganini sanash kifoya.
2. `s` ichida `letter` necha marta uchrashini sanang (masalan `strings.Count`), so'ng shu sonni `100` ga ko'paytirib satr uzunligiga butun songa bo'ling — Go dagi butun sonlar bo'linishi avtomatik pastga yaxlitlaydi.
