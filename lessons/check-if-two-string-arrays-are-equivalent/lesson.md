# Check if Two String Arrays are Equivalent

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Tushuntirish: word1 -> "abc", word2 -> "abc". Teng.
```

## TASK

Sizga ikkita satrlar massivlari `word1` va `word2` berilgan.

Agar `word1` dagi barcha satrlarni birlashtirish natijasi `word2` dagi barcha satrlarni birlashtirish natijasiga teng bo'lsa `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(word1 []string, word2 []string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikkala massivni alohida-alohida yig'ib chiqib solishtirish shart emas — ularni bitta satrga birlashtirib solishtirsa bo'ladi.
2. word1 dagi barcha elementlarni ketma-ket birlashtirib bitta satr hosil qiling, xuddi shunday word2 uchun ham qiling, so'ng ikkala natijaviy satrni tenglikka tekshiring.
