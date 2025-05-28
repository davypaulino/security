# **Criptografias IFSP**  
_Repositório dedicado à implementação de algoritmos de criptografia para atividades da disciplina de Segurança da Informação no IFSP._

## 🔹 **Funcionalidades Implementadas**
✔️ **Cifra de César**  
- [X] 🔐 Encriptar texto usando Cifra de César  
- [X] 🔓 Desencriptar texto usando Cifra de César  
- [X] 🧐 Quebra da Cifra de César por Análise de Frequência  

✔️ **Cifra de Vigenère**  
- [X] 🔐 Encriptar texto usando Cifra de Vigenère  
- [X] 🔓 Desencriptar texto usando Cifra de Vigenère  
- [ ] 🧐 Quebra da Cifra de Vigenère por Análise de Frequência  

## Texto para teste Cesar:

```txt
A Travessia do Tempo

O tempo avanca como um rio, constante e inexoravel. Desde o instante em que nascemos, somos lancados na correnteza da existencia, navegando por aguas calmas e tempestades furiosas. Cada dia que passa carrega consigo historias invisiveis, momentos que se acumulam como paginas de um livro que nunca sabemos ao certo como terminara.

A infancia, esse territorio magico onde tudo e possivel, nos ensina a ver o mundo com olhos de descoberta. Cada folha caída e um misterio, cada sombra na parede um universo paralelo. A imaginacao e fertil e ilimitada, e os dias passam como sonhos entre brincadeiras e risadas. Mas, tal como as estacoes do ano, tudo muda. Crescemos, e com o crescimento, vem as responsabilidades, os desafios e as reflexoes profundas sobre quem somos e para onde estamos indo.

A juventude e marcada por batalhas internas e externas. Buscamos identidade, proposito, pertencimento. Tentamos entender nossos proprios sentimentos, experimentar novas ideias, quebrar padroes e, as vezes, recomecar do zero. E um periodo de intensidade, onde cada emocao parece um furacao e cada decisao ecoa como um trovao no horizonte da vida. Porem, com o tempo, aprendemos que a paz nao esta na ausencia de conflito, mas na compreensao de nossas proprias escolhas.

E entao, vem a maturidade. Olhamos para tras e percebemos que os desafios do passado moldaram nossas jornadas. As memorias tornam-se tesouros, as cicatrizes lembrancas de batalhas superadas. A busca pela felicidade passa a ser menos sobre grandes conquistas e mais sobre apreciar os momentos simples: um cafe ao amanhecer, uma conversa sincera, o abraco de quem amamos.

Mas o tempo nao para. A velhice chega como um sussurro, trazendo consigo uma sabedoria que so pode ser adquirida pela experiencia. O peso dos anos ensina a valorizar o presente, a entender que o agora e tudo o que realmente temos. Aprendemos que as perguntas mais importantes da vida muitas vezes nao tem respostas definitivas – e que esta tudo bem.

O fim da jornada e inevitavel, mas nao deve ser temido. Somos como estrelas que brilham intensamente por um periodo e, eventualmente, retornamos ao infinito do cosmos. O que permanece sao os rastros que deixamos, as historias que compartilhamos, as licoes que transmitimos.

Que possamos viver cada dia com autenticidade, amor e coragem, sabendo que nossa travessia, por mais efemera que seja, deixa marcas eternas no tecido do universo.
```

## Texto para teste Vigenere

**Chave:**
```txt
EuvousaireuprecisomemudarEuvoulaprointeriorvouviverpertodemimEuvousubirnotremquevaipartirdaquiEuprecisosaber
```

**Menssagem:**
```txt
A Travessia do Tempo

O tempo avanca como um rio, constante e inexoravel. Desde o instante em que nascemos, somos lancados na correnteza da existencia, navegando por aguas calmas e tempestades furiosas. Cada dia que passa carrega consigo historias invisiveis, momentos que se acumulam como paginas de um livro que nunca sabemos ao certo como terminara.

A infancia, esse territorio magico onde tudo e possivel, nos ensina a ver o mundo com olhos de descoberta. Cada folha caida e um misterio, cada sombra na parede um universo paralelo. A imaginacao e fertil e ilimitada, e os dias passam como sonhos entre brincadeiras e risadas. Mas, tal como as estacoes do ano, tudo muda. Crescemos, e com o crescimento, vem as responsabilidades, os desafios e as reflexoes profundas sobre quem somos e para onde estamos indo.

A juventude e marcada por batalhas internas e externas. Buscamos identidade, proposito, pertencimento. Tentamos entender nossos proprios sentimentos, experimentar novas ideias, quebrar padroes e, as vezes, recomecar do zero. E um periodo de intensidade, onde cada emocao parece um furacao e cada decisao ecoa como um trovao no horizonte da vida. Porem, com o tempo, aprendemos que a paz nao esta na ausencia de conflito, mas na compreensao de nossas proprias escolhas.

E entao, vem a maturidade. Olhamos para tras e percebemos que os desafios do passado moldaram nossas jornadas. As memorias tornam-se tesouros, as cicatrizes lembrancas de batalhas superadas. A busca pela felicidade passa a ser menos sobre grandes conquistas e mais sobre apreciar os momentos simples: um cafe ao amanhecer, uma conversa sincera, o abraco de quem amamos.

Mas o tempo nao para. A velhice chega como um sussurro, trazendo consigo uma sabedoria que so pode ser adquirida pela experiencia. O peso dos anos ensina a valorizar o presente, a entender que o agora e tudo o que realmente temos. Aprendemos que as perguntas mais importantes da vida muitas vezes nao tem respostas definitivas e que esta tudo bem.

O fim da jornada e inevitavel, mas nao deve ser temido. Somos como estrelas que brilham intensamente por um periodo e, eventualmente, retornamos ao infinito do cosmos. O que permanece sao os rastros que deixamos, as historias que compartilhamos, as licoes que transmitimos.

Que possamos viver cada dia com autenticidade, amor e coragem, sabendo que nossa travessia, por mais efemera que seja, deixa marcas eternas no tecido do universo.
```