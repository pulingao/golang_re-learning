package tools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/iris-contrib/go.uuid"
	jsoniter "github.com/json-iterator/go"
	"hash/crc32"
	"math"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

func ToJson(dd interface{}) string {
	datas, _ := jsoniter.Marshal(dd)
	return string(datas)
}

// 适用于 atm 的spec_text，返回排序后的string
func MyFormatSpecText(str string) string {
	sep := "||"
	sp := strings.Split(str, sep)
	sort.Strings(sp)
	return strings.Join(sp, sep)
}
func MyRawUrlEncode(strin string) string {
	if len(strin) > 0 {
		return strings.Replace(url.QueryEscape(strin), "+", "%20", -1)
	}
	return ""
}
func MyRawUrlDecode(strin string) string {
	if len(strin) > 0 {
		str5 := strings.Replace(strin, "%20", "+", -1)
		str6, err := url.QueryUnescape(str5)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return str6
	}
	return ""
}

func MyUrlEncode(strin string) string {
	if len(strin) > 0 {
		return url.QueryEscape(strin)
	}
	return ""
}

func MyUrlDecode(strin string) string {
	if len(strin) > 0 {
		m, _ := url.QueryUnescape(strin)
		return m
	}
	return ""
}

func MyBase64(strin string) string {
	if len(strin) > 0 {
		return base64.StdEncoding.EncodeToString([]byte(strin))
	} else {
		return ""
	}
}

func MyMd5(strin string) string {
	h := md5.New()
	h.Write([]byte(strin))
	return hex.EncodeToString(h.Sum(nil))
}

func MyCrc(strin string) uint32 {
	if len(strin) < 64 {
		var scratch [64]byte
		copy(scratch[:], strin)
		return crc32.ChecksumIEEE(scratch[:len(strin)])
	}
	return crc32.ChecksumIEEE([]byte(strin))
}

func MaraURI(str_uri string, para_map *map[string]string) {
	ipos := strings.Index(str_uri, "?")
	var str_para []byte
	if ipos >= 0 {
		str_para = []byte(str_uri)[ipos+1:]
	}
	var strarr []string = strings.Split(string(str_para), "&")
	n := len(strarr)

	for i := 0; i < n; i++ {
		var str_para_filed []byte
		var str_para_x []byte
		xpos := strings.Index(strarr[i], "=")
		if xpos >= 0 {
			str_para_filed = []byte(strarr[i])[0:xpos]
			str_para_x = []byte(strarr[i])[xpos+1:]
			(*para_map)[string(str_para_filed)] = string(str_para_x)
		}
	}
}

func DJBHash(tmpstr string) uint64 {
	var hash uint64 = 0
	var hash1 uint64 = 5381

	for i := 0; i < len(tmpstr); i++ {
		hash1 = ((hash1 << 5) + hash1) + uint64([]byte(tmpstr)[i])
	}
	hash1 = hash1 & 0xffffff

	var hash2 uint64 = 0
	for j := 0; j < len(tmpstr); j++ {
		hash2 = uint64([]byte(tmpstr)[j]) + (hash2 << 6) + (hash2 << 16) - hash2
	}
	hash2 = hash2 & 0xffffff
	hash = (hash2 << 24) | hash1
	return hash
}

func TransStrQueryToI64(strkey string, strQuery string) uint64 {
	if len(strkey) < 2 {
		return 0
	}
	i64_res := uint64(0)
	i64_tmp := uint64([]byte(strkey)[0])
	i64_tmp = (i64_tmp << 56)
	i64_res = i64_res | i64_tmp

	i64_tmp = uint64([]byte(strkey)[1])
	i64_tmp = (i64_tmp << 48)
	i64_res = i64_res | i64_tmp

	i64_tmp = DJBHash(strQuery)
	i64_tmp = i64_tmp & 0xffffffffffff
	i64_res = i64_res | i64_tmp
	return i64_res
}

func TransIntQueryToI64(strkey string, i64Query uint64) uint64 {
	if len(strkey) < 2 {
		return 0
	}
	if i64Query > 281474976710655 {
		return 0
	}
	i64_res := uint64(0)
	i64_tmp := uint64([]byte(strkey)[0])
	i64_tmp = (i64_tmp << 56)
	i64_res = i64_res | i64_tmp

	i64_tmp = uint64([]byte(strkey)[1])
	i64_tmp = (i64_tmp << 48)
	i64_res = i64_res | i64_tmp

	i64_tmp = i64Query
	i64_res = i64_res | i64_tmp
	return i64_res
}

func MySysTimeStampM() int64 {
	return time.Now().UnixNano()
}

func MySysTimeStamp() int64 {
	return time.Now().Unix()
}

func MySysTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func MySysTimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func MySysTimeToTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	currentTimeFormat, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), location)
	return currentTimeFormat
}

func MySysTimeDay() string {
	return time.Now().Format("2006-01-02")
}

func MySysTimeDayHour() string {
	return time.Now().Format("2006-01-02_15")
}

func MySysTimeYmdHis() string {
	return time.Now().Format("2006_01_02_15_04_05")
}
func MySysTimeHis() string {
	return time.Now().Format("15_04_05")
}
func MySysTimeYmdDirFileName() string {
	return time.Now().Format("2006/01/02")
}

func StringToStringArray(str string, str_arr *[]string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		//*str_arr = append(*str_arr,tmp_str_arr[i])
		str := strings.TrimSpace(tmp_str_arr[i])
		if str != "" {
			is_repeat := false
			for j := 0; j < len(*str_arr); j++ {
				if (*str_arr)[j] == str {
					is_repeat = true
					break
				}
			}
			if is_repeat == false {
				*str_arr = append(*str_arr, str)
			}
		}
	}
}

