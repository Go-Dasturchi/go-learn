# Word Ladder

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Tushuntirish: Eng qisqa o'tish "hit" -> "hot" -> "dot" -> "dog" -> "cog" bo'lib, jami 5 ta so'zdan iborat.
```

## TASK

Sizga `beginWord` satri, `endWord` satri va `wordList` so'zlar ro'yxati berilgan. Quyidagi qoidalarga ko'ra `beginWord` dan `endWord` gacha bo'lgan eng qisqa ketma-ketlikning uzunligini toping:

1. Har bir ketma-ketlikdagi yonma-yon kelgan ikkita so'z faqat bitta harf bilan farq qilishi kerak.
2. Har bir o'zgartirilgan so'z `wordList` ichida mavjud bo'lishi kerak. `beginWord` ro'yxatda bo'lishi shart emas.

Agar bunday ketma-ketlik mavjud bo'lmasa, `0` ni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(beginWord string, endWord string, wordList []string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir so'zni "graf tuguni" deb, ulardan bir harf farq qiluvchi so'zlarni "qo'shni" deb hisoblang — shunda eng qisqa o'zgartirishlar ketma-ketligini topish oddiy BFS (kenglikni birinchi qidirish) masalasiga aylanadi.
2. beginWord dan boshlab BFS yuriting: har bir so'z uchun uning har bir pozitsiyasidagi harfni 'a' dan 'z' gacha almashtirib barcha mumkin variantlarni hosil qiling, hosil bo'lgan so'z wordList da bor va hali tashrif buyurilmagan bo'lsa uni navbatdagi qatlamga qo'shing. Har bir qatlam bosib o'tilganda qadamlar sonini oshirib boring; endWord uchraganda joriy qadamlar sonini qaytaring.
