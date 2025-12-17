package godist

import (
	"fmt"
	"math"
	"testing"
	"github.com/iannsp/godist"
)

/*
Validação de Teste A/B e Lançamento de Feature (Rollout)

Alice, uma Product Manager utilizando estatistica para provar retorno sobre o investimento 
e minimizar riscos.

O Cenário: Otimização do Funil de Assinatura

Contexto:
Alice é PM de um aplicativo de Streaming de Vídeo (estilo Netflix/Globoplay).
A meta do trimestre é aumentar a conversão da página de assinatura (o momento em 
que o usuário gratuito vira pagante).

* Versão Atual (Controle - A): Taxa de conversão média histórica de 5% (p=0.05p=0.05).

* Nova Versão (Variante - B): Alice e o time de Design criaram uma nova página 
simplificada, com menos campos para preencher.

Ela roda um experimento (Teste A/B) mostrando a nova página para 1.000 usuários.
Resultado do Teste: 60 usuários assinaram (Conversão de 6%).

O Dilema:
A conversão subiu de 5% para 6%. Isso parece ótimo (20% de melhoria!).
Mas Alice se pergunta: "Isso foi sorte? Foi apenas um dia atípico? Ou a nova 
página é realmente melhor?"

Se ela lançar para milhões de usuários e for apenas sorte, a empresa pode perder 
dinheiro a longo prazo.
*/

/*
Qual a probabilidade da versão Antiga gerar 60 ou mais vendas por pura sorte?
*/

var taxaConversaoBase = 0.05
var taxaConversaoObservada = 0.06
var numeroUsuarioTeste = 1000
var conversao60 = 60
var nivelSignificancia = 0.05

/*
    Este teste visa validar se a diferença na conversão é realmente causada pelo
    novo modelo B, ou seja eliminando a possibilidade do modelo A obter esse tipo
    de resultado.
    O resultado é inconclusivo pq para descartar a hipotese desse resultado ter
    sido gerado pelo modelo A a probabilidade deveria ser insignificante (< nivelSignificancia).
*/
func TestQualAProbabilidadeVersaoAntigaGerar60ouMaisVendasHipoteseNula(t *testing.T){
    modeloA, _ := godist.NewBinomial(numeroUsuarioTeste, taxaConversaoBase)

    // O ajuste ocorre pq CDF calcula de 0 a k(k ou menos é a resposta que ele traz)
    // precisamos ajustar para que 60 fique fora do CDF, então:
    // 1 = CDF(59) + P(X >= 60) => P(x >= 60)= 1- CDF(59)
    ajuste := -1
    p60, _ := modeloA.CDF(conversao60 + ajuste) // p(x <= 59)
    probVal := 1- p60 // P(x >= 60)
    // 1 em cada 20 (Ronald Fisher) padrão de desvio considerado significativo.
    fmt.Printf("p60 = %f => %.4f", p60, 1- p60)

    // Se a chance de o modelo A, por sorte, converter >=60 for menor que 
    // nivelSignificancia então a diferença é causada pelo Modelo B.
    // se for maior então o resultado é não é conclusivo já que a "grande" change
    // do próprio modelo A atingir esse resultado +60 conversões por sorte.

    if probVal < nivelSignificancia{
        t.Errorf("Erro de Calculo. Esperado Aprox. %.3f, recebido %.3f", 0.08 , probVal)
    }

    if probVal < nivelSignificancia {
        fmt.Printf("[Aprovado] Não foi sorte, novo modelo é melhor.")
    }else {
        fmt.Printf("[Inconclusivo] %.2f Existe chance de ser somente ruido e o novo modelo não ser melhor.\n Necessário aumentar o tamanho da amostra rodando teste mais dias.\n", probVal * 100)
    }
}
/*
    Este teste valida o resultado obtido com o binomial utilizando "distancia" estatistica.
    Quantos desvios Padrão a nova versão se afastou da velha? (zScore) 
    o zCritico de 1.645 é o nivelSignificancia (5%). Se zCritico > nivelSignificancia ou
    menor que nivelSignificancia * -1 indica significancia estatistica.

    90% => Z = 1.28
    95% => Z = 1.645
    99% => Z = 2.33
*/
func TestComparaCurvaModelosAB( t *testing.T){
    amostra := float64(numeroUsuarioTeste)

    erroPadrao := math.Sqrt( taxaConversaoBase * (1- taxaConversaoBase) / float64(amostra) )
    curvaNormalA := godist.NewNormal( taxaConversaoBase, erroPadrao)

    zScore := curvaNormalA.SingleDataZ(taxaConversaoObservada)

    probVal := 1 - curvaNormalA.CDF(taxaConversaoObservada)
    zCritico := 1.645

    if zScore > zCritico {
        t.Errorf("Erro de Calculo. Esperado Aprox. %.3f, recebido %.3f", 1.450, zScore)
    }
    if zScore > zCritico {
        fmt.Printf("[Aprovado] Não foi sorte, novo modelo é melhor.")
    }else {
        fmt.Printf("[Inconclusivo] %.2f%% Existe chance de ser somente ruido e o novo modelo não ser melhor.\n Necessário aumentar o tamanho da amostra rodando teste mais dias.\n", probVal * 100)
    }
}
