package cmd

import (
	"crgo/infra/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var str string
var mode int8

func init() {
	WordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	WordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下:",
	"1: 全部单词转大写",
	"2: 全部单词转小写",
	"3: 下划线单词转大写驼峰单词",
	"4: 下划线单词转小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var WordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCameLCase(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCameLCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CameCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档")

		}
		log.Printf("输出结果:%s", content)
	},
}
