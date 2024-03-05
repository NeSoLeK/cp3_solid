package solid

import "fmt"

// Интерфейс для устройства
type Device interface {
	Print()
	Scan()
}

// Структура для принтера
type Printer struct{}

func (p *Printer) Print() {
	fmt.Println("Printing document")
}

func (p *Printer) Scan() {
	panic("Printer cannot scan")
}

// Структура для сканера
type Scanner struct{}

func (s *Scanner) Print() {
	panic("Scanner cannot print")
}

func (s *Scanner) Scan() {
	fmt.Println("Scanning document")
}

func PerformPrintJob(d Device) {
	d.Print()
}

func PerformScanJob(d Device) {
	d.Scan()
}

func I() {
	fmt.Print("[Interface Segregation Principle]\n")

	printer := &Printer{}
	scanner := &Scanner{}

	PerformPrintJob(printer)
	PerformScanJob(scanner)
}
