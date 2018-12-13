pipeline {
  agent {
    node {
      label '1'
    }

  }
  stages {
    stage('get code') {
      steps {
        git(url: 'https://github.com/zou2699/demoWeb.git', branch: 'master')
        sh 'go build -o app .'
      }
    }
  }
}