package main

import (
  "testing"
  "fmt"
  "strings"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestVersion(t *testing.T) {

  t.Log("Given a pom content with version and artifactId, and a timestamp"); {
    pom := `<project>
            <version>1.0-SNAPSHOT</version>
            <artifactId>golangTest</artifactId>
            <groupId>com.go.mainTest</groupId>
          </project>`
    timestamp := "1234567890"

    majorVersion, artifactId, mvnVersion, groupId := Version(pom, timestamp)

    if majorVersion == "1.0" {
      t.Logf("\tShould receive a major version of %s - %s", majorVersion, checkMark)
    } else {
      t.Errorf("\tShould have received a major version of %s but got %s - %s", "1.0", majorVersion, ballotX)
    }

    if artifactId == "golangTest" {
      t.Logf("\tShould receive an artifactId of %s - %s", artifactId, checkMark)
    } else {
      t.Errorf("\tShould have received an artifactId of %s but was %s - %s", "golangTest", artifactId, ballotX)
    }

    if mvnVersion == "1.0.1234567890" {
      t.Logf("\tShould receive a mvnVersion of %s - %s", mvnVersion, checkMark)
    } else {
      t.Errorf("\tShould have received a mvnVersion of %s but was %s - %s", "1.0.1234567890", mvnVersion, ballotX)
    }

    if groupId == "com.go.mainTest" {
      t.Logf("\tShould receive a groupId of %s - %s", groupId, checkMark)
    } else {
      t.Errorf("\tShould have received a groupId of %s but was %s - %s", "com.go.mainTest", groupId, ballotX)
    }
  }
}

func ExampleVersion() {
  pom := `<project>
            <version>1.0-SNAPSHOT</version>
            <artifactId>golangTest</artifactId>
            <groupId>com.go.mainTest</groupId>
          </project>`
  timestamp := "1234567890"

  majorVersion, artifactId, mvnVersion, groupId := Version(pom, timestamp)

  fmt.Printf("%s - %s - %s - %s", majorVersion, artifactId, mvnVersion, groupId)
}

func TestFormatBashSource(t *testing.T) {

  t.Log("Given a simple set of majorVersion, artifactId, mvnVersion, and groupId"); {
    majorVersion := "mv"
    artifactId := "ai"
    mvnVersion := "mn"
    groupId := "gi"

    result := FormatBashSource(majorVersion, artifactId, mvnVersion, groupId)

    expectedMajorVersion := fmt.Sprintf("export MAJOR_VERSION=%s;", majorVersion)
    if strings.Contains(result, expectedMajorVersion) {
      t.Logf("\tShould have export '%s' - %s", expectedMajorVersion, checkMark)
    } else {
      t.Errorf("\tShould have had an export of '%s' but none found - %s", expectedMajorVersion, ballotX)
    }

    expectedArtifactId := fmt.Sprintf("export ARTIFACT_ID=%s;", artifactId)
    if strings.Contains(result, expectedArtifactId) {
      t.Logf("\tShould have export '%s' - %s", expectedArtifactId, checkMark)
    } else {
      t.Errorf("\tShould have had an export of '%s' but none found - %s", expectedArtifactId, ballotX)
    }

    expectedMvnVersion := fmt.Sprintf("export COMPONENT_VERSION=%s;", mvnVersion)
    if strings.Contains(result, expectedMvnVersion) {
      t.Logf("\tShould have export '%s' - %s", expectedMvnVersion, checkMark)
    } else {
      t.Errorf("\tShould have had an export of '%s' but none found - %s", expectedMvnVersion, ballotX)
    }

    expectedGroupId := fmt.Sprintf("export GROUP_ID=%s;", groupId)
    if strings.Contains(result, expectedGroupId) {
      t.Logf("\tShould have export '%s' - %s", expectedGroupId, checkMark)
    } else {
      t.Errorf("\tShould have had an export of '%s' but none found - %s", expectedGroupId, ballotX)
    }
  }
}



