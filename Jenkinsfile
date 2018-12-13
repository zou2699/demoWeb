pipeline {
  agent {
    docker {
      image 'golang'
    }

  }
  stages {
    stage('get code') {
      steps {
        git(url: 'https://github.com/zou2699/demoWeb.git', branch: 'master')
      }
    }
    stage('build') {
      steps {
        sh 'go env'
        sh 'go get github.com/gin-gonic/gin'
        sh 'go build -o app .'
      }
    }
    stage('build image') {
      steps {
        sh '/usr/bin/docker build -t zouhl/webdemo .'
      }
    }
  }
}