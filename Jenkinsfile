pipeline {
    agent { docker { image 'golang:1.12' } }
    stages {
        stage('build') {
            steps {
                sh 'cd entryserver && go build'
            }
        }
    }
}