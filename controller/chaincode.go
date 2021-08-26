package controller

import (
	"baas-fabric/service"
	"baas-fabric/util"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

var files = make(map[string]*service.ChainCode,16)

func DeployCC(c *gin.Context)  {
	defer func() {
		if errs :=recover();errs != nil {
			obj :=make(map[string]interface{},16)
			obj["status"] = 400
			obj["msg"] = "出现异常了"
			obj["data"] = nil
			c.JSON(http.StatusBadGateway,obj)
		}
	}()
	name := c.PostForm("codeName")
	version := c.PostForm("version")
	ccp := c.PostForm("ccp")
	ctype := c.PostForm("type")
	channelId := c.PostForm("channelId")
	seq := c.PostForm("seq")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "FAIL")
		return
	}

	md5h  := md5.New()
	mdf,_ := file.Open()
	io.Copy(md5h, mdf)
	md5bt :=md5h.Sum([]byte(""))
	fileId :=fmt.Sprintf("%x",md5bt)
	filename := filepath.Base(file.Filename)
	filepath := "public/upload/" + fileId
	if ok,_ := util.PathExists(filepath); !ok {
		os.Mkdir(filepath,os.ModePerm)
	}
	filepath = filepath + "/" + filename
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		obj :=make(map[string]interface{},16)
		obj["status"] = 400
		obj["msg"] = "fail"
		obj["data"] = nil
		c.JSON(http.StatusBadGateway,obj)
		return
	}
	dest := strings.Split(filepath,".zip")
	chainCodeDest :=dest[0]
	err = util.DeCompressByPath(filepath,chainCodeDest)
	if err!=nil {
		obj :=make(map[string]interface{},16)
		obj["status"] = 400
		obj["msg"] = "文件解压失败"
		obj["data"] = nil
		c.JSON(http.StatusBadGateway,obj)
		return
	}
	ct := pb.ChaincodeSpec_GOLANG
	if strings.ToLower(ctype) == "java" {
		ct = pb.ChaincodeSpec_JAVA
	}
	sint64, _ := strconv.ParseInt(seq, 10, 64)

	chainCode :=&service.ChainCode{
		Seq: sint64,
		Path: chainCodeDest,
		ChannelId: channelId,
		Label: name,
		Version: version,
		Ccp: ccp,
		Type: ct,
	}
	_,ok := files[fileId]
	if !ok {
		files[fileId] = chainCode
	}
	ok ,_ = service.DeployCC(chainCode)
	if !ok {
		obj :=make(map[string]interface{},16)
		obj["status"] = 400
		obj["msg"] = "部署智能合约失败"
		obj["data"] = nil
		c.JSON(http.StatusBadGateway,obj)
		return
	}
	obj :=make(map[string]interface{},16)
	obj["status"] = 200
	obj["msg"] = "SUCCESS"
	obj["data"] = nil
	c.JSON(http.StatusOK,obj)

	return
}

