package main

import (
	"fmt"
	"math"
)

type JurosCompostosMontante struct {
	capitalInicial float64
	taxaDeJuros    float64
	periodo        int
}

func (j *JurosCompostosMontante) calcularMontante() float64 {
	montante := j.capitalInicial * math.Pow((1+j.taxaDeJuros), float64(j.periodo))
	return montante
}

type JurosCompostosCapital struct {
	montante    float64
	taxaDeJuros float64
	periodo     int
}

func (c *JurosCompostosCapital) calcularCapital() float64 {
	capital := c.montante / math.Pow((1+c.taxaDeJuros), float64(c.periodo))
	return capital
}

type JurosCompostosTaxa struct {
	montante       float64
	capitalInicial float64
	periodo        int
}

func (t *JurosCompostosTaxa) calcularTaxa() float64 {
	taxa := math.Pow((t.montante/t.capitalInicial), (1/float64(t.periodo))) - 1
	return taxa
}

type JurosCompostosPeriodo struct {
	montante       float64
	capitalInicial float64
	taxaDeJuros    float64
}

func (p *JurosCompostosPeriodo) calcularPeriodo() float64 {
	periodo := math.Log(p.montante/p.capitalInicial) / math.Log(1+p.taxaDeJuros)
	return periodo
}

type JurosCompostosTaxaEfetiva struct {
	taxaNominal   float64
	numeroPeriodo int
}

func (te *JurosCompostosTaxaEfetiva) calcularTaxaEfetiva() float64 {
	taxaEfetiva := math.Pow(1+te.taxaNominal, 1/float64(te.numeroPeriodo)) - 1
	return taxaEfetiva
}

func (tnc *JurosCompostosTaxaEfetiva) calcularTaxaNominalConvertida() float64 {
	resultadoTaxaEfetiva := tnc.calcularTaxaEfetiva()
	taxaNominalConvertida := resultadoTaxaEfetiva * float64(tnc.numeroPeriodo)
	return taxaNominalConvertida
}

func converterTaxa(taxa float64, dePeriodo string, paraPeriodo string) float64 {
	if dePeriodo == paraPeriodo {
		return taxa
	}
	if dePeriodo == "anual" && paraPeriodo == "mensal" {
		return math.Pow(1+taxa, 1/12.0) - 1
	} else if dePeriodo == "mensal" && paraPeriodo == "anual" {
		return math.Pow(1+taxa, 12) - 1
	}
	return taxa
}

