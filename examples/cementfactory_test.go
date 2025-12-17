package godist

import (
	"fmt"
	"github.com/iannsp/godist"
	_ "log/slog"
	"math"
	"testing"
)

/*
Engenharia de Produção e Controle de Qualidade.

Vamos imaginar que você trabalha como Gerente de Qualidade em uma fábrica de cimento. O problema envolve garantir que os sacos vendidos tenham o peso prometido, mas analisando o risco em lotes de envio.

A fábrica vende sacos de cimento rotulados como 50 kg.
Devido à vibração e imprecisão das máquinas de envase, não é possível colocar exatamente 50.000g em cada saco. O peso oscila naturalmente.

Para evitar processos legais por vender menos do que o anunciado, a fábrica regula a máquina para colocar um pouco a mais, com uma média de 50,5 kg

A fábrica envia o cimento em paletes com 20 sacos.
O cliente (uma grande construtora) tem uma regra de recebimento: Eles pesam todos os 20 sacos. Se houver 2 ou mais sacos abaixo do peso, eles rejeitam o palete inteiro e devolvem a carga.

Questões:
1- Qual é a probabilidade de um único saco sair da máquina com peso menor que o permitido (menos de 50 kg)?

2- Qual é a probabilidade de o palete ser devolvido?

*/

// DADOS do Problema
var mediaFabrica = 50.5
var desvioPadraoMaquina = 0.3
var pesoMinimoAceitavel = 50.0
var probabilidadePesoIdeal = 0.0478
var numeroSacosPorPalete = 20
var expectedProababilidadeZeroDefeito = 0.375
var expectedProababilidadeUmDefeito = 0.375
var expectedProbabilidadeZeroDefeito = 0.375460
var expectedProbabilidadeUmDefeito = 0.376958
var expectedProbabilidadeRejeicaoPalete = 24.758186

/*
Queremos saber a probabilidade de um saco pesar **menos que 50.0 kg**, dado que a máquina produz com média de **50.5 kg** e desvio padrão de **0.3 kg**.
*/
func TestProbabilidadeUnicoSacoTerPesoMenorquePermitido(t *testing.T) {
	probabilidadeEsperada := probabilidadePesoIdeal
	distribuicaoCimento := godist.NewNormal(mediaFabrica, desvioPadraoMaquina)

	fmt.Printf("Média da Máquina: %.2f kg\n", mediaFabrica)
	fmt.Printf("Desvio Padrão: %.2f kg\n", desvioPadraoMaquina)
	fmt.Printf("Limite Mínimo: %.2f kg\n", pesoMinimoAceitavel)
	fmt.Println("------------------------------------------------")

	// Utilizando Less
	probabilidadeDefeito := distribuicaoCimento.Less(pesoMinimoAceitavel)
	fmt.Printf("Existe uma probabildiade de  %.2f%% de um saco defeituoso ser produzido(que pese menos de 50kg)\n",
		probabilidadeDefeito*100)

	if math.Abs(probabilidadeDefeito-probabilidadeEsperada) > 0.0001 {
		t.Errorf("Calcula Falhou. Esperado %.4f, recebido %.4f", probabilidadeEsperada, probabilidadeDefeito)
	}
}

/*
Queremos saber qual é a probabilidade de o palete ser devolvido? Sabendo que são pesados todos os **20 sacos** de cimento no palete, que a **probabilidade de sucesso/falha individual é = 0.0478** (calculada em TestProbabilidadeUnicoSacoTerPesoMenorquePermitido) e que o **critério de rejeição é 2** ou mais defeitos (peso menor pesoMinimoAceitavel) por palete.
*/
func TestProbabilidadePaleteSerDevolvido(t *testing.T) {

	// A abordagem escolhida é a da [Probabilidade Complementar](https://pt.wikipedia.org/wiki/Evento_complementar) para o problema é utilizar a chance de 0 defeitos(todos os 20 sacos dentro do peso) e de 1 defeito (1 saco abaixo do limite de peso) e extrair a chance de 2 ou mais defeitos.
    binomial, _ := godist.NewBinomial( numeroSacosPorPalete, probabilidadePesoIdeal)
	probabilidadeZeroDefeito, err := binomial.PMF(0)

	if err != nil {
		t.Errorf("Falha ao Calcular PMF. %s", err.Error())
	}


    binomial, _ = godist.NewBinomial(numeroSacosPorPalete, probabilidadePesoIdeal)
	probabilidadeUmDefeito, err := binomial.PMF(1)

	if err != nil {
		t.Errorf("Falha ao Calcular PMF. %s", err.Error())
	}

	if math.Abs(probabilidadeZeroDefeito-expectedProbabilidadeZeroDefeito) > 0.001 {
		t.Errorf("Falha Calculo P0def. Expected %f, got %f", expectedProababilidadeZeroDefeito, probabilidadeZeroDefeito)
	}

	if math.Abs(probabilidadeUmDefeito-expectedProbabilidadeUmDefeito) > 0.001 {
		t.Errorf("Falha Calculo P1def. Expected %f, got %f", expectedProababilidadeUmDefeito, probabilidadeUmDefeito)
	}

	probabilidadeAceitacaoPalete := (probabilidadeZeroDefeito + probabilidadeUmDefeito) * 100
	probabilidadeRejeicaoPalete := 100 - probabilidadeAceitacaoPalete

	if math.Abs(probabilidadeRejeicaoPalete-expectedProbabilidadeRejeicaoPalete) > 0.001 {
		t.Errorf("Falha Calculo PRej. Expected %f, got %f", expectedProbabilidadeRejeicaoPalete, probabilidadeRejeicaoPalete)
	}
}
