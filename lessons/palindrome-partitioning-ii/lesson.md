# Palindrome Partitioning II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s = "aab"
Output: 1
Tushuntirish: "aab" satrini ["aa", "b"] qilib bo'lish uchun faqat 1 marta kesish kifoya. Barcha qismlar palindrom.
```

## TASK

Sizga `s` satri berilgan. `s` satrini shunday qismlarga bo'lingki (partition), bunda har bir qism satr palindrom bo'lsin.

Bunday bo'lishlarni amalga oshirish uchun kerak bo'ladigan **eng kam (minimum)** kesishlar sonini toping va qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masalani ikki bosqichda yeching: avval qaysi (i,j) qism-satrlar palindrom ekanini oldindan hisoblab jadval qiling, so'ngra shu jadvaldan foydalanib "minimal kesishlar soni" uchun 1D dinamik dasturlash quring.
2. isPalindrome[i][j] ni "s[i..j] palindrommi" deb belgilang: u true bo'ladi agar s[i]==s[j] va (uzunlik <=2 yoki isPalindrome[i+1][j-1] true) bo'lsa. Keyin cuts[i] ni "s[0..i] ni palindromlarga bo'lish uchun kerakli minimal kesishlar" deb belgilab, agar butun s[0..i] palindrom bo'lsa cuts[i]=0, aks holda barcha j<=i lar bo'yicha, agar s[j..i] palindrom bo'lsa cuts[i] = min(cuts[i], cuts[j-1]+1) qilib yangilang.
