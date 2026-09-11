# Word Break II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
Output: ["cats and dog","cat sand dog"]
```

## TASK

Sizga `s` satri va lug'at vazifasini bajaruvchi, so'zlardan iborat bo'lgan `wordDict` massivi berilgan.

Siz `s` satrining ichiga shunday bo'shliqlar (space) qo'shishingiz kerakki, hosil bo'lgan har bir jumla to'liq `wordDict` dagi so'zlardan tashkil topgan bo'lsin. Barcha mumkin bo'lgan to'g'ri jumlalar ro'yxatini qaytaring. Natijani istalgan tartibda qaytarish mumkin.

Lug'atdagi so'zlar jumla ichida takrorlanishi mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, wordDict []string) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu backtracking (rekursiya) masalasi, lekin bir xil pozitsiyadan boshlangan qidiruv bir necha marta takrorlanib ketishi mumkin — natijalarni pozitsiya bo'yicha keshlab (memoizatsiya) qo'ying, shunda har bir pozitsiya faqat bir marta qayta ishlanadi.
2. dfs(start) funksiyasi "s[start:] dan hosil bo'ladigan barcha mumkin jumlalar ro'yxati"ni qaytarsin. Har bir end (start+1 dan len(s) gacha) uchun s[start:end] lug'atda bormi tekshiring, bo'lsa dfs(end) ni chaqirib qolgan qismning barcha variantlarini oling va joriy so'z bilan birlashtirib natijaga qo'shing (qolgan qism bo'sh bo'lsa faqat so'zning o'zini qo'shing); start==len(s) bo'lganda bazaviy holat sifatida bitta bo'sh jumla qaytaring.
