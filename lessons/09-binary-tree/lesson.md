# 09 — Binary Tree

## THEORY

Oilaviy shajarani (daraxtini) tasavvur qiling — bir ota-onadan bir nechta farzand, ulardan yana farzandlar, va hokazo, shoxlanib boradi. **Daraxt (tree)** — ma'lumotlar tuzilmasi sifatida ham aynan shunday ishlaydi: har bir **tugun (node)** o'z "farzandlariga" (bolalariga) ishora qiladi. **Binary tree (ikkilik daraxt)** — bunda har bir tugun ko'pi bilan **ikkita** bolaga ega bo'ladigan maxsus holat: **chap (Left)** va **o'ng (Right)**.

**Tugun qanday e'lon qilinadi.** "Linked List" darsida ko'rgan `Node`ga o'xshaydi, lekin endi **ikkita** pointer bor:

```go
type Tugun struct {
	Value int
	Left  *Tugun
	Right *Tugun
}
```

**Qo'lda daraxt yasash:**

```go
ildiz := &Tugun{
	Value: 5,
	Left:  &Tugun{Value: 3},
	Right: &Tugun{Value: 8},
}
```

Bu daraxt shunday ko'rinadi:

```
      5
     / \
    3   8
```

**Aylanib chiqish (traversal) — "inorder".** Daraxtni "to'g'ri" tartibda aylanib chiqishning bir necha yo'li bor. Eng ko'p ishlatiladigani — **inorder** (chap → o'zi → o'ng): avval butun **chap** qism daraxtni (rekursiv!) aylaning, keyin joriy tugunning qiymatini oling, so'ng butun **o'ng** qism daraxtni aylaning:

```go
func inorderAylanish(tugun *Tugun) []int {
	if tugun == nil {
		return []int{} // bazaviy holat — bo'sh joy, qo'shadigan narsa yo'q
	}
	natija := inorderAylanish(tugun.Left)
	natija = append(natija, tugun.Value)
	natija = append(natija, inorderAylanish(tugun.Right)...)
	return natija
}
```

Bu — "Recursion" darsida ko'rgan naqshning tabiiy davomi: daraxtlar **o'zi rekursiv tuzilmalar** (har bir tugunning bolasi ham — o'zi bir kichikroq daraxt), shuning uchun ular ustida ishlaydigan funksiyalar deyarli har doim rekursiv yoziladi — bu holatda "for" tsikli bilan yozish ancha murakkabroq bo'lar edi.

**Nega aynan "inorder" deb ataladi va nega foydali.** Agar daraxt maxsus tartibda qurilgan bo'lsa (keyingi darsda ko'radigan **Binary Search Tree**), inorder aylanish natijasi avtomatik ravishda **saralangan** ro'yxat beradi — bu daraxtlarning eng foydali xususiyatlaridan biri. Boshqa aylanish turlari ham bor — **preorder** (o'zi → chap → o'ng) va **postorder** (chap → o'ng → o'zi) — ular daraxtni nusxalash yoki o'chirish kabi boshqa vazifalar uchun qulayroq.

## EXAMPLE

```go
package main

import "fmt"

type Tugun struct {
	Value int
	Left  *Tugun
	Right *Tugun
}

func inorderAylanish(tugun *Tugun) []int {
	if tugun == nil {
		return []int{}
	}
	natija := inorderAylanish(tugun.Left)
	natija = append(natija, tugun.Value)
	natija = append(natija, inorderAylanish(tugun.Right)...)
	return natija
}

func main() {
	ildiz := &Tugun{
		Value: 5,
		Left:  &Tugun{Value: 3},
		Right: &Tugun{Value: 8},
	}
	fmt.Println(inorderAylanish(ildiz))
}
```

Natija:

```
[3 5 8]
```

## TASK

`Tugun` struct'i (`Value int`, `Left *Tugun`, `Right *Tugun`) va `inorderAylanish(tugun *Tugun) []int` funksiyasi berilgan. Funksiyani **rekursiya** yordamida shunday to'ldiringki, u daraxtni **inorder** tartibda (chap → o'zi → o'ng) aylanib, barcha qiymatlarni `[]int` slice qilib qaytarsin. Bo'sh daraxt (`nil`) uchun bo'sh slice qaytarilishi kerak.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bazaviy holat: `if tugun == nil { return []int{} }`.
2. `natija := inorderAylanish(tugun.Left)`, `natija = append(natija, tugun.Value)`, `natija = append(natija, inorderAylanish(tugun.Right)...)`, so'ng `return natija`.
