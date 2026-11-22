# 🐹 QuizGo

Um jogo de perguntas e respostas desenvolvido em **Go (Golang)**, executado diretamente no terminal.

O projeto foi criado com o objetivo de praticar os fundamentos da linguagem Go, explorando organização de código, manipulação de arquivos CSV, tratamento de erros, concorrência com goroutines e comunicação por channels.

## 🚀 Funcionalidades

- 🎮 Quiz interativo executado no terminal.
- 👤 Identificação do jogador pelo nome.
- 📚 Seleção de categorias.
- 🔢 Questões de múltipla escolha com alternativas numéricas.
- ⏱️ Tempo limite de 5 minutos para concluir o quiz.
- 🏆 Pontuação de 10 pontos por resposta correta.
- ✅ Aprovação ao atingir pelo menos 70% da pontuação total.
- 📂 Leitura de perguntas armazenadas em arquivos CSV.
- 📦 Arquivos CSV incorporados ao executável com `embed.FS`.
- 🎨 Mensagens coloridas no terminal.
- 🧪 Testes automatizados.

## 📚 Categorias disponíveis

| Categoria               | Descrição                                     |
| ----------------------- | --------------------------------------------- |
| 🌎 Conhecimentos Gerais | Perguntas sobre assuntos diversos             |
| 🇬🇧 Inglês               | Questões para testar conhecimentos de inglês  |
| 🔢 Matemática           | Desafios de raciocínio e cálculos matemáticos |

## 🛠️ Tecnologias e conceitos

O projeto utiliza **Go**, sem dependências externas declaradas no `go.mod`.

Durante o desenvolvimento, foram explorados:

- **Structs e Maps:** representação do estado do jogo e das categorias.
- **Slices:** armazenamento das perguntas e alternativas.
- **Goroutines:** execução concorrente do cronômetro e da leitura de respostas.
- **Channels:** comunicação e sinalização entre goroutines.
- **Select:** controle da entrada do jogador e do encerramento do tempo.
- **Embed:** incorporação dos arquivos CSV ao executável.
- **Interfaces:** utilização de `fs.FS` para acesso aos arquivos.
- **Tratamento de erros:** validação de entradas e propagação de erros.
- **Testes:** verificação de funcionalidades com o pacote `testing`.

## 📦 Como executar

### Pré-requisitos

- Go 1.27.3 ou versão compatível com o projeto.
- Git.

### 1. Clone o repositório

```bash
git clone https://github.com/felipedestro/quizgo.git
```

### 2. Acesse a pasta

```bash
cd quizgo
```

### 3. Execute o jogo

```bash
go run .
```

### 4. Responda às perguntas

Informe seu nome, selecione uma categoria e escolha a alternativa correta digitando um número de `1` a `4`.

## 🎯 Regras do jogo

1. O jogador informa seu nome.
2. Seleciona uma das três categorias disponíveis.
3. O cronômetro de **5 minutos** é iniciado.
4. Cada resposta correta vale **10 pontos**.
5. Respostas incorretas não acrescentam pontos.
6. Ao finalizar todas as perguntas, o sistema calcula o desempenho.
7. O jogador é aprovado ao alcançar **70% ou mais** da pontuação máxima.

Caso o tempo termine antes da conclusão, o jogo é encerrado com uma mensagem de tempo esgotado.

## 🧪 Testes

Para executar os testes automatizados:

```bash
go test ./...
```

Para visualizar os resultados detalhadamente:

```bash
go test -v ./...
```

## 📁 Estrutura do projeto

```text
quizgo/
├── internal/
│   ├── data/
│   │   ├── quiz_geral.csv
│   │   ├── quiz_ingles.csv
│   │   └── quiz_matematica.csv
│   ├── game/
│   ├── helper/
│   ├── model/
│   ├── quizfile/
│   └── stopwatch/
├── go.mod
└── main.go
```

**Responsabilidades principais:**

- `game`: gerenciamento do estado, fluxo e resultado do quiz.
- `helper`: leitura e validação das entradas.
- `model`: representação das perguntas.
- `quizfile`: processamento dos arquivos CSV.
- `stopwatch`: controle do tempo por meio de goroutines e channels.
- `data`: arquivos contendo as questões de cada categoria.

## 💡 Objetivo do projeto

O QuizGo é um projeto pessoal de estudo, desenvolvido para consolidar conhecimentos em Go por meio da construção de uma aplicação funcional.

Mais do que implementar um jogo de perguntas e respostas, o objetivo é compreender os recursos da linguagem, experimentar diferentes abordagens e evoluir a qualidade do código com a prática.

## 👨‍💻 Autor

Desenvolvido por **Felipe Destro**.

Desenvolvedor Full Stack, apaixonado por tecnologia e sempre buscando aprimorar seus conhecimentos.

- 🐙 **GitHub:** [@felipedestro](https://github.com/felipedestro)
- 💼 **LinkedIn:** [Felipe Destro](https://www.linkedin.com/in/felipe-destro/)

---

⭐ Se você gostou do QuizGo, considere deixar uma estrela no repositório!
