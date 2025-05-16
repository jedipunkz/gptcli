package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gptcli/domain/repository"
	openaiinfra "gptcli/infrastructure/openai"
	"gptcli/usecase"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validModels = map[string]bool{
	"gpt-4.1": true,
	"gpt-4o":  true,
	// "o4-mini": true,
	// "o3":      true,
	// "o3-mini": true,
	// "o1":      true,
}

// モデルのエイリアス
var modelAliases = map[string]string{}

// モデル名を解決する関数
func resolveModelName(modelName string) string {
	// モデル名が直接指定されている場合はそのまま返す
	if validModels[modelName] {
		return modelName
	}
	// エイリアスの場合は解決
	if alias, ok := modelAliases[modelName]; ok {
		return alias
	}
	// それ以外の場合はデフォルトのモデルを返す
	return "gpt-4o"
}

var rootCmd = &cobra.Command{
	Use:   "gptcli",
	Short: "OpenAI CLI client",
}

func initDependencies() (repository.ChatRepository, repository.GenerationRepository, *usecase.ChatUseCase, *usecase.GenerationUseCase) {
	apiKey := viper.GetString("api-key")
	if apiKey == "" {
		fmt.Println("Error: API key not set. Use 'gptcli config set api-key KEY' to set it.")
		os.Exit(1)
	}

	chatRepo := openaiinfra.NewChatClient(apiKey)
	genRepo := openaiinfra.NewGenerationClient(apiKey)
	chatUseCase := usecase.NewChatUseCase(chatRepo)
	genUseCase := usecase.NewGenerationUseCase(genRepo)
	return chatRepo, genRepo, chatUseCase, genUseCase
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start a chat conversation",
	Run: func(cmd *cobra.Command, args []string) {
		model, _ := cmd.Flags().GetString("model")
		systemPrompt, _ := cmd.Flags().GetString("system-prompt")

		_, _, chatUseCase, _ := initDependencies()
		chat := chatUseCase.StartChat(resolveModelName(model))

		if systemPrompt != "" {
			chat.AddMessage("system", systemPrompt)
		}

		fmt.Println("Chat started. Type 'exit' to quit.")
		for {
			fmt.Print("You: ")
			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				break
			}

			response, err := chatUseCase.SendMessage(chat, input, resolveModelName(model))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			fmt.Printf("AI: %s\n", response.Content)
		}
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate [prompt]",
	Short: "Generate text from a prompt",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		model, _ := cmd.Flags().GetString("model")
		temperature, _ := cmd.Flags().GetFloat32("temperature")
		maxTokens, _ := cmd.Flags().GetInt("max-completion-tokens")
		stream, _ := cmd.Flags().GetBool("stream")

		_, _, _, genUseCase := initDependencies()

		response, err := genUseCase.GenerateText(args[0], resolveModelName(model), temperature, maxTokens, stream)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if !stream {
			fmt.Println(response)
		}
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		// モデル名のバリデーション
		if key == "default-model" {
			resolvedModel := resolveModelName(value)
			if !validModels[resolvedModel] {
				fmt.Printf("Error: Invalid model name. Available models are:\n")
				fmt.Println("Standard models:")
				for model := range validModels {
					fmt.Printf("  - %s\n", model)
				}
				fmt.Println("\nAliases:")
				for alias, model := range modelAliases {
					fmt.Printf("  - %s (maps to %s)\n", alias, model)
				}
				os.Exit(1)
			}
			// エイリアスが指定された場合は、そのまま保存
			if _, isAlias := modelAliases[value]; isAlias {
				// エイリアスをそのまま保存
			} else {
				// 標準のモデル名の場合は解決後の名前を保存
				value = resolvedModel
			}
		}

		viper.Set(key, value)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Set %s to %s\n", key, value)
	},
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.gptcli")

	// 設定ファイルが存在しない場合は作成
	configPath := filepath.Join(os.Getenv("HOME"), ".gptcli", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// ディレクトリが存在しない場合は作成
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			fmt.Printf("Error creating config directory: %v\n", err)
			os.Exit(1)
		}
		// 空の設定ファイルを作成
		if err := os.WriteFile(configPath, []byte{}, 0644); err != nil {
			fmt.Printf("Error creating config file: %v\n", err)
			os.Exit(1)
		}
	}

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading config file: %v\n", err)
			os.Exit(1)
		}
	}

	chatCmd.Flags().String("model", "gpt-3.5-turbo", "Model to use")
	chatCmd.Flags().String("system-prompt", "", "System prompt to use")

	generateCmd.Flags().String("model", "gpt-3.5-turbo", "Model to use")
	generateCmd.Flags().Float32("temperature", 0.7, "Temperature for generation")
	generateCmd.Flags().Int("max-completion-tokens", 1000, "Maximum completion tokens to generate")
	generateCmd.Flags().Bool("stream", true, "Stream the response")

	configCmd.AddCommand(configSetCmd)
	rootCmd.AddCommand(chatCmd, generateCmd, configCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
