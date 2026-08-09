# 01 — Frequency Counter

Exercício em Go focado na utilização de `map` para contar a frequência de elementos em uma `slice`.

## Objetivo

Implementar uma função que receba uma `slice` de inteiros e retorne um `map` contendo a quantidade de vezes que cada número aparece.

## Exemplo

```text
Entrada:
[2, 2, 3, 3, 5, 1, 4]

Saída:
2 → 2
3 → 2
5 → 1
1 → 1
4 → 1
```

## Restrições

* Utilizar um `map` para armazenar as frequências.
* Percorrer a `slice` apenas uma vez.
* Não ordenar os elementos.
* Não criar uma nova `slice`.

## Complexidade

| Recurso | Complexidade |
| ------- | ------------ |
| Tempo   | `O(n)`       |
| Espaço  | `O(n)`       |

`n` representa a quantidade de elementos presentes na `slice`.

## Assinatura da função

```go
func frequency(frequency_list []int) map[int]int
```

## Desafio

Implemente a solução sem consultar uma implementação pronta.
