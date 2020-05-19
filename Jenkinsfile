node {
  stage('CI') {
    doDynamicParallelSteps()
  }
}

def doDynamicParallelSteps(){
  tests = [:]
  for (f in ["Branch_1", "Branch_2", "Branch_3"]) {
    tests["${f}"] = {
      node {
        stage("${f}") {
          echo f
        }
      }
    }
  }
  parallel tests
}