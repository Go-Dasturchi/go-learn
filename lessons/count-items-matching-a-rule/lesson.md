# Count Items Matching a Rule

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
Output: 1
Tushuntirish: Faqat bitta element mos keladi: ["computer","silver","lenovo"].
```

## TASK

Sizga `items` nomli massiv berilgan, bunda har bir `items[i] = [typei, colori, namei]` mos ravishda elementning turi, rangi va nomini ifodalaydi.
Shuningdek, ikkita satr `ruleKey` va `ruleValue` berilgan.

Agar quyidagilardan biri to'g'ri bo'lsa, `i`-chi element qoidaga mos keladi deb hisoblanadi:

- `ruleKey == "type"` va `ruleValue == typei`
- `ruleKey == "color"` va `ruleValue == colori`
- `ruleKey == "name"` va `ruleValue == namei`

Qoidaga mos keladigan elementlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(items [][]string, ruleKey string, ruleValue string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
