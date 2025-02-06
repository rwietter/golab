# Manipulação de Strings e Conversão de Tipos

- strings: Contém funções para manipulação de strings, como strings.ToLower(), strings.ToUpper(), strings.Split(), e muitas outras.

## Com Strings

- len(str): Para obter o tamanho de uma string.
- strings.ToUpper(str): Para converter uma string em letras maiúsculas.
- strings.ToLower(str): Para converter uma string em letras minúsculas.
- strings.Split(str, separator): Para dividir uma string em partes com base em um separador.
- strings.Contains(str, substr): Para verificar se uma string contém uma substring.
- strings.TrimSpace(str): Para remover espaços em branco no início e no final de uma string.

## Com Números

- strconv.Itoa(num): Para converter um número inteiro em uma string.
- strconv.ParseInt(str, base, bitsize): Para analisar uma string e obter um número inteiro.
- strconv.ParseFloat(str, bitSize): Para analisar uma string e obter um número de ponto flutuante.

## Com Mapas

- make(map[type]type): Para criar um novo mapa.
- map[key] = value: Para atribuir um valor a uma chave no mapa.
- val, ok := map[key]: Para verificar se uma chave existe no mapa e obter o valor associado.
- delete(map, key): Para excluir uma chave e seu valor do mapa.

## Com Structs

- Acesso a campos: Você pode acessar os campos de um struct diretamente usando o ponto, por exemplo, myStruct.Nome.
- Inicialização: Para criar uma instância de um struct, você pode simplesmente declará-lo e atribuir valores aos campos, por exemplo, pessoa := Pessoa{Nome: "Maurício", Idade: 30}.

# Entrada e Saída

- fmt: Usado para formatação de entrada e saída, incluindo fmt.Println(), fmt.Sprintf(), e fmt.Scan().
- io: Fornece interfaces e funções para entrada/saída, como leitura e escrita em arquivos.
  
# Conversão de Tipos

- strconv: Usado para conversão de tipos, incluindo funções como strconv.Atoi(), strconv.Itoa(), e strconv.ParseFloat().

# Manipulação de Dados

- container: Oferece tipos de dados como listas ligadas e pilhas.
- sort: Possui funções para classificação de slices.
- encoding/json: Para codificação e decodificação de dados em formato JSON.
- encoding/xml: Para codificação e decodificação de dados em formato XML.
- encoding/csv: Para leitura e escrita de dados em formato CSV.
- encoding/gob: Para codificação e decodificação de dados no formato gob.
- Marshal: Para codificação de dados em formato JSON, XML, CSV, ou gob.
- Unmarshal: Para decodificação de dados em formato JSON, XML, CSV, ou gob.

# Tempo e Data

- time: Usado para trabalhar com datas e horários, incluindo time.Now(), time.Parse(), e time.Format().

# Concorrência

- sync: Fornece mecanismos de sincronização, como Mutexes e Cond.
- goroutine: Pacote principal para criação e gerenciamento de goroutines.

# Rede

- net: Para operações de rede, incluindo conexões TCP e UDP, resolução de DNS, etc.

# Web

- http: Usado para criar servidores HTTP e clientes, incluindo http.ListenAndServe(), http.Get(), etc.
- html/template: Para renderização de modelos HTML.