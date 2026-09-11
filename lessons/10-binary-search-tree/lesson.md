# 10 — Binary Search Tree

## THEORY

Kutubxonada kitoblar maxsus qoidaga ko'ra joylashtirilgan deb tasavvur qiling: har bir javon oldida turib, "kerakli kitob nomi bu yerdagi kitobdan oldinmi keyinmi" deb so'raysiz — agar oldin bo'lsa chap tarafga, keyin bo'lsa o'ng tarafga o'tasiz, va bu jarayon takrorlanadi. Har safar qidiruv maydoni **kamayib** boradi. **Binary Search Tree (BST, ikkilik qidiruv daraxti)** — "Binary Tree" darsida ko'rgan oddiy daraxtning, aynan shu qidiruv qulayligini ta'minlaydigan maxsus turi.

**BST xossasi (eng muhim qoida):** har bir tugun uchun:
- Uning **chap** qism daraxtidagi barcha qiymatlar undan **kichik**.
- Uning **o'ng** qism daraxtidagi barcha qiymatlar undan **katta**.

```
      5
     / \
    3   8
   /   / \
  1   6   9
```

Bu yerda `5`ning chap tomonida (`3`, `1`) faqat undan kichiklar, o'ng tomonida (`8`, `6`, `9`) faqat undan kattalar bor — va bu qoida **har bir** tugun uchun, daraxtning har bir sathida amal qiladi.

**Qidiruv — BST xossasidan foydalanib, "Binary Search"ga o'xshab.** "Binary Search" darsida saralangan slice ustida yarmini tashlab, tez qidirishni ko'rgan edingiz. BST'da ham xuddi shunday mantiq ishlaydi, faqat slice o'rniga daraxt tuzilmasi orqali:

```go
func qidir(tugun *Tugun, qiymat int) bool {
	if tugun == nil {
		return false // daraxt (yoki qism daraxt) tugadi, qiymat topilmadi
	}
	if tugun.Value == qiymat {
		return true
	}
	if qiymat < tugun.Value {
		return qidir(tugun.Left, qiymat) // faqat chap tomonni tekshirish kifoya
	}
	return qidir(tugun.Right, qiymat) // faqat o'ng tomonni tekshirish kifoya
}
```

**Nega bu tez.** E'tibor bering: har bir qadamda, agar `qiymat < tugun.Value` bo'lsa, **butun o'ng qism daraxtni tekshirishga hojat yo'q** — BST xossasiga ko'ra, u yerda faqat kattaroq qiymatlar bor, demak qidirilayotgan kichikroq qiymat u yerda bo'la olmaydi. Shu tarzda, har qadamda qidiruv maydoni (taxminan) yarmiga qisqaradi — xuddi "Binary Search"dagidek, `O(log n)` tezlikda (agar daraxt "muvozanatlashgan", ya'ni bir tomonga qattiq "cho'zilib ketmagan" bo'lsa).

**Butun daraxtni oddiy `inorder` bilan qidirish bilan solishtiring.** Agar BST xossasidan foydalanmasdan, "Binary Tree" darsidagi kabi butun daraxtni aylanib chiqib qidirsangiz, bu har doim `O(n)` vaqt oladi — har bir tugunni tekshirish kerak bo'ladi. BST xossasidan foydalanish esa har safar **yarmi** tugunlarni butunlay chetlab o'tish imkonini beradi.

## EXAMPLE

```go
package main

import "fmt"

type Tugun struct {
	Value int
	Left  *Tugun
	Right *Tugun
}

func qidir(tugun *Tugun, qiymat int) bool {
	if tugun == nil {
		return false
	}
	if tugun.Value == qiymat {
		return true
	}
	if qiymat < tugun.Value {
		return qidir(tugun.Left, qiymat)
	}
	return qidir(tugun.Right, qiymat)
}

func main() {
	ildiz := &Tugun{
		Value: 5,
		Left:  &Tugun{Value: 3, Left: &Tugun{Value: 1}},
		Right: &Tugun{Value: 8, Left: &Tugun{Value: 6}, Right: &Tugun{Value: 9}},
	}
	fmt.Println(qidir(ildiz, 6))  // true
	fmt.Println(qidir(ildiz, 7))  // false
}
```

Natija:

```
true
false
```

## TASK

`Tugun` struct'i va `qidir(tugun *Tugun, qiymat int) bool` funksiyasi berilgan (`tugun` — BST xossasiga bo'ysunuvchi daraxt ildizi). Funksiyani **rekursiya** va **BST xossasidan foydalanib** (butun daraxtni emas, faqat kerakli tomonni tekshirib) shunday to'ldiringki, u `qiymat` daraxtda bor-yo'qligini `bool` qilib qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bazaviy holatlar: `if tugun == nil { return false }`, `if tugun.Value == qiymat { return true }`.
2. `if qiymat < tugun.Value { return qidir(tugun.Left, qiymat) }` aks holda `return qidir(tugun.Right, qiymat)`.
