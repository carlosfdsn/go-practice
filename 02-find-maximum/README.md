# 02 — Find Maximum

Exercício em Go focado na iteração de slices, comparação de valores e análise de complexidade.

## Objetivo

Implementar uma função que receba uma `slice` de inteiros e retorne o maior elemento presente nela.

A solução deve percorrer a `slice` apenas uma vez

## Exemplo

```text
Entrada:
[4, 12, 7, 25, 9]

Saída:
25
```

## Restrições

* Não utilizar `sort`.
* Não criar uma nova `slice`.
* Percorrer a `slice` apenas uma vez.
* Utilizar apenas `O(1)` de espaço adicional.

## Complexidade

| Recurso | Complexidade |
| ------- | ------------ |
| Tempo   | `O(n)`       |
| Espaço  | `O(1)`       |

## Assinatura da função

```go
func findMaximum(nums []int) int
```

## Desafio

Implemente a solução sem consultar uma implementação pronta.
