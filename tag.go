package obj3d

/*
typedef struct {
  std::string name;

  std::vector<int> intValues;
  std::vector<float> floatValues;
  std::vector<std::string> stringValues;
} tag_t;
*/

type TagT struct {
	Name         string

	IntValues    []int
	FloatValues  []float64
	StringValues []string
}
