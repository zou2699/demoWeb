pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }

  }
  stages {
    stage('get code') {
      steps {
        git(url: 'https://github.com/zou2699/demoWeb.git', branch: 'master')
      }
    }
    stage('') {
      steps {
        sh 'go build -o app .'
      }
    }
  }
}