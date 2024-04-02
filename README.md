# Vocabulary Quiz

This is a basic Vocabulary Quiz application written in Go. It challenges users with vocabulary questions sourced from a CSV file, making it a great tool for those looking to improve their language skills.

## Prerequisites

To run this application, you need to have Go installed on your system. If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/).

## Installation

1. Clone the repository to your local machine:
   ```
   git clone https://github.com/YourUsername/vocabulary_quiz.git
   ```
2. Navigate to the cloned repository:
   ```
   cd vocabulary_quiz
   ```

## Compiling the Application

To compile the application, run the following command in the root of the project directory:

```
go build -o quiz quiz.go
```

This command compiles the `quiz.go` file and generates an executable named `quiz`.

## Making the Binary Executable

After compiling, you may need to make the binary executable. Run the following command:

```
chmod +x quiz
```

This command grants execution permissions to the `quiz` binary.

## Running the Application

To run the quiz, execute the binary:

```
./quiz
```

## How It Works

The application reads a list of vocabulary questions and their answers from a `words.csv` file and presents them to the user in a quiz format. Users are prompted to enter answers, and the application provides feedback based on their responses.

## Contributing

Contributions to the Vocabulary Quiz are welcome! Please feel free to submit pull requests or open issues to improve the application or suggest new features.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
