# Sudoku Solver

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
```

## TASK

Sizga qisman to'ldirilgan `9 x 9` o'lchamdagi Sudoku doskasi berilgan. Sudoku doskasidagi bo'sh kataklarni to'ldirib yechimni toping.

Yechilgan Sudoku quyidagi shartlarni qanoatlantirishi kerak:

1. Har bir qatorda `1-9` gacha raqamlar takrorlanmasdan ishtirok etishi kerak.
2. Har bir ustunda `1-9` gacha raqamlar takrorlanmasdan ishtirok etishi kerak.
3. Doskadagi har bir `3 x 3` maydonchada `1-9` gacha raqamlar takrorlanmasdan ishtirok etishi kerak.

Bo'sh kataklar `'.'` belgisi bilan ifodalangan. Har bir test keysi uchun faqat bitta yagona to'g'ri yechim bor deb hisoblashingiz mumkin.

Izoh: Aslida funksiya `in-place` bo'lishi kerak, lekin bizning tizimda doskani `[[Character]]` sifatida qaytarishingiz kerak bo'ladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(board [][]string) [][]string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu klassik backtracking (orqaga qaytish) masalasi — bo'sh katakni tanlang, unga 1 dan 9 gacha raqamlarni birma-bir sinab ko'ring, mos kelmasa qaytib boshqa raqamni sinang.
2. Har bir bo'sh katak uchun 1-9 raqamlarni sinab, mos raqam qatorida, ustunida va 3x3 quti (box) ichida takrorlanmasligini tekshiruvchi yordamchi funksiya yozing. Raqamni joylab keyingi bo'sh katak uchun rekursiv chaqiruv qiling; agar rekursiya muvaffaqiyatsiz tugasa (true qaytarmasa), joylagan raqamni bekor qilib ('.' ga qaytarib) keyingi raqamni sinang — barcha katak to'lganda true qaytaring.