func StringToEncodeStringArray(str string, str_arr *[]string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		//*str_arr = append(*str_arr,tmp_str_arr[i])
		str := strings.TrimSpace(tmp_str_arr[i])
		if !strings.Contains(str, "%") {
			tmp_split := strings.Split(str, "/")
			for i := 3; i < len(tmp_split); i++ {
				tmp_split[i] = url.QueryEscape(tmp_split[i])
			}
			str = strings.Join(tmp_split, "/")
		}
		if str != "" {
			*str_arr = append(*str_arr, str)
		}
	}
}

func StringToAsOrEsStringArray(str string, str_arr1 *[]string, str_arr2 *[]string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		//*str_arr = append(*str_arr,tmp_str_arr[i])
		str := strings.TrimSpace(tmp_str_arr[i])
		if strings.Contains(str, "|") {
			strArr := strings.Split(str, "|")
			if len(strArr) > 1 {
				*str_arr2 = append(*str_arr2, MyMd5(strArr[1]+"_"+strArr[0]))
			}
		} else if str != "" {
			*str_arr1 = append(*str_arr1, str)
		}
	}
}

func StringToMd5StringArray(str string, str_arr *[]string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		if strings.Contains(tmp_str_arr[i], ":") { //MAC地址的boss id使用原文
			*str_arr = append(*str_arr, strings.ToUpper(tmp_str_arr[i]))
		} else {
			md5str := strings.ToUpper(MyMd5(strings.ToUpper(tmp_str_arr[i])))
			*str_arr = append(*str_arr, md5str)
		}
	}
}

func StringToIntArray(str string, int_arr *[]int) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		tmp_int, err := strconv.Atoi(tmp_str_arr[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		*int_arr = append(*int_arr, tmp_int)
	}
}

func StrToIntArray(str string, int_arr *[]int) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		tmp_int, err := strconv.Atoi(tmp_str_arr[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		*int_arr = append(*int_arr, tmp_int)
	}
}

func StringToStringValueArray(str string, str_arr *[]string, valueMap map[string]string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		value, exist := valueMap[tmp_str_arr[i]]
		if exist {
			*str_arr = append(*str_arr, value)
		}
		//*str_arr = append(*str_arr,valueMap[tmp_str_arr[i]])
	}
}

func StringToRtStringArray(str string, str_arr *[]string, valueMap map[string]string, accountId string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		if tmp_str_arr[i] == "99999" {
			for key, value := range valueMap {
				tmp := strings.Split(key, "_")
				if len(tmp) > 1 && tmp[1] == accountId {
					strSplit := strings.Split(value, ",")
					*str_arr = append(*str_arr, strSplit...)
					//*str_arr = append(*str_arr, value)
				}
			}
		} else if strings.HasPrefix(tmp_str_arr[i], "99999_") {
			customLabelId := strings.Replace(tmp_str_arr[i], "99999_", "", 1)
			strSplit := strings.Split(valueMap[customLabelId+"_"+accountId], ",")
			*str_arr = append(*str_arr, strSplit...)
		} else {
			*str_arr = append(*str_arr, tmp_str_arr[i])
		}
	}
}

func StringToStringLikeValueArray(str string, str_arr *[]string, valueMap map[string]string, osType string) {
	if len(str) < 3 {
		return
	}
	tmp_str := strings.Replace(str[1:len(str)-1], "\"", "", -1)
	tmp_str = strings.Replace(tmp_str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, ",")
	for i := 0; i < len(tmp_str_arr); i++ {
		//*str_arr = append(*str_arr,valueMap[tmp_str_arr[i]])
		for id, name := range valueMap {
			if strings.HasPrefix(id, tmp_str_arr[i]) {
				*str_arr = append(*str_arr, osType+"_"+name)
			}
		}
	}
}

func IntTostr(i_num int) string {
	return strconv.Itoa(i_num)
}

func Int64Tostr(i_num int64) string {
	return strconv.FormatInt(i_num, 10)
}

func Float64Tostr(f_num float64) string {
	return strconv.FormatFloat(f_num, 'E', -1, 64)
}

// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
func Float64TostrF(f_num float64) string {
	return strconv.FormatFloat(f_num, 'f', -1, 64)
}

func StrToint(str string) int {
	i_num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i_num
}
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

func StrToint64(str string) int64 {
	i_num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i_num
}

func StrTofloat64(str string) float64 {
	f_num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return f_num
}

func GetRandNum(max_num int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(max_num)
}

func GetUUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				//fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func CreateRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func CreateRandomStringNoNumber(l int) string {
	str := "abcdefghjklmnpqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetYmdHisByTimeStamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)

	st := tm.Format("2006-01-02 15:04:05")

	return st
}

func UintToStr(u uint) string {
	b := strconv.Itoa(int(u))
	c := string(b)

	return c
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func StrSplitToIntArray(str string, split string, int_arr *[]int) {
	tmp_str := strings.Replace(str, "\r", "", -1)
	tmp_str = strings.Replace(tmp_str, "\n", "", -1)
	tmp_str_arr := strings.Split(tmp_str, split)
	for i := 0; i < len(tmp_str_arr); i++ {
		tmp_int, err := strconv.Atoi(tmp_str_arr[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		*int_arr = append(*int_arr, tmp_int)
	}
}

// 获取当前日期 前或后的 xx 日 时的时间戳 + 正的就是往后多少日， - 负的就是往前多少日
func GetDaysTimeStamp(day int) (timer int64) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)

	timer = t.AddDate(0, 0, day).Unix()

	return
}

// mc float64 -> int64
func FloatToInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// 切片去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
