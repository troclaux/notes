
## lógica proposicional

> proposição lógica: oração declarativa que é verdadeira ou falsa

- oração: sentido completo com presença de verbo
- sentença aberta: sentença em que não se pode determinar se é verdadeira ou falsa
- sentença declarativa: NÃO são proposições as sentenças exclamativas, interrogativas, imperativas ou optativas
- é verdadeira ou falsa: NÃO são proposições as sentenças abertas, paradoxos ou as frases com alta carga de subjetividade
- quantificadores: "todo", "para todo", "nenhum", "existe", etc

### proposições simples

> [!NOTE]
> sentenças abertas não são proposições
> e.g. `x + 9 = 10`
> de um ponto de vista lógico, não podemos afirmar que esta sentença é verdadeira ou falsa

- proposição simples: não pode ser dividida em proposições menores

> [!NOTE]
> frases que exprimem opinião não são proposições

- negação: `~`, `¬`

```
~(~p) = p = ¬( ¬p)
```

### proposições compostas

> proposição que resulta da combinação de duas ou mais proposições simples que são unidas por conectivos

| tipo | conectivos | notação |
| --- | --- | --- |
| conjunção | e/mas | p^q |
| disjunção inclusiva | ou | pvq |
| disjunção exclusiva | ou ..., ou | p⊻q p⊕v |
| condicional | se ..., então/como/logo/implica/quando/toda vez que/somente se/porque/pois | p -> v |
| bicondicional | se e somente se/assim como/se e só se | p <-> v |

- nem = `^~`
  - pedro não estuda nem trabalha = `~p^~q`
    - p: pedro estuda
    - q: pedro trabalha

- condicional: `p -> q`
  - p é:
    - condição suficiente
    - antecedente
    - precedente
  - q é:
    - condição necessária
    - consequente
    - subsequente

> [!TIP]
> se p, então q => s e => s uficiente

```
p <-> v = ~(p ⊕ v)
```

### ordem de precedência de negação

1. aplicar a negação no menor enunciado possível
2. conjunção `^` ou disjunção `v`, na ordem que aparecer
3. disjunção exclusiva `⊻` ou `⊕`
4. condicional `->`
5. bicondicional `<->`

### tabelas-verdade

condicional

| p | q | p -> q |
| --- | --- | --- |
| V | V | V |
| V | F | F |
| F | V | V |
| F | F | V |

disjunção exclusiva/x-or

| p | q | p ⊕ q |
| --- | --- | --- |
| V | V | F |
| V | F | V |
| F | V | V |
| F | F | F |

### tautologia, contradição e contingência

- tautologia: proposição que é sempre verdadeira
  - e.g. `p v ~p`
- contradição: proposição que é sempre falsa
  - e.g. `p ^ ~p`
- contingência: é verdadeira ou falsa dependendo dos valores (V ou F) atribuídos às proposições que compões ela

### equivalências lógicas

```
p -> q = ~q -> ~p
p -> q = ~p v q
p v q = ~p -> q

~(p ^ q) = ~p v ~q = p -> ~q = q -> ~p
~(p v q) = ~p ^ ~q

~(p -> q) = p ^ ~q

(p -> r) ^ (q -> r) = (p v q) -> r

(p -> q) ^ (p -> r) = p -> (q ^ r)

p ^ q = q ^ p
p v q = q v p

p <-> q = (p -> q) ^ (q -> p)
p <-> q = ~(p ⊻ q)
p ⊻ q = (~p) ⊻ (~q) = (~p) <-> q = p <-> (~q)
```

## Lógica de Primeira Ordem (LPO)

### proposições quantificadas e diagramas lógicos

- sentenças abertas: são expressões que possuem um termo variável
  - não há como atribuir-lhes valor lógico (V ou F)
  - não são proposições
    - proposição lógica: oração declarativa que é verdadeira ou falsa
- sentenças fechadas: são expressões que não contém variáveis livres
  - é verdadeiro ou falso

- quantificador universal (∀) = para todo, para qualquer, qualquer que seja
- quantificador existencial (∃) = existe, algum, pelo menos um

### negação de proposições quantificadas

- proposição universal afirmativa
  - exemplo: todo marinheiro é pescados
- proposição universal negativa
  - exemplo: todo brasileiro **não** é mentiroso
- proposição particular afirmativa
  - exemplo: existe um matemático que é engenheiro
- proposição particular negativa
  - exemplo: existe um matemático que **não** é engenheiro

para negar alguma proposição quantificada:

- proposição universal afirmativa <=> proposição particular negativa
- proposição universal negativa <=> proposição particular afirmativa

passo a passo:

1. encontre o quantificador e o predicado
2. troque o quantificador
  - universal => particular
  - particular => universal
3. troque o predicado pela sua negação
4. negação de proposição concluída

exemplos:

