# unshort

Ferramenta de linha de comando para observar o destino de uma URL encurtada sem a necessidade de abrir em um browser.

## Como Usar?

```bash
unshort https://bit.ly
```

Resultado:

```
REFERER:        https://bit.ly
LOCATION:       https://bitly.com/
```

## Como Criar?

Você pode criar o binário de 2 formas distintas:

### Com Docker

#### Pré Requisitos

- [Docker](https://docker.com/)

#### Como Fazer?

1. Clone este repositório e navegue até a pasta raiz.
2. Execute o comando abaixo no terminal:

```bash
docker run -ti --rm \
  -v $(pwd):/unshort \
  -w /unshort \
  golang:1.19 sh -c "go build -o ./bin/unshort ./cmd"
```

**Obs**: o mesmo comando pode ser feito executando o arquivo `commands`, da seguinte forma:

```bash
./commands
```

### Com Go

#### Pré Requisitos

- [Go](https://go.dev/)

#### Como Fazer?

1. Clone este repositório e navegue até a pasta raiz.
2. Execute o comando abaixo no terminal:

```bash
go build -o ./bin/unshort ./cmd
```

---

