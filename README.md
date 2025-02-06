# Golang

### Run code

```
go run main.go
```

---

### Modules

Módulos são uma coleção de pacotes. Assim como o package.json do NodeJS, o go.mod é o arquivo que contém as informações de dependências do projeto. Para criar um módulo, basta executar o comando abaixo:

```
go mod init package-name
```

### Build

A partir do módulo, podemos gerar um executável do projeto. Para isso, basta executar o comando abaixo:

```go
go build
// go build -o new-name
```

O comando acima irá gerar um executável com o nome do módulo. Para alterar o nome do executável, basta executar o comando abaixo:

```
./binary
```

---

### Privado e Público

Para que um método seja público, basta que o nome do método comece com letra maiúscula. Caso contrário, o método será privado e só poderá ser acessado dentro do pacote.

```go
package sum

/**
 * Public method
 * If the name starts with lowercase, it's private
 */
func Sum(a, b int) int {
	return a + b
}
```

```go
package main

import (
	"fmt"
	"lib/sum"
)

func main() {
	fmt.Println("Hello World")
  // if sum.sum, it will not work because it's private
	fmt.Println("The sum is", sum.Sum(1, 2))
}
```

### Instalando pacotes

Para instalar pacotes, basta executar o comando abaixo:

```
go get package-name
go get github.com/badoux/checkmail
```

Para remover as dependências que não estão sendo utilizadas, basta executar o comando abaixo:

```
go mod tidy
```

# Importações

Este grupo de códigos em parênteses da importação, é a declaração de importação "consignada"

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

Você também pode escrever várias declarações de importação assim

```go
import "fmt"
import "math"
```

## Nomes exportados

Em Go, um nome é exportado se ele começa com uma letra maiúscula. Por exemplo, Pizza é um nome exportado, assim como Pi, que é exportado do pacote math.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // Pi é exportado
	fmt.Println(math.pi) // error pi não é exportado
}
```

## Valores nomeados de retorno

Valores de retorno de Go podem ser nomeados e agirem apenas como variáveis.

Esses nomes devem ser usados para documentar o significado dos valores de retorno.

A declaração return sem argumentos retorna os valores atuais dos resultados. Isto é conhecido como um retorno "limpo".

Instruções de retorno limpas devem ser usadas apenas em funções curtas, como no exemplo mostrado aqui. Elas podem prejudicar a legibilidade em funções mais longas.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17)) // 7 10
}
```

## Erros

- [Go error handling good bad ugly](https://gist.github.com/ljtfreitas/dfbbbb7e03ecb7b02ac5f132d6cabb4b)