1. Todos os habitantes possuem, pelo menos, uma bicicleta.
  - quantificador universal: todos
  - predicado: possuem, pelo menos, uma bicicleta.
2. Existe um habitante
3. que não possui bicileta
4. Existe um habitante que não possui bicicleta

- p: todo brasileiro gosta de futebol
- ~p: existe algum brasileiro que não gosta de futebol
- q: qualquer pessoa consegue passar
- ~q: alguma pessoa não consegue passar = existe pessoa que não consegue passar

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

> [!TIP]
> as letras das formas seguem a ordem das vogais (AEIO)
> começamos com proposições universais e terminamos com particulares
> alternamos entre afirmativa e negativa
> A de Afirmativa
> E de nEgativa

- proposições contrárias
- proposições subcontrárias
- proposições subalternas
- proposições contraditórias

- função-predicado/função proposicional: sentença que depende do valor de uma variável para ser classificada como verdadeira ou falsa
- universo do discurso: conjunto formado pelos valores que a variável de uma função-predicado pode assumir

- exemplo: x é ímpar
  - variável: x
  - predicado: é ímpar
  - função-predicado: x é ímpar

### proposições categóricas

| proposição categórica | representação simbólica |
| --- | --- |
| todo A é B | ∀x( A(x) -> B(x) ) |
| algum A é B | ∃x( A(x) ^ B(x) ) |
| nenhum A é B | ~∃x( A(x) ^ B(x) ) |
| algum A não é B | ∃x( A(x) ^ ~B(x) ) |

> [!WARNING]
> se o negativo vier antes do quantificador, ele se aplica no quantificador
> se o negativo vier antes do predicado, ele se aplica no predicado

### conectivos lógicos

como resolver questões:

1. identificar as afirmações que estão no "formato fácil"
2. desconsiderar o contexto da questão, transformando as afirmações da língua portuguesa para a linguagem proposicional
3. obter os valores lógicos das proposições simples presentes nas afirmações do enunciado (sempre que possível)
4. verificar a resposta que apresenta uma proposição verdadeira

As afirmações do enunciado que apresentam um "formato fácil" são as seguintes:
- Proposição simples (verdadeira ou falsa)
- Conjunção (e/∧) verdadeira
- Disjunção inclusiva (ou/∨) falsa
- Condicional (se...então/→) falsa

## lógica de argumentação

### argumentos dedutivos

- argumento dedutivo: argumento que não produz conhecimento novo
- silogismo: argumento dedutivo composto por duas premissas e uma conclusão

- existem 2 tipos de argumento:
  - argumentos categóricos: argumentos que apresentam proposições categóricas
  - argumentos hipotéticos: argumentos que não apresentam proposições categóricas
    - fazem uso de conectivos

- validade: característica dos argumentos dedutivos, que pode ser válido ou inválido
  - conclusão válida: segue logicamente das premissas, mesmo que as premissas sejam falsas
- veracidade: característica das proposições, que podem ser verdadeiras ou falsas

- argumento dedutivo é valido quando as premissas verdadeiras levam à uma conclusão que é necessariamente verdadeira
- argumento dedutivo é inválido quando premissas consideradas verdadeiras não necessariamente levam à conclusão verdadeira
  - argumento dedutivo inválido = sofisma ou falácia formal

- é possível ter um argumento válido nesses cenários:
  - premissas verdadeiras e conclusão verdadeira
  - premissas falsas e conclusão verdadeira
  - premissas falsas e conclusão falsa
- não é possível ter um argumento válido com premissas verdadeiras e conclusão falsa
- é possível ter um argumento inválido nesses cenários:
  - premissas verdadeiras e conclusão verdadeira
  - premissas verdadeiras e conclusão falsa
  - premissas falsas e conclusão verdadeira
  - premissas falsas e conclusão falsa
- não há uma relação direta entre a validade de um argumento e a veracidade da sua conclusãoválido nesses cenários:


### regra da transitividade condicional

argumento no formato abaixo, independentemente do número de premissas, é sempre válido

```
a -> b
b -> c
c -> d
d -> e

é equivalente a

a -> e
```

### regras de inferência

dadas as premissas, podemos deduzir as conclusões abaixo:

- modus ponens (afirmação do antecedente)
  - premissa 1: p -> q
  - premissa 2: p
  - conclusão: q
- modus tollens (negação do consequente)
  - premissa 1: p -> q
  - premissa 2: ~q
  - conclusão: ~p
- silogismo hipotético
  - premissa 1: p -> q
  - premissa 2: q -> r
  - conclusão: p -> r
- dilema construtivo ou silogismo disjuntivo
  - premissa 1: p -> q
  - premissa 2: r -> s
  - premissa 3: p v r
  - conclusão: q v s
- dilema destrutivo
  - premissa 1: p -> q
  - premissa 2: r -> s
  - premissa 3: ~q ou ~s
  - conclusão: ~p ou ~r
