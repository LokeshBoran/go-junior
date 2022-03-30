package helper

func (gh GeneralHelper) IsNullOrEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch value.(type) {
	case string:
		return value.(string) == ""
	case []string:
		return len(value.([]string)) == 0
	case []int:
		return len(value.([]int)) == 0
	case []int64:
		return len(value.([]int64)) == 0
	case []float64:
		return len(value.([]float64)) == 0
	case []float32:
		return len(value.([]float32)) == 0
	case []bool:
		return len(value.([]bool)) == 0
	case []interface{}:
		return len(value.([]interface{})) == 0
	case map[string]interface{}:
		return len(value.(map[string]interface{})) == 0
	case map[int]interface{}:
		return len(value.(map[int]interface{})) == 0
	case map[int64]interface{}:
		return len(value.(map[int64]interface{})) == 0
	case map[float64]interface{}:
		return len(value.(map[float64]interface{})) == 0
	case map[float32]interface{}:
		return len(value.(map[float32]interface{})) == 0
	case map[bool]interface{}:
		return len(value.(map[bool]interface{})) == 0
	case map[interface{}]interface{}:
		return len(value.(map[interface{}]interface{})) == 0
	default:
		return false
	}
}
