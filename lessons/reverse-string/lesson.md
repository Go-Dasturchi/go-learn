# Reverse String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

## TASK

Sizga belgilar massivi ko'rinishidagi `s` satri berilgan. Massiv elementlarini teskari o'zgartiring (reverse).

Siz buni massivni in-place (ya'ni qo'shimcha xotira ajratmasdan, bor-yo'g'i `O(1)` qo'shimcha xotira bilan) o'zgartirishingiz shart.

_Eslatma: Bu yerda tizim tekshiruvi uchun siz o'zgartirgan massivni qaytarishingiz talab etiladi._

Quyidagi funksiyani to'ldiring:

```go
func solve(s []string) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
