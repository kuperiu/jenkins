// @Library('wolox-ci') _

// node {

//   checkout scm

//   woloxCi('config.yml');
// }


pipeline {
  agent any
  stages {
    stage("Load") {
      steps {
        load("jenkins_test")
      }
    }
  }
}