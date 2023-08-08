package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// define config struct for config.json file
type Config struct {
	ConfigName1 string  `cfg:"config_name1"`
	ConfigName2 string  `cfg:"config_name2"`
	ConfigName3 bool    `cfg:"config_name3"`
	ConfigName4 int64   `cfg:"config_name4"`
	ConfigName5 float64 `cfg:"config_name5"`
}

func ReadFile(configFileName string) {
	file, err := os.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}

	config := parse(file)
	printParsedValue(config)
}

func parse(src []byte) *Config {
	config := &Config{}

	var rawMap map[string]string
	// read config file
	err := json.Unmarshal(src, &rawMap)
	if err != nil {
		panic(err)
	}

	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	n := t.Elem().NumField()
	for i := 0; i < n; i++ {
		field := t.Elem().Field(i)
		fieldVal := v.Elem().Field(i)

		key, ok := field.Tag.Lookup("cfg")
		if !ok || strings.TrimLeft(key, " ") == "" {
			key = field.Name
		}

		value, ok := rawMap[strings.ToLower(key)]
		if ok {
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int64:
				num, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fieldVal.SetInt(num)
				}
			case reflect.Bool:
				boolValue := "true" == value
				fieldVal.SetBool(boolValue)
			case reflect.Float64:
				num, err := strconv.ParseFloat(value, 64)
				if err == nil {
					fieldVal.SetFloat(num)
				}
			default:
				fieldVal.SetString(value)
			}
		}
	}

	return config
}

// print the parsed config value
func printParsedValue(config *Config) {
	fmt.Println("{\n  \"config_name1\": \"配置名称1\",\n  \"config_name2\": \"配置名称2\",\n  \"config_name3\": \"true\",\n  \"config_name4\": \"100\",\n  \"config_name5\": \"10.0\"\n}")
	fmt.Println("**************************")
	fmt.Printf("%+#v", config)
}
