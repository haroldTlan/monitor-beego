package models

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"sort"
	"strconv"
	_ "strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/lifei6671/mindoc/utils"
)

type Extract struct {
	Features interface{} `json:"features"`
	TimeUsed int         `json:"time_used"`
}

type Compare struct {
	Confidence float64 `json:"confidence"`
	TimeUsed   int     `json:"time_used"`
}

type SimilarFile struct {
	Similarity string `json:"similarity"`
	Fid        string `json:"feature"`
}

type One2ManyResult struct {
	Status string `json:"result"`
	TaskId string `json:"taskid"`
	Total  int    `json:"total"`
	Datas  []Data `json:"data"`
}

type Data struct {
	Age         int       `json:"age"`
	Alias       string    `json:"aliasname"`
	Attacharea  int       `json:"attacharea"`
	Createtime  time.Time `json:"creattime"`
	CrimeType   string    `json:"crimetype"`
	Education   string    `json:"education"`
	Gender      string    `json:"gender"`
	Height      int       `json:"height"`
	IdentityId  string    `json:"identityId"`
	ImgbaseId   int64     `json:"imgbaseid"`
	Name        string    `json:"name"`
	Nation      string    `json:"nation"`
	Occupation  string    `json:"occupation"`
	Photourl    string    `json:"photourl"`
	Recruitment string    `json:"recruitment"`
	Remark      string    `json:"remark"`
	Score       float64   `json:"score"` //string
	Uuid        string    `json:"uuid"`
	Width       int       `json:"width"`
}

type One2ManyDatas []Data

func (r One2ManyDatas) Len() int           { return len(r) }
func (r One2ManyDatas) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r One2ManyDatas) Less(i, j int) bool { return r[i].Score > r[j].Score }

// New Compare
func One2Many(threhold, lib string, file multipart.File) (res One2ManyResult, err error) {
	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	path := "/tmp/img.jpg"

	img, err := os.Create(path)
	defer img.Close()

	if _, err = io.Copy(img, file); err != nil {
		return
	}

	// Ensure
	if err = EnsureExist(path); err != nil {
		return
	}

	/*
		var cmd []string
		cmd = append(cmd, "compare.py", threhold, lib, path)
	*/

	//o, err := utils.Execute("./compare", cmd, true)
	//cmd := "python getfeature.py photos2/mark1.jpg"

	/*
		name := "/home/administrator/lib/build_SampleApp_Release/compare"
				name := "python"
				filepath := "/home/administrator/lib/build_SampleApp_Release"
				o, err := utils.NewExecute(filepath, name, cmd, true)

			fmt.Println(o, err)
			o = strings.TrimSpace(o)

			json.Unmarshal([]byte(o), &res)
	*/

	var cmd []string
	cmd = append(cmd, "compare.py", threhold, lib, path)
	filepath := InstallFilePath
	name := "python"

	o, err := utils.Execute(filepath, name, cmd, true)
	var ss []SimilarFile
	json.Unmarshal([]byte(o), &ss)

	//os.Remove(path)

	var datas One2ManyDatas
	for _, i := range ss {

		//sss := []SimilarFile{SimilarFile{Similarity: 0.902, Fid: "55,4811afd6cab6"}, SimilarFile{Similarity: 0.632, Fid: "55,4811acab6"}}
		//for _, i := range sss {
		var d Data
		p, _ := GetPictureByFeature(i.Fid)

		d.Remark = p.LibraryName
		d.ImgbaseId = p.LibraryId
		d.Gender = p.Gender
		d.Name = p.TargetName
		d.Photourl = p.Url

		score, _ := strconv.ParseFloat(i.Similarity, 64)
		d.Score = score

		d.Createtime = time.Now()
		d.Height = 143
		d.Width = 117
		datas = append(datas, d)
	}

	sort.Sort(datas)
	res.Datas = datas
	res.Total = len(ss)
	res.Status = "success"

	return
}

func EnsureExist(path string) (err error) {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return
}
