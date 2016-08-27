package obj

import (
	"bufio"
	"regexp"
	"strings"
	"strconv"
	"io"
)

const (
	newMtlPattern = `(?i)^newmtl `
	ambientColorPattern = `(?i)^Ka `
	emissionColorPattern = `(?i)^Ke `
	diffuseColorPattern = `(?i)^Kd `
	specularColorPattern = `(?i)^Ks `
	specularShininessPattern = `(?i)^Ns `
	alpha1Pattern = `(?i)^d `
	alpha2Pattern = `(?i)^Tr `
	ambientColorMapPattern = `(?i)map_Ka `
	emissionColorMapPattern = `(?i)map_Ke `
	diffuseColorMapPattern = `(?i)map_Kd `
	specularColorMapPattern = `(?i)map_Ks `
	specularShininessMapPattern = `(?i)map_Ns `
	normalMapPattern = `(?i)map_Bump `
	alphaMapPattern = `(?i)map_d `

	bitSize64 = 64
)

var (
	newMtlRegex = regexp.MustCompile(newMtlPattern)
	ambientColorRegex = regexp.MustCompile(ambientColorPattern)
	emissionColorRegex = regexp.MustCompile(emissionColorPattern)
	diffuseColorRegex = regexp.MustCompile(diffuseColorPattern)
	specularColorRegex = regexp.MustCompile(specularColorPattern)
	specularShininessRegex = regexp.MustCompile(specularShininessPattern)
	alpha1Regex = regexp.MustCompile(alpha1Pattern)
	alpha2Regex = regexp.MustCompile(alpha2Pattern)
	ambientColorMapRegex = regexp.MustCompile(ambientColorMapPattern)
	emissionColorMapRegex = regexp.MustCompile(emissionColorMapPattern)
	diffuseColorMapRegex = regexp.MustCompile(diffuseColorMapPattern)
	specularColorMapRegex = regexp.MustCompile(specularColorMapPattern)
	specularShininessMapRegex = regexp.MustCompile(specularShininessMapPattern)
	normalMapRegex = regexp.MustCompile(normalMapPattern)
	alphaMapRegex = regexp.MustCompile(alphaMapPattern)
)

type material map[string]interface{}
type materials map[string]*material

func createMaterial() material {
	m := material{}
	m["ambientColor"] = nil                // Ka
	m["emissionColor"] = nil            // Ke
	m["diffuseColor"] = nil                // Kd
	m["specularColor"] = nil            // Ks
	m["specularShininess"] = nil        // Ns
	m["alpha"] = nil                    // d / Tr
	m["ambientColorMap"] = nil            // map_Ka
	m["emissionColorMap"] = nil            // map_Ke
	m["diffuseColorMap"] = nil            // map_Kd
	m["specularColorMap"] = nil            // map_Ks
	m["specularShininessMap"] = nil        // map_Ns
	m["normalMap"] = nil                // map_Bump
	m["alphaMap"] = nil                    // map_d
	return m
}

func getDefaultMtl() *material {
	m := createMaterial();
	m["diffuseColor"] = []float64{0.5, 0.5, 0.5, 1.0}
	return &m
}

func Parse(r io.Reader) materials {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var Material material
	var Materials materials = materials{}
	var line, name string
	for scanner.Scan() {
		line = strings.Trim(scanner.Text(), " ")
		switch {
		// first case always trigger before others
		// https://github.com/AnalyticalGraphicsInc/obj2gltf/commit/68582967ac56100b819988b7ea8bd40c66542138#commitcomment-18800265
		case newMtlRegex.MatchString(line):
			name = strings.Trim(line[7:], " ")
			Material = createMaterial()
			Materials[name] = &Material
		case ambientColorRegex.MatchString(line):
			Material["ambientColor"] = customParse1(line)
		case emissionColorRegex.MatchString(line):
			Material["emissionColor"] = customParse1(line)
		case diffuseColorRegex.MatchString(line):
			Material["diffuseColor"] = customParse1(line)
		case specularColorRegex.MatchString(line):
			Material["specularColor"] = customParse1(line)
		case specularShininessRegex.MatchString(line):
			Material["specularShininess"] = ParseFloat(strings.Trim(line[3:], " "))
		case alpha1Regex.MatchString(line):
			Material["alpha"] = ParseFloat(strings.Trim(line[2:], " "))
		case alpha2Regex.MatchString(line):
			// https://github.com/AnalyticalGraphicsInc/obj2gltf/commit/68582967ac56100b819988b7ea8bd40c66542138#commitcomment-18801135
			//Material["alpha"] = ParseFloat(strings.Trim(line[3:], " "))
			Material["alpha"] = TrHandle(line)
		case ambientColorMapRegex.MatchString(line):
			Material["ambientColorMap"] = strings.Trim(line[7:], " ")
		case emissionColorMapRegex.MatchString(line):
			Material["emissionColorMap"] = strings.Trim(line[7:], " ")
		case diffuseColorMapRegex.MatchString(line):
			Material["diffuseColorMap"] = strings.Trim(line[7:], " ")
		case specularColorMapRegex.MatchString(line):
			Material["specularColorMap"] = strings.Trim(line[7:], " ")
		case specularShininessMapRegex.MatchString(line):
			Material["specularShininessMap"] = strings.Trim(line[7:], " ")
		case normalMapRegex.MatchString(line):
			Material["normalMap"] = strings.Trim(line[9:], " ")
		case alphaMapRegex.MatchString(line):
			Material["alphaMap"] = strings.Trim(line[6:], " ")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if Material["alpha"] != nil {
		if arr, ok := Material["diffuseColor"].([]float64); ok {
			if alpha, ok := Material["alpha"].(float64); ok {
				arr[3] = alpha
			}
		}
	}
	return Materials
}




/*func (m *material)HasDefault() bool {
	panic("!!!!!!!func (m *materials)HasDefault() bool {")
	return
}

func (m *material)Has(s string) bool {
	panic("!!!!!!!func (m *materials)Has() bool {")
	return
}

func (m *material)SetDef(s string) bool {
	panic("!!!!!!!func (m *materials)SetDef() bool {")
	return
}

func (m *material)GetDefault() material {
	panic("!!!!!!!func (m *materials)GetDefault() materials {")
	return
}

func (m *material)Length() int {
	panic("!!!!!!!func (m *materials)Length() materials {")
	return
}

func (m *material)useMaterial() int {
	panic("!!!!!!!func (m *materials)Length() materials {")
	return
}*/

// alias for strconv.ParseFloat with exception
func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, bitSize64)
	if err != nil {
		panic(err)
	}
	return f
}

// parse mtl string in array
func customParse1(s string) []float64 {
	values := strings.Trim(s[3:], " ")
	arrValues := strings.Split(values, " ")
	return []float64{
		ParseFloat(arrValues[0]),
		ParseFloat(arrValues[1]),
		ParseFloat(arrValues[2]),
		1.0,
	}
}

// handle Tr tag
// https://github.com/AnalyticalGraphicsInc/obj2gltf/commit/68582967ac56100b819988b7ea8bd40c66542138#commitcomment-18801135
func TrHandle(s string) interface {} {
	v := strings.Trim(s[3:], " ")

	va := strings.Split(v, " ")
	return ParseFloat(va[0])
}
