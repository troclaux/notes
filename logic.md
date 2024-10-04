
## introdução às proposições

> proposição lógica: oração declarativa que é verdadeira ou falsa

- oração: sentido completo com presença de verbo
- sentença declarativa: NÃO são proposições as sentenças exclamativas, interrogativas, imperativas ou optativas
- é verdadeira ou falsa: NÃO são proposições as sentenças abertas, paradoxos ou as frases com alta carga de subjetividade
- quantificadores: "todo", "para todo", "nenhum", "existe", etc

## proposições simples

> proposição simples: não pode ser dividida em proposições menores

negação = ~ = ¬

~(~p) = p = ¬(¬p)

## proposições compostas

> proposição que resulta da combinação de duas ou mais proposições simples que são unidas por conectivos

| tipo | conectivos | notação |
| --- | --- | --- | --- |
| conjunção | e/mas | p^q |
| disjunção inclusiva | ou | pvq |
| disjunção exclusiva | ou ..., ou | p⊻q p⊕v |
| condicional | se ..., então/como/logo/implica/quando/toda vez que/somente se/porque/pois | p -> v |
| bicondicional | se e somente se/assim como/se e só se | p <-> v |

p -> q
  - p é condição suficiente para q
  - q é condição necessária para p

> se p, então q => s e => s uficiente

p <-> v = ~(p ⊕ v)

- nem = ^~
  - pedro não estuda nem trabalha = ~p^~q

## conversão da linguagem natural para a proposicional



### ordem de precedência de negação

1. aplicar a negação no menor enunciado possível
2. conjunção(^) ou disjunção(v), na ordem que aparecer
3. disjunção exclusiva(⊻)
4. condicional(->)
5. bicondicional(<->)

## tabela-verdade

VISTO

## tautologia, contradição e contingência

- tautologia: proposição que é sempre verdadeira
  - p v ~p
- contradição: proposição que é sempre falsa
  - p ^ ~p
- contingência: é verdadeira ou falsa dependendo dos valores (V ou F) atribuídos às proposições que compões ela

## equivalências lógicas

p -> q = ~q -> ~p
p -> q = ~p v q
p v q = ~p -> q

~(p ^ q) = ~p v ~q = p -> ~q = q -> ~p
~(p v q) = ~p ^ ~q
~(p -> q) = p ^ ~q

(p -> r) ^ (q -> r) = (p v q) -> r = p -> (q ^ r)

(p <-> q) = (p -> q) ^ (q -> p)

p ^ q = q ^ p
p v q = q v p

## proposições quantificadas e diagramas lógicos

- sentenças abertas: são expressões que possuem um termo variável
  - não há como atribuir-lhes valor lógico (V ou F)
  - não são proposições
    - proposição lógica: oração declarativa que é verdadeira ou falsa
- sentenças fechadas: são expressões que não contém variáveis livres
  - é verdadeiro ou falso

### proposições quantificadas e categóricas

quantificador universal (∀)
 - para todo, para qualquer, qualquer que seja
quantificador existencial (∃)
 - existe, algum, pelo menos um


#### negação de proposições quantificadas

- proposição universal afirmativa
  - exemplo: todo marinheiro é pescados
- proposição universal negativa
  - exemplo: todo brasileiro **não** é mentiroso
- proposição particular afirmativa
  - exemplo: existe um matemático que é engenheiro
- proposição particular negativa
  - exemplo: existe um matemático que **não** é engenheiro

para negar alguma proposição quantificada:
1. encontre o quantificador e o predicado
2. troque o quantificador
  - universal => particular
  - particular => universal
3. troque o predicado pela sua negação
4. negação de proposição concluída

exemplo:
1. Todos os habitantes possuem, pelo menos, uma bicicleta.
  - quantificador universal: todos
  - predicado: possuem, pelo menos, uma bicicleta.
2. Existe um habitante
3. que não possui bicileta
4. Existe um habitante que não possui bicicleta

- predicado: é tudo na oração que se declara sobre o sujeito, seja afirmando ou negando
  - exemplo: todo brasileiro gosta de futebol
    - sujeito = todo brasileiro
    - predicado = gosta de futebol

- proposição universal afirmativa
  - forma A
- proposição universal negativa
  - forma E
- proposição particular afirmativa
  - forma I
- proposição particular negativa
  - forma O

- proposições contrárias
- proposições subcontrárias
- proposições subalternas
- proposições contraditórias


## lógica de primeira ordem (LGO)

- função-predicado/função proposicional: sentença que depende do valor de uma variável para ser classificada como verdadeira ou falsa
- universo do discurso: conjunto formado pelos valores que a variável de uma função-predicado pode assumir

exemplo: x é ímpar
- variável: x
- predicado: é ímpar
- função-predicado: x é ímpar

### LPO e as proposições categóricas

| proposição categórica | representação simbólica |
| --- | --- |
| todo A é B | ∀x( A(x) -> B(x) ) |
| algum A é B | ∃x( A(x) ^ B(x) ) |
| nenhum A é B | ~∃x( A(x) ^ B(x) ) |
| algum A não é B | ∃x( A(x) ^ ~B(x) ) |

### relações e aridade

- relações unárias: 

para resolver questões de conectivos lógicos:
1. identificar as afirmações que se apresentam em algum dos "formatos fáceis"
2. desconsiderar o contexto da questão, transformando as afirmações da língua portuguesa para a linguagem proposicional
3. obter os valores lógicos das proposições simples presentes nas afirmações do enunciado (sempre que possível)
4. verificar a resposta que apresenta uma proposição verdadeira

As afirmações do enunciado que apresentam um "formato fácil" são as seguintes:
- Proposição simples (verdadeira ou falsa)
- Conjunção (e/∧) verdadeira
- Disjunção inclusiva (ou/∨) falsa
- Condicional (se...então/→) falsa
