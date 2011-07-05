package mapping

import "testing"
import "github.com/bmizerany/assert"
import "reflect"

func TestTest(test *testing.T){	
	v := reflect.ValueOf(url)
	f := reflect.Indirect(v);
	maps := f.Field(0).Interface()
	assert.Equal(test,maps.(map[string]string)["url"],"/")
	assert.Equal(test,maps.(map[string]string)["controller"],"hello")
	assert.Equal(test,maps.(map[string]string)["view"],"index")
}