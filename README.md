# Lista 3: Manipulação de erros em Go

## Este Trabalho foi realizado em dupla

### Juntamente de [Pablo Santos](https://github.com/BiscuI)

# Tarefa:

<sub>Última atualização: 26/07/2023</sub>

## Sumário

- [Visão geral e objetivos](#visão-geral-e-objetivos)
- [Tarefas](#tarefas)
- [Conhecimentos necessários](#conhecimentos-necessários)
- [Requisitos](#requisitos)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)
- [Dúvidas e informações](#dúvidas-e-informações)

## Visão geral e objetivos

O objetivo deste exercício de programação é implementar, na linguagem de programação Go, um programa chamado `sudoku-checker` que **verifica** a satisfação das regras do jogo _Sudoku_. O _Sudoku_ consiste no posicionamento lógico de números entre 1 e 9 em cada uma das células vazias em uma grade de tamanho 9 x 9 constituída de nove subgrades (regiões) de tamanho 3 x 3. O quebra-cabeça contém algumas pistas iniciais, que são números inseridos em algumas células, de maneira a permitir uma indução ou dedução dos números em células que estejam vazias.

<p align="center">
   <img width="25%" src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg/1200px-Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg.png" alt="A typical Sudoku puzzle, with nine rows and nine columns that intersect at square spaces. Some of the cells are filled with a number; others are blank cells to be solved." />
   <img width="25%" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/Sudoku_Puzzle_by_L2G-20050714_solution_standardized_layout.svg/1200px-Sudoku_Puzzle_by_L2G-20050714_solution_standardized_layout.svg.png" alt="Sudoku Puzzle by L2G-20050714 solution standardized layout.svg"/>
   <br/>
 <span><em>Exemplo de grade do Sudoku (esquerda) e respectiva solução (direita).<br/>Fonte: Wikipedia.</em></span>
</p>

O preenchimento das células no _Sudoku_ deve obedecer a algumas regras. Em essência, o número a ser inserido não deve ter sido previamente colocado na linha, coluna ou região desejada. Ou seja, cada coluna, linha e região só pode ter um número de 1 a 9. Além disso, os números com os quais a grade do _Sudoku_ é inicialmente preenchida não são passíveis de modificação, isto é, eles são fixos em suas posições e não podem ser sobrescritos. Com isso, a inserção de números nas células deve ocorrer unicamente sobre células vazias.

**Observação:** Este exercício de programação possui foco unicamente na verificação das restrições do _Sudoku_. **Não é necessário implementar um solucionador (_solver_) automático para o jogo.**

## Tarefas

As seguintes tarefas deverão ser realizadas na implementação do programa `sudoku-checker`:

1. Implementar uma função que prepara uma grade do _Sudoku_ com números previamente preenchidos a partir de um arquivo de texto fornecido como entrada ao programa como argumento de linha de comando. Nesta implementação, as células vazias são representadas pelo número 0. Um exemplo de conteúdo desse arquivo de entrada seria:

```bash
5 3 0 0 7 0 0 0 0
6 0 0 1 9 5 0 0 0
0 9 8 0 0 0 0 6 0
8 0 0 0 6 0 0 0 3
4 0 0 8 0 3 0 0 1
7 0 0 0 2 0 0 0 6
0 6 0 0 0 0 2 8 0
0 0 0 4 1 9 0 0 5
0 0 0 0 8 0 0 7 9
```

Considerando que o _Sudoku_ possui um tamanho predefinido, sugere-se utilizar constantes e um _alias_ de tipo para o _array_ bidimensional que representa a grade do jogo:

```go
// Dimensões da grade do Sudoku
const rows, columns = 9, 9

// Grade do Sudoku
type Grid [rows][columns]int8
```

No momento da leitura do arquivo, deve-se fazer uma verificação adicional, por meio de outra função, se o conteúdo constitui uma grade válida para o _Sudoku_, isto é, (i) tem-se um conjunto de números dispostos na forma de uma matriz de tamanho 9 x 9 e (ii) se todos os dígitos estão entre 0 e 9. O arquivo representando a grade inicial do _Sudoku_ pode ser elaborado a partir de exemplos disponíveis na Internet, como no site [Sudoku.com](https://sudoku.com/). Uma vez carregada a grade inicial, esta deverá ser exibida na saída padrão.

2. Implementar uma função que realiza a tentativa de inserção de um dígito na grade do _Sudoku_. O usuário deve fornecer, via entrada padrão, as coordenadas da célula onde se deseja realizar a inserção e o dígito a ser inserido. Neste procedimento, deverão ser realizadas as seguintes verificações:

- Se as coordenadas da célula informadas pelo usuário são válidas
- Se a célula cujas coordenadas foram informadas pelo usuário já está preenchida (apenas células que contenham 0 podem ter novos dígitos inseridos)
- Se o dígito a ser inserido é válido

Caso a inserção seja feita de forma bem sucedida, deve-se apresentar na saída padrão a grade do _Sudoku_ com o dígito inserido na célula indicada pelo usuário.

Para cada tipo de violação de restrição do _Sudoku_, devem ser criados erros distintos, cada um deles com tratamento adequado. Além disso, por questões de modularidade, as verificações deverão ser feitas por funções individuais. Em caso de violação de alguma restrição, o programa deverá exibir uma mensagem de erro correspondente à violação na saída e solicitar ao usuário novas entradas.

Caso o conteúdo representando a grade inicial do _Sudoku_ seja inválido, o programa deve exibir uma mensagem de erro ao usuário e ser encerrado, evidentemente garantindo-se a liberação dos recursos associados ao arquivo de entrada utilizado. Deve-se ainda realizar verificações adicionais quanto à ocorrência de eventuais erros durante a abertura e leitura do arquivo de entrada.

## Conhecimentos necessários

Este exercício explora os seguintes elementos da programação em Go, cujos conhecimentos são, portanto, ora necessários:

- Entrada e saída formatada via _console_
- Entrada via arquivos
- Estruturas de controle de fluxo
- Funções
- _Arrays_
- Criação e tratamento de erros customizados
- Instrução `defer`
- Função `panic`

## Requisitos

A implementação deste trabalho requer os seguintes elementos instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

## Autoria e política de colaboração

Este trabalho deverá necessariamente ser realizado em equipe composta de **até dois estudantes**, sendo importante, dentro do possível, dividir as tarefas igualmente entre os integrantes da equipe. Após a implementação das soluções para os problemas propostos, o arquivo [`author.md`](https://github.com/ufrn-golang/lista3/tree/master/author.md) presente no repositório deverá ser editado preenchendo as informações de identificação dos integrantes da equipe, na seção [Informações de Autoria](https://github.com/ufrn-golang/lista3/tree/master/author.md#identificação-de-autoria).

O trabalho em cooperação entre estudantes da turma é estimulado, sendo admissível a discussão de ideias e estratégias. Contudo, tal interação não deve ser entendida como permissão para utilização de (parte de) código fonte de colegas, o que pode caracterizar situação de plágio. Trabalhos copiados no todo ou em parte de outros colegas ou da Internet serão anulados e receberão nota zero.

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://git-scm.com) serão utilizados para possibilitar a entrega da implementação realizadas. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (_assignment_) no GitHub Classroom. Cada integrante de equipe deverá acessar este [_link_], aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelos integrantes da equipe e pelo docente, sob a organização [`ufrn-golang`](https://github.com/ufrn-golang). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo do arquivo executável gerado a partir da compilação do código fonte.

A implementação do programa objeto deste trabalho deverá ser realizada **até as 23:59 do dia 21 de maio de 2023** no respectivo repositório Git da equipe. Para fins de registro, o endereço do repositório também deverá ser **obrigatoriamente** enviado através da opção _Tarefas_ na Turma Virtual do SIGAA, devendo **um único membro da equipe** realizar esse envio. Além disso, **não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.**

## Avaliação

A avaliação do programa implementado contabilizará nota de até 10,0 pontos. O programa implementado será avaliado de acordo com os seguintes critérios:

- utilização correta dos recursos providos pela linguagem de programação Go;
- corretude da execução do programa implementado, que deve apresentar saída em conformidade com a especificação;
- aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte, e;
- correta utilização do repositório Git, no qual deverá ser registrado todo o histórico da implementação por meio de _commits_.

O não cumprimento de algum dos critérios de avaliação especificados poderá resultar nos seguintes decréscimos, aplicados sobre a nota obtida até então na avaliação:

| Falta                                                                                                                                                         | Decréscimo |
| :------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------: |
| Falta de comentários no código fonte                                                                                                                          |       -10% |
| Uso inadequado de controle de versão com Git                                                                                                                  |       -20% |
| Falta de especificação ou especificação incorreta da autoria do trabalho (arquivo [`author.md`](https://github.com/ufrn-golang/lista2/tree/master/author.md)) |       -20% |
| Código fonte com legibilidade prejudicada (por exemplo, com identação ou nomenclatura inadequada)                                                             |       -30% |
| Implementação realizada em desacordo com as especificações postas para o trabalho                                                                             |       -50% |
| Programa apresenta erros de compilação, não executa ou apresenta saída incorreta                                                                              |       -70% |
| Plagiarismo                                                                                                                                                   |      -100% |

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível:

- enviar _e-mail_ para o endereço;
- enviar mensagem privada diretamente ao docente, utilizando o servidor Discord;
- enviar mensagem no canal de texto `#duvidas` no servidor Discord, ou;
- agendar encontros síncronos pelo canal de áudio `Fale com Prof. Everton` no servidor Discord.