func main() {
	loop := true
	var opcao int

	for loop {
		fmt.Println("Bem-Vindo a Calculadora Financeira")
		fmt.Println("Digite 1 para entrar no menu Montante!")
		fmt.Println("Digite 2 para entrar no menu Capital!")
		fmt.Println("Digite 3 para entrar no menu Taxa!")
		fmt.Println("Digite 4 para entrar no menu Periodo!")
		fmt.Println("Digite 5 para entrar no menu Taxa Nominal e Efetiva!")
		fmt.Println("Digite 6 para encerrar!")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var valorCapitalInicial, valorTaxaDeJuros float64
			var valorPeriodo float64
			var periodoTaxa string

			fmt.Println("Você entrou no menu Montante!")
			fmt.Println("Digite o valor do Capital Inicial:")
			fmt.Scanln(&valorCapitalInicial)
			fmt.Println("Digite o valor da Taxa de Juros (decimal):")
			fmt.Scanln(&valorTaxaDeJuros)
			fmt.Println("Digite o período da Taxa de Juros (mensal/anual):")
			fmt.Scanln(&periodoTaxa)
			fmt.Println("Digite o valor do Periodo (meses):")
			fmt.Scanln(&valorPeriodo)

			valorTaxaDeJuros = converterTaxa(valorTaxaDeJuros, periodoTaxa, "mensal")

			jurosCompostosMontante := JurosCompostosMontante{
				capitalInicial: valorCapitalInicial,
				taxaDeJuros:    valorTaxaDeJuros,
				periodo:        int(valorPeriodo),
			}

			fmt.Printf("O valor do Montante é: R$ %.2f\n", jurosCompostosMontante.calcularMontante())
		case 2:
			var valorMontante, valorTaxaDeJuros float64
			var valorPeriodo float64
			var periodoTaxa string

			fmt.Println("Você entrou no menu Capital!")
			fmt.Println("Digite o valor do Montante:")
			fmt.Scanln(&valorMontante)
			fmt.Println("Digite o valor de Taxa De Juros (decimal):")
			fmt.Scanln(&valorTaxaDeJuros)
			fmt.Println("Digite o período da Taxa de Juros (mensal/anual):")
			fmt.Scanln(&periodoTaxa)
			fmt.Println("Digite o valor do Periodo (meses):")
			fmt.Scanln(&valorPeriodo)

			valorTaxaDeJuros = converterTaxa(valorTaxaDeJuros, periodoTaxa, "mensal")

			jurosCompostosCapital := JurosCompostosCapital{
				montante:    valorMontante,
				taxaDeJuros: valorTaxaDeJuros,
				periodo:     int(valorPeriodo),
			}

			fmt.Printf("O valor do Capital é: R$ %.2f\n", jurosCompostosCapital.calcularCapital())
		case 3:
			var valorMontante, valorCapital float64
			var valorPeriodo float64

			fmt.Println("Você entrou no menu Taxa!")
			fmt.Println("Digite o valor do Montante:")
			fmt.Scanln(&valorMontante)
			fmt.Println("Digite o valor do Capital:")
			fmt.Scanln(&valorCapital)
			fmt.Println("Digite o valor do Periodo (meses):")
			fmt.Scanln(&valorPeriodo)

			jurosCompostosTaxa := JurosCompostosTaxa{
				montante:       valorMontante,
				capitalInicial: valorCapital,
				periodo:        int(valorPeriodo),
			}

			fmt.Printf("O valor da Taxa em decimal é: %.2f\n", jurosCompostosTaxa.calcularTaxa())
		case 4:
			var valorMontante, valorCapitalInicial, valorTaxaDeJuros float64

			fmt.Println("Você entrou no menu Periodo!")
			fmt.Println("Digite o valor do Montante:")
			fmt.Scanln(&valorMontante)
			fmt.Println("Digite o valor do Capital Inicial:")
			fmt.Scanln(&valorCapitalInicial)
			fmt.Println("Digite o valor da Taxa de Juros (decimal):")
			fmt.Scanln(&valorTaxaDeJuros)

			jurosCompostosPeriodo := JurosCompostosPeriodo{
				montante:       valorMontante,
				capitalInicial: valorCapitalInicial,
				taxaDeJuros:    valorTaxaDeJuros,
			}

			fmt.Println("O valor do Periodo em meses é: ", int(jurosCompostosPeriodo.calcularPeriodo()))
		case 5:
			var valorTaxaNominal float64
			var valorNumeroPeriodo int

			fmt.Println("Você entrou no menu Taxa Nominal e Efetiva!")
			fmt.Println("Digite o valor da Taxa Nominal (decimal):")
			fmt.Scanln(&valorTaxaNominal)
			fmt.Println("Digite o valor do Numero de Periodo:")
			fmt.Scanln(&valorNumeroPeriodo)

			jurosCompostosTaxaEfetiva := JurosCompostosTaxaEfetiva{
				taxaNominal:   valorTaxaNominal,
				numeroPeriodo: valorNumeroPeriodo,
			}

			fmt.Printf("O valor da Taxa Efetiva é: %.3f\n", jurosCompostosTaxaEfetiva.calcularTaxaEfetiva())
			fmt.Printf("O valor da Taxa Nominal Convertida é: %.3f\n", jurosCompostosTaxaEfetiva.calcularTaxaNominalConvertida())
		case 6:
			fmt.Println("Você encerrou o programa!")
			loop = false
		default:
			fmt.Println("Valor Inválido!")
		}
	}
}