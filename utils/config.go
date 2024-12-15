package utils

import (
	"encoding/json"
	"grule_study/models"
	"io/ioutil"
	"log"
	"os"
)

// 读取顶级键
func ReadKeys() (keySlice []string, _ error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		// 如果文件不存在，则初始化一个空的 JSON 对象
		if os.IsNotExist(err) {
			data = []byte("{}")
		} else {
			return nil, err
		}
	}
	var configMap map[string]interface{}
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		return nil, err
	}
	for k, _ := range configMap {
		keySlice = append(keySlice, k)
	}
	return keySlice, nil
}

// CreateGrule 将 RuleForm 保存到动态顶级键下，并直接以 rule_name 为键
func CreateGrule(key string, rule models.RuleForm) error {
	// 读取 config.json 文件
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		// 如果文件不存在，则初始化一个空的 JSON 对象
		if os.IsNotExist(err) {
			data = []byte("{}")
		} else {
			return err
		}
	}

	// 解析 JSON 文件为通用 map
	var configMap map[string]map[string]interface{}
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		return err
	}

	// 检查顶级键是否存在，若不存在则初始化
	if _, exists := configMap[key]; !exists {
		configMap[key] = make(map[string]interface{})
	}

	// 将规则字段组织成 map
	ruleValue := map[string]interface{}{
		"rule_desc":      rule.RuleDesc,
		"rule_salience":  rule.RuleSalience,
		"rule_conditions": rule.RuleConditions,
		"rule_logics":    rule.RuleLogics,
	}

	// 添加或更新规则（使用 rule.RuleName 作为键）
	configMap[key][rule.RuleName] = ruleValue

	// 序列化回 JSON 并保存到文件
	updatedData, err := json.MarshalIndent(configMap, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config.json", updatedData, 0644)
	if err != nil {
		return err
	}

	log.Printf("Rule '%s' under key '%s' added successfully!\n", rule.RuleName, key)
	return nil
}
