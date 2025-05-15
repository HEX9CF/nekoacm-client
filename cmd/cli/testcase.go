package cli

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/internal/application/service"
	"nekoacm-common/pkg/utils"
	"os"
	"strings"
)

// 生成测试用例
var TestcaseCmd = &cobra.Command{
	Use:   "testcase",
	Short: "Generate a testcase.",
	Long:  "Generate a testcase.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(" -------- 生成测试用例 -------- ")
		reader := bufio.NewReader(os.Stdin)
		ti, err := readTestcaseInstruction(reader)
		if err != nil {
			return err
		}

		for {
			// 生成测试用例
			fmt.Println("正在生成测试用例...")
			t, err := service.TestcaseGenerate(ti)
			if err != nil {
				log.Println(err)

				// 重试
				err := clearBuffer(reader)
				if err != nil {
					return err
				}
				fmt.Print("生成失败，是否重试(Y/N)?")
				again, _ := reader.ReadString('\n')
				again = strings.TrimSpace(again)
				again = strings.ToLower(again)
				if again != "y" {
					break
				}
				continue
			}

			str, err := utils.PrettyStruct(t)
			if err != nil {
				log.Println(err)
				return err
			}
			fmt.Println(str)

			// 保存到文件
			err = saveTestcaseJson(reader, t)
			if err != nil {
				log.Println(err)
			}

			// 继续生成
			_, err = reader.Discard(reader.Buffered())
			if err != nil {
				return err
			}
			fmt.Print("是否继续生成测试用例(Y/N)?")
			again, _ := reader.ReadString('\n')
			again = strings.TrimSpace(again)
			again = strings.ToLower(again)
			if again != "y" {
				break
			}
		}

		return nil
	},
}

// 读取测试用例信息
func readTestcaseInstruction(reader *bufio.Reader) (dto.TestcaseInstruction, error) {
	ti := dto.TestcaseInstruction{}
	err := clearBuffer(reader)
	if err != nil {
		return dto.TestcaseInstruction{}, err
	}

	// 读取题目信息
	fmt.Println("请输入题目信息：")
	fmt.Print("标题：")
	ti.Title, _ = reader.ReadString('\n')
	ti.Title = strings.TrimSpace(ti.Title)
	fmt.Print("描述：")
	ti.Description, _ = reader.ReadString('\n')
	ti.Description = strings.TrimSpace(ti.Description)
	fmt.Print("输入说明：")
	ti.Input, _ = reader.ReadString('\n')
	ti.Input = strings.TrimSpace(ti.Input)
	fmt.Print("输出说明：")
	ti.Output, _ = reader.ReadString('\n')
	ti.Output = strings.TrimSpace(ti.Output)
	fmt.Print("样例输入：")
	ti.SampleInput, _ = reader.ReadString('\n')
	ti.SampleInput = strings.TrimSpace(ti.SampleInput)
	fmt.Print("样例输出：")
	ti.SampleOutput, _ = reader.ReadString('\n')
	ti.SampleOutput = strings.TrimSpace(ti.SampleOutput)
	fmt.Print("提示：")
	ti.Hint, _ = reader.ReadString('\n')
	ti.Hint = strings.TrimSpace(ti.Hint)
	fmt.Print("标签（以空格分隔）：")
	tagsInput, _ := reader.ReadString('\n')
	ti.Tags = strings.Fields(tagsInput)
	fmt.Print("题解代码：")
	ti.Solution, _ = reader.ReadString('\n')
	ti.Solution = strings.TrimSpace(ti.Solution)

	return ti, nil
}

// 保存测试用例到文件
func saveTestcaseJson(reader *bufio.Reader, t dto.Testcase) error {
	err := clearBuffer(reader)
	if err != nil {
		return err
	}
	fmt.Print("是否保存到文件(Y/N)?")
	save, _ := reader.ReadString('\n')
	save = strings.TrimSpace(save)
	save = strings.ToLower(save)

	if save == "y" {
		path, err := service.TestcaseSaveJson(t)
		if err != nil {
			fmt.Println("保存失败！")
			return err
		}
		fmt.Println("保存成功，文件路径：" + path)
	}

	return nil
}
