package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var host = "http://117.51.148.112:9200/"
var index = "jw"
var typeName = "employee"

// Employee ...
type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		fmt.Println("new client err: ", err)
		os.Exit(-1)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

func isExists(id string) bool {
	var exist bool
	exist, _ = client.Exists().Index(index).Type(typeName).Id(id).Do(context.Background())
	if !exist {
		log.Println("ID may be incorrect! ", id)
		return false
	}
	log.Println(id, "is exist")
	return true
}

func create() {
	// 使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().Index("jw").Type("employee").Id("6").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed employee %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	// 使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().Index("jw").Type("employee").Id("7").BodyJson(e2).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed employee %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().Index("jw").Type("employee").Id("8").BodyJson(e3).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed employee %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
}

func delete() {
	res, err := client.Delete().Index("jw").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

func update() {
	res, err := client.Update().Index("jw").Type("employee").Id("2").Doc(map[string]interface{}{"age": 88}).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)
}

func gets() {
	// 通过id查找
	get1, err := client.Get().Index("jw").Type("employee").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

func query() {
	var res *elastic.SearchResult
	var err error
	res, err = client.Search("jw").Type("employee").Do(context.Background())
	printEmployee(res, err)

	// 字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("jw").Type("employee").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)

	if res.Hits.TotalHits.Value > 0 {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits.Value)
		for _, hit := range res.Hits.Hits {
			var t Employee
			err := json.Unmarshal([]byte(hit.Source), &t) // 另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}
			fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}

	// 条件查询 年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("jw").Type("employee").Query(q).Do(context.Background())
	printEmployee(res, err)

	// 短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("jw").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

	// 分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("jw").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)
}

// 简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("jw").Type("employee").Size(size).From((page - 1) * size).Do(context.Background())
	printEmployee(res, err)
}

// 打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var e Employee
	for _, item := range res.Each(reflect.TypeOf(e)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func main() {
	isExists("2")
	// create()
	// delete()
	// update()
	// gets()
	// query()
	// list(3, 1)
}
