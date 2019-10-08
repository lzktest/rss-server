package main
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/gin-gonic/gin"
		)
func main(){
	resp,err :=http.Get("http://news.sciencenet.cn/htmlnews/2019/9/430434.shtm")
	if err !=nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body,err :=ioutil.ReadAll(resp.Body)
	if err !=nil {
		fmt.Print(err)
	}
	fmt.Print(result)
	r :=gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.String(200,)
	})
	r.Run()
}
func splitRegularString(datastr string)(regularstr []string){
	//分割字符串规则
	
}
