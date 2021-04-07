package video

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type Packet struct {
}

type ProbeFormat struct {
	NBStreams      int       `json:"nb_streams"`       //多媒体中包含的流的个数
	FormatName     string    `json:"format_name"`      //使用的封装模块的名称
	FormatLongName string    `json:"format_long_name"` //封装的完整的名称
	StartTime      mduration `json:"start_time"`       //媒体文件的起始时间
	Duration       mduration `json:"duration"`         //媒体文件的总时间长度
	BitRate        string    `json:"bit_rate"`         //媒体文件的码率
	Size           humanSize `json:"size"`
}

type ProbeStreams struct {
	Index        int    `json:"index"`          //流所在的索引区域
	CodeName     string `json:"code_name"`      //编码名
	CodeLongName string `json:"code_long_name"` //编码全名
	Profile      string `json:"profile"`
	Level        int    `json:"level"`
	HasBFrame    string `json:"has_b_frame"` //包含B帧信息
	CodecType    string `json:"codec_type"`  //编码类型

}

type ProbeResult struct {
	Format  *ProbeFormat    `json:"format"`
	Streams []*ProbeStreams `json:"streams"`
}

func NewProbeFormat(ctx context.Context, fileName string) (result ProbeResult, err error) {
	var cmdStr string
	switch runtime.GOOS {
	case "linux":
		log.Fatal("linux not support yet, coming soon!")
		err = errors.New("platForm err")
		return
	case "darwin":
		cmdStr = fmt.Sprintf("ffprobe -v quiet -print_format json -show_format %s", fileName)
	default:
		log.Fatal("platform not support yet")
		err = errors.New("platForm err")
		return
	}
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", cmdStr)
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("origin:", cmdOut.String())
	err = json.Unmarshal(cmdOut.Bytes(), &result)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Format)
	fmt.Println(*result.Format)
	return
}

type mduration struct {
	time.Duration
}

func (x *mduration) UnmarshalText(text []byte) error {
	s := string(text) + "s"
	var err error
	x.Duration, err = time.ParseDuration(s)
	return err
}

type humanSize struct {
	string
}

//transfer file size kb to human readable
func (x *humanSize) UnmarshalText(test []byte) error {
	sizeStr := string(test)
	sizeInt64, err := strconv.ParseFloat(sizeStr, 64)
	if err != nil {
		return err
	}

	if sizeInt64 > 1024*1024*1024 {
		sizeGB, err := strconv.ParseFloat(fmt.Sprintf("%.2f", sizeInt64/1024/1024/1024), 64)
		if err != nil {
			return err
		}
		x.string = strconv.FormatFloat(sizeGB, 'f', -1, 64) + "GB"
	} else {
		sizeMB, err := strconv.ParseFloat(fmt.Sprintf("%.2f", sizeInt64/1024/1024), 64)
		if err != nil {
			return err
		}
		x.string = strconv.FormatFloat(sizeMB, 'f', -1, 64) + "MB"
	}
	return nil
}
