package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type locale map[string]string

var replaceFilter bool

func main() {
	flag.BoolVar(&replaceFilter, "replace", false, "output with replace filter style")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("required two locales")
	}
	if err := combine(flag.Arg(0), flag.Arg(1), os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func combine(file1, file2 string, w io.Writer) error {
	l1, err := loadLocale(file1)
	if err != nil {
		return err
	}
	l2, err := loadLocale(file2)
	if err != nil {
		return err
	}
	for k, v1 := range l1 {
		v2, ok := l2[k]
		if !ok {
			continue
		}
		if v1 == "" || v2 == "" {
			continue
		}
		switch {
		case replaceFilter:
			fmt.Fprintf(w, "replace_filter %q %q g;\n", regexp.QuoteMeta(v1), v2)
		default:
			fmt.Fprintf(w, "sub_filter %q %q;\n", v1, v2)
		}
	}
	return nil
}

func reload(dst locale, src map[interface{}]interface{}, path []string) error {
	for rk, rv := range src {
		k := fmt.Sprintf("%s", rk)
		curr := append(path, k)
		switch w := rv.(type) {
		case string:
			fullpath := strings.Join(curr, ".")
			dst[fullpath] = w
		case bool:
			fullpath := strings.Join(curr, ".")
			dst[fullpath] = fmt.Sprintf("%t (bool)", w)
		case int:
			fullpath := strings.Join(curr, ".")
			dst[fullpath] = fmt.Sprintf("%d (int)", w)
		case []interface{}:
			fullpath := strings.Join(curr, ".")
			dst[fullpath] = fmt.Sprintf("%v (%T)", w, w)
		case map[interface{}]interface{}:
			err := reload(dst, w, curr)
			if err != nil {
				return err
			}
		default:
			fullpath := strings.Join(append(path, k), ".")
			log.Fatalf("unknown type of value: %s=%T", fullpath, rv)
		}
	}
	return nil
}

func loadLocale(name string) (locale, error) {
	v, err := loadYaml(name)
	if err != nil {
		return nil, err
	}
	l := make(locale)
	if err := reload(l, v, nil); err != nil {
		return nil, err
	}
	return l, nil
}

func loadYaml(name string) (map[interface{}]interface{}, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var v map[interface{}]interface{}
	if err := yaml.Unmarshal(d, &v); err != nil {
		return nil, err
	}
	return v, nil
}
