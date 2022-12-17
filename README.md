# Desafio-Go-Bases-I

Oficina de código: Desafio Go

Objetivo:

A seguir, é colocado um desafio integrador que nos permitirá avaliar os tópicos vistos até agora. 

Enunciado:

Uma CIA aérea pequena possui um sistema de reservas de bilhetes para diferentes países. Ele retorna um arquivo com as informações sobre os bilhetes comprados nas últimas 24 horas. A CIA precisa de um programa para extrair informações sobre as vendas do dia e, assim, analisar as tendências de compra. 

O arquivo em questão é do tipo valores separados com vírgula (CSV), onde os campos são formados por: id, nome, email, país de destino, hora do voo e preço.
Boa sorte!

Desafio:

Realizar um programa que sirva como ferramenta para calcular diferentes dados estatísticos. Para isso, você deverá baixar este repositório, que contém um arquivo CSV com dados gerados e o esqueleto do projeto. 


Atenção! Os exemplos a seguir são somente a modo de orientação. O desafio pode ser resolvido de múltiplas formas. 

Requisito 1:

Uma função que calcule a quantidade de pessoas que viajam para um país determinado.

func GetTotalTickets(destination string) (int, error) {}
(exemplo 1)

 Tip: VS Code nos permite pesquisar uma palavra em um arquivo com  Ctrl + F ou ⌘ + F.
 
Requisito 2: 

Uma ou várias funções que calculam quantas pessoas viajam de madrugada (0 → 6), manhã (7 → 12), tarde (13 → 19), e noite (20 → 23).

func GetCountByPeriod(time string) (int, error) {}

(exemplo 2)

 Dica: Em Go, temos o pacote strings para manipular caracteres.
 
Requisito 3: 

Calcular a média de pessoas que viajam para um determinado país. 

func AverageDestination(destination string) (float64, error) {}
(exemplo 3)

 Dica: A média de x é calculada como: x̄ =  xn 
