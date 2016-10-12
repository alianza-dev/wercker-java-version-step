package main

import (
  "io"
  "io/ioutil"
  "fmt"
  "encoding/xml"
  "os"
  "strings"
  "regexp"
)

type Project struct {
  XMLName xml.Name `xml:"project"`
  ModelVersion string `xml:"modelVersion"`
  GroupId string `xml:"groupId"`
  ArtifactId string `xml:"artifactId"`
  Packaging string `xml:"packaging"`
  Version string `xml:"version"`
  Name string `xml:"name"`
}

func main() {
  pomFile := GetEnvOrDefault("WERCKER_VERSIONING_POM", "pom.xml")
  outFile := GetEnvOrDefault("WERCKER_VERSIONING_OUTFILE", "build_source")
  timeStamp := GetEnvOrDefault("WERCKER_MAIN_PIPELINE_STARTED", "NONE")

  pom := readPom(pomFile)

  majorVersion, artifactId, mvnVersion, groupId := Version(pom, timeStamp)

  sourceContent := FormatBashSource(majorVersion, artifactId, mvnVersion, groupId)
  writeContents(outFile, sourceContent)
}

func GetEnvOrDefault(envName string, defaultValue string) (string) {
  envValues := os.Environ()
  for _, value := range envValues {
    name := strings.Split(value, "=")[0]
    if name == envName {
      return os.Getenv(envName)
    }
  }

  return defaultValue
}

func readPom(pomFile string) (content string) {
  pomContent, err := ioutil.ReadFile(pomFile)
  if err != nil {
    fmt.Print(err)
    content = ""
  } else {
    content = string(pomContent)
  }

  return content
}

func Version(pomContents string, timestamp string) (majorVersion string, artifactId string, mvnVersion string, groupId string) {
  var pom Project
  err := xml.Unmarshal([]byte(pomContents), &pom)
  if err != nil {
    fmt.Print(err)
    return "", "", "", ""
  }

  versionExtractRegex, err := regexp.Compile(`\d+\.\d+`)
  if err != nil {
    fmt.Print(err)
    return "", "", "", ""
  }

  majorVersion = versionExtractRegex.FindString(pom.Version)
  artifactId = pom.ArtifactId
  mvnVersion = fmt.Sprintf("%s.%s", majorVersion, timestamp)
  groupId = pom.GroupId

  return majorVersion, artifactId, mvnVersion, groupId
}

func FormatBashSource(majorVersion string, artifactId string, mvnVersion string, groupId string) (string) {
  return fmt.Sprintf("export MAJOR_VERSION=%s; export ARTIFACT_ID=%s; export COMPONENT_VERSION=%s; export GROUP_ID=%s\n", majorVersion, artifactId, mvnVersion, groupId)
}

func writeContents(o string, contents string) {
  file, err := os.Create(o)
  if err != nil {
    fmt.Print(err)
    return
  }
  defer file.Close()

  contentReader := strings.NewReader(contents)
  multi := io.MultiWriter(os.Stdout, file)
  io.Copy(multi, contentReader)
}

