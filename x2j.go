// x2j - for (mostly) backwards compatibility with x2j package

package mxj

// Wrappers for end-to-end XML to JSON transformation.

/*
import "io"

// FromXml() --> ToJson().
func XmlToJson(xmlVal []byte) ([]byte, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.Json()
}

// FromXml() --> ToJsonWriter().
func XmlToJsonWriter(xmlVal []byte, jsonWriter io.Writer) (*[]byte, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.JsonWriter(jsonWriter)
}

// FromXmlReader() --> ToJson().
func XmlReaderToJson(xmlReader io.Reader) ([]byte, *[]byte, error) {
	m, raw, merr := NewMapXmlReader(xmlReader)
	if merr != nil {
		return nil, raw, merr
	}
	j, jerr := m.Json()
	return j, raw, jerr
}

// FromXmlReader() --> ToJsonWriter().  Handy for bulk transformation of documents.
func XmlReaderToJsonWriter(xmlReader io.Reader, jsonWriter io.Writer) (*[]byte, *[]byte, error) {
	m, xraw, merr := NewMapXmlReader(xmlReader)
	if merr != nil {
		return xraw, nil, merr
	}
	jraw, jerr := m.JsonWriter(jsonWriter)
	return xraw, jraw, jerr
}
*/

// XML wrappers for Map methods implementing tag path and value functions.

/*
// Wrap PathsForKey for XML.
func XmlPathsForTag(xmlVal []byte, tag string) ([]string, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	paths := m.PathsForKey(tag)
	return paths, nil
}

// Wrap PathForKeyShortest for XML.
func XmlPathForTagShortest(xmlVal []byte, tag string) (string, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return "", merr
	}
	path := m.PathForKeyShortest(tag)
	return path, nil
}

// Wrap ValuesForKey for XML.
// 'attrs' are key:value pairs for attributes, where key is attribute label prepended with a hypen, '-'.
func XmlValuesForTag(xmlVal []byte, tag string, attrs ...string) ([]interface{}, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.ValuesForKey(tag, attrs...)
}

// Wrap ValuesForPath for XML.
// 'attrs' are key:value pairs for attributes, where key is attribute label prepended with a hypen, '-'.
func XmlValuesForTagPath(xmlVal []byte, path string, attrs ...string) ([]interface{}, error) {
	m, merr := NewMapXml(xmlVal)
	if merr != nil {
		return nil, merr
	}
	return m.ValuesForPath(path, attrs...)
}
*/

