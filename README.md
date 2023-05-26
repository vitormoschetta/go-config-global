# Configurações Globais

Quando queremos compartilhar um pacote de configurações globais em nosso projeto temos alguns caminhos a seguir, como por exemplo:

## Variáveis Globais
Criar um arquivo de configuração com variáveis exportáveis (primeira letra maiúscula) e importar em todos os arquivos que precisam dessas configurações.

#### Pros
- Simples de implementar

#### Contras
- Os valores não são imutáveis, ou seja, podem ser alterados em qualquer lugar do código.


## Constantes
Criar um arquivo de configuração com constantes e importar em todos os arquivos que precisam dessas configurações.

#### Pros
- Simples de implementar

#### Contras
- É necessário definir os valores das contantes no momento da criação, estaticamente.


## Singleton
Criar uma struct com um método de acesso para a construção e outro método de acesso para a obtenção da instância com seus valores.

#### Pros
- Os valores são imutáveis, ou seja, não podem ser alterados em qualquer lugar do código. 
- Não é necessário definir os valores de configuração no momento da criação, podendo ser definidos dinamicamente.

#### Contras
- Complexidade de implementação

Obs: Apesar de não parecer algo complexo, devemos levar em conta que é necessário uma complexidade maior para casos de singleton thread-safe, por exemplo.