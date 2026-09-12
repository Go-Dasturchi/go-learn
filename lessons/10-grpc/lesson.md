# 10 — gRPC

## THEORY

"HTTP & Backend" bo'limida, REST API — JSON orqali, matn ko'rinishida ma'lumot almashishni ko'rdik. Bu — inson o'qiy oladigan, moslashuvchan format, lekin **katta hajmda** (masalan, mikroservislar orasida, soniyasiga minglab so'rov) — JSON'ni har safar matn sifatida yozish/o'qish (serialize/deserialize qilish) — nisbatan **sekin va og'ir**.

**gRPC** — Google tomonidan yaratilgan, xizmatlar orasida **tezkor** aloqa uchun mo'ljallangan protokol. U ikkita narsaga tayanadi:

1. **Protocol Buffers (protobuf)** — JSON o'rniga, ma'lumotni **ikkilik (binary)** formatda, juda ixcham qilib kodlaydi.
2. **HTTP/2** — bir nechta so'rov/javobni, bitta ulanish orqali, **parallel** yuborish imkonini beradi.

### 1-qadam: kerakli vositalarni o'rnatish

**macOS:**

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

**Ubuntu/Debian:**

```bash
sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Tekshirish:

```bash
protoc --version
```

### 2-qadam: `.proto` fayl yozish

gRPC'da, avval xizmat va xabarlar **`.proto`** faylida tasvirlanadi — bu, tilga bog'liq bo'lmagan "shartnoma":

```protobuf
syntax = "proto3";

package salom;

service SalomlashXizmati {
  rpc Salomlash (SalomSorovi) returns (SalomJavobi);
}

message SalomSorovi {
  string ism = 1;
}

message SalomJavobi {
  string xabar = 1;
}
```

### 3-qadam: Go kodini generatsiya qilish

```bash
protoc --go_out=. --go-grpc_out=. salom.proto
```

Bu buyruq, `.proto` faylidan, Go struct'lari va gRPC server/klient **interfeyslarini avtomatik yaratadi** — qo'lda yozish shart emas.

### 4-qadam: Server va klient yozish

```go
type server struct {
	pb.UnimplementedSalomlashXizmatiServer
}

func (s *server) Salomlash(ctx context.Context, req *pb.SalomSorovi) (*pb.SalomJavobi, error) {
	return &pb.SalomJavobi{Xabar: "Salom, " + req.Ism}, nil
}

// Klient tomonida:
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
client := pb.NewSalomlashXizmatiClient(conn)
javob, _ := client.Salomlash(context.Background(), &pb.SalomSorovi{Ism: "Ali"})
```

**Nega bu — oddiy funksiya chaqirishga o'xshaydi.** gRPC'ning eng katta afzalligi — u, **tarmoq orqali** so'rov yuborishni, xuddi **oddiy Go funksiyasini chaqirgandek** ko'rsatadi (`client.Salomlash(...)`), garchi orqa fonda HTTP/2 va protobuf ishlatilsa ham. Bu — **RPC (Remote Procedure Call — masofaviy protsedura chaqiruvi)** g'oyasining o'zi.

**Nega bu darsda haqiqiy `protoc`/gRPC ishlatilmaydi.** Bu, tashqi vositalar (protoc kompilyatori) va tarmoq ulanishini talab qiladi, checker esa har bir mashqni **tarmoqsiz, izolyatsiyalangan** muhitda tekshiradi. Shu sababli, RPC'ning **markaziy g'oyasini** — "masofaviy xizmatni, mahalliy interfeys kabi chaqirish" — oddiy Go interfeysi orqali mashq qilamiz:

```go
type SalomlashXizmati interface {
	Salomlash(ism string) string
}

type MahalliyXizmat struct{}

func (m *MahalliyXizmat) Salomlash(ism string) string {
	return "Salom, " + ism
}

func Sorov(xizmat SalomlashXizmati, ism string) string {
	return xizmat.Salomlash(ism) // xuddi gRPC klienti kabi — "qayerda" ishlashi muhim emas
}
```

`Sorov` funksiyasi, `xizmat` haqiqatan **shu jarayonda** ishlaydimi, yoki **tarmoq orqali, boshqa serverda** ishlaydimi — bilishga **hojati yo'q**, chunki u faqat interfeys bilan gaplashadi. Bu, aynan "Repository Pattern" va "Clean Architecture" darslarida ko'rgan g'oyaning, tarmoq chegarasiga qo'llanilishi.

## EXAMPLE

```go
package main

import "fmt"

type SalomlashXizmati interface {
	Salomlash(ism string) string
}

type MahalliyXizmat struct{}

func (m *MahalliyXizmat) Salomlash(ism string) string {
	return "Salom, " + ism
}

func Sorov(xizmat SalomlashXizmati, ism string) string {
	return xizmat.Salomlash(ism)
}

func main() {
	xizmat := &MahalliyXizmat{}
	fmt.Println(Sorov(xizmat, "Ali"))
	fmt.Println(Sorov(xizmat, "Vali"))
}
```

Natija:

```
Salom, Ali
Salom, Vali
```

## TASK

`Sorov(xizmat SalomlashXizmati, ism string) string` funksiyasini to'ldiring — u, `xizmat.Salomlash(ism)` ni chaqirib, natijasini qaytarsin (xuddi gRPC klienti, masofaviy xizmatni chaqirgandek, lekin interfeys orqali).

`main()` funksiyasini o'zgartirish shart emas.

**Terminalda sinash uchun** (checker tekshirmaydi): `protoc`ni o'rnatib, yuqoridagi `.proto` faylini yozib, Go kodini generatsiya qilib ko'ring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `return xizmat.Salomlash(ism)`.
2. `Sorov` funksiyasi, `xizmat`ning aniq turini bilmaydi — faqat interfeysni biladi.
